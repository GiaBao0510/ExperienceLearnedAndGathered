using Microsoft.Extensions.DependencyInjection;
using System;
using System.Xml.Linq;
using Microsoft.Extensions.Options;
using Microsoft.Extensions.Configuration;
using Microsoft.Extensions.Configuration.Json;

namespace Buoi33
{
    //1.Inverted Dependency - Đảo ngược phụ phuộc
    interface IclassB
    {
        public void ActionB();
    }
    interface IclassC
    {
        public void ActionC();
    }
    interface IclassE
    {
        public void ActionE();
    }

    class LopB : IclassB
    {
        IclassC C_dependency;

        public LopB(IclassC _cDependency)
        {
            this.C_dependency = _cDependency;
        }
        public void ActionB()
        {
            Console.WriteLine("Action in Class B");
            C_dependency.ActionC();
        }
    }
    class LopA
    {
        IclassB B_dependency;

        public LopA(IclassB B_dependency)
        {
            this.B_dependency = B_dependency;
        }
        public void ActionA()
        {
            Console.WriteLine("Action in Class A");
            B_dependency.ActionB();
        }
    }
    class LopC : IclassC
    {
        public void ActionC() => Console.WriteLine("Action in class C.");
    }
        //Ví dụ thay đổi phụ thuộc
    class LopC1 : IclassC
    {
        public LopC1() => Console.WriteLine("ClassC1 is created");

        public void ActionC() => Console.WriteLine("Action in ClassC1");
    }
    class LopB1 : IclassB
    {
        IclassC C_dependency;
        public LopB1(IclassC c_dependency)
        {
            C_dependency = c_dependency;
            Console.WriteLine("ClassB1 is created");
        }

        public void ActionB()
        {
            Console.WriteLine("Acton in ClassB1");
            C_dependency.ActionC();
        }
    }
    class LopE : IclassE
    {
        public void ActionE() => Console.WriteLine("Action in class E");
    }

    class LopB2 : IclassB
    {
        IclassC c_deependency;
        string message;
        public LopB2(IclassC C_deependency, string tinnhan)
        {
            c_deependency = C_deependency;
            message = tinnhan;
            Console.WriteLine("Class B2 is created");
        }
        public void ActionB() {
            Console.WriteLine(message);
            c_deependency.ActionC();
        }
    }
        //Thiết kế sự phụ thuộc lỏng lẻo 
    class Horn
    {
        private int level;
        public Horn(int dolon) { this.level = dolon; }
        public void Beep() => Console.WriteLine("Beep - Beep - Beep ...");
    }

    class Car
    {
        public Horn horn { get; set; }
        public Car(Horn horn) { horn = horn; }

        public void Beep()
        {
            horn.Beep();
        }
    }

    //Lớp này chứa cách thiết lập cho lớp MyService
    public class MyServiceOptions
    {
        public string data1 { get; set; }
        public int data2 { set; get; }
    }

    public class MyService
    {
        public string data1 { get; set; }
        public int data2 { set; get; }

        public MyService(IOptions<MyServiceOptions> option)
        {
            var _options = option.Value;
            data1 = _options.data1;
            data2 = _options.data2;
        }

        public void PrintData() => Console.WriteLine($"Data1: {data1}, Data2: {data2}");
    }
    class progarm
    {
        //Tạo phương thức để đăng ký cho lớp B2
        public static IclassB createB2(IServiceProvider provider)
        {
            var b2 = new LopB2(
                        provider.GetService<IclassC>(),
                        ">>> Thuc hien trong class B2"
            );
            return b2;
        }

        static void Main(string[] args)
        {
            //Example: Inverted dependency
            IclassC objectC = new LopC1();
            IclassB objectB = new LopB1(objectC);
            LopA objectA = new LopA(objectB);

            objectA.ActionA();

            //Thiết kế sự phục thuộc lỏng léo
            Console.WriteLine("\n\t === Su phu thuoc long leo ===");
            Horn horn = new Horn(8);
            Car xe = new Car(horn);
            xe.horn = horn;
            xe.Beep();

            //3. Thư viện DependencyInjection
            Console.WriteLine("\n\t === Library Dependency Injection ===");
            var cacDichVu = new ServiceCollection();

            //3.1 Đăng ký các dịch vụ
            // Interface C có LopC và LopC1 dùng kế thừa đến
                //Singleton
            cacDichVu.AddSingleton<IclassC, LopC>(); //Tham số 1 là tên dịch vụ, tham số 2 là kiểu dịch vụ đó được tạo ra

            //Transient 
            cacDichVu.AddTransient<IclassB, LopB>();

            //scope
            cacDichVu.AddScoped<IclassE, LopE>();

            //3.2 Tạo ra đối tượng provider
            var nhaCungCap = cacDichVu.BuildServiceProvider();

            //3.3 Lấy ra đối tượng đã đăng ký trong collection

            Console.WriteLine("\nLifetime: Singleton");
            for(int i = 0; i <= 5; i++)
            {
                IclassC X1 = nhaCungCap.GetService<IclassC>();

                //Lấy mã băm của đối tượng GetHashCode(): nếu mã hash giống nhau thì cùng ột đối tượng, ngược lại thì khác
                Console.WriteLine(X1.GetHashCode());
            }

            Console.WriteLine("\nLifetime: Transient");
            for(int i =0; i <= 4; i++)
            {
                IclassB y1 = nhaCungCap.GetService<IclassB>();
                Console.WriteLine(y1.GetHashCode());
            }

            Console.WriteLine("\nLifetime: scope");
            for (int i = 0; i <= 4; i++)
            {
                IclassE z1 = nhaCungCap.GetService<IclassE>();
                Console.WriteLine(z1.GetHashCode());
            }

            //Tạo ra scope mới - Trong mỗi phạm vi truy cập thì có các phiên bản khác nhau
            Console.WriteLine("\n tao cac scope moi");
            using(var scope = nhaCungCap.CreateScope())
            {
                var nhaCungCap1 = scope.ServiceProvider;
                for (int i = 0; i <= 4; i++)
                {
                    IclassE z1 = nhaCungCap1.GetService<IclassE>();
                    Console.WriteLine(z1.GetHashCode());
                }
            }

            /*
             * Bài ví dụ: những dịch vụ được đăng ký vào sẽ tự động được tạo ra. Và khi tạo ra dịch vụ đó,
             * Nếu dịch vụ cần dependency thì các dependency sẽ tự động tạo ra và tự động inject vào đối tượng 
             */
            Console.WriteLine("\nVi du tu dong tao va Inject Dependency vao dich vu");
            //Đăng ký dịch vụ singleton
            var cacDichVu_v2 = new ServiceCollection();
            
            cacDichVu_v2.AddSingleton<LopA, LopA>();
            cacDichVu_v2.AddSingleton<IclassB>(createB2);
            cacDichVu_v2.AddSingleton<IclassC, LopC1>();

            var NhaCungCap_v2 = cacDichVu_v2.BuildServiceProvider();

            LopA A1 = NhaCungCap_v2.GetService<LopA>();
            A1.ActionA();

                //Sử dụng option khởi tạo dịch vụ trong DI
            Console.WriteLine("\n\t Su dung option khoi tao dich vu trong DI");

            //Đăng ký dịch vụ
            var cacDichVu_v3 = new ServiceCollection();
            cacDichVu_v3.AddSingleton<MyService>();
            cacDichVu_v3.Configure<MyServiceOptions>(
                (MyServiceOptions opt) =>
                {
                    opt.data1 = "Hello moi nguoi";
                    opt.data2 = 2024;
                }    
            );

            var provider_v3 = cacDichVu_v3.BuildServiceProvider();
            var myservice = provider_v3.GetService<MyService>(); //Lấy dịch vụ
            myservice.PrintData();

            //Sử dụng cấu hình từ file cho DI Container
            Console.WriteLine("\n\t lay du lieu tu file cau hinh cho DI container");
            /*
             * Lớp ConfigurationBuilder, giúp nạp các cấu hình lưu trong file config, từ đó build ra đối tượng
             * ConfigurationRoot , đối tượng nảy truy cập đến các cấu hình bằng chỉ toán tử chỉ số [key]
            */
            IConfigurationRoot cauHinhGoc;

            /*
            Đối tượng này có khả năng đọc file cấu hình   
            */
            ConfigurationBuilder cauHinhXayDung = new ConfigurationBuilder();
            cauHinhXayDung.SetBasePath("D:/HocTap/Experience/hocC#/Buoi33_DependencyInjection/Buoi33_DependencyInjection");
            cauHinhXayDung.AddJsonFile("cauhinh.json");

            cauHinhGoc = cauHinhXayDung.Build();

            //Để đọc giá trị trong file cấu hình thì dùng phương thức GetSection   
            var data1 = cauHinhGoc.GetSection("MyServiceOptions").GetSection("data1");
            Console.WriteLine("mot phan du lieu duoc doc tu file json: "+data1.Value);

            var cacDichVu_v4 = new ServiceCollection();
            cacDichVu_v4.AddSingleton<MyService>();

            var sectionMyServiceOptions = cauHinhGoc.GetSection("MyServiceOptions");
            cacDichVu_v4.Configure<MyServiceOptions>(sectionMyServiceOptions);

            var provider_v4 = cacDichVu_v4.BuildServiceProvider();
            provider_v4.GetService<MyService>().PrintData();



        }
    }
}

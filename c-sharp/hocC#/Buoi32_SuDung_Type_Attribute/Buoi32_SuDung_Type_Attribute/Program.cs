using System;
using System.ComponentModel.DataAnnotations;
using System.Reflection;
namespace Buoi32
{
    /*
        Ví dụ tạo Attribute riêng như sau:
        Tạo Attribute MoTa:
         - ThongTinChiTiet
    */
    [AttributeUsage(AttributeTargets.Class | AttributeTargets.Property | AttributeTargets.Method)] //Mô tả thuộc tính được sử dụng ở đâu (Lớp. property, methob, ...)
    class MoTaAttribute : Attribute
    {
        public string ThongTinChiTiet { get; set; }

        public MoTaAttribute(string _thongTinChiTiet)
        {
            this.ThongTinChiTiet = _thongTinChiTiet;
        }

    }

    [MoTa("Lop chua thong tin chi tiet tren he thong")]
    class User
    {
        [MoTa("Ten nguoi dung:")]
        [Required(ErrorMessage ="Phai dat ten cho nguoi dung")]
        [StringLength(50, MinimumLength =3, ErrorMessage ="Chu y: Ten phai co do dai tu 3 - 50 ky tu")]
        public string Name { get; set; }

        [MoTa("Tuoi nguoi dung:")]
        [Range(18,90)]
        public int Age { get; set; }

        [MoTa("So dien thoai nguoi dung:")]
        [Phone]
        public string PhoneNumber { get; set; }

        [MoTa("Email nguoi dung:")]
        [EmailAddress(ErrorMessage ="Dia chi email khong hop le")]
        public string Email { get; set; }

        public User(string hoten, int tuoi, string sdt, string email)
        {
            this.Name = hoten;
            this.Age = tuoi;
            this.PhoneNumber = sdt;
            this.Email = email;
        }

        public string ToString()
        {
            return $"Name:{Name} ,Age:{Age} , PhoneNumber:{PhoneNumber}, Email: {Email}";
        }

        //Đánh dấu phương thức lỗi thời bằng 
        [Obsolete("! Phuong thuc nay da loi thoi .Vui long dung phuong thuc khac")]
        public void PrintUsserName() => Console.WriteLine($"Name: {Name}");
    }

    class program
    {
        static void Main(string[] args) {
            int X1 = 100;
            int[] nums = { 1, 4, 7, 8, 5, 3 };


            // --------- 1.Type ------------
            Type type1 = typeof(int),
                type2 = typeof(double);
            var type3 = typeof(Array);

            //Phương thức GetType để lấy thông tin kiểu dữ liệu của một biến
            var thongtin = nums.GetType();

            //Phương thức FullName dùng để lấy tên kiểu dữ liệu mà nó tiếp nhận
            Console.WriteLine($"Ten kieu du lieu: {thongtin.FullName}");

            //Phương thức GetProgeties(): dùng để trả một mảng những thuộc tính của một mảng nào đó
            Console.WriteLine($"\n\tPhuong thuc GetProgeties");
            thongtin.GetProperties().ToList().ForEach(
                (PropertyInfo e) =>
                {
                    Console.WriteLine(e.Name);
                }
            );

            //Phương thức GetFields() dùng để đọc thông tin về các trường dữ liệu
            Console.WriteLine($"\n\tPhuong thuc GetFields");
            thongtin.GetFields().ToList().ForEach(
                (FieldInfo e) =>
                {
                    Console.WriteLine(e.Name);
                }
            );

            //Phương thức GetMethods() dùng để lấy thông tin về các phương thức có trong dữ liệu kiểu mảng
            Console.WriteLine($"\n\tPhuong thuc GetMethods");
            thongtin.GetMethods().ToList().ForEach(
                (MethodInfo e) =>
                {
                    Console.WriteLine(e.Name);
                }
            );

            //ví dụ lấy thông tin thuộc tính, tên thuộc tính, phương thức có trong lớp nào đó
            Console.WriteLine("\n\t ===== Thuc hanh type ====");
            User nguoiDung1 = new User("Nguyen Van A", 18, "nvA@gmail.com", "01239267068");
            nguoiDung1.GetType().GetProperties().ToList().ForEach(
                (p) =>
                {
                    var values = p.GetValue(nguoiDung1);
                    Console.WriteLine($"{p.Name}: {values}");
                }
            );
            nguoiDung1.PrintUsserName();

            // --------- 2.Attribute ------------
            Console.WriteLine("\n\t 2.Attribute");
            //Đọc các thuộc tính của lớp User
            nguoiDung1.GetType().GetProperties().ToList().ForEach(
                (PropertyInfo property) =>
                {
                    foreach(var attr in property.GetCustomAttributes(false))
                    {
                        MoTaAttribute mota = attr as MoTaAttribute;
                        if( mota != null)
                        {
                            var value = property.GetValue(nguoiDung1);
                            var name = property.Name;
                            Console.WriteLine($"{name} - {mota.ThongTinChiTiet} {value}");
                            Console.WriteLine($"{name} - {mota.ThongTinChiTiet} {value}");
                        }
                    }
                }    
            );

            Console.WriteLine("\nCac vi du ve Attribute:");
            ValidationContext context = new ValidationContext(nguoiDung1);

            var result = new List<ValidationResult>();

            //Kiểm tra
            var kq = Validator.TryValidateObject(nguoiDung1, context, result, true);
            Console.WriteLine(kq);
            if(kq == false)
            {
                result.ToList().ForEach(
                    (er)=> {
                        Console.WriteLine(er.MemberNames.First());
                        Console.WriteLine(er.ErrorMessage);
                    }
                );
            }
        }
    }
}

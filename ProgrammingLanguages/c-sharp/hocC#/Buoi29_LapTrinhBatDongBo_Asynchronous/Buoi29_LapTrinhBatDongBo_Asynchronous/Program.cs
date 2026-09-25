using System;

namespace buoi29
{
    class program
    {
        //Phương thức này sẽ thi hàng trong một khoảng thời gian được quy định sẵn
        static void DoSomeThing(int seconds, string messege, ConsoleColor color)
        {
            //1.Khóa đối tượng 
            lock (Console.Out)
            {
                Console.ForegroundColor = color;
                Console.WriteLine($"{messege ,10} ... Start");
                Console.ResetColor();
            };

            //Hàm Sleep này sẽ dừng hoạt động trong dòng bao nhiêu mili giây
            for (int i =1; i<= seconds; i++)
            {
                lock (Console.Out)
                {
                    Console.ForegroundColor = color;
                    Console.WriteLine($"{messege} - {i}");
                    Console.ResetColor();
                }
                Thread.Sleep(500);
            }

            lock (Console.Out)
            {
                Console.ForegroundColor = color;
                Console.WriteLine($"{messege,10} ... End");
                Console.ResetColor();
            };
        }
        
        //Hàm trả về Task
        static Task GardeningDuties()
        {
            Task t3 = new Task(
                () =>
                {
                    DoSomeThing(8, "Gardening Duties", ConsoleColor.White);
                }
            );
            t3.Start();
            return t3;
        }

        static Task CleaningDuties()
        {
            Task x = new Task(
                () =>
                {
                    DoSomeThing(10, "Cleaning duty", ConsoleColor.Magenta);
                }
            );
            x.Start();
            return x;
        }

        //Hàm bất đồng bộ - void
        static Task MoppingTask()
        {
            Task x = new Task(
                () =>
                {
                    DoSomeThing(9, "Mopping Task", ConsoleColor.Cyan);
                }
            );
            x.Start();
            x.Wait();
            return x;
        }
        static Task LawnMowingDuty()
        {
            Task x = new Task(
                () =>
                {
                    DoSomeThing(7, "Lawn Mowing Duty", ConsoleColor.DarkRed);
                }
            );
            x.Start();
            x.Wait();
            return x;
        }

        //Hàm bất đồng bộ - return type
        static Task<string> paymentTask()
        {
            Task<string> X = new Task<string>(
                (object obj) =>
                {
                    string str = (string)obj;
                    DoSomeThing(6, str, ConsoleColor.Magenta);
                    return "Da hoan thanh viec thanh toan";
                }
            ,"Payment task.");
            X.Start();
            return X;
        }

        // Async - await : task void
        static async Task DishwashingDuty()
        {
            Task x = new Task(
                () =>
                {
                    DoSomeThing(5, "Dishwashing Duty", ConsoleColor.Red);
                }   
            );
            x.Start();
            await x; //Tường đương với việc x.Wait(); .Nhưng được cái la nó sẽ trả về luông mà không cần thêm lệnh return
        }

        static async Task GarbageCollection()
        {
            Task x = new Task(
                () =>
                {
                    DoSomeThing(3, "Garbage Collection", ConsoleColor.DarkBlue);
                }
            );
            x.Start();
            await x;
        }

        // Async - await : task return type
        static async Task<string> WindowCleaningDuty()
        {
            Task<string> X = new Task<string>(
                () =>
                {
                    DoSomeThing(4, "Window Cleaning Duty", ConsoleColor.Gray);
                    return "Da hoan thanh viec lau cua so";
                }    
            );
            X.Start();
            var ketQua = await X;
            return ketQua;
        }

        //Ví dụ phương thức bất đồng bộ tải về 1 trang web
        static async Task<string> GetWeb(string url)
        {
            /*
             + HttpClient: Lớp này dùng để tải 1 trang web.
             + GetAsync: Đây là phương thức bất đồng bộ dùng để tải nội dung của một trang web.
             + HttpResponseMessage: Lớp này dùng để nhận kết quả từ việc tải nội dung trên phương thức GetAsync 
             + Content: để lấy hoặc đặt 1 đoạn tin nhắn phản hồi của HTTP
             + ReadAsStringAsync: Đây là một phương thức bất đồng bộ để đọc nội dung trả về từ trang web
            */
            Console.WriteLine("\n\tBat dau tai noi dung");
            HttpClient httpClient = new HttpClient();   
            HttpResponseMessage ketQua = await httpClient.GetAsync(url);
            Console.WriteLine("\n\tBat dau doc noi dung");
            string content = await ketQua.Content.ReadAsStringAsync();
            Console.WriteLine("\n\tTra ve noi dung");
            return content;
        }

        static async Task<string> CobwwebCleaningMission()
        {
            Task<string> X = new Task<string>(
             (object obj) =>
             {
                 string str = (string)obj;
                 DoSomeThing(4, str, ConsoleColor.Magenta);
                 return "Da hoan thanh viec quet don mang nhen";
             }   
            , "Cobwweb Cleaning Mission");
            X.Start();
            string ketqua = await X;
            return ketqua;
        }

        static async Task Main(string[] args) {

            //1. Lập trình Đồng bộ - Khóa đối tượng
            Console.WriteLine("\n\t=== Lock - Synchronous ===");
            DoSomeThing(5, "Task1",ConsoleColor.Green);
            DoSomeThing(6, "Task2", ConsoleColor.Red);
            DoSomeThing(4, "Task3", ConsoleColor.Blue);

            //2. Tạo tác vụ Task
            Console.WriteLine("\n\t=== Task ===");
            //Khởi tạo tác vụ
            Task T2 = new Task(() =>
            {
                DoSomeThing(8, "Task4", ConsoleColor.Yellow);
            });

            Task T3 = new Task(
                (Object obj) =>
                {
                    //Chuyển đối đối tượng về chuỗi
                    string tacvu5 = (String)obj;
                    DoSomeThing(3, tacvu5, ConsoleColor.White);
                }, "Task5"    
            );

            //-- Chạy tác vụ. Các tác vụ chạy trên các luồng khác nhau .Nên xảy ra chạy // với nhau
            //T2.Start();
            //T3.Start();
            //Task.WaitAll(T2, T3);
            //DoSomeThing(4, "Task6", ConsoleColor.Cyan);

            //Task lamVuon = GardeningDuties();
            //Task quetNha = CleaningDuties();
            Task LauNha = MoppingTask();
            Task CatCo = LawnMowingDuty();
            Task NhatRac= GarbageCollection();
            Task RuaChen = DishwashingDuty();

            await NhatRac;
            await RuaChen;

            // Async - await có trẻ về kiểu dữ liệu
            Task<string> GiatDo = new Task<string>(
                () =>
                {
                    DoSomeThing(6, "Laundry duty", ConsoleColor.Blue);
                    return "Rua chen xong";
                }    
            );

            Task<string> NauAn = new Task<string>(
                (object obj) =>
                {
                    string str = (string)obj;
                    DoSomeThing(5, str, ConsoleColor.Yellow);
                    return $"Da nau an xong - {str}";
                }
            ,"Cooking mission");

            GiatDo.Start();
            NauAn.Start();
            Task.WaitAll(NauAn, GiatDo);
            
            //Lấy gia trị trả về của hàm bất đồng bộ
            var kq1 = NauAn.Result; 
            var kq2 = GiatDo.Result;
            Console.WriteLine($"Giat do: {kq2} - Nau an: {kq1}");

            //Hàm Task có kiểu trả về
            Task<string> thanhToan = paymentTask();
            Task.WaitAll(thanhToan);
            var kq3 = thanhToan.Result;
            Console.WriteLine($"Thanh toan: {kq3}");

            //Hàm task async - await có kiểu trả về
            Task<string> DonMangNhen = CobwwebCleaningMission();
            Task<string> lauCuaSo = WindowCleaningDuty();

            var kq4 = await DonMangNhen;
            var kq5 = await lauCuaSo;
            Console.WriteLine($"Quet don mang nhen: {DonMangNhen} - Lau cua so: {lauCuaSo}");

            //Ví dụ thực tế
            var DocNoiDung = GetWeb("https://huyminhcantho.com/");
            var noidung = await DocNoiDung;
            Console.WriteLine(noidung);

            Console.WriteLine("End Program");
            Console.ReadKey();  //Sẽ dừng lại khi người dùng bấm phím bất kỳ
        }
    }
}

using System;

namespace Buoi2
{
    class Program
    {
        static void Main(String[] args) {

            //Biến (Khai báo tường minh)
            string mssv, hovaten, strinput;
                mssv = "b2016947";  //Gán giá trị
            int age = 22;
            char gender;
            float diemso = 9.8f, a, b;
                diemso = (float)10.0; //Ép kiểu dữ liệu
            double pi = 3.14;
            object conan;   //Kiểu dữ liệu đối tượng có thể có nhận bất cứ giá trị nào
                conan = diemso;

            //Biến (Khai báo ngầm định)
            var obj1 = "sj";
            var obj2 = 2;
            var obj3 = 3.14;

            //khai báo Hằng số
            const double PI = 3.14;

            //2. THiết lập màu chữ
            Console.ForegroundColor = ConsoleColor.Black;

            //3. Thiết lập màu nền
            Console.BackgroundColor = ConsoleColor.Yellow;

            //4.Đặt tiêu đề cửa số
            Console.Title = "Buoi thu 2";

            //5.Hiển thị
            Console.WriteLine("Student code: {0}", mssv);
            Console.WriteLine("Age: {0}", age);

            //6.Nhập dữ liệu
            Console.Write("Enter student code: ");
            mssv = Console.ReadLine();

            Console.Write("Enter full name: ");
            hovaten = Console.ReadLine();

            Console.Write("Enter parameter A: ");
            strinput = Console.ReadLine();
            a = float.Parse(strinput);

            Console.Write("Enter parameter B: ");
            strinput = Console.ReadLine();
            b = Convert.ToSingle(strinput);

            Console.WriteLine("Full name: {0} - StudentCode: {1}",hovaten, mssv);
            Console.WriteLine("A + B = {0}", (a + b));

            //7.Thiết lập màu nền, màu chữ về mặc định
            Console.ResetColor();

            Console.WriteLine();

            //8.Bắt sự kiện khi người dùng nhấn vào bất kỳ phím nào
            Console.WriteLine("Enter any key to 'Exit'.");
            Console.ReadKey();
        } 
    }
}

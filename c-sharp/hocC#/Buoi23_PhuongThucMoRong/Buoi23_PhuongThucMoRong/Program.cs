using MyLib;
using System;
using System.Linq;

namespace Buoi23
{
    //Lớp tính, thì bên trong phải là phương thức tĩnh
    static class MyLbrary
    {
        //Thêm từ khóa this và tham số để cho biết phương thức này là mở rộng
        public static void HienThi(this string s, ConsoleColor color)
        {
            Console.ForegroundColor = color;
            Console.WriteLine(s);
            Console.ResetColor();
        }
    }

    class Program
    {
        

        static void Main(string[] args)
        {
            string XinChao = "Hello World!";
            XinChao.HienThi(ConsoleColor.Red);
            "Hello All".HienThi(ConsoleColor.Yellow);

            double X1 = 17.2;
            Console.WriteLine($"Tinh binh phuong: {X1.BinhPhuong()}");
            Console.WriteLine($"Tinh Can Bac Hai: {X1.CanBacHai()}");
            Console.WriteLine($"Tinh Sin: {X1.Sin()}");
            Console.WriteLine($"Tinh Cos: {X1.Cos()}");
        }
    }
}

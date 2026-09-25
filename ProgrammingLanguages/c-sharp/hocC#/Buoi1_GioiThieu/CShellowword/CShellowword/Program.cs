using System;

namespace CShelloword
{
    class Program
    {
        static void Main(string[] args) {
            Console.WriteLine("Hello ,World!");
            System.Console.WriteLine("Xin chào");

            Console.WriteLine(Tong(5,10));
        }
        
        /// <summary>
        /// Trả về tổng 2 số nguyên
        /// </summary>
        /// <param name="a">số nguyên 1</param>
        /// <param name="b">Số nguyên 2</param>
        /// <returns>Tính tổng (a + b)</returns>
        static int Tong(int a, int b)
        {
            return a + b;
        }
    }
}
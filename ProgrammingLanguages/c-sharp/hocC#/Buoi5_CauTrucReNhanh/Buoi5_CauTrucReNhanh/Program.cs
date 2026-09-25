using System;

namespace Buoi5
{
    class program
    {
        static void Main(string[] args)
        {
            
            int a, b, c, max;
            Console.Write("Enter a integer A: ");
            a = int.Parse(Console.ReadLine());
            Console.Write("Enter a integer B: ");
            b = int.Parse(Console.ReadLine());
            Console.Write("Enter a integer C: ");
            c = int.Parse(Console.ReadLine());


            //Ví dụ kiểm tra chẳn lẻ
            if (a % 2 == 0)
            {
                Console.WriteLine("-> {0} is an even number", a);
            }
            else
            {
                Console.WriteLine("-> {0} is an old number", a);
            }

            //Ví dụ tìm số lớn nhất 3 ngôi
            max = (a > b) ? (a > c) ? a : c : (b > c) ? b : c;

            Console.WriteLine("Langest nunber: {0}",max);

            //Ví dụ kiểm tra loại học lực
            Console.WriteLine();
            float dtb;
            Console.Write("Enter your aveeage score: ");
            dtb = float.Parse(Console.ReadLine());

            if (dtb >= 8.0 && dtb <= 10)
            {
                Console.WriteLine("Well-studing");
            }
            else if (dtb >= 6.5 && dtb < 8)
            {
                Console.WriteLine("Academic pretty");
            }
            else if (dtb >= 4.5 && dtb < 6.5)
            {
                Console.WriteLine("Learning capacity is average");
            }
            else
            {
                Console.WriteLine("Weak learning capacity");
            }
        }
    }
}

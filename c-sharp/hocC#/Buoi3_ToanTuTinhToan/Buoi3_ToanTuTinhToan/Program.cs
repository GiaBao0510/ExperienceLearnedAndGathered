using System;

namespace Buoi3
{
    class Program
    {
        static void Main(string[] args)
        {
            float A = 50f,
                B = 15f;
            float kq = A + B * (A / B) % A;
            
            //Toán tử gán
            int x = 10;
            x += 15;

            Console.WriteLine(" {0} + {1} = {2} ",A,B,(A+B));
            Console.WriteLine(" {0} - {1} = {2} ", A, B, (A - B));
            Console.WriteLine(" {0} * {1} = {2} ", A, B, (A * B));
            Console.WriteLine(" {0} / {1} = {2} ", A, B, (A / B));
            Console.WriteLine(" {0} % {1} = {2} ", A, B, (A % B));
            Console.WriteLine(" {0} + {1} = {2} ", A, B, (A + B));
            Console.WriteLine("Result: {0}", kq);
            Console.WriteLine("X: {0}", x);

            //Toán tử tăng giảm
            x++;
            Console.WriteLine("X++: {0}", x);

            ++x;
            Console.WriteLine("++X: {0}", x);

            x--;
            Console.WriteLine("X--: {0}", x);

            --x;
            Console.WriteLine("--X: {0}", x);
        }
    }
}
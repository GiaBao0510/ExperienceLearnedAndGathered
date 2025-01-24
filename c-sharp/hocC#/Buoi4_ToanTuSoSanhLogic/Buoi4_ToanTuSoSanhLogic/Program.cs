using System;

namespace Buoi4
{
    class progarm
    {
        static void Main(string[] args)
        {
            int a = 15, b = 18;
            bool dung = true, sai = false;

            //Hiển thị kết quả so sánh
            Console.WriteLine("{0} > {1} : {2}",a,b,(a>b));
            Console.WriteLine("{0} >= {1} : {2}", a, b, (a >= b));
            Console.WriteLine("{0} < {1} : {2}", a, b, (a < b));
            Console.WriteLine("{0} <= {1} : {2}", a, b, (a <= b));
            Console.WriteLine("{0} == {1} : {2}", a, b, (a == b));
            Console.WriteLine("{0} != {1} : {2}", a, b, (a != b));

            //Hiển thị kết quả login
            Console.WriteLine();
            Console.WriteLine("{0} && {1} : {2}", dung, sai, (dung && sai));
            Console.WriteLine("{0} || {1} : {2}", dung, sai, (dung || sai));
            Console.WriteLine("!{0} : {1}", dung, !dung);
        }

    }
}

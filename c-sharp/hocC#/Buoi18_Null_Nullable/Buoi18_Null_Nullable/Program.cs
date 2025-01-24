using System;
using System.Net.Security;

namespace buoi18
{
    class X1
    {
        public void XinChao() => Console.WriteLine("Hello world");
    }

    class program
    {
        static void Main(string[] args)
        {
            X1 x = null;
            x?.XinChao();

            //nullable
            int? H = null;
            //H = 100;
            Console.WriteLine($"Co gia tri: {H.HasValue} - H: {H}");
        }
    }
}
 
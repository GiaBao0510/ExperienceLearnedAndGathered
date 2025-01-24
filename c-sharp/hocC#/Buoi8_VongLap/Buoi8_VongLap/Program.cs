using System;

namespace buoi7
{
    class program
    {
        static void Main(string[] args)
        {
            //Example for
            int n, m, k;
            Console.Write("For: enter the number of repeat: ");
            n = int.Parse(Console.ReadLine());
            for(int i =1; i <= n; i++)
            {
                Console.WriteLine("Count {0}",i);
            }

            Console.Write("While: enter the number of repeat: ");
            m = int.Parse(Console.ReadLine());
            int dem = 1;
            while(dem <= m)
            {
                Console.WriteLine("{0} * {1} = {2}",m,dem,(dem*m));
                dem++;
            }

            Console.Write("do...while: enter the number of repeat: ");
            dem = 1;
            k = int.Parse(Console.ReadLine());
            do
            {
                Console.WriteLine("{0} * {1} = {2}", k, dem, (dem * k));
                dem++;
            } while (dem <= k);
        }
    }
}

using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;

namespace MyNameSpace
{
    internal class MyNameSpace
    {
        public static void Hello()
        {
            Console.WriteLine("Hello All!");
        }
    }

    //Khai báo namespace con
    namespace buoi15_child
    {
        class program
        {
            public static void Hello()
            {
                Console.WriteLine("Hello VietNam!");
            }
        }
    }
}

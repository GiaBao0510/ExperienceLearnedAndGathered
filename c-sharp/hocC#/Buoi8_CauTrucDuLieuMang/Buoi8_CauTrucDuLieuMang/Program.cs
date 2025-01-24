using System;
using System.Linq;  //Mở rộng lấy thêm phương thức trong mảng

namespace buoi8
{
    class program
    {
        static void Main(string[] args)
        {

           // ### 1. Introducing arrays

            int soLuong;
            Console.Write("Enter the number of names: ");
            soLuong = int.Parse(Console.ReadLine());

            //For example, the list data type is String
            string[] ds;
            int[] soNguyen = { 100, 2, 52, 76, 9, 15, 1, 8}; //assigned upon initialization
            ds = new string[soLuong];

            //Enter each name
            for(int i = 0; i< soLuong; i++)
            {
                Console.WriteLine(); //Space
                Console.Write("Enter name {0}: ", i);
                ds[i] = Console.ReadLine();
            }

            //show
            Console.WriteLine("[{0}]", string.Join(", ", ds));

           // ### 2. array methods

            //Show array length
            Console.WriteLine("Length: {0}", ds.Length);

            //Number of dimensions in the array
            Console.WriteLine("Rank: {0}", ds.Rank);

            //Get minimum value
            Console.WriteLine("Min value: {0}", soNguyen.Min());

            //Get Max value
            Console.WriteLine("Min value: {0}", soNguyen.Max());

            //Sum of elements
            Console.WriteLine("Sum of elements: {0}", soNguyen.Sum());

            //Sort up ascending
            Array.Sort(soNguyen);
            Console.WriteLine("[{0}]", string.Join(", ",soNguyen));


            // ### 3. Create Matrix
            int[,] members = new int[3,3] { 
                {1, 2, 3 }, 
                { 4, 5, 6} ,
                { 7, 8, 9}
            };

            //Show
            for(int i = 0; i<3;i++)
            {

                Console.WriteLine();

                for (int j = 0; j<3; j++)
                {
                    Console.Write(members[i, j]);
                    Console.Write(" ");
                }
            }
        }
    }
}

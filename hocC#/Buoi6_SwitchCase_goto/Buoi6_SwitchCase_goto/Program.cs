using System;

namespace Buoi6
{
    class program
    {
        static void Main(string[] args)
        {
            // Example switch case
            int thu;

            L1: //Đặt nhãn, nơi được để điều hướng
            Console.Write("\nEnter the day of the week: ");
            thu = int.Parse(Console.ReadLine());

            switch (thu)
            {
                case 2:
                    Console.WriteLine("Monday");
                break;

                case 3:
                    Console.WriteLine("Tuesday");
                    break;

                case 4:
                    Console.WriteLine("Wednesday");
                    break;

                case 5:
                    Console.WriteLine("Thursday");
                    break;

                case 6:
                    Console.WriteLine("Friday");
                    break;

                case 7:
                    Console.WriteLine("Saturday");
                    break;

                case 8:
                    Console.WriteLine("Sunday");
                    break;

                default:
                    Console.WriteLine("Not Found");
                    goto L1; //Nếu trường hợp không tìm thấy thì điều hướng về L1 , để nhập lại
                break ;

            }

            //Exmaple GoTo



            //waiting
            Console.WriteLine("Enter to Exit");
            Console.Read();
        }
    }
}

using System;

namespace buoi20
{
    public delegate void ShowLog(string message);

    class program {

        static void news(string x)
        {
            Console.ForegroundColor = ConsoleColor.Green;
            Console.WriteLine(x);
        }

        static void warning(string x)
        {
            Console.ForegroundColor = ConsoleColor.Red;
            Console.WriteLine(x);
        }

        static public int Sum(int a, int b) => a + b;

        static public int Hieu(int a, int b) => a - b;

        static void Main(string[] args) {

            //1. Delegate
            //Biến có kiểu delegate
            ShowLog doanTinNhan1 = null; //Có thể nhận giá trị null       
            
            doanTinNhan1 += news;


            //Dùng toán tử += để cho biến này có thể thi hành nhiều phương thức cùng lúc
            doanTinNhan1 += warning;
            doanTinNhan1 += news;
            doanTinNhan1 += news;
            doanTinNhan1 += news;
            doanTinNhan1 += warning;

            doanTinNhan1.Invoke("Tim tao dau qua man");

            //2. Action Delegate  
            Action hanhdong1;               //delegate void()
            Action<string, int> hanhdong2;  //delegate void(string a, int b)
            Action<string> hanhdong3;
            
            hanhdong3 = news;
            hanhdong3?.Invoke("VaiDan");

            //3. Func
            Func<int, int, int> action1;    //delegate int Ham1(int x, int y)
            action1 = Sum;

            Console.WriteLine($"Tong: {action1(15,84)}");
        }
    }

}

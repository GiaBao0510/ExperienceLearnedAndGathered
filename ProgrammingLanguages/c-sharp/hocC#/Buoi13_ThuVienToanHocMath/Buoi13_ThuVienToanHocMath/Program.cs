using System;

namespace buoi13
{
    class program
    {
        static void Main(string[] args)
        {
            int X = 5, X1 = 25;
            float Y = 6.75f;

            Console.WriteLine($"Pi: {Math.PI}");
            Console.WriteLine($"E: {Math.E}");
            Console.WriteLine($"The largest number between {8} and {19} is {Math.Max(8, 19)}");
            Console.WriteLine($"The smallest number between {8} and {19} is {Math.Min(8, 19)}");

            //Lấy trị tuyệt đối
            Console.WriteLine($"Absolute value {-15} is {Math.Abs(-15)}");

            //Hàm sign trả về 0 (Nếu số đó == 0), 1(Nếu số đó > 0) ,-1(Nếu số đó < 0)
            Console.WriteLine($"Sign of {0} is {Math.Sign(0)} ");
            Console.WriteLine($"Sign of {-18} is {Math.Sign(-18)} ");
            Console.WriteLine($"Sign of {8} is {Math.Sign(8)} ");

            //Lượng giác sin, cos, tan, Asin, Acos, Atan
            double corner = 60;
            double rad = Math.PI * corner / 180;
             
            Console.WriteLine($"Sin({rad}) = {Math.Sin(rad / 2)}");
            Console.WriteLine($"Cos({rad}) = {Math.Cos(rad / 2)}");
            Console.WriteLine($"Tan({rad}) = {Math.Tan(rad / 2)}");
            Console.WriteLine($"ASin({rad}) = {Math.Asin(rad / 2)}");
            Console.WriteLine($"ACos({rad}) = {Math.Acos(rad / 2)}");
            Console.WriteLine($"Âtn({rad}) = {Math.Atan(rad / 2)}");

            //sqrt: tính lũy thừa
            Console.WriteLine($"Sqrt({X1}) = {Math.Sqrt(X1)}");

            //pow: tính mũ
            Console.WriteLine($"pow({X},3) = {Math.Pow(X,3)}");

            //log(x): tính logarit cơ số 10
            Console.WriteLine($"log({X}) = {Math.Log(X)}");

            //log2(x): tính logarit cơ số 2
            Console.WriteLine($"log2({X}) = {Math.Log2(X)}");

            //Làm tròn về số nguyên gần nhất
            Console.WriteLine($"Round({Y}) = {Math.Round(Y)}");

            //làm tròn lên
            Console.WriteLine($"Ceilling({Y}) = {Math.Ceiling(Y)}");

            //làm tròn xuống
            Console.WriteLine($"Floor({Y}) = {Math.Floor(Y)}");

            //Cắt đi phần thập phân của 1 số
            Console.WriteLine($"Truncate({Y}) = {Math.Truncate(Y)}");
        }
    }
}

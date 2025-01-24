using System;

namespace buoi9
{
    class program
    {
        //1.Tạo phương thức hello
        public static void hello()
        {
            Console.WriteLine("Hello World!");
        }

        //2.Tạo phương thức tính tích
        public static int ProductOf2Integers(int a, int b)
        {
            return a * b;
        }

        //3. Ví dụ truyền tham số
        public static void FullName(string lastname, string firstname)
        {
            Console.WriteLine($"Xin chào {lastname+" "+firstname}");
        }

        //4. truyền tham trị
        public static void binhphuong(int x)
        {
            x = x * x;
            Console.WriteLine("The square is {0}", x);
;        }

        //5. truyền tham chiếu
        public static void binhphuong2(ref int x)
        {
            x = x * x;
            Console.WriteLine("The square is {0}", x);
            ;
        }

        //6. Tạo 1 lớp count để thực hiện động tác truyền tham chiếu
        class Count
        {
            public int c = 0;
        }

        //6.1 Tạo phương thức dùng để tăng giá trị của thuộc tính c trong lớp count
        static void dem(Count count)
        {
            count.c++;
        }

        //Main
        static void Main(string[] args) {
            hello();
            buoi9.program.hello();//Viết một cách đầy đủ
            var result = buoi9.TinhToan.SumOf2Integers(5, 19);
            float x = 25.5f, y = 8.15f;

            FullName("Pham", "Gia Bao");
            FullName(lastname: "Nguyen", firstname: "Van A");

            Console.WriteLine("===================");
            Console.WriteLine("Product of 2 integers: {0}", ProductOf2Integers(5, 8));
            Console.WriteLine("Sum of 2 integers: {0}", result);
            Console.WriteLine("Sum of 2 real numbers: {0}", TinhToan.SumOf2Integers(x,y));

            Console.WriteLine("========Tham chiếu, tham trị==========");
            int a = 2, b = 3;
            binhphuong(a);
            Console.WriteLine(a);
            binhphuong2(ref b);
            Console.WriteLine(b);

            Console.WriteLine("===================");
            Count C = new Count();
            Console.WriteLine("Khoi tao: {0}", C.c);
            dem(C);
            Console.WriteLine("Da dem: {0}", C.c);

        }
    }
}

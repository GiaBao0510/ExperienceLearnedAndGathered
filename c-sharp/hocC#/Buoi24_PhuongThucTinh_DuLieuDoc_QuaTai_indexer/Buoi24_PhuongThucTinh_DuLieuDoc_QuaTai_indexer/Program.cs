using System;
using System.Runtime.CompilerServices;

namespace buoi24
{
    //1.Phương thức tĩnh - Thuộc tính tĩnh
    class webpage
    {
        //Thuộc tính tĩnh
        public static int number { set; get; }

        //Phuong thức tĩnh
        public static void ShowNumberOfVisit()
        {
            Console.WriteLine($"So luot truy cap: {number}");
        }

        public void Count()
        {
            number++;
        }
    }

    //2.Chỉ có đọc
    class Student
    {
        public readonly string name;
        public readonly string MSSV = "B" + DateTime.Now.ToString("yyyy MM dd H m s").Replace(" ","");

        //Chỉ được gán khi vừa khởi tạo
        public Student(string _name)
        {
            this.name = _name;
        }
    }

    //3. Quá tải toán tử
    class Vector
    {
        public double x, y;

        public Vector(double x, double y)
        {
            this.x = x;
            this.y = y;
        }

        public void VectorInformation()
        {
            Console.WriteLine($"X = {x}, Y = {y}");
        }

        //Tạo toán tử cộng giữa 2 vector. VD: vector3 = vector1 + vector2
        public static Vector operator+(Vector a, Vector b)
        {
            return new Vector(a.x + b.x, a.y + b.y);
        }

        //Tạo toán tử trừ giữa 2 vector. VD: vector3 = vector1 - vector2
        public static Vector operator -(Vector a, Vector b)
        {
            return new Vector(a.x + b.x, a.y + b.y);
        }

        //Tạo toán tử cộng giữa 1 vector với 1 số nào đó. VD: vector3 = vector1 + 10
        public static Vector operator +(Vector a, double value)
        {
            return new Vector(a.x + value, a.y + value);
        }

        // ======= Indexer ======
        public double this[int i]
        {
            set
            {
                switch (i)
                {
                    case 0:
                        x = value;
                        break;
                    case 1:
                        y = value;
                        break;
                    default:
                        throw new Exception("Chi so khong xac dinh.");
                }
            }

            get
            {
                switch (i)
                {
                    case 0:
                        return x;
                    case 1:
                        return y;
                    default:
                        throw new Exception("Chi so khong xac dinh.");
                }
            }
        }
    }

    class program
    {
        static void Main(string[] args)
        {
            //== 1
            webpage web1 = new webpage(),
                web2 = new webpage();

            web1.Count();
            web2.Count();
            webpage.ShowNumberOfVisit();

            //==2
            Student sv = new Student("Pham Gia Bao");
            Console.WriteLine($"Ho te: {sv.name} - MSSV: {sv.MSSV}");

            //===3.
            Vector v1 = new Vector(15, 16),
                    v2 = new Vector(-5,10);

            var v3 = v1 + v2;
            v1.VectorInformation();
            v2.VectorInformation();
            v3.VectorInformation();


            //===4.
            Vector v4 = new Vector(15,10);
            v4[0] = 2;
            v4[1] = 20;
            v4.VectorInformation();
        }
    }
}

using System;

namespace buoi11
{
    class student : IDisposable
    {
        //THUỘC TÍNH
        public string fullname { set; get; }
        public int age { get; set; }

        //PHƯƠNG THỨC
        public student(string hoten, int tuoi)
        {
            this.fullname = hoten;
            this.age = tuoi;
        }

        public void Dispose()
        {
            Console.WriteLine("cancelled!");
        }
    }
    class program
    {
        static void Main(string[] args)
        {
            VuKhi KiemThan = new VuKhi();
            VuKhi Sung1 = new VuKhi("74dfe","Ak47",500);
            VuKhi Sung2 = new VuKhi("54ej", "Tieu lien", 300);

            Sung2.doSatThuong = 350;
            Sung2.NoiSanXuat = "VietNam";
            
            //Hiển thị
            Sung1.show();
            Sung2.show();

            //Giải phóng bộ nhớ
            Sung2 = null;
            Sung1 = null;
            KiemThan = null;

            //===== using
            using (student S = new student("Pham Gia Bao", 22))
            {
                //...
            }//Tự động giải phóng bộ nhớ ,khi đối tượng ra khỏi using

        }
    }
}


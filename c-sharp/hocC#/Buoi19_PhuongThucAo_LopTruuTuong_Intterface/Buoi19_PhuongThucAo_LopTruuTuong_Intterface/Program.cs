using System;

namespace buoi19
{
    //1.Lớp trừu tượng
    abstract class product {
        protected int ID;
        protected string ProductType; 

        //Phương thức ảo
        public virtual string ProductInformation()
        {
            return $"ID: {this.ID}.\n Loai San Pham: {this.ProductType}.\n";
        }

        //Phương thức trừu tượng
        public abstract void ShowProduct();
    }

    class SmartPhone: product
    {
        private string productName;
        protected float price;

        public SmartPhone(string tenSP, float giaca)
        {
            this.ID = 102;
            this.ProductType = "Dien thoai";
            this.price = giaca;
            this.productName = tenSP;
        }

        //Tạo phương thức ghi đè phương thức ảo - Nạp chồng phương thức
        public override string ProductInformation()
        {
            return $"\t ============== \n"+base.ProductInformation();
        }

        public void Show() => Console.WriteLine($"{this.ProductInformation()} Ten san pham: {this.productName}.\n Gia: {this.price}.\n");

        //Ghi đè lên phương thức trừu tượng
        public override void ShowProduct()
        {
            Console.WriteLine("Lop ke thua:");
            Console.WriteLine($"ID: {this.ID}.\n Product type: {this.ProductType}");
            Console.WriteLine($"Product name: {this.productName}.\n Price: {this.price}");
        }
    }

    //2.Giao diện
    interface hinhhoc
    {
        public double Acreage();        //Diện tích
        public double Perimeter();      //Chu vi
    }

    class HinhChuNhat : hinhhoc
    {
        private double width, length;

        public HinhChuNhat(double width, double length) {
            this.width = width;
            this.length = length;
        }

        public double Acreage()
        {
            return this.width * this.length;
        }

        public double Perimeter()
        {
            return 2*(this.width + this.length);
        }
    }

    class program
    {
        static void Main(string[] args)
        {
            SmartPhone IP = new SmartPhone("Iphone", 5000.0f);
            IP.Show();
            IP.ShowProduct();

            //Interface
            HinhChuNhat A1 = new HinhChuNhat(5,20);
            Console.WriteLine($"Chu vi: {A1.Perimeter()}. \nDien tich: {A1.Acreage()}");
        }
    }
}

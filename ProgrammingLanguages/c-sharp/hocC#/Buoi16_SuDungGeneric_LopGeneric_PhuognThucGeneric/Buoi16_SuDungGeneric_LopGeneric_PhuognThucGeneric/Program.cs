using System;

namespace buoi16
{
    //Kiểu generic áp dụng trên lớp
    class Product<KieuDaiDien>
    {
        public KieuDaiDien IDproduct;
        public KieuDaiDien ProductName;

        public Product(KieuDaiDien _id, KieuDaiDien _name)
        {
            this.IDproduct = _id;
            this.ProductName = _name;
        }

        public void Show()
        {
            Console.WriteLine($"ID: {this.IDproduct}.\n Product: {this.ProductName}.\n");
        }

    }

    class program
    {
        //Kiểu generic áp dụng trên phương thức 
        static void permutation<KieuDuLieu>(ref KieuDuLieu a,ref KieuDuLieu b)
        {
            KieuDuLieu temp = a;
            a = b; 
            b = temp;
        }
        static void Main(string[] args)
        {
            string a = "A", b = "B";
            Console.WriteLine($"A: {a} - B: {b}");

            permutation(ref a, ref b);
            Console.WriteLine($"A: {a} - B: {b}");

            //sử dụng lớp kiểu generic
            Product<string> sanpham = new Product<string>("123","But chi");
            sanpham.Show();

            Product<int> sanpham2 = new Product<int>(12, 500);
            sanpham2.Show();
        }
    }
}



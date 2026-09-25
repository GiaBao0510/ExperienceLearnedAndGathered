using System;

namespace buoi12
{
    class program
    {
        //Khai báo struct
        public struct Product
        {
            //Dữ liệu
            public string name;
            public float price;

            //Phương thức
            public string getNameProduct()
            {
                return this.name;
            }
            public float getPrice()
            {
                return this.price;
            }

            //Phương thức tạo
            public Product(string tenhang, float giaca)
            {
                this.name = tenhang;
                this.price = giaca;
            }

            public void show()
            {
                Console.WriteLine($"- Product: {this.name} - Price: {this.price}.");
            }
        }

        //Kiểu dữ liệu liệt kê enum
        enum HOCLUC { 
            Kem = 0, //Có thể thiêt lập giá trị bên trong dữ liệu
            Yeu, 
            TrungBinh, 
            Kha, 
            Gioi, 
            XuatSat = 10 
        }

        static void Main(string[] args)
        {

                //Struct
            Product product1;
            product1.name = "pen";
            product1.price = 1.25f;

            Console.WriteLine($"Product: {product1.getNameProduct()} - Price: {product1.getPrice()}");

            Product product2 = new Product("cup", 0.8f);
            product2.show();

            //Ví dụ về kiểu liệt kê
            HOCLUC hocluc;
            hocluc = HOCLUC.XuatSat;            //Lấy dữ liệu bên trong kiểu liệt kê
            hocluc = (HOCLUC)10;              //Chuyển số nguyên sang dữ liệu dựa trên số má nó nhận được
            int so = (int)HOCLUC.XuatSat;   //Xem giá trị tại dữ liệu xuất sắt này nhận được

            switch (hocluc)
            {
                case HOCLUC.Kem:
                    Console.WriteLine("Hoc luc kem");
                    break;
                case HOCLUC.Yeu:
                    Console.WriteLine("Hoc luc Yeu");
                    break;
                case HOCLUC.TrungBinh:
                    Console.WriteLine("Hoc luc trung binh");
                    break;
                case HOCLUC.Kha:
                    Console.WriteLine("Hoc luc kha");
                    break;
                case HOCLUC.Gioi:
                    Console.WriteLine("Hoc luc gioi");
                    break;
                case HOCLUC.XuatSat:
                    Console.WriteLine("Hoc luc xuat sat");
                    break;
            }
            Console.WriteLine($"Number: {so}");
;
            
        }
    }
}


using System;
using System.Collections.Generic;

namespace buoi26
{
    class Product
    {
        public string NameProduct { get; set; }
        public string ID { get; set; }
        public double price { get; set; }
        public string origin { get; set; }
    }

        class program
    {
        static void Main(string[] args)
        {
                //1. Khởi tạo danh sách
            List<int> ds_songuyen = new List<int>();
            
            //Thêm từng dữ liệu vào cuối danh sách
            ds_songuyen.Add(5);
            ds_songuyen.Add(15);
            ds_songuyen.Add(115);

            //Thêm mảng dữ liệu vào trong danh sách
            ds_songuyen.AddRange(new int[] {8,12,30,25,35,45});

            //Lấy số lượng phần tử
            Console.WriteLine($"Do dai: {ds_songuyen.Count}");

            //Lấy giá trị phần tử tại vị trí
            Console.WriteLine($"Gia tri tai vi tri 2: {ds_songuyen[2]}");

            //Thêm phần tử tại vị trí cụ thể
            ds_songuyen.Insert(2, 333);

            //Xóa phần tử tại vị trí cụ thể
            ds_songuyen.RemoveAt(2);

            //Xóa phần tử có giá trị đầu tiên đã tồn tại trong mảng
            ds_songuyen.Remove(250);

            //Duyệt các phần tử trong danh sách
            Console.WriteLine("[{0}]", string.Join(", ", ds_songuyen));

                //2. Tìm kiếm các phần tử trong danh sách

            //FIND(): Tìm kiếm giá trị số chẳn phân tử đầu tiên trong danh sách. Tham số trong đây là Delegate
            var KQ = ds_songuyen.Find(
                (e) =>
                {
                    return e %2 ==0;
                }   
            );
            Console.WriteLine(KQ);

            //FINDALL(): TÌm kiếm các phần tử chẳn trong mảng
            var KQ1 = ds_songuyen.FindAll(
                (e) =>
                {
                    return e % 2 == 0;
                }
            );
            Console.WriteLine("So chan: [{0}]", string.Join(", ", KQ1));

            //Tìm ID sản phẩm
            List<Product> cuaHang1 = new List<Product>()
            {
                new Product()
                {
                   ID = "sp01" ,NameProduct = "Ban chai", origin="Kim long", price=7.500
                },
                new Product()
                {
                   ID = "sp02" ,NameProduct = "Bit ong hut", origin="Kim long", price=18.500
                },
                new Product()
                {
                   ID = "sp03" ,NameProduct = "Ghe da", origin="Kim long", price=180.500
                },
                new Product()
                {
                   ID = "sp04" ,NameProduct = "Dep", origin="Kim long", price=80.500
                },
                new Product()
                {
                   ID = "sp05" ,NameProduct = "Tai nghe", origin="Kim long", price=300.500
                },
                new Product()
                {
                   ID = "sp06" ,NameProduct = "Ban", origin="Kim long", price=250.500
                }
            };

            var timSP = cuaHang1.Find(
                (e) =>
                {
                    return e.ID == "sp03";
                }
            );
            if(timSP  != null)
            {
                Console.WriteLine($"ID: {timSP.ID} - {timSP.NameProduct} - {timSP.origin} - {timSP.price}");
            }


            // Sắp xếp phần tử trong danh sách
                //Ban đầu
            Console.WriteLine("\t------- Ban dau:");
            foreach(var i in cuaHang1)
            {
                Console.WriteLine($"ID: {i.ID} - {i.NameProduct} - {i.origin} - {i.price}");
            }
                //Sắp xếp
            cuaHang1.Sort(
                (p1,p2) =>
                {
                    if (p1.price == p2.price) return 0;
                    if (p1.price > p2.price) return 1;
                    return -1;
                }
            );
            Console.WriteLine("\t ------------ Sau khi sap xep:");
            foreach (var i in cuaHang1)
            {
                Console.WriteLine($"ID: {i.ID} - {i.NameProduct} - {i.origin} - {i.price}");
            }

            //2. Kiểu SortedList
            Console.WriteLine("\t ===================== ");
                //Key có kiểu string và value là đối tượng Product
            SortedList<string, Product> Shops = new SortedList<string, Product>();
            
            Shops["kimloan"] = new Product() { ID = "sp01", NameProduct = "Ban chai", origin = "Kim long", price = 7.500 };
            Shops["thanhlong"] = new Product() {ID = "sp06" ,NameProduct = "Ban", origin="Kim long", price=250.500 };
            Shops["hauhottoc"] = new Product() { ID = "sp04", NameProduct = "Dep", origin = "Kim long", price = 80.500 };
            Shops["trihottoc"] = new Product() { ID = "sp02", NameProduct = "Bit ong hut", origin = "Kim long", price = 18.500 };
            Shops.Add("ThanhLiem", new Product() { ID = "sp05", NameProduct = "Tai nghe", origin = "Kim long", price = 300.500 } );

            //Đọc thông tin 
            var cuaHang = Shops["kimloan"];
            Console.WriteLine($"ID: {cuaHang.ID} - {cuaHang.NameProduct} - {cuaHang.origin} - {cuaHang.price}");

            Console.WriteLine("\t ===========Keys - value ========== ");
            foreach(KeyValuePair<string, Product> item in Shops)
            {
                var key = item.Key;
                var value = item.Value;
                Console.WriteLine($"{key} - {value.NameProduct}");
            }

            //Xóa thông tin trong SortedLisst dựa trên key
            Shops.Remove("hauhottoc");

            //Xóa tất cả phần tử
            Shops.Clear();
        }
    }
}
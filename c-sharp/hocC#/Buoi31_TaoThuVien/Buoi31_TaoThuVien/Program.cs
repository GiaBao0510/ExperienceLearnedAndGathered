using System;
using Newtonsoft.Json;  //Sử dụng thư viện mới tải về cho dự án
using XTL; //Thêm thư viện vừa tạo

namespace Buoi31
{
    class Product
    {
        public string Name { set; get; }
        public DateTime Expiry { set; get; }
        public string[] Sizes { set; get; }

        public Product(string tenSP, DateTime hansudung, string[] kichco)
        {
            this.Name= tenSP;
            this.Expiry = hansudung;
            this.Sizes = kichco;
        }

        public string ToString()
        {
            return $"San pham: {this.Name} - Han su dung: {this.Expiry}. - Kich co: {string.Join(", ", this.Sizes)}";
        }
    }

    class program
    {
        static void Main(string[] args) {
            Product sanpham = new Product("Apple", new DateTime(2024, 08, 09), new string[] { "small" });
            string json = JsonConvert.SerializeObject(sanpham);  //Chuyển đối đối tượng c# sang chuỗi json
            
            Console.WriteLine($"\n\t Chuyen doi doi tuong thanh chuoi json:");
            Console.WriteLine(json);

            Console.WriteLine($"\n\t Chuyen doi chuoi json thanh doi tuong:");
            string input = @"
            {
                'Name':'HomeLander',            
                'Expiry':'2024-12-09T00:00:00',
                'Sizes': ['small', 'huge', 'big']
            }".Replace("'", "\"");

            var sanphamchuyendoi = JsonConvert.DeserializeObject<Product>(input);
            Console.WriteLine(sanphamchuyendoi.ToString());

            Console.WriteLine($"\n\t Chuyen doi chuoi so thanh van ban:");
            double a = 100;
            var kq = Utils.NumberToText(a);
            Console.WriteLine(kq);
            Utils.HelloAll();
        }
    }
}
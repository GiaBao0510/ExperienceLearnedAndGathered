using System;
using System.Linq;
using System.Collections.Generic;
using System.ComponentModel.DataAnnotations;

namespace buoi30
{
    public class Product
    {
        public int idProduct { get; set; }
        public string nameProduct { get; set; }
        public double Price { get; set; }
        public string[] Color { get; set; }
        public int Brand { get; set; }

        //Phương thức khởi tạo
        public Product(int _id, string _name, double _price, string[] _color, int brand)
        {
            this.idProduct = _id;
            this.nameProduct = _name;
            this.Price = _price;
            this.Color = _color;
            this.Brand = brand;
        }

        public override string ToString()
            => $"ID: {this.idProduct}, Product: {this.nameProduct}, Price: {this.Price}, Color: {string.Join(", ", this.Color)}, Brand: {this.Brand}";
    }
    
    public class Brand
    {
        public int IDbrand { get; set; }
        public string NameBrand { get; set; }
    
    class program
    {
            static void Main(string[] args)
            {
                Product cuaHang = new Product(3, "Banh bong lan", 5000, new string[] { "red", "green", "blue" }, 2);
                //Console.WriteLine(cuaHang.ToString());

                //Tạo danh sách thương hiệu mẫu
                var Brands = new List<Brand>() {
                    new Brand{IDbrand=0, NameBrand="Kinh do"},
                    new Brand{IDbrand=1, NameBrand="Thang Long"},
                    new Brand{IDbrand=2, NameBrand="KFC"},
                };

                    //Tạo danh sách sản phẩm
                var ListProduct = new List<Product>() {
                    new Product(0,"Ao thun", 25000, new string[] { "red","yellow"}, 2),
                    new Product(6,"Ao ba lo", 555000, new string[] { "red","yellow"}, 1),
                    new Product(1,"Ao vest", 180000, new string[] { "red","blue"}, 3),
                    new Product(4,"Ao mu trum", 20000, new string[] { "violet","blue"}, 3),
                    new Product(7,"Ao dai", 350000, new string[] { "red","yellow","black"}, 2),
                    new Product(5,"Ao la", 18000, new string[] { "red","blue"}, 3),
                    new Product(2,"Ao so mi", 85000, new string[] { "red","green"}, 1),
                    new Product(3,"Ao tay dai", 105000, new string[] { "red","blue","white"}, 2),
                };

                //Truy vấn giá > 30000
                Console.WriteLine("\n\t===== truy van gia > 30000 ==== ");
                var query1 = from p in ListProduct where p.Price > 30000 select p;
                foreach (var item in query1)
                {
                    Console.WriteLine(item.ToString());
                }

                Console.WriteLine("\n\t=== Phuong thuc Select() ===");
                //Phương thức Select - ví dụ truy vấn lấy tên các phần tử trong danh sách
                var TenCacSP = ListProduct.Select(
                        (p) =>
                        {
                            return p.nameProduct + " " + p.Price + " VNĐ.";
                        }
                );
                foreach(var item in TenCacSP)
                {
                    Console.WriteLine($"Product: {item}");
                }

                //Phương thức Where - Ví dụ tìm sản phẩm có giá >= 18000 và <= 100000
                Console.WriteLine("\n\t=== Phuong thuc Where() ===");
                var truyvan1 = ListProduct.Where(
                    (p) => { return p.Price >= 18000 && p.Price<=100000; }    
                );
                foreach (var item in truyvan1)
                {
                    Console.WriteLine($"Find vest: {item}");
                }

                //Phương thức selectMany - ví dụ trả về tập hợp phần tử của thuộc tính nếu thuộc tính đó là mảng
                Console.WriteLine("\n\t=== Phuong thuc SelectMany() ===");
                var truyvan2 = ListProduct.SelectMany(
                    (p) => { return p.Color; }
                );
                foreach (var item in truyvan2)
                {
                    Console.WriteLine($"Color: {item}");
                }

                //Phương thức Min, Max, Sum, Average
                Console.WriteLine("\n\t=== Phuong thuc Min, Max, Sum, Average ===");
                Console.WriteLine($"Gia do lon nhat: {ListProduct.Max(p => p.Price)}");
                Console.WriteLine($"Gia do nho nhat: {ListProduct.Min(p => p.Price)}");
                Console.WriteLine($"Tong tien cac gia do: {ListProduct.Sum(p => p.Price)}");
                Console.WriteLine($"Trung binh tong tien cac gia do: {ListProduct.Average(p => p.Price)}");

                //Phương thức join
                Console.WriteLine("\n\t=== Phuong thuc join: ===");
                var truyvan3 = ListProduct.Join(Brands, e=>e.Brand, x=>x.IDbrand, (e, x) => {
                    return $"Product: {e.nameProduct} - Brand: {x.NameBrand}";
                });
                foreach (var item in truyvan3)
                {
                    Console.WriteLine($"{item}");
                }

                //Phương thức Groupjoin
                Console.WriteLine("\n\t=== Phuong thuc GroupJoin: ===");
                var truyvan4 = Brands.GroupJoin(ListProduct, b => b.IDbrand, p => p.Brand, 
                    (brand, products) =>
                    {
                        return new
                        {
                            Brand = brand.NameBrand,
                            Products = products
                        };
                    }
                );
                foreach (var grp in truyvan4)
                {
                    Console.WriteLine($"{grp.Brand}");
                    foreach(var sanpham in grp.Products)
                    {
                        Console.WriteLine(sanpham);
                    }
                }

                //Phương thức take
                Console.WriteLine("\n\t=== Phuong thuc Take: ===");
                foreach(var item in ListProduct.Take(3))
                {
                    Console.WriteLine(item.ToString());
                }

                //Phương thức skip
                Console.WriteLine("\n\t=== Phuong thuc Skip: ===");
                foreach (var item in ListProduct.Skip(5))
                {
                    Console.WriteLine(item.ToString());
                }

                //Phương thức oderby - ví dụ sắp xếp phần tử có giá tăng dần
                var truyvan5 = ListProduct.OrderBy(p => p.Price);
                Console.WriteLine("\n\t=== Phuong thuc OrderBy: ===");
                foreach (var item in truyvan5)
                {
                    Console.WriteLine(item.ToString());
                }

                //Phương thức OrderByDescending - ví dụ sắp xếp phần tử có giá giảm dần
                var truyvan6 = ListProduct.OrderByDescending(p => p.Price);
                Console.WriteLine("\n\t=== Phuong thuc OrderByDescending: ===");
                foreach (var item in truyvan6)
                {
                    Console.WriteLine(item.ToString());
                }

                //Phương thức Reverse
                var truyvan7 = ListProduct.AsEnumerable().Reverse();
                Console.WriteLine("\n\t=== Phuong thuc OrderByDescending: ===");
                foreach (var item in truyvan7)
                {
                    Console.WriteLine(item.ToString());
                }

                //Phương thức groupby
                var group = ListProduct.GroupBy( p => p.Brand);
                Console.WriteLine("\n\t=== Phuong thuc GroupBy: ===");
                foreach (var nhom in group)
                {
                    Console.WriteLine(nhom.Key);
                    foreach(var item in nhom)
                    {
                        Console.WriteLine(item);
                    }
                }

                //Phương thức Distinct
                Console.WriteLine("\n\t=== Phuong thuc Distinct: ===");
                ListProduct.SelectMany(p => p.Color).Distinct().ToList().
                    ForEach( e => Console.WriteLine($"Color: {e}"));

                //Phương thức Single
                Console.WriteLine("\n\t=== Phuong thuc Single: ===");
                var truyVan8 = ListProduct.Single(p => p.Price == 18000);
                Console.WriteLine($"Tim ket qua don: {truyVan8}");

                //Phương thức any
                Console.WriteLine("\n\t=== Phuong thuc any: ===");
                var truyvan9 = ListProduct.Any(p => p.Price == 545121);
                Console.WriteLine($"Tim san pham có gia 545121: {truyvan9}");

                //Phương thức all
                Console.WriteLine("\n\t=== Phuong thuc all: ===");
                var truyvan10 = ListProduct.All(p => p.Price > 10000);
                Console.WriteLine($"Tat ca san pham có gia tren 10000: {truyvan10}");

                //Phương thức count
                Console.WriteLine("\n\t=== Phuong thuc count: ===");
                var truyvan11 = ListProduct.Count();
                Console.WriteLine($"so luong san pham: {truyvan11}");

                /*Ví dụ thực tế tìm tên sản phẩm và tên thương hiệu sản phẩm 
                  có giá từ 20000 - 100000 và sắp xếp sản phẩm theo giá giảm dần
                */
                Console.WriteLine("\n\tVi du thuc te:");
                ListProduct.Where(p => p.Price >= 20000 && p.Price <= 100000)
                .OrderByDescending(p => p.Price)
                .Join(Brands, p => p.Brand, b => b.IDbrand, (sp, thuonghieu) =>
                {
                    return new
                    {
                        SanPham = sp.nameProduct,
                        Gia = sp.Price,
                        TenThuongHieu = thuonghieu.NameBrand
                    };
                })
                .ToList().ForEach(info =>
                {
                    Console.WriteLine($"{info.SanPham, 15} - {info.Gia} - {info.TenThuongHieu}");
                });

                Console.WriteLine("\n\tVi du thuc te cach 2:");
                var viDuThucTe = from p in ListProduct
                                 from color in p.Color
                                 where p.Price >= 10000 && p.Price <= 50000
                                 orderby p.Price
                                 select new
                                 {
                                     ID = p.idProduct,
                                     SanPham = p.nameProduct,
                                     GiaBan = p.Price,
                                     Mau = p.Color,
                                 }                      
                    ;
                viDuThucTe.ToList().ForEach(
                    infor =>
                    {
                        Console.WriteLine($"ID: {infor.ID} - San pham: {infor.SanPham} - Gia Ban: {infor.GiaBan} - Color: {string.Join(", ", infor.Mau)}");
                    }
                );

                // Từ khóa let
                Console.WriteLine("\n\tTu khoa let:");
                var truyvan12= from p in ListProduct
                      group p by p.Price into grp
                      orderby grp.Key
                      let soluong = "So luong la: " + grp.Count()
                      select new
                      {
                          GiaBan = grp.Key,
                          CacSanPham = grp.ToList(),
                          SoLuong = soluong
                      };
                truyvan12.ToList().ForEach(item =>
                {
                    Console.WriteLine($"Gia ban: {item.GiaBan}");
                    Console.WriteLine($"So luong: {item.SoLuong}");
                    item.CacSanPham.ForEach(p => Console.WriteLine(p));
                });

                //Từ khóa join dùng để kết hợp 2 nguồn dữ liệu
                Console.WriteLine("\n\tTu khoa join:");
                var truyvan13 = from p in ListProduct
                                join B in Brands on p.Brand equals B.IDbrand
                                select new
                                {
                                    SanPham = p.nameProduct,
                                    GiaBan = p.Price,
                                    ThuongHieu = B.NameBrand
                                };
                truyvan13.ToList().ForEach(item =>
                {
                    Console.WriteLine($"\nSan pham: {item.SanPham}" +
                        $"\nGia Ban: {item.GiaBan}" +
                        $"\nThuong hieu: {item.ThuongHieu}");
                });

            }
        }
    }
}

using System;
using System.Linq;
using System.Threading.Tasks;
using System.Collections.Generic;
using System.ComponentModel.DataAnnotations;
using System.ComponentModel.DataAnnotations.Schema;
using Microsoft.EntityFrameworkCore;
using Microsoft.EntityFrameworkCore.ChangeTracking;
using Microsoft.EntityFrameworkCore.Diagnostics;
using Microsoft.EntityFrameworkCore.Infrastructure;
using Microsoft.EntityFrameworkCore.Metadata.Builders;
using Microsoft.Extensions.DependencyInjection;
using Microsoft.Extensions.Logging;
using Pomelo.EntityFrameworkCore.MySql;

namespace Buoi40
{

    class Program
    {

        //Phương thức tĩnh tạo csdl
        static void CreateDatabase()
        {
            //Sử dụng từ khóa using để tự động giải phóng tài nguyên
            using var dbContext = new ProductDBContext();

            //Lấy tên csdl
            string DBname = dbContext.Database.GetDbConnection().Database;
            Console.WriteLine($"DBname: {DBname}");

            //Tạo csdl
            bool taoCSDL = dbContext.Database.EnsureCreated();
            if (taoCSDL == true)
            {
                Console.WriteLine("Da tao duoc csdl");
            }
            else
            {
                Console.WriteLine("CSDL nay da ton tai. ");
            }

        }

        //Xóa cơ sở dữ liệu
        static void DeleteDatabse()
        {
            using var dbContext = new ProductDBContext();

            bool xoaCSDL = dbContext.Database.EnsureDeleted();
            if (xoaCSDL == true)
            {
                Console.WriteLine("Xoa csdl thanh cong");
            }
            else
            {
                Console.WriteLine("Xoa csdl that bai");
            }
        }

        //Chèn dòng vào bảng trong csdl
        public static void InsertProduct()
        {
            using var dbContext = new ProductDBContext();
            
            //Thao tác này chỉ là thêm 1 dòng dữ liệu
            var sp = new Product();

            sp.ProductName = "Dien thoai";
            sp.ProductProvider = "Can Tho";

            dbContext.Add(sp);                  //Thêm đối tượng
            int num = dbContext.SaveChanges();  //Lưu lại câu lệnh và trả về số dòng bị tác đọng
            Console.WriteLine($"So dong bi tac dong: {num}");

            //Thao tác này là thêm nhiều dòng dữ liệu
            var ArrSP = new object[]
            {
                new Product(){ProductName = "Thuoc ke",ProductProvider="Can Tho" },
                new Product(){ProductName = "Tam bong",ProductProvider="Can Tho" },
                new Product(){ProductName = "Mat Kinh",ProductProvider="Bac lieu" },
            };

            dbContext.AddRange(ArrSP);
            int num2 = dbContext.SaveChanges();  //Lưu lại câu lệnh và trả về số dòng bị tác đọng
            Console.WriteLine($"So dong bi tac dong: {num2}");
        }

        //In tất cả sp
        public static void ReadProducts()
        {
            using var dbContext = new ProductDBContext();
            var cacSP = dbContext.products.ToList();
            cacSP.ForEach(Sp => Sp.PrintProduct());
        }

        //Lấy sp theo ID
        public static void getProduct(int ID)
        {
            using var dbContext = new ProductDBContext();
            
            //Lấy thông tin đầu tiên
            var _sp = (from sp in dbContext.products
                      where sp.ProductId == ID
                      select sp ).FirstOrDefault();
            
            if(_sp != null)
            {
                _sp.PrintProduct();
            }
            else
            {
                Console.WriteLine("null");
            }
        }

        //xóa dòng vào bảng trong csdl
        public static void DeleteProduct(int ID)
        {
            using var dbContext = new ProductDBContext();
            var _sp = (from sp in dbContext.products
                       where sp.ProductId == ID
                       select sp).FirstOrDefault();

            if (_sp != null)
            {
                dbContext.Remove( _sp );
                int num = dbContext.SaveChanges();
                Console.WriteLine($"So dong bi tac dong: {num}");
            }
            else
            {
                Console.WriteLine("Khong tim thay");
            }
        }

        //cập nhật dòng vào bảng trong csdl
        public static void UpdateProduct(int ID, string nameProduct)
        {
            using var dbContext = new ProductDBContext();
            var _sp = (from sp in dbContext.products
                       where sp.ProductId == ID
                       select sp).FirstOrDefault();
            
            if (_sp != null)
            {
                _sp.ProductName = nameProduct;
                int num = dbContext.SaveChanges();
                Console.WriteLine($"So dong bi tac dong: {num}");
            }
            else
            {
                Console.WriteLine("Khong tim thay");
            }
        }

        static void Main(string[] args)
        {
            //Entity -> Database, class
            // Database on MySQL : isuzu_local kế thừ từ lớp DbContext
            // -- ChuVu

            //CreateDatabase();
            //DeleteDatabse();
            //InsertProduct();
            //ReadProducts();
            getProduct(2);
            //UpdateProduct(1, "LapTopChiNa");
            //DeleteProduct(3);

        }
    }
}



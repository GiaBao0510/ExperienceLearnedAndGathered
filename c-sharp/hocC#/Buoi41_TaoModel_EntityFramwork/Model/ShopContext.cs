using System;
using System.Collections.Generic;
using System.Linq;
using System.Threading.Tasks;
using Microsoft.EntityFrameworkCore;
using Microsoft.Extensions.DependencyInjection;
using Microsoft.EntityFrameworkCore.Infrastructure;
using Microsoft.Extensions.Logging;

namespace Buoi41_TaoModel_EntityFramwork.Model
{
    public class ShopContext : DbContext
    {
        protected string connect_string=
            @"Data Source=127.0.0.1,3306;
            Initial Catalog=shopdata;
            User ID=root;Passsword=";
        public DbSet<Product> products{set;get;}
        public DbSet<Category> Categorys{set;get;}

        protected override void OnConfiguring(DbContextOptionsBuilder optionsBuilder)
        {
            base.OnConfiguring(optionsBuilder);

            //Tạo ILoggerFactory
            ILoggerFactory loggerFactory = LoggerFactory.Create(builder => builder.AddConsole());

            //Chỉ định phiên bản mysql
            var serverVersion = new MySqlServerVersion(new Version(9,0,1));
            optionsBuilder.UseLoggerFactory(loggerFactory);
            optionsBuilder.UseMySql(connect_string,serverVersion);
        }

        //Tạo DB
        public async Task CreateDB(){
            String DBname = Database.GetDbConnection().Database;

            Console.WriteLine($"Tạo: {DBname}");
            bool result = await Database.EnsureCreatedAsync();
            string resultstring = result ? "Tạo thành công" : "Đã có Database trước đó";
            Console.WriteLine($"CSDL: {DBname}:{resultstring}"); 
        }

        //Xóa DB
        public async Task DeleteDB(){
            String DBname = Database.GetDbConnection().Database;

            Console.WriteLine($"Có muốn thật sự xóa {DBname} này(y) ? ");
            string input = Console.ReadLine();

            //Kiểm tra nếu câu trả lời là y thì xóa
            if(input.ToLower() == "y"){
                bool delete = await Database.EnsureDeletedAsync();
                string deletionInfor = delete ? "Xóa thành công":"Không thể xóa";
                Console.WriteLine($"{DBname}:{deletionInfor}");
            }
        }

    }
}
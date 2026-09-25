using Microsoft.EntityFrameworkCore;
using MySql.EntityFrameworkCore.Extensions;
using Microsoft.Extensions.Logging;

namespace Buoi40
{
    //Biểu diễn cho 1 csdl
    public class ProductDBContext: DbContext
    {
        //Tạo ra 1 loggerfactory với tham số là builder
        public static ILoggerFactory loggerfactory = LoggerFactory.Create(builder =>
        {
            builder.AddFilter(DbLoggerCategory.Query.Name,LogLevel.Information);
            builder.AddConsole();
        });

        //Dbset
        public DbSet<Product> products { set; get; }

        //Chuỗi kết nối
        private const string connectString = "Data Source=127.0.0.1,3306; " +
                                            "Initial Catalog= isuzu_local_temp;" +
                                            "User ID=root;" +
                                            "Password=";    

        //Bất kỳ khi nào đối tượng DbContext được tạo mới, thì nó sẽ thi hành phương thức override
        //Chúng ta thường làm tròn phương thức này để cấu hình kết nối đến cơ sở dữ liệu
        //Phương thức này chạy khi DbContext mới được tạo ra
        protected override void OnConfiguring(DbContextOptionsBuilder optionsBuider)
        {
            base.OnConfiguring(optionsBuider);

            //Chỉ định phiên ản Mysql
            var serverVersion = new MySqlServerVersion(new Version(9,0,1));
            optionsBuider.UseLoggerFactory(loggerfactory);  //Dùng logger
            optionsBuider.UseMySql(connectString, serverVersion); //Để cho biết cái contexxt này làm việc với Mysql
        } 
    }
}
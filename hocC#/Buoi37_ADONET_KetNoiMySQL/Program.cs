using Microsoft.Extensions.Configuration;
using MySql.Data.MySqlClient;
using System;
using System.Data;
using System.Data.Common;
using System.Data.SqlClient;
using System.Security.Cryptography.X509Certificates;
using Microsoft.Extensions.Configuration.Json;
using System.IO;
using System.Runtime.CompilerServices;


namespace Buoi37_ADONET_KetNoiMySQL;

class Program
{

    //Kết nối đến cơ sở dữ liệu trong file json
    public static void GetConnectString()
    {
        var configBuilder = new ConfigurationBuilder()
            .SetBasePath(Directory.GetCurrentDirectory())   //Lấy file config ở thư mục hiện tại
            .AddJsonFile("appconfig.json");                 //nạp config định dạng JSON

        var configurationroot = configBuilder.Build();      //Tạo Configurationroot
        
        var connection = new MySqlConnection(configurationroot["csdl:ketnoi"]);            //Điểm truy cập đến cơ sở dũ liệu đã lưu trong file config

        Console.WriteLine($"{"ConnectionString",17} : {configurationroot["csdl:ketnoi"]}");
        Console.WriteLine($"DataSource: {connection.Database}");

        //Bắt sự kiện trạng thái kết nối thay đổi
        connection.StateChange += (object sender, StateChangeEventArgs e) =>
        {
            Console.WriteLine($"Ket noi thay doi: {e.CurrentState}, trang thai truoc: {e.OriginalState}");
        };

        //Mở kết nối
        connection.Open();

        //Dùng sqlcommand thi hành lệnh sql
        using(DbCommand command = connection.CreateCommand())
        {
            //Câu truy vấn SQL
            command.CommandText = "SELECT * FROM pq_user";
            var reader = command.ExecuteReader();
            //Đọc kết quả truy vấn
            Console.WriteLine("\n\tCác người dùng:");
            Console.WriteLine($"{"USERMA",10} USER_TEM");
            while (reader.Read())
            {
                Console.WriteLine($"{reader["USER_MA"],10} {reader["USER_TEN"]}");
            }
        }

        //Không dùng thì đóng kết nối
        connection.Close();
    }

    //Kết nối và thao tác với csdl MySQL
    public class Exam1
    {
        public static void Test()
        {
            //Tạo chuỗi kết nối bằng SqlConnectionStringBuilder
            var stringBuilder = new MySqlConnectionStringBuilder();
            stringBuilder["server"] = "baokhoagold.ddns.net";
            stringBuilder["Database"] = "qltv_khaihang";
            stringBuilder["Port"] = "3306";
            stringBuilder["User Id"] = "user_test";
            stringBuilder["Password"] = "12345";

            String sqlConnectionString = stringBuilder.ToString();

            var connection = new MySqlConnection(sqlConnectionString);

            Console.WriteLine($"{"ConnectionString",17} : {stringBuilder}");
            Console.WriteLine($"{"DataSource",17} : {connection.Database}");

            //Bắt sự kiện trạng thái kết nối thay đổi
            connection.StateChange += (object sender, StateChangeEventArgs e) =>
            {
                Console.WriteLine($"Kết nối thay đổi: {e.CurrentState}, trạng thái trước: {e.OriginalState}");
            };

            //Mở kết nối
            connection.Open();

            //Dùng SqlCommand thi hành
            using (DbCommand command = connection.CreateCommand())
            {
                //Câu truy vấn SQL
                command.CommandText = "SELECT * FROM loai_hang LIMIT 5";
                var reader = command.ExecuteReader();
                //Đọc kết quả truy vấn
                Console.WriteLine("\n\t Cac loai vang:");
                Console.WriteLine($"{"LoaiMa",10} {"TenSanPham"}");
                while (reader.Read())
                {
                    Console.WriteLine($"{reader["LOAIMA"],10} {reader["LOAI_TEN"]}");
                }
            }

            //Không dùng đến kết nối thì giải phóng (Đóng kết nối)
            connection.Close();
        }
    }

        //Tham số (Paramenters) trong SqlCommand
    public static void Paramenters()
    {
        var sqlConnecstring = "Data Source=127.0.0.1,3306; Initial Catalog= isuzu_local;User ID=root;Password=";
        var connection = new MySqlConnection(sqlConnecstring);
        connection.Open();  //Mở kết nối

        //Thi hành câu lệnh
        using(var cmd = connection.CreateCommand())
        {
            //Chuỗi truy vấn
            string querryString = "SELECT * FROM dc_danh_muc_vtpt " +
                "WHERE DVTID > @SoLuongMacDinh LIMIT 5";  //Đưa tham số vào câu truy vấn

            //Mệnh đề truy vấn
            cmd.CommandText = querryString;

            //Thiết lập tham số
            cmd.Parameters.AddWithValue("@SoLuongMacDinh", 2);

            //Khởi tạo 1 SqlParameter và thêm vào
            var danhmuc = new MySqlParameter("@SoLuongMacDinhV2", 5);  //Tạo tham số
            cmd.Parameters.Add(danhmuc);                        //Thêm vào SqlCommand

            //Đọc câu truy vấn
            using var reader = cmd.ExecuteReader();

            if (reader.HasRows)
            {
                //Nếu gọi lần đầu nó sẽ đọc dòng dữ liệu đầu tiên
                //Sau đó con trỏ sẽ nhảy sang dòng tiếp theo để đọc
                //Nếu đọc cuối dòng thì nó sẽ trả về false ,ngược lại là true
                while (reader.Read())
                {
                    var VTPTID = reader["VTPTID"];
                    var TEN_VTPT = reader["TEN_VTPT"];
                    Console.WriteLine($"{VTPTID} - {TEN_VTPT}");
                }  
            }
            else
            {
                Console.WriteLine("Khong co du lieu");
            }
        }

        connection.Close(); //Đóng kết nối
    }

    //1. ExecuteScalar .Ví dụ chèn 1 dòng mới vào bảng và trả về giá trị định danh của dòng mới chèn vào ID
    public static void exampple_ExecuteScalar()
    {
        var sqlConnecstring = "Data Source=127.0.0.1,3306; Initial Catalog= isuzu_local;User ID=root;Password=";
        var connection = new MySqlConnection(sqlConnecstring);

        //Mở kết nối
        connection.Open();
        using(var cmd = connection.CreateCommand())
        {
            string sql = "SELECT COUNT(*)SoLuong FROM dc_danh_muc_vtpt " +
                "WHERE DVTID > 10";

            cmd.CommandText = sql;

            //Thiết lập tham số
            cmd.Parameters.AddWithValue("@DVTID", 10);

            //Đọc kết quả try vấn - Trả về giá trị của dòng 1 cột 1
            var reader = cmd.ExecuteScalar();
            Console.WriteLine(reader);
        }
        connection.Close();
    }

    //ExecuteNonQuery - Nó không trả về kết quả, mà nó trả về tổng số dòng bị tác động
    //Thường là Insert - update
    public static void example_ExecuteNoQuery()
    {
        var sqlConnecstring = "Data Source=127.0.0.1,3306; Initial Catalog= isuzu_local;User ID=root;Password=";
        var connection = new MySqlConnection(sqlConnecstring);

        int SoLanThem = 5;

        //Mở kết nối
        connection.Open();

        using(var cmd = connection.CreateCommand())
        {
            string sql = "INSERT INTO " +
                "danh_muc_cong_ld(CONGLAODONGID, DVTID, TEN_DVT, SO_LUONG_MAC_DINH, TEN_CONG_LD) " +
                "VALUES(@CONGLAODONGID,@DVTID,@TEN_DVT,@SO_LUONG_MAC_DINH,@TEN_CONG_LD)";
            cmd.CommandText = sql;
            
            cmd.Parameters.AddWithValue("@CONGLAODONGID", "AK-FRR-TMVSI-CK1");
            cmd.Parameters.AddWithValue("@DVTID", 2);
            cmd.Parameters.AddWithValue("@TEN_DVT", "Công");
            cmd.Parameters.AddWithValue("@SO_LUONG_MAC_DINH", 10);
            cmd.Parameters.AddWithValue("@TEN_CONG_LD", "Công tháo lắp, sửa chữa");

            var result = 0;
            for (int i= 0; i < SoLanThem; i++)
            {
                result = cmd.ExecuteNonQuery();
            }
            Console.WriteLine($"So dong bi tac dong: {result}");
        }

        //Đóng kết nối
        connection.Close();
    }

    //StoredProcedure  - Xây dụng thủ tục
    

    static void Main(string[] args)
    {
        //Paramenters();
        //exampple_ExecuteScalar();
        example_ExecuteNoQuery();

    }
}

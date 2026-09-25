using System;
using System.Linq;

namespace buoi17
{
    //Lop sinh viên
    class student
    {
        public string mssv;
        public string hoten;
        public string noisinh;
        public int namsinh;
    }

    class sinhvien
    {
        public string Name;

        public void HelloWorld()
        {
            Console.WriteLine(Name);
        }
    }

    class Program
    {

        static void PrintInformation(dynamic obj)
        {
            obj.Name = "GiaBao";
            obj.HelloWorld();

        }

        static void Main(string[] args) {

            //>>> 1.Đối tượng vô danh
            var sanPham = new
            {
                TenSP = "Sach",
                Gia = 15.000f,
                NamXuatBan = 2024
            };
            Console.WriteLine($"Ten San pham: {sanPham.TenSP}\n. Gia: {sanPham.Gia}. \nNam san xuat: {sanPham.NamXuatBan}");

            //>>> 2.Tạo đối tượng vô danh và truy vấn trong Linq
            List<student> SinhViens = new List<student>()
            {
                new student(){mssv = "b2016941" ,hoten = "Nguyen Van A", noisinh = "CanTho", namsinh = 2002},
                new student(){mssv = "b2016942" ,hoten = "Nguyen Van B", noisinh = "CaMau", namsinh = 2001},
                new student(){mssv = "b2016943" ,hoten = "Nguyen Van C", noisinh = "AnGiang" , namsinh = 2003},
                new student(){mssv = "b2016944" ,hoten = "Nguyen Van D", noisinh = "NamDinh", namsinh = 2002}
            };

            //Truy vấn dữ liệu trong danh sách chứa các đối tượng
            var timSinhVien = from sv in SinhViens
                              where sv.namsinh == 2002
                              select new
                              {
                                  MSSV = sv.mssv,
                                  HoTen = sv.hoten,
                                  NoiSinh = sv.noisinh,
                                  NamSinh = sv.namsinh
                              };
            foreach (var sv in timSinhVien)
            {
                Console.WriteLine($"Ma so sinh vien: {sv.MSSV}\n." +
                    $"Ho ten : {sv.HoTen}.\n" +
                    $"Noi sinh: {sv.NoiSinh}.\n" +
                    $"Nam sinh: {sv.NamSinh}."
                    );
            }

            //>>> 3. Kiểu động
            sinhvien X = new sinhvien();
            PrintInformation(X);
        }
    }
}

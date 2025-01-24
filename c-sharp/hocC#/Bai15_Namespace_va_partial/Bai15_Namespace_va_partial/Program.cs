using System;
using MyNameSpace; //Sử dụng namespace khác
using SanPham;
using GhiTat = MyNameSpace.buoi15_child.program;

//Sử dụng phương thức tính trong lớp nào đó, mà không cần ghi lại tên lớp đó
using static System.Console;
using static System.Math;

namespace buoi15
{
    public class program
    {
        public static void Hello (){
            WriteLine("Hello World!");
            WriteLine(Math.PI);
        }

        static void Main(string[] args) {

            //Cách truy cập đến thành phần từ nào trong lớp từ namespace
            buoi15.program.Hello();

            //Truy cập đến namespace con trên cùng namespace
            MyNameSpace.buoi15_child.program.Hello();

            //Ghi tắt đường dẫn đến lớp trong namespace khác
            GhiTat.Hello();

            //ví dụ về partial
            Products sanpham = new Products("0123","Banh quy",15000,"Thom ngon");
            sanpham.nhasanxuat = new Products.manufactory();
            sanpham.nhasanxuat.address = "Can Tho";
            sanpham.nhasanxuat.FactoryName = "Kim dong";
            WriteLine(sanpham.getInformaton());
        }
    }
}

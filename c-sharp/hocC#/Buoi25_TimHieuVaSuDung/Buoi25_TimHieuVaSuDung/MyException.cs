using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;

namespace MyException
{
    //Xây dụng lớp bắt lỗi tên rỗng
    internal class _NameEmptyException : Exception
    {
        public _NameEmptyException(): base("Khong duoc de ten rong") {}
    }

    //Xây dựng lớp bắt lỗi tuổi không hợp lệ
    class _AgeException : Exception
    {
        public int Age { get; set; }
        public _AgeException() : base("Tuoi khong hop le, Tuoi phai > 0 hoac <= 100") { }

        public void Detail()
        {
            Console.WriteLine("Tuoi phai trong khoang tu 1 - 100");
        }


    }

}

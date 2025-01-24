using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;

namespace SanPham
{
    public partial class Products
    {
        public string description { set; get; }
        //Tạo ra đối lượng lớp con
        public manufactory nhasanxuat { set; get; }

        //Khai báo lớp con
        public class manufactory
        {
            public string FactoryName { set; get; }
            public string address { set; get; }
        }

        
    }
}

using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;

namespace SanPham
{
    public partial class Products
    {
            //Thuộc tính
        private string IDproduct;
        public string productname { set; get; }
        public decimal price { set; get; }

            //Phương thức

        //Các hàm khởi tạo
        public Products()
        {
            this.IDproduct = "null";
            this.productname = "null";
            this.price = 0;
            this.description = "null";
        }
        public Products(string ID, string tenSP, decimal gia, string mota)
        {
            this.IDproduct = ID; 
            this.productname = tenSP;
            this.price = gia;
            this.description = mota;
        }

        //Lấy thông tin
        public string getInformaton()
        {
            return @$"
                ID: {this.IDproduct}.\n 
                Product: {this.productname}.\n 
                Price: {this.price}.\n
                Description: {this.description}.\n
            ";
        }
    }
}

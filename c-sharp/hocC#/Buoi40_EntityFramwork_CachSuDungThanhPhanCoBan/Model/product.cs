using System;
using System.ComponentModel.DataAnnotations;
using System.ComponentModel.DataAnnotations.Schema;

namespace Buoi40
{
    //Khởi tạo tên bảng
    [Table("Product")]
    public class Product
    {
        [Key] //Tương đương với privary key
        public int ProductId { get; set; }

        [Required]          //NotNull
        [StringLength(50)]   //Quy định dài tối đa 50 ký tụ
        public string ProductName { get; set; }

        [StringLength(50)]
        public string ProductProvider { get; set; }

        public void PrintProduct() => Console.WriteLine($"{ProductId} - {ProductName} - {ProductProvider}");
    }
}
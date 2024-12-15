using System;
using System.Collections.Generic;
using System.Linq;
using System.Threading.Tasks;
using System.ComponentModel.DataAnnotations;
using System.ComponentModel.DataAnnotations.Schema;

namespace Buoi41_TaoModel_EntityFramwork.Model
{
    [Table("Product")]      //Ánh xạ bảng product vào trong csdl
    public class Product
    {
        [Key]               //Primary key
        public int ProductID{set;get;}

        [Required]
        [StringLength(50)]          //Not null & length <=50
        public string Name{set;get;}

        [Column(TypeName="Money")]      //Cột kiểu Money trong MySQL
        public decimal Price {set;get;}
    }
}
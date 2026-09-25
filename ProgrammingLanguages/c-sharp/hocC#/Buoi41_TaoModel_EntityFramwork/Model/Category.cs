using System;
using System.Collections.Generic;
using System.Linq;
using System.Threading.Tasks;
using System.ComponentModel.DataAnnotations;
using System.ComponentModel.DataAnnotations.Schema;

namespace Buoi41_TaoModel_EntityFramwork.Model
{
    [Table("Category")]
    public class Category
    {
        [Key]
        public int CategoryID {set;get;}
        
        [StringLength(100)]
        public string Name{set;get;}
        
        [Column(TypeName="ntext")]
        public string Description {set;get;}
    }
}
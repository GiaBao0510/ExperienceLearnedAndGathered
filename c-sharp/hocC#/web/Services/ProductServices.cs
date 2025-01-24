using System.Collections.Generic;
using System.Linq;

public class Product
{
    public string ID { set; get; }

    public string Name { set; get; }
    public double price{ set; get; }
}

public class ProductServices
{
    List<Product> products = new List<Product>();   //Danh sách sản phẩm
    
    public ProductServices()
    {
        products.AddRange(new Product[]
        {
            new Product(){ID = "sp01", Name="But bi", price = 5500},
            new Product(){ID = "sp02", Name="Quan da", price = 130000},
            new Product(){ID = "sp03", Name="Nhan Rong", price = 12513000},
            new Product(){ID = "sp04", Name="Tra sua truyen thong", price = 25000},
            new Product(){ID = "sp05", Name="Cap moc treo do", price = 20000},
            new Product(){ID = "sp06", Name="op lung dien thoai", price = 80000},
        });
    }

    //Tìm sản phẩm theo ID
    public Product FindProduct(string id)
    {
        var qr = from p in products where p.ID == id select p;
        return qr.FirstOrDefault(); //Có thì trả về kết quả .Ngược lại trả về Null
    }
}
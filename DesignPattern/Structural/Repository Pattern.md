**Repository Pattern** là một mẫu thiết kế phổ biến trong phát triển phần mềm, đặt biệt trong kiến trúc **Clean Architecture.** ==Mẫu này hoạt động như một phần trung gian giữa logic nghiệp vụ của ứng dụng và tầng dữ liệu==, giúp tách biệt logic truy cập dữ liệu khỏi logic nghiệp vụ.

---
## **Mục đính của Repository Pattern:** 

###### 1.**Trừu tượng hóa tầng dữ liệu:** Che giấu chi tiết cách dữ liệu lưu trữ và truy xuất.
###### 2.**Tạo điều kiện kiểm thử:** Dễ dàng tạo mock repository cho kiểm thử
###### 3.**Tách biệt quan tâm:** Giúp phân chia rõ ràng giữa xử lý logic nghiệp vụ và truy xuất dữ liệu
###### 4.**Tăng khả năng bảo trì:** Thay đổi nguồn dữ liệu không ảnh hưởng đến logic nghiệp vụ

---
## **1️⃣Cấu trúc mẫu dữ án:**

Cấu trúc mẫu dữ án như sau:
```csharp
/ProjectName
│── /Controllers
│    ├── ProductController.cs
│
│── /Services
│    ├── ProductService.cs
│
│── /Repositories
│    ├── IRepository.cs
│    ├── IProductRepository.cs
│    ├── ProductRepository.cs
│
│── /Models
│    ├── Product.cs
│
│── /Data
│    ├── AppDbContext.cs
│
│── Program.cs
```

---
## **2️⃣Thiết lập Repository Pattern trong .Net:**

### **📌 Bước 1: Tạo Model (Entity)**
📂 **`/Models/Product.cs`**

```csharp
namespace ProjectName.Models
{
    public class Product
    {
        public int Id { get; set; }
        public string Name { get; set; }
        public decimal Price { get; set; }
    }
}
```

### **📌 Bước 2: Tạo Database Context (Entity Framework Core)**
📂 **`/Data/AppDbContext.cs`**

```csharp
using Microsoft.EntityFrameworkCore;
using ProjectName.Models;

namespace ProjectName.Data
{
    public class AppDbContext : DbContext
    {
        public AppDbContext(DbContextOptions<AppDbContext> options) 
	        : base(options) { }

        public DbSet<Product> Products { get; set; }
    }
}
```

### **📌 Bước 3: Tạo Interface Repository chung (Generic Repository)**
📂 **`/Repositories/IRepository.cs`**

```csharp
namespace ProjectName.Repositories
{
    public interface IRepository<T> where T : class
    {
        Task<IEnumerable<T>> GetAllAsync();
        Task<T> GetByIdAsync(int id);
        Task AddAsync(T entity);
        Task UpdateAsync(T entity);
        Task DeleteAsync(int id);
    }
}
```

### **📌 Bước 4: Tạo Interface Repository riêng cho Product**
📂 **`/Repositories/IProductRepository.cs`**

```csharp
using ProjectName.Models;

namespace ProjectName.Repositories
{
    public interface IProductRepository : IRepository<Product>
    {
        Task<IEnumerable<Product>> GetProductsByPriceAsync(decimal minPrice);
    }
}
```

### **📌 Bước 5: Cài đặt Repository**
📂 **`/Repositories/ProductRepository.cs`**

```csharp
using Microsoft.EntityFrameworkCore;
using ProjectName.Data;
using ProjectName.Models;

namespace ProjectName.Repositories
{
    public class ProductRepository : IProductRepository
    {
        private readonly AppDbContext _context;

        public ProductRepository(AppDbContext context)
        {
            _context = context;
        }

        public async Task<IEnumerable<Product>> GetAllAsync()
        {
            return await _context.Products.ToListAsync();
        }

        public async Task<Product> GetByIdAsync(int id)
        {
            return await _context.Products.FindAsync(id);
        }

        public async Task AddAsync(Product entity)
        {
            await _context.Products.AddAsync(entity);
            await _context.SaveChangesAsync();
        }

        public async Task UpdateAsync(Product entity)
        {
            _context.Products.Update(entity);
            await _context.SaveChangesAsync();
        }

        public async Task DeleteAsync(int id)
        {
            var product = await _context.Products.FindAsync(id);
            if (product != null)
            {
                _context.Products.Remove(product);
                await _context.SaveChangesAsync();
            }
        }

        public async Task<IEnumerable<Product>> 
	        GetProductsByPriceAsync(decimal minPrice)
        {
            return await _context.Products
	            .Where(p => p.Price >= minPrice).ToListAsync();
        }
    }
}
```

### **📌 Bước 6: Tạo Service Layer (Xử lý logic nghiệp vụ)**
📂 **`/Services/ProductService.cs`**

```csharp
using ProjectName.Models;
using ProjectName.Repositories;

namespace ProjectName.Services
{
    public class ProductService
    {
        private readonly IProductRepository _productRepository;

        public ProductService(IProductRepository productRepository)
        {
            _productRepository = productRepository;
        }

        public async Task<IEnumerable<Product>> GetAllProductsAsync()
        {
            return await _productRepository.GetAllAsync();
        }

        public async Task AddProductAsync(Product product)
        {
            await _productRepository.AddAsync(product);
        }
    }
}
```

### **📌 Bước 7: Tạo API Controller**
📂 **`/Controllers/ProductController.cs`**

```csharp
using Microsoft.AspNetCore.Mvc;
using ProjectName.Models;
using ProjectName.Services;

namespace ProjectName.Controllers
{
    [ApiController]
    [Route("api/products")]
    public class ProductController : ControllerBase
    {
        private readonly ProductService _productService;

        public ProductController(ProductService productService)
        {
            _productService = productService;
        }

        [HttpGet]
        public async Task<IActionResult> GetAll()
        {
            var products = await _productService.GetAllProductsAsync();
            return Ok(products);
        }

        [HttpPost]
        public async Task<IActionResult> Create(Product product)
        {
            await _productService.AddProductAsync(product);
            return Ok("Product added successfully");
        }
    }
}
```

### **📌 Bước 8: Cấu hình Dependency Injection**
📂 **`Program.cs`**

```csharp
using Microsoft.EntityFrameworkCore;
using ProjectName.Data;
using ProjectName.Repositories;
using ProjectName.Services;

var builder = WebApplication.CreateBuilder(args);

// Đăng ký DbContext
builder.Services.AddDbContext<AppDbContext>
(options =>
	options.UseSqlServer
	(builder.Configuration.GetConnectionString("DefaultConnection"))
);

// Đăng ký Repository và Service
builder.Services.AddScoped<IProductRepository, ProductRepository>();
builder.Services.AddScoped<ProductService>();

// Cấu hình API
builder.Services.AddControllers();
builder.Services.AddEndpointsApiExplorer();
builder.Services.AddSwaggerGen();

var app = builder.Build();

app.UseSwagger();
app.UseSwaggerUI();

app.UseAuthorization();
app.MapControllers();

app.Run();
```

---
## **3️⃣Lợi ích của Repository Pattern**

**✅ Tách biệt concern:** phân tách rõ ràng giữa logic nghiệp vụ và truy cập dữ liệu
**✅ Khả năng kiểm thử** dễ dàng tạo mock cho repository trong unit testing
**✅ Khả năng thay đổi nguồn dữ liệu** thay đổi từ SQL Server sang MongoDB hoặc bất kỳ CSDL nào khác chỉ cần thay đổi implementation của repository
**✅ Code có tổ chức** tập trung logic truy cập dữ liệu mỗi entity trong các lớp riêng biệt.
**✅ Tính tái sử dụng:** Sử dụng lại repository cho nhiều service khác nhau.

---
## **4️⃣ Lưu ý khi triển khai:**

##### 1.**Không sử dụng Generic Repository khi không cần thiết:** nếu các entity có nhiều thao tác cụ thể, nên tạo repository riêng.
##### 2.**Cân nhắc Unit of Work:** với các giao dịch phức tạp, nên sử dụng kết hợp Unit of Work pattern  |✨ [Repository Pattern + Unit of Work](obsidian://open?vault=CuuAmChanKinh&file=c-sharp%2FBackEnd_Develop%2FDesignPattern_DotNet%2FDataAccessPatterns%2FRepository%20Pattern%20%2B%20Unit%20of%20Work)0
##### 3.**Repository chỉ truy cập dữ liệu:** Không đặt logic nghiệp vụ trong respository.
##### 4.**Cân nhắc Specification Pattern:** Kết hợp với Specification Pattern để xây dụng truy vấn phức tạp.


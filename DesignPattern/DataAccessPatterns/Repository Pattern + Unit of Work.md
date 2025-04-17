## **1️⃣ Unit of Work là gì?**

**Unit of Work(UoW)** là một mẫu thiết **(Design Pattern)** giúp quản lý giao dịch trong ứng dụng. Nó đảm bảo rằng tất cả các thay đổi trong một phiên làm việc **(unit of work)** sẽ được thực thi như một khối duy nhất. Nếu có lỗi xảy ra, tất cả các thay đổi sẽ được **rollback** để tránh dữ liệu bị lỗi hoặc không nhất quán

📌 **Lợi ích của Unit of Work:**
✅ **Đảm bảo tính nhất quán của dữ liệu (tất cả hoặc không có gì).**
✅ **Giảm số lần truy cập cơ sở dữ liệu** giúp tăng hiệu suất
✅ **Kiểm soát giao dịch tốt hơn khi làm với nhiều** respository.
✅ **Dễ dàng mở rộng và kiểm thử (Unit Test).**

---
## **2️⃣ Kiến trúc Repository Pattern + Unit of Work trong .NET**

📂 `Cấu trúc thư mục`

```csharp
/ProjectName
│── /Controllers
│    ├── ProductController.cs
│    ├── OrderController.cs
│
│── /Services
│    ├── ProductService.cs
│    ├── OrderService.cs
│
│── /Repositories
│    ├── IRepository.cs
│    ├── IProductRepository.cs
│    ├── IOrderRepository.cs
│    ├── ProductRepository.cs
│    ├── OrderRepository.cs
│
│── /UnitOfWork
│    ├── IUnitOfWork.cs
│    ├── UnitOfWork.cs
│
│── /Models
│    ├── Product.cs
│    ├── Oeder.cs
│
│── /Data
│    ├── AppDbContext.cs
│
│── Program.cs
```

---
## **3️⃣ Cài đặt Repository Pattern + Unit of Work trong .NET**

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

📂 **`/Models/Order.cs

```csharp
using System;
using System.Collections.Generic;

namespace ProjectName.Models
{
    public class Order
    {
        public int Id { get; set; }
        public DateTime OrderDate { get; set; }
        public string CustomerName { get; set; }
        public string ShippingAddress { get; set; }
        public decimal TotalAmount { get; set; }
        public List<OrderItem> Items { get; set; } = new List<OrderItem>();
    }

    public class OrderItem
    {
        public int Id { get; set; }
        public int OrderId { get; set; }
        public int ProductId { get; set; }
        public Product Product { get; set; }
        public int Quantity { get; set; }
        public decimal UnitPrice { get; set; }
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
        public DbSet<Order> Order { get; set; }
        public DbSet<OrderItem> OrderItem { get; set; }
        
        protected override void OnModelCreating(ModelBuilder modelBuilder) {
	        modelBuilder.ApplyConfigurationsFromAssembly
		        (typeof(ApplicationDbContext).Assembly); 
        }
        
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

📂 **`/Repositories/IOrderRepository.cs

```csharp
using System.Collections.Generic;
using System.Threading.Tasks;
using ProjectName.Models;

namespace ProjectName.Repositories
{
    public interface IOrderRepository
    {
        Task<IEnumerable<Order>> GetByCustomerNameAsync(string customerName);
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
        }

        public async Task UpdateAsync(Product entity)
        {
            _context.Products.Update(entity);
        }

        public async Task DeleteAsync(int id)
        {
            var product = await _context.Products.FindAsync(id);
            if (product != null)
            {
                _context.Products.Remove(product);
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

📂 **`/Repositories/OrderRepository.cs

```csharp
using Microsoft.EntityFrameworkCore;
using ProjectName.Data;
using ProjectName.Models;

namespace ProjectName.Repositories
{
    public class OrderRepository : IOrderRepository
    {
        private readonly AppDbContext _context;

        public OrderRepository(AppDbContext context)
        {
            _context = context;
        }

        public async Task<IEnumerable<Order>> GetAllAsync()
        {
            return await _context.Order.ToListAsync();
        }

        public async Task<Order> GetByIdAsync(int id)
        {
            return await _context.Order.FindAsync(id);
        }

        public async Task AddAsync(Order entity)
        {
            await _context.Order.AddAsync(entity);
        }

        public async Task UpdateAsync(Order entity)
        {
            _context.Order.Update(entity);
        }

        public async Task DeleteAsync(int id)
        {
            var Order = await _context.Order.FindAsync(id);
            if (Order != null)
            {
                _context.Products.Remove(Order);
            }
        }

        public async Task<IEnumerable<Order>> 
	        GetByCustomerNameAsync(string customerName)
        {
            return await _context.Order
	            .Where(o => o.CustomerName == customerName).ToListAsync();
        }
    }
}
```

### **📌 Bước 6: Tạo Interface cho Unit of Work**
📂 **`/UnitOfWork/IUnitOfWork.cs`**

```csharp
using ProjectName.Repositories;

namespace ProjectName.UnitOfWork
{
    public interface IUnitOfWork : IDisposable
    {
        IProductRepository Products { get; }
        IOrderRepository Products { get; }
        Task<int> SaveChangesAsync();
    }
}
```

### **📌 Bước 7: Cài đặt Unit of Work**
📂 **`/UnitOfWork/UnitOfWork.cs`**

```csharp
using ProjectName.Data;
using ProjectName.Repositories;

namespace ProjectName.UnitOfWork
{
    public class UnitOfWork : IUnitOfWork
    {
        private readonly AppDbContext _context;
        public IProductRepository Products { get; }
        public IOrderRepository Order { get; }
        private bool _disposed = false;

        public UnitOfWork(
	        AppDbContext context, 
	        IProductRepository productRepository,
	        IOrderRepository orderRepository
	    )
        {
            _context = context;
            Products = productRepository;
            Order = orderRepository;
        }

        public async Task<int> SaveChangesAsync()
        {
            return await _context.SaveChangesAsync();
        }

		protected virtual void Dispose(bool disposing){
			if(!_disposed){
				if(disposing)
					_context.Dispose();
					
				_disposed = true;
			}
		}

        public void Dispose()
        {
            Dispose(true); 
            GC.SuppressFinalize(this);
        }
    }
}
```

### **📌 Bước 8: Tạo Service Layer**
📂 **`/Services/ProductService.cs`**

```csharp
using ProjectName.Models;
using ProjectName.UnitOfWork;

namespace ProjectName.Services
{
    public class ProductService
    {
        private readonly IUnitOfWork _unitOfWork;

        public ProductService(IUnitOfWork unitOfWork)
        {
            _unitOfWork = unitOfWork;
        }

        public async Task<IEnumerable<Product>> GetAllProductsAsync()
        {
            return await _unitOfWork.Products.GetAllAsync();
        }

        public async Task AddProductAsync(Product product)
        {
            await _unitOfWork.Products.AddAsync(product);
            await _unitOfWork.SaveChangesAsync(); // Lưu thay đổi
        }
    }
}
```

📂 **`/Services/OrderService.cs

```csharp
using ProjectName.Models;
using ProjectName.UnitOfWork;

namespace ProjectName.Services
{
    public class OrderService
    {
        private readonly IUnitOfWork _unitOfWork;

        public OrderService(IUnitOfWork unitOfWork)
        {
            _unitOfWork = unitOfWork;
        }

        public async Task<Order> GetOrderAsync(int id) { 
	        return await _unitOfWork.Orders.GetByIdAsync(id); 
        }

		public async Task<IEnumerable<Order>> GetAllOrdersAsync(){ 
			return await _unitOfWork.Orders.GetAllAsync(); 
		}

		public async Task<bool> UpdateOrderAsync(Order order){
			try{
				_unitOfWork.Order.Update(order);
				await _unitOfWork.SaveChangesAsync();
				return true;
			}catch (Exception) { 
				return false; 
			}
		}

		public async Task<bool> CancelOrderAsync(int orderId){
			try{
				var order = await _unitOfWork.Order.GetByIdAssync(orderId);
				if(order == null)
					return false;

				//Hoàn lại số lượng sản phẩm
				foreach(var item in order.Items){
					var product = await _unitOfWork.Products
						.GetByIdAssync(item.ProductId);

					if(product != null){
						product.StockQuantity += item.Quantity;
						_unitOfWork.Products.Update(product);
					}
				}

				//Xóa đơn hàng
				_unitOfWork.Order.Delete(order);
				await _unitOfWork.CompleteAssync();
				return true;
				
			}
			catch (Exception) { 
				return false; 
			}
		}
    }
}
```
### **📌 Bước 9: Tạo API Controller**
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

📂 **`/Controllers/OrderController.cs

```csharp
using System.Threading.Tasks;
using Microsoft.AspNetCore.Mvc;
using YourProject.Application.Services;
using YourProject.Domain.Entities;

namespace ProjectName.Controllers
{
    [Route("api/[controller]")]
    [ApiController]
    public class OrdersController : ControllerBase
    {
        private readonly OrderService _orderService;

        public OrdersController(OrderService orderService)
        {
            _orderService = orderService;
        }

        [HttpGet]
        public async Task<IActionResult> GetAll()
        {
            var orders = await _orderService.GetAllOrdersAsync();
            return Ok(orders);
        }

        [HttpGet("{id}")]
        public async Task<IActionResult> GetById(int id)
        {
            var order = await _orderService.GetOrderAsync(id);
            if (order == null)
            {
                return NotFound();
            }
            return Ok(order);
        }

        [HttpPost]
        public async Task<IActionResult> Create(Order order)
        {
            var result = await _orderService.PlaceOrderAsync(order);
            if (result)
            {
                return CreatedAtAction(
	                nameof(GetById), new { id = order.Id }, order
                );
            }
            return BadRequest("Không thể tạo đơn hàng");
        }

        [HttpPut("{id}")]
        public async Task<IActionResult> Update(int id, Order order)
        {
            if (id != order.Id)
            {
                return BadRequest();
            }

            var result = await _orderService.UpdateOrderAsync(order);
            if (result)
            {
                return NoContent();
            }
            return BadRequest("Không thể cập nhật đơn hàng");
        }

        [HttpDelete("{id}")]
        public async Task<IActionResult> Cancel(int id)
        {
            var result = await _orderService.CancelOrderAsync(id);
            if (result)
            {
                return NoContent();
            }
            return BadRequest("Không thể hủy đơn hàng");
        }
    }
}
```

### **📌 Bước 10: Cấu hình Dependency Injection**
📂 **`Program.cs`**

```csharp
builder.Services.AddScoped<IProductRepository, ProductRepository>();
builder.Services.AddScoped<IUnitOfWork, UnitOfWork>();
builder.Services.AddScoped<ProductService>();
builder.AddScoped<OrderService>();
```

---
## 4️⃣**Lợi ích của Repository + Unit of Work:**

###### 1.**Tính toàn vẹn dữ liệu:** Tất cả thay đổi được lưu cùng lúc, tránh tình trạng lưu nửa vời.
###### 2.**Tái sử dụng:** Các `Repository` chia sẻ `cùng` DbContext, giảm overhead
###### 3.**Dễ mở rộng:** Thêm entity mới chỉ cần thêm Repository mới và cập `UnitOfWork`.

---
## **5️⃣Kết luận**

✅ **Unit of Work quản lý giao dịch của nhiều repository**.  
✅ **Giảm số lần truy cập database, giúp cải thiện hiệu suất**.  
✅ **Dễ dàng mở rộng và kiểm thử**.
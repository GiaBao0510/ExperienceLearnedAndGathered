## **Mô hình MVC là gì?**

Mô hình Model-View-Controller(MVC) là một mấu kiến trúc phần tách ứng dụng thành 3 thành phần chính: Model, View & Controller. Mỗi thành phần kiến trúc được xây dụng để xử lý công việc cụ thể của một ứng dụng. MVC tách lớp logic nghiệp vụ và lớp hiển thị ra riêng biệt. HIện đang là mẫu kiến trúc phần mềm được sử dụng phổ biến.

![](https://images.viblo.asia/ed9ea401-6c09-4ca8-9aa1-8b143956db8f.png)

---
## **Kiến trúc MVC!**

![](https://static.vietnix.vn/wp-content/uploads/2022/03/cac-thanh-phan-cua-mvc.webp)

###### **MVC quan trong bao gồm:**

- *Model*: phần này bao gồm tất cả dữ liệu và logic nghiệp vụ của ứng dụng. Đồng thời Model cũng chứa các lớp mô tả business logic và định nghĩa business rule cho dữ liệu
- *View*: phần này trình bày hiển thị dữ liệu cho người dùng hoặc xử lý tương tác của người dùng.
- *Controller*: là phần quan trọng nhất trong mô hình, nó liên kết phần Model với View
###### **View**
- View là một phần của ứng dụng đại điện cho việc ==trình bày dữ liệu==.
- Là giao diện người dùng (**User Interface**)
- View được tại bởi các dữ liệu mà chúng lấy dữ liệu từ model. Một View yêu cầu Model cung cấp đầy đủ dữ liệu để cho nó hiển thị ra cho người dùng.
- View chính ==là nơi chứa giao diện== (ví dụ: nút bấm, khung nhập, menu,...) nó đảm nhiệm hiển thị dữ liệu và giúp người dùng tương tác với hệ thống.
- Nơi nhận dữ liệu từ **Controller** và hiển thị. 
- Bộ phân này ==đại diện cho các thành phần UI== (ví dụ: HTML, XML, ...)
>*Ví dụ*: Trong ứng dụng mua sắm, View sẽ xác định cách hiển thị giỏ hàng cho người dùng và nhận dữ liệu từ Model để hiển thị. View sẽ bao gồm tất cả các thành phần UI như hiển thị nút bấm, danh sách thả xuống, v.v. mà người dùng cuối cùng tương tác.

###### **Controller**
- **Controller** là một phần của ứng dụng nhằm để x==ử lý tương tác của người dùng==. Bộ điều khiển diễn giải đầu vào từ người dùng, thông báo cho model và view để thay đổi khi thích hợp.
- Là ==thành phần **trung gian** giữa **Model** và **View**==
- **Controller** là nơi để tiếp nhận những yêu cầu xử lý từ người dùng, nó sẽ gồm những class/function xử lý nhiều nghiệp vụ logic giúp lấy dụng dữ liệu thông tin cần thiết nhờ vào các nghiệp vụ lớp Model cung cấp và hiển thị dữ liệu đó ra cho người dùng nhờ vào lớp View.
- Controller gửi các lệnh đến model để làm thay đổi trạng thái của nó (Ví dụ: ta thêm mới 1 User hoặc chỉnh sửa thông tin từ 1 User). Controller cũng gửi các lệnh đến view liên quan của nó để thay đổi cách hiển thị view (Ví dụ: Xem thông tin 1 User).
>*Ví dụ:* Trong ứng dụng mua sắm, ở giỏ hàng của người dùng, bạn có thể thêm các button cho phép người dùng thêm hoặc xóa các mặt hàng.
>
>Những hành động này của người dùng yêu cầu Model phải được cập nhật, do đó, đầu vào sẽ được gửi đến Controller, sau đó Controller sẽ thao tác với Model cho phù hợp, sau đó Controller sẽ gửi dữ liệu cập nhật đến View.

###### **Model**
- Thành phần model l==ưu trử dữ liệu và logic nghiệp vụ liên quan== của nó. Bao gồm các class function xử lý các tác vụ như truy vấn, thêm ,sửa ,xóa dữ liệu. Ví dụ như một đối tượng Controller lấy thông tin khách hàng từ cơ sở dữ liệu. Nó thao các dữ liệu và gửi lại cơ sở dữ liệu hoặc sử dụng nó để hiển thị dữ liệu.
-  ==Có nhiệm vụ thao tác với **Database**==
- Nó chứa **tất cả các hàm**, **các phương thức truy vấn trực tiếp** với dữ liệu.
- **Controller** sẽ thông qua các hàm, phương thức đó để **lấy dữ liệu rồi gửi qua View**
- Nếu trạng thái của dữ liệu này thay đổi thì Model thường sẽ thông báo cho View (để màn hình có thể thay đổi khi cần) và đôi khi là Controller (Nếu cần Logic khác để cập nhật View).
>*Ví dụ*: Giả sử bạn đang phát triển một ứng dụng mua sắm. Ở đây, Model sẽ chỉ định giỏ hàng sẽ bao gồm những dữ liệu nào — như mặt hàng, giá cả, v.v. — và những dữ liệu nào đã có sẵn trong giỏ hàng.

---
## **Cách Thức hoạt động**

![](https://static.vietnix.vn/wp-content/uploads/2021/05/Luong-di-trong-cua-mo-hinh-MVC.webp)

- Controller tương tác qua lại với View.
-  Controller tương tác qua lại với Mode.
- Model và View không có sự tương tác với nhau trực tiếp mà nó phải thông qua Controller.

**Ví dụ:** Khi người dùng ấn đăng nhập từ *View* thì sẽ gửi *request* đến *Controller*. *Controller* sẽ gọi *Model* để xử lý logic đăng nhập từ người dùng và trả về thông tin kết quả từ *Model* sang *Controller* rồi gửi đến *View*. 

---
## **Ứng dụng mô hình MVC vào lập trình như thế nào?**

- Ngôn ngữ lập trình và framework mà bạn dùng phụ thuộc nhiều hơn vào mục đích nghề nghiệp. Nhưng lập trình MVC dưới dạng kiến trúc sẽ luôn là sự lựa chọn khả thi để phát triển nghề nghiệp của bạn.
- Ví dụ: mọi người đang dần chuyển từ Dotnet MVC sang Dotnet Core. Nhưng hiện nay vẫn còn nhu cầu về [Django](https://vietnix.vn/django-la-gi/) cũng sử dụng MVC.

---
## **Ưu/nhược điểm:**

##### **Ưu điểm:**
- Bảo trì code dễ dàng, dễ dàng mở rộng và phát triển.
- Việc phát triển các thành phần khác nhau có thể thực hiện song song.
- Giúp tránh sự phức tạp bằng cách cho ứng dụng thành ba đơn vị chính: Model, View, Controller.
- Cung cấp hộ trợ tốt nhất cho phát triển theo hướng thử nghiệm.
- Cung cấp khả năng phân tách rõ ràng các mối quan tâm
- Tất cả các đối tượng được phân loại và đối tượng độc lập với nhau để có thể kiểm tra chúng một cách riêng biệt.

##### **Nhược điểm:**
- Khó đọc, thay đổi, kiểm tra và sử dụng lại mô hình này.
- Không có hộ trợ xác thực chính thức.
- Tăng độ phức tạp và tính kém hiệu quả của dữ liệu.
- Bảo trì nhiều code trong Controller.
- Khó triển khai.

---
### **Ví dụ về cách áp dụng MCV vào bài toán bán hàng:**

Cấu trúc thư mục
```text
MyMvcApp/
├── Controllers/           # Chứa các Controller
│   ├── HomeController.cs
│   ├── ProductController.cs
│   └── AccountController.cs
├── Models/               # Chứa các Model
│   ├── Product.cs
│   ├── User.cs
│   └── ViewModels/
│       └── ProductViewModel.cs
├── Views/                # Chứa các View
│   ├── Home/
│   │   ├── Index.cshtml
│   │   └── About.cshtml
│   ├── Product/
│   │   ├── Index.cshtml
│   │   ├── Details.cshtml
│   │   └── Create.cshtml
│   ├── Shared/
│   │   ├── _Layout.cshtml
│   │   └── _ViewStart.cshtml
│   └── _ViewImports.cshtml
├── wwwroot/              # Chứa static files
│   ├── css/
│   ├── js/
│   └── images/
├── Data/                 # Chứa DbContext (nếu dùng Entity Framework)
│   └── ApplicationDbContext.cs
├── Services/             # Chứa business logic
│   └── ProductService.cs
├── appsettings.json
├── Program.cs
└── Startup.cs
```

Tệp tin ***Models/Product.cs***:
```csharp
using System.ComponentModel.DataAnnotations;

namespace MyMvcApp.Models
{
    public class Product
    {
        public int Id { get; set; }
        
        [Required(ErrorMessage = "Tên sản phẩm là bắt buộc")]
        [StringLength(100, ErrorMessage = "Tên sản phẩm không được vượt quá 100 ký tự")]
        public string Name { get; set; }
        
        [Range(0, double.MaxValue, ErrorMessage = "Giá phải lớn hơn 0")]
        public decimal Price { get; set; }
        
        public string Description { get; set; }
        
        public DateTime CreatedDate { get; set; }
        
        public bool IsActive { get; set; }
    }
}
```

Tệp tin ***Models/ViewModels/ProductViewModel.cs***:
```csharp
namespace MyMvcApp.Models.ViewModels
{
    public class ProductViewModel
    {
        public int Id { get; set; }
        public string Name { get; set; }
        public decimal Price { get; set; }
        public string Description { get; set; }
        public string FormattedPrice => Price.ToString("N0") + " VND";
        public string CreatedDateString => CreatedDate.ToString("dd/MM/yyyy");
        public DateTime CreatedDate { get; set; }
    }
}
```

Tệp tin ***Controllers/ProductController.cs***:
```csharp
using Microsoft.AspNetCore.Mvc;
using MyMvcApp.Models;
using MyMvcApp.Models.ViewModels;
using MyMvcApp.Services;

namespace MyMvcApp.Controllers
{
    public class ProductController : Controller
    {
        private readonly IProductService _productService;
        
        public ProductController(IProductService productService)
        {
            _productService = productService;
        }
        
        // GET: Product
        public async Task<IActionResult> Index()
        {
            var products = await _productService.GetAllProductsAsync();
            var viewModels = products.Select(p => new ProductViewModel
            {
                Id = p.Id,
                Name = p.Name,
                Price = p.Price,
                Description = p.Description,
                CreatedDate = p.CreatedDate
            }).ToList();
            
            return View(viewModels);
        }
        
        // GET: Product/Details/5
        public async Task<IActionResult> Details(int id)
        {
            var product = await _productService.GetProductByIdAsync(id);
            if (product == null)
            {
                return NotFound();
            }
            
            var viewModel = new ProductViewModel
            {
                Id = product.Id,
                Name = product.Name,
                Price = product.Price,
                Description = product.Description,
                CreatedDate = product.CreatedDate
            };
            
            return View(viewModel);
        }
        
        // GET: Product/Create
        public IActionResult Create()
        {
            return View();
        }
        
        // POST: Product/Create
        [HttpPost]
        [ValidateAntiForgeryToken]
        public async Task<IActionResult> Create(Product product)
        {
            if (ModelState.IsValid)
            {
                product.CreatedDate = DateTime.Now;
                product.IsActive = true;
                
                await _productService.CreateProductAsync(product);
                TempData["Success"] = "Sản phẩm đã được tạo thành công!";
                return RedirectToAction(nameof(Index));
            }
            
            return View(product);
        }
        
        // GET: Product/Edit/5
        public async Task<IActionResult> Edit(int id)
        {
            var product = await _productService.GetProductByIdAsync(id);
            if (product == null)
            {
                return NotFound();
            }
            
            return View(product);
        }
        
        // POST: Product/Edit/5
        [HttpPost]
        [ValidateAntiForgeryToken]
        public async Task<IActionResult> Edit(int id, Product product)
        {
            if (id != product.Id)
            {
                return NotFound();
            }
            
            if (ModelState.IsValid)
            {
                await _productService.UpdateProductAsync(product);
                TempData["Success"] = "Sản phẩm đã được cập nhật thành công!";
                return RedirectToAction(nameof(Index));
            }
            
            return View(product);
        }
        
        // GET: Product/Delete/5
        public async Task<IActionResult> Delete(int id)
        {
            var product = await _productService.GetProductByIdAsync(id);
            if (product == null)
            {
                return NotFound();
            }
            
            return View(product);
        }
        
        // POST: Product/Delete/5
        [HttpPost, ActionName("Delete")]
        [ValidateAntiForgeryToken]
        public async Task<IActionResult> DeleteConfirmed(int id)
        {
            await _productService.DeleteProductAsync(id);
            TempData["Success"] = "Sản phẩm đã được xóa thành công!";
            return RedirectToAction(nameof(Index));
        }
    }
}
```

Tệp tin ***Services/ProductService.cs***: 
```csharp
using MyMvcApp.Models;
using MyMvcApp.Data;
using Microsoft.EntityFrameworkCore;

namespace MyMvcApp.Services
{
    public interface IProductService
    {
        Task<IEnumerable<Product>> GetAllProductsAsync();
        Task<Product> GetProductByIdAsync(int id);
        Task CreateProductAsync(Product product);
        Task UpdateProductAsync(Product product);
        Task DeleteProductAsync(int id);
    }
    
    public class ProductService : IProductService
    {
        private readonly ApplicationDbContext _context;
        
        public ProductService(ApplicationDbContext context)
        {
            _context = context;
        }
        
        public async Task<IEnumerable<Product>> GetAllProductsAsync()
        {
            return await _context.Products
                .Where(p => p.IsActive)
                .OrderByDescending(p => p.CreatedDate)
                .ToListAsync();
        }
        
        public async Task<Product> GetProductByIdAsync(int id)
        {
            return await _context.Products
                .FirstOrDefaultAsync(p => p.Id == id && p.IsActive);
        }
        
        public async Task CreateProductAsync(Product product)
        {
            _context.Products.Add(product);
            await _context.SaveChangesAsync();
        }
        
        public async Task UpdateProductAsync(Product product)
        {
            _context.Products.Update(product);
            await _context.SaveChangesAsync();
        }
        
        public async Task DeleteProductAsync(int id)
        {
            var product = await _context.Products.FindAsync(id);
            if (product != null)
            {
                product.IsActive = false; // Soft delete
                await _context.SaveChangesAsync();
            }
        }
    }
}
```

Tệp tin ***Data/ApplicationDbContext.cs***:
```csharp
using Microsoft.EntityFrameworkCore;
using MyMvcApp.Models;

namespace MyMvcApp.Data
{
    public class ApplicationDbContext : DbContext
    {
        public ApplicationDbContext(DbContextOptions<ApplicationDbContext> options)
            : base(options)
        {
        }
        
        public DbSet<Product> Products { get; set; }
        
        protected override void OnModelCreating(ModelBuilder modelBuilder)
        {
            base.OnModelCreating(modelBuilder);
            
            modelBuilder.Entity<Product>(entity =>
            {
                entity.HasKey(e => e.Id);
                entity.Property(e => e.Name).IsRequired().HasMaxLength(100);
                entity.Property(e => e.Price).HasColumnType("decimal(18,2)");
                entity.Property(e => e.Description).HasMaxLength(500);
            });
        }
    }
}
```

Tệp tin ***Startup.cs***:
```csharp
using Microsoft.EntityFrameworkCore;
using MyMvcApp.Data;
using MyMvcApp.Services;

var builder = WebApplication.CreateBuilder(args);

// Add services to the container
builder.Services.AddControllersWithViews();

// Add Entity Framework
builder.Services.AddDbContext<ApplicationDbContext>(options =>
    options.UseSqlServer(builder.Configuration.GetConnectionString("DefaultConnection")));

// Add custom services
builder.Services.AddScoped<IProductService, ProductService>();

var app = builder.Build();

// Configure the HTTP request pipeline
if (!app.Environment.IsDevelopment())
{
    app.UseExceptionHandler("/Home/Error");
    app.UseHsts();
}

app.UseHttpsRedirection();
app.UseStaticFiles();
app.UseRouting();
app.UseAuthorization();

app.MapControllerRoute(
    name: "default",
    pattern: "{controller=Home}/{action=Index}/{id?}");

app.Run();
```

---
## **Tham khảo:**

1. [Tìm hiểu mô hình MVC dành cho người mới bắt đầu: Cấu trúc và ví dụ](https://viblo.asia/p/tim-hieu-mo-hinh-mvc-danh-cho-nguoi-moi-bat-dau-cau-truc-va-vi-du-V3m5WLDyKO7)
2. [Mô hình MVC là gì? Ví dụ và cách ứng dụng MVC vào lập trình](https://vietnix.vn/tim-hieu-mo-hinh-mvc-la-gi/)
3. 
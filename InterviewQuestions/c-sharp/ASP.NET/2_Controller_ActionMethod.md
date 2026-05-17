## **Controller là gì?**
Controller là thành phần trung tâm trong mô hình MVC (Model-View-Controller), đóng vai trò như cầu nối giữa người dùng và ứng dụng.

**Định nghĩa:**
- Controller là các lớp (class) chứa logic xử lý các HTTP request từ client
- Mỗi Controller thường quản lý một nhóm chức năng liên quan
- Các lớp Controller thường thừa kế từ lớp `Controller` hoặc `ControllerBase` của ASP.NET

**Ví dụ về một Controller đơn giản:**

```csharp
public class HomeController : Controller
{
    public IActionResult Index()
    {
        return View();
    }
    
    public IActionResult About()
    {
        return View();
    }
}
```

---
## **Action Method là gì?**

Action Method là các phương thức công khai (public methods) bên trong Controller, chịu trách nhiệm xử lý các HTTP request cụ thể.

### **Đặc điểm của Action Method:**

1. **Phải là public method**
2. **Không thể là static method**
3. **Không thể có generic parameters**
4. **Thường có kiểu trả về là `IActionResult` hoặc các kiểu kế thừa từ nó**

### **Ví dụ các Action Method:**

```csharp
public class ProductController : Controller
{
    // Action Method trả về View
    public IActionResult Index()
    {
        return View();
    }
    
    // Action Method với tham số
    public IActionResult Details(int id)
    {
        var product = GetProductById(id);
        return View(product);
    }
    
    // Action Method xử lý POST request
    [HttpPost]
    public IActionResult Create(Product product)
    {
        if (ModelState.IsValid)
        {
            SaveProduct(product);
            return RedirectToAction("Index");
        }
        return View(product);
    }
}
```

---
## **Routing - Ánh xạ URL đến Action Method**

Routing là cơ chế ánh xạ URL từ trình duyệt đến Action Method tương ứng trong Controller.

### **Cơ chế Routing mặc định:**

```
URL Pattern: /{Controller}/{Action}/{id?}
```

**Ví dụ:**

- `/Home/Index` → HomeController.Index()
- `/Product/Details/5` → ProductController.Details(5)
- `/User/Profile` → UserController.Profile()

### **Cấu hình Route trong Program.cs:**

```csharp
app.MapControllerRoute(
    name: "default",
    pattern: "{controller=Home}/{action=Index}/{id?}");
```

### **Custom Routing với Attributes:**

```csharp
[Route("products")]
public class ProductController : Controller
{
    [Route("")]
    [Route("list")]
    public IActionResult Index() { ... }
    
    [Route("details/{id:int}")]
    public IActionResult Details(int id) { ... }
    
    [Route("category/{category}")]
    public IActionResult ByCategory(string category) { ... }
}
```

---
## **Vòng đời của Controller**

### **Quan trọng: Controller Lifecycle**

**Đặc điểm quan trọng:**
- Mỗi HTTP request sẽ tạo ra một **instance mới** của Controller
- Controller object có vòng đời ngắn (per-request)
- Sau khi xử lý xong request, Controller object sẽ bị hủy

**Quá trình xử lý request:**
1. **Request đến** → ASP.NET nhận HTTP request
2. **Route Resolution** → Xác định Controller và Action Method
3. **Controller Instantiation** → Tạo instance mới của Controller
4. **Action Execution** → Thực thi Action Method
5. **Response Generation** → Tạo HTTP response
6. **Controller Disposal** → Hủy Controller instance

### **Dependency Injection và Controller:**

```csharp
public class ProductController : Controller
{
    private readonly IProductService _productService;
    
    // Constructor injection
    public ProductController(IProductService productService)
    {
        _productService = productService;
    }
    
    public IActionResult Index()
    {
        var products = _productService.GetAllProducts();
        return View(products);
    }
}
```

**Lưu ý về DI Lifetime:**

- **Transient**: Tạo instance mới mỗi lần request (mặc định cho Controller)
- **Scoped**: Một instance per HTTP request
- **Singleton**: Một instance duy nhất cho toàn ứng dụng

---

## **Các kiểu trả về của Action Method**

Action Method có thể trả về nhiều loại kết quả khác nhau thông qua interface `IActionResult`.

### **1. ViewResult - Trả về View**

```csharp
public IActionResult Index()
{
    return View(); // Trả về view tương ứng
}

public IActionResult Details(int id)
{
    var model = GetProductById(id);
    return View(model); // Trả về view với model
}
```

### **2. JsonResult - Trả về JSON**

```csharp
public IActionResult GetProductData(int id)
{
    var product = GetProductById(id);
    return Json(product);
}
```

### **3. RedirectResult - Chuyển hướng**

```csharp
public IActionResult Create(Product product)
{
    SaveProduct(product);
    return RedirectToAction("Index"); // Chuyển hướng đến action khác
}

public IActionResult GoToGoogle()
{
    return Redirect("https://www.google.com"); // Chuyển hướng external
}
```

### **4. ContentResult - Trả về text thuần**

```csharp
public IActionResult GetPlainText()
{
    return Content("Hello World!", "text/plain");
}
```

### **5. FileResult - Trả về file**

```csharp
public IActionResult DownloadFile()
{
    byte[] fileBytes = GetFileBytes();
    return File(fileBytes, "application/pdf", "document.pdf");
}

public IActionResult ViewImage()
{
    return PhysicalFile("/path/to/image.jpg", "image/jpeg");
}
```

### **6. EmptyResult - Không trả về gì**

```csharp
public IActionResult DoSomething()
{
    // Thực hiện một số logic
    ProcessData();
    return new EmptyResult(); // Hoặc return new EmptyResult();
}
```

### **7. PartialViewResult - Trả về Partial View**

```csharp
public IActionResult LoadPartialContent()
{
    var model = GetData();
    return PartialView("_PartialViewName", model);
}
```

### **8. StatusCodeResult - Trả về HTTP Status Code**

```csharp
public IActionResult NotFound()
{
    return StatusCode(404);
}

public IActionResult BadRequest()
{
    return BadRequest("Invalid data");
}

public IActionResult Unauthorized()
{
    return Unauthorized();
}
```

---

## **HTTP Verbs và Action Methods**

### **Sử dụng HTTP Verb Attributes:**

```csharp
public class ApiController : ControllerBase
{
    [HttpGet]
    public IActionResult GetAll()
    {
        return Ok(data);
    }
    
    [HttpGet("{id}")]
    public IActionResult GetById(int id)
    {
        return Ok(data);
    }
    
    [HttpPost]
    public IActionResult Create([FromBody] Model model)
    {
        return CreatedAtAction(nameof(GetById), new { id = model.Id }, model);
    }
    
    [HttpPut("{id}")]
    public IActionResult Update(int id, [FromBody] Model model)
    {
        return NoContent();
    }
    
    [HttpDelete("{id}")]
    public IActionResult Delete(int id)
    {
        return NoContent();
    }
}
```

---

## **Best Practices cho Controller**

### **1. Keep Controllers Thin (Giữ Controller gọn nhẹ)**

```csharp
// ❌ Bad - Logic phức tạp trong Controller
public class ProductController : Controller
{
    public IActionResult Create(Product product)
    {
        // Validate
        if (string.IsNullOrEmpty(product.Name)) { ... }
        if (product.Price <= 0) { ... }
        
        // Business logic
        product.CreatedDate = DateTime.Now;
        product.Slug = GenerateSlug(product.Name);
        
        // Save to database
        _context.Products.Add(product);
        _context.SaveChanges();
        
        return RedirectToAction("Index");
    }
}

// ✅ Good - Sử dụng Service
public class ProductController : Controller
{
    private readonly IProductService _productService;
    
    public ProductController(IProductService productService)
    {
        _productService = productService;
    }
    
    public IActionResult Create(Product product)
    {
        if (!ModelState.IsValid)
            return View(product);
            
        _productService.CreateProduct(product);
        return RedirectToAction("Index");
    }
}
```

### **2. Sử dụng Model Binding hiệu quả**

```csharp
// Binding từ URL parameters
public IActionResult Details(int id) { ... }

// Binding từ Form data
[HttpPost]
public IActionResult Create(Product product) { ... }

// Binding từ JSON body
[HttpPost]
public IActionResult CreateApi([FromBody] Product product) { ... }

// Binding từ Query string
public IActionResult Search([FromQuery] string term) { ... }
```

### **3. Error Handling**

```csharp
public IActionResult Details(int id)
{
    try
    {
        var product = _productService.GetById(id);
        if (product == null)
            return NotFound();
            
        return View(product);
    }
    catch (Exception ex)
    {
        _logger.LogError(ex, "Error getting product {Id}", id);
        return StatusCode(500, "Internal server error");
    }
}
```

---

## **Tóm tắt**

**Điểm quan trọng cần nhớ:**

1. **Controller** là trung tâm xử lý request trong MVC
2. **Action Method** là các phương thức xử lý logic cụ thể
3. **Routing** ánh xạ URL đến Action Method
4. **Controller instance** được tạo mới cho mỗi request
5. **IActionResult** cung cấp nhiều cách trả về response
6. **Best practice**: Giữ Controller gọn nhẹ, sử dụng Dependency Injection

Controller và Action Method là nền tảng của ứng dụng ASP.NET MVC, hiểu rõ về chúng sẽ giúp bạn xây dựng ứng dụng web hiệu quả và dễ bảo trì.
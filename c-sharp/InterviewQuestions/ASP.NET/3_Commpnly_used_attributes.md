# Các Attribute thường dùng trong ASP.NET

Attribute là các metadata được sử dụng để cung cấp thông tin bổ sung cho compiler và runtime về cách xử lý các class, method, hoặc parameter. Trong ASP.NET, attributes đóng vai trò quan trọng trong việc cấu hình routing, validation, và model binding.

---

## **1. Nhóm Attribute kiểm soát Controller và Action**

### **`[NonAction]` - Loại trừ method khỏi Action**

**Mục đích:** Đánh dấu một phương thức public trong Controller để nó **không** trở thành Action Method.

**Khi nào sử dụng:**

- Khi bạn có helper methods trong Controller
- Khi muốn tái sử dụng code nhưng không muốn expose ra ngoài

**Ví dụ:**

```csharp
public class HomeController : Controller
{
    public IActionResult Index()
    {
        var data = GetFormattedData();
        return View(data);
    }
    
    [NonAction]
    public string GetFormattedData()
    {
        // Helper method - không phải Action Method
        return "Formatted data";
    }
    
    // URL: /Home/Index ✅ Có thể access
    // URL: /Home/GetFormattedData ❌ Không thể access (404 error)
}
```

### **`[NonController]` - Loại trừ class khỏi Controller**

**Mục đích:** Đánh dấu một class kế thừa từ Controller để nó **không** được coi là endpoint.

**Đặc điểm quan trọng:**

- **Có tính chất thừa kế**: Các class con cũng sẽ bị ảnh hưởng
- Class được đánh dấu sẽ không được map vào bất kỳ URL nào

**Ví dụ:**

```csharp
[NonController]
public class BaseController : Controller
{
    protected ILogger Logger { get; set; }
    
    protected IActionResult HandleError(Exception ex)
    {
        Logger.LogError(ex, "An error occurred");
        return StatusCode(500);
    }
}

// Class này kế thừa BaseController cũng sẽ bị ảnh hưởng bởi [NonController]
public class HomeController : BaseController
{
    public IActionResult Index()
    {
        return View();
    }
}

// Giải pháp: Ghi đè bằng cách không kế thừa hoặc sử dụng interface
public class HomeController : Controller
{
    public IActionResult Index()
    {
        return View();
    }
}
```

---

## **2. Nhóm Attribute cho HTTP Methods**

### **HTTP Verb Attributes**

Các attribute này xác định loại HTTP request mà Action Method có thể xử lý.

#### **`[HttpGet]` - Xử lý GET Request**

```csharp
[HttpGet]
public IActionResult GetProducts()
{
    var products = _productService.GetAll();
    return Json(products);
}

// Có thể kết hợp với route template
[HttpGet("products/{id:int}")]
public IActionResult GetProduct(int id)
{
    var product = _productService.GetById(id);
    return Json(product);
}
```

#### **`[HttpPost]` - Xử lý POST Request**

```csharp
[HttpPost]
public IActionResult CreateProduct([FromBody] Product product)
{
    if (!ModelState.IsValid)
        return BadRequest(ModelState);
        
    _productService.Create(product);
    return CreatedAtAction(nameof(GetProduct), new { id = product.Id }, product);
}
```

#### **Các HTTP Verb khác:**

```csharp
[HttpPut("{id}")]
public IActionResult UpdateProduct(int id, [FromBody] Product product)
{
    _productService.Update(id, product);
    return NoContent();
}

[HttpDelete("{id}")]
public IActionResult DeleteProduct(int id)
{
    _productService.Delete(id);
    return NoContent();
}

[HttpPatch("{id}")]
public IActionResult PatchProduct(int id, [FromBody] JsonPatchDocument<Product> patch)
{
    // Xử lý partial update
    return NoContent();
}

[HttpHead]
public IActionResult CheckProductExists(int id)
{
    var exists = _productService.Exists(id);
    return exists ? Ok() : NotFound();
}

[HttpOptions]
public IActionResult GetOptions()
{
    Response.Headers.Add("Allow", "GET,POST,PUT,DELETE");
    return Ok();
}
```

---

## **3. Nhóm Attribute Model Binding (`From*`)**

Các attribute này chỉ định nguồn dữ liệu để ASP.NET lấy giá trị cho các parameter của Action Method.

### **`[FromQuery]` - Lấy từ Query String**

```csharp
public IActionResult Search([FromQuery] string keyword, [FromQuery] int page = 1)
{
    // URL: /Product/Search?keyword=laptop&page=2
    var results = _productService.Search(keyword, page);
    return View(results);
}

// Có thể bind vào object
public IActionResult Filter([FromQuery] ProductFilter filter)
{
    // URL: /Product/Filter?Category=Electronics&MinPrice=100&MaxPrice=500
    return View(filter);
}

public class ProductFilter
{
    public string Category { get; set; }
    public decimal? MinPrice { get; set; }
    public decimal? MaxPrice { get; set; }
}
```

### **`[FromForm]` - Lấy từ Form Data**

```csharp
[HttpPost]
public IActionResult CreateProduct([FromForm] Product product)
{
    // Dữ liệu từ HTML form hoặc form-data
    if (ModelState.IsValid)
    {
        _productService.Create(product);
        return RedirectToAction("Index");
    }
    return View(product);
}

// Với file upload
[HttpPost]
public IActionResult UploadFile([FromForm] IFormFile file, [FromForm] string description)
{
    if (file != null && file.Length > 0)
    {
        // Xử lý file upload
        var fileName = Path.GetFileName(file.FileName);
        // ...
    }
    return View();
}
```

### **`[FromBody]` - Lấy từ Request Body**

```csharp
[HttpPost]
public IActionResult CreateProductApi([FromBody] Product product)
{
    // Dữ liệu JSON từ request body
    if (!ModelState.IsValid)
        return BadRequest(ModelState);
        
    _productService.Create(product);
    return CreatedAtAction(nameof(GetProduct), new { id = product.Id }, product);
}

// Với complex object
[HttpPost]
public IActionResult BulkCreate([FromBody] List<Product> products)
{
    _productService.CreateBulk(products);
    return Ok();
}
```

### **`[FromHeader]` - Lấy từ HTTP Header**

```csharp
public IActionResult GetUserInfo([FromHeader] string authorization, 
                                [FromHeader(Name = "X-API-Key")] string apiKey)
{
    // authorization sẽ lấy từ header "Authorization"
    // apiKey sẽ lấy từ header "X-API-Key"
    
    if (string.IsNullOrEmpty(apiKey))
        return Unauthorized();
        
    return Ok();
}

// Ví dụ với custom header
public IActionResult ProcessRequest([FromHeader(Name = "X-Correlation-ID")] string correlationId)
{
    _logger.LogInformation("Processing request with correlation ID: {CorrelationId}", correlationId);
    return Ok();
}
```

### **`[FromRoute]` - Lấy từ URL Route**

```csharp
[Route("products/{id:int}/reviews/{reviewId:int}")]
public IActionResult GetReview([FromRoute] int id, [FromRoute] int reviewId)
{
    // URL: /products/5/reviews/123
    // id = 5, reviewId = 123
    var review = _reviewService.GetReview(id, reviewId);
    return Json(review);
}

// Thường không cần thiết vì ASP.NET tự động bind từ route
public IActionResult GetProduct(int id) // Tự động lấy từ route
{
    return Json(_productService.GetById(id));
}
```

### **`[FromServices]` - Lấy từ Dependency Injection Container**

```csharp
public IActionResult GetData([FromServices] IProductService productService,
                            [FromServices] ILogger<HomeController> logger)
{
    // Lấy service trực tiếp từ DI container
    logger.LogInformation("Getting product data");
    var products = productService.GetAll();
    return Json(products);
}

// Thường dùng khi cần service tạm thời hoặc optional
public IActionResult ProcessData(int id, [FromServices] ICacheService cache = null)
{
    var data = GetData(id);
    
    // Sử dụng cache nếu có
    cache?.Set($"data_{id}", data);
    
    return Json(data);
}
```

---

## **4. Routing Attributes**

### **`[Route]` - Định tuyến tùy chỉnh**

**Cấp Controller:**

```csharp
[Route("api/[controller]")]
public class ProductsController : ControllerBase
{
    // Tất cả action sẽ có prefix "api/Products"
}
```

**Cấp Action Method:**

```csharp
public class ProductsController : Controller
{
    [Route("")]
    [Route("products")]
    [Route("products/list")]
    public IActionResult Index()
    {
        // Có thể truy cập qua: /, /products, /products/list
        return View();
    }
    
    [Route("products/{id:int}")]
    public IActionResult Details(int id)
    {
        // URL: /products/5
        return View();
    }
    
    [Route("products/{category}/{id:int}")]
    public IActionResult CategoryDetails(string category, int id)
    {
        // URL: /products/electronics/5
        return View();
    }
}
```

**Route Constraints:**

```csharp
[Route("products/{id:int:min(1)}")]          // id phải là số nguyên >= 1
[Route("products/{slug:alpha}")]             // slug chỉ chứa chữ cái
[Route("products/{code:regex(^[A-Z]{{3}}$)}")]  // code phải là 3 chữ cái hoa
[Route("users/{email:email}")]               // email phải đúng format
[Route("files/{filename:file}")]             // filename có thể chứa dấu chấm
```

---

## **5. Validation Attributes**

### **Model Validation Attributes**

```csharp
public class Product
{
    [Required(ErrorMessage = "Tên sản phẩm là bắt buộc")]
    [StringLength(100, MinimumLength = 2, ErrorMessage = "Tên sản phẩm phải từ 2-100 ký tự")]
    public string Name { get; set; }
    
    [Range(0.01, double.MaxValue, ErrorMessage = "Giá phải lớn hơn 0")]
    public decimal Price { get; set; }
    
    [EmailAddress(ErrorMessage = "Email không hợp lệ")]
    public string ContactEmail { get; set; }
    
    [Url(ErrorMessage = "URL không hợp lệ")]
    public string Website { get; set; }
    
    [RegularExpression(@"^\d{10}$", ErrorMessage = "Số điện thoại phải có 10 chữ số")]
    public string PhoneNumber { get; set; }
}
```

### **Action Method Validation**

```csharp
[HttpPost]
[ValidateAntiForgeryToken] // Chống CSRF attack
public IActionResult Create([Bind("Name,Price,Description")] Product product)
{
    if (!ModelState.IsValid)
    {
        return View(product);
    }
    
    _productService.Create(product);
    return RedirectToAction("Index");
}
```

---

## **6. Authorization và Authentication Attributes**

```csharp
[Authorize] // Yêu cầu đăng nhập
public class AdminController : Controller
{
    [AllowAnonymous] // Cho phép truy cập ẩn danh
    public IActionResult Login()
    {
        return View();
    }
    
    [Authorize(Roles = "Admin,Manager")] // Yêu cầu role cụ thể
    public IActionResult Dashboard()
    {
        return View();
    }
    
    [Authorize(Policy = "MinimumAge")] // Yêu cầu policy tùy chỉnh
    public IActionResult RestrictedContent()
    {
        return View();
    }
}
```

---

## **7. Caching và Performance Attributes**

```csharp
[ResponseCache(Duration = 300)] // Cache 5 phút
public IActionResult Index()
{
    return View();
}

[ResponseCache(Duration = 0, Location = ResponseCacheLocation.None, NoStore = true)]
public IActionResult NoCache()
{
    return View();
}

[OutputCache(Duration = 60)] // .NET 7+
public IActionResult CachedContent()
{
    return View();
}
```

---

## **8. Content Negotiation Attributes**

```csharp
[Produces("application/json")] // Chỉ trả về JSON
public class ApiController : ControllerBase
{
    [Consumes("application/json")] // Chỉ nhận JSON
    [HttpPost]
    public IActionResult Create([FromBody] Product product)
    {
        return Ok();
    }
    
    [Produces("application/xml", "application/json")]
    public IActionResult GetData()
    {
        // Có thể trả về XML hoặc JSON tùy thuộc Accept header
        return Ok(data);
    }
}
```

---

## **9. Custom Attributes**

### **Tạo Custom Action Filter:**

```csharp
public class LogActionAttribute : ActionFilterAttribute
{
    public override void OnActionExecuting(ActionExecutingContext context)
    {
        var logger = context.HttpContext.RequestServices.GetRequiredService<ILogger<LogActionAttribute>>();
        logger.LogInformation("Action {Action} executing", context.ActionDescriptor.DisplayName);
    }
}

// Sử dụng
[LogAction]
public IActionResult Index()
{
    return View();
}
```

### **Custom Validation Attribute:**

```csharp
public class FutureDateAttribute : ValidationAttribute
{
    public override bool IsValid(object value)
    {
        if (value is DateTime date)
        {
            return date > DateTime.Now;
        }
        return false;
    }
}

// Sử dụng
public class Event
{
    [FutureDate(ErrorMessage = "Ngày sự kiện phải là tương lai")]
    public DateTime EventDate { get; set; }
}
```

---

## **10. Best Practices**

### **1. Sử dụng Attribute một cách có ý thức**

```csharp
// ✅ Good - Rõ ràng và có mục đích
[HttpPost]
[ValidateAntiForgeryToken]
[Authorize(Roles = "Admin")]
public IActionResult DeleteProduct(int id)
{
    _productService.Delete(id);
    return RedirectToAction("Index");
}

// ❌ Bad - Quá nhiều attribute không cần thiết
[HttpGet]
[HttpPost]
[AllowAnonymous]
[Authorize]
public IActionResult ConfusingAction()
{
    return View();
}
```

### **2. Tổ chức Attribute theo thứ tự**

```csharp
// Thứ tự khuyến nghị:
// 1. Route attributes
// 2. HTTP verb attributes
// 3. Authorization attributes
// 4. Validation attributes
// 5. Custom attributes

[Route("products/{id:int}")]
[HttpGet]
[Authorize]
[ResponseCache(Duration = 300)]
[LogAction]
public IActionResult Details(int id)
{
    return View();
}
```

### **3. Sử dụng Model Binding hiệu quả**

```csharp
// ✅ Good - Rõ ràng về nguồn dữ liệu
public IActionResult Search([FromQuery] string term, 
                           [FromQuery] int page,
                           [FromHeader] string authorization)
{
    return View();
}

// ✅ Good - Sử dụng model để group parameters
public IActionResult Search([FromQuery] SearchRequest request)
{
    return View();
}
```

---

## **Tóm tắt**

**Các điểm quan trọng cần nhớ:**

1. **Attributes** là metadata cung cấp thông tin cho framework
2. **`[NonAction]`** và **`[NonController]`** loại trừ khỏi routing
3. **HTTP Verb attributes** xác định loại request được xử lý
4. **`From*` attributes** chỉ định nguồn dữ liệu cho model binding
5. **`[Route]`** cho phép tùy chỉnh URL mapping
6. **Validation attributes** giúp kiểm tra dữ liệu đầu vào
7. **Authorization attributes** kiểm soát quyền truy cập
8. **Custom attributes** mở rộng chức năng theo nhu cầu

Sử dụng attributes một cách thông minh sẽ giúp code của bạn rõ ràng, dễ bảo trì và có hiệu suất tốt hơn.
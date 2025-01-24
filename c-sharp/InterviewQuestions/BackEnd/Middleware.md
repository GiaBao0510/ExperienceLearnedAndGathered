### **1. Middleware là gì?**
- Là các thành phần phần mềm được kết hợp vào một pipeline để xử lý các yêu cầu HTTP và phản hồi HTTP. Đây là một khái niệm trung gian trong kiến trúc ASP.NET Core cho phép việc thêm. sắp xếp hoặc các bước xử lý và phản hồi

---
### **2. Cách thức hoạt động của Middleware?** 
Khi một yêu cầu(request) đến ứng dụng:
1. **Request** sẽ đi qua các middleware theo thứ tự mà chúng được đăng ký trong pipeline
2. Mỗi middleware có thể:
	- Xử lý yêu cầu và trả về kết quả ngay lập tức( chặn pineline).
	- Gửi yêu cầu xuống middleware tiếp theo.
3. Response từ ứng dụng được xử lý ngược trở lại qua pipeline, đi qua các middleware theo thứ tự ngược lại.

---
### **3. Các đặc điểm của Middleware?**
1. **Thành phần xử lý yêu cầu:**
	- Middleware có thể chỉnh sửa, xác thực hoặc theo dõi yêu cầu tước khi nó đến logic chính trong ứng dụng
1. **Thành phần xử lý phản hồi:**
	- Middleware cũng có thể sử đổi phản hồi trước khi gửi nó trở lại máy khách.
2. **Pipeline:**
	- Middleware được sắp xếp theo một thứ tự, tạo thành một pipeline xử lý tuần tự

---
### **4. Middleware pipeline order?**

![middleware pipeline order](https://www.shekhali.com/wp-content/uploads/2022/01/Middleware-Ordering-in-ASP.NET-Core-1024x652.webp)

---
### **5. Giải thích về ExceptionHandler?**
- Xử lý ngoại lệ xảy ra trong pipeline và trả về thông tin lỗi có kiểm soát.
- Dùng để bắt lỗi từ người dùng hoặc hệ thống.
- Thường nằm ở đầu pipeline.
```
public void Configure(IApplicationBuilder app, IWebHostEnviroment env){
	if(env.IsDevelopment()){
		app.UseDeveloperExceptionPage();  //Hiển thị chi tiết lỗi khi phát triển
	}else{
		app.UseExceptionHandler("/Home/Error");  //Chuyển hướng đến một trang lỗi tùy chỉnh
	}
}
```

---
### **6. Giải thích về HSTS?**
- Thường HSTS đi cặp với HttpsRedirection
- Bật HTTP Strict Transport Security(HSTS), buộc trình duyệt chỉ sử dụng HTTPS.
- Cải thiện bảo mật.
- Dùng trong môi trưởng sản xuất.
- Được khuyến nghị sử dụng
```
public void Configure(IApplicationBuilder app, IWebHostEnviroment env){
	if(!env.IsDevelopment()){
		app.UseHsts();   //Kích hoạt HSTS
	}
}
```

---
### **7. Giải thích về HttpsRedirection?**
- Tự động chuyển hướng các yêu cầu HTTP sang HTTPS.
- Bảo vệ ứng dụng khỏi các cuộc tấn công thông qua kết nối HTTP không an toàn
```
public void Configure(IApplicationBuilder app, IWebHostEnviroment env){
	app.UseHttpsRedirection();
}
```

---
### **8. Giải thích về CookiePolicy? (Tham khảo)**
- Cung cấp các chính sách về cookie để kiểm soát cookie được gửi và nhận từ client.
- Hỗ trợ chính sách bảo mật như ==Secure, HttpOnly, SameSite.==
```
public void ConfigureServices(IServiceCollection services){
	services.Configure<CookiePolicyOptions>(options => {
		options.AddPolicy( "AllowSpecifiOrigin" ,policy =>{
			options.CheckConsentNeeded = context => true; // Yêu cầu sự đồng ý trước khi lưu cookie. 
			options.MinimumSameSitePolicy = SameSiteMode.Strict; // Cookie chỉ được gửi từ cùng một site.
		});
	});
}

public void Configure(IApplicationBuilder app, IWebHostEnviroment env){
	app.UseCookiePolicy();  //kích hoạt CookiePolicy
}
```

---
### **9. Giải thích về StaticFiles?**
- Phục vụ các tệp tĩnh như HTML, CSS, JavaScript, hình ảnh
- Thường đặt trong thư mục ==wwwroot==.
- Có thể cấu hình đường dẫn hoặc tệp tùy chỉnh
```
public void Configure(IApplicationBuilder app, IWebHostEnviroment env){
	app.UseStaticFiles();
}
```


---
### **10. Giải thích về Routing?**
- Ánh xạ các yêu cầu HTTP đến các endpoint được định nghĩa
- kết hợp với ==UseEndpoints== để định nghĩa các route.
```
public void Configure(IApplicationBuilder app, IWebHostEnviroment env){
	app.UseRouting();
	app.UseEndpoints(endpoints =>{
		endpoints.MapGet("/", async context =>{
			await context.Response.WriteAsync("Hello, world!");
		});
	});
}
```


---
### **11. Giải thích về RateLimiter?**
- Giới hạn số lượng yêu cầu (request) từ một client trong một khoảng thời gian nhất định.
- Bảo vệ hệ thống khỏi các cuộc tấn công DDoS hoặc lạm dụng tài nguyên.
```
public void ConfigureServices(IServiceCollection services){
	services.AddRateLimiter(options => {
		options.AddFixedWindowLimiter( "FixedPolicy" ,c =>{
				c.PermitLimit = 1000;  //Tối đa 1000 res
				c.Window = TimeSpan.FromSeconds(1);//trong 1 s
		});
	});
}

public void Configure(IApplicationBuilder app, IWebHostEnviroment env){
	app.UseRateLimiter();
}
```

---
### **12. Giải thích về CORS?**
- Kiểm soát cho phép chia sẻ tài nguyên giữa các domain
- Tiện khi giữa FrontEnd và BackEnd ở các domain khác nhau.
```
public void ConfigureServices(IServiceCollection services){
	services.AddCors(options => {
		options.AddPolicy( "AllowSpecifiOrigin" ,policy =>{
			policy.WithOrigins("*")
				.AllowCredentials()
				.AllowAnyMethod()
				.AllowAnyHeader();
		});
	});
}

public void Configure(IApplicationBuilder app, IWebHostEnviroment env){
	app.UseCors("AllowSpecifiOrigin");
}
```


---
### **13. Giải thích về Authentication?**
- Xác thực người dùng dựa trên các scheme như JWT, Cookie, OAuth
```
public void ConfigureServices(IServiceCollection services){
	services.AddAuthentication("Bearer")
		.AddJwtBearer(options =>{
			options.Authority = "https://example.com";
			options.Audience = "api1";
		});
}

public void Configure(IApplicationBuilder app, IWebHostEnviroment env){
	app.UseAuthentication();
}
```

---
### **14. Giải thích về Authorization?**
- Kiểm tra quyền truy cập dựa trên thông tin xác thực
```
public void Configure(IApplicationBuilder app, IWebHostEnviroment env){
	app.UseAuthorization();
}
```

---
### **15. Giải thích về Custom Middleware?**
- Phần này người dùng sẽ tự định nghĩa dựa trên yêu câu thực tế để xử lý các yêu cầu và phản hồi.
- Thêm logic để xử lý cụ thể.
```
public class LoggingMiddleware{
	private readonly RequestDeledate _next;

	public LoggingMiddleware(RequestDeledate next){
		_next = next;
	}

	public async Task Invoke(HttpContext context){
		Console.WriteLine($"Request: {context.Request.Method} {context.Request.Path}"); 
		await _next(context); // Tiếp tục đến middleware tiếp theo
	}
}

public void Configure(IApplicationBuilder app, IWebHostEnviroment env){
	app.UseMiddleware<LoggingMiddleware>();
}
```


---
### **16. Giải thích về EndPoints?**
- Kết thúc pipeline và trả về phản hồi cho client
- Có kết hợp với ==UseRouting== để định nghĩa các hành động cuối.
- Sẽ routing đến Action Controller API.
```
public void Configure(IApplicationBuilder app) { 
	app.UseRouting();
	app.UseEndpoints(endpoints => { 
		endpoints.MapControllers(); // Mapping các controller 
	}); 
}
```

### **1. ASP.NET Web API là gì?**
- ASP.NET Web API là một Framework giúp đơn giản hóa việc xây dựng các dịch vụ HTTP cho một loạt Client (app & web) trên nền tảng .Net Framework
- Bằng cách sử dụng ASP.NET Web API, có thể tạo ra các dịch vụ non-SOAP như chuỗi XML hoặc JSON,... với các ưu điểm:
	- Tạo dịch vụ hướng tài nguyên sử dụng đầy đủ tính năng của giao thực HTTP.
	- Tiếp cận các dịch vụ dễ dàng từ nhiều loại khách hàng như trình duyệt, thiết bị di động,...

### **2. Giải thích cách sử dụng của HttpResponseMessage?**
- ==HttpResponseMessage== cho phép làm việc với giao thức HTTP( ví dụ với thuộc tính headers) và thống nhất kiểu trả về. Để dễ hiểu, thì ==HttpResponseMessage== là một cách để trả về một message/data dựa trên hành động từ người dùng.
```
public HttpResponseMessage GetEmployee(int id){
	Employee emp = EmployeeContext.Employee.Where(e=>e.id == id).FirstOrDefault();
	if(emp != null){
		return Request.CreateResponse<Employee>(HttpStatusCode.OK, emp);
	}else{
		return Request.CreateResponse<Employee>(HttpStatusCode.NotFound, "Employee not found");
	}
}
```
### **3. Những lợi ích của việc sử dụng ASP.NET Web API là gì?**


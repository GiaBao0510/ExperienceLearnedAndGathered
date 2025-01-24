### **1. .NET Standard là gì?**
- **.NET Standard** là một cách để đặt tả chính thức của các **.NET APIs** với mục đích có sẵn trên tất cả các triển khai .NET.
### **2. .NET Core là gì?**
- Nền tảng .NET Core là một .NET stack mới được tối ưu hóa để phát triển mã nguồn mở và phân phối trên NuGet.
- .NET Core có 2 thành phần chính. Gồm một runtime nhỏ được xây dựng từ cùng một cơ sở mã như .NET Framework CLR. .NET Core runtime bao gồm cùng một GC và JIT (RyuJIT), nhưng không bao gồm các tính năng như Application Domains hoặc Code Access Security. Runtime này được phân phối qua NuGet, như 1 phần của APS.NET Core package.
- .NET Core bao gồm các class libraries này phần lớn là các mã giống như các .NET Framework class libraries, nhưng đã được tính theo yếu tố( loại bỏ các dependencies) để cho phép gửi một bộ libraries nhỏ hơn .Các libraries này được vận chuyển dưới dạng các _System.*_ NuGet package trên NuGet.org 

### **3. .NET Framework là gì?**
- .NET là một Framework, là một tập hợp các lớp thư viện có thể tái sử dụng do Microsoft cung cấp để sử dụng trong các ứng dụng .NET khác và để phát triển, xây dựng và triển khai nhiều loại ứng dụng trên nền tảng Window bao gồm:
	- **Console Applications**
	- **Window Forms Applications**
	- **Windows Presentation Foundation (WPF) Application.**
	- **Web Applications**
	- **Web Services**
	- **Window Services**
	- **Services-oriented applications** sử dụng **Windows Communication Foundation(WCF)**.
	- **Workflow-enabled application** sử dụng **Windows Workflow**
	- **Foundation(WF)**
### **4. Sự khác biệt giữa String và string trong C # là gì?**
- **string** là kiểu dữ liệu chuỗi trong C# cho **System.String**. Vì vậy, về mặt kỹ thuật, không có sự khác biệt. Tựa như **int** do với **System.Int32**
- _Ví dụ: string_
```
string place = "World";
```
- _Ví dụ String nếu cần tham chiếu đến lớp cụ thể:_
```
string greet = String.Format("Hello {0}!",place);
```

### **5. .NET application domain là gì?**
- Là một lớp biệt lập được cung cấp từ **.NET runtime** .Như vậy các app domains tồn tại cùng trong một process (một process có thể có nhiều app domains) và có không gian địa chỉ ảo của riêng chúng.
- App domains hữu ích vì:
	- Chúng ít tốn kém hơn các process đầy đủ.
	- chúng là đa luồng.
	- Có thể stop một cái mà không cần phải kill tất cả trong process đó.
	- Phân tách resources/ config/ v.v.
	- Mỗi app domain chạy ở cấp độ bảo mật riêng.

### **6. CLR là gì?**
- **CLR (Common Language Runtime)** và nó là một Execution Enviroment. Nó hoạt động như là một layer giữa Operating System và các ứng dụng được .NET tuân theo **CLS (Common Language Specification).**
- Chức năng chính của **Common Language Runtime (CLR)** là chuyển đổi _Managed code_ thành _native code_, sau đó sẽ thực thi chương trình.

![[Pasted image 20241214204729.png]]
### **7. Liệt kê một số CLR services?**
- Assembly Resolver
- Assembly Loader
- Type Checker
- COM marshalled
- Debug Manager
- Thread Support
- IL to Native compiler
- Exception Manager
- Garbage Collector
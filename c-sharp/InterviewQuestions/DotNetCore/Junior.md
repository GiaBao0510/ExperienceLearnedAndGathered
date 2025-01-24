### **1. .NET Standard là gì?**
1. **.NET Standard** giải quyết được các vấn đề chia sẻ mã nguồn từ những nhà phát triển .NET trên tất cả các nền tảng bằng cách cung cấp API mà họ mong đợi à yêu thích trên các môi trường mà họ cần: Desktop applications, Mobile apps & Games và Cloud services.
2. **.NET Standard** là **một tập hợp các API** mà **tất cả** các nền tảng .NET **phải triển khai**. Điều này **hợp nhất các nền tảng .NET** và n**găn chặn sự phân mảnh trong tương lai.**
3. **.NET Standard 2.0** sẽ được thực thi bởi **.NET Framework, .NET Core và Xamarin**. Đối với **.NET Core**, điều này sẽ thêm nhiều API hiện có mà được yêu cầu.
4. **.NET Standard 2.0** bao gồm một miếng đệm tương thích cho các **.NET Framework binaries**, làm tăng đáng kể bộ thư viện mà hiện bạn đang có thể them khảo từ các thư viện .NET Standard của mình
5. **.NET Standard sẽ thay thể các Portable Class Libraries(PCLs)** làm công cụ để xây dựng các thư viện .NET đa nền tảng.
![[Pasted image 20241214211708.jpg]]

### **2. Kể ra các đặt tính của .NET Core?**
- **Triển khai linh hoạt**: có thể đưa vào ứng dụng hoặc được cài đặt song song với user hoặc toàn bộ máy.
- **Đa nền tảng:** chạy trên Window, macOS và Linux, có thể chuyển sang hệ điều hành khác. Các hệ điều hành (OS), CPU và Application scenarios được hỗ trợ sẽ phát triển theo thời gian, được cung cấp bởi Microsoft, các công ty và cá nhân khác.
- **Công cụ Command-line:** Tất cả các Product scenarios có thể được thực hiện tại command-line
- **Tính tương thích:** .NET Core tương thích với .NET Framework, Xamarin và Mono thông qua .NET Standard Library.
- **Được hỗ trợ bỏ Microft**: .NET Core được hỗ trợ bởi Microsoft, theo .NET Core Support.

### **3. Sự khác biệt giữa SDK và Runtime trong .NET Core là gì?**
- **SDK ( Software Development Kit/ Bộ phát triển phần mềm):** là tất cả những thứ cần thiết/giúp phát triển ứng dụng .NET Core dễ dàng hơn, chẳng hạn nhưu CLI và trình biên dịch (Compiler).
- **Runtime** là "máy ảo" lưu trữ/ chạy ứng dụng và trừu tượng hóa tất cả các tương tác với hệ điều hành cơ sở.

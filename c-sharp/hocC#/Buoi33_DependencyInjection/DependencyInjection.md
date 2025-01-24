## **Dependency Injection(Mối quan hệ phụ thuộc):**

#### **1.1 Dependency Inversion:**
là nguyên lý cuối cùng trong nguyên lý SOLID, trong đó:
	- Các modules cấp cao không nên phụ phuộc vào các modules cấp thấp. Cả 2 nên phụ thuộc và _Abstraction._
	- Interface (abstracion) không nên phụ thuộc và chi tiết, mà ngược lại (Các class giao tiếp với nhau thông qua Interface, không phải thông qua Implementation).
**=>** Với cách code thông thường thì module cấp cao sẽ phụ thuộc vào module cấp thấp. Khi mà module cấp thấp thay đổi thì sẽ kéo theo hàng loạt module cấp cao thay đổi. Điều này dẫn đến giảm khả năng bảo trì code. 
* Vì vậy nên theo **Depandency Inversion Principle**, các module cùng phụ thuộc vào 1 Interface không đổi. Nên là thay đổi module cấp thấp mà không ảnh hưởng gì tới module cấp cao.

#### **1.2 Inversion of Control (IoC) / Dependency inversion - Đảo ngược phụ thuộc:**
- Đảo ngược điểu kiển là một nguyên lý được thiết kế trong công nghệ phần mềm trong đó các thành phần nó dựa vào vấn đề để làm việc bị đảo ngược quyền điều kiển khi so với lập trình hướng thủ tục
- Đây là một _Design pattern_ được tạo ra để code có thể tuân thủ nguyên lý **Dependency Inversion**. Có nhiều cách hiện thực pattern này: ServiceLocator, Event, Delegate,... **Dependency Injection** là một trong những cách đó.

#### **1.3 Dependency Injection (DI):**
- Là một nguyên tắc triển khai nguyên tắc đảo ngược sự phụ thuộc,
- Đưa các **Dependence** vào trong một lớp thì thao tác này gọi là **Injection.** 
- Thiết kế sự phụ thuộc **cứng** là các lớp có mối quan hệ phụ thuộc với nhau. _Ví dụ_ như lớp A muốn thực hiện thao tác nào đó phải thông qua lớp B thì khi lớp B bị thay đổi thì sẽ kéo theo lớp A bị thao đổi theo.
- Thiết kế sự phụ thuộc **lỏng lẻo** là các lớp có mối quan hệ phụ thuộc với nhau. _Ví dụ_ như bên trong lớp A có khai báo đối tượng kiểu lớp B và phương thức khỏi tạo của lớp A có tham số là đối tượng lớp B. Như vậy _Dependency_ được đưa vào lớp A thông qua phương thức khởi tạo (_Cách này gọi là Inject bằng phương thức khởi tạo_). Thì khi bên trong lớp B _bị thay đổi_  thì không có kéo theo lớp A phải thay đổi theo
- Là một design pattern được [ASP.Net](http://asp.net/) hỗ trợ. Đây là một kỹ thuật để hiện thực hóa Inversion of Control Pattern (có thể xem là một design pattern riêng ). Các module phụ thuộc (dependency) sẽ được inject vào module cấp cao. Có thể hiểu 1 cách đơn giản như sau:
	- Các Module không giao tiếp trực tiếp với nhau ,mà phải thông qua Interface. Module cấp thấp sẽ Implement interface, module cấp cao sẽ gọi module cấp thấp thông qua interface.
	- _Example:_ Để giao tiếp với database, ta có giao diện IDdatabase và các module cấp thấp như: XMLDatabase, SQLDatabase, MongoDatabase,. Module cấp cao là CustomerBusiness sẽ chỉ sử dụng giao diện IDdatabase.
	- Việc khởi tạo các module cấp thấp sẽ do DI Container thực hiện.
	- Việc module nào gắn với Interface nào sẽ được config trong code hoặc trong file XML.
	- DI được dùng để làm giảm sự phụ thuộc giữa các Module, dễ dàng hơn trong việc thay đổi module, bảo trì code và testing.
#### **1.4 Các phương pháp Inject:**
- Đưa _Dependency_ vào trong lớp khởi tạo của đối tượng.
- **Inject** thông qua **setter**. _Ví dụ:_ lớp A gọi thuộc tính có kiểu là lớp B và gán thuộc tính đó với đối tượng có kiểu là lớp B.
#### **1.5 Thư viện Dependency inject:**
- Có đối tượng **DI container (Class ServiceCollection)** đối tượng này có khả năng:
	- đăng ký các dịch vụ (Lớp) vào bên trong nó.
	- hỗ trợ lấy dịch vụ (Đối tượng). khi nó lấy ra đối tượng thì nó sẽ khởi tạo đối tượng đó. Và nếu đối tượng đó cần **Dependency** mà **Dependency** chưa có thì nó sẽ tự động tạo ra **Dependency** đó và cũng tự Inject vào dịch vụ đang dùng đến
- Trước tiên thì đảm bảo rằng đã tính hợp Package [Microsoft.Extensions.DependencyInjection](https://www.nuget.org/packages/Microsoft.Extensions.DependencyInjection/) vào dự án:
```
dotnet add package Microsoft.Extensions.DependencyInjection
```

- Kiểu **ServiceProvider** cho phép lấy ra các dịch vụ đã được đăng ký **ServiceProvider**.
- **ServiceLifeTime - Vòng đời của dịch vụ:** thì khi đăng ký các dịch vụ vào trong **ServiceCollection** thì có đối tượng **ServiceDescriptor** chứa thông tin về dịch vụ đó, căn cứ vào **ServiceDescriptor** để xác định dịch vụ đó sẽ tồn tại được trong bao lâu. LifeTime có kiểu **ServiceLifeTime (Kiểu enum)** có các giá trị cụ thể như sau:
	- **Singleton 0:** chỉ có một phiên bản duy nhất cho dịch vụ đó. Tức là một đối tượng duy nhất cho dịch vụ đó cho suốt vòng đời **ServiceProvider**
	- **Scoped 1:** Trong mỗi phạm vi được tạo ra bởi **ServiceProvider** thì có mỗi phiên bản cho dịch vụ đó
	- **Transient 2:** Mỗi lần lấy dịch vụ đó ra thì đối tượng mới của dịch vụ đó được tạo ra
- Khi chúng ta có đối tượng **ServiceCollection** thì các dịch vụ nó đăng ký vào thông qua các phương thức:
	- **_AddSingleton<ServiceType, ImplementationType>()_**:
	- **_AddSingleton< ServiceType>()_**:
	- **_AddTransient<ServiceType, ImplementationType>() Hoặc AddTransient< ServiceType>()_**:
	- **_AddScoped<ServiceType, ImplementationType>()_**:
	- **_BuildServiceProvider()_**:
- Trong các phương thức đăng ký dịch vụ thì phải chỉ ra kiểu dịch vụ và kiểu mà dịch vụ đó sẽ triển khai.
- Sau khi đăng ký tất cả các dịch vụ rồi thì gọi phương thức **_BuildServiceProvider()_** để nó sinh ra một lớp **ServiceProvider**.
- Trong lớp **ServiceProvider** có các phương thức như:
	- **_GetService< ServiceType>():_** dùng để lấy ra đối tượng dịch vụ thông qua tên dịch vụ đã đăng ký
	- **_GetRequiredService(ServiceType):_**
	- **_CreateScope():_** gọi phương thức này để tạo ra một cái phạm vi
#### **1.6 Sử dụng Delegate ,Factory đăng ký dịch vụ:**
- Các phương thức để đăng dịch vụ vào **ServiceCollection** như **AddSingleton, AddSingleton, AddTransient** còn có phiên bản _(nạp chồng)_ nó nhận tham số là _delegate_ trả về đối tượng dịch vụ có kiểu **ImplementationType.** 
- Khai báo một phương thức tĩnh để cung cấp các cơ chế cho việc tạo ra đối tượng mong muốn gọi là **Factory**.
#### **1.7 Sử dụng IOptions Inject thông số cho dịch vụ:**
- Khi một dịch vụ nào đó được đăng ký trong DI, nếu nó cần các tham số để khởi tạo thì trong **ServiceCollection** có hỗ trợ sử dụng giao diện **IOptions** để mà phân tách giữa các dịch vụ và các thiết lập truyền vào dể khởi tạo dịch vụ.
- Đầu tiên cần phải thêm gói sao:
```
dotnet add package Microsoft.Extensions.Options
```
- Sử dụng namespace:
```
using Microsoft.Extensions.Options;
```
- Khi tham số có kiểu IOptions, thì nó được **inject** từ một tập các IOptions riêng biệt với các dịch vụ và các IOptions nạp vào **ServiceCollection** bằng phương thức **Configure**.
- Để đăng ký 1 option thì cần phải sử dụng phương thức **Configure<tên lớp thiết lập khởi tạo>();** trong đối tượng kiểu **ServiceCollection**.

#### **1.8 Nạp cấu hình File vào ứng dụng:**
- Các giá trị dữ liệu trong lớp cấu hình có thể được lưu trữ bên các tệp tin bên ngoài có đuôi: .XML, Json, INI,...
- Đầu tiên thì cần thêm gói sau:
```
dotnet add package Microsoft.Extensions.Configuration
dotnet add package Microsoft.Extensions.Options.ConfigurationExtensions
```
- Sau đó muốn dùng định dạng file nào để đọc giá trị thì thêm gói đó, chẳng hạn như: 
```
dotnet add package Microsoft.Extensions.Configuration.Json
dotnet add package Microsoft.Extensions.Configuration.Ini
dotnet add package Microsoft.Extensions.Configuration.Xml
```
- Sử dụng namespace với file kiểu json:
```
using Microsoft.Extensions.Configuration;
using Microsoft.Extensions.Configuration.Json;
```
##### **ConfigurationBuilder:**
- Lớp **ConfigurationBuilder**, giúp nạp các cấu hình lưu trong file config, từ đó build ra đối tượng _ConfigurationRoot_ , đối tượng nảy truy cập đến các cấu hình bằng chỉ toán tử chỉ số [key].
- Đối tượng này có khả năng đọc file cấu hình
- Phương thức **SetBasePath()** trong đối tượng **ConfigurationBuilder** dùng để chỉ ra đường dẫn.
- Phương thức **AddJsonFile()** thì dùng để đọc các giá trị được ghi sẳn trong file json vào đối tượng.
- Phương thức **Build** giúp trả về một đối tượng **ConfigurationRoot
- Để đọc giá trị trong file cấu hình thì dùng phương thức **GetSection(Tên key muốn đọc)**

### **Dependency Inject trong ứng dụng ASP.NET:**



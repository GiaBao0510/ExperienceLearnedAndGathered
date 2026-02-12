
## **SOLID là gì?**

**SOLID** trong  [lập trình hướng đối tượng](https://200lab.io/blog/oop-la-gi/) là tập hợp các nguyên lý thiết kế phần mềm nhằm giúp lập trình viên có thể tạo ra các ==hệ thống có thể dễ dàng bảo trì, mở rộng và có tính ổn định==. Các nguyên lý này bao gồm: Single Repository, Open/Closed, Liskov Substitution, Interface Segregation và Dependency Inversion.

---
## **Tổng quan về SOLID:**

![](https://images.viblo.asia/fa9b80a1-9398-44ae-b50a-71e665b63ab5.png)

**SOLID** là 5 nguyên tắc cơ bản, giúp xây dựng kiến trúc phần mềm tốt. Bạn có có thể thấy tất cả các Design Pattern điều dựa trên các nguyên tắc này. SOLID được  ghép lại từ 5 chữ viết tắt đầu tiên của 5 nguyên tắt này:
###### 1.**S**ingle resposibility principle (SRP)
###### 2.**O**pen/Closed principle (OCP)
###### 3.**L**iskov substitution principle (LSP)
###### 4.**I**nterface segregation principle (ISP)
###### 5.**D**ependency injection principle (DIP)

---
## **Các nguyên lý trong SOLID:**

#### 1.**S**ingle resposibility principle (SRP)

Đây là nguyên lý đầu tiên của **SOLID** trong lập trình hướng đối tượng, nghĩa là:
- **Mỗi lớp (Class)** trong chương trình chỉ nên ==có một trách nhiệm duy nhất==.
- Nói cách khác là, một lớp chỉ nên có ==một lý do duy nhất== để thay đổi.


![](https://images.viblo.asia/7ae08c24-26d9-447a-b460-4a23f8c85cf0.png)

- Khi một **class có nhiều trách nhiệm**, nó sẽ trở nên ==khó quản lý hơn, khó kiểm tra lỗi, mở rộng và bảo trì==. Nếu cần thay đổi logic tính toán hoặc định dạng, thì có thể sẽ phải thay đổi toàn bộ lớp này, điều này làm cho mã nguồn trở nên phức tạp và khó quản lý ➡️Vì vậy nên để **mỗi Class chỉ nên đảm nhận một nhiệm vụ duy nhất**. Điều này giúp mã nguồn trở nên đơn giản hơn, dễ bào trì và mở rộng. 

***Ví dụ bằng code:***

- **Không tuân theo SRP:**
```csharp
namespace SRP{
	
	public class Employee
	{
		public int Employee_id {set; get;}
		public string Employee_name {set; get;}

		// This method used to insert into employee table
		public bool AddEmployee(Employee e){

			// Insert into employee table.
			return true;
		}

		// Method to generate report
		public void GenerateReport(Employee e){
			// Report generation with employee data using crystal report.
		}
	}
}
```

Thấy rằng lớp Class `Employee` này chịu 2 trách nhiệm. Một là thao tác thêm thông tin `Employee` mới và bảng và một là tạo báo cáo. Lớp `Employee` không nên đảm nhận nhiệm vụ tạo báo cáo vì giả sử một ngày khách hàng sẽ yêu cầu phải tạo ra báo cáo với các định dạng như: `PDF`, `Excel`,..., thì class này sẽ thay đổi cho phù hợp

- **Tuân theo SRP:**
```csharp
	public class ReportGeneration{
		// Method to generate report
		public void GenerateReport(Employee e){
			// Report reneration with employee data.
		}
	}
```

Vì thế tuân theo SRP, là chỉ một class đảm nhận một trách nhiệm duy nhất, nên sẽ tạo lớp `ReportGeneration` để chịu trách nhiệm cho việc báo cáo, mà không gây ảnh hướng đến lớp `Employee`

***Ví dụ thực tế:***
> Nếu bạn thuê một người thợ làm tất cả mọi việc—từ xây tường, lắp điện, sửa ống nước đến sơn nhà—thì khi cần sửa chữa một phần nào đó, ví dụ như hệ thống ống nước, bạn sẽ phải gọi người này. Điều này có thể gây rắc rối vì việc sửa chữa có thể ảnh hưởng đến các công việc khác mà anh ta đã làm.
>Ngược lại, nếu bạn có từng nhóm thợ chuyên trách—thợ nền chỉ xây tường, thợ điện chỉ lắp đặt hệ thống điện, thợ ống nước chỉ lo về đường ống, và thợ sơn chỉ chuyên sơn nhà—thì khi cần sửa ống nước, bạn chỉ cần gọi thợ ống nước mà không ảnh hưởng đến các phần khác. Điều này giúp công việc sửa chữa dễ dàng hơn và quản lý ngôi nhà hiệu quả hơn.

#### 2.**O**pen/Closed principle (OCP)

Nguyên tắc này đề cập đến cách thiết kế phần mềm sao cho nó có thể mở rộng bằng cách thêm tính năng mới mà không cần phải sửa đổi mã nguồn hiện có. Điều này giúp cho mã nguồn có thể ổn định hơn, tránh việc sửa đổi dẫn đến lỗi không mong muốn.

![](https://statics.cdn.200lab.io/2024/08/solid-la-gi-single-open-closed-principle.png?width=1200)

***Ví dụ bằng code:***

- **Không tuân theo OCPSRP:**
```csharp
public class ReportGeneration{

	// Report type
	public string ReportType {get; set;}

	// Method to generate report
	public void GenerateReport(Employee e){

		if(ReportType == "Excel"){
			// Report generation with employee data in Excel.
		}

		if(ReportType == "PDF"){
			// Report generation with employee data in PDF.
		}
	}
}
```

Thấy rằng phương phức `GenerateReport` tại lớp `ReportGeneration` có quá nhiều mệnh đề `if`. Nếu như khách hàng muốn tạo bảng báo cáo dạng `.doc` thì sẽ viết thêm 1 lần if nữa.

- **Tuân theo OCP:**
```csharp
public interface IReportGeneration{
	// Method to generate report
	public void GenerateReport(Employee e){}
} 

// Class to generate PDF report
public class PDFReportGeneration: IReportGeneration{
	
	public void GenerateReport(Employee e){
		// Generate PDF report.
	}
}

// Class to generate Excel report
public class ExcelReportGeneration: IReportGeneration{
	
	public void GenerateReport(Employee e){
		// Generate Excel report.
	}
}
```

- Nếu muốn đưa ra một định dạng báo cáo các, thì chỉ cần kế thừa `interface` `IReportGeneration`. Vì `IReportGeneration` là `interface` nên nó không có triển khai chi tiết trong các phương thức.

***Ví dụ thực tế:***
> Ban đầu, chiếc xe chỉ sử dụng động cơ xăng. Sau một thời gian, bạn muốn bổ sung tùy chọn động cơ điện. Nếu thiết kế ban đầu không linh hoạt, bạn có thể phải tháo rời và thay đổi gần như toàn bộ chiếc xe để lắp đặt động cơ điện. Điều này sẽ rất phức tạp, tốn nhiều thời gian và chi phí.
> 
>Ngược lại, nếu ngay từ đầu bạn thiết kế xe theo cách cho phép động cơ dễ dàng thay thế và mở rộng, thì việc chuyển từ động cơ xăng sang động cơ điện sẽ đơn giản hơn rất nhiều mà không cần thay đổi toàn bộ kết cấu xe. Đây chính là nguyên tắc **Open/Closed Principle (OCP)** – thiết kế mở rộng dễ dàng mà không cần chỉnh sửa phần cốt lõi.

#### 3.**L**iskov substitution principle (LSP)

Nguyên lý này được lấy theo tên của **nhà khoa học máy tính Barcara Liskov** và nó nói rằng: Các đối tượng của một lớp con phải có thể thay thế cho các đối tượng của lớp cha mà không làm thay đổi đi tính đúng đắng của chương trình.
- Class con **không nên phá vỡ các định nghĩa và hành vi** của Class cha

![](https://statics.cdn.200lab.io/2024/08/solid-la-gi-liskov-substitution-principle.png?width=1200)

***Ví dụ bằng code:***

- **Không tuân theo LSP:**
 
Giả sử có `Class Employee` là lớp cha của `Class ContractualEmployee` và `Class CasualEmployee`, như sau

```csharp
class abstract class Employee{
	
	public virtual string GetProjectDetails(int employee_id){
		return "Base Project";
	}
	
	public virtual string GetEmployeeDetails(int employee_id){
		return "Base employee";
	}
}

class CasualEmployee: Employee{
	
	public override string GetProjectDetails(int employee_id){
		return "chill Project";
	}

	public override string GetEmployeeDetails(int employee_id){
		return "chill employee";
	}
}

class ContractualEmployee: Employee{
	
	public override string GetProjectDetails(int employee_id){
		return "chill Project";
	}

	// May be for contractual employee we do not need to store the details into database.
	public override string GetEmployeeDetails(int employee_id){
		throw new NotImplementedException();
	}
}
```

Hãy thử đoạn code dưới đây và nó đã vi phạm nguyên tắc
```csharp
List<Employee> EmployeeList = new List<Employee>();

EmployeeList.Add(new ContractualEmployee());
EmployeeList.Add(new CasualEmployee());

foreach(Employee e in EmployeeList)
	e.GetEmployeeDetails(119);
```

Sau khi chạy sẽ thấy với `Contractual Employee`, nó sẽ ném ra một exception khi method `GetEmployeeDetails(int employee_id)` chưa được triển khai, và điều này vi phạm LSP.

- **Tuân theo LSP:**

Giải pháp là tác chúng thành 2 interface khác nhau. Một là Iproject, hai là Iemployee và triển khai theo từng type khác nhau.

```csharp
public interface IEmployee{
	string GetEmployeeDetails(int employee_id);
}

public interface IProject{
	string GetProjectDetails(int employee_id);
}
```

Giờ thì contractual employee sẽ triển khai IEmployee nhưng không cps IProject. Điều này giúp tuân theo nguyên tắc LSP.

***Ví dụ thực tế:***

>**Hãy tưởng tượng bạn có một chiếc xe đạp** mà bạn thường dùng để đi từ nhà đến trường. Một ngày nọ, bạn quyết định nâng cấp lên một chiếc xe đạp điện. Bạn mong muốn rằng, dù thay đổi phương tiện, bạn vẫn có thể tiếp tục đi học như bình thường mà không gặp trở ngại gì.
>Tương tự, trong lập trình, nếu một đối tượng của lớp con thay thế cho một đối tượng của lớp cha, chương trình vẫn phải hoạt động đúng như trước. Nếu việc thay thế này gây lỗi hoặc làm thay đổi hành vi mong đợi, thì bạn đã vi phạm **nguyên lý Liskov Substitution Principle (LSP)**.

#### 4.**I**nterface segregation principle (ISP)

Với nguyên tắc này giúp ==đảm bảo rằng các lớp không bị ép buộc phải triển khai các phương thức mà chúng không sử dụng==. Nói cách khác là, ==thay vì tạo một giao diện lớn (fat interface)== với nhiều phương thức mà không phải tất cả các lớp điều cần, vậy nên ==tách nó thành các interface nhỏ hơn==, ==mỗi interface chỉ có những phương thức liên quan đến một phạm vi cụ thể.==

![](https://statics.cdn.200lab.io/2024/08/solid-la-gi-interface-segregation-principle.png?width=800)
- Thay vì dùng interface lớn, ta nên tách thành nhiều interface nhỏ, với mục đích khác nhau.

***Ví dụ bằng code:***

- **Không tuân theo ISP:**
Có những nhân viên chỉ cần chức năng export excel, có nhân viên cần cả 2 chức năng. Lúc này nếu gom chung lại thì ta phải implement lại tất cả phương thức

```csharp
public interface IExportReport{
	void ExportExcelReport();
	void ExportPDFReport();
}
```

- **Tuân theo ISP:**
Tách nhỏ mỗi loại export ra thành các interface nhỏ.
```csharp
public interface IExportPdfFile{
	void ExportPDFReport();
}

public interface IExportExcelFile{
	void ExportExcelReport();
}
```

***Ví dụ thực tế:***

>Hãy tưởng tượng bạn làm việc trong một công ty với nhiệm vụ chính là xử lý đơn hàng. Tuy nhiên, công ty lại yêu cầu bạn phải kiêm luôn các công việc như quản lý tài chính, điều hành nhân sự và phát triển sản phẩm—những nhiệm vụ không liên quan đến công việc chính của bạn. Điều này sẽ khiến bạn quá tải, mất tập trung và không thể hoàn thành tốt nhiệm vụ của mình.

>Tương tự, trong lập trình, nếu một lớp bị ép buộc triển khai quá nhiều phương thức không liên quan đến nhiệm vụ chính của nó, mã nguồn sẽ trở nên phức tạp, khó bảo trì và khó mở rộng. Để tránh điều này, **nguyên lý Interface Segregation (ISP)** khuyến khích chia nhỏ giao diện để mỗi lớp chỉ cần thực hiện đúng trách nhiệm của mình.

#### 5.**D**ependency injection principle (DIP)

Đây là nguyên tắc cuối cùng trong **SOLID**, nó nhấn mạnh việc ==giảm sự phụ thuộc giữa các module cấp cao và module cấp thấp== trong hệ thống bằng cách sử dụng các `abstraction`(lớp trừu tượng hoặc `interface`) thay vì các concretion (cụ thể, chi tiết triển khai).
- Nguyên tắc này nói rằng **không nên viết code gắn chặt với nhau**. Vì việc này sẽ **gây ra sự cực kỳ khó khăn trong việc bảo trì** khi ứng dụng lớn dần.
- Cố gắng viết các class ít phụ thuộc nhất có thể.

![](https://statics.cdn.200lab.io/2024/08/solid-la-gi-dependency-inversion-principle.png?width=800)

***Ví dụ bằng code:***

- **Không tuân theo DIP:**

Giả sử chúng ta có một hệ thống thông báo sau khi lưu vài thông tin vào trong DB:

```csharp
public class Email{
	public void SendEmail{
		//Code to send email
	}
}

public class Notification{
	
	private Email _email;
	
	public Notification(){
		_email = new Email();
	}
	public void PromotionalNotification(){
		_email.SendEmail();
	}
}
```

Giờ `class Notification `hoàn toàn phụ thuộc vào `class Email`, vì nó chỉ gửi một loại thông báo. Nếu mà khách hàng yêu cầu gửi thông báo như `SMS` chẳng hạn? thì chúng ta cũng phải thay đổi thông báo? Để có thể giúp nó giảm phụ thuộc vào nhau thông qua code sau:

```csharp
public interface IMessenger{
	void SendMessage();
}

public class Email : IMessenger{
	void SendMessage(){
		//code to send email
	}
}

public class SMS : IMessenger{
	void SendMessage(){
		//code to send SMS
	}
}

public class Notification{
	private IMessenger _iMessenger;
	
	public Notification(){
		_iMessenger = new Email();
	}
	public void DoNotify(){
		_iMessenger.SendMessage();
	}
}
```

Tuy nhiện `class Notification` vẫn phụ thuộc vào `Email class`. Nhưng đây là sử dụng **dependency injection** để làm cho chúng **giảm sự phụ thuộc**. Có 3 loại DI, `Contructor Injection,` `Property Injection` và` Method Injection`. 
- **Tuân theo DIP:**
##### Constructor Injection
```csharp
public class Notification{
	private IMessenger _iMessenger;
	
	public Notification(IMessenger iMessenger){
		_iMessenger = iMessenger;
	}
	public void DoNotify(){
		_iMessenger.SendMessage();
	}
}
```
##### Property Injection
```csharp
public class Notification{
	private IMessenger _iMessenger;

	public Notification(){}
	public IMessenger MessageService{
		private get;
		set { _iMessenger = value; }
	}
	public void DoNotify(){
		_iMessenger.SendMessage();
	}
}
```
##### Method Injection
```csharp
public class Notification{

	public void DoNotify(IMessenger pMessenger){
		pMessenger.SendMessage();
	}
}
```

***Ví dụ thực tế:***
> Hãy tưởng tượng bạn đang xây dựng một ngôi nhà. Trong nhà có một hệ thống điện (module cấp cao) và nhiều loại bóng đèn khác nhau (module cấp thấp). Nếu hệ thống điện chỉ hỗ trợ một loại bóng đèn cụ thể, thì mỗi khi bạn muốn thay đổi bóng đèn, bạn sẽ phải sửa lại toàn bộ hệ thống điện—việc này rất phức tạp và tốn kém.
>
>Thay vào đó, nếu hệ thống điện được thiết kế để hoạt động với bất kỳ loại bóng đèn nào thông qua một **ổ cắm tiêu chuẩn** (một dạng **abstraction**), thì bạn có thể dễ dàng thay bóng đèn mà không cần thay đổi hệ thống điện.
>
>Tương tự trong lập trình, **nguyên lý Dependency Inversion Principle (DIP)** khuyến khích module cấp cao không phụ thuộc trực tiếp vào module cấp thấp mà thay vào đó, cả hai nên phụ thuộc vào một abstraction. Điều này giúp hệ thống linh hoạt hơn, dễ mở rộng và bảo trì.

---
## **Kết luận:** 

Mỗi nguyên lý trong **SOLID** từ **SRP, OCP, LSP, ISP đến DIP** đều hướng đến mục tiêu giúp mã nguồn trở nên **rõ ràng, ít lỗi và linh hoạt hơn**. Khi được áp dụng đúng cách, SOLID không chỉ nâng cao chất lượng mã mà còn giúp phần mềm dễ dàng thích nghi với những thay đổi trong tương lai. Điều này không chỉ **giảm thiểu rủi ro** mà còn **tăng hiệu suất làm việc** của nhóm phát triển, giúp dự án phát triển bền vững và dễ bảo trì hơn.
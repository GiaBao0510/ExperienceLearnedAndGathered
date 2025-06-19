![](https://images.viblo.asia/ead224e3-b111-4052-a80c-23e152e1877a.png)

### **Buider Design Pattern là gì?**

![](https://hanam88.com/images/posts/113102-09112022-Builder-Design-Pattern-Ha-nam-88.jpg)

**Builder Pattern** là một trong những **Creational pattern** - những mẫu thiết kế cho việc khởi tạo đối tượng lớp.

**Builder Pattern** được tạo ra ==để xây dựng một đối tượng phức tạp bằng cách sử dụng các đối tượng đơn giản và sử dụng tiếp cận từng bước==, việc xây dựng các đối tượng độc lập với các đối tượng khác.

---
### **Tại sao cần sử dụng Builder Pattern?**

##### **Tạo object phức tạp với nhiều tham số:**
- tránh constructor với nhiều parameter.
- Cho phép tạo object theo từng bước.
- Dễ đọc và maintain code.
##### **Tạo các biến thể khác nhau của cùng một object:**
- Có thể tạo nhiều loại configuration khác nhau.
- Flexible trong việc thiết lập thuộc tính.
##### **Immutable object:**
- Object được tạo ra không thể thay đổi sau khi build.
- Thread-safe

---
### **Kiến trúc**

![](https://freetuts.net/upload/tut_post/images/2022/06/14/5620/factory-03.png)

Chúng ta sẽ thấy **Builder Pattern** sẽ có 4 thành phần chính:
- **Builder:**  Là một **abstract** hoặc **Interface** xác định tất cả các bước được sử dụng để xây dựng một sản phẩm.
- **ConcreteBuilder:** Lớp này sẽ triển khai giao diện **Builder**. Chịu trách nhiệm cho việc xây dựng và lắp ráp các bộ phận riêng của sản phẩm.
- **Director:** **Class Director** lấy các quy trình đơn lẻ từ **Builder** và xác định trình tự để xây dụng sản phẩm
- **Product:** lớp này xác định các bộ phận khác nhau để tạo ra sản phẩm.

---
### **Ưu & nhược điểm**

##### **Ưu điểm:**
- Xây dựng các đối tượng theo từng bước, trì hoãn các bước xây dựng hoặc chạy các bước một cách đệ quy.
- Có thể sử dụng lại cùng một Construction Code khi xây dựng các thể hiện khác nhau của sản phẩm.
- Nguyên tắc Trách Nhiệm Đơn lẻ - SRP. Có thể tách biệt Contruction Code phức tạp khỏi Business Logic Layer của sản phẩm.
- Cho phép bạn thay đổi các thể hiện khác nhau của từng sản phẩm.
- Tính đóng gói code cho construction.
- Cung cấp khả năng kiểm soát các bước quy trình Construction.

##### **Nhược điểm:**
- Độ phức tạp tổng thể của mã tăng lên vì cần xây dựng nhiều class mới.
- Mỗi ConcreteBuilder riêng biệt phải được tạo cho từng loại sản phẩm.
- Các lớp Builder phải có thể thay thế được.

---
### **Khi nào sử dụng**

Builder pattern được sử dụng khi:
- Sử dụng **Builder Pattern** để tránh việc "**telescopic constructor**" (Gọi là **telescopic constructor** là khi một class chứa nhiều constructor với nhiều tham số khác nhau sẽ gây ra khó khăn cho lập trình viên để phải nhớ và sử dụng cho đúng). **Builder Pattern** cho phép xây dựng object từng bước, chỉ sử dụng những bước bạn thực sự cần. Sau khi triển khai pattern, bạn không phải nhồi nhét hàng tá tham số vào các constructor của mình nữa.
- Sử dụng **Builder Pattern** là khi muốn code của mình có thể tạo các thể hiện khác nhau của một số sản phẩm. **Builder Pattern** có thể được áp dụng khi việc xây dựng các bản trình bày khác nhau của sản phẩm bao gồm các bước tương tự chỉ khác nhau về chi tiết.
- Sử dụng **Builder Pattern** để tạo ra cây *Composite* hoặc các đối tượng phức tạp khác. **Builder Pattern** cho phép bạn tạo sản phẩm theo từng bước. Bạn có thể trì hoãn việc thực hiện một số bước mà không làm hỏng sản phẩm cuối cùng. Thậm chí bạn có thể gọi đệ quy các bước, điều này rất hữu ích khi bạn cần xây dựng một cây đối tượng. Một Builder không để lộ sản phẩm chưa hoàn thành khi đang chạy các bước xây dụng. Điều này ngăn cho client code tìm nạp kết quả không đầy đủ.

---
### **Ví dụ minh họa**

***Ví dụ:*** Về việc tạo báo cáo với 3 phần nội dung: **ReportHeader**, **ContentReport** và **ReportFooter**. Bản báo cáo sẽ được trình bày dưới dạng là Excel và PDF.
```csharp
/// <summary>
/// Bước 1 (Product): Tạo Class Report bao gồm: ReportType, ReportHeader,
/// ContentReport và ReportFooter.
/// </summary>
public class Report
{
    public string ReportType { get; set; }
    public string ReportHeader { get; set; }
    public string ReportContent { get; set; }
    public string ReportFooter { get; set; }
    public void DisplayReport()
    {
        Console.WriteLine($"Report Type: {ReportType}");
        Console.WriteLine($"Report Header: {ReportHeader}");
        Console.WriteLine($"Report Content: {ReportContent}");
        Console.WriteLine($"Report Footer: {ReportFooter}");
    }
}

/// <summary>
/// Bước 2: Tạo class ReportBuilder
/// </summary>
public abstract class ReportBuilder
{
    protected Report reportOpt;
    public abstract void SetUpType();
    public abstract void SetUpHeader();
    public abstract void SetUpContent();
    public abstract void SetUpFooter();
  
    public void CreateNewReport()
		=> reportOpt = new Report();

    public Report GetReport()
    {
        return reportOpt;
    }
}

/// <summary>
/// Bước 3: Vì là dang hai dạng báo cáo Excel và PDF
/// Nên cần tạo hai class con kế thừa từ ReportBuilder.
/// </summary>
public class ExcelReport : ReportBuilder
{
    public override void SetUpType()
        => reportOpt.ReportType = "Excel";

    public override void SetUpHeader()
        => reportOpt.ReportHeader = "Excel Report Header";

    public override void SetUpContent()
        => reportOpt.ReportContent = "Excel Report Content";

    public override void SetUpFooter()
        => reportOpt.ReportFooter = "Excel Report Footer";
}

public class PDFReport : ReportBuilder
{
    public override void SetUpType()
        => reportOpt.ReportType = "PDF";

    public override void SetUpHeader()
        => reportOpt.ReportHeader = "PDF Report Header";

    public override void SetUpContent()
        => reportOpt.ReportContent = "PDF Report Content";

    public override void SetUpFooter()
        => reportOpt.ReportFooter = "PDF Report Footer";
}

/// <summary>
/// Tạo Class ReportDirector có một phương thức chung là MakeReport()
/// Có tham số đầu vào là ReportBuilder.
/// </summary>
public class ReportDirector
{
    public Report MakeReport(ReportBuilder reportBuilder)
    {
        reportBuilder.CreateNewReport();
        reportBuilder.SetUpType();
        reportBuilder.SetUpHeader();
        reportBuilder.SetUpContent();
        reportBuilder.SetUpFooter();
        return reportBuilder.GetReport();
    }
}

public class Solution
{
    public static void Main(string[] args)
    {
        Report report;
        ReportDirector reportDirector = new ReportDirector();
  
        // Tạo báo cáo Excel
        ExcelReport excelReport = new ExcelReport();
        report = reportDirector.MakeReport(excelReport);
        report.DisplayReport();
        Console.WriteLine("\n---------------------------------\n");

        //Tạo báo cáo PDF
        PDFReport pDFReport = new PDFReport();
        report = reportDirector.MakeReport(pDFReport);
        report.DisplayReport();
    }
}
```

- Kết quả sau khi chạy:
```powershell
Report Type: Excel
Report Header: Excel Report Header
Report Content: Excel Report Content
Report Footer: Excel Report Footer

---------------------------------

Report Type: PDF
Report Header: PDF Report Header
Report Content: PDF Report Content
Report Footer: PDF Report Footer
```


***Ví dụ:*** Một shop cần sản xuất các sản phẩm xe cộ gồm:  **MotorCycle**, **Car**, **Scooter**. Mỗi loại gồm 4 thành phần là **Frame**, **Engine**, **Wheels**, **Doors.** Vậy áp dụng Builder Pattern sẽ được các thành phần sau:

- **Product:** (Vehicle)
- **Builder:** (VehicleBuilder)
- **ConcreteBuilder:** (MotorCycleBuilder, CarBuilder, ScooterBuilder)
- **Director:** (Shop)
```csharp
/// <summary>
/// Lớp Vehicle đóng vai trò là Product
/// </summary>
class Vehicle
{
    //Loại xe
    private string _vehicleType;

    //Khai báo dictionary chứa các thành phần của 1 chiếc xe (Frame, Engine, Wheels, Doors)
    private Dictionary<string, string> _parts = new Dictionary<string, string>();

    //Hàm khởi tạo
    public Vehicle(string vehicleType)
    {
        _vehicleType = vehicleType;
    }

    //Chỉ mục
    public string this[string key]
    {
        get { return _parts[key]; }
        set { _parts[key] = value; }
    }

    //Hàm hiển thị thông tin
    public void Show()
    {
       Console.WriteLine($"Vehicle Type: {_vehicleType}");
        Console.WriteLine($"Frame: {_parts["Frame"]}");
        Console.WriteLine($"Engine: {_parts["Engine"]}");
        Console.WriteLine($"Wheels: {_parts["Wheels"]}");
        Console.WriteLine($"Doors: {_parts["Doors"]}");
    }
}

/// <summary>
/// Lớp Abstract VehicleBuilder đóng vai trò là Builder
/// </summary>
abstract class VehicleBuilder
{

    //Khai báo trường kiểu vehicle
    protected Vehicle vehicle;

    //Định nghĩa thuộc tính Get để lấy vehicle
    public Vehicle Vehicle
    {
        get { return vehicle; }
    }

    //Khai báo các phương thức Abstract để build
    public abstract void BuildFrame();
    public abstract void BuildEngine();
    public abstract void BuildWheels();
    public abstract void BuildDoors();
}

/// <summary>
/// ConcreteBuilder cho xe ô tô
/// </summary>
class MotoCycleBuilder : VehicleBuilder
{
    public MotoCycleBuilder()
    {
        vehicle = new Vehicle("Motorcycle");
    }
  
    public override void BuildFrame()
        => vehicle["Frame"] = "Motorcycle Frame";

    public override void BuildEngine()
        => vehicle["Engine"] = "500cc";
  
    public override void BuildDoors()
        => vehicle["Doors"] = "0";

    public override void BuildWheels()
        => vehicle["Wheels"] = "2";
}

class CarBuilder : VehicleBuilder
{
    public CarBuilder()
    {
        vehicle = new Vehicle("Car");
    }

    public override void BuildFrame()
        => vehicle["Frame"] = "Car Frame";

    public override void BuildEngine()
        => vehicle["Engine"] = "2500cc";
  
    public override void BuildDoors()
        => vehicle["Doors"] = "4";

    public override void BuildWheels()
        => vehicle["Wheels"] = "4";
}

class ScooterBuilder : VehicleBuilder
{
    public ScooterBuilder()
    {
        vehicle = new Vehicle("Scooter");
    }

    public override void BuildFrame()
        => vehicle["Frame"] = "Scooter Frame";

    public override void BuildEngine()
        => vehicle["Engine"] = "50cc";
  
    public override void BuildDoors()
        => vehicle["Doors"] = "0";
  
    public override void BuildWheels()
        => vehicle["Wheels"] = "2";
}

/// <summary>
/// Lớp Shop đóng vai trò như Director
/// </summary>
class Shop
{
    public void Construct(VehicleBuilder vehicleBuilder)
    {

        vehicleBuilder.BuildFrame();
        vehicleBuilder.BuildEngine();
        vehicleBuilder.BuildWheels();
        vehicleBuilder.BuildDoors();
    }
}

public class Solution
{
    public static void Main(string[] args)
    {

        //Khai báo đối tượng Builder
        VehicleBuilder builder;

        //Tạo đối tượng Shop
        Shop shop = new Shop();

        //Tạo và hiển thị xe Moto
        builder = new MotoCycleBuilder();
        shop.Construct(builder);
        builder.Vehicle.Show();

        //Tạo và hiển thị xe Ô tô
        builder = new CarBuilder();
        shop.Construct(builder);
        builder.Vehicle.Show();

  

        //Tạo và hiển thị xe Scooter
        builder = new ScooterBuilder();
        shop.Construct(builder);
        builder.Vehicle.Show();
    }
}
```

- Kết quả sau khi chạy:
```powershell
Vehicle Type: Motorcycle
Frame: Motorcycle Frame
Engine: 500cc
Wheels: 2
Doors: 0
Vehicle Type: Car
Frame: Car Frame
Engine: 2500cc
Wheels: 4
Doors: 4
Vehicle Type: Scooter
Frame: Scooter Frame
Engine: 50cc
Wheels: 2
Doors: 0
```


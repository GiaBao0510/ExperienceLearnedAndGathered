![](https://images.viblo.asia/2f439de5-da4c-4c24-9d57-f02cede94c30.png)

### **Giới thiệu:**

Đây là loại **Design pattern** thuộc nhóm **Behavior Pattern.**
- **Bí danh:** Policy pattern.
- **Mục đích ra đời:** Định nghĩa một tập hợp các thuật toán giống nhau, encapsulate chúng và khiến chúng có thể thay thế cho nhau. Strategy làm cho thuộc toán độc lập khỏi client sử dụng nó.
- **Tần suất sử dụng:** Khá cao.

![](https://fxstudio.dev/wp-content/uploads/2024/05/strategy_pattern.webp)

---
### **Mục đích ra đời:**

**Giả sử như bạn đang xây dựng app giúp người dùng tìm đường đi trong thành phố.**

Ban đầu, app này chỉ hỗ trợ cho người đi bộ, nếu bạn đã viết thuận toán tìm đường đi cho người đi bộ và sử dụng nó trực tiếp ở những nơi cần thiết trong mã nguồn. Tuy nhiên, khi yêu cầu phát triển, ứng dụng cần hỗ trợ thêm tìm đường đi cho các phương tiện như: Xe hơi, xe đạp,...

Nếu tiếp tục chèn thêm thuật toán mới vào nơi đã có sẳn cho thuật toán cũ, bạn sẽ phải sửa đổi nhiều đoạn mã. Việc này khiến cho chương trình khó bảo trì và mở rộng. Đặc biệt là những phần đang hoạt động vốn đã ổn định.

Để giải quyết vấn đề này, đó là áp dụng **Strategy Pattern**:
$\to$ Thay vì viết "cứng" thuật toán trong mã, ta nên tách mỗi thuật toán ra thành một lớp riêng gọi là **strategy**.
$\to$ Những đoạn mã cần sử dụng thuật toán sẽ chỉ tương tác với các strategy này thông quan 1 **interface chung**, mà không cần quan tâm đến việc cụ thể là strategy nào đang được dùng.

Kết quả là:
- Phần mã sử dụng thuật toán và phần định nghĩa thuật toán phải **tách biệt rõ gàng**.
- Ta có thể dễ dang **thêm hoặc thay đổi thuật toán** mà không cần sửa phần còn lại của chương trình

---
### **Tính chất cơ bản**

- Strategy Pattern cho phép thay đổi hành vi của một đối tượng tại thời gian chạy (**runtime**) bằng cách thay đổi thuật toán mà nó sử dụng
- Điều này giống với mẫu Delegation, cả hai đều dựa trên **protocol** để tăng tính linh hoạt.
- Tuy nhiên, khác với Delegation, Strategy Pattern sử dụng một họ các đối tượng, và các thuật toán (strategies) có thể dễ dang thay đổi thời gian chạy, trong khi delegate thường được sử dụng cố định trong thời gian chạy

---
### **Kiến trúc:**

![](https://images.viblo.asia/9e6ad788-fd88-4935-a53c-c15d161e7331.png)

**Các thành phần trong mô hình:**
- *Context:* Class này sử dụng các **strategy object** và chỉ giao tiếp với các s**trategy objec**t thông qua **interface**.
- *Strategy:* Cung cấp một interface chung cho các *Context* giao tiếp với các strategy object.
- *Concrete Strategy:* Implement các thuật toán khác nhau cho từng Context.
- *Client:* Có trách nhiệm tạo ra các **strategy object** và truyền cho các Context sử dụng.

---
### **Khi nào sử dụng:**

- Khi có nhiều câu lệnh điều kiện xung quanh một số thuật toán được định nghĩa trước.
- Khi một class mà có nhiều hành vi liên quan đến nhau.
- Khi cần hoán đổi nhiều thuật toán trong quá trình chạy ứng dụng.
- Muốn sử các biến thể khác nhau của một xử lý trong một đối tượng và có thể chuyển đổi giữa các xử lý trong runtime.
- Khi có nhiều lớp tương đương và chỉ khác cách chúng thực thi một vài hành vi.
- Khi muốn tách biệt business logic của một lớp khỏi implementation details của các xử lý.

---
### **Ưu & nhược điểm:**

##### **Ưu điểm:**
- Không cần phải sử dụng lại câu lệnh **if...else** hay **switch...case**. Vì các câu lệnh điều kiện rất khó thực hiện init test. Chưa kể, càng mở rộng ứng dụng dụng, câu lệnh điều kiện càng phức tạp hơn.
- Dễ dàng mở rộng thêm các tính năng cho ứng dụng
- Dễ đọc code hơn. Vì tách ra các hàm riêng biệt và có thên hàm. Do đó, chỉ cần đọc tên hàm là biết hàm đó cần làm gì.
- Có thể thay thế các thuật toán linh hoạt với nhau.
- Tách biệt phần thuật toán khỏi phần sử dụng thuật toán.
- Có thể thay thế việc kế thừa bằng việc encapsulate thuật toán.
- Tăng tính open-closed: khi thay đổi code trong phần context.

##### **Nhược điểm:**
- Không nên áp dụng nếu chỉ có một vài xử lý và hiếm khi thay đổi.
- Client phải nhận biết được sự khác biệt giữa các strategy.

---
### **Ví dụ minh họa:**

**Ví dụ về chế độ chơi game**:
```csharp
using System;
using System.Collections.Generic;

// Tạo strategy
public interface IGameLevel
{
	void HandleLevel(int level);
}

//Tạo context
public class GameContext
{
	private IGameLevel GameLevel;

	public GameContext(IGameLevel GameLevel){
		this.GameLevel = GameLevel;
	}

	public void SetStrategy(IGameLevel GameLevel){
		this.GameLevel = GameLevel;
	}

	public void CurrentGameLevel(int level){
		this.GameLevel.HandleLevel(level);
	}
}

//Tạo các Concrete Strategy
public class LevelEasy: IGameLevel{
	public void HandleLevel(int level){
		Console.WriteLine($"Level Easy: {level}");
		// Cấu hình game: ít enemy, slow speed,...
	} 
}

public class LevelNormal: IGameLevel{
	public void HandleLevel(int level){
		Console.WriteLine($"Level Normal: {level}");
	} 
}

public class LevelDifficult: IGameLevel{
	public void HandleLevel(int level){
		Console.WriteLine($"Level Difficult: {level}");
	} 
}

public class Program{
	public static void Main(){
		var ctx = new GameContext(new LevelEasy());
		ctx.CurrentGameLevel(1);
		ctx.SetStrategy(new LevelNormal());
		ctx.CurrentGameLevel(2);
		ctx.SetStrategy(new LevelDifficult());
		ctx.CurrentGameLevel(3);
	}
}
```

**Ví dụ xuất file thiết kết (strategy pattern - DI)**: 
```csharp
using Microsoft.Extensions.Logging;
using Microsoft.Extensions.DependencyInjection;

public interface IExport
{
    void ExportFile(string fileName);
    string ExportType { get; }      //Xác định strategy
}


#region ==== Tạo các Concrete Strategy ====
public class ExportJPG : IExport
{
    private readonly ILogger<ExportJPG> _logger;

    public ExportJPG(ILogger<ExportJPG> logger)
        => _logger = logger;

    public string ExportType => "JPG";

    public void ExportFile(string fileName)
        => _logger.LogInformation($"Export file: {fileName}.JPG");
}

public class ExportPDF: IExport
{
    private readonly ILogger<ExportJPG> _logger;
  
    public ExportPDF(ILogger<ExportJPG> logger)
        => _logger = logger;

    public void ExportFile(string fileName)
    {
        _logger.LogInformation($"Export file: {fileName}.PDF");
    }

    public string ExportType => "PDF";
}

public class ExportPNG : IExport
{
    private readonly ILogger<ExportJPG> _logger;

    public ExportPNG(ILogger<ExportJPG> logger)
        => _logger = logger;

    public string ExportType => "PNG";


    public void ExportFile(string fileName)
    {
        _logger.LogInformation($"Export file: {fileName}.PNG");
    }
}
#endregion


#region ==== Strategy Factory để quản lý các strategy ====
public interface IExportStrategyFactory
{
    IExport GetExportStrategy(string exportType);
    IEnumerable<string> GetAvailableExportTypes();
}

public class ExportStrategyFactory : IExportStrategyFactory
{
    private readonly IEnumerable<IExport> _exportsStrategies;

    public ExportStrategyFactory(IEnumerable<IExport> exportsStrategies)
    {
        _exportsStrategies = exportsStrategies;
    }

    public IExport GetExportStrategy(string exportType)
    {
        var strategy = _exportsStrategies.FirstOrDefault(
            e => e.ExportType.Equals(exportType,StringComparison.OrdinalIgnoreCase)
        );

        return strategy ?? throw new 
	        ArgumentException($"Export type '{exportType}' is not supported.");
    }

    public IEnumerable<string> GetAvailableExportTypes()
    {
        return _exportsStrategies.Select(e => e.ExportType);
    }
}
#endregion

#region     ==== Export Service (Thay thể cho ExportContext) ====
public interface IExportService
{
    void ExportFile(string fileName, string exportType);
    void ExportMultipleFiles(string fileName, params string[] exportFile);
    IEnumerable<string> GetSupportedFormats();
}

public class ExportService : IExportService
{
    private readonly IExportStrategyFactory _strategyFactory;

    private readonly ILogger<ExportService> _logger;

    public ExportService(IExportStrategyFactory strategyFactory, ILogger<ExportService> logger)
    {
        _strategyFactory = strategyFactory;
        _logger = logger;
    }


    public void ExportFile(string fileName, string exportType)
    {
        try
        {
            var strategy = _strategyFactory.GetExportStrategy(exportType);

            _logger.LogInformation
	            ($"Exporting file: {fileName} with type: {exportType}");

            strategy.ExportFile(fileName);

            _logger.LogInformation
	            ($"Successfully exported file: {fileName} as {exportType}");
        }

        catch(Exception ex)
        {
            _logger.LogError($"Error exporting file: {fileName} with type: {exportType}. Exception: {ex.Message}");
            throw;
        }
    }

    public void ExportMultipleFiles(string fileName, params string[] exportTypes)
    {
        foreach( var exportType in exportTypes)
        {
            ExportFile(fileName, exportType);
        }
    }

    public IEnumerable<string> GetSupportedFormats()
    {
        return _strategyFactory.GetAvailableExportTypes();
    }
}
#endregion 

#region  ==== Dependency Injection Setup ====
public static class ServiceCollectionExtensions
{
    public static IServiceCollection AddExportServices(this IServiceCollection services)
    {
        services.AddTransient<IExport, ExportJPG>();
        services.AddTransient<IExport, ExportPDF>();
        services.AddTransient<IExport, ExportPNG>();

        services.AddSingleton<IExportStrategyFactory, ExportStrategyFactory>();

        services.AddScoped<IExportService, ExportService>();

        return services;
    }
}
#endregion
```

---

## Ví dụ minh họa bằng Go

> Các ví dụ gốc dùng C#. Dưới đây là bản chuyển sang Go idiomatic: dùng `interface` cho Strategy, `struct` cho Context và Concrete Strategy.

### Ví dụ 1: Chế độ chơi game (độ khó)

```go
package main

import "fmt"

// Strategy: interface chung cho các thuật toán xử lý level
type GameLevel interface {
	HandleLevel(level int)
}

// Context: sử dụng GameLevel strategy hiện tại, không quan tâm strategy cụ thể nào
type GameContext struct {
	level GameLevel
}

func NewGameContext(level GameLevel) *GameContext {
	return &GameContext{level: level}
}

func (c *GameContext) SetStrategy(level GameLevel) {
	c.level = level
}

func (c *GameContext) CurrentGameLevel(level int) {
	c.level.HandleLevel(level)
}

// ---- Concrete Strategy ----

type LevelEasy struct{}

func (LevelEasy) HandleLevel(level int) {
	fmt.Printf("Level Easy: %d\n", level)
	// Cau hinh game: it enemy, toc do cham,...
}

type LevelNormal struct{}

func (LevelNormal) HandleLevel(level int) {
	fmt.Printf("Level Normal: %d\n", level)
}

type LevelDifficult struct{}

func (LevelDifficult) HandleLevel(level int) {
	fmt.Printf("Level Difficult: %d\n", level)
}

func main() {
	ctx := NewGameContext(LevelEasy{})
	ctx.CurrentGameLevel(1)

	ctx.SetStrategy(LevelNormal{})
	ctx.CurrentGameLevel(2)

	ctx.SetStrategy(LevelDifficult{})
	ctx.CurrentGameLevel(3)
}
```


### Ví dụ 2: Xuất file theo nhiều định dạng (Strategy + Factory)

> Bản C# dùng `IServiceCollection`/`ILogger` của ASP.NET Core để đăng ký và tiêm (inject) các strategy qua DI container. Go không có DI container tiêu chuẩn đi kèm ngôn ngữ, nên cách idiomatic hơn là **wiring thủ công, tường minh** qua các hàm khởi tạo (`NewXxx`) — đây là cách làm phổ biến và được khuyến khích trong cộng đồng Go, giúp thấy rõ mọi phụ thuộc ngay tại nơi khởi tạo.

```go
package main

import (
	"fmt"
	"log"
	"strings"
)

// Strategy: interface chung cho các loại export
type Exporter interface {
	ExportType() string
	ExportFile(fileName string)
}

// ---- Concrete Strategy ----

type ExportJPG struct{ logger *log.Logger }

func (e ExportJPG) ExportType() string { return "JPG" }
func (e ExportJPG) ExportFile(fileName string) {
	e.logger.Printf("Export file: %s.JPG", fileName)
}

type ExportPDF struct{ logger *log.Logger }

func (e ExportPDF) ExportType() string { return "PDF" }
func (e ExportPDF) ExportFile(fileName string) {
	e.logger.Printf("Export file: %s.PDF", fileName)
}

type ExportPNG struct{ logger *log.Logger }

func (e ExportPNG) ExportType() string { return "PNG" }
func (e ExportPNG) ExportFile(fileName string) {
	e.logger.Printf("Export file: %s.PNG", fileName)
}

// ---- Strategy Factory: quản lý và tra cứu strategy theo tên ----

type ExportStrategyFactory struct {
	strategies map[string]Exporter
}

func NewExportStrategyFactory(logger *log.Logger) *ExportStrategyFactory {
	all := []Exporter{
		ExportJPG{logger: logger},
		ExportPDF{logger: logger},
		ExportPNG{logger: logger},
	}

	m := make(map[string]Exporter, len(all))
	for _, s := range all {
		m[strings.ToUpper(s.ExportType())] = s
	}
	return &ExportStrategyFactory{strategies: m}
}

func (f *ExportStrategyFactory) Get(exportType string) (Exporter, error) {
	s, ok := f.strategies[strings.ToUpper(exportType)]
	if !ok {
		return nil, fmt.Errorf("export type %q is not supported", exportType)
	}
	return s, nil
}

func (f *ExportStrategyFactory) AvailableTypes() []string {
	types := make([]string, 0, len(f.strategies))
	for t := range f.strategies {
		types = append(types, t)
	}
	return types
}

// ---- Export Service (tương đương ExportService trong bản C# gốc) ----

type ExportService struct {
	factory *ExportStrategyFactory
	logger  *log.Logger
}

func NewExportService(factory *ExportStrategyFactory, logger *log.Logger) *ExportService {
	return &ExportService{factory: factory, logger: logger}
}

func (s *ExportService) ExportFile(fileName, exportType string) error {
	strategy, err := s.factory.Get(exportType)
	if err != nil {
		s.logger.Printf("Error exporting file: %s with type: %s. %v", fileName, exportType, err)
		return err
	}
	s.logger.Printf("Exporting file: %s with type: %s", fileName, exportType)
	strategy.ExportFile(fileName)
	s.logger.Printf("Successfully exported file: %s as %s", fileName, exportType)
	return nil
}

func (s *ExportService) ExportMultipleFiles(fileName string, exportTypes ...string) []error {
	var errs []error
	for _, t := range exportTypes {
		if err := s.ExportFile(fileName, t); err != nil {
			errs = append(errs, err)
		}
	}
	return errs
}

func main() {
	logger := log.Default()
	factory := NewExportStrategyFactory(logger)
	service := NewExportService(factory, logger)

	service.ExportMultipleFiles("report", "JPG", "PDF", "PNG")
}
```
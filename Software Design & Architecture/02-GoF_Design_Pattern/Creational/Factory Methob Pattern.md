# Factory Method Pattern

> Tài liệu này giải thích mẫu thiết kế **Factory Method** — thuộc nhóm **Creational Pattern** (nhóm mẫu khởi tạo) — dành cho lập trình viên mới học, có liên hệ thực tế với **Golang/Backend**.

## Mục lục

1. [Factory Pattern là gì?](https://claude.ai/chat/de82c3b1-20af-49cd-a508-785244f774d4#1-factory-pattern-l%C3%A0-g%C3%AC)
2. [Vấn đề cần giải quyết](https://claude.ai/chat/de82c3b1-20af-49cd-a508-785244f774d4#2-v%E1%BA%A5n-%C4%91%E1%BB%81-c%E1%BA%A7n-gi%E1%BA%A3i-quy%E1%BA%BFt)
3. [Simple Factory — bước đệm trước Factory Method](https://claude.ai/chat/de82c3b1-20af-49cd-a508-785244f774d4#3-simple-factory--b%C6%B0%E1%BB%9Bc-%C4%91%E1%BB%87m-tr%C6%B0%E1%BB%9Bc-factory-method)
4. [Factory Method Pattern là gì?](https://claude.ai/chat/de82c3b1-20af-49cd-a508-785244f774d4#4-factory-method-pattern-l%C3%A0-g%C3%AC)
5. [Kiến trúc](https://claude.ai/chat/de82c3b1-20af-49cd-a508-785244f774d4#5-ki%E1%BA%BFn-tr%C3%BAc)
6. [Khi nào nên sử dụng](https://claude.ai/chat/de82c3b1-20af-49cd-a508-785244f774d4#6-khi-n%C3%A0o-n%C3%AAn-s%E1%BB%AD-d%E1%BB%A5ng)
7. [Ưu & nhược điểm](https://claude.ai/chat/de82c3b1-20af-49cd-a508-785244f774d4#7-%C6%B0u--nh%C6%B0%E1%BB%A3c-%C4%91i%E1%BB%83m)
8. [Ví dụ minh họa bằng Go](https://claude.ai/chat/de82c3b1-20af-49cd-a508-785244f774d4#8-v%C3%AD-d%E1%BB%A5-minh-h%E1%BB%8Da-b%E1%BA%B1ng-go)

---

## 1. Factory Pattern là gì?

"Factory Pattern" là cách gọi chung cho một nhóm giải pháp giúp **tách logic khởi tạo đối tượng ra khỏi nơi sử dụng đối tượng đó**. Thay vì rải rác các lệnh `new Xxx()` khắp nơi trong chương trình, ta gom việc quyết định "tạo cái gì, tạo như thế nào" vào một chỗ — thường gọi là "factory" (nhà máy/xưởng sản xuất).

Trong thực tế, người ta thường chia thành 3 cách tiếp cận:

- **Simple Factory:** một hàm/lớp duy nhất, nhận tham số rồi quyết định trả về đối tượng nào (thường dùng `if/switch`). Đây là cách đơn giản, dễ hiểu nhất, **không phải** một trong 23 mẫu thiết kế kinh điển của GoF (Gang of Four) nhưng lại là idiom rất phổ biến trong thực tế và thường được dùng làm bước đệm trước khi học Factory Method.
- **Factory Method:** một mẫu thiết kế chính thức thuộc bộ 23 mẫu GoF, dùng **tính đa hình (polymorphism)** — mỗi lớp con tự quyết định loại đối tượng nó tạo ra, thay vì dùng `switch/case`. Đây là trọng tâm của tài liệu này.
- **Abstract Factory:** cũng là một mẫu GoF chính thức, mở rộng ý tưởng của Factory Method để tạo ra **cả một họ đối tượng liên quan với nhau** (ví dụ: bộ UI theo từng theme — button, checkbox, scrollbar cùng phong cách). Đây là một mẫu khác, không được trình bày chi tiết trong tài liệu này.

Trong Factory Method Pattern, việc tạo đối tượng được thực hiện mà **không để lộ chi tiết cách tạo ra nó** cho phía người dùng (client). Người dùng chỉ cần gọi và nhận được đối tượng mong muốn, không cần quan tâm nó được khởi tạo như thế nào bên trong.

---

## 2. Vấn đề cần giải quyết

Hãy tưởng tượng bạn đang xây dựng một ứng dụng quản lý logistics. Phiên bản đầu tiên chỉ hỗ trợ vận tải bằng xe tải, nên phần lớn mã nguồn nằm trong lớp `Truck` (Xe tải).

Sau một thời gian, ứng dụng trở nên phổ biến. Mỗi ngày bạn nhận được nhiều yêu cầu tích hợp dịch vụ logistics đường biển vào ứng dụng.

![Factory Method problem](https://refactoring.guru/images/patterns/diagrams/factory-method/problem1-en-1.5x.png)

Đó là tin tốt về mặt kinh doanh, nhưng phần mã nguồn hiện tại đang bị ràng buộc chặt chẽ với lớp `Truck`. Thêm `Ship` (Tàu) vào ứng dụng đòi hỏi phải chỉnh sửa nhiều nơi trong hệ thống. Về sau, nếu có thêm một loại hình vận chuyển khác, bạn sẽ lại phải lặp lại toàn bộ những thay đổi đó.

Kết quả là mã nguồn dần chằng chịt các câu lệnh điều kiện (`if/switch`) để quyết định hành vi tùy theo loại đối tượng vận chuyển — càng thêm loại mới, đoạn điều kiện càng phình to và khó bảo trì.

---

## 3. Simple Factory — bước đệm trước Factory Method

Trước khi đến với Factory Method, hãy xem một cách tiếp cận đơn giản hơn — **Simple Factory** — để thấy rõ vấn đề mà Factory Method thực sự giải quyết.

Giả sử có 3 kiểu `Dog`, `Cat`, `Duck` cùng hiện thực một interface `Animal`. Khi cần tạo một `Animal` mà chưa biết trước là loại nào (tùy điều kiện cụ thể), cách viết trực quan nhất là:

![Simple factory conditional example](https://images.viblo.asia/1ea931d6-4432-4990-ab2a-94ea05b47913.png)

```csharp
IAnimal animal;

if (...) {
    animal = new Dog();
} else if (...) {
    animal = new Cat();
} else if (...) {
    animal = new Duck();
}
```

Cách này gây ra:
- **Trùng lặp logic khởi tạo** nếu cần dùng lại ở nhiều nơi trong chương trình.
- **Khó bảo trì** khi cần sửa đổi hoặc mở rộng thêm loại `Animal` mới.

![Simple factory solution](https://images.viblo.asia/02bc95d1-e578-4cd3-9853-1a35e0dd25e9.png)

**Giải pháp bước đầu — Simple Factory:** gom toàn bộ logic khởi tạo vào một hàm/lớp duy nhất, giúp mã ngắn gọn, dễ quản lý hơn:

```csharp
public class AnimalFactory {
    public static IAnimal CreateAnimal(AnimalType type) {
        switch (type) {
            case AnimalType.Cat: return new Cat();
            case AnimalType.Dog: return new Dog();
            case AnimalType.Duck: return new Duck();
            default: return null;
        }
    }
}
```

Lợi ích: giảm lặp code, dễ thay đổi logic khởi tạo, tập trung logic vào một nơi.

> **Lưu ý quan trọng:** đây **chưa phải** là Factory Method Pattern theo đúng định nghĩa GoF, dù trông có vẻ tương tự. `AnimalFactory.CreateAnimal` vẫn dùng `switch/case` để quyết định loại đối tượng — nghĩa là logic rẽ nhánh vẫn nằm nguyên trong factory, chỉ là được gom vào một chỗ thay vì rải rác. Factory Method Pattern (ở mục tiếp theo) đi xa hơn: nó **loại bỏ hoàn toàn `switch/case`**, thay bằng tính đa hình — mỗi lớp con tự quyết định loại đối tượng nó tạo ra.

---
## 4. Factory Method Pattern là gì?

![Factory Method overview](https://images.viblo.asia/fa6170a7-ffd8-44a4-940f-2ccc47849fcd.png)

**Factory Method** là một mẫu thiết kế thuộc nhóm **Creational Pattern**. Vì dựa trên tính đa hình (polymorphism), nó còn được gọi là **Virtual Constructor** (hàm khởi tạo ảo). Factory Method giải quyết bài toán khởi tạo đối tượng mà **không chỉ định trước chính xác lớp cụ thể nào sẽ được tạo** — quyết định đó được **ủy quyền cho lớp con**.

Cụ thể: Factory Method định nghĩa một phương thức (gọi là _factory method_) trong một lớp cha để tạo đối tượng, nhưng để cho **lớp con ghi đè (override)** phương thức đó nhằm chỉ rõ đối tượng cụ thể nào được tạo ra.

**Mục đích:**

- Cung cấp một cách khởi tạo object thông qua một interface/phương thức chung.
- Che giấu chi tiết xử lý logic của việc khởi tạo khỏi phía client.
- Giảm sự phụ thuộc giữa client và các lớp cụ thể, giúp dễ mở rộng.
- Tránh phải sửa đổi hàng loạt câu lệnh điều kiện mỗi khi thêm loại sản phẩm mới.

![Factory Method purpose](https://images.viblo.asia/6ed7d8a5-7e91-4666-8156-1a0676b2c912.png)

Quay lại ví dụ logistics: cả `Truck` và `Ship` đều hiện thực interface `Transport`, với một phương thức chung là `Deliver` (giao hàng) — nhưng mỗi lớp cài đặt phương thức này theo cách riêng: xe tải giao hàng đường bộ, tàu giao hàng đường biển. Lớp `RoadLogistics` có factory method trả về đối tượng `Truck`, trong khi lớp `SeaLogistics` có factory method trả về đối tượng `Ship`.

![Factory Method solution](https://refactoring.guru/images/patterns/diagrams/factory-method/solution3-en-1.5x.png)

Điểm mấu chốt: đoạn mã sử dụng factory method (gọi là **mã khách - client code**) **không cần biết** sự khác biệt giữa các sản phẩm cụ thể do các lớp con trả về. Client chỉ làm việc với kiểu trừu tượng `Transport`, biết rằng mọi đối tượng vận chuyển đều có phương thức `Deliver`, nhưng không quan tâm cơ chế bên trong của từng loại.

---

## 5. Kiến trúc

![Factory Method structure](https://images.viblo.asia/87b847da-a31e-47ba-83c5-3b4090d80893.png)

Các thành phần trong mô hình:

1. **Product (Sản phẩm):** định nghĩa giao diện chung cho tất cả các đối tượng có thể được tạo ra bởi Creator (người tạo) và các lớp con của nó.
2. **ConcreteProduct (Sản phẩm cụ thể):** các cách hiện thực hóa khác nhau của giao diện Product.
3. **Creator (Người tạo):** khai báo **factory method**, trả về một đối tượng thuộc kiểu Product. Kiểu trả về của phương thức này phải khớp với giao diện Product. Creator có thể tự định nghĩa một cài đặt mặc định cho factory method, trả về một ConcreteProduct mặc định nào đó. Creator cũng thường chứa sẵn logic nghiệp vụ cốt lõi, và gọi đến factory method bên trong logic đó để lấy đối tượng Product cần dùng.
4. **ConcreteCreator (Người tạo cụ thể):** ghi đè factory method để trả về một instance của một ConcreteProduct cụ thể.

Một điểm cần lưu ý: mặc dù tên gọi là "Creator" (người tạo), việc **tạo ra sản phẩm không phải là trách nhiệm chính** của lớp này. Thông thường, Creator đã chứa sẵn logic nghiệp vụ cốt lõi liên quan đến sản phẩm; factory method chỉ giúp tách phần "tạo đối tượng nào" ra khỏi phần "làm gì với đối tượng đó". Ví dụ dễ hình dung: một công ty phần mềm lớn có thể có bộ phận đào tạo lập trình viên mới, nhưng chức năng chính của công ty vẫn là viết mã nguồn, chứ không phải "sản xuất" lập trình viên.

---

## 6. Khi nào nên sử dụng

Nên cân nhắc Factory Method khi:

- Bạn có một lớp cha (hoặc interface) với nhiều lớp con, và cần **quyết định lớp con nào sẽ được tạo dựa trên dữ liệu đầu vào hoặc ngữ cảnh**, mà không muốn client tự làm việc này. Việc chuyển trách nhiệm khởi tạo từ client sang Factory giúp client không phụ thuộc trực tiếp vào các lớp cụ thể.
- Bạn **chưa biết trước** sau này sẽ cần thêm những loại sản phẩm (lớp con) nào. Khi cần mở rộng, chỉ cần tạo thêm một ConcreteProduct và một ConcreteCreator tương ứng — implement thêm factory method cho loại mới, mà không cần sửa code đã có (tuân theo nguyên lý Open/Closed — mở để mở rộng, đóng để sửa đổi).

---
## 7. Ưu & nhược điểm

**Ưu điểm:**
- Che giấu chi tiết xử lý logic của việc khởi tạo khỏi client.
- Hạn chế sự phụ thuộc trực tiếp giữa Creator và các ConcreteProduct cụ thể.
- Dễ dàng mở rộng: thêm loại sản phẩm mới mà không cần sửa đổi các đoạn code đã có (tuân theo nguyên lý Open/Closed).
- Gom các đoạn code tạo ra sản phẩm vào một chỗ, giúp dễ theo dõi và bảo trì.

Nhờ những ưu điểm trên, Factory Method thường được dùng trong các thư viện/framework, nơi người dùng thư viện chỉ cần lấy được đối tượng mong muốn, không cần quan tâm nó được tạo ra như thế nào bên trong.

**Nhược điểm:**
- Mã nguồn có thể trở nên phức tạp hơn vì cần thêm nhiều lớp/interface mới (Product, ConcreteProduct, Creator, ConcreteCreator) để triển khai đầy đủ mẫu này — với hệ thống đơn giản, việc này có thể là dư thừa.
- Refactor một lớp có sẵn (chưa dùng Factory Method) sang dùng Factory Method có thể ảnh hưởng đến client hiện tại, do phải thay đổi cách client lấy đối tượng.

> Factory Method Pattern không yêu cầu constructor phải là private, và bản thân mẫu này dựa trên kế thừa/interface (tức là _khuyến khích_ mở rộng thông qua các ConcreteCreator, không hề ngăn cản việc đó). 

---

## 8. Ví dụ minh họa bằng Go

> Các ví dụ gốc dùng C#. Go không có khái niệm `class`/kế thừa, nên các ví dụ dưới đây dùng `struct` + `interface` — vốn là cách tiếp cận hướng đối tượng thông dụng của Go — để minh họa rõ sự khác biệt giữa Simple Factory và Factory Method thật sự.

### Ví dụ 1: Simple Factory (tương ứng ví dụ `AnimalFactory` ở Mục 3)

```go
package main

import "fmt"

type Animal interface {
	Sound() string
}

type Dog struct{}

func (Dog) Sound() string { return "Gau gau" }

type Cat struct{}

func (Cat) Sound() string { return "Meo meo" }

type Duck struct{}

func (Duck) Sound() string { return "Cap cap" }

// NewAnimal là một Simple Factory: logic rẽ nhánh vẫn nằm trong switch,
// chỉ được gom vào một hàm duy nhất thay vì rải rác nhiều nơi.
func NewAnimal(kind string) (Animal, error) {
	switch kind {
	case "dog":
		return Dog{}, nil
	case "cat":
		return Cat{}, nil
	case "duck":
		return Duck{}, nil
	default:
		return nil, fmt.Errorf("unknown animal type: %s", kind)
	}
}

func main() {
	a, err := NewAnimal("cat")
	if err != nil {
		panic(err)
	}
	fmt.Println(a.Sound())
}
```

### Ví dụ 2: Factory Method thật sự (tương ứng ví dụ Truck/Ship ở Mục 4 và 5)

Ví dụ này thể hiện đúng bản chất Factory Method: **không có `switch/case`** nào quyết định loại `Transport` — quyết định đó nằm ở việc bạn chọn dùng `RoadLogistics{}` hay `SeaLogistics{}`, mỗi lớp tự "ghi đè" factory method của riêng mình.

```go
package main

import "fmt"

// ---- Product ----
type Transport interface {
	Deliver()
}

type Truck struct{}

func (Truck) Deliver() { fmt.Println("Giao hang bang xe tai, tren duong bo") }

type Ship struct{}

func (Ship) Deliver() { fmt.Println("Giao hang bang tau, tren duong bien") }

// ---- Creator ----
// Logistics khai báo factory method CreateTransport. Logic nghiệp vụ dùng
// chung (PlanDelivery) chỉ thao tác qua interface Transport, không biết
// và không cần biết loại Transport cụ thể nào được tạo ra.
type Logistics interface {
	CreateTransport() Transport
}

func PlanDelivery(l Logistics) {
	transport := l.CreateTransport() // factory method quyết định loại transport
	fmt.Println("Len ke hoach giao hang...")
	transport.Deliver()
}

// ---- ConcreteCreator ----
type RoadLogistics struct{}

func (RoadLogistics) CreateTransport() Transport { return Truck{} }

type SeaLogistics struct{}

func (SeaLogistics) CreateTransport() Transport { return Ship{} }

func main() {
	PlanDelivery(RoadLogistics{})
	PlanDelivery(SeaLogistics{})
}
```

**Kết quả:**

```text
Len ke hoach giao hang...
Giao hang bang xe tai, tren duong bo
Len ke hoach giao hang...
Giao hang bang tau, tren duong bien
```

So sánh với Ví dụ 1: ở đây, việc thêm một loại vận chuyển mới (ví dụ `AirLogistics` giao hàng bằng máy bay) chỉ cần thêm một struct `Plane` và một `AirLogistics` mới — hoàn toàn không đụng đến `PlanDelivery` hay các ConcreteCreator đã có.

### Ví dụ 3: Factory kết hợp với Dependency Injection (tương ứng ví dụ Notification ở tài liệu gốc)

Bản C# gốc dùng `IServiceProvider` của ASP.NET Core (một DI container) để tra cứu instance theo tên chuỗi bên trong factory. Go không có DI container tiêu chuẩn đi kèm ngôn ngữ, nên cách idiomatic hơn là dùng một **registry các hàm khởi tạo (map[string]func() ...)**:

```go
package main

import (
	"context"
	"fmt"
)

type NotificationService interface {
	Send(ctx context.Context, message, recipient string) error
}

type EmailService struct{}

func (EmailService) Send(ctx context.Context, message, recipient string) error {
	fmt.Printf("Email sent to %s: %s\n", recipient, message)
	return nil
}

type SmsService struct{}

func (SmsService) Send(ctx context.Context, message, recipient string) error {
	fmt.Printf("SMS sent to %s: %s\n", recipient, message)
	return nil
}

// NotificationFactory ánh xạ tên loại thông báo tới hàm khởi tạo tương ứng.
type NotificationFactory struct {
	creators map[string]func() NotificationService
}

func NewNotificationFactory() *NotificationFactory {
	return &NotificationFactory{
		creators: map[string]func() NotificationService{
			"email": func() NotificationService { return EmailService{} },
			"sms":   func() NotificationService { return SmsService{} },
		},
	}
}

func (f *NotificationFactory) Create(kind string) (NotificationService, error) {
	ctor, ok := f.creators[kind]
	if !ok {
		return nil, fmt.Errorf("unknown notification type: %s", kind)
	}
	return ctor(), nil
}

func main() {
	factory := NewNotificationFactory()
	ctx := context.Background()

	if svc, err := factory.Create("email"); err == nil {
		svc.Send(ctx, "Hello", "user123@gmail.com")
	}
	if svc, err := factory.Create("sms"); err == nil {
		svc.Send(ctx, "Hello", "0123456987")
	}
}
```

> **Ghi chú:** trong bản C# gốc, factory gọi `_serviceProvider.GetRequiredService<T>()` để lấy service theo tên — đây là biểu hiện của **Service Locator**, một cách dùng DI container thường bị xem là anti-pattern khi kết hợp với factory, vì nó "giấu" các phụ thuộc thực sự của class bên trong một lệnh gọi runtime thay vì khai báo tường minh qua constructor. Cách dùng `map[string]func() ...` ở trên tránh được vấn đề này vì các hàm khởi tạo được khai báo tường minh ngay trong factory, không cần tra cứu qua container.

---

### Đề xuất mở rộng

Sau khi nắm vững Factory Method, có thể tìm hiểu thêm:

- **Simple Factory vs Factory Method vs Abstract Factory:** so sánh kỹ hơn 3 cách tiếp cận đã nhắc ở Mục 1, đặc biệt là khi nào cần "nâng cấp" từ Simple Factory lên Factory Method hoặc Abstract Factory.
- **Abstract Factory Pattern:** mẫu nâng cao hơn, dùng khi cần tạo cả một họ đối tượng liên quan với nhau, không chỉ một đối tượng đơn lẻ.
- **Builder Pattern:** một Creational Pattern khác, phù hợp khi đối tượng cần khởi tạo có nhiều tham số tùy chọn phức tạp, thay vì chỉ chọn giữa vài loại cố định như Factory Method.
- **Nguyên lý Open/Closed (trong SOLID):** hiểu rõ nguyên lý này giúp thấy rõ lý do Factory Method giúp mở rộng hệ thống mà không cần sửa code cũ.
- **Dependency Injection và Service Locator anti-pattern:** tìm hiểu sâu hơn sự khác biệt giữa "constructor injection" (khai báo tường minh) và "service locator" (tra cứu ẩn qua container) đã nhắc ở Ví dụ 3.
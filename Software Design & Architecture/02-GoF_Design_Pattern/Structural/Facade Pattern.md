# Facade Pattern

> Tài liệu này giải thích mẫu thiết kế **Facade** — một mẫu thiết kế thuộc nhóm **Structural Pattern** (nhóm mẫu cấu trúc) — dành cho lập trình viên mới học, có liên hệ thực tế với **Golang/Backend**.

## Mục lục

1. [Giới thiệu](https://claude.ai/chat/de82c3b1-20af-49cd-a508-785244f774d4#1-gi%E1%BB%9Bi-thi%E1%BB%87u)
2. [Vấn đề & Giải pháp](https://claude.ai/chat/de82c3b1-20af-49cd-a508-785244f774d4#2-v%E1%BA%A5n-%C4%91%E1%BB%81--gi%E1%BA%A3i-ph%C3%A1p)
3. [Kiến trúc](https://claude.ai/chat/de82c3b1-20af-49cd-a508-785244f774d4#3-ki%E1%BA%BFn-tr%C3%BAc)
4. [Khi nào nên sử dụng Facade](https://claude.ai/chat/de82c3b1-20af-49cd-a508-785244f774d4#4-khi-n%C3%A0o-n%C3%AAn-s%E1%BB%AD-d%E1%BB%A5ng-facade)
5. [Pseudocode minh họa](https://claude.ai/chat/de82c3b1-20af-49cd-a508-785244f774d4#5-pseudocode-minh-h%E1%BB%8Da)
6. [Cách triển khai](https://claude.ai/chat/de82c3b1-20af-49cd-a508-785244f774d4#6-c%C3%A1ch-tri%E1%BB%83n-khai)
7. [Ví dụ minh họa bằng Go](https://claude.ai/chat/de82c3b1-20af-49cd-a508-785244f774d4#7-v%C3%AD-d%E1%BB%A5-minh-h%E1%BB%8Da-b%E1%BA%B1ng-go)
8. [Ưu & nhược điểm](https://claude.ai/chat/de82c3b1-20af-49cd-a508-785244f774d4#8-%C6%B0u--nh%C6%B0%E1%BB%A3c-%C4%91i%E1%BB%83m)
9. [Tài liệu tham khảo](https://claude.ai/chat/de82c3b1-20af-49cd-a508-785244f774d4#9-t%C3%A0i-li%E1%BB%87u-tham-kh%E1%BA%A3o)

---

## 1. Giới thiệu

![Facade overview](https://refactoring.guru/images/patterns/content/facade/facade.png?id=1f4be17305b6316fbd548edf1937ac3b)

- **Facade** là một mẫu thiết kế (design pattern) thuộc nhóm **Structural Pattern** (nhóm mẫu cấu trúc — các mẫu tập trung vào cách tổ chức và kết hợp các lớp/đối tượng với nhau).
- **Facade** cung cấp một **giao diện đơn giản (simplified interface)** thay cho một nhóm các giao diện phức tạp bên trong một **hệ thống con (subsystem)**. Nói cách khác, Facade định nghĩa một giao diện ở **cấp độ cao hơn (higher-level interface)**, giúp người dùng dễ dàng thao tác với hệ thống con mà không cần hiểu chi tiết bên trong.
- Facade cho phép client (mã nguồn gọi đến) giao tiếp với hệ thống con **thông qua một giao diện chung duy nhất**, thay vì phải tự gọi trực tiếp đến từng thành phần bên trong hệ thống con đó. Mục tiêu là **che giấu các hoạt động phức tạp bên trong**, làm cho hệ thống con dễ sử dụng hơn.
- **Tần suất sử dụng:** khá cao — đặc biệt phổ biến khi tích hợp với thư viện/framework bên thứ ba, hoặc khi cần gom nhiều service nhỏ lại thành một điểm gọi duy nhất (tương tự vai trò của một **service layer** trong kiến trúc backend).

---

## 2. Vấn đề & Giải pháp

### Vấn đề gặp phải (Problem)

Trong một ứng dụng phức tạp, có nhiều hành động cần được thực hiện theo một thứ tự nhất định, và các hành động này thường được lặp lại ở nhiều nơi khác nhau trong hệ thống. Ví dụ:

- Khi làm việc với một thư viện bên thứ ba phức tạp, ta cần khởi tạo nhiều đối tượng, theo dõi trạng thái của chúng, và đảm bảo gọi đúng thứ tự logic.
- Khi thao tác với một tập hợp lớn các đối tượng, việc lặp lại quy trình khởi tạo và quản lý trạng thái trở nên rườm rà.

Điều này dẫn đến:

- Phải viết lại cùng một đoạn mã nhiều lần ở các vị trí khác nhau trong ứng dụng.
- Nếu quy trình thay đổi, phải sửa tất cả các đoạn mã liên quan — mất thời gian và dễ gây lỗi.

Hệ quả:

- Mã nguồn khó bảo trì.
- Khả năng xuất hiện lỗi cao mỗi khi có thay đổi.
- Logic nghiệp vụ bị gắn chặt (tightly coupled) với các thành phần bên thứ ba, làm giảm tính linh hoạt của hệ thống.

![Facade problem](https://images.viblo.asia/ac5b8b8c-0aad-46c5-94ae-5400c1990d9e.png)

### Giải pháp (Solution)

**Facade Pattern** giải quyết vấn đề này bằng cách đóng vai trò như một lớp trung gian, cung cấp các phương thức đơn giản hóa thao tác với hệ thống phức tạp bên trong.

Cách hoạt động:

- Thay vì gọi trực tiếp từng thành phần bên trong, client chỉ cần gọi một phương thức duy nhất trong Facade.
- Facade xử lý toàn bộ các bước cần thiết (gọi đúng thứ tự, đúng thành phần) và trả về kết quả.
- Khi quy trình cần thay đổi, ta chỉ cần sửa Facade, không ảnh hưởng đến phần còn lại của ứng dụng.

Lợi ích:

- Giảm sự phụ thuộc của các thành phần trong hệ thống vào chi tiết cài đặt bên trong.
- Dễ bảo trì và mở rộng hơn.
- Mã nguồn gọn gàng, dễ đọc hơn.

![Facade solution](https://images.viblo.asia/51de245b-2a3a-43d8-b88f-f68e847e032c.png)

---

## 3. Kiến trúc

![Facade structure](https://images.viblo.asia/e0ce5777-c04e-4ac4-8303-c44d0901168d.png)

Các thành phần trong mô hình:

- **Facade:** nắm rõ hệ thống con nào đảm nhiệm việc đáp ứng yêu cầu của client, và chuyển yêu cầu đó đến đúng đối tượng trong hệ thống con tương ứng.
- **Additional Facade:** có thể tạo thêm để tránh việc một Facade duy nhất trở nên quá phức tạp. Có thể được dùng bởi cả client lẫn Facade chính.
- **Complex Subsystem:** bao gồm nhiều đối tượng khác nhau, cài đặt các chức năng thực tế của hệ thống con, xử lý công việc được Facade gọi đến. Các lớp này **không cần biết** đến sự tồn tại của Facade và không tham chiếu ngược lại nó.
- **Client:** đối tượng sử dụng Facade để tương tác với hệ thống con, thay vì gọi trực tiếp vào từng thành phần bên trong.

> **Lưu ý:** các đối tượng Facade thường được triển khai theo mẫu **Singleton** (vì thông thường chỉ cần một Facade duy nhất cho một hệ thống con), nhưng đây **không phải là yêu cầu bắt buộc** của Facade Pattern — chỉ là cách triển khai phổ biến trong thực tế.

---

## 4. Khi nào nên sử dụng Facade

**Khi bạn cần một giao diện đơn giản, tinh gọn để tương tác với một hệ thống con phức tạp.**

Theo thời gian, các hệ thống con thường trở nên phức tạp hơn — kể cả khi áp dụng thêm các mẫu thiết kế khác cũng thường sinh ra nhiều lớp mới. Hệ thống con có thể trở nên linh hoạt và tái sử dụng tốt hơn, nhưng đổi lại là khối lượng cấu hình và **mã nguồn lặp lại (boilerplate code)** — tức là những đoạn mã khởi tạo/cấu hình lặp đi lặp lại, không mang giá trị logic nghiệp vụ — mà client phải viết ngày càng tăng. Facade giải quyết vấn đề này bằng cách cung cấp một lối tắt đến các tính năng được dùng nhiều nhất, đáp ứng phần lớn nhu cầu của client mà không cần chạm vào chi tiết bên trong.

Cụ thể, nên cân nhắc dùng Facade khi:

- **Muốn gom nhóm chức năng để client dễ sử dụng.** Khi hệ thống có quá nhiều lớp, người dùng khó nắm được luồng xử lý; khi có nhiều hệ thống con với giao diện riêng lẻ, việc phối hợp sử dụng chúng trở nên khó khăn. Facade tạo ra một giao diện đơn giản cho toàn bộ hệ thống phức tạp đó.
- **Muốn giảm sự phụ thuộc (coupling)** — tức mức độ các thành phần bị "dính chặt" vào nhau — giữa các hệ thống con. Dùng Facade để định nghĩa cổng giao tiếp chung cho mỗi hệ thống con, khiến các hệ thống con chỉ giao tiếp với nhau qua cổng giao diện chung đó thay vì phụ thuộc trực tiếp lẫn nhau.
- **Muốn tổ chức hệ thống con thành các lớp (layer).** Có thể tạo một Facade cho từng cấp độ/lớp của hệ thống con, và để các lớp bên trong chỉ giao tiếp với nhau qua các Facade đó. Cách tiếp cận này khá giống với mẫu thiết kế [Mediator](https://refactoring.guru/design-patterns/mediator) — bạn có thể tìm hiểu thêm để so sánh sự khác biệt.
- **Khi client phụ thuộc quá nhiều vào chi tiết cài đặt.** Facade tách biệt phần sử dụng (client) khỏi phần cài đặt (hệ thống con), giúp hệ thống con dễ thay thế/nâng cấp trong tương lai mà không ảnh hưởng đến client.
- **Cần đóng gói, che giấu một thuật toán/quy trình phức tạp** đằng sau một interface dễ dùng.

**Ví dụ đời thực:** khi bạn gọi điện đến tổng đài của một shop để đặt hàng, tổng đài chính là Facade của toàn bộ các dịch vụ và phòng ban trong cửa hàng (kho hàng, thanh toán, vận chuyển...). Bạn chỉ cần một giao diện đơn giản qua điện thoại để đặt hàng, dù phía sau tổng đài có thể đang phối hợp với nhiều phòng ban khác nhau.

![Facade real-world example](https://images.viblo.asia/25947d05-05cc-429b-9547-815ba7edeee3.png)

---

## 5. Pseudocode minh họa

Ví dụ dưới đây minh họa mẫu Facade giúp đơn giản hóa việc tương tác với một framework chuyển đổi video phức tạp.

![Facade pseudocode diagram](https://refactoring.guru/images/patterns/diagrams/facade/example-1.5x.png)

Thay vì để mã nguồn tương tác trực tiếp với hàng loạt lớp của framework, ta tạo một lớp Facade để đóng gói và ẩn các chức năng đó khỏi phần còn lại của mã nguồn. Nhờ vậy, khi cần nâng cấp lên phiên bản framework mới hoặc thay thế bằng framework khác, thay đổi duy nhất cần thực hiện chỉ nằm trong phần triển khai của lớp Facade.

```pseudocode
// Đây là một số lớp thuộc framework chuyển đổi video phức tạp
// của bên thứ ba. Chúng ta không kiểm soát được mã nguồn này,
// nên không thể đơn giản hóa nó trực tiếp.

class VideoFile
// ...

class OggCompressionCodec
// ...

class MPEG4CompressionCodec
// ...

class CodecFactory
// ...

class BitrateReader
// ...

class AudioMixer
// ...

// Tạo một lớp facade để ẩn đi sự phức tạp của framework
// đằng sau một giao diện đơn giản. Đây là sự đánh đổi
// giữa tính năng đầy đủ và sự đơn giản khi sử dụng.
class VideoConverter is
    method convert(filename, format):File is
        file = new VideoFile(filename)
        sourceCodec = (new CodecFactory).extract(file)
        if (format == "mp4")
            destinationCodec = new MPEG4CompressionCodec()
        else
            destinationCodec = new OggCompressionCodec()
        buffer = BitrateReader.read(filename, sourceCodec)
        result = BitrateReader.convert(buffer, destinationCodec)
        result = (new AudioMixer()).fix(result)
        return new File(result)

// Mã nguồn ứng dụng không cần phụ thuộc vào hàng chục lớp
// của framework phức tạp. Nếu sau này đổi framework, ta chỉ
// cần viết lại lớp facade.
class Application is
    method main() is
        convertor = new VideoConverter()
        mp4 = convertor.convert("funny-cats-video.ogg", "mp4")
        mp4.save()
```

---

## 6. Cách triển khai

1. Xem xét liệu có thể cung cấp một giao diện **đơn giản hơn** so với giao diện hiện có của hệ thống con hay không. Bạn đang đi đúng hướng nếu giao diện này giúp client không còn phải phụ thuộc trực tiếp vào nhiều lớp bên trong hệ thống con.
2. Khai báo và triển khai giao diện đó trong một lớp Facade mới. Facade chuyển hướng các lời gọi từ client đến đúng đối tượng trong hệ thống con. Facade cũng nên chịu trách nhiệm khởi tạo hệ thống con và quản lý vòng đời của nó, trừ khi client đã tự làm việc này.
3. Để tận dụng tối đa lợi ích của mẫu này, hãy đảm bảo **mọi** tương tác giữa client và hệ thống con đều đi qua Facade. Nhờ vậy, client sẽ được bảo vệ trước các thay đổi bên trong hệ thống con — khi hệ thống con nâng cấp, ta chỉ cần sửa bên trong Facade.
4. Nếu Facade trở nên quá cồng kềnh (đảm nhiệm quá nhiều việc), hãy cân nhắc tách bớt chức năng sang một Facade mới, tinh gọn hơn (xem thêm khái niệm **Additional Facade** ở [Mục 3](https://claude.ai/chat/de82c3b1-20af-49cd-a508-785244f774d4#3-ki%E1%BA%BFn-tr%C3%BAc)).

---

## 7. Ví dụ minh họa bằng Go

> Ví dụ gốc tham khảo dùng C#. Dưới đây là bản chuyển sang Go theo phong cách idiomatic: dùng `struct` + method thay cho `class`, và dùng `sync.Once` để triển khai Facade dạng Singleton (thay cho kiểu kiểm tra `if instance == nil` thủ công — cách này **không an toàn** khi có nhiều goroutine gọi đồng thời).

### Ví dụ 1: Facade cho hệ thống bán hàng với nhiều dịch vụ

```go
package main

import (
	"fmt"
	"sync"
)

// ---- Các subsystem: mỗi service không biết gì về Facade ----

type AccountService struct{}

func (AccountService) GetAccount(email string) {
	fmt.Println("Getting the account of", email)
}

type EmailService struct{}

func (EmailService) SendMail(mailTo string) {
	fmt.Println("Sending an email to", mailTo)
}

type PaymentService struct{}

func (PaymentService) PaymentByCash() {
	fmt.Println("Payment by cash")
}

func (PaymentService) PaymentByPaypal() {
	fmt.Println("Payment by Paypal")
}

type ShippingService struct{}

func (ShippingService) FreeShipping() {
	fmt.Println("Free Shipping")
}

func (ShippingService) StandardShipping() {
	fmt.Println("Standard Shipping")
}

type SmsService struct{}

func (SmsService) SendSMS(mobilePhone string) {
	fmt.Println("Sending a message to", mobilePhone)
}

// ---- Facade ----

type ShopFacade struct {
	account  AccountService
	email    EmailService
	payment  PaymentService
	shipping ShippingService
	sms      SmsService
}

var (
	shopInstance *ShopFacade
	shopOnce     sync.Once
)

// GetShopFacade trả về thể hiện duy nhất của ShopFacade (Singleton).
func GetShopFacade() *ShopFacade {
	shopOnce.Do(func() {
		shopInstance = &ShopFacade{}
	})
	return shopInstance
}

func (f *ShopFacade) BuyProductByCashWithFreeShipping(email string) {
	f.account.GetAccount(email)
	f.payment.PaymentByCash()
	f.shipping.FreeShipping()
	f.email.SendMail(email)
	fmt.Println("Done")
}

func (f *ShopFacade) BuyProductByPaypalWithStandardShipping(email, mobilePhone string) {
	f.account.GetAccount(email)
	f.payment.PaymentByPaypal()
	f.shipping.StandardShipping()
	f.email.SendMail(email)
	f.sms.SendSMS(mobilePhone)
	fmt.Println("Done")
}

func main() {
	facade := GetShopFacade()
	facade.BuyProductByCashWithFreeShipping("baob2016947@student.ctu.edu.vn")
	facade.BuyProductByPaypalWithStandardShipping(
		"baob2016947@student.ctu.edu.vn",
		"0123456789",
	)
}
```

**Kết quả sau khi chạy:**

```text
Getting the account of baob2016947@student.ctu.edu.vn
Payment by cash
Free Shipping
Sending an email to baob2016947@student.ctu.edu.vn
Done
Getting the account of baob2016947@student.ctu.edu.vn
Payment by Paypal
Standard Shipping
Sending an email to baob2016947@student.ctu.edu.vn
Sending a message to 0123456789
Done
```

### Ví dụ 2: Facade cho hệ thống đặt hàng

```go
package main

import (
	"fmt"
	"sync"
)

type Product struct{}

func (Product) GetProductDetails() {
	fmt.Println("Chi tiet san pham")
}

type Payment struct{}

func (Payment) MakePayment() {
	fmt.Println("Thuc hien thanh toan thanh cong")
}

type Invoice struct{}

func (Invoice) SendInvoice() {
	fmt.Println("Gui hoa don thanh toan thanh cong")
}

// OrderFacade gom Product, Payment, Invoice thành một luồng đặt hàng duy nhất.
type OrderFacade struct {
	product Product
	payment Payment
	invoice Invoice
}

var (
	orderInstance *OrderFacade
	orderOnce     sync.Once
)

func GetOrderFacade() *OrderFacade {
	orderOnce.Do(func() {
		orderInstance = &OrderFacade{}
	})
	return orderInstance
}

func (f *OrderFacade) PlaceOrder() {
	fmt.Println("Bat dau dat hang")
	f.product.GetProductDetails()
	f.payment.MakePayment()
	f.invoice.SendInvoice()
	fmt.Println("Dat hang thanh cong")
}

func main() {
	GetOrderFacade().PlaceOrder()
	fmt.Println("Ket thuc chuong trinh")
}
```

**Kết quả:**

```text
Bat dau dat hang
Chi tiet san pham
Thuc hien thanh toan thanh cong
Gui hoa don thanh toan thanh cong
Dat hang thanh cong
Ket thuc chuong trinh
```

---

## 8. Ưu & nhược điểm

**Ưu điểm:**

- Tách mã nguồn của bạn ra khỏi sự phức tạp của hệ thống con.
- Hệ thống tích hợp qua Facade đơn giản hơn, vì chỉ cần tương tác với Facade thay vì hàng loạt đối tượng khác.
- Tăng khả năng độc lập và khả năng di chuyển (portability) của hệ thống con, giảm sự phụ thuộc giữa các thành phần.
- Có thể đóng gói nhiều hàm/lớp được thiết kế chưa tốt phía sau một giao diện có thiết kế tốt hơn, giúp client không bị ảnh hưởng bởi những phần code chưa tối ưu đó.

**Nhược điểm:**

- Lớp Facade có thể trở nên quá lớn, ôm quá nhiều nhiệm vụ và hàm chức năng — dễ vi phạm nguyên tắc **SRP (Single Responsibility Principle)**, một trong 5 nguyên tắc thiết kế hướng đối tượng thuộc bộ nguyên tắc **SOLID**.
- Dễ dẫn đến việc phá vỡ các nguyên tắc trong **SOLID** nói chung nếu Facade "ôm đồm" quá nhiều logic nghiệp vụ thay vì chỉ đóng vai trò điều phối.
- Với hệ thống đơn giản, không quá phức tạp, việc thêm một lớp Facade có thể trở nên dư thừa và làm tăng số lớp không cần thiết.

---

## 9. Tài liệu tham khảo

1. [Facade Design Pattern - Trợ thủ đắc lực của Developers](https://viblo.asia/p/facade-design-pattern-tro-thu-dac-luc-cua-developers-924lJBLNlPM)
2. [Facade Pattern – Đơn giản hóa tất cả](https://topdev.vn/blog/facade-pattern-don-gian-hoa-tat-ca/)
3. [Facade Design Pattern trong C# - Cách triển khai và ví dụ](https://freetuts.net/facade-design-pattern-trong-c-sharp-5624.html)
4. [Facade — Refactoring.Guru](https://refactoring.guru/design-patterns/facade)

---

### Đề xuất mở rộng

Sau khi nắm vững Facade, có thể tìm hiểu thêm:

- **Mediator Pattern:** mẫu thiết kế gần giống Facade khi tổ chức hệ thống con thành các lớp giao tiếp qua điểm trung gian — nên so sánh để phân biệt rõ khi nào dùng Facade, khi nào dùng Mediator.
- **Singleton Pattern:** vì Facade thường được triển khai kèm Singleton (`sync.Once` trong Go) — nên hiểu rõ đánh đổi giữa việc dùng Singleton cho Facade và việc dùng dependency injection để dễ test hơn.
- **Adapter Pattern:** một mẫu Structural khác, dễ nhầm với Facade — Adapter tập trung vào việc "chuyển đổi" một interface không tương thích sang interface mong muốn, còn Facade tập trung vào việc "đơn giản hóa" một interface phức tạp có sẵn.
- **Service Layer / Application Service (trong kiến trúc backend):** khái niệm gần gũi với Facade trong thực tế phát triển backend — một layer gom nhiều use-case/service nhỏ lại thành các hàm nghiệp vụ cấp cao hơn cho tầng handler/controller gọi vào.
- **SOLID Principles**, đặc biệt là **SRP (Single Responsibility Principle)**: hiểu rõ nguyên tắc này giúp nhận biết khi nào một Facade đã "phình to" và cần được tách nhỏ.
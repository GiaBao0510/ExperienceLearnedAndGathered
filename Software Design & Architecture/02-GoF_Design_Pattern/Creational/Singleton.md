# Singleton Pattern

> Tài liệu này giải thích mẫu thiết kế **Singleton** — một trong những mẫu thiết kế (design pattern) phổ biến nhất — dành cho lập trình viên mới học, có liên hệ thực tế với **Golang/Backend**.

## Mục lục

1. [Singleton là gì?](https://claude.ai/chat/de82c3b1-20af-49cd-a508-785244f774d4#1-singleton-l%C3%A0-g%C3%AC)
2. [Vấn đề Singleton giải quyết](https://claude.ai/chat/de82c3b1-20af-49cd-a508-785244f774d4#2-v%E1%BA%A5n-%C4%91%E1%BB%81-singleton-gi%E1%BA%A3i-quy%E1%BA%BFt)
3. [Giải pháp](https://claude.ai/chat/de82c3b1-20af-49cd-a508-785244f774d4#3-gi%E1%BA%A3i-ph%C3%A1p)
4. [Ví dụ đời thực](https://claude.ai/chat/de82c3b1-20af-49cd-a508-785244f774d4#4-v%C3%AD-d%E1%BB%A5-%C4%91%E1%BB%9Di-th%E1%BB%B1c)
5. [Kiến trúc](https://claude.ai/chat/de82c3b1-20af-49cd-a508-785244f774d4#5-ki%E1%BA%BFn-tr%C3%BAc)
6. [Pseudocode minh họa](https://claude.ai/chat/de82c3b1-20af-49cd-a508-785244f774d4#6-pseudocode-minh-h%E1%BB%8Da)
7. [Triển khai Singleton trong Go](https://claude.ai/chat/de82c3b1-20af-49cd-a508-785244f774d4#7-tri%E1%BB%83n-khai-singleton-trong-go)
8. [Đặc điểm của Singleton](https://claude.ai/chat/de82c3b1-20af-49cd-a508-785244f774d4#8-%C4%91%E1%BA%B7c-%C4%91i%E1%BB%83m-c%E1%BB%A7a-singleton)
9. [Khi nào nên / không nên dùng Singleton](https://claude.ai/chat/de82c3b1-20af-49cd-a508-785244f774d4#9-khi-n%C3%A0o-n%C3%AAn--kh%C3%B4ng-n%C3%AAn-d%C3%B9ng-singleton)
10. [Các bước triển khai tổng quát](https://claude.ai/chat/de82c3b1-20af-49cd-a508-785244f774d4#10-c%C3%A1c-b%C6%B0%E1%BB%9Bc-tri%E1%BB%83n-khai-t%E1%BB%95ng-qu%C3%A1t)
11. [Ưu & nhược điểm](https://claude.ai/chat/de82c3b1-20af-49cd-a508-785244f774d4#11-%C6%B0u--nh%C6%B0%E1%BB%A3c-%C4%91i%E1%BB%83m)
12. [Tài liệu tham khảo](https://claude.ai/chat/de82c3b1-20af-49cd-a508-785244f774d4#12-t%C3%A0i-li%E1%BB%87u-tham-kh%E1%BA%A3o)

---

## 1. Singleton là gì?

**Singleton** là một mẫu thiết kế (design pattern) thuộc nhóm **Creational Patterns** (nhóm mẫu khởi tạo đối tượng). Nó đảm bảo:

- Một lớp (class) chỉ có **duy nhất một thể hiện (instance)** trong suốt vòng đời của ứng dụng.
- Cung cấp **một điểm truy cập toàn cục (global access point)** đến thể hiện đó, để bất kỳ đâu trong chương trình cũng có thể lấy ra và sử dụng.

![Singleton overview](https://images.viblo.asia/8cc36217-fa29-496b-a2ab-03a5286d8b6b.png)

---

## 2. Vấn đề Singleton giải quyết

Singleton giải quyết đồng thời hai vấn đề. Việc gộp chung hai trách nhiệm này vào một lớp cũng chính là điểm gây tranh cãi của mẫu thiết kế này (sẽ nói rõ hơn ở phần [Nhược điểm](https://claude.ai/chat/de82c3b1-20af-49cd-a508-785244f774d4#11-%C6%B0u--nh%C6%B0%E1%BB%A3c-%C4%91i%E1%BB%83m)):

**Vấn đề 1 — Đảm bảo một lớp chỉ có duy nhất một thể hiện**

Lý do phổ biến nhất để kiểm soát số lượng thể hiện của một lớp là để kiểm soát quyền truy cập vào một **tài nguyên dùng chung** (shared resource) — ví dụ: kết nối cơ sở dữ liệu, tệp cấu hình, hoặc một logger ghi log ra cùng một tệp.

Cơ chế hoạt động: khi bạn "tạo" một đối tượng lần thứ hai, thay vì nhận về một đối tượng hoàn toàn mới, bạn sẽ nhận lại **chính đối tượng đã được tạo ở lần đầu tiên**.

> Lưu ý: hành vi này **không thể** đạt được chỉ bằng hàm khởi tạo (constructor) thông thường, vì theo thiết kế mặc định của hầu hết ngôn ngữ lập trình, mỗi lần gọi constructor luôn trả về một đối tượng mới.

![Singleton comic](https://refactoring.guru/images/patterns/content/singleton/singleton-comic-1-en.png)

**Vấn đề 2 — Cung cấp một điểm truy cập toàn cục**

Đây là ý tưởng gợi nhớ đến **biến toàn cục (global variable)**: tiện lợi vì có thể truy cập từ bất kỳ đâu, nhưng rủi ro vì bất kỳ đoạn mã nào cũng có thể vô tình ghi đè giá trị của biến đó.

Singleton giữ lại sự tiện lợi này (truy cập từ mọi nơi) nhưng **bảo vệ thể hiện khỏi bị ghi đè** bởi mã nguồn khác, vì việc khởi tạo và quản lý instance được gom vào một chỗ duy nhất — thay vì rải rác khắp chương trình.

> Trong thực tế, người ta thường gọi một thành phần là "Singleton" ngay cả khi nó chỉ giải quyết một trong hai vấn đề trên (thường là vấn đề 1).

---

## 3. Giải pháp

Mọi cách triển khai Singleton đều dựa trên hai bước cốt lõi:

1. **Ẩn constructor mặc định** bằng cách đặt nó ở chế độ `private` (riêng tư), để không lớp nào khác có thể dùng từ khóa `new` (hoặc tương đương) để tạo trực tiếp đối tượng.
2. **Tạo một phương thức khởi tạo tĩnh (static creation method)** đóng vai trò thay thế cho constructor. Phương thức này:
    - Gọi constructor `private` để tạo đối tượng **trong lần gọi đầu tiên**, rồi lưu vào một trường tĩnh (static field).
    - Ở tất cả các lần gọi sau, chỉ trả về đối tượng đã lưu trữ đó — không tạo mới.

Vì phần code nào có quyền truy cập lớp Singleton đều chỉ có thể gọi phương thức tĩnh này, nên mọi lần gọi đều nhận về **cùng một đối tượng duy nhất**.

---

## 4. Ví dụ đời thực

Chính phủ là một ví dụ điển hình cho mẫu Singleton: một quốc gia chỉ có duy nhất một chính phủ chính thức tại một thời điểm. Cho dù thành phần nhân sự bên trong có thay đổi, danh xưng "Chính phủ của quốc gia X" vẫn luôn là **một điểm truy cập chung**, đại diện cho nhóm người đang nắm quyền điều hành.

---

## 5. Kiến trúc

Về mặt cấu trúc, một Singleton gồm ba thành phần:

- **Constructor `private`**: ngăn client (mã gọi từ bên ngoài) tự khởi tạo đối tượng của lớp.
- **Biến `static private`** (thường đặt tên là `instance`): lưu trữ thể hiện duy nhất của lớp.
- **Phương thức `public static`** (thường đặt tên là `getInstance()`): là cách duy nhất để các lớp khác truy cập vào instance đó.

![Singleton structure](https://refactoring.guru/images/patterns/diagrams/singleton/structure-en-indexed.png)

---

## 6. Pseudocode minh họa

Ví dụ dưới đây minh họa một lớp kết nối cơ sở dữ liệu (`Database`) được triển khai theo mẫu Singleton. Vì constructor là `private`, cách duy nhất để lấy đối tượng là gọi `getInstance()`. Phương thức này tạo đối tượng ở lần gọi đầu tiên, và trả về chính đối tượng đó ở mọi lần gọi sau.

```pseudocode
class Database is
    // Trường lưu trữ thể hiện singleton — phải là static
    private static field instance: Database

    // Constructor luôn ở chế độ private để ngăn khởi tạo trực tiếp bằng `new`
    private constructor Database() is
        // Một số logic khởi tạo, ví dụ: mở kết nối tới database server
        // ...

    // Phương thức tĩnh kiểm soát việc truy cập tới instance
    public static method getInstance() is
        if (Database.instance == null) then
            acquireThreadLock() and then
                // Kiểm tra lại lần nữa, phòng trường hợp một luồng (thread)
                // khác đã khởi tạo instance trong lúc luồng này chờ lock
                if (Database.instance == null) then
                    Database.instance = new Database()
        return Database.instance

    // Logic nghiệp vụ chạy trên instance duy nhất này
    public method query(sql) is
        // Mọi truy vấn của ứng dụng đều đi qua đây,
        // nên có thể thêm logic throttling hoặc caching tại đây
        // ...

class Application is
    method main() is
        Database foo = Database.getInstance()
        foo.query("SELECT ...")

        Database bar = Database.getInstance()
        bar.query("SELECT ...")
        // foo và bar trỏ tới cùng một object
```

Đoạn `acquireThreadLock()` ở trên gọi là kỹ thuật **double-checked locking**: kiểm tra `instance == null` hai lần — một lần trước khi khóa (để tránh khóa không cần thiết khi instance đã tồn tại), một lần sau khi khóa (để tránh hai luồng cùng tạo instance). Đây là kỹ thuật kinh điển trong các ngôn ngữ như Java/C++, nhưng **không nên sao chép y nguyên sang Go** — xem phần tiếp theo.

---

## 7. Triển khai Singleton trong Go

> **Ghi chú kỹ thuật quan trọng:** mẫu double-checked locking ở pseudocode trên, nếu viết "thô" bằng `if instance == nil` như vậy trong Go mà không dùng cơ chế đồng bộ hóa đúng cách, **không an toàn** trong môi trường đa luồng (goroutine) do mô hình bộ nhớ (memory model) của Go không đảm bảo thứ tự nhìn thấy giá trị giữa các goroutine nếu thiếu `sync.Mutex`/`sync/atomic`. Go cung cấp sẵn công cụ chuẩn cho đúng bài toán này: **`sync.Once`**.

### Cách triển khai idiomatic (khuyến nghị)

```go
package config

import "sync"

type Config struct {
    Env string
}

var (
    instance *Config
    once     sync.Once
)

// GetInstance trả về thể hiện duy nhất của Config.
// sync.Once đảm bảo hàm khởi tạo bên trong chỉ chạy đúng một lần,
// kể cả khi nhiều goroutine cùng gọi GetInstance() đồng thời.
func GetInstance() *Config {
    once.Do(func() {
        instance = &Config{Env: "production"}
    })
    return instance
}
```

`sync.Once.Do()` tự xử lý phần khóa (lock) và kiểm tra an toàn, nên bạn không cần tự viết double-checked locking thủ công. Đây là cách phổ biến nhất để triển khai Singleton trong các dự án Go thực tế — ví dụ: singleton cho logger, connection pool, hoặc client cấu hình dùng chung.

### Ví dụ áp dụng: Logger dùng chung

```go
package logger

import (
    "log"
    "os"
    "sync"
)

var (
    logInstance *log.Logger
    once        sync.Once
)

func GetLogger() *log.Logger {
    once.Do(func() {
        logInstance = log.New(os.Stdout, "APP: ", log.LstdFlags)
    })
    return logInstance
}
```

Ở bất kỳ package nào trong ứng dụng, `logger.GetLogger()` luôn trả về cùng một `*log.Logger`, đảm bảo toàn bộ log được ghi qua cùng một cấu hình.

---

## 8. Đặc điểm của Singleton

- **Duy nhất (unique):** chỉ tồn tại một thể hiện của lớp trong suốt vòng đời ứng dụng. Ví dụ điển hình: quản lý cấu hình, kết nối cơ sở dữ liệu, logging, thread pool / worker pool.
- **Toàn cục (global access):** có thể truy cập thể hiện này từ bất kỳ đâu trong ứng dụng.
- **Kiểm soát khởi tạo:** thể hiện chỉ được tạo khi thực sự cần (**lazy initialization** — khởi tạo trễ, tức là chỉ tạo ở lần gọi đầu tiên, không tạo sẵn ngay khi chương trình chạy).

---

## 9. Khi nào nên / không nên dùng Singleton

### Nên dùng khi:

- Cần **đúng một** đối tượng duy nhất trong toàn bộ ứng dụng, và việc có thêm bản sao thứ hai sẽ gây lỗi hoặc lãng phí tài nguyên.
- **Quản lý cấu hình:** ứng dụng chỉ cần một nguồn cấu hình duy nhất.
- **Quản lý kết nối cơ sở dữ liệu:** tái sử dụng connection/connection pool thay vì mở kết nối mới liên tục.
- **Hệ thống ghi log:** một logger dùng chung, đảm bảo log không bị phân mảnh hoặc ghi chồng chéo.
- Bạn cần kiểm soát chặt chẽ hơn so với việc dùng biến toàn cục thông thường — vì Singleton bảo vệ instance khỏi bị ghi đè tùy tiện, điều mà biến toàn cục không đảm bảo được.

> Về lý thuyết, giới hạn "chỉ một instance" không phải bất biến tuyệt đối — bạn có thể chỉnh sửa `getInstance()` để cho phép tạo tối đa N instance nếu bài toán yêu cầu, mà không cần thay đổi phần còn lại của code.

### Không nên dùng khi:

- Lớp cần hỗ trợ **nhiều instance** trong các ngữ cảnh khác nhau (ví dụ: nhiều kết nối tới nhiều database khác nhau).
- Việc viết **unit test** trở nên khó khăn: vì Singleton giữ trạng thái toàn cục (global state) xuyên suốt, rất khó mock hoặc thay thế instance giữa các test case, dễ khiến các test ảnh hưởng lẫn nhau.
- Có nguy cơ bị lạm dụng như một biến toàn cục "trá hình", làm giảm tính linh hoạt, khó bảo trì, và che giấu các phụ thuộc ẩn (hidden dependencies) giữa các phần của chương trình. Trong nhiều trường hợp, **dependency injection** là giải pháp thay thế nên cân nhắc.

---

## 10. Các bước triển khai tổng quát

1. Thêm một **trường tĩnh, riêng tư** (private static field) vào lớp để lưu trữ thể hiện singleton.
2. Khai báo một **phương thức khởi tạo tĩnh, công khai** (public static method) để lấy thể hiện đó.
3. Triển khai **lazy initialization** bên trong phương thức tĩnh này: tạo đối tượng mới ở lần gọi đầu tiên, lưu vào trường tĩnh; các lần gọi sau chỉ trả về đối tượng đã lưu.
4. Đặt **constructor** của lớp ở chế độ **riêng tư** (private) — chỉ phương thức tĩnh nội bộ mới gọi được, các đối tượng khác thì không.
5. Rà soát mã nguồn phía client, thay mọi lệnh gọi constructor trực tiếp bằng lệnh gọi phương thức khởi tạo tĩnh.

> Trong Go không có khái niệm `private constructor` theo nghĩa OOP truyền thống (Go không có class/constructor); thay vào đó, ta dùng **quy ước không export** (tên biến/hàm viết thường, không viết hoa chữ cái đầu) để giới hạn phạm vi truy cập trong package — như ví dụ ở [Mục 7](https://claude.ai/chat/de82c3b1-20af-49cd-a508-785244f774d4#7-tri%E1%BB%83n-khai-singleton-trong-go).

---

## 11. Ưu & nhược điểm

**Ưu điểm:**

- **Đảm bảo tính duy nhất:** chỉ có một thể hiện trong toàn bộ ứng dụng.
- **Tiết kiệm tài nguyên:** đặc biệt quan trọng khi việc khởi tạo đối tượng tốn kém (ví dụ: mở kết nối database).
- Đối tượng chỉ được khởi tạo đúng một lần, tại thời điểm được yêu cầu lần đầu.
- Kiểm soát chặt chẽ việc truy cập tới instance duy nhất.
- Giảm số lượng biến toàn cục cần quản lý (giảm ô nhiễm namespace).
- **Truy cập toàn cục:** đơn giản hóa việc chia sẻ một đối tượng dùng chung giữa các phần của chương trình.

**Nhược điểm:**

- **Vi phạm nguyên tắc SRP (Single Responsibility Principle):** Singleton vừa lo việc khởi tạo/quản lý vòng đời, vừa cung cấp logic nghiệp vụ — hai trách nhiệm gộp vào một lớp.
- **Khó viết unit test:** vì Singleton mang trạng thái toàn cục, khó cô lập (isolate) hoặc mock instance giữa các test case.
- **Rủi ro trong môi trường đa luồng:** nếu triển khai sai (ví dụ thiếu đồng bộ hóa đúng cách), có thể tạo ra nhiều hơn một instance hoặc gây race condition. Trong Go, dùng `sync.Once` sẽ tránh được rủi ro này (xem [Mục 7](https://claude.ai/chat/de82c3b1-20af-49cd-a508-785244f774d4#7-tri%E1%BB%83n-khai-singleton-trong-go)).

---

## 12. Tài liệu tham khảo

1. https://viblo.asia/p/singleton-design-pattern-tro-thu-dac-luc-cua-developers-Qbq5QBkJKD8
2. https://viblo.asia/p/hoc-singleton-pattern-trong-5-phut-4P856goOKY3
3. [Singleton — Refactoring.Guru](https://refactoring.guru/design-patterns/singleton)

---

### Đề xuất mở rộng

Sau khi nắm vững Singleton, có thể tìm hiểu thêm:

- **Dependency Injection (DI):** giải pháp thay thế Singleton trong nhiều trường hợp, giúp code dễ test hơn vì các phụ thuộc được truyền vào tường minh thay vì ẩn trong global state.
- **`sync.Once` và gói `sync` trong Go:** tìm hiểu sâu hơn các primitive đồng bộ hóa khác như `sync.Mutex`, `sync.RWMutex`, `sync/atomic`.
- **Các Creational Pattern khác:** Factory Method, Abstract Factory, Builder, Prototype — để so sánh khi nào nên dùng Singleton thay vì các mẫu này.
- **Anti-pattern "Service Locator":** một cách dùng sai phổ biến khi kết hợp Singleton với việc "tra cứu" dịch vụ toàn cục, dẫn tới code khó kiểm soát phụ thuộc.
- **Testing với Singleton:** kỹ thuật interface + constructor injection để có thể "swap" implementation thật bằng mock trong unit test, ngay cả khi production code dùng Singleton.
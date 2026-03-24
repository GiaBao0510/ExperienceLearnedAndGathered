# Luồng Thực Thi của Một Chương Trình Go

## 1. Tổng quan

Khi một chương trình Go được khởi động, nó không đơn giản là chạy thẳng vào hàm `main()`. Trước đó, Go cần **chuẩn bị môi trường** bằng cách nạp các package phụ thuộc, khởi tạo biến, hằng và gọi các hàm `init()`. Chỉ sau khi tất cả đã sẵn sàng, chương trình mới thực sự bắt đầu tại `main.main()`.

---

## 2. Điểm bắt đầu: `main.main()`

Mọi chương trình Go đều bắt đầu thực thi từ hàm `main()` trong package `main`. Đây là **điểm vào duy nhất** của chương trình.

```go
package main

import "fmt"

func main() {
    fmt.Println("Chương trình bắt đầu!")
}
```

> **Lưu ý:** Tên package phải là `main` và hàm phải tên là `main()` — không thể đặt tên khác.

---
## 3. Thứ tự khởi tạo khi import package

Khi `package main` import các package khác, Go sẽ khởi tạo chúng theo một thứ tự nhất định trước khi chạy `main()`.

### Thứ tự thực hiện cho mỗi package:

1. **Import các package phụ thuộc** (đệ quy — nếu package đó lại import package khác thì xử lý tiếp theo thứ tự tương tự)
2. **Khởi tạo các hằng số** (`const`)
3. **Khởi tạo các biến** (`var`)
4. **Gọi hàm `init()`** (nếu có)

### Ví dụ minh họa:

```go
// package database
package database

import "fmt"

var Connection = "localhost:5432"

func init() {
    fmt.Println("[database] init() được gọi")
}
```

```go
// package main
package main

import (
    "fmt"
    _ "myapp/database" // import để kích hoạt init()
)

func init() {
    fmt.Println("[main] init() được gọi")
}

func main() {
    fmt.Println("[main] main() bắt đầu")
}
```

**Kết quả in ra:**

```
[database] init() được gọi
[main] init() được gọi
[main] main() bắt đầu
```

---

## 4. Hàm `init()` — Những điều cần biết

Hàm `init()` là một hàm đặc biệt trong Go, dùng để **khởi tạo trạng thái ban đầu** cho một package.

### Đặc điểm quan trọng:

|Đặc điểm|Mô tả|
|---|---|
|Không có tham số|`init()` không nhận tham số và không trả về giá trị|
|Không thể gọi thủ công|Các hàm khác **không thể** gọi `init()` trực tiếp|
|Có thể khai báo nhiều lần|Một file hoặc một package có thể có **nhiều hàm `init()`**|
|Thứ tự gọi|Gọi theo thứ tự xuất hiện trong file (từ trên xuống dưới)|
|Chỉ chạy một lần|Dù package được import nhiều lần, `init()` chỉ chạy **đúng một lần**|

### Ví dụ một package có nhiều `init()`:

```go
package mypackage

import "fmt"

func init() {
    fmt.Println("init() lần 1")
}

func init() {
    fmt.Println("init() lần 2")
}
```

**Kết quả:**

```
init() lần 1
init() lần 2
```

---

## 5. Package được import nhiều lần

Trong một dự án lớn, cùng một package có thể được import từ nhiều nơi khác nhau. Go đảm bảo rằng mỗi package **chỉ được khởi tạo đúng một lần**, dù nó được import bao nhiêu lần đi nữa.

```
main
 ├── import A
 │    └── import C   ← C được import lần 1
 └── import B
      └── import C   ← C được import lần 2 (bỏ qua, không khởi tạo lại)
```

> Go sẽ chỉ chạy `init()` của package `C` **một lần duy nhất**.

---

## 6. Sơ đồ tổng thể luồng thực thi

```
Chương trình Go khởi động
         │
         ▼
  Import các package phụ thuộc (đệ quy từ sâu nhất trước)
         │
         ▼
  Khởi tạo const & var ở package-level
         │
         ▼
  Gọi hàm init() (theo thứ tự xuất hiện)
         │
         ▼
  Thực thi main.main()
         │
         ▼
  Chương trình kết thúc
```

> Hình ảnh minh họa chi tiết hơn: [https://zalopay-oss.github.io/go-advanced/images/ch1-11-init.ditaa.png](https://zalopay-oss.github.io/go-advanced/images/ch1-11-init.ditaa.png)

---

## 7. Goroutine và luồng chính

Trước khi `init()` hoàn tất, **tất cả code đều chạy trên một Goroutine duy nhất** — đó là Goroutine chính (`main goroutine`).

Nếu bạn khởi chạy một Goroutine mới bên trong `init()` hoặc `main()`, Goroutine đó **chỉ được thực thi sau khi `init()` kết thúc** và chương trình đã đi vào `main.main()`.

### Ví dụ:

```go
package main

import "fmt"

func init() {
    go func() {
        fmt.Println("Goroutine trong init()") // Có thể chưa chạy ngay
    }()
    fmt.Println("init() đang chạy trên goroutine chính")
}

func main() {
    fmt.Println("main() bắt đầu")
    // Cần thêm cơ chế đồng bộ (sync.WaitGroup, channel...) 
    // để đảm bảo goroutine con hoàn thành
}
```

> **Lưu ý:** Khi `main()` kết thúc, toàn bộ chương trình sẽ dừng lại — kể cả các Goroutine đang chạy dở. Vì vậy, cần dùng cơ chế đồng bộ nếu bạn muốn chắc chắn các Goroutine hoàn tất.

---
## 8. Tóm tắt

![](https://zalopay-oss.github.io/go-advanced/images/ch1-11-init.ditaa.png)

```
Thứ tự khởi tạo:
  package phụ thuộc sâu nhất  →  const  →  var  →  init()  →  main()

Quy tắc cần nhớ:
  ✔ init() chạy tự động, không thể gọi thủ công
  ✔ Một package có thể có nhiều init()
  ✔ Mỗi package chỉ được khởi tạo một lần dù import nhiều nơi
  ✔ Goroutine chỉ chạy sau khi init() hoàn tất
  ✔ Khi main() kết thúc, toàn bộ chương trình dừng lại
```
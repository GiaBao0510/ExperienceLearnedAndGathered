#### **1. Go là gì?**

![](https://techvccloud.mediacdn.vn/zoom/300_180/280518386289090560/2022/4/28/golang-la-gi-165111502649548114519-13-0-350-600-crop-1651115032844167630610.jpg)

**Go** là một ngôn ngữ tổng quát được thiết với mục đích lập trình hệ thống. Ban đầu nó được phát triển tại Google vào năm 2007 bởi Robert Griesemer, Rob Pike và Ken Thompson.
Ngôn ngữ này được định hình mạnh và tĩnh, cung cấp hỗ trợ sẵn để thu gom dữ liệu ra và hỗ trợ lập trình đồng thời. Các chương trình được xây ==dựng bằng các sử dụng các packages==, để quản lý hiệu quả các **dependencies**. Triển khai lập trình **Go** ==sử dụng biên dịch và liên kết mô hình để tạo thực thi các tệp nhị phân==.

---
### **Tại sao dùng golang?**

Do Golang có hiệu suất cao tương đương với C/C++, cú pháp đơn giản dễ học, và khả năng xử lý lượng đồng thời cao (concurrency) vượt trội với Goroutines. Nó lý tưởng cho backend, microservice và hệ thống đám mây, giúp tối ưu tài nguyên phần cứng, biên dịch nhanh và dễ dàng bảo trì

Dưới đây là các lý do chi tiết tại sao nên dùng Golang:

- **Hiệu suất vượt trội (Performance):** Go là ngôn ngữ biên dịch (compiled) sang mã máy, giúp tốc độ xử lý nhanh chóng, gần tương đương với C/C++.
- **Xử lý đồng thời (Concurrency) dễ dàng:** Goroutines và Channels tích hợp giúp Go quản lý hàng triệu tiến trình nhẹ nhàng, tối ưu cho các hệ thống phân tán và dịch vụ mạng.
- **Quản lý bộ nhớ hiệu quả:** Go tích hợp cơ chế thu gom rác (garbage collection) tự động, giúp quản lý bộ nhớ tốt mà không đòi hỏi nhiều sự can thiệp.
- **Triển khai (Deployment) dễ dàng:** Go biên dịch ra một file nhị phân duy nhất (single binary), không phụ thuộc vào các thư viện bên ngoài, giúp việc triển khai lên server cực kỳ gọn nhẹ.
- **Được hỗ trợ bởi Google:** Đảm bảo độ tin cậy, sự phát triển lâu dài và phù hợp cho các cơ sở hạ tầng quy mô lớn
---
#### **2. Go là một language, framework hay library mới?**

**Go** không phải là một library và không phải là một framework, đó là một ngôn ngữ mới. ==Go hầu hết thuộc họ C (cú pháp cơ bản)==, với đầu vào đáng kể họ Pascal/ Modula / Oberon (khai báo ,gói).

**Go** có một thư viện rộng lớn, được gọi là **runtime**, là một phần của mọi chương trình **Go**. Mặc dù đó là ngôn ngữ trung tâm, **runtime** của Go tương tự như libc, thư viện C. Tuy nhiên, điều quan trọng cần hiểu là **runtime** của Go không bao gồm máy ảo, chẳng hạng như không bao gồm Java runtime.

Các chương trình Go được biên dịch trước sang mã máy gốc.

---
### **3. Ưu/ nhược điểm của Golang là gì?**

Go lang được đánh giá cao nhờ hiệu suất mạnh mẽ, cú pháp đone giản và khả năng xử lý đồng thời vượt trội. Tuy nhiên, ngôn ngữ này vẫn tồn tại một số điểm hạn chế so với các ngôn ngữ lập trình khác

#### **Ưu điểm**

Nhờ hiệu suất xử lý cao, gần tương đương với các ngôn ngữ biên dịch như C/C++. Khả năng hỗ trợ lập trình (Concurrency) với Goroutines giúp Golang xử lý hàng triệu tiến trình một cách nhanh chóng và nhẹ nhàng, rất phù hợp cho hệ thống phân tán microservices và dịch vụ web hiệu suất cao.
Bên cạnh đó thì cú pháp của Golang được thiết kế đơn giản, dễ đọc, dễ học và thân thiện với người mới

#### **Nhược điểm**

Golang không hỗ trợ kế thừa trong lập trình hướng đối tượng, khiến việc thiết kế hệ thống theo OOP có phân hạn chế. Mặc dù Go có hỗ trợ Interface nhưng vẫn thiếu đi những tính năng nâng cao như abstract class hay methob overloading.

Go chưa thật sự mạnh trong các lĩnh vực như trí tuệ nhân tạo (AI), học máy (machine learning), và xử lý dữ liệu lớn (big data), nơi mà Python đang chiếm ưu thế tuyệt đối. Ngoài ra, việc phát triển ứng dụng frontend, game hoặc phần mềm giao diện người dùng bằng Golang là không phù hợp do thiếu thư viện và công cụ hỗ trợ.

Cuối cùng, hệ thống quản lý thư viện bên ngoài (module) của Go vẫn còn đơn giản so với các hệ sinh thái lớn như Node.js (npm) hay Python (pip).


---
### 4. **Làm thế nào phát hiện `Data race` trong Go? và cách giải quyết chúng như thế nào?** [trả lời.](Go/Go_Interview/Data-Race_golang)
---
#### 5. **Concurrency trong golang [Trả lời](33.1_Goroutine_channel.md)**

---
#### 6. **Channel là gì?  Sự khác nhau giữa channel buffer và unbuffer [Trả lời](33.1_Goroutine_channel.md)**

---
### 7.  **Giải thích method & interface**

#### 7.1. Method (Phương thức)

Trong Go, **Method** thực chất là một hàm (function), nhưng có thêm một tham số đặc biệt gọi là **receiver** (bộ nhận). Receiver này nằm giữa từ khóa `func` và tên hàm.

Thay vì gói mọi thứ vào một class, Go cho phép bạn "gắn" hàm vào một kiểu dữ liệu bất kỳ (thường là `struct`).

```go
type Rectangle struct {
    width, height float64
}

// Đây là một Method. (r Rectangle) là receiver.
func (r Rectangle) Area() float64 {
    return r.width * r.height
}
```

### Pointer vs. Value Receiver:

Đây là điểm quan trọng  cần nhớ:
- **Value Receiver (`r Rectangle`):** Go sẽ tạo một bản sao của đối tượng. Mọi thay đổi bên trong method sẽ không ảnh hưởng đến đối tượng gốc.
- **Pointer Receiver (`r *Rectangle`):** Go dùng địa chỉ ô nhớ. Method có thể thay đổi dữ liệu của đối tượng gốc và hiệu năng tốt hơn vì không phải copy dữ liệu.

#### 7.2. Interface (Giao diện)

Nếu Method là **hành động**, thì Interface là **bản hợp đồng**. Nó định nghĩa một tập hợp các chữ ký hàm (method signatures). Một kiểu dữ liệu được coi là "thực thi" (implement) interface đó nếu nó sở hữu tất cả các method mà interface yêu cầu.

**Điểm đặc biệt nhất của Go:** Bạn không cần từ khóa `implements`. Chỉ cần bạn có đủ method, Go sẽ tự hiểu bạn đã "ký hợp đồng". Đây gọi là **Implicit Implementation** (thực thi ngầm định).

```go
// Định nghĩa "hợp đồng" Shape
type Shape interface {
    Area() float64
}

// Đối tượng Square
type Square struct {
    side float64
}

// Square thực thi Area() -> Tự động thỏa mãn interface Shape
func (s Square) Area() float64 {
    return s.side * s.side
}

func PrintArea(sh Shape) {
    fmt.Println("Diện tích là:", sh.Area())
}
```

##### 7.3. Sự khác biệt chính

| **Đặc điểm** | **Method**                                   | **Interface**                                          |
| ------------ | -------------------------------------------- | ------------------------------------------------------ |
| **Bản chất** | Là hàm gắn liền với một kiểu dữ liệu cụ thể. | Là một kiểu dữ liệu trừu tượng định nghĩa các hành vi. |
| **Mục đích** | Thực thi logic xử lý cho đối tượng.          | Tạo ra sự đa hình (polymorphism), giúp code linh hoạt. |
| **Khai báo** | Phải có receiver.                            | Chỉ chứa tên hàm, tham số và kiểu trả về.              |
#### Tại sao Go lại làm như vậy?

Cách thiết kế này giúp Go đạt được sự linh hoạt tuyệt vời:
1. **Tính đóng gói (Encapsulation):** Bạn có thể thêm method cho bất kỳ type nào trong package của mình (không chỉ struct).
2. **Tính đa hình (Polymorphism):** Bạn có thể viết một hàm nhận vào một `Interface`, và hàm đó sẽ hoạt động với bất kỳ object nào thỏa mãn interface đó, giúp code cực kỳ dễ mở rộng.
3. **Hạn chế phụ thuộc:** Bạn có thể định nghĩa interface ở nơi sử dụng, thay vì nơi tạo ra đối tượng.

---
#### **8. Slice là gì?** [Trả lời](28_Slice.md)

---
#### **9.  Pointer trong go [Trả lời](21_Pointer.md)

---

10. 
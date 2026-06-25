# Mô hình MVC

## Mô hình MVC là gì?

Mô hình **Model-View-Controller (MVC)** là một mẫu kiến trúc phần mềm tách ứng dụng thành ba thành phần có trách nhiệm rõ ràng: **Model**, **View** và **Controller**. Mục tiêu cốt lõi của MVC là tách biệt lớp logic nghiệp vụ (business logic) khỏi lớp hiển thị (presentation), giúp mã nguồn dễ bảo trì, kiểm thử và mở rộng hơn.

![](https://images.viblo.asia/ed9ea401-6c09-4ca8-9aa1-8b143956db8f.png)

---

## Kiến trúc MVC

![](https://static.vietnix.vn/wp-content/uploads/2022/03/cac-thanh-phan-cua-mvc.webp)

MVC gồm ba thành phần chính:

- **Model:** Chứa dữ liệu và toàn bộ logic nghiệp vụ của ứng dụng. Model định nghĩa cấu trúc dữ liệu, các quy tắc nghiệp vụ (business rules), và chịu trách nhiệm tương tác với nguồn dữ liệu (database, external API...).
- **View:** Chịu trách nhiệm hiển thị dữ liệu cho người dùng. View nhận dữ liệu từ Model (thông qua Controller) và trình bày dưới dạng giao diện người dùng.
- **Controller:** Đóng vai trò trung gian giữa Model và View. Controller nhận request từ người dùng, điều phối Model để xử lý logic, rồi chọn View phù hợp để hiển thị kết quả.

---

### View

View là thành phần đại diện cho giao diện người dùng (UI). Nhiệm vụ duy nhất của View là trình bày dữ liệu — nó không chứa logic nghiệp vụ.

- Nhận dữ liệu từ Controller và render ra giao diện (HTML, JSON, XML...).
- Có thể bao gồm các thành phần UI như nút bấm, ô nhập liệu, danh sách thả xuống.
- Không trực tiếp giao tiếp với Model.

**Ví dụ:** Trong ứng dụng mua sắm, View hiển thị giỏ hàng bao gồm danh sách sản phẩm, giá, nút "Thêm vào giỏ" và nút "Thanh toán". Toàn bộ dữ liệu trong đó đến từ Controller, không phải View tự truy vấn.

---

### Controller

Controller là thành phần xử lý tương tác của người dùng và điều phối luồng dữ liệu giữa Model và View.

- Nhận và diễn giải request từ người dùng (HTTP request, form submission, button click...).
- Gọi Model để thực hiện nghiệp vụ hoặc truy vấn dữ liệu.
- Chọn và truyền dữ liệu đến View phù hợp để hiển thị kết quả.
- Không chứa logic nghiệp vụ phức tạp — đây là trách nhiệm của Model (hoặc Service layer).

**Ví dụ:** Khi người dùng nhấn "Xóa sản phẩm khỏi giỏ hàng", request được gửi đến Controller. Controller gọi Model để xóa sản phẩm khỏi database, sau đó trả dữ liệu giỏ hàng mới cho View để render lại.

---

### Model

Model là thành phần chứa dữ liệu và toàn bộ logic nghiệp vụ của ứng dụng.

- Định nghĩa cấu trúc dữ liệu và các business rules.
- Thực hiện các thao tác với nguồn dữ liệu: truy vấn, thêm, sửa, xóa (CRUD).
- Không biết gì về cách dữ liệu được hiển thị — Model hoàn toàn độc lập với View.
- Khi trạng thái dữ liệu thay đổi, Model thông báo cho Controller (và trong một số biến thể, trực tiếp cho View qua Observer pattern).

**Ví dụ:** Trong ứng dụng mua sắm, Model `Cart` định nghĩa giỏ hàng bao gồm danh sách mặt hàng, giá từng mặt hàng, tổng giá trị, và các phương thức như `addItem()`, `removeItem()`, `calculateTotal()`.

---

## Luồng hoạt động

![](https://static.vietnix.vn/wp-content/uploads/2021/05/Luong-di-trong-cua-mo-hinh-MVC.webp)

```
User → Request → Controller → Model (xử lý nghiệp vụ, truy vấn dữ liệu)
                                  │
                                  ▼
              Controller ← Dữ liệu kết quả
                  │
                  ▼
                View (render giao diện) → Response → User
```

Nguyên tắc cốt lõi: **Model và View không giao tiếp trực tiếp với nhau**. Mọi trao đổi dữ liệu đều phải qua Controller.

**Ví dụ luồng đăng nhập:**

1. Người dùng nhập email/password và nhấn "Đăng nhập" trên **View**.
2. **View** gửi request đến **Controller**.
3. **Controller** gọi **Model** để kiểm tra thông tin đăng nhập với database.
4. **Model** trả về kết quả (thành công / thất bại) cho **Controller**.
5. **Controller** chọn View phù hợp: chuyển hướng đến trang chủ nếu thành công, hoặc hiển thị thông báo lỗi nếu thất bại.

---

## Ưu và nhược điểm

### Ưu điểm

- **Tách biệt rõ ràng các trách nhiệm (Separation of Concerns):** Mỗi thành phần có nhiệm vụ rõ ràng, giúp mã nguồn dễ đọc và bảo trì.
- **Phát triển song song:** Frontend (View) và backend (Model, Controller) có thể phát triển độc lập với nhau.
- **Dễ kiểm thử:** Model và Controller có thể được kiểm thử đơn vị (unit test) độc lập mà không cần render giao diện.
- **Tái sử dụng:** Cùng một Model có thể phục vụ nhiều View khác nhau (ví dụ: cùng dữ liệu sản phẩm nhưng render khác nhau cho web và mobile API).
- **Hỗ trợ Test-Driven Development (TDD)** tốt nhờ khả năng kiểm thử từng thành phần riêng biệt.

### Nhược điểm

- **Có thể dẫn đến Fat Controller:** Nếu không có Service layer riêng, logic nghiệp vụ phức tạp tích tụ vào Controller, làm giảm khả năng bảo trì.
- **Tăng số lượng file và lớp trừu tượng:** Với ứng dụng nhỏ, MVC có thể là over-engineering — cần thêm nhiều file và cấu trúc so với cách tiếp cận đơn giản hơn.
- **Không có cơ chế xác thực tích hợp sẵn:** Validation logic phải được tự xây dựng hoặc dùng thư viện bên ngoài.
- **Business logic dễ bị phân tán:** Nếu không có quy ước rõ ràng trong team, logic có thể xuất hiện ở cả Controller lẫn Model, gây khó khăn khi debug và refactor.

---

## Ứng dụng MVC trong thực tế

MVC được áp dụng rộng rãi trong hầu hết các framework web phổ biến:

|Framework|Ngôn ngữ|Ghi chú|
|---|---|---|
|ASP.NET Core MVC|C#|Framework chính thức của Microsoft|
|Django|Python|MTV (Model-Template-View) — biến thể của MVC|
|Ruby on Rails|Ruby|Convention over Configuration|
|Spring MVC|Java|Phổ biến trong môi trường doanh nghiệp|
|Gin / Echo|Go|Framework nhẹ, thường kết hợp với Repository pattern|
|Laravel|PHP|Full-stack MVC framework|

---

## Ví dụ cấu trúc thư mục

### C# — ASP.NET Core MVC

```
MyMvcApp/
├── Controllers/
│   ├── HomeController.cs
│   ├── ProductController.cs
│   └── AccountController.cs
├── Models/
│   ├── Product.cs
│   ├── User.cs
│   └── ViewModels/
│       └── ProductViewModel.cs
├── Views/
│   ├── Home/
│   │   ├── Index.cshtml
│   │   └── About.cshtml
│   ├── Product/
│   │   ├── Index.cshtml
│   │   ├── Details.cshtml
│   │   └── Create.cshtml
│   ├── Shared/
│   │   ├── _Layout.cshtml
│   │   └── _ViewStart.cshtml
│   └── _ViewImports.cshtml
├── wwwroot/
│   ├── css/
│   ├── js/
│   └── images/
├── Data/
│   └── ApplicationDbContext.cs
├── Services/
│   └── ProductService.cs
├── appsettings.json
└── Program.cs
```

> `Startup.cs` đã được hợp nhất vào `Program.cs` kể từ ASP.NET Core 6. Đã cập nhật cấu trúc thư mục.

### Go — Gin framework với Repository pattern

```
project-golang/
├── go.mod
├── go.sum
├── init.sql
├── internal/
│   ├── common/
│   ├── config/
│   ├── controller/
│   │   ├── auth/
│   │   │   ├── auth.controller.go
│   │   │   └── dto/
│   │   │       └── auth.dto.go
│   │   └── user/
│   ├── initialize/
│   │   ├── mysql.go
│   │   ├── router.go
│   │   ├── run.go
│   │   └── service.go
│   ├── middleware/
│   │   ├── cors.go
│   │   └── validation.go
│   ├── model/
│   │   ├── auth/
│   │   │   └── auth.entity.go
│   │   └── user/
│   ├── repository/
│   │   ├── auth/
│   │   │   └── auth.repository.go
│   │   ├── repositories.go
│   │   └── user/
│   │       └── user.repository.go
│   ├── router/
│   │   ├── auth/
│   │   │   ├── auth.router.go
│   │   │   ├── enter.go
│   │   │   └── rbac.router.go
│   │   ├── enter.go
│   │   └── user/
│   │       └── enter.go
│   └── service/
│       ├── auth/
│       │   ├── auth.service.go
│       │   └── auth.service.impl.go
│       └── user/
├── pkg/
│   └── response/
│       ├── codeErr.go
│       └── response.go
├── scripts/
├── tests/
└── utils/
```

Lưu ý: Cấu trúc Go ở trên áp dụng MVC kết hợp với **Repository pattern** và **Service layer** — đây là thực hành phổ biến trong backend Go để tránh Fat Controller:

- `controller/`: Nhận HTTP request, gọi service, trả response. Không chứa business logic.
- `service/`: Chứa toàn bộ business logic.
- `repository/`: Chịu trách nhiệm truy vấn database.
- `model/`: Định nghĩa struct dữ liệu (entity).

---

## Tài liệu tham khảo

1. [Tìm hiểu mô hình MVC dành cho người mới bắt đầu: Cấu trúc và ví dụ](https://viblo.asia/p/tim-hieu-mo-hinh-mvc-danh-cho-nguoi-moi-bat-dau-cau-truc-va-vi-du-V3m5WLDyKO7)
2. [Mô hình MVC là gì? Ví dụ và cách ứng dụng MVC vào lập trình](https://vietnix.vn/tim-hieu-mo-hinh-mvc-la-gi/)

---

## Thông tin bổ sung

### 1. MVC và các biến thể trong thực tế

MVC thuần túy thường được mở rộng hoặc biến thể trong thực tế:

- **MVP (Model-View-Presenter):** View thụ động hơn — Presenter chịu trách nhiệm cập nhật View trực tiếp. Phổ biến trong Android development (trước khi MVVM thống trị).
- **MVVM (Model-View-ViewModel):** ViewModel đóng gói trạng thái View và hỗ trợ data binding hai chiều. Phổ biến trong WPF, Angular, Vue.js, Jetpack Compose.
- **MTV (Model-Template-View):** Tên gọi của Django cho mô hình tương tự MVC, trong đó "Template" tương đương với View, và "View" của Django tương đương với Controller.

Hiểu rõ các biến thể này giúp người đọc không bị nhầm lẫn khi đọc tài liệu của các framework khác nhau.

### 2. Service layer — Tránh antipattern Fat Controller

Trong các ứng dụng thực tế, Controller thuần MVC thường không đủ để tổ chức business logic phức tạp. Giải pháp phổ biến là thêm **Service layer** nằm giữa Controller và Model:

```
Request → Controller → Service → Repository → Database
                          ↑
                   (Business logic)
```

- **Controller:** Chỉ xử lý HTTP concerns (parse request, trả response, xử lý lỗi HTTP).
- **Service:** Chứa toàn bộ business logic, có thể tái sử dụng qua nhiều Controller.
- **Repository:** Đóng gói toàn bộ logic truy vấn database, trả về entity cho Service.

Đây là kiến trúc mà cấu trúc Go trong phần ví dụ đang áp dụng.

### 3. Khi nào nên và không nên dùng MVC

**Nên dùng MVC khi:**

- Ứng dụng web hoặc API có quy mô vừa đến lớn.
- Cần phát triển song song giữa frontend và backend.
- Yêu cầu kiểm thử độc lập từng thành phần.

**Cân nhắc kiến trúc khác khi:**

- Ứng dụng rất nhỏ, đơn giản — MVC có thể là overhead không cần thiết.
- Hệ thống phân tán phức tạp (microservices) — có thể cần kiến trúc như Clean Architecture, Hexagonal Architecture, hoặc DDD (Domain-Driven Design) để quản lý độ phức tạp tốt hơn.
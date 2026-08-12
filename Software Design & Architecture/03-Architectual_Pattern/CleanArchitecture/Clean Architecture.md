

Clean architecture v2 · MD
# Clean Architecture: Kiến trúc phần mềm hướng vào trong
 
> **Lưu ý nhỏ trước khi vào bài**: yêu cầu của bạn có kèm theo phần nhắc kiểm tra lại quy trình cấp Access Token từ Refresh Token - đây có vẻ là nội dung còn sót lại từ mẫu yêu cầu trước (tài liệu RefreshToken), không liên quan đến tài liệu Clean Architecture lần này, nên mình sẽ bỏ qua phần đó và tập trung rà soát đúng nội dung Clean Architecture bên dưới.
 
## Mục lục
 
1. [Clean Architecture là gì?](#clean-architecture-là-gì)
2. [Đứng trên vai những gã khổng lồ](#đứng-trên-vai-những-gã-khổng-lồ)
3. [Mục tiêu cốt lõi](#mục-tiêu-cốt-lõi)
4. [Bốn tầng chính trong Clean Architecture](#bốn-tầng-chính-trong-clean-architecture)
5. [Dependency Rule - Quy tắc phụ thuộc](#dependency-rule---quy-tắc-phụ-thuộc)
6. [Ưu điểm của Clean Architecture](#ưu-điểm-của-clean-architecture)
7. [Nhược điểm của Clean Architecture](#nhược-điểm-của-clean-architecture)
8. [Áp dụng Clean Architecture sao cho hợp lý](#áp-dụng-clean-architecture-sao-cho-hợp-lý)
9. [Kết luận](#kết-luận)
10. [Mở rộng](#mở-rộng)
---
 
## Clean Architecture là gì?
 
Clean Architecture là một kiến trúc phần mềm do Robert C. Martin (thường được gọi là "Uncle Bob") giới thiệu năm 2012, dựa trên việc kết hợp nhiều khái niệm, quy tắc và mô hình kiến trúc đã được kiểm chứng trước đó, nhằm đề xuất một cách chuẩn hóa để xây dựng ứng dụng phần mềm.
 
Kiến trúc này được xây dựng trên tư tưởng **"độc lập"**, kết hợp với các nguyên lý thiết kế hướng đối tượng - đại diện tiêu biểu nhất là **Dependency Inversion Principle (DIP)**, một trong năm nguyên lý SOLID. "Độc lập" ở đây nghĩa là logic nghiệp vụ cốt lõi của ứng dụng không phụ thuộc vào framework, thư viện hay công cụ kiểm thử cụ thể nào.
 
## Đứng trên vai những gã khổng lồ
 
Clean Architecture không phải là ý tưởng hoàn toàn mới, mà kế thừa và tổng hợp từ nhiều kiến trúc trước đó đã cùng chia sẻ một mục tiêu: tách biệt rõ ràng giữa logic nghiệp vụ và các chi tiết kỹ thuật bên ngoài. Các kiến trúc có ảnh hưởng trực tiếp đến Clean Architecture bao gồm:
 
- **Hexagonal Architecture (Ports & Adapters)** của Alistair Cockburn.
- **DCI (Data, Context, Interaction)** của James Coplien và Trygve Reenskaug.
- **BCE (Boundary-Control-Entity)** của Ivar Jacobson.
- **Onion Architecture** của Jeffrey Palermo.
Dù khác nhau đôi chút về chi tiết, các kiến trúc này đều có chung mục tiêu cốt lõi:
 
- **Độc lập với framework**: kiến trúc không phụ thuộc vào sự tồn tại của một thư viện/framework cụ thể nào, framework chỉ đóng vai trò là công cụ hỗ trợ.
- **Độc lập với cơ chế phân phối (delivery mechanism)**: logic nghiệp vụ không quan tâm dữ liệu đến từ REST API, CLI, hay message queue.
- **Khả năng cô lập khi kiểm thử**: logic nghiệp vụ có thể được test độc lập, không cần khởi động web server, kết nối database hay bất kỳ thành phần ngoại vi nào.
## Mục tiêu cốt lõi
 
Kiến trúc của Clean Architecture được chia thành 4 tầng (layer), tuân theo một quy tắc phụ thuộc duy nhất: **các tầng bên trong không được biết bất kỳ điều gì về các tầng bên ngoài**. Nói cách khác, quan hệ phụ thuộc trong mã nguồn luôn phải "hướng vào trong".
 
![Sơ đồ các tầng trong Clean Architecture](https://images.viblo.asia/594d086c-73c2-4802-9266-55c7914f3f43.jpeg)
 
Nguyên tắc này chính là Dependency Inversion Principle được áp dụng ở cấp độ kiến trúc, không chỉ ở cấp độ class. Vòng tròn bên trong hoàn toàn không biết gì về các vòng tròn bên ngoài. Khi dữ liệu được truyền qua một ranh giới (boundary) giữa các tầng, dữ liệu đó luôn được chuyển đổi sang dạng thuận tiện nhất cho tầng phía trong tiếp nhận, thay vì buộc tầng trong phải hiểu định dạng của tầng ngoài.
 
## Bốn tầng chính trong Clean Architecture
 
### 1. Entities (Domain Layer)
 
Đây là tầng trong cùng và quan trọng nhất - nơi chứa các thực thể (entity) đại diện cho đối tượng nghiệp vụ cùng các quy tắc nghiệp vụ cốt lõi (core business rules) gắn liền với đối tượng đó. Một entity có thể là một object hoặc một cụm object liên quan - ví dụ trong use case tạo người dùng, entity chính là đối tượng `User` cùng các business rule gắn với nó (định dạng email hợp lệ, độ dài mật khẩu tối thiểu...).
 
Tầng này không phụ thuộc vào bất kỳ framework nào, có thể chạy và test độc lập mà không cần bất kỳ thành phần hạ tầng nào (web server, database). Đây là lý do vì sao Entities dễ test, dễ bảo trì và phát triển nhất trong bốn tầng.
 
### 2. Use Cases (Application Layer)
 
Tầng này chứa logic nghiệp vụ ở cấp độ ứng dụng cụ thể (application-specific business rules), mô tả luồng xử lý của từng use case: tương tác với Entities như load, xử lý, lưu lại entity. Có thể hình dung Use Case như một **orchestrator** điều phối luồng xử lý của một request.
 
Use Case không quan tâm dữ liệu đến từ đâu, được truyền qua giao thức nào, hay hiển thị ra sao ở tầng ngoài - đó là trách nhiệm của tầng Interface Adapters. Entities và Use Cases cùng nhau tạo thành phần được gọi là **core business logic** - phần lõi hoàn toàn độc lập với framework, UI hay database.
 
### 3. Interface Adapters (Presentation Layer)
 
Tầng này chịu trách nhiệm chuyển đổi định dạng dữ liệu giữa bên ngoài (web request, dữ liệu từ database) và bên trong (application/domain layer), theo cả hai chiều. Các thành phần điển hình ở tầng này gồm: controller, presenter, các object chuyển đổi dữ liệu (DTO), ví dụ `CreateUserCommand` hoặc `UserDTO` trong use case tạo người dùng.
 
Ví dụ minh họa: thông tin người dùng có thể chứa nhiều trường nhạy cảm (email, số điện thoại, địa chỉ), nhưng không phải lúc nào giao diện cũng cần hiển thị đầy đủ tất cả - Interface Adapters chịu trách nhiệm định dạng lại dữ liệu sao cho phù hợp với từng nhu cầu hiển thị hoặc lưu trữ cụ thể, mà không cần quan tâm dữ liệu sẽ được hiển thị chi tiết như thế nào (đó là việc của tầng Frameworks & Drivers).
 
### 4. Frameworks & Drivers (Infrastructure Layer)
 
Đây là tầng ngoài cùng, tập hợp các công cụ cụ thể phục vụ nhu cầu thực tế: giao diện web, thiết bị, hệ quản trị cơ sở dữ liệu... Đây là tầng "nhẹ" nhất về lượng code cần viết, vì phần lớn logic đã nằm ở các tầng trong.
 
Trên thực tế, đây là nơi duy nhất "biết tất cả" - vì tầng này chịu trách nhiệm khởi tạo các đối tượng cụ thể (implementation) cho các tầng bên trong thông qua Interface, quá trình này thường được gọi là **Dependency Injection** hoặc **Setup Dependencies**.
 
> **Quy tắc bắt buộc**: mã nguồn của một class thuộc tầng trong **không được phép** tham chiếu trực tiếp đến mã nguồn của một class thuộc tầng ngoài hơn nó.
 
## Dependency Rule - Quy tắc phụ thuộc
 
Dependency Rule là khía cạnh quan trọng nhất, đóng vai trò then chốt để Clean Architecture đạt được tính linh hoạt, dễ bảo trì và mở rộng. Quy tắc này phát biểu cụ thể như sau:
 
- Quan hệ phụ thuộc trong mã nguồn (source code dependency) chỉ được phép hướng vào trong, hướng về phía các chính sách cấp cao hơn (higher-level policies).
- Không một thành phần nào ở vòng tròn trong được phép biết bất cứ điều gì về một thành phần ở vòng tròn ngoài - kể cả tên hàm, tên class, tên biến, hay bất kỳ thực thể phần mềm nào được đặt tên ở tầng ngoài.
- Tương tự, các định dạng dữ liệu (data format) được khai báo ở tầng ngoài (ví dụ struct ánh xạ với bảng database, hay JSON response) không nên được sử dụng trực tiếp ở tầng trong.
Để tuân thủ quy tắc này trong khi tầng trong vẫn cần "gọi" đến chức năng của tầng ngoài (ví dụ Use Case cần lưu dữ liệu xuống database), Clean Architecture dùng kỹ thuật đảo ngược phụ thuộc thông qua **Interface**: tầng trong định nghĩa interface (ví dụ `UserRepository`), còn tầng ngoài chịu trách nhiệm implement interface đó (ví dụ `PostgresUserRepository`). Nhờ vậy, tầng trong chỉ phụ thuộc vào interface do chính nó định nghĩa, không phụ thuộc vào chi tiết triển khai cụ thể ở tầng ngoài.
 
![Ví dụ minh họa use case tạo user tuân theo Dependency Rule](https://tech.cybozu.vn/static/19832d1ef9943c85f424564b44cf457f/d00b9/create-user-usecase-explain.webp)
 
## Ưu điểm của Clean Architecture
 
- **Mạch lạc, dễ nhìn (Screaming Architecture)**: chỉ cần nhìn vào cấu trúc thư mục/package cũng có thể hiểu được mục đích và cơ chế hoạt động của ứng dụng, thay vì phải đọc chi tiết từng file.
- **Linh hoạt**: logic nghiệp vụ độc lập, không phụ thuộc vào framework, database hay application server cụ thể - có thể thay đổi công nghệ hạ tầng mà không cần viết lại business logic.
- **Dễ kiểm thử (testable)**: nhờ tách biệt qua Interface, việc viết mock/stub để test từng tầng độc lập trở nên đơn giản, không cần khởi động toàn bộ hệ thống.
- **Chia để trị hiệu quả trong ứng dụng lớn**: code ở đúng tầng của nó, hạn chế tình trạng "code ở đâu cũng được, chạy được là được" khiến hệ thống dần trở nên khó kiểm soát.
- **Dễ bảo trì và mở rộng**: vì các tầng độc lập thông qua Interface, việc mở rộng hoặc thay đổi một tầng ít ảnh hưởng đến các tầng còn lại, hạn chế breaking change và giảm nhu cầu refactor diện rộng.
## Nhược điểm của Clean Architecture
 
- **Cồng kềnh, phức tạp hơn**: cần viết nhiều class/interface hơn so với cách viết thông thường. Với ứng dụng đơn giản, ít tính năng hoặc vòng đời ngắn, việc áp dụng đầy đủ Clean Architecture có thể tạo ra sự phức tạp không cần thiết.
- **Tính trừu tượng cao (indirection)**: việc gọi qua nhiều lớp interface trung gian giúp code linh hoạt hơn, nhưng cũng đồng nghĩa với việc phải "nhảy" qua nhiều lớp trừu tượng khi đọc/debug code, và có thể ảnh hưởng nhẹ đến hiệu năng runtime so với gọi trực tiếp.
- **Không phù hợp để code nhanh kiểu "mì ăn liền"**: do tuân thủ Dependency Inversion, mọi tương tác giữa các tầng đều cần đi qua Interface được định nghĩa trước, khó áp dụng lối viết code tùy tiện, thiếu kế hoạch.
- **Đường cong học tập cao hơn với người mới**: cần hiểu rõ về SOLID, đặc biệt là Dependency Inversion, trước khi áp dụng đúng cách; áp dụng sai có thể dẫn đến kiến trúc phức tạp mà không mang lại lợi ích tương xứng.
## Áp dụng Clean Architecture sao cho hợp lý
 
Không phải ứng dụng nào cũng cần đầy đủ 4 tầng Entities, Use Cases, Interface Adapters và Frameworks & Drivers. Trong thực tế, phần lớn kỹ sư chỉ tham khảo tư tưởng của Clean Architecture để xây dựng một kiến trúc phù hợp hơn với quy mô dự án của mình, thay vì áp dụng máy móc 100% mô hình gốc.
 
Điều cốt lõi cần ghi nhớ: thay vì gom toàn bộ business logic vào một class hoặc một hàm duy nhất, nên chia thành các layer với trách nhiệm riêng biệt, giao tiếp với nhau thông qua Interface thay vì phụ thuộc trực tiếp vào implementation cụ thể.
 
Một đề xuất đơn giản hóa thường gặp khi xây dựng REST API là chỉ dùng **3 tầng cơ bản**: **Transport** (nhận request, trả response), **Business** (xử lý logic nghiệp vụ), **Repository/Storage** (thao tác với database). Đây không phải là best practice chính thức của Clean Architecture, mà là một cách áp dụng thực tế, đơn giản hóa cho các service REST API thông thường.
 
Trước khi áp dụng, một hàm xử lý business thường trông như sau - toàn bộ logic dồn vào một hàm duy nhất, khó test và khó mở rộng:
 
![Ví dụ hàm xử lý gộp toàn bộ logic trước khi áp dụng Clean Architecture](https://statics.cdn.200lab.io/2022/05/huge-function-before-clean-architecture.png?width=800)
 
Sau khi áp dụng nguyên tắc tách tầng và "hướng phụ thuộc vào trong" ở mức đơn giản hóa:
 
![Mô hình 3 tầng đơn giản hóa từ Clean Architecture](https://statics.cdn.200lab.io/2022/05/simple-layers-clean-architecture-1.png?width=800)
 
Ví dụ cụ thể với API cập nhật sản phẩm (Update Product):
 
![Ví dụ luồng xử lý API Update Product theo 3 tầng](https://statics.cdn.200lab.io/2022/05/example-basic-clean-architecture-update-product-api-1.png?width=800)
 
Ở tầng Storage, có thể triển khai song song hai lựa chọn MongoDB hoặc MySQL - cả hai cùng tuân theo một Interface chung (giống nhau về tên hàm, tham số đầu vào/đầu ra), chỉ khác nhau ở phần implementation cho từng hệ database. Điểm mấu chốt là tầng Business không bao giờ gọi trực tiếp vào implementation cụ thể, mà luôn gọi thông qua Interface.
 
### Ví dụ minh họa (Go)
 
```go
// Tầng Business (Use Case) chỉ phụ thuộc vào Interface do chính nó định nghĩa,
// không biết gì về việc dữ liệu được lưu ở PostgreSQL, MongoDB hay nơi nào khác.
type ProductRepository interface {
    FindByID(ctx context.Context, id string) (*Product, error)
    Update(ctx context.Context, p *Product) error
}
 
type ProductService struct {
    repo ProductRepository
}
 
func (s *ProductService) UpdateProductPrice(ctx context.Context, id string, newPrice float64) error {
    product, err := s.repo.FindByID(ctx, id)
    if err != nil {
        return err
    }
    if err := product.ChangePrice(newPrice); err != nil { // business rule nằm trong Entity
        return err
    }
    return s.repo.Update(ctx, product)
}
 
// Tầng Storage: implementation cụ thể cho PostgreSQL, nằm ở tầng ngoài,
// implement lại Interface ProductRepository đã định nghĩa ở tầng trong.
type PostgresProductRepository struct {
    db *sql.DB
}
 
func (r *PostgresProductRepository) FindByID(ctx context.Context, id string) (*Product, error) {
    // triển khai truy vấn cụ thể với PostgreSQL/sqlc
    return nil, nil
}
 
func (r *PostgresProductRepository) Update(ctx context.Context, p *Product) error {
    // triển khai câu lệnh UPDATE cụ thể
    return nil
}
```
 
Nhờ `ProductService` chỉ phụ thuộc vào interface `ProductRepository`, việc viết unit test cho logic nghiệp vụ trở nên đơn giản: chỉ cần tạo một mock implementation của `ProductRepository` (ví dụ bằng `uber-go/mock`), hoàn toàn không cần kết nối database thật.
 
## Kết luận
 
Clean Architecture không phải là một công thức bắt buộc phải tuân thủ tuyệt đối, mà là một tập hợp nguyên tắc giúp tách biệt logic nghiệp vụ khỏi các chi tiết kỹ thuật dễ thay đổi. Ngay cả Uncle Bob cũng từng nhấn mạnh rằng: cuối cùng, "Clean Architecture" cũng chỉ là một cái tên - việc không tuân thủ triệt để kiến trúc này không hẳn là xấu, và việc tuân thủ máy móc cũng không hẳn là tốt nếu nó không phù hợp với quy mô và bối cảnh thực tế của dự án. Điều quan trọng nhất vẫn là hiểu đúng bản chất của Dependency Rule và áp dụng linh hoạt tùy theo nhu cầu.
 
### Mở rộng
 
Một số hướng tìm hiểu thêm để nâng cao kiến thức về chủ đề này:
 
- **SOLID Principles**: đặc biệt là Dependency Inversion Principle (DIP) và Interface Segregation Principle (ISP) - nền tảng lý thuyết trực tiếp của Clean Architecture.
- **Hexagonal Architecture (Ports & Adapters)** và **Onion Architecture**: tìm đọc bài viết gốc của Alistair Cockburn và Jeffrey Palermo để hiểu rõ nguồn gốc tư tưởng mà Clean Architecture kế thừa.
- **Domain-Driven Design (DDD)**: khái niệm Entity, Value Object, Aggregate trong DDD có liên hệ chặt chẽ với tầng Entities của Clean Architecture, đặc biệt hữu ích khi thiết kế domain layer phức tạp.
- **Package by Feature vs Package by Layer trong Go**: nghiên cứu cách tổ chức thư mục dự án Go theo Clean Architecture sao cho vẫn giữ được tính đơn giản, tránh over-engineering với dự án nhỏ.
- **Wire hoặc fx**: các công cụ Dependency Injection phổ biến trong hệ sinh thái Go, giúp tự động hóa việc khởi tạo và kết nối các tầng ở Frameworks & Drivers layer thay vì inject thủ công.
- **So sánh Clean Architecture với kiến trúc đơn giản hơn (MVC truyền thống)**: hiểu rõ khi nào nên đầu tư vào Clean Architecture đầy đủ và khi nào một kiến trúc đơn giản hơn là lựa chọn hợp lý hơn, tránh áp dụng kiến trúc phức tạp cho dự án không thực sự cần đến.
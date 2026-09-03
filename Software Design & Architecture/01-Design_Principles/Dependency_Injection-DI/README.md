# Giới thiệu Dependency Injection

## Mục lục

1. [Đặt vấn đề](https://claude.ai/chat/d6f7164f-e12c-439e-8152-336e1aa0d9c5#1-%C4%91%E1%BA%B7t-v%E1%BA%A5n-%C4%91%E1%BB%81)
2. [Dependency Injection (DI) là gì?](https://claude.ai/chat/d6f7164f-e12c-439e-8152-336e1aa0d9c5#2-dependency-injection-di-l%C3%A0-g%C3%AC)
3. [Dependency Inversion Principle (DIP) — nền tảng của DI](https://claude.ai/chat/d6f7164f-e12c-439e-8152-336e1aa0d9c5#3-dependency-inversion-principle-dip--n%E1%BB%81n-t%E1%BA%A3ng-c%E1%BB%A7a-di)
4. [Khi nào nên dùng DI?](https://claude.ai/chat/d6f7164f-e12c-439e-8152-336e1aa0d9c5#4-khi-n%C3%A0o-n%C3%AAn-d%C3%B9ng-di)
5. [Ba loại Dependency Injection](https://claude.ai/chat/d6f7164f-e12c-439e-8152-336e1aa0d9c5#5-ba-lo%E1%BA%A1i-dependency-injection)
6. [Trách nhiệm của Dependency Injection](https://claude.ai/chat/d6f7164f-e12c-439e-8152-336e1aa0d9c5#6-tr%C3%A1ch-nhi%E1%BB%87m-c%E1%BB%A7a-dependency-injection)
7. [Nguyên tắc hoạt động của DI trong thực tế](https://claude.ai/chat/d6f7164f-e12c-439e-8152-336e1aa0d9c5#7-nguy%C3%AAn-t%E1%BA%AFc-ho%E1%BA%A1t-%C4%91%E1%BB%99ng-c%E1%BB%A7a-di-trong-th%E1%BB%B1c-t%E1%BA%BF)
8. [Lợi ích và bất lợi](https://claude.ai/chat/d6f7164f-e12c-439e-8152-336e1aa0d9c5#8-l%E1%BB%A3i-%C3%ADch-v%C3%A0-b%E1%BA%A5t-l%E1%BB%A3i)
9. [Tổng kết](https://claude.ai/chat/d6f7164f-e12c-439e-8152-336e1aa0d9c5#9-t%E1%BB%95ng-k%E1%BA%BFt)
10. [Mở rộng](https://claude.ai/chat/d6f7164f-e12c-439e-8152-336e1aa0d9c5#m%E1%BB%9F-r%E1%BB%99ng)

> **Phần thực hành:** Xem ví dụ áp dụng DI trong Go tại tài liệu `DI_Golang_Example.md`.

---

## 1. Đặt vấn đề

Với những dự án có độ phức tạp cao, ngoài việc thiết kế tính năng cho ứng dụng, việc tổ chức code luôn là vấn đề được đặt lên hàng đầu. Tổ chức code tốt giúp lập trình viên dễ dàng bảo trì, cũng như mở rộng code về sau.

Để tiết kiệm chi phí và thời gian cho công đoạn này nhưng vẫn đem lại hiệu quả cao, việc nắm vững các design pattern sẽ giúp ích rất nhiều. **Dependency Injection** là một dạng design pattern được thiết kế với mục đích ngăn chặn sự phụ thuộc trực tiếp giữa các class, giúp code dễ hiểu hơn, trực quan hơn, phục vụ cho mục đích bảo trì và nâng cấp code.

Nếu sau khi đọc xong thuật ngữ trên bạn vẫn chưa hiểu gì thì hãy coi đây là chuyện bình thường — các định nghĩa về một [design pattern](https://vi.wikipedia.org/wiki/M%E1%BA%ABu_thi%E1%BA%BFt_k%E1%BA%BF_ph%E1%BA%A7n_m%E1%BB%81m) thường khá trừu tượng (abstract), không đi vào cụ thể. Cách áp dụng Dependency Injection sẽ phụ thuộc vào từng tình huống cụ thể, cũng như kỹ năng của developer.

---

## 2. Dependency Injection (DI) là gì?

Theo định nghĩa của Wikipedia:

> Dependency injection là một kĩ thuật trong đó một object (hoặc một static method) cung cấp các dependencies của một object khác. Một dependency là một object mà object kia có thể sử dụng (một service).

Định nghĩa trên khá khó hiểu ở lần đọc đầu tiên, nên hãy cùng phân tích từng phần.

**Dependency (hay dependent)** nghĩa là "phụ thuộc vào sự hỗ trợ của một cái gì đó". Ví dụ đời thường: nếu chúng ta phụ thuộc quá nhiều vào smartphone để hoàn thành công việc hàng ngày, có thể nói chúng ta đang "dependent" (phụ thuộc) vào smartphone.

Trong lập trình, khi class A sử dụng một số chức năng của class B, ta nói **class A có quan hệ phụ thuộc (dependency) với class B**.

![](https://images.viblo.asia/dd6d54c4-7fd4-48a0-a2ba-ec7ee0651949.jpeg)

Thông thường, class A sẽ tự tay khởi tạo (new) đối tượng B mà nó cần. **Dependency Injection** là việc chuyển giao nhiệm vụ khởi tạo object đó cho một thành phần khác (thay vì để class A tự làm), và class A chỉ việc trực tiếp sử dụng dependency được cung cấp sẵn.

![](https://images.viblo.asia/1057e16d-643e-4512-a6cc-2b6ce9cb9898.png)

---

## 3. Dependency Inversion Principle (DIP) — nền tảng của DI

Trước khi đi tiếp, bạn nên dành chút thời gian ôn lại nguyên lý [SOLID](https://topdev.vn/blog/solid-la-gi/) trong [OOP](https://gpcoder.com/2232-4-tinh-chat-cua-lap-trinh-huong-doi-tuong-trong-java/) — vì DI bắt nguồn trực tiếp từ nguyên lý cuối cùng trong SOLID, ứng với chữ **D**: **Dependency Inversion Principle (DIP)** — nguyên lý đảo ngược sự phụ thuộc. Nội dung của nguyên lý này:

- Các module cấp cao không nên phụ thuộc vào các module cấp thấp. Cả hai nên phụ thuộc vào **abstraction** (interface/lớp trừu tượng).
- Interface (abstraction) không nên phụ thuộc vào chi tiết triển khai (implementation); ngược lại, chi tiết triển khai nên phụ thuộc vào interface. Các class giao tiếp với nhau thông qua interface, không phải trực tiếp qua implementation.

**Ba khái niệm dễ gây nhầm lẫn với nhau:** [Dependency Inversion Principle (DIP), Inversion of Control (IoC), và Dependency Injection (DI)](https://topdev.vn/blog/inversion-of-control-va-dependency-injection/) — cả ba đều hướng đến cùng một mục tiêu: tạo ra ứng dụng ít kết dính (loosely coupled), dễ mở rộng (flexible), giúp lập trình viên tập trung vào business logic thay vì lo lắng chuyện khởi tạo object. Sự khác biệt nằm ở mức độ trừu tượng:

```
IoC (Inversion of Control)  → là một NGUYÊN LÝ/hướng đi tổng quát
DIP (Dependency Inversion)  → là sự CỤ THỂ HÓA của nguyên lý đó, áp dụng cho quan hệ phụ thuộc giữa các module
DI  (Dependency Injection)  → là một KỸ THUẬT hiện thực hóa DIP trong code thực tế
```

---

## 4. Khi nào nên dùng DI?

Một số tình huống nên cân nhắc dùng DI:

- Khi cần inject các giá trị cấu hình (configuration) cho một hoặc nhiều module khác nhau.
- Khi cần inject cùng một dependency cho nhiều module khác nhau.
- Khi cần dùng một service được cung cấp sẵn bởi container (VD: kết nối database, logger dùng chung).
- Khi cần tách biệt cách triển khai dependency giữa các môi trường khác nhau. Chẳng hạn, ở môi trường dev chỉ cần log lại việc gửi mail (không gửi thật), còn ở môi trường production cần gửi mail qua một API thật sự — miễn là cả hai đều implement cùng một interface, phần code còn lại của ứng dụng không cần thay đổi gì.

---

## 5. Ba loại Dependency Injection

Về cơ bản, có 3 cách triển khai Dependency Injection:

1. **Constructor injection**: dependency được cung cấp thông qua constructor (hàm khởi tạo) của class.
2. **Setter injection**: class cung cấp một setter method để các thành phần khác gán (set) dependency vào sau khi object đã được khởi tạo.
3. **Interface injection**: dependency tự cung cấp một hàm injector, dùng để "tiêm" chính nó vào bất kỳ client nào được truyền vào — client phải implement một interface có sẵn setter method dành riêng cho việc nhận dependency.

> Ở tài liệu thực hành Go (`DI_Golang_Example.md`), bạn sẽ thấy ví dụ minh họa cụ thể cho **Constructor injection** và **Setter injection** — đây cũng là hai kiểu phổ biến nhất trong Go, vì ngôn ngữ này không hỗ trợ interface injection theo đúng nghĩa (không có cơ chế reflection/annotation tự động tương tự Java).

---

## 6. Trách nhiệm của Dependency Injection

Tổng hợp lại, trách nhiệm cốt lõi của DI gồm:

1. Tạo ra các object.
2. Quản lý sự phụ thuộc (dependencies) giữa các đối tượng.
3. Cung cấp (inject) các dependency được yêu cầu cho đối tượng, từ bên ngoài đối tượng đó.

Nhờ vậy, nếu trong tương lai cách triển khai của một dependency thay đổi (VD: đổi từ gửi SMS sang gửi Email), thành phần chịu trách nhiệm inject chỉ cần cấp lại đúng object cần thiết — class sử dụng dependency đó hoàn toàn không cần sửa code.

---

## 7. Nguyên tắc hoạt động của DI trong thực tế

- Các module không giao tiếp trực tiếp với nhau, mà thông qua interface. Module cấp thấp sẽ implement interface; module cấp cao gọi module cấp thấp thông qua interface đó, không quan tâm chi tiết triển khai bên trong.
- Việc khởi tạo các module cấp thấp thường do một **DI Container / IoC Container** đảm nhiệm (trong các framework có hỗ trợ sẵn).
- Việc "module nào gắn với interface nào" thường được cấu hình qua file properties, file XML, hoặc Annotation. Annotation là cách phổ biến trong nhiều framework — ví dụ `@Inject` với [CDI](http://www.cdi-spec.org/), `@Autowired` với [Spring](https://spring.io/) (Java), hay `@ManagedProperty` với [JSF](https://www.oracle.com/technetwork/java/javaee/javaserverfaces-139869.html).

> **Lưu ý cho lập trình viên Go:** các ví dụ Annotation ở trên đều thuộc hệ sinh thái Java — Go **không có** cơ chế annotation/reflection tự động tương tự để "tự động" wiring dependency. Trong Go, việc inject dependency thường được thực hiện một trong hai cách: **thủ công, tường minh** (truyền trực tiếp qua constructor function hoặc gán trực tiếp vào struct field — như ví dụ ở `DI_Golang_Example.md`), hoặc thông qua các thư viện DI dành riêng cho Go như [Google Wire](https://github.com/google/wire) (sinh code lúc compile-time, không dùng reflection) hoặc [Uber Fx](https://github.com/uber-go/fx) (dùng dependency graph lúc runtime). Cách tiếp cận tường minh (không dùng container ẩn) thường được cộng đồng Go ưa chuộng hơn, vì phù hợp với triết lý "rõ ràng hơn là ngầm định" (explicit over implicit) của ngôn ngữ này.

---

## 8. Lợi ích và bất lợi

### Lợi ích

1. Giúp viết Unit Test dễ dàng hơn — vì có thể thay dependency thật bằng một bản giả lập (mock/stub) khi test, không cần phụ thuộc vào hệ thống thật (database, API bên ngoài...).
2. Giảm boilerplate code, vì việc khởi tạo dependency được đảm nhiệm bởi một thành phần khác thay vì lặp lại logic khởi tạo ở nhiều nơi.
3. Mở rộng dự án dễ dàng hơn — thêm một cách triển khai mới của dependency (VD: thêm `PushNotificationService` bên cạnh `SMSService`/`EmailService`) không đòi hỏi sửa code ở nơi sử dụng dependency.
4. Giúp liên kết lỏng (loose coupling) giữa các thành phần trong dự án, đúng với tinh thần của DIP đã nêu ở mục 3.

### Bất lợi

1. Khá phức tạp để học ở giai đoạn đầu; nếu lạm dụng quá đà (inject mọi thứ, kể cả những phần không thực sự cần thiết) có thể khiến code khó theo dõi hơn thay vì dễ hơn.
2. Một số lỗi vốn có thể phát hiện ngay lúc biên dịch (compile time) lại bị đẩy sang lúc chạy chương trình (runtime) — đặc biệt đúng với các framework DI dùng reflection/container ẩn (ít gặp hơn trong Go nếu dùng cách tường minh hoặc Google Wire, vì Wire sinh code lúc compile-time).
3. Có thể ảnh hưởng đến chức năng auto-complete hoặc "Find references" của một số IDE, do dependency thực tế chỉ được xác định lúc runtime (với các cách triển khai dùng container/reflection).

---

## 9. Tổng kết

- **Dependency Injection** là một kỹ thuật hiện thực hóa nguyên lý **DIP** (chữ D trong SOLID) — chuyển việc khởi tạo dependency ra khỏi class sử dụng nó, giúp giảm kết dính giữa các thành phần.
- Ba khái niệm **IoC → DIP → DI** có quan hệ từ tổng quát đến cụ thể: IoC là nguyên lý, DIP là sự cụ thể hóa cho quan hệ phụ thuộc, DI là kỹ thuật hiện thực hóa trong code.
- Có 3 kiểu triển khai DI: **Constructor injection**, **Setter injection**, **Interface injection** — trong Go, hai kiểu đầu phổ biến nhất.
- Lợi ích lớn nhất của DI là **khả năng test dễ dàng hơn** (thay dependency thật bằng mock) và **liên kết lỏng** giữa các module — nhưng cần tránh lạm dụng để không làm code khó hiểu hơn.

---

### Mở rộng

- **Google Wire vs Uber Fx**: hai hướng tiếp cận DI phổ biến nhất trong hệ sinh thái Go hiện nay — một bên sinh code lúc compile-time (Wire), một bên dùng dependency graph lúc runtime (Fx). Đáng tìm hiểu sự đánh đổi giữa hai cách này.
- **Service Locator pattern**: một pattern hay bị nhầm lẫn với DI — cả hai đều nhằm giảm việc class tự khởi tạo dependency, nhưng khác nhau ở chỗ Service Locator để class chủ động "hỏi xin" dependency từ một registry trung tâm, trong khi DI "đẩy" dependency vào class từ bên ngoài mà class không cần biết registry tồn tại.
- **Mocking trong Unit Test với Go**: tìm hiểu thư viện [testify/mock](https://github.com/stretchr/testify) hoặc [gomock](https://github.com/uber-go/mock) để tạo mock implementation cho interface — đây chính là ứng dụng thực tế của lợi ích "dễ viết Unit Test hơn" đã nêu ở mục 8.
- **Dependency Injection Container trong các ngôn ngữ khác**: so sánh cách Spring (Java), NestJS (TypeScript), hay .NET Core tự động hóa việc wiring dependency bằng container, để hiểu rõ hơn vì sao Go lại chọn hướng tiếp cận tường minh, thủ công hơn.
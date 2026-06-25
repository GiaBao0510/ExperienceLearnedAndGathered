## **Mediator Pattern là gì?**

![](https://images.viblo.asia/d6bebfc5-b8d7-4895-979a-a3c9eaabaffa.png)

**Mediator pattern (Mẫu thiết kế trung gian)** là một trong những mẫu thiết kết thuộc nhóm **Behavioral Pattern**. Được sử dụng để quản lý giao tiếp giữa các đối tượng.
**Mediator pattern** hoạt động như là một cầu nối.
**Mục tiêu** là ==giảm sự phụ thuộc trực tiếp giữa các thành phần==, đối tượng trong hệ thống.
Chúng tương tác với nhau thông qua một đối tượng trung gian (Mediator).

---
## **Tại sao Mediator lại quan trọng?**

Khi một hệ thống phức tạp, nơi mà có nhiều đối tượng cần giao tiếp qua lại với nhau. Thì sao đây là các lý do nên cần dùng:
1. **Giảm sự phụ thuộc giữa các thành phần:** Các đối tượng không cần biết chi tiết của nhau, thay vào đó chỉ cần giao tiếp với Mediator.
2. **Tăng tính linh hoạt và khả năng mở rộng:** Với Mediator, việc thêm hoặc sửa đổi các đối tượng dễ dàng hơn mà không ảnh hưởng đển toàn bộ hệ thống.
3. **Cải thiện khả năng bảo trì:** Giảm sự kết nối lẫn nhau giúp mã nguồn dễ hiểu và bảo trì hơn, đặc biệt là với các dự án dài hạn.

---
## **Khi nào nên sử dụng?**

Thường thì Mediator áp dụng trong các tình huống sau:

- **Ứng dụng GUI phức tạp:** Trong các ứng dụng giao diện người dùng, có nhiều người dùng giao tiếp với nhau
- **Hệ thống có nhiều module:** Nếu có nhiều Module cần tương tác với nhau, Mediator giúp chúng tách biệt và dễ bảo trì hơn.

---
## **Các Thành phần?**

Mediator pattern bao gồm các thành phần chính như sau:
###### 1. **Mediator (Đối tượng trung gian):** 
Định nghĩa các phương thức để các đối tượng giao tiếp với nhau
###### 2. **Concrete Mediator (Đối tượng trung gian cụ thể):**
- Để đóng gói mối quan hệ giữa các componet khác nhau. Các **Concrete Mediator** thường quản lý các tham chiếu của các component kể cả quản lý vòng đời của chúng. 
- Thực thi các phương thức giao tiếp được định nghĩa trong Mediator, chịu trách quản lý và điều phối các đối tượng.
###### 3. **Colleagues (Các đối tượng giao tiếp):** 
Là các đối tượng cần giao tiếp với nhau, nhưng không giao tiếp trực tiếp với nhau mà phải thông qua Mediator, được khai báo với kiểu là Mediator interface. Các đối tượng không cần phải quan tâm các lớp thực sự của Mediator.

---
## **Ví dụ:**

---
## **Ưu & nhược điểm:**

**Ưu điểm:**
- **Giúp đảm bảo nguyên tắt Signle Repository principle(SRP):** Tách thành phần giao tiếp giữa các thành phần (component) khác sang một nơi khác (*Nghĩa là*: trích xuất sự liên lạc giữa các thành phần vào trong một nơi duy nhất => dễ bảo trì )
- **Giúp đảm bảo nguyên tắt Open/Closed principle(OCP):** bằng cách thêm Mediator mới mà không ảnh hưởng đến các componet hiện có.
- Tái sử dụng component dễ dàng hơn
- Giảm khớp nối giữa các componet
- Giúp đơn giản hóa giao tiếp giữa các đối tượng. ==Mediator sẽ thay thế cho quan hệ nhiều nhiều (many to many) thành quan hệ một nhiều (one to many) giữa các component với một mediator==

**Nhược điểm:**
- Phức tạp hóa hệ thống nếu không thiết kế cẩn thận.
- Mediator Pattern có thể khiến cho Mediator trở nên quá tải khi nhiều đối tượng ([Gob Object](https://vi.wikipedia.org/wiki/%C4%90%E1%BB%91i_t%C6%B0%E1%BB%A3ng_th%C6%B0%E1%BB%A3ng_%C4%91%E1%BA%BF_(L%E1%BA%ADp_tr%C3%ACnh_m%C3%A1y_t%C3%ADnh)#:~:text=Trong%20l%E1%BA%ADp%20tr%C3%ACnh%20h%C6%B0%E1%BB%9Bng%20%C4%91%E1%BB%91i,th%E1%BB%A9c%20(anti%2Dpattern).))

---
## **Tham khảo:**

1.  [Mediator Pattern Là Gì? Tìm Hiểu Về Mediator Design Pattern](https://t3h.com.vn/tin-tuc/mediator-pattern-la-gi-tim-hieu-ve-mediator-design-pattern#)
2. [Mediator Design Pattern – Collaborate via me](https://topdev.vn/blog/mediator-design-pattern-collaborate-via-me/)
3. [ Mediator Design Pattern - Trợ thủ đắc lực của Developers](https://viblo.asia/p/mediator-design-pattern-tro-thu-dac-luc-cua-developers-m68Z0jVj5kG)
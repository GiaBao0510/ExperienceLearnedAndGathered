## **HTTP Request:**

**HTTP request** có tất cả 9 loại method, 2 loại được sử dụng phổ biến là **GET** và **POST**.
- **GET:** được sử dụng để lấy thông tin từ server theo URI đã cung cấp.
- **HEAD:** giống với GET nhưng response trả về không có body, chỉ có header.
- **POST:** gửi thông tin tới server thông qua các biểu mẫu http.
- **PUT:** ghi đè tất cả thông tin của đối tượng với những gì được gửi lên.
- **PATCH:** ghi đè các thông tin được thay đổi của đối tượng.
- **DELETE:** xóa tài nguyên trên server.
- **CONNECT:** thiết lập một kết nối tới server theo URI
- **OPTIONS:** mô tả các tùy chọn giao tiếp cho resource.
- **TRACE:** Thực hiện một bài tét loop - back đường dẫn đến resource.

---
## **Authorization:**

Hiện tại có 3 cơ chế Authorizie chính:
- HTTP Basic.
- JSON Web Token (JWT).
- OAith2
Tùy thuộc vào server mà chọn loại Authorization với mức độ phù hợp.

---
## **Quản lý version của API:**

- Khi thiết lập api cho app ios hay client side, chúng ta nên đặt version cho các api. Ví dụ như endpoint sau: **api/v1/users**
- Điều này sẽ giúp cho hệ thống sau khi nâng cấp lên version mới vẫn hõ trợ các api của version cũ, cũng như giúp cho việc sửa chữa và bảo trì dễ dàng hơn.
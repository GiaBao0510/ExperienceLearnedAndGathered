### **HTTP là gì?**

**HTTP (Hypertext Transfer Protocol)**, là nền tảng giao tiếp giữa *client-server* trên các ứng dụng web. Các phương thức HTTP đóng vai trò quan trọng trong việc xác định cách thức giao tiếp xảy ra. Mỗi phương thức có một mục đích cụ thể và hiểu được sắc thái (nuances) của chúng là điều cần thiết để phát triển web hiệu quả.

---

### **Các phương thức HTTP (HTTP Methods):**

Các phương thức **HTTP** còn được gọi là **HTTP verbs**, chịu trách nhiệm xác định các hành động có thể được thực hiện trên các tài nguyên xác định bởi mã định danh tài nguyên **URI (Uniform Resource Identifiers)**.

HTTP định nghĩa một số phương thức, mỗi phương thức phục vụ một mục đích cụ thể trong việc xử lý yêu cầu (request) và phản hồi (response) tài nguyên. Các phương thức HTTP phổ biến gồm: GET, POST, PUT, DELETE, HEAD, OPTION, CONNECT và TRACE.

<video>
	<source src="https://ant.ncc.asia/wp-content/uploads/2024/02/GC0VxG7a8AAF9oz.mp4"/>
</video>

##### **GET:**

**Mục đích:** Lấy thông tin từ server.
**Đặc điểm:**
- Idempotent (thực hiện nhiều lần cho kết quả giống nhau)
- Dữ liệu được gửi qua URL parameters.
- Có giới hạn kích thước dữ liệu (~2KB).
- Không thể thay đổi dữ liệu trên server.
- Có thể được cache.
- Chỉ hỗ trợ các dữ liệu kiểu String.
- Có thể được bookmark (đánh dấu rồi xem lại sau) do được lưu trong lịch sử trình duyệt.
***Ví dụ:***
```http
GET https://example.com/api/users?id=123 HTTP/1.1
Host: example.com
```

##### **POST:**

**Mục đích:** Gửi dữ liệu đến server để tạo/cập nhật tài nguyên.
**Đặc điểm:**
- Không idempotent (thực hiện nhiều lần có thể tạo ra nhiều resource khác nhau).
- Dữ liệu được gửi trong body của request.
- Không gới hạn kích thức dữ liệu.
- Không được cache mặc định.
- Không thể Bookmark
- An toàn hơn GET vì dữ liệu không hiển thị trong URL.
***Ví dụ:***
```http
POST https://example.com/api/users HTTP/1.1
Host: example.com
Content-Type: application/json

{
	"name":"Gia Bao",
	"email":"GiaBao123456@gmail.com",
	"pwd":"HelloAll123",
}
```

##### **PUT:**

**Mục đích:** Cập nhật toàn bộ tài nguyên hoặc tạo mới nếu chưa tồn tại
**Đặc điểm:**
- Idempotent (thực hiện nhiều lần sẽ cho nhiều kết quả khác nhau).
- Thay thế hoàn toàn tài nguyên hiện có. 
- Cần gửi toàn bộ dữ liệu của tài nguyên
***Ví dụ:***
```http
PUT https://example.com/api/users?id=123 HTTP/1.1
Host: example.com
Content-Type: application/json

{
	"id":123,
	"name": "Pham Gia Bao",
	"email": "phamgiabao2202@gmail.com",
	"phone": "0123456789",
	"address": "HCM, Vietnam"
}
```
- Khi bạn muốn cập nhật toàn bộ thông tin cá nhân.

##### **PATCH:**

**Mục đích:** Cập nhật một phần tài nguyên
**Đặc điểm:**
- Không idempotent.
- Chỉ cập nhật các trường được chỉ định.
- Tiết kiệm băng thông hơn PUT.
***Ví dụ:***
```http
PATCH https://example.com/api/users/123 HTTP/1.1
Host: example.com
Content-Type: application/json

{
  "email": "nguyenvana_new@example.com"
}
```
- Khi bạn muốn cập nhật một phần thông tin mà không thay đổi các thông tin khác.

##### **DELETE:**

**Mục đích:** Xóa tài nguyên server
**Đặc điểm:**
- Idempotent (Xóa nhiều lần trong cùng một tài nguyên, Kết quả vẫn như nhau).
- Có thể trả về body, hoặc không
***Ví dụ:*** 
```http
DELETE https://example.com/api/users/123 HTTP/1.1 
Host: example.com
```

##### **HEAD:**

**Mục đích:** Giống như GET nhưng chỉ nhận header, không có body
**Đặc điểm:**
- Idempotent
- Tiết kiệm băng thông
- Thường dùng để kiểm tra tài nguyên mà không tải về.
***Ví dụ:***
```http
HEAD https://example.com/large-file.zip HTTP/1.1
Host: example.com
```
- Khi bạn muốn kiểm tra xem một file có tồn tại không, kiểm tra kích thước hoặc thời gian sửa đổi mà không cần trả về.
##### **CONNECT:**

**Mục đích:** Thiết lập một kết nối tunnel đến server.
**Đặc điểm:**
- Thường dùng cho HTTPS hoặc kết nối qua proxy.
- Chuyển đổi kết nối HTTP sang TCP.

***Ví dụ:***
```http
CONNECT secure.example.com:443 HTTP/1.1
Host: secure.example.com
```
- Khi bạn muốn truy cập website thông qua HTTPS, trình duyệt sẽ sử dụng CONNECT để thiết lập tunnel qua proxy.

##### **OPTIONS:**

**Mục đích:** Lấy thông tin về các phương thức HTTP được hỗ trợ
**Đặc điểm:**
- Idempotent.
- Thường dùng trong CORS (Cross-Origin Resource Sharing)
- Trả về header "Allow" liệt kê các method được phép.
***Ví dụ:***
```http
OPTIONS https://example.com/api/users HTTP/1.1
Host: example.com
```

Response:
```
HTTP/1.1 200 OK
Allow: GET, POST, PUT, PATCH, DELETE
Access-Control-Allow-Origin: *
Access-Control-Allow-Methods: GET, POST, PUT, PATCH, DELETE
```
##### **TRACE:**

**Mục đích:** Kiểm tra đường đi của request đến server.
**Đặc điểm:**
- Idempotent
- Dùng để debug.
- Server phản hồi chính xác nội dung đã nhận
- Thường bị vô hiệu hóa vì lý do bảo mật
***Ví dụ:***
```http
TRACE https://example.com/api HTTP/1.1
Host: example.com
Custom-Header: test-value
```

response:
```
HTTP/1.1 200 OK
Content-Type: message/http

TRACE https://example.com/api HTTP/1.1
Host: example.com
Custom-Header: test-value
```
- Khi bạn cần debug để xem request đã thực sự đến server như thế nào, đặc biệt khi đi qua nhiều proxy.

---
### **So sánh:**

### GET vs POST

|Tiêu chí|GET|POST|
|---|---|---|
|**Dữ liệu**|Trong URL|Trong body request|
|**Bảo mật**|Kém hơn (dữ liệu hiện trong URL)|Tốt hơn (dữ liệu ẩn trong body)|
|**Cache**|Có thể cache|Thường không được cache|
|**Idempotent**|Có|Không|
|**Bookmark**|Có thể|Không thể|
|**Kích thước**|Giới hạn (~2KB)|Không giới hạn|
|**Encoding**|application/x-www-form-urlencoded|application/x-www-form-urlencoded hoặc multipart/form-data|
|**Ứng dụng**|Lấy dữ liệu, tìm kiếm|Gửi form, upload file, cập nhật dữ liệu|
- GET được sử dụng để lấy dữ liệu và không làm thay đổi trạng thái server, thường dùng cho thao tác đọc.
- POST dùng để gửi dữ liệu và có thể thay đổi trạng thái server, thường dùng cho thao tác tạo mới.
- GET không an toàn vì dữ liệu nằm trong URL và có giới hạn độ dài, trong khi POST an toàn hơn với dữ liệu trong body và không giới hạn kích thước.

### PUT vs PATCH

|Tiêu chí|PUT|PATCH|
|---|---|---|
|**Cập nhật**|Toàn bộ tài nguyên|Một phần tài nguyên|
|**Idempotent**|Có|Không|
|**Yêu cầu dữ liệu**|Toàn bộ dữ liệu|Chỉ dữ liệu cần thay đổi|
|**Băng thông**|Sử dụng nhiều hơn|Tiết kiệm hơn|
|**Ngữ nghĩa**|Thay thế|Sửa đổi|
- PUT thay thế hoàn toàn một tài nguyên, yêu cầu gửi toàn bộ dữ liệu và là idempotent (nhiều request cho cùng kết quả)
- PATCH chỉ cập nhật một phần tài nguyên, gửi những trường cần thay đổi, tiết kiệm băng thông hơn và không nhất thiết là idempotent.
- Ví dụ, với PUT cần gửi tất cả thông tin người dùng kể cả các trường không thay đổi, còn PATCH chỉ cần gửi trường cần cập nhật như email mới.
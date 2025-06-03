## **[HTTP Request](obsidian://open?vault=CuuAmChanKinh&file=General_interview%2FRESTAPI%2FHTTP_Methobs):**

**HTTP request** có tất cả 9 loại method, 2 loại được sử dụng phổ biến là **GET** và **POST**.
- **GET:** được sử dụng để lấy thông tin từ server theo URI đã cung cấp.
- **HEAD:** giống với GET nhưng response trả về không có body, chỉ có header.
- **POST:** gửi thông tin tới server thông qua các biểu mẫu http.
- **PUT:** ghi đè tất cả thông tin của đối tượng với những gì được gửi lên.
- **PATCH:** ghi đè các thông tin được thay đổi của đối tượng.
- **DELETE:** xóa tài nguyên trên server.
- **CONNECT:** thiết lập một kết nối tới server theo URI
- **OPTIONS:** mô tả các tùy chọn giao tiếp cho resource.
- **TRACE:** Thực hiện một bài test loop - back đường dẫn đến resource.

---
## **Authorization:**

Hiện tại có 3 cơ chế Authorizie chính:
- HTTP Basic.
- JSON Web Token (JWT).
- OAuth2.
Tùy thuộc vào server mà chọn loại Authorization với mức độ phù hợp.

---
## **Quản lý version của API:**

- Khi thiết lập api cho app ios hay client side, chúng ta nên đặt version cho các api. Ví dụ như endpoint sau: **api/v1/users**
- Điều này sẽ giúp cho hệ thống sau khi nâng cấp lên version mới vẫn hõ trợ các api của version cũ, cũng như giúp cho việc sửa chữa và bảo trì dễ dàng hơn.

---
### **Các nguyên tắc (Constraints) của kiến trúc REST:**

###### 1. **Client-Server Architecture:** tách biệt client và server, cho phép các thành phần phát triển độc lập với nhau
###### 2. **Statelessness:** Mỗi yêu cầu từ client đến server phải chứa đầy đủ thông tin để hiểu và xử lý yêu cầu. Server không lưu trạng thái client giữa các request.
###### 3. **Cacheability:** Các phản hồi có thể được đánh dấu là cacheable hoặc non-cacheable. Nếu cacheable, client có thể tái sử dụng dữ liệu phản hồi cho các yêu cầu tương tự sau này
###### 4. **Uniform Interface:** Giao diện thống nhất giữa các thành phần, đơn giản hóa kiến trúc và tăng tính rõ ràng. Gồm 4 nguyên tắc:
- Resources identification in requests (*Định dạng nguồn tài nguyên trong các yêu cầu*).
- Resources mainpulation through representations (*Thao tác tài nguyên thông qua các biểu diễn*).
- Self-descriptive messages (**tin nhắn tự mô tả**)
- Hypermedia as the engin of application state (HATEOAS).
###### 5. **Layered System:** Kiến trúc phân lớp co phép các thành phần truy gian (intermediary) như proxy, load balancer, API gateway.
###### 6. **Code on Demand (optional):** Server có thể tạm thời mở rộng chức năng client bằng cách chuyển mã thực thi.

---
### **HATEOAS (Hypermedia as the engin of application state)**

**HATEOAS là một ràng buộc quan trọng của REST**, cho phép client tương tác với API thông qua các hypermedia (liên kết) được cung cấp trong response. Một API RESTful đầy đủ phải cung cấp hyperlinks trong mỗi response để client có thể biết các hành động có thể thực hiện tiếp theo.

*Ví dụ:*
```json
{
	"account":{
		"account_number": "12345",
		"balance":{
			"currency":"usd",
			"value": 100.00
		},
		"links":{
			"deposits": "/acounts/12345/deposits",
			"withdrawals": "/acounts/12345/withdrawals",
			"transfers":"/acounts/12345/transfers",
			"close": "/acounts/12345/close"
		}
	}
}
```

---
#### **Best Practices cho RESTful API Design:**

###### 1. **Resource naming:**  nên sử dụng danh từ số nhiều cho tên resource (e.g: `/users` thay cho `/user`)
###### 2. **Versioning:** Có nhiều cách version API (URL path, query parameter, header). 
###### 3. **Filtering, sorting, pagination:** Sử dung query parameters
###### 4. **Error handling:** response với status với message phù hợp
###### 5. **Documentation:** Sử dụng công cụ như Swagger/OpenAPI để tạo tài liệu
###### 6. **Rate limiting:** Bảo vệ API khỏi lạm dụng
###### 7. **API security:** Nhiều cấp độ bảo mật (transport, Access, Message level).

---
#### **Content Negitiation:**

**RESTful API** nên hỗ trợ nhiều định dạng dữ liệu, và client có thể chọn định dạng mong muốn thông qua HTTP headers:
- **Content-Type:** xác định định dạng của request body
- **Accept:** Xác định định dạng response mong muốn

---
### **Indempotency và Safety:**

##### - **Idempotent methods:** thực hiện nhiều lần vẫn cho kết quả giống nhau (POST, PUT, DELETE)
##### - **Safe methods:** không làm thay đổi trạng thái của server (GET, HEAD, OPTIONS).

---
#### **RESTful API thường áp dụng ở tầng nào trong project?**

RESTful API thường áp dụng ở phía back-end, cụ thể:
1. **Controller/ API Layer**: đây là tầng tiếp nhận **HTTP requests**, xử lý và trả về **responses**, trong kiến trúc MVC, đây là tầng **controller**:
2. **Service layer:** lớp trung gian giữa controllers và repositories, chứa business logic của ứng dụng.
3. **Persistence layer:** tương tác với cơ sở dữ liệu, thường được triển khai qua các respositories hoặc DAOs (Data Access Objects).

Trong kiến trúc N-tier phổ biến:
- **Presentation Layer:** Client-side (SPA, mobiles Apps).
- **API layer:** RESTful API controllers.
- **Business Logic Layer:** Services
- **Data Access Layer:** Repositories
- **Database Layer:** Cơ sở dữ liệu

---
### Ưu/nhược điểm của RESTful API

#### Ưu điểm:

1. **Dễ hiểu và sử dụng**: Dựa trên HTTP, một giao thức quen thuộc và phổ biến.
2. **Stateless**: Không duy trì trạng thái giữa các request, giúp khả năng mở rộng tốt hơn.
3. **Cacheability**: Có thể cache responses để tăng hiệu suất.
4. **Khả năng tương tác**: Có thể tương tác với các client khác nhau (web, mobile, desktop).
5. **Khả năng mở rộng**: Cho phép mở rộng và phát triển API theo thời gian mà không ảnh hưởng đến clients hiện có.
6. **Tách biệt client-server**: Client và server có thể phát triển độc lập.
7. **Công cụ phong phú**: Có nhiều framework, thư viện và công cụ hỗ trợ.
8. **Web standard**: Dựa trên các tiêu chuẩn web đã được chứng minh.

#### Nhược điểm:

1. **Over-fetching/Under-fetching**: Client có thể nhận được nhiều hoặc ít dữ liệu hơn mức cần thiết.
2. **Multiple round trips**: Có thể cần nhiều requests để lấy đủ dữ liệu phức tạp.
3. **Versioning challenges**: Quản lý versions của API có thể phức tạp.
4. **Lack of strong typing**: Không có kiểm tra kiểu dữ liệu tự động như một số công nghệ khác.
5. **Documentation overhead**: Cần duy trì tài liệu riêng (Swagger/OpenAPI).
6. **Performance overhead với nested resources**: Trong trường hợp dữ liệu lồng nhau phức tạp, hiệu suất có thể giảm.

---

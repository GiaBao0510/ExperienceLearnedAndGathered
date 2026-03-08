# Tổng Quan Về API

## **API là gì?**

API (Application Programming Interface - Giao diện Lập trình Ứng dụng) là cầu nối cho phép các chương trình, ứng dụng khác nhau giao tiếp và trao đổi dữ liệu với nhau.

**Ví dụ thực tế:**
- Khi bạn đăng nhập Facebook bằng tài khoản Google → đó là API của Google đang làm việc
- Khi xem bản đồ Google Maps trên trang web đặt đồ ăn → đó là API của Google Maps
- Khi thanh toán online qua VNPay, Momo → đó là API thanh toán

![](https://cole.edu.vn/wp-content/uploads/2024/05/1-2.png)

**Phạm vi sử dụng:**
- API nội bộ: giao tiếp giữa các module trong cùng một hệ thống
- API công khai: cho phép bên thứ ba sử dụng (ví dụ: Facebook API, Google Maps API)
- API có thể áp dụng cho: web, mobile app, game, IoT, AI,...

---
## **Phân loại API**

Có nhiều loại API khác nhau, mỗi loại phù hợp với từng mục đích sử dụng:

### 1. **RESTful API** (phổ biến nhất)

- Sử dụng giao thức HTTP
- Dễ học, dễ sử dụng
- Phù hợp cho web và mobile app

### 2. **SOAP API**

- Sử dụng XML
- Bảo mật cao, nghiêm ngặt
- Thường dùng trong ngân hàng, tài chính

### 3. **GraphQL API**

- Cho phép client yêu cầu chính xác dữ liệu cần thiết
- Tránh lấy thừa hoặc thiếu dữ liệu
- Phù hợp với ứng dụng phức tạp

### 4. **gRPC API**

- Hiệu năng cao, tốc độ nhanh
- Sử dụng Protocol Buffers
- Phù hợp cho microservices

### 5. **WebSocket API**

- Kết nối 2 chiều thời gian thực
- Dùng cho chat app, game online, trading

---

## **RESTful API** (REST API)

REST (Representational State Transfer) là một kiến trúc thiết kế API phổ biến nhất hiện nay.

![](https://images.ctfassets.net/vwq10xzbe6iz/5sBH4Agl614xM7exeLsTo7/9e84dce01735f155911e611c42c9793f/rest-api.png)
### **Cách hoạt động:**

```
[Client/Trình duyệt] 
    ↓ (Request + Method)
[REST API]
    ↓ (Truy vấn)
[Database/Cơ sở dữ liệu]
    ↓ (Trả dữ liệu)
[REST API]
    ↓ (Response: JSON/XML)
[Client/Trình duyệt]
```

### **Ví dụ thực tế:**

Bạn muốn xem danh sách sản phẩm trên website:
1. Trình duyệt gửi request: `GET /api/products`
2. REST API nhận request, truy vấn database
3. Database trả về danh sách sản phẩm
4. REST API gửi response dạng JSON về trình duyệt
5. Trình duyệt hiển thị danh sách sản phẩm

### **Các HTTP Methods thường dùng:**

| Method      | Ý nghĩa                       | Ví dụ                    |
| ----------- | ----------------------------- | ------------------------ |
| **GET**     | Lấy dữ liệu                   | Xem danh sách sản phẩm   |
| **POST**    | Tạo mới dữ liệu               | Đăng ký tài khoản        |
| **PUT**     | Cập nhật toàn bộ              | Sửa thông tin profile    |
| **PATCH**   | Cập nhật một phần             | Đổi mật khẩu             |
| **DELETE**  | Xóa dữ liệu                   | Xóa bài viết             |
| **HEAD**    | Lấy metadata (không lấy body) | Kiểm tra file có tồn tại |
| **OPTIONS** | Kiểm tra methods được hỗ trợ  | CORS preflight           |
![](https://images.viblo.asia/f94a70b4-94f3-4640-a903-ad1d4eed4cbc.png)
### **Các định dạng Response:**

- **JSON** (phổ biến nhất): dễ đọc, nhẹ, dễ xử lý
- **XML**: cấu trúc rõ ràng nhưng nặng hơn

---
## **Giao thức HTTP**

HTTP (HyperText Transfer Protocol) là giao thức truyền tải dữ liệu trên web - giống như "ngôn ngữ" để client và server giao tiếp với nhau.

![](https://www.corero.com/wp-content/uploads/2024/07/HTTP-Process-768x367.png)
### **Cơ chế hoạt động:**

1. **Client gửi Request** đến Server, bao gồm:
    - URL (địa chỉ tài nguyên)
    - Method (GET, POST,...)
    - Headers (thông tin bổ sung)
    - Body (dữ liệu gửi đi, nếu có)
    
2. **Server xử lý Request** và gửi **Response** về, bao gồm:
    - Status Code (mã trạng thái)
    - Headers (thông tin về response)
    - Body (dữ liệu trả về)

### **Ví dụ HTTP Request:**

```
GET /api/users/123 HTTP/1.1
Host: example.com
Authorization: Bearer token123
Content-Type: application/json
```

### **Ví dụ HTTP Response:**

```
HTTP/1.1 200 OK
Content-Type: application/json
{
  "id": 123,
  "name": "Nguyễn Văn A",
  "email": "a@example.com"
}
```

---

## **HTTP Status Code**

HTTP Status Code là những con số máy chủ (server) trả về để cho biết kết quả xử lý request - thành công, lỗi, hoặc cần làm gì tiếp theo.

### **Phân loại theo nhóm:**

|Nhóm|Ý nghĩa|
|---|---|
|**1xx**|Thông tin (Informational)|
|**2xx**|Thành công (Success)|
|**3xx**|Chuyển hướng (Redirection)|
|**4xx**|Lỗi Client (Client Error)|
|**5xx**|Lỗi Server (Server Error)|

### **Các mã trạng thái phổ biến:**

| Mã      | Tên                   | Ý nghĩa                                  | Ví dụ                                |
| ------- | --------------------- | ---------------------------------------- | ------------------------------------ |
| **200** | OK                    | Yêu cầu thành công                       | Lấy dữ liệu thành công               |
| **201** | Created               | Tạo mới thành công                       | Đăng ký tài khoản thành công         |
| **204** | No Content            | Thành công nhưng không có dữ liệu trả về | Xóa dữ liệu thành công               |
| **400** | Bad Request           | Request sai cú pháp hoặc thiếu dữ liệu   | Gửi form thiếu trường bắt buộc       |
| **401** | Unauthorized          | Chưa đăng nhập hoặc token hết hạn        | Truy cập trang yêu cầu đăng nhập     |
| **403** | Forbidden             | Không có quyền truy cập                  | User thường truy cập trang Admin     |
| **404** | Not Found             | Không tìm thấy tài nguyên                | Truy cập URL không tồn tại           |
| **405** | Method Not Allowed    | Method không được hỗ trợ                 | Gửi POST đến endpoint chỉ hỗ trợ GET |
| **500** | Internal Server Error | Lỗi server (lỗi code, database,...)      | Server bị crash, lỗi logic           |
| **503** | Service Unavailable   | Server tạm thời không khả dụng           | Server đang bảo trì                  |

### **Cách nhớ nhanh:**

![](https://images.viblo.asia/full/80981e61-d77a-4619-87af-6058320c0790.png)

---

## **JSON là gì?**

JSON (JavaScript Object Notation) là định dạng dữ liệu nhẹ, dễ đọc, được sử dụng phổ biến để trao đổi dữ liệu giữa client và server.

### **Đặc điểm:**

- Sử dụng cặp `key: value`
- Dễ đọc cho con người
- Dễ parse (phân tích) cho máy tính
- Nhẹ hơn XML

### **Cấu trúc JSON:**

#### **1. Object (Đối tượng):**

```json
{
  "id": 1,
  "name": "Nguyễn Văn A",
  "age": 20,
  "isStudent": true,
  "address": {
    "city": "Hà Nội",
    "district": "Cầu Giấy"
  }
}
```

#### **2. Array (Mảng):**

```json
{
  "students": [
    {
      "id": 1,
      "name": "Nguyễn Văn A"
    },
    {
      "id": 2,
      "name": "Trần Thị B"
    }
  ]
}
```

### **Các kiểu dữ liệu trong JSON:**

- **String** (chuỗi): `"Xin chào"`
- **Number** (số): `123`, `45.67`
- **Boolean** (logic): `true`, `false`
- **Null** (rỗng): `null`
- **Object** (đối tượng): `{...}`
- **Array** (mảng): `[...]`

### **Ví dụ Response API thực tế:**

```json
{
  "status": "success",
  "message": "Lấy dữ liệu thành công",
  "data": {
    "products": [
      {
        "id": 1,
        "name": "Laptop Dell",
        "price": 15000000,
        "inStock": true
      },
      {
        "id": 2,
        "name": "iPhone 15",
        "price": 25000000,
        "inStock": false
      }
    ],
    "total": 2
  }
}
```

---

## **Tóm tắt**

1. **API** là cầu nối cho các ứng dụng giao tiếp với nhau
2. **RESTful API** là loại API phổ biến nhất, sử dụng HTTP protocol
3. **HTTP Methods** (GET, POST, PUT, DELETE,...) xác định hành động muốn thực hiện
4. **HTTP Status Code** cho biết kết quả xử lý request (2xx = thành công, 4xx = lỗi client, 5xx = lỗi server)
5. **JSON** là định dạng dữ liệu phổ biến để trao đổi thông tin qua API
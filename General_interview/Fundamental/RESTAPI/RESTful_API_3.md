### **Các nguyên tắc của RESTful API**

Khi thiết kế RESTfulAPI, cần phải tuân thủ một số nguyên tắc cơ bản cần tuân thủ đển đảm bảo API dễ hiểu:
##### 1. **Sử dụng các phương thức HTTP phù hợp:**
- **GET:** Lấy dữ liệu
- **POST:** Tạo mới tài nguyên
- **PUT:** Cập nhật toàn bộ tài nguyên
- **PATCH:** Cập nhật một phần tài nguyên
- **DELETE:** Xóa tài nguyên
##### 2. **Sử dụng danh từ cho các endpoint (không sử dụng động từ):**
- API nên đại diện cho các tài nguyên (resource) bằng danh từ, *ví dụ:* `/users`, `products`, thay vì `/getUsers` hay `createProduct`.
- Tài nguyên nên được tổ chức theo dạng phân cấp nếu cần: `/users/{id}/orders`.

##### 3. **Sử dụng định dạng URL nhất quán:**
- Sử dụng danh từ số nhiều cho tài nguyên (*ví dụ*: `/users` hay `/user`)
- Tránh sử dụng các ký tự đặc biệt hoặc các định dạng không chuẩn trong URL.
- Sử dụng dấu gạch ngang `-` thay vì dấu gạch dưới `_` trong URL để dễ đọc

##### 4. **Hộ trợ trạng thái không trạng thái:**
- Mỗi yêu cầu (requset) từ client phải chứa đầy đủ thông tin để server xử lý, không dựa vào trạng thái trước đó.
- Sử dụng token (như JWT) để quản lý phiên thay vì lưu trạng thái trên server.

##### 5. **Sử lý mã trạng thái HTTP đúng cách**
| Mã  | Ý nghĩa                            |
| --- | ---------------------------------- |
| **200** | OK – Thành công                    |
| **201** | Created – Tạo mới thành công       |
| **204** | No Content – Xoá thành công        |
| **400** | Bad Request – Request sai          |
| **401** | Unauthorized – Chưa xác thực       |
| **403** | Forbidden – Không được phép        |
| **404** | Not Found – Không tìm thấy         |
| **500** | Internal Server Error – Lỗi server |

##### 6. **HATEOAS (Hypermedia as the Engine of Application State):**
- API nên cung cấp các liên kết (links) trong phản hồi để giúp client điều hướng hoặc thực hiện các hành động tiếp theo, *ví dụ*:
```json
{
  "id": 1,
  "name": "John Doe",
  "links": {
    "self": "/users/1",
    "orders": "/users/1/orders"
  }
}
```

##### 7. **Định dạng dữ liệu nhất quán (thường là json):**
- Tên trường (field names) nên dùng kiểu **camlCase** hoặc **snake_case** và nhất quán toàn bộ API.
- Tránh lồng dữ liệu quá sâu.

##### 8. **Hỗ trợ phân trạng, lọc và sắp xếp:**
- Sử dụng query paramethers để hỗ trợ phân trang (`?page=1&size=10`), lọc (`?status=active`) và sắp xếp (`sort=name,asc`)
- Ví dụ: `/users?page=2&size=20&sort=name,desc`

##### 9. **Versioning (quản lý phiên bản):**
- Sử dụng versioning để quản lý thay đổi API, ví dụ: `/v1/users` hoặc thông quan header (`Accept: application/vnd.example.v1+json`)
- tránh phá vỡ các client hiện tại khi cập nhật API.

##### 10. **Bảo mật:**
- Sử dụng **HTTPS** để mã hóa dữ liệu truyền tải.
- Xác thực và phân quyền bằng **OAuth**, **API keys**, hoặc JWT.
- Giới hạn tỷ lệ yêu cầu (rate limiting) để ngăn chặn lạm dụng.

##### 11. **Xử lý lỗi rõ ràng**
- Trả về thông báo lỗi có cấu trúc, ví dụ:
```json
{
  "error": {
    "code": 400,
    "message": "Invalid input data",
    "details": ["Field 'name' is required"]
  }
}
```

##### 12. **Tài liệu API**
- Cung cấp tài liệu rõ ràng (ví dụ: sử dụng OpenAPI/Swagger) để mô tả các endpoint, phương thức, tham số, và phản hồi.
- Đảm bảo tài liệu luôn được cập nhật.

---
### **API Security:**
### **API Testing:**
### **API Design Patterns:**
### **API Documentation:**
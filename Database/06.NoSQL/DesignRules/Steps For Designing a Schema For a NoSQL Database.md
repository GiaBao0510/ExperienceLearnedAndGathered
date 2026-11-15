## **Các bước thiết kế schema cho cơ sở dữ liệu NoSQL**

Mặc dù một số có thể cho rằng NoSQL không cần thiết kế schema do tính linh hoạt và không ràng buộc của nó nhưng thực tế thì thiết kế schema vẫn mang lại nhiều lợi ích trong việc tổ chức dữ liệu. Schema giúp tối ưu hóa hiệu suất truy vấn, đảm bảo tính nhất quán và tổ chức logic của dữ liệu cũng như hỗ trợ quá trình phát triển và bảo trì ứng dụng. Dưới đây là một số bước và phương pháp phổ biến để thiết kế schema trong NoSQL:

### **1. Thiết kế Containers**

- Trong ngữ cảnh này, “container” đề cập đến các đơn vị lưu trữ dữ liệu trong cơ sở dữ liệu NoSQL, tương tự như bảng trong RDBMS.
- Xác định các Containers dựa trên Query Patterns: Xây dựng các container dựa trên mẫu truy vấn để mỗi container phục vụ một chức năng duy nhất.
- Cân bằng thiết kế Container: Đảm bảo container tương ứng với một module/class code để thực thi một chức năng duy nhất, giúp trong việc phát triển và bảo trì hệ thống. Điều quan trọng là phải cân nhắc giữa việc thiết kế một container cho mỗi truy vấn so với thiết kế một container để thực hiện nhiều truy vấn. Do đó, việc sử dụng thuật toán chuẩn hóa ngược (denormalization) đóng vai trò quan trọng trong thiết kế NoSQL.

![](https://itviec.com/blog/wp-content/uploads/2024/04/nosql-container-vippro.jpg)
Cách Container chứa cơ sở dữ liệu NoSQL. (Nguồn: @henriquesd.medium.com)

### **2. Điều chỉnh Dữ liệu**

- Sử dụng kỹ thuật nhúng (Embedding) hoặc tham chiếu (Referencing) để giảm bớt sự chuẩn hoá dữ liệu giữa các containers.
- Xác định mối quan hệ , phép nối cần thực hiện là quan hệ 1-1 (One-to-One),  quan hệ 1-nhiều (One-to-Many) hay quan hệ nhiều-nhiều (Many-to-Many) giữa các đối tượng (entities) và áp dụng kỹ thuật giảm chuẩn hóa tương ứng.
- Lựa chọn giữa Embedding và Referencing dựa trên số lượng và tần suất truy cập của các thuộc tính.
![](https://itviec.com/blog/wp-content/uploads/2024/04/nosql-tham-chieu-vippro.jpg)
*Dùng tham chiếu (Referencing) để giảm bớt sự chuẩn hóa dữ liệu. (Nguồn: @datensen)*

### **3. Thiết kế Primary Keys**

- Lựa chọn các thuộc tính liên quan đến doanh nghiệp làm Primary Keys để tối ưu hóa việc truy cập dữ liệu.
- Phân loại các mẫu truy vấn và thiết kế Primary Keys để đáp ứng yêu cầu của các truy vấn quan trọng và chính xác.
- Thiết kế Primary Keys để hỗ trợ việc tham gia dữ liệu trên ứng dụng mà không cần sử dụng các chỉ số phức tạp.

### **4. Thiết kế Indexes**

- Tạo các Indexes dựa trên các mẫu truy vấn để cải thiện hiệu suất của các truy vấn không quan trọng hoặc dùng cho mục đích khai thác dữ liệu.
- Tránh việc tạo quá nhiều Indexes cho mỗi truy vấn và tối ưu hóa các Indexes để giảm bớt tác động đến hiệu suất hệ thống.
- Thiết kế các loại Indexes phù hợp với nhu cầu cụ thể của ứng dụng như B-Tree, Array, Geo-spatial và Full Text Search.
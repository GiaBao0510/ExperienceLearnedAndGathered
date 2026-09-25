#### Hàm **Main**: 
- Sẽ chạy tất cả các câu lệnh xuất hiện bên trong nó. 
- Cũng là điểm bắt đầu để chạy chương trình. 
- Phương thức **Main** là phương thức tĩnh không thuộc về một đối tượng cụ thể nào của một lớp nào đó.
- Phương thức **Main** không có kiểu trả về. Vì có kiểu trả về là _void_.
- Hàm **Main** có đối số là mảng các chuỗi. Và mảng các chuỗi được truyền và khi thực hiện các lệnh.
- Hàm **Main** chỉ khai báo bên trong một lớp nào đó

#### **Command line on terminal:**
- _dotnet build_: Câu lệnh này dùng đề build một ứng dụng c# là file có đuôi là .dll.
- _dotnet [file name]_: câu lệnh này dùng để chạy tệp tin có đuôi là .dll.
- _dotnet run_: câu lệnh này dùng để chạy chương trình khi và chỉ khi ở tại thư mục của chương trình cần chạy.
- dotnet publish -c [Tên thư mục cần lưu]: dùng để đóng gói dự án. Sau khi hoàn thành.

### Ghi chú trong tệp chương trình: _Khi thực hiện chương trình sẽ bỏ qua các dòng ghi chú._
- //: Đây là đoạn ghi chú được sử dụng trên một dòng đơn.
- /*  */: đây là đoạn ghi chú được sử dụng trên nhiều dòng. 
- ///: dùng để miêu tả chi tiết ghi chú về hàm, phương thức, lớp,... Theo chuẩn C#.


#### **Bài Học:** 
- Câu lệnh: **Console.WriteLine()** dùng để hiển thị thông tin.
- Phím tắt biên dịch:  **ctrl + F5**.
- Khai báo phương thức bên trong lớp.
- **Phương thức tĩnh** có thể được gị từ bất kỳ đâu. Vì nó không thuộc về một loại đối tượng cụ thể nào.
- Nếu mà một file chương trình mà tồn tại **hai hoặc nhiều lớp**, mà _mỗi lớp điều có hàm Main(Tức 2 điểm mồi)_ thì không thể chạy chương trình đó được.
- Trên toàn bộ dự án chỉ có duy nhất hàm Main thôi.
- **namespace:** dùng để gôm các lớp có cùng đặc điểm trong một khu vực chung. Mục đích là dễ quản lý và tránh bị xung đột.
- **_using System:_** ý nghĩa của đoạn code này dùng để sử dụng namespace _using system_. Đây là namespace hệ thống của dotnet.
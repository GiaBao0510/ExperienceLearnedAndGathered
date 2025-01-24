## **1. File(tệp tin):**
- Trong hệ thống máy tính thì dữ liệu được lưu trữ trên các đĩa, các đĩa thì tổ chức thành các thư mục, các file.
- Lớp **DriveInfo** lớp này cung cấp thông tin về ổ đĩa. Các phương thức của lớp này thì vào tệp tin program.sc để xem chi tiết.
- Lớp **File** giúp ta làm việc với file như: Đọc file, sửa file, xóa file, di chuyển file, đổi tên file,...

## **2. Folder(thư mục):**
- Để làm việc với thư mục thì dùng lớp **Directory**.

## **3. Path(đường dẫn):**
- Để làm việc với đường dẫn thì dùng lớp **Path**, cho phép hỗ trợ quản lý các đường dẫn đến file thư mục. Nhất là hỗ trợ cross-plattom thì lớp tĩnh System.IO.path chứa các phương thức tĩnh thực thi với phương thức đó.

## **4. FileStream():**
- Dùng để lưu thông tin được truyền từ mạng vào trong file.
- Lớp **File Stream** dùng để tạo ra các đối tượng dùng để đọc và ghi dữ liệu ra file. Do **Stream** không là tài nguyên quản lý bới GC. _Nên cấm dư nó vào cấu trúc string để tự động gọi giải phóng tài nguyên_.
- Khi làm việc với **Stream** thì làm việc với mảng byte dữ liệu.
- Để chuyển đổi dữ liệu sang mảng Byte thì quan tâm đến các phương thức như: **GetBytes**, **ToInt32**, **ToDouble**,... trong lớp **BitConverter**
- Để chuyển đổi dữ liệu kiểu chuỗi sang mảng **Byte** thì dùng phương thức **UTF8 từ lớp Encoding**.
## **Introduction**

Nếu hệ thống phòng thủ của tổ chức không thể ngăn chặn thành công một cuộc tấn công mạng thì ưu tiên tiếp theo của tổ chức đó là phát hiện cuộc tấn công mạng. Hy vọng rằng việc phát hiện xảy ra trong khi cuộc tấn công đang diễn ra hoặc thậm chí tốt hơn là trước khi vi phạm xảy ra.

Trong mô-đun này, chúng ta sẽ xem xét một số phương pháp mà các tổ chức sử dụng để phát hiện các cuộc tấn công.

![](https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcRDi8jdxiI9-omTXOLDaJ8jYJKwcaTzqOgR1Z4VlZuLzsC4_5Yy3PUcbks&s=10)

---

## **Antimalware software**

Phương pháp tiêu chuẩn để tấn công hệ thống là phần mềm độc hại và biện pháp tiêu chuẩn để phát hiện phần mềm độc hại là phần mềm chống phần mềm độc hại. **Antimalware software** hay phần mềm chống vi-rút là phần mềm chuyên dụng có chức năng phát hiện, cách ly và thậm chí tiêu diệt phần mềm độc hại trên máy tính hoặc mạng. Một số chương trình chống phần mềm độc hại nổi tiếng bao gồm Malwarebytes, McAfee Antivirus và Windows Defender Antivirus. Bạn có thể cài đặt phần mềm này cục bộ trên một thiết bị hoặc chạy và quản lý phần mềm trên máy chủ tập trung.

Phần mềm chống phần mềm độc hại phát hiện phần mềm độc hại bằng cách quét tất cả các tệp thiết bị để tìm chữ ký. **malware signature** là một mẫu thuộc tính tương ứng với phần mềm độc hại đã biết. Sau khi xác định chữ ký trong tệp, phần mềm sẽ xóa tệp, cách ly tệp hoặc cảnh báo cho bạn rằng tệp có thể bị nhiễm.

---

## **Logging**

Phương pháp tiêu chuẩn để tấn công hệ thống là phần mềm độc hại và biện pháp tiêu chuẩn để phát hiện phần mềm độc hại là phần mềm chống phần mềm độc hại. Phần mềm chống phần mềm độc hại hay phần mềm chống vi-rút là phần mềm chuyên dụng có chức năng phát hiện, cách ly và thậm chí tiêu diệt phần mềm độc hại trên máy tính hoặc mạng. Một số chương trình chống phần mềm độc hại nổi tiếng bao gồm Malwarebytes, McAfee Antivirus và Windows Defender Antivirus. Bạn có thể cài đặt phần mềm này cục bộ trên một thiết bị hoặc chạy và quản lý phần mềm trên máy chủ tập trung.

Phần mềm chống phần mềm độc hại phát hiện phần mềm độc hại bằng cách quét tất cả các tệp thiết bị để tìm chữ ký. Chữ ký phần mềm độc hại là một mẫu thuộc tính tương ứng với phần mềm độc hại đã biết. Sau khi xác định chữ ký trong tệp, phần mềm sẽ xóa tệp, cách ly tệp hoặc cảnh báo cho bạn rằng tệp có thể bị nhiễm.



> **<u>Ví dụ</u>**
> Mục nhật ký sau đây hiển thị định dạng nhật ký mà máy chủ web Apache sử dụng:
> 
> 9.12.156.2 - bob [11/01/2020:14:16:34 -0700] “GET /index.html HTTP/1.0” 200 4066
> 
> Lưu ý rằng mục nhật ký này mô tả một người dùng có tên bob đang truy cập một trang web cụ thể với trạng thái và thời gian được ghi chú.



---



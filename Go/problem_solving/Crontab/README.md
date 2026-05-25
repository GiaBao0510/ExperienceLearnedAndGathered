## **Cronjob/ Crontab là gì?**
Nó là một hành động được thực hiện định kỳ trên máy tính (ví dụ công việc A cứ cách một tiếng sẽ thực hiện 1 lần, công việc B cứ cách một ngày sẽ thực hiện 1 lần,...). Vậy trường hợp nào sẽ áp dụng điều này.

---
## **Kịch bản nào sẽ sử dụng đến Crontab/Cronjob?**

Ví dụ:
- Khi chúng ta thanh toán đơn hàng, như giả xử nó bị lỗi hoặc chúng ta để đơn hàng đó quá lâu (ví dụ đơn hàng này chỉ có 60p để thanh toán) mà mình không thực hiện thì nó sẽ đưa vào trong cronjob để hủy
- Trên Youtube người dùng thiết lập cho video sẽ đăng vào khung 9 giờ tối, thì đển 9 giờ tối nó sẽ tự động công khai video của người dùng đó trên nền tảng
- Người dùng đăng ký bằng email để muốn xem webinar trên amazon vào lúc 9 giờ tối (tại ngày nào đó cụ thể), thì trước khi đến khoảng thời gian đó thì nó sẽ tự động gửi thông báo email cho người dùng đã đăng ký email đó. (Tức là có một tác vụ nó sẽ scan tất cả trong Database để lấy ra những email đã đăng ký trước đó, để nó có thể gửi thông báo tước 10 phút. Số lượng email này có thể vài chục, vài trăm hoặc vài triệu record )
- ...

---
## **Phương pháp nào triển khai Crontab/Cronjob hiệu quả**
 Công cụ xử lý:
 - redis
 - rabitMQ => message delay
 - cronjob/crontab 

---
## Sự bất cẩn của BackEnd

Ví dụ có trường hợp: Giả sử có tác vụ một phút gọi lên (get list database) 100.000 records email để thông qua amazon service rồi send email qua SMTP, nhìn thì thấy quy trình này diễn ra khá là chơn chu. Nhưng nếu một ngày SMTP có vấn đề (như nghẽn mạng), thì 1 phút chưa lấy đủ 100.000 records email. 
Thì thời điểm tiếp theo nó tác động lại lần nữa thì nó lấy 100.000 records tiếp theo. Nhưng bên chỗ 100.000 records trước đó chưa xử lý xong thì nó lại lấy luôn dữ liệu trước đó. Thì lúc đó có thể là 1 email bị gửi đến 2 lần mà chung 1 nội dung (Đây là tác vụ thực tế kinh điển bị ban, bị push ,bị duplicate, spam,...)
Thì đây là lý do phần Cronjob thực hiện chưa đúng.

---
## **Giải quyết vấn đề trên**

Để giải quyết vấn đề trên thì ta có thể sử dụng package của go là thư viện [này]([robfig/cron: a cron library for go](https://github.com/robfig/cron)).











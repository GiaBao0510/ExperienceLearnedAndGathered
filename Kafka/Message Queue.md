## **1.Message Queue là gì?**

Message Queue là một cơ chế trong lập trình và kiến trúc phần mềm, được sử dụng để truyền thông tin (thông điệp) giữa các thành phần của hệ thống mà không cần chúng tương tác với nhau.

![](https://statics.cdn.200lab.io/2024/01/message-queue-small.png?width=1200)
## **2.Các thành phần chính Message Queue**

![](https://hookdeck.com/_astro/producer-sending-messages-to-a-consumer-through-a-message-queue.CmVeoKrF_Frhzd.webp)
+ **Producer:** đây là bộ phận tạo ra thông tin để tương tác với bộ phận khác. Thông tin này sẽ được truyền vào Message queue
+ **Consumer:** bộ phận/ thành phần trong hệ thống nhận thông tin và xử lý thông tin từ Producer thông qua Message queue
+ **Message:** thông tin (thông diệp) thường là ở dạng TEXT hoặc JSON, đôi khi có thể ở dạng Binary do Producer tạo ra
+ **Broker:** Nơi đây sẽ xử lý Message và quản lý Message Queue để đảm bảo Producer và Consumer truyền thông tin được cho nhau. Broker giúp định tuyến thông tin, quản lý tình trạng của hàng đợi, và đảm bảo rằng thông tin được chuyển giao đúng cách
+ **Channel:** Là cơ chế truyền thông tin giữa Producer và Consumer thông qua Message Queue. Channel đóng vai trò như một cầu nối để truyền thông điệp qua lại giữa các bên

## **3.Các thức hoạt động**

![](https://statics.cdn.200lab.io/2024/01/sqs_seo_queue.1dc710b63346bef869ee34b8a9a76abc014fbfc9.png?width=1200)
##### **Bước 1:** producer tạp ra thông điệp cần truyền đi. Thông tin này có thể là dữ liệu hoặc các thông tin bổ sung. Thông tin này sẽ được truyền vào Message queue thông qua channel và được lưu trữ tạm thời tại đây

***Ví dụ:*** nếu chúng ta phải cập nhật danh mục, tên, mô tả về một sản phẩm
Thông điệp mà Producer tạo và đưa vào Message queue sẽ trông giống như sau:
```json
{
	"product_id": 12345,
	"category": 567,
	"productname": "Nike Zoom ZX 2024",
	"description": "Made in Vietnam"
}
```
##### **Bước 2:** Consumer sẽ lấy thông điệp của Producer thông Message queue. Thông tin thường được lấy theo cơ chế FIFO (first in - first out), tuy nhiên vẫn có thể can thiệp vào cơ chế này bằng cách định ra các ưu tiên.
##### **Bước 3:** Sau khi Consumer lấy được thông tin sẽ tiếp tục xử lý và thực hiện các hành động tuỳ thuộc vào yêu cầu của hệ thống.

Quá trình này tạo ra một mô hình truyền và nhận thông tin linh hoạt và không đồng bộ giữa các thành phần của hệ thống. Producer và consumer không cần biết về sự tồn tại của nhau; thay vào đó, chúng tương tác thông qua hàng đợi được quản lý bởi Broker. Điều này giúp tăng tính mở rộng và giảm sự phụ thuộc giữa các thành phần trong hệ thống

## **4.Ưu và Nhược điểm**

#### ***Ưu điểm:***

![](https://www.cloudamqp.com/img/blog/illu-cloudamqp-architechture.jpg)

- **Bất đồng bộ:** Message Queue hỗ trợ truyền thông diệp giữa các thành phần mà không đòi hỏi chúng phải chờ đợi nhau. Điều này cải thiện hiệu suất và tăng tính mở rộng của hệ thống
- **Tính độc lập:** Producer và Consumer không cần biết về sự tồn tại của nhau. Sự phân tách này giúp giảm sự phụ thuộc giữa các thành phần và tăng khả năng mở rộng của hệ thống
- **Xử lý lưu lượng cao:** Message Queue có thể xử lý lượng lớn thông điệp lớn và đồng thời từ nhiều nguồn mà không gây ảnh hưởng đến hiệu suất của hệ thống
- **Đảm bảo giao tiếp tin cậy:** Hệ thống Message Queue đảm bảo rằng việc trao đổi thông tin giữa Producer và Consumer được chính xác và xử lý đúng cách
- **Giảm lỗi chồng chéo:** Message Queue giúp giảm lỗi chống chéo bằng cách loại bỏ giữa các thành phần, giảm khả năng lỗi do sự phụ thuộc và giao tiếp trực tiếp
- **Khả năng phục hồi:** Do các thành phần hoạt động độc lập với nhau nên khi một thành phần gặp sự cố thì thành phần kia vẫn có thể hoạt động bình thường. Việc bảo trì, sửa chữa hrrj thống cũng không quá phức tạp 
#### ***Nhược điểm:***

- ***Phức tạp hoá hệ thống:*** Sử dụng Message Queue có thể làm phức tạp hoá hệ thống và tốn kém. Đối với các hệ thống nhỏ, đôi khi triển khai Message Queue là không cần thiết
- ***Độ trễ:*** Việc trao đổi thông tin bất đồng bộ giữa các thành phần sẽ có một số độ trễ nhất định
- ***Chi phí xử lý:*** Message Queue sẽ tăng tải của hệ thống nếu phải xử lý lượng lớn thông tin
- ***Quản lý và theo dõi:*** Khi hệ thống có nhiều hàng đợi, hoặc có nhiều Producer/ consumer thì việc quản lý theo dõi hoạt động động của Message queue sẽ gặp nhiều khó khăn
- ***Khó xử lý đồng bộ:*** Khi hệ thống cần xử lý đồng bộ giữa các service thì Message queue không phải là lựa chọn hàng đầu mà phải chọn các cơ chế khác phù hợp hơn Rest hoặc rGPC 

## **5.Ứng dụng trong thực tế**
Trong thực tế, Message queue giải quyết được rất nhiều vấn đề quan trọng trong hệ thống như:

- Message được lưu giữ trong hàng đợi (queue) nên khi các thành phần xử lý gặp lỗi hoặc bộ phận trong hệ thống gặp sự cố thì không mất dữ liệu. Khi hệ thống được phục hồi thì có thể tiếp tục lấy message trong queue để xử lý tiếp.
- Khi số lượng Message quá lớn, thì cơ chế xử lý bất đồng bộ sẽ phát huy hiệu quả. Các Message sẽ được xử lý dần dần cho tới khi hoàn tất mà không sợ bị thất thoát thông tin hoặc gây quá tải cho hệ thống.
- Các thành phần hoạt động tách biệt nên dễ dàng mở rộng hệ thống. Trong thực tế có những thời điểm lượng message tăng cao thì có thể tăng lượng consumer lên để xử lý.

Do những đặc điểm ưu việt trên mà Message queue được sử dụng nhiều trong các hệ thống  phân tán và phức tạp. Dưới đây là một số trường hợp thực tế:

- **Xử lý đơn hàng và Thanh Toán Trực Tuyến:**

Trong môi trường thương mại điện tử, khi một đơn hàng được đặt thành công, thông tin  đơn hàng có thể được xếp vào hàng đợi để  hệ thống xác nhận đơn hàng xử lý, xử lý thanh toán và cập nhật trạng thái đơn hàng.

- **Xử Lý Sự Kiện Real-time**

Trong ứng dụng cần xử lý sự kiện real-time như theo dõi và phân tích dữ liệu sensor, Message Queue có thể đóng vai trò là cầu nối để truyền thông điệp giữa các thành phần xử lý sự kiện.

- **Chia Sẻ Dữ Liệu Giữa Ứng Dụng và Dịch Vụ**

Message Queue được sử dụng để chia sẻ dữ liệu giữa các ứng dụng và dịch vụ khác nhau. Điều này giúp đơn giản hóa tích hợp và giảm sự phụ thuộc trực tiếp giữa các thành phần.

***Ví dụ:*** có một trang web cho phép người dùng tải video từ hệ thống thì nó sẽ có các thành phần tác vụ sau:

- ***API service:*** là 1 Producer. nhận thông tin (URL Video) từ phía người dùng và đưa thông tin này vào message queue
- ***Processing Service:*** Vừa là consumer vừa là producer. Service này đọc URL Video từ message queue, bắt đầu tải file Video về encode lại, lưu vào Server. Sau khi encode xong, nó đưa URL của file đã encode vào message queue.
- ***Upload Service:*** Khi nhận được message từ processing server, nó sẽ upload video đó lên Amazon 3
![](https://www.cloudamqp.com/img/blog/rabbitmq-beginners-updated.png)
### ***Một số Message queue được dùng hiện nay:***

- [Kafka](https://lcdung.top/apache-kafka-la-gi/)
- [Pulsar](https://lcdung.top/apache-pulsar-incubator/)
- RabitMQ
- ActiveMQ
- SQS
- ZeroMQ
- MSMQ
- IronMQ
- Kinesis
- RocketMQ

## **6.Tài liệu tham khảo:**
1.https://toidicodedao.com/2019/10/08/message-queue-la-gi-ung-dung-microservice/
2.https://200lab.io/blog/message-queue-la-gi
3.https://topdev.vn/blog/message-queue-la-gi/
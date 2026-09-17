
![](https://www.imperva.com/learn/wp-content/uploads/sites/13/2019/01/http2.jpg)
### HTTP/2 là gì?

HTTP/2, hay Hypertext Transfer Protocol phiên bản 2, là một bản nâng cấp lớn của [giao thức HTTP/1.1](https://azweb.com.vn/http-la-gi/). Được chính thức phê duyệt vào năm 2015 bởi Lực lượng chuyên trách về kỹ thuật liên mạng (IETF), HTTP/2 được xây dựng dựa trên giao thức SPDY do Google phát triển. Mục tiêu chính của nó không phải là thay đổi ngữ nghĩa của HTTP mà là cải thiện cách dữ liệu được định dạng và vận chuyển giữa client (trình duyệt) và server.

Về cơ bản, HTTP/2 ra đời để giải quyết các vấn đề cố hữu về hiệu suất của HTTP/1.1, vốn không được thiết kế cho các trang web phức tạp và giàu tài nguyên như hiện nay. Mục tiêu cốt lõi của giao thức này là giảm độ trễ, cho phép tải nhiều tài nguyên song song trên một kết nối duy nhất và nén dữ liệu hiệu quả hơn. Kết quả là tốc độ tải trang nhanh hơn, trải nghiệm người dùng mượt mà hơn và sử dụng tài nguyên mạng hiệu quả hơn, đồng thời vẫn giữ nguyên tính tương thích với các phương thức, mã trạng thái và tiêu đề quen thuộc của HTTP/1.1.

### Sự khác biệt giữa HTTP/2 và HTTP/1.1

Sự khác biệt căn bản giữa HTTP/2 và HTTP/1.1 nằm ở cách chúng xử lý và truyền tải dữ liệu. HTTP/1.1 hoạt động theo cơ chế tuần tự, nghĩa là trình duyệt phải gửi một yêu cầu và chờ nhận phản hồi trước khi gửi yêu cầu tiếp theo trên cùng một kết nối. Điều này gây ra một hiện tượng gọi là “Head-of-Line Blocking”, làm chậm đáng kể tốc độ tải trang khi có nhiều tài nguyên cần được xử lý. Để khắc phục, các trình duyệt thường phải mở nhiều kết nối [TCP/IP](https://azweb.com.vn/tcp-ip-la-gi/) song song, gây lãng phí tài nguyên và tăng độ phức tạp.

Ngược lại, HTTP/2 giới thiệu một cấu trúc hoàn toàn mới dựa trên luồng nhị phân (binary framing layer). Thay vì gửi văn bản thuần túy, HTTP/2 chuyển đổi các yêu cầu và phản hồi thành các khung nhị phân nhỏ hơn. Điều này cho phép nó hỗ trợ đa luồng (multiplexing), tức là gửi và nhận nhiều yêu cầu, phản hồi cùng lúc trên một kết nối TCP duy nhất. Hơn nữa, HTTP/2 còn có khả năng nén tiêu đề (header compression) và cho phép server chủ động đẩy tài nguyên (server push), giúp giảm thiểu số lượng yêu cầu và tăng tốc độ tải trang một cách ấn tượng.

![](https://cdn.prod.website-files.com/5ff66329429d880392f6cba2/676d4500b15423be50fc505f_6149cbd7fd4bdd7c82f55cc6_http1%2520vs%2520http2.png)

---
## Kỹ thuật chính của HTTP/2

#### **Đa luồng (Multiplexing)**

Đa luồng là một trong những cải tiến đột phá và quan trọng nhất của HTTP/2. Trong giao thức HTTP/1.1, nếu bạn muốn tải 10 tài nguyên (như ảnh, file CSS, file JavaScript), trình duyệt phải gửi 10 yêu cầu riêng biệt, thường là trên nhiều kết nối TCP khác nhau. Mỗi kết nối lại phải trải qua quá trình “bắt tay” (handshake) tốn thời gian, và các yêu cầu trên một kết nối phải chờ đợi nhau, gây ra tình trạng tắc nghẽn.

HTTP/2 giải quyết triệt để vấn đề này bằng cách cho phép gửi và nhận nhiều luồng dữ liệu (streams) cùng một lúc trên một kết nối TCP duy nhất. Hãy tưởng tượng HTTP/1.1 là một con đường một làn xe, nơi các xe phải nối đuôi nhau di chuyển. Còn HTTP/2 giống như một xa lộ nhiều làn, nơi các phương tiện có thể di chuyển song song mà không cản trở nhau. Lợi ích của đa luồng là giảm đáng kể độ trễ mạng, loại bỏ hoàn toàn tình trạng Head-of-Line Blocking, giúp website tải nhanh hơn rất nhiều, đặc biệt là các trang có nhiều tài nguyên nhỏ.

![](https://coolicehost.com/images/http2-how-it-works.jpg)
#### **Nén header (Header Compression)**

Mỗi khi trình duyệt gửi một yêu cầu đến server, nó đều đính kèm một loạt thông tin trong phần header, chẳng hạn như loại trình duyệt, ngôn ngữ chấp nhận, cookie, và nhiều thông tin khác. Trong HTTP/1.1, các header này được gửi dưới dạng văn bản thuần túy và thường lặp đi lặp lại giữa các yêu cầu, gây lãng phí băng thông không cần thiết. Khi một trang web có hàng trăm yêu cầu, tổng dung lượng của các header này có thể trở nên đáng kể.

HTTP/2 giới thiệu một cơ chế nén header thông minh gọi là HPACK. Thay vì gửi toàn bộ header cho mọi yêu cầu, HPACK duy trì một bảng tra cứu các header đã được gửi trước đó ở cả client và server. Trong các yêu cầu tiếp theo, nó chỉ cần gửi các giá trị đã thay đổi hoặc một chỉ mục tham chiếu đến các header đã có trong bảng. Kỹ thuật này giúp giảm đáng kể kích thước dữ liệu cần truyền đi, đặc biệt hiệu quả với các ứng dụng di động có băng thông hạn chế. Kết quả là thời gian tải trang được cải thiện và tài nguyên mạng được sử dụng hiệu quả hơn.
#### **Giữ kết nối lâu dài (Connection Persistence)**
#### **Stream**
#### **Stream ID**
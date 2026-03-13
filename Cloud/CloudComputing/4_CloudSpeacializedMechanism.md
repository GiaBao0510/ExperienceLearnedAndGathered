# Các Cơ Chế Hình Thành Chức Năng Đám Mây

---

## 1. Bộ Tiếp Nhận Yêu Cầu Co Giãn Tự Động (Automated Scaling Listener)

**Automated Scaling Listener** là một dịch vụ tác nhân (service agent) theo dõi và truy vết sự giao tiếp giữa người tiêu dùng dịch vụ đám mây và các dịch vụ đám mây, phục vụ mục tiêu **co giãn tự động**.

- Thường được triển khai cạnh **tường lửa (firewall)**, từ đó có thể theo dõi các thông tin liên quan đến trạng thái tải.
- **Tải (Workload)** có thể được xác định dựa vào:
    - Khối lượng các yêu cầu được tạo ra từ người tiêu dùng đám mây.
    - Các xử lý của hệ thống phía sau (Back-end) được kích hoạt bởi các kiểu yêu cầu khác nhau (ví dụ: lượng dữ liệu đầu vào lớn sinh ra từ nhiều xử lý đồng thời).

**Các kiểu đáp ứng khi phát hiện thay đổi tải:**

- **Tự động tăng/giảm** tài nguyên CNTT dựa trên các ngưỡng đã được thiết lập trước bởi người tiêu dùng đám mây.
- **Tự động báo hiệu** người quản trị khi tải vượt ngưỡng trên hoặc xuống dưới ngưỡng dưới, để họ chủ động điều chỉnh cấp phát tài nguyên.

![](https://patterns.arcitura.com/wp-content/uploads/2018/08/fig2-80.png)

**Mô tả sơ đồ trên:** Ba người tiêu dùng đám mây cùng truy cập đồng thời một dịch vụ đám mây (1). Automated Scaling Listener mở rộng và tạo 3 phiên bản (instance) mới cho dịch vụ được yêu cầu (2). Người tiêu dùng thứ tư gửi yêu cầu sử dụng dịch vụ (3). Vì đã được lập trình chỉ cho phép tối đa 3 phiên bản, Automated Scaling Listener từ chối yêu cầu thứ tư và gửi cảnh báo đến người quản trị biết tải đã vượt giới hạn (4). Người quản trị truy cập từ xa vào môi trường quản trị để tăng số lượng phiên bản tối đa được phép (5).

---

**Kịch bản: Co giãn theo chiều đứng (Scale Up) và di dời máy ảo (Live Migration)**

> _Một người tiêu dùng đám mây tạo và khởi động một máy chủ ảo với 8 nhân CPU ảo và 16 GB RAM ảo (1). VIM tạo máy chủ ảo theo yêu cầu, máy ảo tương ứng được cấp phát lên Máy chủ Vật lý 1, cùng với 3 máy ảo khác đang hoạt động (2). Nhu cầu từ người tiêu dùng làm mức sử dụng CPU của máy chủ ảo vượt 80% trong 60 giây liên tiếp (3). Automated Scaling Listener đang chạy trên Hypervisor phát hiện cần phải scale up và ra lệnh cho VIM xử lý (4)._

![](https://patterns.arcitura.com/wp-content/uploads/2018/08/fig2-81.png)

> _VIM xác định rằng không thể scale up máy chủ ảo trên Máy chủ Vật lý 1 (vì tài nguyên đang được dùng bởi các máy ảo khác), nên tiến hành **di dời trực tiếp (live migrate)** máy chủ ảo sang Máy chủ Vật lý 2 – máy có đủ tài nguyên hơn (5)._

![](https://patterns.arcitura.com/wp-content/uploads/2018/08/fig2-82.png)

> _Sau khi di dời, mức sử dụng CPU/RAM của máy chủ ảo duy trì dưới 15% trong 60 giây liên tiếp (6). Automated Scaling Listener phát hiện cần scale down và ra lệnh cho VIM (7), VIM tiến hành giảm tài nguyên của máy chủ ảo (scale down) trong khi máy ảo vẫn đang hoạt động trên Máy chủ Vật lý 2 (8)._

---

## 2. Bộ Cân Bằng Tải (Load Balancer)

Một trong những cách phổ biến để thực hiện **co giãn theo chiều ngang** là **phân phối tải đồng đều** giữa nhiều tài nguyên CNTT, giúp tăng hiệu năng và năng lực xử lý vượt qua giới hạn của một tài nguyên đơn lẻ.

> **Giải thích về "tài nguyên tạo ra tải" và "tài nguyên xử lý tải":**
> 
> - **Tài nguyên tạo ra tải (Workload Generator):** Là người dùng hoặc ứng dụng gửi các yêu cầu (request) lên hệ thống. Ví dụ: hàng nghìn người dùng cùng truy cập một website là nguồn tạo ra tải.
> - **Tài nguyên xử lý tải (Workload Handler):** Là các máy chủ ảo, dịch vụ hoặc ứng dụng phía sau nhận và xử lý các yêu cầu đó. Ví dụ: 3 máy chủ web đứng sau một Load Balancer cùng chia nhau xử lý các request.
> - **Load Balancer** đứng ở giữa, điều phối luồng request từ phía tạo ra tải đến các tài nguyên xử lý tải sao cho không có máy chủ nào bị quá tải.

![](https://www.cloud4u.com/upload/medialibrary/5a6/0_CCK15OF3DizmOITk.png)

**Các kiểu phân phối tải thường gặp:**

- **Phân phối không đồng đều (Weighted Distribution):** Các tài nguyên có năng lực lớn hơn sẽ nhận nhiều tải hơn tương ứng.
- **Theo độ ưu tiên (Priority-based):** Tải được lập lịch, xếp hàng, hủy bỏ hoặc phân phối dựa theo mức độ ưu tiên.
- **Dựa trên nội dung (Content-based):** Các yêu cầu được phân phối đến tài nguyên CNTT khác nhau tùy theo nội dung của yêu cầu (ví dụ: request về hình ảnh đi vào một nhóm máy chủ riêng, request API đi vào nhóm khác).

![](https://patterns.arcitura.com/wp-content/uploads/2018/08/fig2-116.png)

**Mô tả sơ đồ trên:** Load Balancer được triển khai như một service agent, phân phối trong suốt các yêu cầu đến từ người tiêu dùng đám mây ra hai phiên bản dịch vụ dự phòng, qua đó tối đa hóa hiệu năng.

**Cấu hình của Load Balancer:**

- Được cài đặt với các quy tắc về hiệu năng và chất lượng dịch vụ (QoS) với mục tiêu sử dụng tài nguyên tối ưu, tránh quá tải và tối đa hóa thông lượng.

**Các dạng triển khai Load Balancer:**

- **Bộ chuyển mạch mạng nhiều tầng (Multi-layer Network Switch):** Là thiết bị mạng hoạt động ở nhiều tầng giao thức (tầng 4 đến tầng 7), có thể đưa ra quyết định định tuyến dựa trên nội dung yêu cầu (ví dụ: URL, cookie, header HTTP), không chỉ dựa trên địa chỉ IP đơn thuần.
- **Thiết bị phần cứng tận hiến (Dedicated Hardware Appliance):** Là thiết bị phần cứng vật lý được thiết kế riêng chỉ để thực hiện chức năng cân bằng tải, không dùng cho mục đích nào khác, mang lại hiệu năng cao và độ ổn định tốt (ví dụ: các thiết bị F5 BIG-IP, Citrix ADC).
- **Hệ thống dựa trên phần mềm tận hiến:** Thường là phần mềm chạy trên hệ điều hành máy chủ chuyên dụng (ví dụ: HAProxy, NGINX).
- **Tác tử dịch vụ (Service Agent):** Thường được điều khiển bởi phần mềm quản trị đám mây, hoạt động như một proxy trung gian.

> **Làm rõ về vị trí và vai trò của Load Balancer:** Load Balancer được đặt trên **đường giao tiếp giữa tài nguyên tạo tải và tài nguyên xử lý tải**. Về phía người tiêu dùng dịch vụ, Load Balancer thường **trong suốt** – tức là người dùng không biết mình đang giao tiếp với Load Balancer hay máy chủ thật; họ chỉ thấy một địa chỉ duy nhất. Phía sau đó, Load Balancer đóng vai trò **proxy trừu tượng hóa** các tài nguyên xử lý tải – nó che giấu sự tồn tại của nhiều máy chủ phía sau, điều phối yêu cầu đến máy chủ phù hợp nhất tại thời điểm đó.

---

## 3. Bộ Theo Dõi Cam Kết Mức Độ Dịch Vụ (SLA Monitor)

**SLA (Service Level Agreement)** là hợp đồng cam kết chất lượng dịch vụ giữa nhà cung cấp đám mây và người tiêu dùng, quy định các chỉ số như: thời gian hoạt động (uptime), thời gian phản hồi, thông lượng tối thiểu, v.v.

**SLA Monitor** (Bộ theo dõi cam kết mức độ dịch vụ) được dùng để **đặc biệt theo dõi hiệu năng thực thi** của các dịch vụ, đảm bảo chúng đáp ứng đầy đủ các yêu cầu chất lượng đã cam kết trong SLA.

Dữ liệu thu thập được bởi SLA Monitor được xử lý bởi **hệ thống quản lý SLA** để tổng hợp thành các **chỉ số báo cáo về SLA**. Mục đích của việc tổng hợp này là:

- **Kiểm chứng** liệu nhà cung cấp có đang đáp ứng đúng các cam kết trong hợp đồng hay không.
- **Phát hiện sớm** các vi phạm SLA để có biện pháp can thiệp kịp thời.
- **Cơ sở tính phí và bồi thường:** Nếu SLA bị vi phạm, người tiêu dùng có thể yêu cầu bồi thường hoặc giảm phí dựa trên dữ liệu báo cáo này.
- **Tối ưu hóa dịch vụ:** Giúp nhà cung cấp biết chỗ nào cần cải thiện hạ tầng hoặc năng lực xử lý.

---

## 4. Bộ Theo Dõi Trả Theo Sử Dụng (Pay-Per-Use Monitor)

Cơ chế này **đo lường mức độ sử dụng** các tài nguyên CNTT trên đám mây tương ứng với các thông số giá đã định trước, và tạo ra các **nhật ký sử dụng (usage logs)** phục vụ mục đích tính phí và lập hóa đơn.

**Các thông số thường được theo dõi:**

- Số lượng thông điệp yêu cầu/phản hồi.
- Khối lượng dữ liệu được truyền tải.
- Mức độ tiêu thụ băng thông.

![](https://patterns.arcitura.com/wp-content/uploads/2018/08/fig2-123.png)

**Mô tả sơ đồ trên (theo dõi vòng đời dịch vụ):** Người tiêu dùng đám mây yêu cầu tạo một phiên bản mới của dịch vụ (1). Tài nguyên được khởi tạo và Pay-Per-Use Monitor nhận sự kiện "start" từ phần mềm tài nguyên (2). Monitor lưu mốc thời gian bắt đầu vào cơ sở dữ liệu nhật ký (3). Người tiêu dùng sau đó yêu cầu dừng dịch vụ (4). Monitor nhận sự kiện "stop" (5) và lưu mốc thời gian kết thúc vào nhật ký (6). Khoảng thời gian giữa start và stop chính là cơ sở để tính phí.

![](https://patterns.arcitura.com/wp-content/uploads/2018/08/fig3-53.png)

**Mô tả sơ đồ trên (theo dõi từng yêu cầu):** Người tiêu dùng dịch vụ đám mây gửi thông điệp yêu cầu đến dịch vụ (1). Pay-Per-Use Monitor chặn giữa thông điệp (2), chuyển tiếp đến dịch vụ (3a), đồng thời ghi lại thông tin sử dụng theo các thước đo đã định nghĩa (3b). Dịch vụ trả kết quả về cho người tiêu dùng (4).

---

## 5. Bộ Theo Dõi Sự Kiện (Audit Monitor)

Cơ chế này được dùng để **thu thập dữ liệu theo vết (log) các sự kiện** liên quan đến mạng và tài nguyên CNTT, nhằm hỗ trợ các yêu cầu tuân thủ quy định bắt buộc (compliance) – ví dụ: quy định bảo vệ dữ liệu, kiểm toán bảo mật.

![](https://patterns.arcitura.com/wp-content/uploads/2018/08/fig2-78.png)

**Mô tả sơ đồ:** Người tiêu dùng đám mây gửi yêu cầu đăng nhập kèm thông tin xác thực (1). Audit Monitor chặn giữa thông điệp (2) và chuyển tiếp đến dịch vụ xác thực (3). Dịch vụ xác thực xử lý và sinh ra phản hồi (4). Audit Monitor chặn giữa thông điệp phản hồi và lưu toàn bộ thông tin sự kiện đăng nhập vào cơ sở dữ liệu nhật ký theo chính sách kiểm toán của tổ chức (5). Quyền truy cập được cấp và phản hồi được gửi về cho người tiêu dùng (6).

---

## 6. Hệ Thống Phòng Chống Lỗi (Failover System)

Cơ chế này được dùng để **tăng độ tin cậy và tính sẵn dùng** của các tài nguyên CNTT bằng cách sử dụng kỹ thuật **phân cụm (Clustering)** và bổ sung thêm các thiết bị/phiên bản dự phòng.

Hệ thống phòng chống lỗi thường được dùng cho:

- **Các chương trình quan trọng (mission-critical programs):** Là những ứng dụng mà nếu bị gián đoạn sẽ gây thiệt hại nghiêm trọng, ví dụ: hệ thống ngân hàng, hệ thống y tế, thanh toán trực tuyến.
- **Các dịch vụ là "điểm chết" duy nhất (Single Point of Failure – SPOF):** Đây là các dịch vụ/thành phần mà nếu chúng ngừng hoạt động, sẽ kéo theo hàng loạt ứng dụng khác phụ thuộc vào chúng cũng bị tê liệt theo. Ví dụ: một dịch vụ xác thực (authentication service) dùng chung cho 20 ứng dụng – nếu nó chết thì cả 20 ứng dụng đó đều không thể đăng nhập được.

Cơ chế **nhân bản tài nguyên** đôi khi cũng được dùng bởi hệ thống phòng chống lỗi để tạo ra các phiên bản dự phòng, sẵn sàng thay thế khi phiên bản chính không còn hoạt động.

Có **2 cấu hình cơ bản**:

### 6.1 Active-Active

Tất cả các tài nguyên CNTT dự phòng **đều đang hoạt động cùng nhau** và cùng chia sẻ tải. Cần có bộ cân bằng tải để phân phối yêu cầu giữa các phiên bản. Khi một phiên bản bị hỏng, nó được loại ra khỏi danh sách của bộ cân bằng tải, các phiên bản còn lại tiếp tục xử lý.

**Ưu điểm:** Tận dụng tối đa tài nguyên, hiệu năng cao, thời gian phục hồi gần như bằng 0.

![63](https://media.geeksforgeeks.org/wp-content/uploads/20240509184654/Active-Active-Architecture-\(3\).webp)

**Kịch bản xử lý lỗi Active-Active:**

- Khi phát hiện lỗi ở một phiên bản, hệ thống phòng chống lỗi ra lệnh cho bộ cân bằng tải chuyển toàn bộ tải sang phiên bản còn lại.
- Phiên bản bị lỗi được phục hồi hoặc nhân bản thành phiên bản mới. Sau đó, hệ thống phòng chống lỗi ra lệnh cho bộ cân bằng tải phân phối lại tải bình thường.

### 6.2 Active-Passive

Các tài nguyên dự phòng ở **chế độ chờ (standby)** và sẽ được kích hoạt khi tài nguyên chính không còn hoạt động; tải sẽ được chuyển hướng sang tài nguyên dự phòng vừa kích hoạt.

**Phù hợp cho:** Các ứng dụng không có trạng thái (stateless), hoặc các hệ thống đòi hỏi tính đơn giản và tiết kiệm chi phí hơn Active-Active.

![](https://media.geeksforgeeks.org/wp-content/uploads/20240509184608/Active-Passive-Architecture-\(3\).webp)

**Kịch bản xử lý lỗi Active-Passive:**

- Hệ thống phòng chống lỗi liên tục theo dõi trạng thái hoạt động của dịch vụ đang là active. Phiên bản active đang nhận và xử lý yêu cầu từ người tiêu dùng đám mây.

![](https://media.geeksforgeeks.org/wp-content/uploads/20240509184614/Active-Passive-Architecture-Failover-\(2\).webp)

- Khi phiên bản active gặp lỗi, hệ thống phát hiện và kích hoạt phiên bản standby để thay thế, đồng thời chuyển hướng toàn bộ tải về phiên bản mới. Phiên bản bị lỗi sau đó được phục hồi hoặc nhân bản thành phiên bản standby mới, trong khi phiên bản vừa kích hoạt tiếp tục đóng vai trò active.

---

## 7. Bộ Ảo Hóa (Hypervisor)

**Hypervisor** là thành phần nền tảng của hạ tầng ảo hóa, được sử dụng chủ yếu để **tạo ra các phiên bản máy chủ ảo (VM)** trên một máy chủ vật lý.

**Đặc điểm và giới hạn của Hypervisor:**

- Bị **giới hạn trên một máy chủ vật lý** – chỉ có thể tạo các máy ảo từ tài nguyên của máy chủ đó.
- Chỉ có thể gán máy ảo vào các **vùng tài nguyên cùng nằm trên máy chủ vật lý** của nó.
- Chỉ có **tính năng quản trị cơ bản** cho máy ảo: thay đổi năng lực, tắt/bật máy ảo.

**VIM (Virtual Infrastructure Manager)** bổ sung một tập các tính năng quản trị nâng cao hơn, điều phối **nhiều Hypervisor trên nhiều máy chủ vật lý** khác nhau, cho phép:

- Di dời máy ảo từ máy chủ vật lý này sang máy chủ khác.
- Quản lý toàn bộ vòng đời của máy ảo từ một giao diện trung tâm.

**Ảo hóa bare-metal:** Hypervisor có thể cài đặt trực tiếp lên phần cứng máy chủ (không qua hệ điều hành – gọi là _Type 1 Hypervisor_ hay _bare-metal hypervisor_) để điều khiển, chia sẻ và lập lịch sử dụng các tài nguyên phần cứng (CPU, RAM, I/O). Nhờ đó, tài nguyên phần cứng xuất hiện trước hệ điều hành máy chủ ảo như thể chúng là tài nguyên vật lý riêng của máy ảo đó.

> **Ví dụ thực tế:** VMware ESXi, Microsoft Hyper-V, KVM là các Hypervisor phổ biến trong môi trường doanh nghiệp và đám mây.

![](https://patterns.arcitura.com/wp-content/uploads/2018/08/fig3-51.png)

**Mô tả sơ đồ:** Các máy chủ ảo được tạo ra bởi các Hypervisor riêng lẻ trên từng máy chủ vật lý. Cả ba Hypervisor được điều phối và quản lý tập trung bởi cùng một VIM.

---

## 8. Cụm Tài Nguyên (Resource Cluster)

Cơ chế này nhóm nhiều phiên bản tài nguyên CNTT lại với nhau và cho chúng **hoạt động như một tài nguyên duy nhất**, qua đó tăng năng lực tính toán, cân bằng tải và khả năng sẵn dùng.

**Kiến trúc cụm tài nguyên** dựa trên các **giao tiếp mạng tốc độ cao, tận hiến** giữa các phiên bản tài nguyên để trao đổi thông tin về:

- Phân phối tải, lập lịch, chia sẻ dữ liệu, đồng bộ hóa.

Các hoạt động này được đảm trách bởi một **nền tảng quản lý cụm** – là middleware chạy phân tán trên tất cả các nút của cụm, cho phép các tài nguyên phân tán xuất hiện như một tài nguyên CNTT thống nhất từ bên ngoài.

### Cụm Máy Chủ (Server Cluster)

Nhiều máy chủ được gom thành cụm để tăng năng lực và khả năng sẵn dùng. Các Hypervisor chạy trên các máy vật lý khác nhau có thể được cấu hình để **chia sẻ trạng thái thực thi** của các máy ảo (trang bộ nhớ, trạng thái thanh ghi), xây dựng thành **máy ảo được gom cụm**. Trong cấu hình này:

- Các máy chủ vật lý đòi hỏi truy cập vào **vùng lưu trữ chung**.
- Máy ảo có thể **di dời trực tiếp (live migrate)** từ máy chủ vật lý này sang máy khác trong khi vẫn đang chạy.
- Quá trình di dời **trong suốt** đối với hệ điều hành và ứng dụng bên trong máy ảo.
- Thường dùng để **giảm tải** từ máy chủ vật lý đang bị quá tải sang máy chủ còn tài nguyên trống.

### Cụm Cơ Sở Dữ Liệu (Database Cluster)

Được thiết kế để **cải thiện khả năng sẵn dùng của dữ liệu**. Các cụm này có tính năng **đồng bộ hóa dữ liệu** giữa nhiều thiết bị lưu trữ trong cụm để đảm bảo tính nhất quán. Năng lực dự phòng thường dựa trên cấu hình Active-Active hoặc Active-Passive.

### Cụm Bộ Dữ Liệu Lớn (Large Dataset Cluster)

Được thiết kế đặc biệt để xử lý **dữ liệu lớn (Big Data)**. Dữ liệu được **phân đoạn (sharding)** và **phân tán** ra nhiều nút trong cụm sao cho:

- Mỗi nút chỉ xử lý **một phần nhỏ của toàn bộ dữ liệu**, không cần biết đến các nút khác.
- Các nút hoạt động **độc lập với nhau** – không cần giao tiếp liên tục giữa các nút như trong các kiểu cụm khác.

> **Ví dụ dễ hiểu:** Hãy tưởng tượng cần đếm số lần xuất hiện của từng từ trong 1 triệu tài liệu. Thay vì một máy xử lý hết, ta chia: máy 1 xử lý tài liệu 1–200.000, máy 2 xử lý 200.001–400.000, v.v. Mỗi máy làm việc độc lập, cuối cùng gộp kết quả lại. Đây chính là nguyên lý của mô hình **MapReduce** được dùng trong Hadoop và các hệ thống Big Data hiện đại.

_Một số ví dụ về cụm tài nguyên:_

![](https://patterns.arcitura.com/wp-content/uploads/2018/08/fig2-130.png)

Cân bằng tải và nhân bản tài nguyên được triển khai thông qua một Hypervisor hỗ trợ cụm. Một mạng SAN chuyên dụng được dùng để kết nối cụm lưu trữ với cụm máy chủ, cho phép chia sẻ chung thiết bị lưu trữ đám mây. Điều này đơn giản hóa quá trình nhân bản lưu trữ, được thực hiện độc lập tại cụm lưu trữ.

![](https://patterns.arcitura.com/wp-content/uploads/2018/08/fig3-55.png)

Một cụm máy chủ liên kết lỏng lẻo (loosely coupled) có tích hợp bộ cân bằng tải. Không có lưu trữ chia sẻ. Nhân bản tài nguyên được dùng để sao chép thiết bị lưu trữ đám mây qua mạng bởi phần mềm cụm.

---

## 9. Bộ Trung Chuyển Đa Thiết Bị (Multi-Device Broker)

Cơ chế này được sử dụng để thực hiện **chuyển đổi dữ liệu trong lúc thực thi (runtime)** nhằm giúp một dịch vụ đám mây có thể được truy cập từ **nhiều loại chương trình và thiết bị** khác nhau.

> **Ví dụ dễ hiểu:** Một dịch vụ đám mây trả về dữ liệu định dạng XML. Tuy nhiên ứng dụng trên điện thoại di động cần JSON, còn ứng dụng IoT cần định dạng nhị phân. Multi-Device Broker đứng ở giữa, tự động chuyển đổi dữ liệu sang đúng định dạng mà mỗi loại thiết bị yêu cầu – mà không cần thay đổi gì ở dịch vụ gốc.

![](https://patterns.arcitura.com/wp-content/uploads/2018/08/fig2-120.png)

**Mô tả sơ đồ:** Multi-Device Broker chứa logic ánh xạ cần thiết để chuyển đổi dữ liệu trao đổi giữa dịch vụ đám mây và các loại thiết bị tiêu dùng khác nhau. Trong sơ đồ này, Multi-Device Broker được triển khai như một dịch vụ đám mây có API riêng. Cơ chế này cũng có thể được triển khai như một service agent chặn giữa thông điệp trong lúc thực thi để thực hiện các chuyển đổi cần thiết.

---

## 10. Cơ Sở Dữ Liệu Quản Lý Trạng Thái (State Management Database)

**Cơ sở dữ liệu quản lý trạng thái** là một thiết bị lưu trữ dùng để chứa **dữ liệu trạng thái tạm thời (temporary state data)** của các phần mềm đang chạy.

Thay vì lưu dữ liệu trạng thái trong **bộ nhớ RAM** (vốn bị giới hạn và mất đi khi chương trình tắt), các ứng dụng có thể ghi dữ liệu trạng thái tạm thời xuống cơ sở dữ liệu này. Điều này mang lại hai lợi ích:

- **Giảm lượng RAM cần tiêu thụ**, giúp ứng dụng có thể mở rộng quy mô dễ dàng hơn.
- **Dữ liệu trạng thái được chia sẻ** giữa nhiều phiên bản của dịch vụ, hỗ trợ co giãn theo chiều ngang.

Cơ sở dữ liệu trạng thái thường được dùng bởi các **dịch vụ đám mây có tác vụ thực thi dài** (long-running tasks), nơi cần lưu trạng thái trung gian của quá trình xử lý giữa các bước.

> **Ví dụ thực tế:** Redis và Memcached là các cơ sở dữ liệu lưu trữ trạng thái phổ biến trong kiến trúc đám mây, thường dùng để lưu phiên đăng nhập (session), kết quả tính toán tạm thời, hoặc hàng đợi tác vụ.
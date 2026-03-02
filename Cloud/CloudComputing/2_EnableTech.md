# Các Công Nghệ Tiền Đề Cho Điện Toán Đám Mây

---

## 1. Kiến Trúc Mạng Internet và Mạng Băng Thông Rộng

### 1.1 Internet và Mạng Băng Thông Rộng

- Tất cả các đám mây phải kết nối vào một mạng, đặc biệt là mạng Internet.
- Internet cho phép việc cấp phát tài nguyên CNTT từ xa và hỗ trợ truy cập từ khắp mọi nơi.
- Khả năng của các đám mây vì thế phát triển đồng hành với những tiến bộ về kết nối Internet và chất lượng dịch vụ mạng.
- Các nhà cung cấp dịch vụ Internet (ISP - Internet Service Providers) triển khai các mạng đường trục lớn nhất trong mạng Internet để kết nối các router liên kết mạng của các quốc gia khác nhau lại với nhau.

![](https://cdnintech.com/media/chapter/59149/1695905654/media/F2.png)

---

### 1.2 Kiến Trúc Mạng Internet

Kết nối quốc tế được thực hiện thông qua một sơ đồ mạng **phân cấp 3 tầng (tiers)**:

![](https://www.cisco.com/content/dam/cisco-cdc/site/images/legacy/assets/swa/img/anchor-info/network-architecture-628x353.png)
#### Tầng 1 (Tier 1) – Nhà cung cấp xương sống toàn cầu

Đây là các nhà cung cấp dịch vụ Internet lớn nhất thế giới, vận hành các tuyến cáp quang biển và đất liền vượt lục địa, đại dương để kết nối nhiều mạng trên toàn cầu. Tầng 1 sẽ kết nối xuống các nhà cung cấp cấp vùng thuộc Tầng 2.

> **Ví dụ thực tế:** AT&T, Verizon, Deutsche Telekom, NTT Communications, Telia Carrier. Đây là các "siêu ISP" – họ không trả phí để trao đổi dữ liệu với nhau mà kết nối theo dạng _peering_ (trao đổi lưu lượng ngang hàng miễn phí).

#### Tầng 2 (Tier 2) – Nhà cung cấp cấp vùng / quốc gia

Đây là các ISP hoạt động ở phạm vi quốc gia hoặc khu vực. Họ mua băng thông từ Tầng 1 và bán lại hoặc cung cấp kết nối cho các ISP nhỏ hơn ở Tầng 3.

> **Ví dụ thực tế:** Ở Việt Nam, VNPT và Viettel thuộc nhóm này – họ vận hành hạ tầng mạng trên toàn quốc và kết nối với các Tier 1 quốc tế qua các cáp quang biển.

#### Tầng 3 (Tier 3) – ISP cục bộ / nhà cung cấp cuối cùng

Đây là các ISP phục vụ trực tiếp người dùng cuối (cá nhân, hộ gia đình, doanh nghiệp nhỏ). Họ mua kết nối từ Tầng 2 và cung cấp Internet tới từng địa chỉ cụ thể.

> **Ví dụ thực tế:** Các nhà mạng địa phương, nhà cung cấp cáp quang FTTH tại các thành phố, tỉnh thành.

#### Cách các tầng kết nối với nhau

Các tầng kết nối với nhau thông qua các **điểm trao đổi Internet (IXP – Internet Exchange Point)** – là các địa điểm vật lý (thường là một trung tâm dữ liệu lớn) nơi nhiều ISP cùng kết nối với nhau để trao đổi dữ liệu trực tiếp, giảm độ trễ và chi phí thay vì phải đi vòng qua nhau. Ngoài IXP, các ISP cấp dưới cũng có thể mua **kết nối thuê bao** (transit) từ ISP cấp trên để được truyền dữ liệu ra toàn cầu.

![](https://techvccloud.mediacdn.vn/280518386289090560/2024/10/24/ixp-la-gi-17297612879231277276399-0-0-719-1280-crop-17297612940172088491874.jpg)

```
[Người dùng cuối]
       |
  [Tầng 3 – ISP cục bộ]
       |
  [Tầng 2 – ISP vùng/quốc gia]
       |
  [Tầng 1 – ISP toàn cầu] ←→ [Tầng 1 khác trên toàn thế giới]
```

---

### 1.3 Cơ Chế Liên Mạng

**Mạng chuyển gói không kết nối (Connectionless Packet Switching / Datagram Networks):** Dòng dữ liệu được chia thành nhiều **gói (packet)** có kích thước giới hạn và được chuyển tiếp từ nơi gửi đến nơi nhận thông qua các switch và router của các mạng trung gian. Một số giao thức thường dùng:

- **MAC (Medium Access Control Protocol):** điều khiển truy cập đường truyền.
- **IP (Internet Protocol):** định tuyến gói tin trên mạng Internet.

**Liên kết nối dựa trên router (Router-based Interconnectivity):** Router là thiết bị kết nối vào nhiều mạng. Khi một gói tin đến, router sẽ **đọc địa chỉ đích** trong gói tin và chuyển tiếp gói đó đến nhánh mạng tiếp theo, hướng đến đích cuối cùng.

---

### 1.4 Mô Hình Phân Tầng Giao Thức của Internet

Mạng Internet hoạt động dựa trên mô hình phân tầng – mỗi tầng thực hiện một vai trò cụ thể và giao tiếp với tầng kề trên/dưới theo quy tắc chuẩn (giao thức). Mô hình phổ biến nhất là **TCP/IP** với 4 tầng:

|Tầng|Tên|Vai trò|Ví dụ giao thức|
|---|---|---|---|
|4|Ứng dụng (Application)|Giao tiếp với người dùng và ứng dụng|HTTP, FTP, DNS, SMTP|
|3|Vận chuyển (Transport)|Đảm bảo truyền dữ liệu tin cậy|TCP, UDP|
|2|Mạng (Internet/Network)|Định tuyến gói tin qua các mạng|IP|
|1|Truy cập mạng (Network Access)|Truyền dữ liệu trên đường vật lý|Ethernet, Wi-Fi|

---

### 1.5 Những Mối Quan Tâm về Kỹ Thuật và Thương Mại

**Vấn đề kết nối (Connection Issues):**

- Người dùng từ bên ngoài truy cập vào _private cloud_ thông qua mạng riêng ảo (VPN) thiết lập trên mạng Internet.
- Người dùng từ bên trong mạng nội bộ truy cập vào đám mây bên ngoài thông qua Internet.
- Nên sử dụng cùng một giao thức cho truy cập tài nguyên nội bộ và tài nguyên trên đám mây để đảm bảo tính di động cho người dùng.

**Vấn đề băng thông và độ trễ mạng (Network Bandwidth and Latency Issues):**

![](https://www.keycdn.com/img/support/network-latency.png)

- Băng thông end-to-end (giữa hai đầu gửi–nhận) được xác định bởi khả năng truyền tải của các đường truyền trung gian được chia sẻ.
- Các ISP cần sử dụng mạng băng thông rộng cho đường trục để đảm bảo kết nối giữa hai đầu cuối.
- **Độ trễ** là khoảng thời gian để một gói tin đi từ nút này đến nút kia trên mạng; độ trễ sẽ tăng khi gói tin phải đi qua nhiều router trung gian.
- Các giải pháp CNTT cần đánh giá tác động của băng thông và độ trễ lên yêu cầu nghiệp vụ:
    - Một số ứng dụng đòi hỏi băng thông lớn để chuyển dữ liệu đến/từ đám mây.
    - Một số ứng dụng đòi hỏi độ trễ thấp để đảm bảo thời gian phản hồi nhanh.

**Chọn nhà cung cấp đám mây và nhà cung cấp truy cập đám mây (Cloud Provider và Cloud Carrier Selection):**

- **Nhà cung cấp đám mây (Cloud Provider):** Là tổ chức cung cấp tài nguyên tính toán, lưu trữ, và dịch vụ phần mềm cho người dùng (ví dụ: AWS, Google Cloud, Microsoft Azure).
![](https://www.weavertech.us/wp-content/uploads/2023/05/Comparing-Amazon-AWS-Google-Cloud-and-Microsoft-Azure-Choosing-the-Right-Cloud-Provider-for-Your-Business.png)
    
- **Nhà cung cấp truy cập đám mây (Cloud Carrier):** Là tổ chức đóng vai trò **trung gian kết nối mạng** giữa người dùng và nhà cung cấp đám mây. Họ không cung cấp tài nguyên điện toán, mà đảm bảo **đường truyền mạng** đạt chất lượng dịch vụ (QoS) nhất định – bao gồm băng thông, độ trễ, và độ ổn định. Nói đơn giản, Cloud Carrier giống như "hãng vận tải" còn Cloud Provider giống như "kho hàng".
![](https://www.techtarget.com/rms/onlineimages/whatis-carrier_cloud.png)
    
> **Ví dụ:** Một doanh nghiệp thuê máy chủ từ AWS (Cloud Provider), nhưng cần một đường truyền tốc độ cao, ổn định từ văn phòng tới AWS – họ sẽ ký hợp đồng với VNPT hay Viettel đóng vai trò Cloud Carrier để cung cấp đường truyền đó.

![](https://viettelcare.com.vn/wp-content/uploads/2020/05/Danh-sach-cac-nha-cung-cap-internet-cap-quang-Viet-Nam-ISP.jpg)
    
- Chất lượng kết nối giữa người dùng và đám mây phụ thuộc vào nhiều ISP trên đường đi. Đảm bảo chất lượng dịch vụ xuyên suốt nhiều ISP là bài toán phức tạp trong thực tế.
    
- Người dùng và nhà cung cấp đám mây có thể cần dùng **nhiều Cloud Carrier** để đảm bảo mức độ kết nối và độ tin cậy cho ứng dụng.
    
---

## 2. Công Nghệ Trung Tâm Dữ Liệu

### 2.1 Trung Tâm Dữ Liệu (Data Center)

Trung tâm dữ liệu (Data Center) là cơ sở hạ tầng đặc biệt được thiết kế để **tập trung với mật độ cao** các tài nguyên CNTT như máy chủ, cơ sở dữ liệu, thiết bị mạng, thiết bị viễn thông và hệ thống phần mềm.

![](https://static.vncdn.vn/vnetwork.vn/pub/websites/uploads/1/61/data-center.jpg)

**Lợi ích của việc tập trung tài nguyên CNTT vào một Data Center:**

- Chia sẻ nguồn cung cấp điện hiệu quả hơn.
- Tăng hiệu quả sử dụng tài nguyên CNTT thông qua chia sẻ chung.
- Tăng cường khả năng tiếp cận cho nhân viên kỹ thuật CNTT.

---
### 2.2 Công Nghệ và Thành Phần Liên Quan đến Data Center

#### Công Nghệ Ảo Hóa

![](https://cdn-media.sforum.vn/storage/app/media/wp-content/uploads/2019/12/virtualization-1-1.jpg)

- Các trung tâm dữ liệu bao gồm cả tài nguyên CNTT vật lý và tài nguyên được ảo hóa.
- **Lớp tài nguyên vật lý** bao gồm hạ tầng phần cứng để vận hành các hệ thống tính toán, mạng và phần mềm.
- **Lớp ảo hóa** cung cấp các công cụ để vận hành và quản lý tài nguyên – thường dựa trên một nền tảng ảo hóa trừu tượng hóa phần cứng thành các thành phần ảo có thể cấp phát, theo dõi và điều khiển dễ dàng hơn.
#### Chuẩn Hóa và Mô Đun Hóa

- Data Center được xây dựng dựa trên phần cứng đã được **chuẩn hóa** và thiết kế theo kiến trúc **mô đun hóa**, kết hợp nhiều khối hạ tầng giống nhau để hỗ trợ khả năng mở rộng, tăng trưởng và thay thế nhanh phần cứng.
- Đây là yêu cầu then chốt để giảm chi phí đầu tư và vận hành, vì nó giảm thiểu sự phức tạp trong các quy trình mua sắm, triển khai, vận hành và bảo trì.

#### Tự Động Hóa

- Data Center có các nền tảng đặc biệt để **tự động hóa** các tác vụ như cấp phát, cấu hình, vá lỗi, theo dõi mà không cần giám sát liên tục.
- Các tiến bộ trong nền tảng quản lý data center hỗ trợ **tính toán tự động** để cho phép tự cấu hình và tự phục hồi.

#### Vận Hành và Quản Lý Từ Xa

- Hầu hết các tác vụ vận hành và quản trị tài nguyên CNTT trong data center được thực hiện qua **giao diện dòng lệnh từ xa** và các hệ thống quản lý.
- Nhân viên kỹ thuật không cần trực tiếp đến phòng máy chủ, ngoại trừ các thao tác đặc biệt như thay thế phần cứng hoặc kéo dây mạng.

#### Độ Sẵn Sàng Cao (High Availability)
![](https://kb.syncplify.com/uploads/images/gallery/2023-07/SBOha1.png)
- Bất kỳ sự gián đoạn nào của data center đều ảnh hưởng nghiêm trọng đến sự liên tục trong kinh doanh.
- Data Center được thiết kế với **mức dự phòng cao**: nguồn điện dự phòng không gián đoạn (UPS), hệ thống cáp dự phòng, cân bằng tải, và các phần cứng cụm (cluster) để đối phó với sự cố.

#### Thiết Kế, Vận Hành và Quản Lý An Toàn

- Các yêu cầu về an ninh bao gồm kiểm soát truy cập cả ở tầng vật lý lẫn logic, cùng các chiến lược phục hồi sau sự cố.
- Xu hướng hiện nay là **chuyển tài nguyên CNTT ra các data center bên ngoài** do nhiều rào cản trong việc tự xây dựng và vận hành. Tuy nhiên, các mô hình thuê ngoài truyền thống thường đòi hỏi cam kết dài hạn và không hỗ trợ tính co giãn linh hoạt – những vấn đề mà điện toán đám mây giải quyết được (truy cập mọi nơi qua Internet, cấp phát theo yêu cầu, co giãn nhanh, trả theo mức sử dụng).

#### Tiện Ích Hỗ Trợ (Facilities)

Các tiện nghi của data center được thiết kế phù hợp cho thiết bị tính toán, lưu trữ và mạng, thường bao gồm: nguồn điện, hệ thống cáp, hệ thống làm mát (nhiệt, quạt, điều hòa), hệ thống phòng cháy chữa cháy và các hệ thống hỗ trợ khác.

---

#### Phần Cứng Tính Toán

Phần lớn việc xử lý trong data center được thực hiện bởi các **máy chủ tiêu chuẩn** với khả năng tính toán và lưu trữ lớn. Một số công nghệ phần cứng thường đi kèm với máy chủ mô đun:

- Các **tủ rack** tích hợp sẵn ổ cắm nguồn, mạng và hệ thống làm mát.
- Hỗ trợ nhiều kiến trúc xử lý: x86-32 bit, x86-64 bit, RISC.
- Kiến trúc CPU đa nhân với hàng trăm nhân trên diện tích nhỏ của một rack.
- Các thiết bị dự phòng **gắn nóng (hot-swap)**: đĩa cứng, bộ nguồn, card mạng – có thể thay thế mà không cần tắt hệ thống.

---
#### Phần Cứng Lưu Trữ

**Các công nghệ cho hệ thống lưu trữ:**

- **Mảng đĩa cứng (Hard Disk Arrays):** Cho phép phân tán và nhân bản dữ liệu trên nhiều đĩa để tăng tốc độ và dự phòng. Công nghệ điển hình: **RAID**.
- **Lưu trữ đệm vào ra (I/O Caching):** Tăng tốc độ truy cập đĩa bằng cách đệm dữ liệu tạm thời trong bộ nhớ nhanh hơn.
- **Đĩa cứng gắn nóng (Hot-swappable Hard Disks):** Cho phép tháo/lắp đĩa cứng mà không cần tắt nguồn.
- **Ảo hóa lưu trữ (Storage Virtualization):** Thực hiện thông qua các ổ đĩa ảo được chia sẻ giữa nhiều máy.
- **Cơ chế chụp ảnh (Snapshotting):** Lưu lại trạng thái bộ nhớ của máy ảo tại một thời điểm vào tập tin, có thể phục hồi lại sau.

**Các thiết bị lưu trữ mạng:**

- **SAN (Storage Area Network):** Thiết bị lưu trữ kết nối vào mạng, cho phép truy cập ở mức **khối dữ liệu (block-level)** sử dụng chuẩn SCSI. Phù hợp với các ứng dụng hiệu năng cao như cơ sở dữ liệu.
- **NAS (Network Attached Storage):** Mảng đĩa cứng kết nối vào mạng và cho phép truy cập dữ liệu bằng các giao thức dịch vụ tập tin như **NFS** hoặc **SMB**. Phù hợp với chia sẻ file trong mạng nội bộ.

---

#### Phần Cứng Mạng

Data Center cần nhiều thành phần phần cứng mạng để đảm bảo kết nối. Một data center điển hình có thể chia thành **5 hệ thống mạng thành phần**:

1. **Liên kết nối mạng bên trong và bên ngoài:** Bao gồm router backbone kết nối mạng WAN bên ngoài với mạng LAN bên trong, cùng các thiết bị an ninh biên như tường lửa (Firewall) và cổng VPN.
    
2. **Bộ cân bằng tải và tăng tốc web tầng 2:** Bao gồm bộ tiền xử lý XML, thiết bị mã hóa/giải mã, switch tầng 7 thực hiện định tuyến theo nội dung.
    
3. **Giàn chuyển mạch mạng cục bộ:** Tạo thành các mạng LAN nội bộ trong data center, cung cấp kết nối hiệu năng cao và dự phòng. Các switch tốc độ hàng chục Gbps hỗ trợ các chức năng ảo hóa mạng: chia VLAN, tổng hợp đường truyền, cân bằng tải.
    
4. **Giàn chuyển mạch cho SAN:** Kết nối server với các hệ thống lưu trữ SAN, thường dùng kênh truyền quang (Fiber Channel – FC), FC over Ethernet (FCoE), và InfiniBand.
    
5. **Cổng NAS:** Cung cấp điểm kết nối cho các thiết bị NAS và các giao thức truyền dữ liệu giữa SAN và NAS.
    

Năm hệ thống mạng trên cùng nhau cải thiện **khả năng dự phòng và độ tin cậy**, cho phép data center duy trì dịch vụ ngay cả khi xảy ra hỏng hóc ở một thành phần nào đó.

Các đường cáp quang siêu tốc được sử dụng để tổng hợp nhiều kênh truyền riêng lẻ vào một sợi cáp nhờ kỹ thuật **ghép kênh (multiplexing)**, cải thiện tốc độ và giảm độ trễ, đặc biệt khi data center trải rộng qua nhiều địa điểm.

---

### 2.3 Kỹ Thuật Ảo Hóa (Virtualization)

**Ảo hóa** là quá trình chuyển đổi các tài nguyên CNTT vật lý thành các tài nguyên CNTT ảo (dựa trên phần mềm mô phỏng).

**Các loại tài nguyên có thể được ảo hóa:**

|Tài nguyên vật lý|Tài nguyên ảo tương ứng|
|---|---|
|Máy chủ vật lý|Máy chủ ảo (Virtual Machine)|
|Thiết bị lưu trữ|Ổ đĩa ảo (Virtual Disk)|
|Router, Switch vật lý|Mạng LAN ảo (VLAN), Switch ảo|
|UPS vật lý|UPS ảo|

**Các bước tạo ra một máy chủ ảo:**

1. Cấp phát một máy chủ vật lý để ảo hóa.
2. Cài đặt hệ điều hành lên máy chủ vật lý (host OS).
3. Cài đặt phần mềm ảo hóa (Hypervisor).
4. Tạo máy ảo (VM) trên nền phần mềm ảo hóa.

> **Lưu ý:** Hệ điều hành của máy chủ ảo (_guest OS_) hoạt động **độc lập** với hệ điều hành của máy chủ vật lý (_host OS_).

Phần mềm ảo hóa bao gồm các dịch vụ chuyên biệt để quản lý máy ảo. Phần mềm này thường được gọi là **Hypervisor** (hay Virtual Machine Monitor/Manager). Một số Hypervisor phổ biến: VMware ESXi, Microsoft Hyper-V, KVM.

**Các tính năng và lợi ích nổi bật của ảo hóa:**

- **Độc lập với phần cứng:** Máy ảo không phụ thuộc vào phần cứng cụ thể, nên có thể dễ dàng di chuyển sang máy chủ vật lý khác mà không gặp vấn đề tương thích.
    
- **Hợp nhất máy chủ (Server Consolidation):** Nhiều máy ảo chạy trên cùng một máy vật lý, giúp tận dụng tối đa tài nguyên và giảm số lượng phần cứng cần mua.
    
- **Nhân bản tài nguyên:** Máy ảo được lưu dưới dạng **ảnh đĩa ảo (virtual disk image)** – một tập tin nhị phân. Nhờ đó có thể dễ dàng:
    
    - Tạo ra các **template chuẩn** sẵn sàng triển khai nhanh.
    - Di dời và mở rộng máy ảo linh hoạt.
    - **Quay lui (rollback)** về trạng thái trước nhờ cơ chế chụp ảnh (snapshot).
    - Hỗ trợ kinh doanh liên tục (Business Continuity) nhờ sao lưu và phục hồi hiệu quả.

**Hai loại ảo hóa chính:**

- **Ảo hóa dựa trên hệ điều hành (OS-based Virtualization):** Máy ảo chạy bên trên host OS. Nhược điểm: host OS tiêu thụ tài nguyên, lời gọi hệ thống phải đi qua nhiều lớp làm giảm hiệu năng, cần bản quyền cho cả host OS và guest OS.
    
- **Ảo hóa dựa trên phần cứng (Hardware-based Virtualization):** Máy ảo tương tác trực tiếp với phần cứng thông qua Hypervisor, giúp cải thiện hiệu năng. Cần đảm bảo các trình điều khiển (driver) phần cứng được hỗ trợ trong Hypervisor.
    

**Quản lý ảo hóa:**

Các phần mềm ảo hóa hiện đại cung cấp các công cụ **Quản lý Hạ Tầng Ảo Hóa (VIM – Virtualization Infrastructure Management)** để quản lý tập trung toàn bộ tài nguyên ảo hóa, thông qua một **Bộ điều khiển (Controller)** chạy trên một máy chủ chuyên dụng. Điều này cho phép tự động hóa các thao tác quản trị và giảm bớt công sức vận hành đáng kể.
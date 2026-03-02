# Các Cơ Chế Hạ Tầng Đám Mây (Cloud Infrastructure Mechanisms)

Các cơ chế hạ tầng đám mây là những **khối xây dựng nền tảng** của môi trường đám mây. Chúng thiết lập các thành phần cơ bản tạo nên kiến trúc công nghệ đám mây.

---

## 1. Đường Biên Mạng Luận Lý (Logical Network Perimeter)

**Đường biên mạng luận lý** được định nghĩa là sự cô lập một môi trường mạng khỏi phần còn lại của mạng truyền thông. Nó hình thành một vùng ranh giới bao bọc và cô lập một nhóm các tài nguyên CNTT dựa trên đám mây – những tài nguyên này về mặt vật lý có thể nằm phân tán ở nhiều nơi khác nhau.

> **Giải thích dễ hiểu hơn:** Hãy hình dung bạn có một công ty với nhiều chi nhánh đặt ở các tỉnh thành khác nhau. Mỗi chi nhánh có máy chủ riêng, nhưng bạn muốn tất cả chúng hoạt động như thể đang nằm trong cùng một mạng nội bộ bảo mật – tách biệt hoàn toàn với người dùng bên ngoài Internet. Đó chính là ý tưởng của Logical Network Perimeter: **tạo ra một "vùng mạng riêng ảo"** bảo mật trên nền hạ tầng vật lý phân tán. Điểm khác biệt với mạng nội bộ (LAN) thông thường là ranh giới này hoàn toàn được định nghĩa bằng phần mềm (logic), không phụ thuộc vào vị trí địa lý hay cáp vật lý.

Cơ chế này được triển khai để:

- Cô lập tài nguyên CNTT khỏi những người không được cấp quyền truy cập.
- Cô lập tài nguyên CNTT khỏi các đối tượng không phải người dùng hợp lệ.
- Cô lập tài nguyên CNTT giữa các người tiêu dùng đám mây khác nhau (mỗi khách hàng không nhìn thấy tài nguyên của nhau).
- Kiểm soát băng thông được cấp phát cho các tài nguyên đã được cô lập.

Đường biên mạng luận lý thường được hình thành thông qua các thiết bị mạng ảo hóa, bao gồm:

- **Tường lửa ảo (Virtual Firewall):** Là một tài nguyên CNTT thực hiện việc lọc luồng giao thông (traffic) vào/ra một mạng, kiểm soát giao tiếp của mạng đó với Internet và các mạng bên ngoài.
- **Mạng LAN ảo (VLAN – Virtual LAN):** Là tài nguyên CNTT dùng để cô lập các môi trường mạng trong cùng một nền tảng hạ tầng đám mây. Nhiều VLAN có thể cùng chạy trên cùng một hạ tầng vật lý nhưng hoàn toàn tách biệt nhau về mặt logic.

![](https://conradjohns.weebly.com/uploads/1/0/0/2/100257408/untitled_17_orig.png)

---

## 2. Máy Chủ Ảo (Virtual Server)

**Máy chủ ảo (Virtual Server)** – hay còn gọi là **máy ảo (Virtual Machine / VM)** – là một phần mềm giả lập hoàn chỉnh một máy chủ vật lý. Các nhà cung cấp đám mây sử dụng máy chủ ảo để **chia sẻ một máy chủ vật lý cho nhiều người dùng** bằng cách cấp cho mỗi người một **phiên bản (instance)** riêng biệt.

![](https://www.cherryservers.com/v3/assets/blog/2025-05-15/img-01.jpg)

Máy chủ ảo là **khối xây dựng cơ bản** của môi trường đám mây. Mỗi máy chủ ảo có thể được dùng để triển khai (host) các tài nguyên CNTT, các giải pháp dựa trên đám mây và các cơ chế điện toán đám mây khác.

![](https://www.researchgate.net/publication/324235969/figure/fig1/AS:622404095188993@1525404108862/Example-state-of-the-art-virtualization-platform.png)

Việc tạo ra một phiên bản máy chủ ảo từ một **tập tin ảnh (image file)** là tiến trình cấp phát tài nguyên có thể được hoàn thành **nhanh chóng và theo yêu cầu** (on-demand). Đây là một trong những ưu điểm nổi bật của đám mây so với việc mua và cài đặt phần cứng vật lý truyền thống.

Người tiêu dùng đám mây thuê một máy chủ ảo có thể **tùy biến môi trường** trên máy chủ ảo của mình **hoàn toàn độc lập** với môi trường của các máy chủ ảo của những người dùng khác đang chạy trên cùng một máy chủ vật lý.

---

## 3. Thiết Bị Lưu Trữ Đám Mây (Cloud Storage Device)

### 3.1 Tổng Quan

Cơ chế lưu trữ đám mây đại diện cho các thiết bị lưu trữ được thiết kế đặc biệt cho việc **cấp phát dựa trên đám mây**. Các thiết bị này có thể được ảo hóa và cung cấp với **dung lượng có thể tăng lên theo nhu cầu** người sử dụng.

**Các thiết bị lưu trữ mạng phổ biến bao gồm:**

|Thiết bị|Mô tả|
|---|---|
|**SAN (Storage Area Network)**|Mạng lưu trữ chuyên dụng, cung cấp truy cập ở mức khối (block-level), tốc độ cao|
|**NAS (Network Attached Storage)**|Thiết bị lưu trữ gắn vào mạng, truy cập theo giao thức tập tin (NFS, SMB)|
|**Object Storage**|Lưu trữ dữ liệu dưới dạng đối tượng (object) có địa chỉ URL riêng, phù hợp đám mây|
|**Ổ đĩa ảo (Virtual Disk)**|Ổ đĩa được ảo hóa, gắn vào máy chủ ảo như ổ đĩa thật|
|**Distributed File System**|Hệ thống tập tin phân tán trải rộng nhiều máy chủ (ví dụ: HDFS, GlusterFS)|

Các thiết bị lưu trữ mạng có thể được **công bố ra bên ngoài** để truy cập từ xa thông qua các **dịch vụ lưu trữ đám mây (Cloud Storage Services)**. Các dịch vụ đám mây phổ biến hiện nay bao gồm:

- **Amazon S3 (Simple Storage Service):** Lưu trữ đối tượng (object storage) quy mô lớn, truy cập qua API HTTP/REST.
- **Google Cloud Storage:** Tương tự S3, tích hợp hệ sinh thái Google Cloud.
- **Microsoft Azure Blob Storage:** Lưu trữ dữ liệu phi cấu trúc trên nền Azure.
- **Dropbox / Google Drive / OneDrive:** Dịch vụ lưu trữ tập tin dành cho người dùng cá nhân và doanh nghiệp nhỏ.
- **Các dịch vụ cơ sở dữ liệu đám mây:** Amazon RDS, Google Cloud SQL, Azure SQL Database (lưu trữ có cấu trúc).

Vấn đề quan tâm chính đối với lưu trữ đám mây là **an toàn, toàn vẹn và độ tin cậy của dữ liệu**. Ngoài ra còn có:

- Vấn đề **pháp lý** khi dữ liệu được lưu trữ vượt ra ngoài biên giới quốc gia (data sovereignty).
- Vấn đề **hiệu năng** đối với các cơ sở dữ liệu lớn.

---

### 3.2 Các Mức Độ Lưu Trữ Đám Mây

Các cơ chế lưu trữ đám mây cung cấp các đơn vị lưu trữ dữ liệu ở những mức độ trừu tượng khác nhau:

|Mức độ|Mô tả|Ví dụ|
|---|---|---|
|**Tập tin (Files)**|Dữ liệu được nhóm thành tập tin, tổ chức trong thư mục|NFS, Google Drive|
|**Khối (Blocks)**|Đơn vị dữ liệu nhỏ nhất, gần với phần cứng, truy cập độc lập|SAN, AWS EBS|
|**Bộ dữ liệu (Datasets)**|Dữ liệu tổ chức theo bảng (table) và mẫu tin (record)|CSV, cơ sở dữ liệu|
|**Đối tượng (Objects)**|Dữ liệu + metadata được nhóm thành tài nguyên web có địa chỉ URL|AWS S3, Azure Blob|

![](https://bashfulbytes.com/static/cloudstorage.png)

Mỗi mức lưu trữ có một số loại **giao diện kỹ thuật (API/Interface)** tương ứng để truy cập, phù hợp với loại thiết bị lưu trữ hoặc dịch vụ đám mây cụ thể.

---

### 3.3 Giao Diện Lưu Trữ Mạng (Network Storage Interface)

Nhóm giao diện lưu trữ mạng bao gồm các thiết bị tuân theo các giao thức công nghiệp chuẩn. Có thể chia thành 2 loại chính:

**Lưu trữ khối (Block Storage) – sử dụng giao thức SCSI:**

- Dữ liệu được chia thành các **khối (block)** có kích thước cố định, đây là đơn vị nhỏ nhất có thể lưu trữ và truy xuất.
- Định dạng lưu trữ **rất gần với phần cứng** – hệ điều hành làm việc trực tiếp với các khối giống như ổ đĩa vật lý.
- **Giao thức:** SCSI (Small Computer System Interface) và các biến thể như iSCSI (qua mạng IP).

**Lưu trữ tập tin (File Storage) – sử dụng giao thức SMB, CIFS, NFS:**

- **SMB (Server Message Block) / CIFS (Common Internet File System):** Giao thức của Microsoft, phổ biến trong môi trường Windows, cho phép chia sẻ tập tin và thư mục qua mạng.
- **NFS (Network File System):** Giao thức của Linux/Unix, cho phép mount (gắn kết) thư mục từ xa như thể đang ở máy cục bộ.
- Dữ liệu được lưu trong các tập tin có thể khác nhau về kích thước và định dạng, tổ chức trong hệ thống thư mục quen thuộc.

> **Vì sao lưu trữ khối nhanh hơn lưu trữ tập tin?** Lưu trữ mức khối nhanh hơn vì nó **bỏ qua các tầng trung gian**. Với lưu trữ tập tin, khi đọc/ghi dữ liệu, hệ thống phải xử lý qua nhiều lớp: giao thức mạng (NFS/SMB) → hệ thống tập tin → sau đó mới đến ổ đĩa. Ngược lại, lưu trữ khối giao tiếp **trực tiếp với ổ đĩa ở mức thấp nhất**, không qua hệ thống tập tin trung gian, giảm độ trễ và tăng thông lượng đáng kể. Đó là lý do các cơ sở dữ liệu hiệu năng cao thường dùng lưu trữ khối thay vì lưu trữ tập tin.

---

### 3.4 Giao Diện Lưu Trữ Đối Tượng (Object Storage Interface)

**Lưu trữ đối tượng (Object Storage)** là kiểu lưu trữ hiện đại được thiết kế cho môi trường đám mây quy mô lớn. Điểm khác biệt so với lưu trữ tập tin hay lưu trữ khối là:

- Mỗi **đối tượng (object)** bao gồm: **dữ liệu thực sự** (nội dung file) + **siêu dữ liệu (metadata)** mô tả dữ liệu đó + một **định danh duy nhất (ID hoặc URL)**.
- Dữ liệu không được tổ chức theo cây thư mục mà được truy cập **qua URL hoặc ID**, giống như một tài nguyên web.
- Hỗ trợ **nhiều loại dữ liệu**: hình ảnh, video, tài liệu, bản sao lưu, log file, dữ liệu đa phương tiện...
- Dễ dàng mở rộng quy mô đến **hàng petabyte** mà không cần cấu trúc phân cấp phức tạp.

**Cách truy cập:** Thông qua **giao thức REST** hoặc các **Web Services API** sử dụng HTTP/HTTPS. Ví dụ: để tải lên một file ảnh, bạn gửi một HTTP PUT request đến endpoint của dịch vụ; để tải về, bạn dùng HTTP GET với URL của đối tượng đó.

> **Ví dụ thực tế:** Amazon S3, Google Cloud Storage, Azure Blob Storage đều là các dịch vụ lưu trữ đối tượng. Khi bạn upload ảnh lên Facebook hay xem video trên YouTube, dữ liệu thực chất được lưu trong các hệ thống object storage quy mô lớn.

---

### 3.5 Giao Diện Lưu Trữ Cơ Sở Dữ Liệu (Database Storage Interface)

Các cơ chế lưu trữ đám mây dựa trên **giao diện cơ sở dữ liệu** thường hỗ trợ thêm **ngôn ngữ truy vấn** (như SQL) bên cạnh các tác vụ lưu trữ thông thường. Việc quản lý lưu trữ được thực hiện thông qua **API chuẩn** hoặc giao diện quản trị.

Được chia thành 2 nhóm dựa vào cấu trúc lưu trữ:

---

### 3.6 Lưu Trữ Dữ Liệu Quan Hệ (Relational Data Storage)

Là kiểu lưu trữ dữ liệu **truyền thống**, dựa trên **cơ sở dữ liệu quan hệ** và hệ thống quản lý cơ sở dữ liệu quan hệ (RDBMS – Relational Database Management System).

**Đặc điểm:**

- Tổ chức dữ liệu theo **Tables (bảng)**, **Rows (hàng)**, và các **mối quan hệ giữa các bảng (Table Relationships)**.
- Đảm bảo **toàn vẹn dữ liệu**, tránh dư thừa thông qua chuẩn hóa.
- Sử dụng ngôn ngữ truy vấn **SQL (Structured Query Language)**.
- Các giải pháp phổ biến: MySQL, PostgreSQL, Microsoft SQL Server, Oracle DB.

**Hạn chế trong môi trường đám mây:**

- Gặp vấn đề về **co giãn (scalability)** và hiệu năng khi dữ liệu lớn.
- Hỗ trợ **co giãn theo chiều đứng** (nâng cấp phần cứng mạnh hơn) tốt hơn, nhưng tốn kém hơn **co giãn theo chiều ngang** (thêm nhiều máy chủ).
- Với nhiều mối quan hệ phức tạp và dung lượng lớn sẽ trở nên **chậm**, đặc biệt khi truy cập từ xa.

---

### 3.7 Lưu Trữ Dữ Liệu Không Quan Hệ (Non-Relational Data Storage / NoSQL)

**NoSQL** (Not Only SQL) chấp nhận cấu trúc **linh hoạt hơn** cho dữ liệu lưu trữ, ít nhấn mạnh đến việc định nghĩa các mối quan hệ và chuẩn hóa dữ liệu.

**Mục tiêu:** Giảm độ phức tạp và chi phí tính toán thường phát sinh trong cơ sở dữ liệu quan hệ, đặc biệt khi xử lý dữ liệu lớn (Big Data) hoặc dữ liệu không có cấu trúc cố định.

**Ưu điểm:**

- Cho phép **co giãn theo chiều ngang** dễ dàng (thêm nhiều máy chủ).
- Hiệu năng cao hơn với dữ liệu lớn và truy cập đồng thời nhiều.
- Các nhà cung cấp đám mây thường hỗ trợ NoSQL với **khả năng sẵn dùng và mở rộng cao** trong môi trường nhiều máy chủ.

**Đánh đổi (Trade-offs):**

- Không thể hiện đầy đủ các mô hình nghiệp vụ phức tạp.
- Không hỗ trợ đầy đủ các tính năng của RDBMS như **transaction** (giao dịch nguyên tử) và **JOIN** (kết hợp bảng).
- Dữ liệu sau khi chuyển từ SQL sang NoSQL sẽ trở nên **phi chuẩn hóa** (denormalized) – kích thước dữ liệu tăng lên nhưng đổi lại truy vấn đơn giản và nhanh hơn.

> **Ví dụ thực tế:** MongoDB (lưu trữ document), Redis (lưu trữ key-value), Cassandra (lưu trữ column-family), Amazon DynamoDB.

---

### 3.8 Sử Dụng Dịch Vụ Lưu Trữ Đám Mây

> **Gợi ý hình ảnh minh họa:** Bạn có thể tìm kiếm các hình ảnh với từ khóa sau để đặt vào đây:
> 
> - `"cloud storage service diagram"` – sơ đồ tổng quan các dịch vụ lưu trữ đám mây.
> - `"AWS S3 vs EBS vs EFS comparison"` – so sánh các loại lưu trữ của AWS.
> - `"cloud storage types infographic"` – infographic so sánh file/block/object storage.

---

### 3.9 Máy Chủ Ảo và Thiết Bị Lưu Trữ Đám Mây

> **Gợi ý hình ảnh minh họa:** Bạn có thể tìm kiếm các hình ảnh với từ khóa sau để đặt vào đây:
> 
> - `"virtual machine attached cloud storage diagram"` – sơ đồ máy ảo gắn kết với thiết bị lưu trữ.
> - `"VM instance with block storage and object storage"` – mối quan hệ giữa VM và các loại lưu trữ.
> - `"cloud compute and storage architecture"` – kiến trúc tổng thể tính toán + lưu trữ đám mây.

---

## 4. Bộ Theo Dõi Sử Dụng Đám Mây (Cloud Usage Monitor)

### 4.1 Tổng Quan

**Bộ theo dõi sử dụng đám mây (Cloud Usage Monitor)** là một phần mềm nhẹ, hoạt động độc lập, chịu trách nhiệm **thu thập và xử lý dữ liệu về mức độ sử dụng tài nguyên CNTT** trên đám mây.

Tùy thuộc vào loại **thước đo (metrics)** được chọn và cách thức thu thập dữ liệu, Bộ theo dõi có thể tồn tại dưới nhiều dạng khác nhau. Mỗi loại được thiết kế để chuyển tiếp dữ liệu thu thập được về một **cơ sở dữ liệu nhật ký (log database)** dùng cho xử lý và báo cáo về sau.

Có 3 loại bộ theo dõi chính:

- **Tác nhân theo dõi (Monitor Agent)**
- **Tác nhân tài nguyên (Resource Agent)**
- **Tác nhân truy vấn (Polling Agent)**

---

### 4.2 Tác Nhân Theo Dõi (Monitor Agent)

Là một **phần mềm trung gian**, hoạt động **dựa trên sự kiện**, tồn tại dưới dạng một **tác nhân dịch vụ (Service Agent)** đặt dọc theo các đường giao tiếp mạng. Nó theo dõi và phân tích dòng dữ liệu một cách **trong suốt** (transparent) – tức là không can thiệp vào luồng dữ liệu, chỉ quan sát và ghi nhận.

> **Ví dụ:** Giống như một camera an ninh đặt tại hành lang – nó ghi lại tất cả hoạt động đi qua mà không làm gián đoạn luồng người.

---

### 4.3 Tác Nhân Tài Nguyên (Resource Agent)

**Tác nhân tài nguyên** là một module thu thập dữ liệu sử dụng bằng cách **tương tác trực tiếp với tài nguyên phần mềm** tương ứng dựa trên sự kiện.

Được dùng để theo dõi các **thước đo sử dụng** dựa trên các sự kiện đã được định nghĩa trước ở mức phần mềm tài nguyên, ví dụ như:

- Khởi tạo (provisioning) một máy ảo mới.
- Tạm dừng (suspend) một dịch vụ.
- Khởi động lại (resume) một tài nguyên.
- Co giãn theo chiều đứng (vertical scaling) – thay đổi cấu hình tài nguyên.

---

### 4.4 Tác Nhân Truy Vấn (Polling Agent)

**Tác nhân truy vấn** là một module xử lý thu thập dữ liệu bằng cách **định kỳ gửi truy vấn (poll)** đến các tài nguyên CNTT để kiểm tra trạng thái.

Thường được dùng để theo dõi định kỳ trạng thái của tài nguyên, ví dụ như: tài nguyên đang hoạt động bình thường hay đã kết thúc/lỗi.

> **So sánh nhanh 3 loại tác nhân:**
> 
> - **Monitor Agent** – "nghe ngóng" luồng dữ liệu đi qua.
> - **Resource Agent** – "lắng nghe sự kiện" từ chính tài nguyên khi có thay đổi.
> - **Polling Agent** – "hỏi thăm định kỳ" tài nguyên xem tình trạng như thế nào.

---

## 5. Nhân Bản Tài Nguyên (Resource Replication)

**Nhân bản tài nguyên** được định nghĩa là việc **tạo ra nhiều phiên bản (instance) giống hệt nhau** của cùng một tài nguyên CNTT.

Cơ chế này được thực hiện khi cần cải thiện:

- **Khả năng sẵn dùng (Availability):** Nếu một phiên bản bị lỗi, các phiên bản còn lại vẫn hoạt động bình thường.
- **Hiệu năng (Performance):** Phân tán tải công việc ra nhiều phiên bản để phục vụ nhiều người dùng đồng thời hơn.

**Kỹ thuật ảo hóa** được sử dụng để triển khai cơ chế nhân bản – nhờ ảo hóa, việc tạo ra một bản sao của máy chủ ảo chỉ đơn giản là sao chép tập tin ảnh đĩa (image file) và khởi chạy phiên bản mới.

<div style="text-align: center;"> <img src="https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcRx8JIwJiTJkiHfPIv3ECWobC_DF4ID_n45PQ&s" width="500" /> </div>

---

## 6. Môi Trường Sẵn Dùng (Ready-Made Environment)

**Môi trường sẵn dùng** là thành phần cơ bản trong mô hình dịch vụ **PaaS (Platform as a Service)**. Đây là một nền tảng dựa trên đám mây bao gồm một tập hợp các tài nguyên CNTT đã được **cài đặt sẵn, cấu hình và sẵn sàng để người tiêu dùng đám mây sử dụng ngay**, ví dụ như: cơ sở dữ liệu, phần mềm trung gian (middleware), công cụ phát triển và công cụ quản lý.

<div style="text-align: center;"> <img src="https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcS776NkQU-RrmAQI9uTYkTLau20hfJqdJh0-g&s" width="500" /> </div>

Môi trường này được người tiêu dùng đám mây sử dụng để **phát triển và triển khai dịch vụ và phần mềm lên đám mây** mà không cần lo lắng về việc quản lý hạ tầng bên dưới.

**Phần mềm trung gian (Middleware)** được tích hợp sẵn trong các nền tảng **đa thuê bao (multi-tenant)** để hỗ trợ việc phát triển và triển khai các ứng dụng web một cách đồng thời cho nhiều khách hàng khác nhau.

Một số nhà cung cấp đám mây cung cấp môi trường thực thi linh hoạt cho các dịch vụ đám mây, tính phí dựa trên **hiệu năng thực thi** hoặc **theo gói giá đã mua** trước.

> **Ví dụ thực tế:** Google App Engine, Heroku, Microsoft Azure App Service, AWS Elastic Beanstalk – đây đều là các "môi trường sẵn dùng" theo mô hình PaaS. Nhà phát triển chỉ cần đẩy code lên, nền tảng sẽ tự lo phần còn lại (máy chủ, cơ sở dữ liệu, cân bằng tải...).
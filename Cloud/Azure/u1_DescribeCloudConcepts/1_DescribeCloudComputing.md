
### **Mô hình trách nhiệm chung là gì?**

![](https://www.influentialsoftware.com/wp-content/webp-express/webp-images/uploads/2024/01/cloud-storage-background-business-network-design-900x450.jpg.webp)

Bắt đầu với một trung tâm dữ liệu truyền thống của công ty. Tại đây công ty sẻ có trách nhiệm cho việc duy trì không gian vật lý, đảm bảo an ninh và duy trì hoặc thay thế máy chủ nếu có bất kỳ chuyện gì đó xảy ra. Bộ phận IT trong công ty sẽ chịu trách nhiệm cho việc bảo trì tất cả cấu trúc cơ sở hạ tần và phần mềm để duy trì hoạt động của trung tâm dữ liệu. Họ cũng có thể chịu trách nhiệm giữa cho tất cả hệ thống được vá lỗi và sử dụng phiên bản chính xác.

Với  **mô hình trách nhiệm chung**, những sự trách nhiệm sẽ được chia sẽ giữa nhà cung cấp đám mây và người tiêu dùng. **Trách nhiệm của bên cung cấp đám mây** gồm: Bảo mật vật lý, cung cấp nguồn, làm mát và kết nối mạng. Người tiêu dùng sẽ không được bố trí làm việc với trung tâm dữ liệu, không có nghĩa là bân người tiêu dùng sẽ không có bất kỳ trách nhiệm nào.

Đối với một số việc thì trách nhiệm sẽ dựa trên hoàn cảnh. Nếu bạn sử dụng **Cloud SQL Database**, thì bên cung cấp cloud sẽ có trách nhiệm duy trì cơ sở dữ liệu thực tế. Tuy nhiên bạn vẫn sẽ chịu trách nhiệm về việc dữ liệu đưa vào cơ sở dữ liệu. Nếu bạn triển khai máy ảo và cài đặt **SQL database** trên nó, thì bạn sẽ chịu trách nhiệm về bản vá và cập nhật database. Cũng như việc duy trì cơ sở dữ liệu và thông tin lưu trữ trong cơ sở dữ liệu.

Với ==trung tâm dữ liệu tại chỗ thì bạn sẽ chịu toàn bộ trách nhiệm về mọi thứ.== Với điện toán đám mây thì trách nhiệm đó sẽ được thay đổi. **Mô hình trách nhiệm chung** sẽ gắn chặt chẽ với loại dịch vụ đám mây: ***Infrastucture as a service (IaaS), platform as a service (PaaS) và Software as a service (SaaS)***. 

- **(IaaS)** sẽ đặt trách nhiệm cao nhất lên **người tiêu dùng** với nhà cung cấp đám mây bắt đầu chịu trách nhiệm cho bảo mật mức vật lý cơ bản, nguồn điện và kết nối. 
- Ở bên kia đầu quang phổ, **(SaaS)** sẽ đặt trách nhiệm cao nhất lên phía nhà cung cấp đám mây.
- **(Paas)**, được đặt là trung gian giữa **(SaaS)** và **(IaaS)**, nằm đâu đó ở giũa và phân bổ tránh nhiệm đến cả 2 là người tiêu dùng và nhà cung cấp đám mây.

Sau đây là **sơ đồ nêu bật trách nhiệm chung** sẽ thông báo ai sẽ chịu trách nhiệm gì dựa trên loại dịch vụ đám mây
![](https://learn.microsoft.com/en-us/training/wwl-azure/describe-cloud-compute/media/shared-responsibility-b3829bfe.svg)

Khi sủ dụng nhà cloud provider, bạn sẽ luôn có trách nhiệm cho:
- Thông tin và dữ liệu được lưu trữ trong đám mây.
- Các thiết bị kết cho phép kết nối đến đám mây của bạn (điện thoại, máy tính, v.v).
- Những tài khoản và định dạng mọi người, các dịch vụ và các thiết bị trong tổ chức của bạn,

Bên nhà cung cấp đám mây sẽ chịu trách nhiệm về:
- Trung tâm dữ liệu vật lý,
- Kết nối mạng vật lý.
- Các máy chủ vật lý.

Mô hình dịch vụ của bạn sẽ xác định trách nhiệm cho:
- Hệ điều hành.
- Điều khiển mạng
- Các ứng dụng.
- Định dạng và cơ sở hạ tầng.
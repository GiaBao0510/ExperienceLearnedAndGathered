### **Connection Pool là gì?**

![](https://cdn.prod.website-files.com/62f2651b2519afce5252ae74/667adbf0131d5ad800f04042_Configuring%20a%20Connection%20Pool_1%403x.png)

Trước khi để hiểu rõ về **Connection Pool** là gì. Thì chúng ta sẽ nói qua cách mà ứng dụng thông thường tương tác với cơ sở dữ liệu. Thông thường khi ứng dụng cần truy cập dữ liệu, nó sẽ thực hiện các bước gồm: thiết lập kết nối đến **CSDL**, thực hiện các câu truy vấn cần thiết và đóng kết nối. ==Quá trình thiết lập và đóng kết nối là một hoạt động tốn thời gian và tài nguyên.==

**Connection pool là gì?** **Connection pool** (hay còn được gọi là "Bể kết nối") là một kỹ thuật quản lý kết nối **CSDL**, duy trì một tập hợp các kết nối đã được thiết lập và sẵn sàng để sử dụng. Thay vì thiết lập một kết nối mới cho mỗi yêu cầu, ứng dụng sẽ lấy một kết nối có sẵn từ **connection pool**, sử dụng nó, và sau đó trả lại **connection pool** để các yêu cầu khác có thể sử dụng.

Hãy tưởng tượng **Connection Pool** giống như một bãi đậu xe chung. Thay vì mỗi lần cần xe, bạn phải gọi một chiếc taxi mới đến (tốn thời gian chờ đợi), thì bạn có thể đến bãi đậu xe và lấy ngay một chiếc xe có sẵn để sử dụng. Sau khi dùng xong, bạn không vứt bỏ chiếc xe mà trả nó về bãi đậu để người khác có thể sử dụng tiếp. Điều này giúp tiết kiệm thời gian và tài nguyên, thay vì phải tạo mới một chiếc xe mỗi lần cần di chuyển.

Tương tự, trong lập trình, **Connection Pool** giúp ==tái sử dụng các kết nối cơ sở dữ liệu thay vì tạo mới mỗi lần==, giúp cải thiện hiệu suất hệ thống và giảm tải cho máy chủ.

---
### ***Tại sao Connection Pooling lại quan trọng?***

![](https://cdn.hashnode.com/res/hashnode/image/upload/v1679465367726/39d66d75-53e6-49fb-8a50-0de67f834617.png?w=1600&h=840&fit=crop&crop=entropy&auto=compress,format&format=webp)

Mở một kết nối mới tới **CSDL** có thể tiêu tốn từ 20ms đến 50ms hoặc lâu hơn neeys máy chủ đang bị quá tải. Nếu mỗi request mở và đóng một kết nối nối mới, hệ thống sẽ nhanh chóng gặp phải tình trạng **nghẽn**. Dưới đây là những lý do nên áp dụng **connection pooling**:

###### 1. **Hiệu suất cao hơn:** Tiết kiệm thời gian mở kết nối mới.
###### 2. **Giảm tải cho máy chủ database:** Ít kết nối mới hơn, máy chủ sẽ không phải  xử lý liên tục mọi kết nối
###### 3. **Tiết kiệm tài nguyên hệ thống:** Giảm lượng bộ nhớ CPU tiêu thụ do không phải tạo và hủy kết nối liên tục.
###### 4. **Đáp ứng tốt với hệ thống có lượng truy vấn cao:** Ứng dụng có khả năng phục vụ nhiều người dùng hơn mà không gặp các vấn đề về hiệu suất.

---
### **Các Connection Pool hoạt động: Vòng đòi của một kết nối**

Để hiểu sâu hơn về **Connection pool**, chúng ta hãy xem xét qua vòng đời của **connection pool**
###### 1. **Khởi tạo (initialization)**
Khi ứng dụng khởi động, connection pool sẽ được tạo và cấu hình. Quá trình này bao gồm việc: ==thiết lập số lượng kết nối ban đầu đến cơ sở dữ liệu==(Thường được gọi là số lượng kết nối tối thiểu). Các thông số cấu hình khác có thể bao gồm: ==số lượng kết nối tối đa== mà **connection pool** có thể quản lý, ==thời gian chờ tối đa để lấy một kết nối== và các thiết lập liên quan đến việc kiểm tra tính hợp lệ của kết nối.

###### 2. **Yêu cầu kết nối (Request a connection)**
Khi ứng dụng cần tương tác với cơ sở dữ liệu, thay vì tạo ra một kết nối mới, nó sẽ gửi yêu cầu đến **connection pool** để lấy một kết nối đang rảnh.

###### 3. **Mượn kết nối (Borrowing a connection)**
Nếu có một kết nối rảnh trong **conection pool**, nó sẽ được "mượn" cho ứng dụng sử dụng. Ứng dụng sẽ sử dụng kết nối này để thực hiện các truy vấn và thao tác với cơ sở dữ liệu.  

###### 4. **Sử dụng kết nối (Using the connection)**
Trong thời gian kết nối được mượn, ứng dụng có toàn quyền sử dụng nó để tương tác với cơ sở dữ liệu.

###### 5. **Trả lại kết nối (Returning a connection)**
Sau khi hoàn thành các thao tác cần thiết, ứng dụng sẽ trả lại kết nối cho **Connection pool**. Lúc này, kết nối sẽ trở lại trạng thái rảnh và sẵn sàng phục vụ các yêu cầu khác. Điều quan trọng là ứng dụng phải đảm bảo trả lại kết nối sau khi sử dụng xong để tránh tình trạng cạn kiệt kết nối trong **connection tool**.

###### 6. **Quản lý kết nối (Connection management)**
**Connection pool** có trách nhiệm quản lý các kết nối trong "Bể". Nó có thể bao gồm các hoạt động:
- **Kiểm tra tính hợp lệ của kết nối:** Định kỳ kiểm tra xem các kết nối trong "Bể" còn hoạt động tốt hay không. Nếu một kết nối bị lỗi, **connection pool** có thể đóng nó và tạo một kết nối mới để thay thế.
- **Quản lý số lượng kết nối:** Đảm bảo số lượng kết nối trong "bể" luôn nằm trong giới hạn đã cấu hình từ trước (tối thiểu & tối đa). Nếu số lượng trong bể  đang được sử dụng vượt quá một ngưỡng nhất định, **connection pool** có thể tạo thêm kết nối mới (nếu chưa đạt đến giới hạn tối đa). Nếu có quá nhiều kết nối rảnh trong một thời gian dài, **connection pool** có thể đóng bớt các kết nối để tiết kiệm tài nguyên
- **Xử lý timeout:** Nếu một ứng dụng giữ một kết nối quá lâu mà không trả lại, **connection pool** sẽ thu hồi kết nối đó để đảm bảo tính sẵn sàng cho các yêu cầu khác.

###### 7. **Mở rộng và giới hạn kết nối (Scaling and Connection Limits)**
**Connection pool** có thể được cấu hình để tự động điều chỉnh số lượng kết nối dựa trên tải của ứng dụng. *Ví dụ*: khi lưu lượng truy cập tăng cao, **connection pool** có thể tự động tạo thêm kết nối mới (trong phạm vi giới hạn tối đa). Việc thiết lập giới hạn tối đa cho số lượng kết nối là rất quan trọng để tránh làm việc quá tải cơ sở dữ liệu

---
### **Lợi ích khi sử dụng Connection Pool**

Việc sử dụng **connection pool** mang lại nhiều lợi ích đáng kể cho hiệu suất và độ ổn định của ứng dụng.

![](https://miro.medium.com/v2/resize:fit:720/format:webp/1*zR8NPu1FEV4ndiRSY05YDA.png)
###### 1. **Cải thiện hiệu suất (improved performance):**
Lợi ích lớn nhất của **connection pool** chính là việc ==cải thiện hiệu suất đáng kể==. Việc thiết lập một kết nối mới đến **CSDL** là một quá trình tốn thời gian, bao gồm việc thiết lập kên giao tiếp, xác thực, và khởi tạo phiên làm việc. Khi áp dụng **connection pool**, các kết nối đã được thiết lập sẳn, giúp giảm thiểu đáng kế thời gian chờ đợi cho mỗi yêu câu truy cập dữ liệu.

###### 2. **GIảm tiêu thụ tài nguyên (Reduced Resouce Consumption):**
Việc tái sử dụng các ==kết nối đã được thiết lập hiệu quả hơn so với việc liên tục tạo và hủy kết nối mới==. Điều này giúp giảm tải các ứng dụng và máy chủ **CSDL**, tiết kiệm, tài nguyên hệ thống như CPU, bộ nhớ và băng thông mạng.

###### 3. **Nâng cao khả năng mở rộng (Enhanced Scalability):**
Với **connection pool**, ==ứng dụng có thể xử lý nhiều yêu cầu truy cập dữ liệu đồng thời hơn mà không gặp phải tình trạng nghẽn cổ chai do việc thiết lập kết nối chậm==. Điều này giúp ứng dụng có khả năng mở rộng tốt hơn để đáp ứng lượng truy cập ngày càng tăng.

###### 4. **Quản lý kết nối tốt hơn (Better Connection Management):**
**Connection pool** cung cấp một cơ chế tập trung để quản lý các kết nối CSDL. Nó giúp theo dõi của các kết nối, xử lý các kết nối bị lỗi và đảm bảo rằng luôn có đủ kết nối sẳn sàng để phục vụ cho ứng dụng.

###### 5. **Đơn giản hóa mã ứng dụng (Simplified Application Code):**
Khi sử dụng **connection pool**, các nhà phát triển không cần phải lo lắng về việc thiết lập và đóng kết nối cho mỗi lần truy cập dữ liệu. Thay vào đó, họ chỉ cần lấy một kết nối từ **connection pool** và trả lại sau khi sử dụng xong. Điều này giúp mã ứng dụng trở nên gọn gàng và dễ bảo trì hơn.

###### 6. **Tăng độ tin cậy (Increased Reliability):**
Bằng cách tự động quản lý và tái sử dụng các kết nối, **connection pool** giúp giảm thiểu nguy cơ xảy ra lỗi do việc thiết lập kết nối không thành công hoặc do số lượng kết nối vượt quá giới hạn của CSDL. ĐIều này góp phần tăng độ tin cậy cho ứng dụng.

---
### **Triển khai Connection Pool:**

Hầu hết các ngôn ngữ lập trình và framwork phát triển web phổ biến điều cung cấp các thư viện **connection pool** mạnh mẽ:

##### 1. **Thư viện và Framework (Libraries and Frameworks)**
Dưới đây là một số ví dụ về thư viện **connection pool** trong các ngôn ngữ lập trình khác nhau:

- **Java:** HikariCP, Apache Commons DBCP, C3P0. HikariCP thường được đánh giá là một trong những thư viện **connection pool** hiệu suất cao nhất cho Java.
- **Python:** SQLAlchemy (thường được sử dụng với các driver cơ sở dữ liệu như Psycopg2 cho PostgreSQL hoặc PyMySQL cho MySQL), aiomysql (cho asynchronous MySQL).
- **Node.js:** mysql, pg, sequelize (ORM tích hợp sẵn **connection pool**).
- **.NET:** System.Data.SqlClient (cung cấp sẵn **connection pool**).

##### 2. **Cấu hình (Configuration):**
Khi sử dụng thư viên **connection pool**, bạn cần cấu hình các thông số quan trong như: 
- **Số lượng kết nối tối thiểu (Minimun Idle Connection):** Số lượng kết nối luôn được duy trong trong "bể" ngay cả khi không có yêu cầu nào.
- **Số lượng kết nối tối đa (Maximum Active Connections):** Số lượng tối đa mà "bể" có thể được tạo ra.
- **Thời gian chờ kết nối tối đa (Maximum wait Time):** Thời gian tối đa mà một yêu cầu chờ để lấy được một kết nối từ "bể" khi tất cả các kết nối đều đang được sử dụng.
- **Thời gian tồn tại tối đa của kết nối (Maximum LifeTime):** Thời gian tối đa mà một kết nối có thể tồn tại trong "bể" trước khi bị đóng và thay thế. 
- **Thời gian kiểm tra kết nối (Connection Timeout):** thời gian chờ thiết lập một kết nối mới.

##### 3. **Các thực hành tốt nhất (Best Practies):**
Để sử dụng **connection pool** một cách hiệu quả, bạn nên tuân theo một số thực hành tốt nhất sau:
- **Đóng kết nối sau khi sử dụng:** Luôn đảm bảo rằng bạn trả lại kết nối cho **connection pool** sau khi đã hoàn thành các thao tác cần thiết, thường là trong khối **finally** để đảm bảo kết nối được trả ngay cả khi có lỗi xảy ra
- **Theo dõi thống kê của connection pool:** Hầy hết các thư viện **connection pool** đều cung cấp các ==thông tin thống kế về số lượng kết nối đang được sử dụng, số lượng kết nối rảnh, thời gian chờ đợi, .v.v==. Theo dõi các thông số này giúp bạn điều chỉnh cấu hình **connection pool** ==cho phù hợp với tải của ứng dụng==.
- **Thử nghiệm với các cấu hình khác nhau:** Tìm ra cấu hình **connection pool** tối ưu cho ứng dụng của bạn thông qua việc thử nghiệm với các giá trị khác nhau cho các thông số cấu hình


---
### **Các trường hợp sử dụng Connection Pool**
**Connection pool** là một kỹ thuật quan trọng và được sử dụng rộng rãi trong hầu hết các ứng dụng có tương tác với cơ sở dữ liệu, bao gồm:
- **Ứng dụng web:** Các trang web và ứng dụng web thường xuyên tương tác với cơ sở dữ liệu để hiển thị thông tin và xử lý yêu cầu của người dùng.
- **Ứng dụng doanh nghiệp:** Các ứng dụng quản lý nghiệp vụ, hệ thống ERP, CRM thường xuyên làm việc với cơ sở dữ liệu lớn.
- **Microservices:** Trong kiến trúc microservices, mỗi service có thể tương tác với một hoặc nhiều cơ sở dữ liệu, và việc sử dụng **connection pool** là rất quan trọng để đảm bảo hiệu suất của từng service.
- **Bất kỳ ứng dụng nào thường xuyên tương tác với cơ sở dữ liệu:** Dù là ứng dụng desktop, ứng dụng di động hay các hệ thống backend, nếu có nhu cầu truy cập cơ sở dữ liệu thường xuyên, việc sử dụng **connection pool** sẽ mang lại lợi ích đáng kể.

---
### **Cách cấu hình Connection Poolin trong c#:**

Để bắt đầu, dưới đây là một chuỗi kết nối đơn giản có bật connection pooling:
```csharp
string connectionString = "Server=mysql.example.com;Port=3306;Database=E_commerce;Uid=ecommerce_user;Pwd=SecurePass123!;Connection Timeout=30;DefaultCommandTimeout=30;Pooling=true;Min Pool Size=10;Max Pool Size=200;";
```

Giải thích:
- **Pooling=true:** kích hoạt connection pooling (mặc định đã được bật).
- **Min pool size=5:** Đảm bảo luôn có ít nhất 5 kết nối trong pool.
- **Max pool size=100:** Tối đa 100 kết nối cùng một lúc.

Nếu bạn không muốn sử dụng connection pooling (trong trường) 

---
### **Ví dụ thực tế: Connection Pooling trong c#:**

Sau đây là đoạn mã C# đơn giản để mô phỏng việc mở và đóng kết nối nhiều lần. Nếu connection pooling hoạt động đúng, tốc độ truy vấn sẽ nhanh hơn rất nhiều so với việc tạo mới kết nối mỗi lần

```csharp
using System;
using MySql.Data.MySqlClient;

string connectionString = "Server=mysql.example.com;Port=3306;Database=E_commerce;Uid=ecommerce_user;Pwd=SecurePass123!;Connection Timeout=30;DefaultCommandTimeout=30;Pooling=true;Min Pool Size=10;Max Pool Size=200;";

for(int i = 0; i < 1000; i++){
	
	using(var connection = new MySqlConnection(connectionString))
	{
		connection.Open();	
		Console.WriteLine($"Kết nối thứ {i + 1} đã mở thành công.");
	} 

	Console.WriteLine("Hoàn thành việc sử dụng connection pooling."); 
	Console.ReadLine();
}
```

---
### **Các vấn đề tiềm ẩn khi sử dụng Connection Pool:**

- **Rò rỉ kết nối (Connection Leak)**: Nếu ứng dụng không trả lại kết nối sau khi sử dụng (do lỗi lập trình hoặc ngoại lệ không được xử lý), Connection Pool có thể cạn kiệt kết nối, dẫn đến lỗi "timeout" hoặc "pool exhausted".
- **Cấu hình không tối ưu**: Nếu số lượng kết nối tối đa quá thấp, ứng dụng sẽ bị nghẽn khi lưu lượng tăng; ngược lại, nếu quá cao, cơ sở dữ liệu có thể bị quá tải.
- **Kết nối "treo" (Stale Connections)**: Kết nối trong pool có thể trở nên không hợp lệ (do mạng gián đoạn hoặc máy chủ CSDL khởi động lại) nếu không được kiểm tra định kỳ.

---
### **Kết luận:**

 **Connection pooling là một trong những tối ưu quan trọng như đơn giản nhất khi làm việc với cơ sở dữ liệu.**

---
### **Tài liệu tham khảo:**

1. [Connection Pool Là Gì? Tối Ưu Hiệu Suất Kết Nối Cơ Sở Dữ Liệu Cho Ứng Dụng Của Bạn](https://t3h.com.vn/tin-tuc/connection-pool-la-gi-toi-uu-hieu-suat-ket-noi-co-so-du-lieu-cho-ung-dung-cua-ban#)
2.  [Tối Ưu Hiệu Suất Ứng Dụng Bằng Connection Pooling Trong C#](https://kungfutech.edu.vn/posts/toi-uu-hieu-suat-ung-dung-bang-connection-pooling-trong-c)
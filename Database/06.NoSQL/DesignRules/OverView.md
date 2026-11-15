## **Thiết kế cơ sở dữ liệu NoSQL là gì?**

NoSQL (thường được hiểu là "Not Only SQL") là nhóm cơ sở dữ liệu không dựa trên mô hình bảng quan hệ truyền thống. Thiết kế cơ sở dữ liệu NoSQL không chỉ đơn giản là xác định cách lưu trữ dữ liệu, mà còn bao gồm một loạt quyết định về tổ chức và quản lý dữ liệu trong ==môi trường có schema linh hoạt và phân tán==. Trái ngược với cơ sở dữ liệu quan hệ truyền thống (RDBMS), nơi schema thường được xác định trước, thiết kế NoSQL thường tập trung vào việc tối ưu hóa truy cập dữ liệu và khả năng mở rộng của hệ thống.

Thông qua việc phân tích yêu cầu của dự án và ==đánh giá các mẫu truy vấn==, các lập trình viên xác định cách tổ chức dữ liệu phù hợp với nhu cầu của ứng dụng. Điều này bao gồm việc chọn khóa (trong nhiều hệ thống là khóa phân vùng - partition key, quyết định dữ liệu nằm ở node nào), quyết định các trường (field) cần được index và chọn mô hình dữ liệu phù hợp: key-value, document, column-family, đồ thị (graph) hoặc chuỗi thời gian (time-series).

> - Thuật ngữ "khóa chính (primary key)" thường dùng cho RDBMS. Trong NoSQL, phần quan trọng nhất của khóa là khóa phân vùng nên đã bổ sung.

---

## Tại sao lại cần phải có NoSQL

- NoSQL được phát triển xuất phát từ yêu cầu cần những cơ sở dữ liệu có khả năng lưu trữ lượng dữ liệu cực lớn, truy vấn với tốc độ cao, mở rộng ngang (horizontal scaling) bằng cách thêm nhiều máy phổ thông thay vì nâng cấp một máy chủ cực mạnh, và có khả năng chịu lỗi tốt.
- Đây là những yêu cầu mà cơ sở dữ liệu quan hệ khó đáp ứng hiệu quả khi mở rộng ngang, vì các đặc tính như join giữa nhiều bảng và giao dịch (transaction) trên nhiều node làm tăng đáng kể độ phức tạp và chi phí.
- Lượng dữ liệu mà các hệ thống cần xử lý ngày càng lớn. Ví dụ, Google và Facebook phải lưu trữ và xử lý lượng dữ liệu cực lớn mỗi ngày. Các hệ thống nội bộ của họ (như Bigtable của Google, Cassandra do Facebook khởi xướng) là những ví dụ tiêu biểu thúc đẩy sự phát triển của NoSQL.

![https://images.viblo.asia/a0352191-139d-4ca0-a4c9-5669017544e6.png](https://images.viblo.asia/a0352191-139d-4ca0-a4c9-5669017544e6.png)

---

## **Điểm mạnh của thiết kế cơ sở dữ liệu NoSQL**

Một điểm mạnh của thiết kế cơ sở dữ liệu NoSQL là tính linh hoạt, cho phép dễ dàng thay đổi cấu trúc dữ liệu để đáp ứng các yêu cầu kinh doanh mới hoặc mô hình truy vấn khác nhau. Điều này thường dẫn đến quá trình thiết kế được thực hiện theo cách lặp lại và có thể điều chỉnh theo thời gian, khi yêu cầu của ứng dụng thay đổi hoặc khi dữ liệu và môi trường vận hành phát triển.

Tuy nhiên, schema linh hoạt không có nghĩa là không có schema. Cơ sở dữ liệu không còn ép buộc cấu trúc, nên trách nhiệm đó chuyển sang tầng ứng dụng (cách tiếp cận thường gọi là schema-on-read). Vì vậy, việc điều chỉnh và cải thiện schema, bao gồm quản lý phiên bản của cấu trúc tài liệu và chuyển đổi dữ liệu cũ, vẫn là một phần quan trọng trong việc duy trì và phát triển hệ thống NoSQL.

Ngoài tính linh hoạt, các điểm mạnh khác như khả năng mở rộng và tính sẵn sàng cao được trình bày ở phần "Một số đặc điểm chung". Các điểm mạnh này đi kèm đánh đổi: NoSQL thường hạn chế join và truy vấn ad-hoc, thường phải nhân bản dữ liệu để đọc nhanh, và các đảm bảo về tính nhất quán khác với RDBMS. Do đó, quyết định thiết kế cần cân nhắc cả hiệu suất, khả năng mở rộng và mức nhất quán mà ứng dụng yêu cầu.

---

## **Cách thiết kế cơ sở dữ liệu RDBMS và NoSQL có gì khác nhau?**

![https://itviec.com/blog/wp-content/uploads/2024/04/nosql-sql-vippro-1600x1037.png](https://itviec.com/blog/wp-content/uploads/2024/04/nosql-sql-vippro-1600x1037.png) _So sánh các kiểu thiết kế dữ liệu NoSQL và RDBMS (sử dụng SQL). (Nguồn: @scylladb)_

### **RDBMS**

- Áp dụng cấu trúc schema được xác định trước.
- Dữ liệu được tổ chức thành các bảng, mỗi bảng đại diện cho một thực thể riêng biệt.
- Thiết kế đi từ dữ liệu: xác định thực thể và quan hệ trước, sau đó dùng SQL (với join) để truy vấn theo nhiều cách khác nhau.
- Áp dụng kỹ thuật chuẩn hóa dữ liệu (normalization) để giảm dữ liệu trùng lặp, từ đó tăng tính nhất quán và toàn vẹn của dữ liệu.
- Quan hệ giữa các bảng được xác định rõ ràng và được bảo vệ bởi schema cùng các ràng buộc của cơ sở dữ liệu (khóa chính, khóa ngoại,...).

> "thiết kế đi từ dữ liệu" để tạo đối trọng rõ ràng với "thiết kế đi từ truy vấn" ở mục NoSQL bên dưới. Nếu thiếu cặp đối lập này, người đọc khó thấy được khác biệt cốt lõi giữa hai cách thiết kế.

### **NoSQL**

- Khuyến khích cách thiết kế đi từ truy vấn (còn gọi là thiết kế theo access pattern), thường gắn với mô hình phát triển "ứng dụng trước" hoặc "API trước": các truy vấn phục vụ chức năng cụ thể của ứng dụng được xác định trước khi xác định mô hình dữ liệu và thực thể.
- Dữ liệu được tổ chức theo mẫu truy cập, thường ở dạng phi chuẩn hóa (denormalization), tức là chấp nhận lưu trùng dữ liệu để một truy vấn chỉ cần đọc ở một nơi.
- Thay vì chuẩn hóa, NoSQL ưu tiên phi chuẩn hóa để tăng tốc độ đọc và giúp phát triển nhanh, linh hoạt hơn khi yêu cầu thay đổi. Đổi lại, khi dữ liệu trùng lặp thay đổi, ứng dụng phải tự đảm bảo các bản sao được cập nhật đồng bộ.
- Các loại cơ sở dữ liệu NoSQL gồm cơ sở dữ liệu hướng tài liệu (document), key-value, column-family và hướng đồ thị (graph). Mỗi loại cung cấp cơ chế lưu trữ và truy vấn khác nhau.

---

## **Các quy tắc khi thiết kế cơ sở dữ liệu NoSQL**

- Thiết kế NoSQL trước hết xuất phát từ cách ứng dụng sẽ truy vấn dữ liệu, và nhấn mạnh các mẫu truy cập hơn các mô hình dữ liệu trừu tượng. Vì vậy, bước đầu tiên của quy trình thiết kế là liệt kê đầy đủ các mẫu truy cập (access pattern) của ứng dụng: truy vấn nào đọc, truy vấn nào ghi, tần suất và độ trễ chấp nhận được. Danh sách này là căn cứ để chọn cấu trúc dữ liệu và chỉ mục.
- Quan hệ giữa các dữ liệu vẫn được xem xét, nhưng dưới góc nhìn của truy vấn: dữ liệu thường được đọc cùng nhau thì nên nhúng (embed) vào cùng một tài liệu, còn dữ liệu thay đổi độc lập hoặc phát triển không giới hạn thì nên tách ra và tham chiếu (reference).
- Thiết kế cũng cần xem xét tần suất dữ liệu thay đổi, tỷ lệ đọc/ghi, lượng dữ liệu cần lưu trữ và các yêu cầu về tính sẵn sàng, hiệu suất và tính nhất quán.
- Thiết kế còn bao gồm việc chọn loại cơ sở dữ liệu NoSQL phù hợp với ứng dụng cụ thể trong số các loại đã nêu ở trên (key-value, document, column-family, graph).

---

## Một số đặc điểm chung

Các đặc điểm dưới đây là xu hướng chung của nhiều hệ thống NoSQL, đặc biệt là các hệ thống phân tán như Cassandra hay DynamoDB. Không phải mọi cơ sở dữ liệu NoSQL đều có đủ các đặc điểm này.

- High Scalability (khả năng mở rộng cao): Mở rộng ngang bằng cách thêm node. Giới hạn thực tế phụ thuộc vào kiến trúc hệ thống và thiết kế khóa phân vùng (ví dụ phân vùng nóng - hot partition làm một node bị quá tải dù cả cụm còn dư tài nguyên).
- High Availability (tính sẵn sàng cao): Dữ liệu được nhân bản (replication) sang nhiều node, nên nếu một node (thường là máy phổ thông - commodity machine) bị lỗi thì toàn bộ hệ thống vẫn tiếp tục phục vụ.
- Atomicity (tính nguyên tử): Phần lớn hệ thống đảm bảo tính nguyên tử ở mức một tài liệu hoặc một dòng. Giao dịch nhiều tài liệu tùy thuộc hệ thống (ví dụ MongoDB hỗ trợ từ phiên bản 4.0, nhiều hệ thống khác hạn chế hoặc không hỗ trợ).
- Consistency (tính nhất quán): Nhiều hệ thống chấp nhận nhất quán cuối cùng (eventual consistency), nghĩa là sau khi cập nhật, các bản sao có thể chưa đồng bộ ngay nên có lúc đọc ra dữ liệu cũ. Nhiều hệ thống cho phép điều chỉnh mức nhất quán theo từng thao tác.
- Durability (độ bền dữ liệu): Dữ liệu thường được ghi vào log trên đĩa (commit log hoặc write-ahead log) trước khi xác nhận, và có thể được giữ thêm trong bộ nhớ để tăng tốc truy cập. Mức đảm bảo cụ thể khác nhau giữa các hệ thống và có thể cấu hình (ví dụ Redis mặc định ghi xuống đĩa theo chu kỳ nên có thể mất dữ liệu gần nhất nếu máy lỗi).
- Deployment Flexibility (linh hoạt triển khai): Với nhiều hệ thống, việc thêm hoặc bớt node được cụm tự động nhận biết và tự phân phối lại dữ liệu, ít cần can thiệp thủ công. Hệ thống không đòi hỏi cấu hình phần cứng mạnh hay đồng nhất.

---

## **5 mẫu thiết kế cơ sở dữ liệu NoSQL thông dụng nhất**

Cơ sở dữ liệu NoSQL ngày càng phổ biến nhờ tính linh hoạt, khả năng mở rộng và hiệu suất khi xử lý các tập dữ liệu lớn và phức tạp. Tuy nhiên, thiết kế một cơ sở dữ liệu NoSQL đòi hỏi cách tiếp cận khác với cơ sở dữ liệu quan hệ truyền thống.

> [Xem chi tiết ở đây](https://claude.ai/chat/1c20f49b-a59f-4c04-9c2d-e1da3e71a72d).

### **Mẫu Key-value**

Một trong những mẫu thiết kế NoSQL đơn giản và phổ biến nhất là key-value. Trong mẫu này, mỗi mục dữ liệu được lưu dưới dạng một cặp khóa-giá trị, trong đó mỗi khóa là duy nhất. Giá trị có thể là bất cứ thứ gì: chuỗi, số, tài liệu hoặc dữ liệu nhị phân.

Mẫu key-value lý tưởng cho các tình huống cần truy cập nhanh và đơn giản đến dữ liệu bằng khóa làm định danh. Hạn chế của mẫu này là cơ sở dữ liệu thường coi giá trị như một khối mờ (opaque), nên khó truy vấn theo nội dung bên trong giá trị.

Ví dụ, có thể dùng mẫu key-value để lưu thiết lập người dùng, dữ liệu phiên (session) hoặc dữ liệu cache. Các hệ thống tiêu biểu: Redis, Amazon DynamoDB.

### **Mẫu Tài liệu (Document)**

Trong mẫu này, mỗi mục dữ liệu được lưu dưới dạng một tài liệu độc lập (thường là JSON hoặc định dạng tương tự), chứa các trường và giá trị mô tả chính nó. Tài liệu thường tự chứa các thông tin liên quan bằng cách nhúng (embed) dữ liệu lồng nhau, nên một lần đọc có thể lấy đủ thông tin mà không cần tham chiếu đến tài liệu khác. Các tài liệu được gom trong một collection, nhưng không bắt buộc phải có cùng cấu trúc, nên cấu trúc có thể thay đổi linh hoạt theo nhu cầu của từng bản ghi.

Mẫu tài liệu phù hợp cho các tình huống cần lưu trữ dữ liệu phức tạp và đa dạng, có thể thay đổi về cấu trúc và kích thước.

Ví dụ, có thể dùng mẫu tài liệu để lưu danh mục sản phẩm, hồ sơ khách hàng hoặc bài đăng trên blog. Các hệ thống tiêu biểu: MongoDB, Couchbase.

### **Mẫu Tổ hợp (Aggregate)**

Mẫu tổ hợp là một biến thể của mẫu tài liệu, tập trung vào việc nhóm các mục dữ liệu liên quan vào một tài liệu duy nhất, tương ứng với một đơn vị nghiệp vụ trọn vẹn.

Mẫu tổ hợp hữu ích khi cần thực hiện thao tác trên một tập dữ liệu mà tất cả các mục thuộc cùng một nhóm. Vì thao tác ghi trên một tài liệu thường có tính nguyên tử, nếu một phần của thao tác không thành công thì toàn bộ thao tác bị hủy, nhờ đó dữ liệu trong nhóm không bị sai lệch hay mất mát. Tính nguyên tử này chỉ áp dụng trong ranh giới của một tài liệu; thao tác trải dài trên nhiều tài liệu không được đảm bảo như vậy ở đa số hệ thống.

Ví dụ, có thể dùng mẫu tổ hợp để lưu đơn hàng kèm các dòng sản phẩm, giỏ hàng hoặc hồ sơ người dùng kèm địa chỉ.

### **Mẫu Bucket**

Mẫu bucket là một biến thể khác của mẫu tài liệu, nhằm giảm số lượng tài liệu trong một collection bằng cách nhóm nhiều mục dữ liệu tương tự vào một tài liệu duy nhất (một "bucket").

Mẫu bucket hữu ích cho dữ liệu được ghi với khối lượng rất lớn và thường được truy vấn theo một khoảng nhất định, điển hình là dữ liệu chuỗi thời gian. Ví dụ, thay vì lưu mỗi lần đo của cảm biến thành một tài liệu riêng, ta gom tất cả các lần đo trong một giờ của cảm biến đó vào một tài liệu. Cách này giúp giảm số lượng tài liệu và kích thước chỉ mục, nhờ đó truy vấn theo khoảng thời gian nhanh hơn.

Ví dụ, có thể dùng mẫu bucket để lưu dữ liệu cảm biến, dữ liệu nhật ký hoặc dữ liệu lịch sử.

### **Mẫu Đồ thị (Graph)**

Mẫu đồ thị là một mẫu thiết kế cơ sở dữ liệu NoSQL chuyên biệt, biểu diễn dữ liệu dưới dạng đồ thị gồm các nút (node) và cạnh (edge). Mẫu này cho phép các truy vấn mạnh mẽ và hiệu quả đi qua nhiều tầng kết nối giữa các mục dữ liệu (ví dụ "bạn của bạn của bạn"), vốn là loại truy vấn cần rất nhiều phép join trong RDBMS. Đổi lại, mẫu này có thể đòi hỏi nhiều tài nguyên và kỹ năng để triển khai và duy trì.

Mẫu đồ thị phù hợp nhất khi cần phân tích các mối quan hệ có sự phụ thuộc phức tạp giữa các mục dữ liệu.

Ví dụ, có thể dùng mẫu đồ thị để lưu mạng xã hội hoặc xây dựng hệ thống gợi ý. Hệ thống tiêu biểu: Neo4j.

---

### **Thiết kế thương mại điện tử NoSQL khác biệt như thế nào?**

NoSQL thường được sử dụng cho phần danh mục sản phẩm của các trang web thương mại điện tử vì tốc độ truy vấn, tính linh hoạt và tính sẵn sàng của nó. Thiết kế NoSQL cho thương mại điện tử khác với thiết kế SQL. Một trang web mua sắm thường bán nhiều loại sản phẩm, và mỗi loại có những thuộc tính riêng để mô tả.

> _Ví dụ, khi một người mua tìm kiếm áo sơ mi, họ sẽ muốn lọc theo kiểu dáng, màu sắc, kích cỡ và có thể là thương hiệu._ _Khi tìm ổ cứng, họ muốn lọc theo kích thước, giao diện và loại ổ đĩa._
> 
> _Thiết kế bằng SQL cho trường hợp này rất khó và tốn kém, vì không thể liệt kê trước mọi thuộc tính của mọi sản phẩm hiện tại và tương lai thành các cột cố định. Có thể dùng các cách như bảng thuộc tính dạng EAV (Entity-Attribute-Value) hoặc cột JSON, nhưng EAV làm truy vấn phức tạp và chậm hơn._
> 
> _Trong khi đó, thiết kế NoSQL được đơn giản hóa vì mỗi sản phẩm không cần có cùng tất cả các trường thuộc tính._

---

## Mở rộng

### 1. Nhúng hay tham chiếu khi mô hình hóa quan hệ

Quyết định này ảnh hưởng trực tiếp đến hiệu năng đọc và chi phí cập nhật.

|Phương án|Ưu điểm|Nhược điểm|
|:--|:--|:--|
|Nhúng (embed) dữ liệu liên quan vào cùng tài liệu|Một lần đọc là đủ, thao tác ghi nguyên tử trong một tài liệu|Tài liệu có thể phình to; dữ liệu trùng lặp phải cập nhật ở nhiều nơi; bị giới hạn kích thước tài liệu|
|Tham chiếu (reference) bằng khóa tới tài liệu khác|Dữ liệu không trùng lặp, mỗi phần cập nhật độc lập|Cần nhiều lần đọc (hoặc phép nối ở tầng ứng dụng), không có đảm bảo nguyên tử giữa các tài liệu ở nhiều hệ thống|
|Kết hợp: nhúng phần ít thay đổi và hay đọc cùng nhau, tham chiếu phần phát triển không giới hạn|Cân bằng giữa tốc độ đọc và kiểm soát kích thước|Thiết kế phức tạp hơn, cần hiểu rõ mẫu truy cập|

### 2. Dữ liệu chuỗi thời gian: mẫu Bucket hay giải pháp chuyên dụng

|Phương án|Ưu điểm|Nhược điểm|
|:--|:--|:--|
|Mẫu Bucket tự cài đặt trên cơ sở dữ liệu tài liệu|Không cần thêm hệ thống, kiểm soát hoàn toàn cấu trúc|Phải tự xử lý logic gom nhóm, cập nhật và truy vấn|
|Tính năng chuỗi thời gian có sẵn của cơ sở dữ liệu (ví dụ time series collection của MongoDB từ phiên bản 5.0)|Cơ sở dữ liệu tự lo việc gom nhóm và nén, ít mã phải viết|Phụ thuộc tính năng và giới hạn của từng sản phẩm|
|Cơ sở dữ liệu chuỗi thời gian chuyên dụng (ví dụ TimescaleDB, InfluxDB)|Tối ưu sẵn cho ghi khối lượng lớn và truy vấn theo thời gian, có sẵn hàm tổng hợp|Thêm một hệ thống phải vận hành|

### 3. Dùng một loại cơ sở dữ liệu hay kết hợp nhiều loại (polyglot persistence)

Ví dụ thương mại điện tử cho thấy các phần khác nhau của một hệ thống có nhu cầu khác nhau.

|Phương án|Ưu điểm|Nhược điểm|
|:--|:--|:--|
|Một RDBMS có hỗ trợ JSON (ví dụ PostgreSQL với kiểu JSONB)|Một hệ thống duy nhất, vẫn có giao dịch ACID và join, đồng thời lưu được thuộc tính linh hoạt|Khả năng mở rộng ngang kém hơn các hệ thống NoSQL phân tán|
|NoSQL cho danh mục sản phẩm, RDBMS cho đơn hàng và thanh toán|Mỗi phần dùng công cụ phù hợp nhất|Phải vận hành nhiều hệ thống và xử lý đồng bộ dữ liệu giữa chúng|
|Chỉ dùng NoSQL cho toàn hệ thống|Kiến trúc thống nhất, mở rộng dễ|Phải tự bảo đảm tính toàn vẹn cho các nghiệp vụ cần giao dịch chặt chẽ|

### 4. Lọc theo nhiều thuộc tính (faceted search) và mô hình column-family

- Lọc sản phẩm theo nhiều thuộc tính trong ví dụ thương mại điện tử thường được giao cho công cụ tìm kiếm chuyên dụng (ví dụ Elasticsearch, OpenSearch) thay vì chỉ dựa vào chỉ mục của cơ sở dữ liệu. Đổi lại phải vận hành thêm một hệ thống và đồng bộ dữ liệu giữa hai nơi.
- Mô hình column-family (wide-column, ví dụ Cassandra, HBase) được nhắc đến trong tài liệu nhưng chưa có mục riêng. Nên bổ sung một mục mô tả mô hình này, vì đây là một trong bốn loại cơ sở dữ liệu NoSQL chính và có cách thiết kế khá khác (thiết kế bảng theo từng truy vấn, khóa phân vùng và khóa sắp xếp).

---

### **[Xem tiếp](https://claude.ai/chat/1c20f49b-a59f-4c04-9c2d-e1da3e71a72d)**
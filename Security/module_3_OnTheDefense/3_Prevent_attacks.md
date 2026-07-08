# **Introduction**

Mối quan tâm hàng đầu của một tổ chức trong an ninh mạng là ngăn chặn một cuộc tấn công thành công xảy ra. Trong mô-đun này, bạn sẽ tìm hiểu một số phương pháp phổ biến mà các tổ chức có thể áp dụng để ngăn chặn các cuộc tấn công.

![](https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcT4nhVNLs9N2h0XojursNd4sEiaHwEJBhnwwAUDJuj-R2SVNr4J8HUTTUU&s=10)

---

# **Goal of attack prevention strategies**

Không tồn tại một chiến lược bảo mật hoàn hảo, bất khả chiến bại. Kẻ tấn công luôn phát triển các kỹ thuật mới để vượt qua các biện pháp bảo mật, khiến các hệ thống phòng thủ hiện tại trở nên lỗi thời.

Một chiến lược bảo mật thực tế và hiệu quả hơn <mark>tập trung vào việc làm cho các cuộc tấn công thành công trở nên khó thực hiện hơn</mark>. Bằng cách đó, các tổ chức có thể ngăn chặn hầu hết các kẻ tấn công, trừ những kẻ quyết tâm nhất. Hãy tưởng tượng rằng việc xâm nhập hệ thống của một tổ chức nào đó sẽ tiêu tốn của kẻ tấn công 100.000 đô la Mỹ nguồn lực. Nếu hệ thống bị xâm nhập chỉ đáng giá 80.000 đô la Mỹ đối với kẻ tấn công, thì kẻ tấn công khó có khả năng tấn công. Do đó, hệ thống phòng thủ có thể hoạt động hiệu quả bất chấp những khiếm khuyết của nó.

Một chiến lược bảo mật thực tế nhưng hiệu quả thường bao gồm việc <mark>chồng lớp phòng thủ, cập nhật và vá lỗi hệ thống thường xuyên, và đầu tư vào thông tin tình báo về mối đe dọa</mark> để luôn đi trước các mối đe dọa mới nổi. Những biện pháp này làm tăng thời gian, chi phí và nỗ lực cần thiết để hoàn thành một cuộc tấn công thành công, từ đó làm giảm đáng kể khả năng xảy ra vi phạm.

Mục tiêu của an ninh mạng là giảm rủi ro hoạt động xuống mức chấp nhận được bằng cách đưa ra sự kết hợp đúng đắn giữa giáo dục, quy trình và công nghệ.

Với mục tiêu đó, hãy cùng xem xét một số chiến lược mà các tổ chức sử dụng để ngăn chặn các cuộc tấn công mạng:

• **Giảm thiểu bề mặt tấn công** 

• **Tạo vùng phi quân sự (DMZ)** 

• **Tuân thủ nguyên tắc quyền hạn tối thiểu** 

• **Quản lý các lỗ hổng phần mềm** 

• **Sử dụng phòng thủ nhiều lớp**



---

## **Reduce the attack surface (Giảm thiểu bề mặt tấn công)**

![](https://datacipher.com/wp-content/uploads/2025/02/7-Essential-Factors-to-Help-You-Choose-the-Right-Attack-Surface-Reduction-Company-visual-selection-1.png)

Một trong những khái niệm đầu tiên cần xem xét là bề mặt tấn công. Bề mặt tấn công là tất cả các điểm trong một hệ thống mà kẻ tấn công có thể cố gắng xâm nhập, tác động hoặc thu thập dữ liệu. Bề mặt tấn công bao gồm tất cả các điểm dễ bị tổn thương, chẳng hạn như giao diện, giao thức và dịch vụ.

Bề mặt tấn công càng lớn, nguy cơ xâm nhập càng cao. Do đó, mục tiêu chính của bất kỳ chiến lược bảo mật tốt nào là giữ cho bề mặt tấn công càng nhỏ càng tốt. Làm như vậy sẽ giảm thiểu các lỗ hổng, khiến hệ thống trở nên kém hấp dẫn hơn và khó bị kẻ tấn công xâm nhập hơn.

> <u>Ví dụ: </u>Hãy xem xét một tổ chức có hệ thống ghi nhận thanh toán. Để giảm bề mặt tấn công của hệ thống, tổ chức này hạn chế quyền truy cập hệ thống của nhân viên tại một số địa điểm văn phòng nhất định. Kết quả là, nhóm an ninh mạng của họ có thể bỏ qua lưu lượng truy cập internet bên ngoài ở vùng biên, làm giảm đáng kể bề mặt tấn công cho những kẻ tấn công tiềm năng. Thay vì có vô số địa chỉ IP để thực hiện một cuộc tấn công, kẻ tấn công trước tiên phải xâm nhập vào một thiết bị đáng tin cậy, và sau đó sử dụng nó cho các cuộc tấn công tiếp theo. Lớp phức tạp bổ sung này làm tăng thách thức cho kẻ tấn công.

Các tổ chức tiếp tục cung cấp cho nhân viên nhiều cách hơn để truy cập vào các hệ thống nội bộ. Ví dụ, nhiều tổ chức cung cấp các phương thức truy cập từ xa và cho phép nhân viên truy cập hệ thống từ thiết bị cá nhân. Những tính năng này đã giúp cải thiện khả năng truy cập, dịch vụ và tính linh hoạt trong công việc. <mark>Nhưng chúng cũng làm tăng diện tích bề mặt tấn công</mark>. Do đó, việc thiết lập một **ranh giới bảo mật** (security perimeter — ranh giới phân định giữa hệ thống nội bộ đáng tin cậy và mạng bên ngoài không đáng tin cậy) trở nên khó khăn hơn. Các tổ chức phải nhận thức được ranh giới bảo mật của mình và chủ động giám sát nó.

---

## **Create a demilitarized zone (DMZ) [Tạo vùng phi quân sự]**

Một phần thiết yếu của thiết kế hệ thống bảo mật là vùng phi quân sự (DMZ). <mark>DMZ là một phân đoạn mạng nằm giữa mạng nội bộ, mạng riêng của tổ chức và internet</mark>. Nó là một vùng đệm bổ sung thêm một lớp bảo mật. Để truy cập vào mạng nội bộ, kẻ tấn công phải đi qua DMZ. Ngay cả khi chúng vượt qua được DMZ, mạng nội bộ vẫn an toàn.

Thông thường, DMZ chứa các máy chủ cần phải truy cập được từ internet rộng hơn nhưng vẫn cần được bảo vệ. Các máy chủ này có thể bao gồm máy chủ web, email, giao thức truyền tệp (FTP) và máy chủ DNS (Domain Name System - hệ thống phân giải tên miền). Các máy chủ này chứa dữ liệu công khai, chứ không phải dữ liệu nội bộ nhạy cảm.

![](https://cpcontents.adobe.com/fr/dynamic-protected/544394deea0445958ee17b7ecbf43393/protected/account/2135/resources/7868768/7868768/content/scormcontent/assets/CyberF_NetworkSegregation.jpg)

Sơ đồ này minh họa hoạt động của vùng DMZ. Người dùng bên ngoài hợp pháp có thể truy cập các máy chủ và ứng dụng nằm trong vùng DMZ, phân đoạn giữa tường lửa bên ngoài và tường lửa bên trong. Tuy nhiên, người dùng này không thể truy cập các máy chủ và ứng dụng nhạy cảm hơn trong mạng nội bộ phía sau tường lửa nội bộ. Ví dụ, khách hàng có thể đặt hàng qua hệ thống thanh toán kỹ thuật số hoặc truy cập email. Tuy nhiên, họ không thể truy cập hồ sơ nhân viên hoặc dữ liệu tài chính của khách hàng khác trên mạng nội bộ.

---

## **Follow the principle of least privilege**

Các tổ chức phải quyết định <mark>mức độ quyền hạn phù hợp cho các ứng dụng và nhân viên</mark>. Họ nên làm điều đó theo nguyên tắc quyền hạn tối thiểu: <mark>cấp cho người dùng hoặc ứng dụng số quyền ít nhất cần thiết để hoàn thành chức năng của chúng.</mark>

> <u>Ví dụ: </u>Một tổ chức thiết lập cơ sở dữ liệu nhân sự (HR) sao cho các nhà quản lý chỉ có quyền truy cập đọc vào dữ liệu cho các vị trí công việc mà họ quản lý. Nếu kẻ tấn công đánh cắp thông tin đăng nhập của một nhà quản lý cụ thể, thì kẻ tấn công chỉ có thể xâm phạm tính bảo mật của những hồ sơ cụ thể đó. Kẻ tấn công không thể sửa đổi chúng vì chúng chỉ có quyền đọc. Ngoài ra, kẻ tấn công không thể truy cập vào các ứng dụng cho các lĩnh vực khác của doanh nghiệp.

Bằng cách đưa ra biện pháp kiểm soát này, tổ chức giảm thiểu hậu quả của một cuộc tấn công thành công so với một hệ thống ít bị hạn chế hơn. Và bằng cách giảm thiểu hậu quả, tổ chức giảm thiểu giá trị rủi ro. Bạn có thể thấy thuật ngữ quân sự "bán kính vụ nổ" (blast radius) được áp dụng trong ngữ cảnh này, trong đó bán kính biểu thị phạm vi ảnh hưởng của một cuộc tấn công. <mark>Giảm quyền hạn là một cách hiệu quả để hạn chế bán kính vụ nổ.</mark>

---

## **Manage software vulnerabilities**

Quản lý lỗ hổng phần mềm bao gồm việc giám sát và giảm thiểu lỗ hổng trong hệ thống phần mềm. Nó bao gồm quản lý bản vá, quy trình áp dụng các bản vá và cập nhật hệ thống khi có bản sửa lỗi.

Khi nhà cung cấp tạo phiên bản phần mềm mới, họ có thể quyết định ngừng hỗ trợ phiên bản cũ hơn (còn gọi là *end-of-life* - EOL). Ví dụ: Microsoft không còn hỗ trợ các phiên bản Windows cũ hơn, chẳng hạn như Windows 7. Kết quả là công ty không còn phát hành các bản vá bảo mật, khiến các thiết bị Windows 7 trở thành mục tiêu dễ dàng cho những kẻ tấn công. Khi có thể, tổ chức nên sử dụng các phiên bản phần mềm mà nhà cung cấp hỗ trợ. Mặt khác, họ nên triển khai các biện pháp kiểm soát bù đắp (compensating controls) để giảm rủi ro liên quan đến các lỗ hổng đã biết. Các biện pháp kiểm soát này có thể bao gồm việc vô hiệu hóa một số tính năng nhất định, tăng cường bảo mật mạng hoặc tăng cường giám sát để phát hiện hoạt động đáng ngờ.

Nhìn chung, việc cập nhật phần mềm và ứng dụng lên phiên bản mới nhất giúp giảm đáng kể nguy cơ bị tấn công thành công. Tuy nhiên, các phiên bản phần mềm mới cũng có thể gây ra lỗ hổng mới. Các tổ chức phải kiểm tra với nhà cung cấp để biết những gì và khi nào có bản vá. Nếu nhà cung cấp không có biện pháp khắc phục lỗ hổng bảo mật, các tổ chức có thể triển khai các biện pháp kiểm soát bù đắp tạm thời, chẳng hạn như vô hiệu hóa một tính năng hoặc quay lại phiên bản trước đó.

Một tổ chức có thể sử dụng trình quét lỗ hổng (vulnerability scanner) để đánh giá phần mềm nào dễ bị tấn công cụ thể. Trình quét lỗ hổng là một ứng dụng quét hệ thống để tìm các lỗ hổng đã biết, chẳng hạn như phần mềm lỗi thời, thiếu bản vá, cài đặt bị định cấu hình sai hoặc mật khẩu yếu. Một số trình quét lỗ hổng hoạt động dựa trên mạng, kiểm tra lỗ hổng bằng cách thử nghiệm tích cực (dynamic scanning). Những trình quét khác quét mã nguồn tĩnh để tìm các lỗi có thể xảy ra (static scanning). Cả hai loại máy quét đều tạo ra thông tin có giá trị để xác định lỗ hổng trước khi kẻ tấn công thực hiện.

---

## **Use defense in depth**

Một cân nhắc quan trọng khác đối với việc phòng thủ là các tổ chức phải sử dụng cách tiếp cận theo lớp. Thuật ngữ phòng thủ theo chiều sâu bắt nguồn từ quân đội; <mark>thay vì dựa vào một hình thức phòng thủ duy nhất, bạn xếp lớp chúng</mark>. Ý nghĩa tương tự trong CNTT. Phòng thủ theo chiều sâu là chiến lược trong đó bạn sử dụng nhiều lớp kiểm soát để bảo vệ tài sản.

![](https://cpcontents.adobe.com/fr/dynamic-protected/544394deea0445958ee17b7ecbf43393/protected/account/2135/resources/7868768/7868768/content/scormcontent/assets/CyberF_DefenseInDepth.png)

Ví dụ: một tổ chức có thể áp dụng các lớp bảo mật sau:

• Phòng thủ mạng như tường lửa

• Biện pháp bảo vệ thiết bị như máy quét phần mềm độc hại

• Bảo vệ dữ liệu như mã hóa

Để một cuộc tấn công thành công, nó phải thỏa hiệp hoặc phá vỡ tất cả các lớp này, đây là một thách thức đáng kể đối với kẻ tấn công.

---

## **Knowledge check**

Câu hỏi 1: Augustyna là trưởng bộ phận CNTT của công ty cô ấy và cô ấy đang giúp thiết lập cơ sở dữ liệu Nhân sự (HR) của công ty mình. Để bảo vệ dữ liệu, cô dự định áp dụng nguyên tắc đặc quyền tối thiểu.

Augustyna nên thực hiện hành động nào sau đây để áp dụng nguyên tắc đặc quyền tối thiểu?

- Cấp cho tất cả nhân viên quyền truy cập đọc và ghi vào tất cả dữ liệu của công ty.
- Cấp cho nhân viên quyền truy cập đọc và ghi chỉ vào dữ liệu mà họ cần. **[Đáp án đúng]**
- Cấp cho người quản lý quyền chỉ đọc dữ liệu đối với các vai trò công việc mà họ quản lý.
- Chỉ cấp cho bộ phận nhân sự quyền truy cập đọc tất cả dữ liệu của công ty.

Giải thích: Nguyên tắc đặc quyền tối thiểu có nghĩa là bạn nên cấp cho mỗi người dùng số quyền ít nhất cần thiết để hoàn thành công việc của họ. Trong trường hợp này, Augustyna chỉ nên cấp cho mỗi nhân viên quyền truy cập đọc và ghi đối với dữ liệu họ cần, chẳng hạn như dữ liệu liên quan đến bộ phận của họ.

Câu hỏi 2: Roger là quản trị viên hệ thống tại tổ chức của anh ấy. Anh ta xác định một lỗ hổng tiềm ẩn trong một trong những hệ thống phần mềm quan trọng của tổ chức. Anh ta liên hệ với nhà cung cấp phần mềm và họ thông báo rằng họ không còn hỗ trợ phần mềm đó nữa.

Roger nên làm gì trong tình huống này?

- Bỏ qua các lỗ hổng tiềm ẩn.
- Thực hiện các biện pháp kiểm soát bồi thường. **[Đáp án đúng]**
- Cấp cho người dùng quyền truy cập đầy đủ vào phần mềm.
- Tắt hoàn toàn hệ thống.

Giải thích: Nếu có thể, Roger nên khuyến khích tổ chức của mình nâng cấp lên các phiên bản phần mềm mà nhà cung cấp hỗ trợ. Mặt khác, anh ta nên thực hiện các biện pháp kiểm soát bù đắp để giảm rủi ro liên quan đến các lỗ hổng đã biết. Ví dụ: anh ta có thể tăng cường giám sát để phát hiện hoạt động đáng ngờ.

---

## **Bổ sung kiến thức**

Các điểm dưới đây không thay đổi nội dung gốc, chỉ bổ sung góc nhìn hiện đại/thực chiến để tài liệu hữu ích hơn khi áp dụng vào công việc backend thực tế:

1. 

2. **Zero Trust như một bước phát triển của DMZ/perimeter security**: Mô hình DMZ dựa trên giả định "trong mạng nội bộ = tin cậy, ngoài internet = không tin cậy". Các kiến trúc hiện đại (đặc biệt trong môi trường cloud, microservice) đang chuyển sang **Zero Trust** — không mặc định tin cậy bất kỳ request nào dù đến từ trong hay ngoài mạng, luôn xác thực và kiểm tra quyền hạn ở từng lớp (identity-aware proxy, mTLS giữa các service). Đây là kiến thức nên biết khi làm backend với Kubernetes/service mesh.

3. **Liên hệ với công việc Golang backend**: Nguyên tắc least privilege áp dụng trực tiếp khi thiết kế RBAC/ABAC cho API, phân quyền theo role trong JWT claims, hoặc khi cấp IAM role cho service trên cloud (chỉ cấp đúng permission cần thiết cho service account, không dùng quyền admin toàn cục).

4. **CVE và CVSS**: Khi nói về "lỗ hổng đã biết" trong phần quản lý lỗ hổng phần mềm, nên biết hai khái niệm liên quan: **CVE** (Common Vulnerabilities and Exposures — mã định danh chuẩn cho từng lỗ hổng đã công bố) và **CVSS** (Common Vulnerability Scoring System — thang điểm đánh giá mức độ nghiêm trọng). Các công cụ như `govulncheck` (Go), `trivy`, `snyk` dùng dữ liệu CVE để quét lỗ hổng trong dependency của dự án Go.

5. **Defense in depth trong kiến trúc microservice**: Ngoài 3 lớp ví dụ trong bài (network, endpoint, data), một backend engineer nên biết thêm các lớp phổ biến trong thực tế: rate limiting/WAF ở API Gateway, input validation ở tầng ứng dụng, secrets management (Vault) thay vì hardcode credential, và audit logging để phát hiện bất thường sau khi các lớp khác đã bị vượt qua.

# Debug Workflow: Quy trình gỡ lỗi phần mềm hiệu quả

## Mục lục

1. [Bug là gì?](#bug-là-gì)
2. [Debug là gì?](#debug-là-gì)
3. [Debugger là gì và vì sao nên dùng thay vì print?](#debugger-là-gì-và-vì-sao-nên-dùng-thay-vì-print)
4. [Quy trình chuẩn để Debug](#quy-trình-chuẩn-để-debug)
5. [Các thao tác cơ bản khi dùng Debugger](#các-thao-tác-cơ-bản-khi-dùng-debugger)
6. [Các chiến lược Debugging phổ biến](#các-chiến-lược-debugging-phổ-biến)
7. [Kết luận](#kết-luận)
8. [Mở rộng](#mở-rộng)

---

## Bug là gì?

Bug là lỗi phát sinh trong mã nguồn của một chương trình cụ thể, khiến chương trình hoạt động không đúng như thiết kế. Bug thường được phát hiện sau khi sản phẩm phát hành hoặc trong quá trình thử nghiệm; khi đó, người dùng phải tránh sử dụng tính năng bị lỗi cho đến khi nhận được bản vá từ đội ngũ phát triển.

Cần lưu ý bug chỉ là một loại vấn đề trong nhiều loại vấn đề mà một chương trình có thể gặp phải. Một chương trình có thể chạy không phát sinh bug nhưng vẫn khó sử dụng hoặc không đáp ứng đúng mục tiêu thiết kế ban đầu (vấn đề về usability hoặc thiết kế nghiệp vụ) - loại vấn đề này thường khó phát hiện qua kiểm thử tự động hơn nhiều so với bug thông thường. Một chương trình được phát triển với quy trình kiểm soát chất lượng tốt thường có mật độ lỗi trên mỗi nghìn dòng mã (bug density) thấp hơn đáng kể. Đây là lý do vì sao Testing và Debugging luôn là hai hoạt động quan trọng xuyên suốt vòng đời phát triển phần mềm.

## Debug là gì?

Khi mới học lập trình, hầu hết mọi người đều từng dùng câu lệnh in giá trị biến ra màn hình để kiểm tra. Ví dụ với Go:

```go
sum := 10 + 5
fmt.Println(sum) // In giá trị biến sum ra console
```

Cách này đơn giản, dễ làm quen, nhưng chỉ phù hợp khi mới tiếp xúc với lập trình hoặc cần kiểm tra nhanh một giá trị đơn lẻ. Khi làm việc với hệ thống phức tạp hơn, cách tiếp cận hiệu quả hơn nhiều là sử dụng **debugger** - công cụ chuyên dụng để theo dõi và phân tích trạng thái chương trình khi đang chạy.

## Debugger là gì và vì sao nên dùng thay vì print?

Hầu hết IDE hiện đại và nhiều text editor đều tích hợp sẵn debugger. Đây là công cụ hữu ích để tìm và sửa lỗi trong chương trình, với nhiều khả năng vượt xa việc chỉ in giá trị biến:

- Theo dõi luồng thực thi của chương trình - nếu chương trình dừng đột ngột, debugger giúp xác định chính xác dòng code gây ra sự cố.
- Xem giá trị của biến, kể cả các object phức tạp, một cách trực quan mà không cần thêm/xóa các dòng `Println` thủ công.
- Xem log, call stack và trạng thái bộ nhớ tại một thời điểm cụ thể trong quá trình chạy.

Debugger có thể giúp phát hiện nhiều loại lỗi khác nhau:

- Lỗi cú pháp (syntax error).
- Lỗi đánh máy (typo).
- Lỗi logic (chương trình chạy nhưng cho kết quả sai).
- Lỗi triển khai (implementation error - hiện thực sai so với thiết kế ban đầu).

## Quy trình chuẩn để Debug

Về bản chất, debug là quá trình: bắt đầu từ một sự cố quan sát được, cô lập phần mã nguồn gây ra sự cố, sau đó khắc phục và xác minh lại. Debugger hỗ trợ tái tạo lại chính xác điều kiện mà bug xảy ra, kiểm tra trạng thái chương trình tại thời điểm đó để xác định nguyên nhân gốc rễ.

Dù công cụ debugger giữa các IDE có khác nhau đôi chút, quy trình debug hiệu quả thường gồm 5 bước cốt lõi sau:

1. **Tái hiện lỗi**: chạy lại chương trình với dữ liệu đầu vào tương tự thời điểm xảy ra lỗi, nhằm xác định chính xác bước nào khiến chương trình dừng hoặc trả về kết quả sai. Nếu không tái hiện được lỗi một cách ổn định, các bước tiếp theo sẽ rất khó thực hiện.
2. **Đọc log và thông báo lỗi**: xem lại execution history hoặc bật ghi log chi tiết để tìm thông báo lỗi gốc (stack trace, error message) - đây thường là manh mối quan trọng nhất để khoanh vùng vấn đề.
3. **Đặt breakpoint và kiểm tra từng bước**: chèn điểm dừng tại các vị trí nghi ngờ để kiểm tra giá trị biến ngay tại thời điểm chương trình chạy đến đó.
4. **Cô lập nguyên nhân**: kiểm tra dữ liệu đầu vào và đầu ra ở từng bước xử lý (từng hàm, từng lớp logic) để xác định chính xác công đoạn nào khiến dữ liệu bị sai lệch.
5. **Sửa lỗi và kiểm chứng**: cập nhật lại logic, sau đó chạy lại nhiều kịch bản dữ liệu khác nhau (bao gồm cả các edge case) để đảm bảo lỗi đã được khắc phục triệt để và không phát sinh lỗi mới (regression).

## Các thao tác cơ bản khi dùng Debugger

### Breakpoint

Breakpoint là điểm đánh dấu trên một dòng code; khi chương trình chạy đến dòng có breakpoint, quá trình thực thi sẽ tạm dừng. Tại đây, có thể xem giá trị các biến, xem log, hoặc chọn đi tiếp từng dòng, hoặc cho chương trình chạy tiếp bình thường.

Cách đặt breakpoint: click chuột vào lề trái của dòng code cần dừng. Click thêm lần nữa để bỏ breakpoint.

![Cách đặt breakpoint trong IDE](https://images.viblo.asia/b7d25ddc-a7d8-4d9d-9bd0-5ee6ae8a3529.png)

Breakpoint thường hiển thị dưới dạng chấm tròn màu đỏ ở lề trái; hình dạng có thể thay đổi đôi chút tùy trạng thái (đã được đi qua hay chưa). Khi chương trình tạm dừng tại breakpoint, dòng code đó sẽ được highlight.

**Lưu ý quan trọng**: dòng code đang được highlight là dòng **chưa** thực thi - nó chỉ được thực thi khi tiếp tục đi qua dòng tiếp theo.

### Local Variables và Watch

Khi chương trình dừng tại breakpoint, có hai công cụ chính để xem giá trị biến:

- **Local Variables**: hiển thị các biến cục bộ (local) trong hàm hiện tại, tự động cập nhật danh sách khi chuyển sang hàm khác.
- **Watch**: dùng để theo dõi biến toàn cục (global) hoặc bất kỳ biến/biểu thức nào cần quan sát; giá trị trong Watch được cập nhật liên tục kể cả khi chương trình đang dừng. Khác với Local Variables, cần thêm biến vào Watch thủ công (thường bằng cách chuột phải vào biến và chọn "Add to watches").

![Xem giá trị biến qua Local Variables và Watch](https://images.viblo.asia/8d8c43ea-c519-44cf-acb0-64bc3cc60d51.png)

### Step Over, Step Into, Step Out và các thao tác điều khiển khác

Khi chương trình đang tạm dừng tại breakpoint, có thể điều khiển tiếp bằng các thao tác sau:

- **Step Over**: thực thi dòng code hiện tại rồi dừng lại ở dòng tiếp theo, không đi sâu vào bên trong các hàm được gọi.
- **Step Into**: nếu dòng hiện tại có lời gọi hàm, thao tác này sẽ nhảy vào bên trong hàm đó để tiếp tục debug từng bước.
- **Step Out**: thoát khỏi hàm hiện tại, quay trở lại nơi hàm được gọi.
- **Run to Cursor**: cho chương trình chạy tiếp cho đến khi đến đúng vị trí con trỏ đang đặt.
- **Resume**: cho chương trình chạy tiếp bình thường, không dừng theo từng dòng nữa.
- **Stop**: dừng hẳn chương trình đang debug.

![Các thao tác Step trong debugger](https://images.viblo.asia/57fa3b13-fc80-44c7-8506-7c7f76fc4eac.jpg)

### Các công cụ hỗ trợ khác

- **Expression Evaluation**: cho phép nhập và tính toán một biểu thức bất kỳ ngay tại thời điểm chương trình đang dừng, kể cả gọi hàm hoặc truy cập biến hiện có trong phạm vi (scope) hiện tại.
- **Console**: nơi hiển thị log được in ra trong quá trình chạy chương trình.
- **Call Stack**: hiển thị danh sách các hàm đã được gọi theo thứ tự, giúp xác định nhanh hàm nào được gọi cuối cùng trước khi xảy ra lỗi - từ đó đặt breakpoint đúng vị trí cần kiểm tra.

## Các chiến lược Debugging phổ biến

Dựa trên quy trình chung ở trên, có một số chiến lược tiếp cận khác nhau khi debug, tùy vào tính chất của lỗi:

- **Backtracking**: bắt đầu từ dòng code nơi phát hiện triệu chứng của bug, sau đó lần ngược lại theo luồng thực thi để tìm ra nguyên nhân gốc rễ. Phù hợp khi lỗi có luồng thực thi rõ ràng, không quá phức tạp.
- **Phương pháp loại bỏ nguyên nhân (Cause Elimination)**: liệt kê danh sách các nguyên nhân tiềm ẩn có thể gây ra lỗi, sau đó lần lượt kiểm tra và loại trừ từng nguyên nhân cho đến khi tìm ra nguyên nhân thực sự. Phù hợp khi có nhiều giả thuyết về nguyên nhân gây lỗi.
- **Program Slicing**: chỉ tập trung vào một "lát cắt" nhỏ của chương trình - tức tập hợp các câu lệnh thực sự ảnh hưởng đến giá trị của một biến hoặc kết quả cụ thể đang bị nghi ngờ - thay vì phải xem xét toàn bộ mã nguồn. Cách này giúp thu hẹp phạm vi điều tra đáng kể trong các hệ thống lớn.
- **Shotgun Debugging**: cách tiếp cận thử - sai (trial and error), sửa nhiều chỗ nghi ngờ cùng lúc dựa trên phỏng đoán rồi kiểm tra xem lỗi có biến mất hay không. Đây là chiến lược **kém hệ thống nhất** trong bốn cách trên, dễ khiến người debug vô tình sửa "trúng" lỗi mà không thực sự hiểu nguyên nhân gốc rễ, hoặc tệ hơn là che giấu bug thay vì khắc phục triệt để. Nên chỉ dùng như bước thăm dò ban đầu, không nên dừng lại ở đây nếu chưa xác minh rõ nguyên nhân.

## Kết luận

Debug không đơn thuần là việc "tìm và sửa lỗi", mà là một quy trình có phương pháp: tái hiện lỗi, thu thập thông tin, thu hẹp phạm vi nghi vấn, xác định nguyên nhân gốc rễ rồi mới sửa và kiểm chứng lại. Việc thành thạo debugger - thay vì chỉ dựa vào việc in giá trị biến ra console - giúp tiết kiệm đáng kể thời gian khi làm việc với hệ thống phức tạp, đặc biệt trong môi trường backend nơi luồng dữ liệu thường đi qua nhiều lớp xử lý khác nhau.

### Mở rộng

Một số hướng tìm hiểu thêm để nâng cao kỹ năng debug:

- **Delve (dlv)**: debugger chuyên dụng cho Go, hỗ trợ đặt breakpoint, step qua goroutine, kiểm tra biến ngay trong terminal hoặc tích hợp với VS Code/GoLand - công cụ nên thành thạo khi phát triển backend bằng Golang.
- **Remote Debugging**: kỹ thuật debug một chương trình đang chạy trên server hoặc container từ xa (ví dụ dùng Delve ở chế độ `headless` kết hợp kết nối từ IDE), rất hữu ích khi lỗi chỉ xuất hiện trong môi trường staging/production mà không tái hiện được ở local.
- **Structured Logging**: khi hệ thống chạy ở production, không phải lúc nào cũng gắn được debugger trực tiếp; kỹ thuật ghi log có cấu trúc (structured logging, ví dụ với Zap trong Go) giúp thay thế một phần vai trò của debugger để truy vết lỗi sau khi sự cố đã xảy ra.
- **Distributed Tracing**: trong kiến trúc microservices, một request có thể đi qua nhiều service khác nhau; các công cụ tracing (OpenTelemetry, Jaeger) giúp theo dõi toàn bộ hành trình của request, tương tự như "debug" ở cấp độ toàn hệ thống thay vì một chương trình đơn lẻ.
- **Debug test tự động (unit test) thay vì debug thủ công**: viết test case tái hiện lại chính xác điều kiện gây lỗi giúp việc debug lặp lại nhanh hơn nhiều so với thao tác thủ công qua UI mỗi lần kiểm tra.
- **Panic và Recover trong Go**: tìm hiểu cách Go xử lý lỗi nghiêm trọng (panic) và cách dùng `recover` đúng cách, tránh việc dùng `recover` như một công cụ "che giấu" lỗi thay vì debug triệt để.
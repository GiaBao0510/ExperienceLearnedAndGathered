# **Debug workflow**

## **Bug là gì?**
Bug là một lỗi mã hóa trong một chương trình cụ thể. Bug thường được phát hiện sau khi sản phẩm được phát hành hoặc trong quá trình thử nghiệm công khai. Khi điều này xảy ra, người dùng phải tìm cách tránh sử dụng mã lỗi hoặc nhận bản vá từ các nhà phát triển phần mềm.

Bug chỉ là một loại vấn đề mà một chương trình có thể mắc phải. Các chương trình có thể chạy không có Bug và vẫn khó sử dụng hoặc bị lỗi trong một số mục tiêu chính. Loại lỗi này khó kiểm tra hơn Bug rất nhiều. Một chương trình được thiết kế tốt sẽ được phát triển với quy trình được kiểm soát dẫn đến ít lỗi hơn trên mỗi nghìn dòng mã. Đây là lý do tại sao quá trình Testing và Debugging là vô cùng quan trọng.

## **Debug là gì?**

Nhớ lại lúc mới học lập trình, hẳn ai cũng đã từng dùng các câu lệnh print giá trị của biến ra màn hình để xem. Ví dụ như đoạn code Java sau.

```
sum := 10 + 5;
fmt.Println(sum);  // Xem giá trị biến sum
```

Cách này khá đơn giản và dễ làm quen. Tuy nhiên nó chỉ hợp với các bạn newbie, mới lần đầu tiếp xúc với code. Nếu bạn đã code được một thời gian, hãy tập sử dụng một công cụ khác có tên là **debugger**.

Hầu hết các IDE và một số text editor đã có debugger. Debugger là một tool siêu hữu ích để thực hiện việc debug - tìm và fix các lỗi trong chương trình. Ngoài ra mục đích của debugger còn nhiều hơn thế:

- Theo dõi luồng chạy chương trình (nếu chương trình bị stop đột ngột thì debug có thể biết được dòng nào bị lỗi)
- Xem giá trị các biến, object phức tạp dễ dàng
- Xem các log được in ra console, call stack,...

Tại đây, các Debugger sẽ tìm kiếm:

- Lỗi cú pháp
- Lỗi đánh máy
- Lỗi trong logic
- Lỗi triển khai

---
## **Quy trình chuẩn để Debug?**

Để Debug một chương trình, người dùng phải bắt đầu với một sự cố, cô lập mã nguồn của sự cố và sau đó khắc phục nó. Các công cụ Debug (được gọi là trình gỡ lỗi) được sử dụng để xác định các Bug mã hóa ở các giai đoạn phát triển khác nhau. Chúng được sử dụng để tái tạo các điều kiện mà Bug đã xảy ra, sau đó kiểm tra trạng thái chương trình tại thời điểm đó và xác định nguyên nhân. 

Debugger có thể theo dõi từng bước việc thực thi chương trình bằng cách đánh giá giá trị của các biến và dừng việc thực thi ở bất cứ nơi nào được yêu cầu để lấy giá trị của các biến hoặc đặt lại các biến chương trình. Một số gói ngôn ngữ lập trình cung cấp trình Debug để kiểm tra mã lỗi trong khi mã đang được viết tại thời điểm chạy.

Các IDE khác nhau sẽ có `debugger` khác nhau, nên sẽ có một tí khác biệt. Nhưng chung quy lại thì debug chỉ gồm một số bước sau:

1. Đặt `breakpoint` ở các dòng cần tạm dừng để debug
2. Chạy chương trình với debug mode
3. Thao tác với chương trình, sao cho chạy tới chỗ đặt breakpoint
4. Khi IDE dừng tại breakpoint, thì thực hiện xem value các biến, xem log,... để kiểm tra bug.
5. Sau đó đi tiếp từng dòng code tiếp theo, xem sự thay đổi các biến sau từng câu lệnh
6. Tiếp tục chạy chương trình bình thường, hoặc dừng chương trình.

Quy trình debug workflow gồm 5 bước cốt lõi: Tái hiện lỗi, phân tích thông báo, kiểm tra từng bước bằng điểm dừng (breakpoint), cô lập nguyên nhân, và kiểm tra lại toàn bộ luồng tự động hóa.

Các bước thực hiện
Nhận diện và tái hiện lỗi: Chạy lại luồng với dữ liệu đầu vào tương tự để xác định chính xác thời điểm workflow bị dừng hoặc trả về kết quả sai.Đọc log và mã lỗi: Xem lịch sử chạy (execution history) hoặc bật tính năng ghi log chi tiết để tìm thông báo lỗi gốc.
Đặt điểm dừng (Breakpoint): Chèn các điểm dừng tại các bước trung gian nghi ngờ để kiểm tra giá trị của biến số ngay tại thời điểm đó.
Khoanh vùng nguyên nhân: Kiểm tra dữ liệu đầu vào và đầu ra ở từng node/nước chuyển giao để xem dữ liệu bị lệch ở khâu nào.
Sửa lỗi và kiểm chứng: Cập nhật lại logic, chạy thử nghiệm (test run) nhiều lần với các kịch bản dữ liệu khác nhau để đảm bảo ổn định

#### **2.1 Breakpoint**

Là thứ được đánh dấu lên dòng code, khi chương trình chạy tới dòng có breakpoint thì sẽ bị tạm dừng. Lúc này bạn có thể dùng các tool của debugger để xem giá trị các biến, xem log,... Hoặc bạn có thể đi tiếp từng dòng tiếp theo, hoặc cho chương trình chạy bình thường (không bị tạm dừng nữa).

Cách đặt breakpoint: Click chuột vào lề trái của dòng đó. Lúc này breakpoint sẽ được toggle, click thêm lần nữa để xóa (như hình).

![](https://images.viblo.asia/b7d25ddc-a7d8-4d9d-9bd0-5ee6ae8a3529.png)

Breakpoint thường là có dạng hình tròn màu đỏ, dấu tròn này có thể khác đôi chút để biểu thị trạng thái breakpoint (đã được đi qua hay chưa). Khi chương trình tạm dừng tại breakpoint thì dòng đó sẽ được highlight.

Lưu ý: Dòng code được highlight sẽ chưa thực hiện, chỉ khi bạn đi tiếp qua dòng khác thì dòng trước đó mới được thực thi.

#### **2.2. Local variables & watches**

Khi chương trình dừng tại `breakpoint`, chúng ta có thể dùng 2 tool là `Local variables` và `watches` để xem giá trị các biến:

- `Local variables` chỉ để xem các biến trong function hiện tại (local). Khi qua function khác thì danh sách biến sẽ được update lại.
- `Watch` dùng để xem các biến `global`, hoặc bất cứ biến nào. Value của watch được theo dõi ngay cả khi chương trình dừng.

Local variables thì sẽ tự động update danh sách biến trong function. Tuy nhiên, với watch thì bạn phải tự thêm các biến vào thủ công (chuột phải vào biến, Add to watches).

![](https://images.viblo.asia/8d8c43ea-c519-44cf-acb0-64bc3cc60d51.png)

#### **2.3. Step over, các step khác, stop/resume program**
Khi IDE tạm dừng tại `breakpoint`, thì muốn tiếp tục đi qua các dòng tiếp theo cần thực hiện:

- Step over: đi tiếp dòng code tiếp theo (di chuyển xuống một dòng, sau vị trí đặt breakpoints)
- Step into: nhảy vào bên trong hàm (dòng code hiện tại chứa lời gọi hàm)
- Step out: từ trong hàm nhảy ra ngoài, trở về nơi gọi hàm
- Run to cursor: chạy tiếp tục cho tới dòng có con trỏ
- Watch: là nơi mà hiển thị các thông tin của biến, giá trị trả về của function, để chúng ta xem và phán đoán bug, đồng thời có thể đưa ra được cách giải quyết cái bug của mình.

![](https://images.viblo.asia/57fa3b13-fc80-44c7-8506-7c7f76fc4eac.jpg)

Ngoài ra, đôi lúc bạn sẽ muốn tiếp tục chạy chương trình bình thường (không phải dùng Step over từng dòng nữa). Lúc này bạn dùng command Resume để tiếp tục chương trình, hoặc stop để dừng lại.

#### **2.4. Các tool khác**
Debugger cũng có các tool khác để hỗ trợ như:

- **Expression (evalution)**: chương trình tính toán biểu thức, đặc biệt có thể nhập bất kì tên biến nào vào tính cũng được 👏
- **Console**: quá quen thuộc rồi, đây là nơi in ra các log
- **Call stack**: stack chứa danh sách các hàm, lệnh đã được gọi. Chúng ta có thể biết được function nào được gọi cuối cùng, để tìm ra và đặt breakpoint ở đó nhanh chóng

---
### **Các chiến lược Debugging phổ biến**
Dựa trên quy trình chung đã đề cập ở trên, có rất nhiều chiến lược Debugging khác nhau. Cụ thể:

- `Backtracking` – trình săn lỗi bắt đầu từ câu lệnh mà tại đó một triệu chứng Bug đã được phát hiện và theo dõi mã nguồn ngược về lỗi thực tế.
- `Phương pháp Loại bỏ Nguyên nhân` – Debugger tạo ra một danh sách các nguyên nhân tiềm ẩn gây ra lỗi và chạy các bài kiểm tra để xác định nguồn gốc của Bug.
`Program Slicing` – người kiểm tra đảm bảo chất lượng (QA) chạy một nhóm các câu lệnh chương trình trong chương trình (lát cắt) bao gồm các điều kiện cụ thể.
- `Shotgun Debugging` – một cách tiếp cận thử-và-sai để gỡ lỗi dựa trên các dự đoán có tính logic của nhà phát triển.

**1. Sequential (tuần tự)**
- Mô tả: Một lõi CPU xử lý từng tác vụ một cách tuần tự, ví dụ Task 1 xong rồi mới đến Task 2.
- Chạy một chương trình rồi chờ xong, mới chạy chương trình khác.
![](https://media.beehiiv.com/cdn-cgi/image/fit=scale-down,format=auto,onerror=redirect,quality=80/uploads/asset/file/ac4b1636-2b73-4b14-a51b-d92c0031c096/image.png?t=1720731877)

**2. Concurrent (Đồng thời)**
- Mô tả: Một lõi CPU xử lý các tác vụ đan xen nhau (ví dụ xử lý một phần của Task 1, rồi chuyển sang xử lý một phần của Task 2, rồi quay lại Task 1,...) nên có tính đồng thời
- Các tác vụ có thể hoàn thành gần cùng lúc, dù vẫn chỉ chạy trên một lõi CPU.
- Đa nhiệm trên một lõi CPU (multitasking sử dụng kỹ thuật như time slicing).
![](https://media.brightdata.com/2024/06/Example-of-3-tasks-running-concurrently-1.png)

**3. Parallel (song song)**
- Mô tả: Hai lõi CPU, mỗi lõi xử lý một tác vụ riêng biệt từ đầu đến cuối (Các tác vụ này thường là các tác vụ nặng). Các tác vụ này không đan xen, nhưng được xử lý cùng lúc.
- Đây là tính song song vì mỗi tác vụ chỉ dùng một luồng và không bị ngắt quãng.
- Hai tiến trình độc lập chạy trên hai lõi khác nhau.
![](https://jenkov.com/images/java-concurrency/concurrency-vs-parallelism-2.png)
    

**4. Concurrent, Parallel (Đồng thời và song song)**
- Mô tả: Nhiều lõi CPU xử lý các tác vụ vừa song song nhau vừa đồng thời, tức là xử lý nhiều tác vụ cùng lúc và mỗi tác vụ có thể được chia nhỏ để chạy đồng thời.
- Đây là hình thức mạnh nhất, tận dụng cả đa lõi CPU và kỹ thuật xử lý đồng thời.
- Hệ thống đa luồng (multithreading) trên nhiều lõi CPU.
![](https://jenkov.com/images/java-concurrency/concurrency-vs-parallelism-3.png)
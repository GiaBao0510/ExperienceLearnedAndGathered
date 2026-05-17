![](https://www.albahari.com/threading/NewThread.png)
### **Thread trong C# là gì?**

**Thread** là một đường thực thi (execution path) của một chương trình. Mỗi **Thread** được định nghĩa một dòng điều kiển duy nhất. Một chương trình có thể có nhiều thread chạy song song và mỗi **Thread** thực hiện một công việc cụ thể.

Lớp `System.Threading.Thread` trong c# được sử dụng để làm việc với các **Thread**. Nó cho phép tạo và truy cập các **Thread** riêng lẻ của một ứng dụng đa luồng. **Thread** đầu tiên được thực thi trong một tiến trình được gọi là **Thread** chính và nó tự động thực thi khi chương trình c# bắt đầu. Các **Thread** được tạo bởi lớp `Thread` được gọi là **Thread** con của **Thread** chính.

Theo ví dụ dưới đây, bạn có thể truy cập một **Thread** bằng các sử dụng thuộc tính `CurrentThread` của lớp `Thread`.

```csharp
public class Solution
{
    public static void Main(string[] args)
    {
        Console.WriteLine("Multiple Thread in C#:");
        Thread thread = Thread.CurrentThread;
        thread.Name = "Main Thread";
        Console.WriteLine($"Thread Name: {thread.Name}");
        Console.WriteLine($"Thread ID: {thread.ManagedThreadId}");
    }
}
```

- Kết quả sau khi chạy:
```powershell
PS D:\HocTap\Practice\LeetCode\test\TestLeetCode> dotnet run
Multiple Thread in C#:
Thread Name: Main Thread
Thread ID: 1
```

---
### **Vòng đời của Thread trong C#**

Vòng đời của `Thread` ==được bắt đầu khi một đối tượng của lớp `Thread` được tạo và kết thúc==. Khi **Thread** đó được kết thúc hoặc hoàn thành thực thi.  Dưới đây là một số trạng thái vòng đời:
![](https://www.simplilearn.com/ice9/free_resources_article_thumb/C%23-Threads-Article/C%23-Threads-Lifecycle-img1.png)

- **Unstarted State:** Là trạng thái khi instance của **Thread** được tạo, nhưng phương thức `Start` chưa được gọi.
- **Ready State:** Là trạng thái của **Thread** đó sẳn sàng để chạy và đợi CPU cycle.
- **Not Runnable State:** là trạng thái mà **Thread** không thể thực thi. Khi:
	- Phương thức `sleep` đã được gọi.
	- Phương thức `Wait` đã được gọi.
	- Bị ngăn chặn bởi hoạt động I/O.
- **Dead State**: Là trạng thái khi **Thread** hoàn thành sự thực thi hoặc bị hủy bỏ.

![](https://www.dotnetinterviewquestions.in/contentpics/d10.jpg)

---
#### **Tạo Thread trong C#:**

Để tạo một **Thread** mới, chỉ cần khởi tạo đối tượng từ lớp `Thread` và truyền cho nó một tham số là một đối tượng từ lớp `ThreadStart` hoặc `ParameterizedThreadStart`. 
- Lớp `ThreadStart` là một delegate (ủy quyền) không có tham số.
- Lớp `ParameterizedThreadStart` là một delegate có một tham số kiểu object. 

Các **delegate** này sẽ chỉ định phương thức mà `Thread` sẽ được thực thi khi khởi động. 

Sau khi tạo **Thread**, bạn chỉ cần gọi phương thức `Start` để bắt đầu thực thi `Thread`. Bạn cũng có thể gọi phương thức `Start` với một tham số kiểu `object` nếu Thread được tạo bởi delegate `ParameterizedThreadStart`.

Sau đây là ví dụ minh họa cho việc tạo và khởi động một **Thread** trong C#.
```csharp
public class Solution
{
    public static void Main(string[] args)
    {
        Console.WriteLine("Multiple Thread in C#:");
        Console.WriteLine("--------------------");

        Thread thread = new Thread(new ThreadStart(DemoThread));
        thread.Start();

        Console.WriteLine($"Main thread continues running");
    }

    public static void DemoThread()
    {

        Console.WriteLine("Thread Start:");

        for (int i = 0; i < 10; i++)
        {
            Console.WriteLine($"Thread ID: {Thread.CurrentThread.ManagedThreadId} - Count: {i}");
            Thread.Sleep(100); // Simulate work
        }

        Console.WriteLine("Thread End.");
    }
}
```

- Kết quả sau khi chạy:
```powershell
PS D:\HocTap\Practice\LeetCode\test\TestLeetCode> dotnet run
Multiple Thread in C#:
--------------------
Main thread continues running
Thread Start:
Thread ID: 4 - Count: 0
Thread ID: 4 - Count: 1
Thread ID: 4 - Count: 2
Thread ID: 4 - Count: 3
Thread ID: 4 - Count: 4
Thread ID: 4 - Count: 5
Thread ID: 4 - Count: 6
Thread ID: 4 - Count: 7
Thread ID: 4 - Count: 8
Thread ID: 4 - Count: 9
Thread End.
```

---
### **Quản lý Thread trong C#**

Trong C#, bạn có thể quản lý các luồng (Thread) thông qua các thuộc tính và phương thức của lớp `Thread`. Dưới đây là một số thuộc tính và phương thức thường dùng:

### Các thuộc tính của lớp `Thread`:

- **Name**: Cho phép ==đặt hoặc lấy tên== của Thread, giúp dễ dàng nhận diện khi xử lý đa luồng.
- **Priority**: Xác định ==mức độ ưu tiên== của Thread. Các mức ưu tiên bao gồm: `Lowest`, `BelowNormal`, `Normal`, `AboveNormal`, `Highest` và `TimeCritical`.
- **IsBackground**: Cho biết ==Thread có phải là luồng nền hay không==. Luồng nền sẽ tự động kết thúc khi tất cả các luồng chính kết thúc.
- **IsAlive**: Trả về giá trị boolean, cho biết ==Thread hiện tại có đang hoạt động hay không.==
- **ThreadState**: ==Xác định trạng thái hiện tại== của Thread. Trạng thái có thể là: `Unstarted`, `Running`, `WaitSleepJoin`, `Suspended`, `Stopped` hoặc `Aborted`.

### Các phương thức của lớp `Thread`:

- **Join()**: Dùng để ==chờ một Thread kết thúc trước khi tiếp tục thực thi luồng hiện tại==. Có thể truyền vào thời gian chờ tối đa dưới dạng số nguyên (milliseconds) hoặc `TimeSpan`.
- **Abort()**: ==Yêu cầu dừng một Thread==. Khi gọi phương thức này, một ngoại lệ `ThreadAbortException` sẽ được ném ra trong Thread bị hủy. Tuy nhiên, phương thức này ==không được khuyến khích== sử dụng vì có thể gây ra các trạng thái không ổn định.
- **Suspend()**: ==Tạm dừng Thread.== Tuy nhiên, phương thức này đã bị loại bỏ (deprecated) vì không an toàn trong môi trường đa luồng.
- **Resume()**: ==Tiếp tục thực thi một Thread đã bị tạm dừng==. Phương thức này cũng đã bị loại bỏ và không nên sử dụng trong các ứng dụng mới.

```csharp
public class Solution
{
    public static void Main(string[] args)
    {

        Console.WriteLine("Multiple Thread in C#:");
        Console.WriteLine("--------------------");
        Thread thread = new Thread(new ThreadStart(() => DemoThread(5)));
        Thread thread2 = new Thread(new ThreadStart(() => DemoThread2(5)));

        thread2.Name = "DemoThread2";
        thread.Name = "DemoThread";

        thread2.Priority = ThreadPriority.Highest;
        thread.Priority = ThreadPriority.Lowest;

        thread2.IsBackground = true;
        thread.IsBackground = false;

        Console.WriteLine($"Thread Name: {thread.Name}, Priority: {thread.Priority}, IsBackground: {thread.IsBackground}, Thread State: {thread.ThreadState}");
        Console.WriteLine($"Thread Name: {thread2.Name}, Priority: {thread2.Priority}, IsBackground: {thread2.IsBackground}, Thread State: {thread.ThreadState}");

        thread.Start();
        thread2.Start();

        Console.WriteLine($"Thread Name: {thread.Name}, Priority: {thread.Priority}, IsBackground: {thread.IsBackground}, Thread State: {thread.ThreadState}");
        Console.WriteLine($"Thread Name: {thread2.Name}, Priority: {thread2.Priority}, IsBackground: {thread2.IsBackground}, Thread State: {thread.ThreadState}");

        thread.Join(); // Wait for thread to finish
        thread2.Join(); // Wait for thread2 to finish

        Console.WriteLine($"Thread Name: {thread.Name}, Thread State: {thread.ThreadState}");
        Console.WriteLine($"Thread Name: {thread2.Name}, Thread State: {thread.ThreadState}");
    }

    public static void DemoThread(int n)
    {
        Console.WriteLine("Thread Start:");
        for (int i = 0; i < n; i++)
        {
            Console.WriteLine($"Thread ID: {Thread.CurrentThread.ManagedThreadId} - Count: {i}");
            Thread.Sleep(100); // Simulate work
        }
        Console.WriteLine("Thread End.");
    }

    public static void DemoThread2(int n)
    {
        Console.WriteLine("Thread Start:");
        for (int i = 0; i < n; i++)
        {
            Console.WriteLine($"Thread ID: {Thread.CurrentThread.ManagedThreadId} - Count: {i}");
            Thread.Sleep(100); // Simulate work
        }
        Console.WriteLine("Thread End.");
    }
}
```

- Sau khi chạy:
```powershell
PS D:\HocTap\Practice\LeetCode\test\TestLeetCode> dotnet run
Multiple Thread in C#:
--------------------
Thread Name: DemoThread, Priority: Lowest, IsBackground: False, Thread State: Unstarted
Thread Name: DemoThread2, Priority: Highest, IsBackground: True, Thread State: Unstarted
Thread Start:
Thread Name: DemoThread, Priority: Lowest, IsBackground: False, Thread State: Running
Thread ID: 5 - Count: 0
Thread Name: DemoThread2, Priority: Highest, IsBackground: True, Thread State: Running
Thread Start:
Thread ID: 4 - Count: 0
Thread ID: 5 - Count: 1
Thread ID: 5 - Count: 2
Thread ID: 4 - Count: 1
Thread ID: 5 - Count: 3
Thread ID: 5 - Count: 4
Thread End.
Thread ID: 4 - Count: 2
Thread ID: 4 - Count: 3
Thread ID: 4 - Count: 4
Thread End.
Thread Name: DemoThread, Thread State: Stopped
Thread Name: DemoThread2, Thread State: Stopped
```

---
### **Sự khác nhau giữa Thread và Task trong .NET?**

Trong .NET **Task** và **Thread** đều là hai khái niệm liên quan đến việc xử lý đa luồng (multithreading) trong ứng dụng

**Thread** là một luồng thực thi riêng biệt, nó ==có thể thực hiện các tác vụ đồng thời với các Thread trong cùng một ứng dụng==. Mỗi Thread điều sẽ có một bộ đếm lệnh (instruction pointer) riêng, một bộ đếm dữ liệu (data buffer) riêng, và các nguồn tài nguyên khác (như các biến, tài nguyên hệ thống, v.v) được chia sẽ với các thread khác

**Task** là một khái niệm cao hơn, nó mô tả tác vụ (task) cần thực hiện trong một ứng dụng.

---
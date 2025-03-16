
Async/await là một tính năng quan trong trọng C# giúp lập trình viên viết code bất đồng bộ một cách dễ dàng hơn.

![](https://images.viblo.asia/290d5f3d-8f8b-4fcf-ae47-a81a54bb4da2.jpg)
## **Khái niệm cơ bản:**

**Bất đồng bộ (Asynchronous)** là khả năng thực hiện các tác vụ mà không cần phải chờ tác vụ trước hoàn thành. Trong lập trình, điều này rất hữu ích khi:
- Thao tác với I/O (đọc/ghi file, network request).
- Xử lý tác vụ tốn thời gian.
- Gửi UI responsive trong ứng dụng.

**async/await** trong c# cung cấp cú pháp đơn giản để viết code bất đồng bộ mà vẫn có cấu trúc code tương tự như code đồng bộ. Nó đơn giản hóa việc xử lý các tác vụ mất thời gian (như gọi API, đọc file, truy vấn csdl) mà không chặn luồng chính (main thread).

---
## **Tại sao cần async/await?**

- Trong lập trình **đồng bộ**, các tác vụ (như tả dữ liệu từ web) sẽ **chặn luồng chính**, làm cho ứng dụng bị **"đơ"**.
- **Async/await** cho phép thực hiện các tác vụ này trong nền (background) mà vẫn giữ giao diện người dùng (UI) hoặc luồng chính hoạt động mượt mà.

---
## **Cách thực họat động async/await**

1. **async**: đánh dấu một phương thức có thể chứa các lệnh await. Cho phép nó có thể trả về một `Task` hoặc `Task<T>` và sử dụng `await` bên trong.
2. **await**: Chờ đợi một task hoàn thành mà không chặn luồng chính.

Khi một phương thức gặp từ khóa `await`, nó sẽ:
- Trả quyền điều khiển về cho phương thức gọi.
- Tiếp tục thực thi tác vụ được `await` hoàn thành.

---
## **Ví dụ trong thực tế:**

***Ví dụ 1 tải dữ liệu từ web***
- Khi ứng dùn của bạn cần tải dữ liệu từ internet, quá trình này có thể mất nhiều thời gian. Nếu không sử dụng **async/await**, ứng dụng sẽ ==đóng băng== trong lúc chờ dữ liệu tải về.

***Ví dụ 2 xử lý nhiều file:***
- Khi cần đọc và xử lý nhiều file lớn, sử dụng **async/await** giúp ứng dụng có thể đọc nhiều file đồng thời thay vì đọc từng file.

***Ví dụ 3 Hãy tưởng tượng bạn đang nấu ăn trong bếp:***

- **Tình huống đồng bộ (synchronous)**:
	- Bạn đặt nồi nước lên bếp để đun sôi. Trong khi chờ nước sôi (mất 5 phút), bạn đứng yên và không làm gì khác, chỉ nhìn nồi nước. Sau 5 phút, nước sôi, bạn tiếp tục nấu mì.
	- => Thời gian bị lãng phí vì bạn không làm gì trong lúc chờ.
	
- **Tình huống bất đồng bộ với async/await**:
	-  Bạn đặt nồi nước lên bếp (gọi là một tác vụ bất đồng bộ). Thay vì đứng chờ, bạn đi cắt rau, chuẩn bị gia vị (các tác vụ khác). Khi nước sôi (tác vụ hoàn tất), bạn quay lại để nấu mì.
	- => Hiệu quả hơn vì bạn làm được nhiều việc trong lúc chờ.

---
## **Ví dụ trong code:**

***Ví dụ 1 Tải dữ liệu từ web:***

```
public async Task<string> DownloadDataAsync(string url){
	using(HttpClient client = new HttpClient){

		//Await sẽ giải phóng thread hiện tại trong khi chờ tải xong
		string result = await client.GetStringAsync(url); 
		return result;
	}
}
```

***Ví dụ 2 xử lý file:***

```
public async Task<string> ReadFileAsync(String filePath){
	
	using(StreamReader reader = new StreamReader(filePath)){
		
		//Đọc file mà không block thread chính
		string content = await reader.ReadToEndAsync();
		return content;
	}
}
```

***Ví dụ 3 gọi nhiều API đồng thời***

```
public async Task<(string weather, string news)> GetDailyInforAsync(){

	//Tạo các task mà không chờ chúng hoàn thành ngay lập tức
	Task<string> weatherTask = GetWeatherAsync();
	Task<string> newsTask = GetNewsAsync();

	//Chờ cả 2 hoàn thành
	await Task.WhenAll(weatherTask, newsTask);

	//Lấy kết quả từ các task đã hoàn thành
	return (weatherTask, newsTask);
}
```

***Ví dụ 4 sử dụng async/await trong UI:***

```
private async void btnDownload_click(object sender, EventArgs e){

	btnDownload.Enabled = false;
	statusLable.Text = "Đang tải...";

	try{
	
		string data = await DownloadDataAsync("https://example.com/api/data");
		resultTextBox.Text = data;
		statusLable.Text = "Tải thành công!";
		
	}catch(Exception ex){
	
		statusLable.Text = "Lỗi" + ex.Message;
		
	}finally{

		btnDownload.Enabled = true;
	}
}
```

***Ví dụ 5: Đồng bộ (Synchronous) - Chặn luồng***
```
using System;

class Program
{
    static void Main()
    {
        Console.WriteLine("Bắt đầu tải dữ liệu...");
        string data = DownloadData(); // Giả lập tải dữ liệu mất 3 giây
        Console.WriteLine("Dữ liệu: " + data);
        Console.WriteLine("Hoàn tất!");
    }

    static string DownloadData()
    {
        System.Threading.Thread.Sleep(3000); // Giả lập thời gian chờ 3 giây
        return "Dữ liệu từ server";
    }
}
```
- Kết quả:
```
Bắt đầu tải dữ liệu...
(Đợi 3 giây)
Dữ liệu: Dữ liệu từ server
Hoàn tất!
```
- Luồng chính bị chặn 3 giây tại ==DownloadData(),== không làm được gì trong lúc chờ.

***Ví dụ 6: Bất đồng bộ với Async/Await***
```
using System;
using System.Threading.Tasks;

class Program
{
	// Main phải là async trong C# 7.1+
    static async Task Main(string[] args) 
    {
        Console.WriteLine("Bắt đầu tải dữ liệu...");
	    
	    // Bắt đầu tác vụ bất đồng bộ
        Task<string> downloadTask = DownloadDataAsync(); 
        Console.WriteLine("Đang làm việc khác trong khi chờ...");

		// Chờ kết quả mà không chặn luồng
        string data = await downloadTask; 
        Console.WriteLine("Dữ liệu: " + data);
        Console.WriteLine("Hoàn tất!");
    }

    static async Task<string> DownloadDataAsync()
    {
        await Task.Delay(3000); // Giả lập thời gian chờ 3 giây
        return "Dữ liệu từ server";
    }
}
```
- **Kết quả:**
```
Bắt đầu tải dữ liệu...
Đang làm việc khác trong khi chờ...
(Đợi 3 giây)
Dữ liệu: Dữ liệu từ server
Hoàn tất!
```
- **Giải thích:**
	- `DownloadDataAsync()` trả về `Task<string>` và chạy bất đồng bộ.
	- `await` không chặn luồng chính, nên "Đang làm việc khác..." được in ngay lập tức.
	- Khi tác vụ hoàn tất (sau 3 giây), `data` nhận giá trị và chương trình tiếp tục.

---
## **Những điều quan trọng cần biết:**

1.**Kiểu trả về:** phương thức async thường trả về:
- `Task` nếu không có giá trị trả về.
- `Task<T>` nếu trả về giá trị kiểu T.
- `void` chỉ dùng event handler.

2.**ConfigureAwait:** khi không cần quay lại context gốc (thường là UI thread), hãy sử dụng:
```
await someTask.ConfigureAwait(false);
```

3.**Xử lý ngoại lệ:** Exception trong code bất đồng bộ được bảo toàn:
```
try{
	await PotentiallyFailingMethodAsync();
	
}cath(Exception ex){
	// Xử lý ngoại lệ
}
```

4.**Async all the way:** nên sử dụng async/await xuyên suốt từ đầu đến cuối chuỗi lời gọi, không nên trộn lẫn với code đồng bộ.

5.**Đánh dấu async:** Phương thức chứa `await` phải có từ khóa `async` và trả về `Task` hoặc `Task<T>`.

6.**Không chặn luồng**: Tránh dùng `.Result` hoặc `.Wait()` trên `Task` vì chúng chặn luồng, làm mất ý nghĩa của `async`. Thay vào đó, dùng `await`.

---
## **Khi nào nên dùng Async/Await?**

- **Nên dùng** khi gọi I/O (như mạng, file, csdl), vì đây là các tác vụ thường mất thời gian và không cần CPU liên tục.
- **Không cần dùng** với tác vụ CPU-bound (như tính toán nặng), bạn ==nên dùng luồng== (`Thread)` hoặc `Task.Run`.

---
## **Các lỗi thường gặp:**

1.**Deadlock:** khi kết hợp async/await với .Wait() hoặc .Result
2.**Fire and forget:** quên xử lý exception trong void async
3.**Không sử dụng await trong phương thức async.**
4.**Sử dụng async không cần thiết** cho các tác vụ không I/O.

---
## **Kết luận:**

- **Async/await** giúp viết mã bất đồng bộ dễ đọc, không chặn luồng chính.
- Trong thực tế, nó giống như làm việc đa nhiệm: bạn bắt đầu một công việc lâu, làm việc khác trong lúc chờ, rồi quay lại khi công việc hoàn tất.
- Trong code, nó rất phổ biến trong ứng dụng web, mobile, hoặc bất kỳ nơi nào cần gọi API, đọc file, hoặc truy vấn dữ liệu.
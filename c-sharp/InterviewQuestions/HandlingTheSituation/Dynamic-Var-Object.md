
### **Dynamic**

Từ khóa `dynamic` là từ khóa để khai báo kiểu `dynamic` cho biến. Mà kiểu dữ liệu này sẽ được xác định ở thời điểm thực thi.

Cú pháp:
```csharp
dynamic <tên biến>;
```


**Đặc điểm:**
- Các đối tượng thuộc kiểu `dynamic` sẽ ==không xác định được kiểu dữ liệu== cho đến khi ==chương trình được thực thi==. Tức là ==trình biên dịch sẽ bỏ qua tất cả lỗi cú pháp==, và ==việc kiểm tra này sẽ thực hiện khi chương trình thực thi==.

*Ví dụ:*
```csharp
public static async Task Main(string[] args)
{
	dynamic name = "PhamGiaBao";
	name++;  //Sẽ báo lỗi ở đây lúc thực thi, vì toán tử ++ không hỗ trợ kiểu chuỗi.
}
```

*Kết quả:*
```powershell
Unhandled exception. Microsoft.CSharp.RuntimeBinder.RuntimeBinderException: Operator '++' cannot be applied to operand of type 'string'
   at CallSite.Target(Closure, CallSite, Object)
   at System.Dynamic.UpdateDelegates.UpdateAndExecute1[T0,TRet](CallSite site, T0 arg0)
   at Solution.Main(String[] args) in D:\HocTap\Practice\LeetCode\test\TestLeetCode\Program.cs:line 17
   at Solution.<Main>(String[] args)
```

- Hỗ trợ `Dynamic programming` (lập trình động) cho ngôn ngữ lập trình sử dụng kiểu dữ liệu tĩnh như C#.
- Cho phép gọi phương thức và thuộc tính trên một biến kiểu `dynamic` mà không cần biết kiểu dữ liệu thực sự của nó.
- Có thể gây ra lỗi thời gian chạy nếu cố gắn gọi phương thức hoặc thuộc tính không tồn tại của đối tượng.
- Giúp cải thiện khả năng tương thích với các ngôn ngữ và nền tảng (frameworks) động
- Có thể ép kiểu qua lại với các kiểu dữ liệu khác một cách bình thường.

---
### **Var**

Từ khóa `var` là từ khóa để khai báo một cách ngầm định kiểu dữ liệu và kiểu `anonymous` cho biến.

**Tính chất:**
- Kiểu dữ liệu của biến được xác đinh ở thởi điểm biên dịch.
- Kiểu dữ liệu `var` được sử dụng để khai báo biến với giá trị được gán cho nó.
- Không thay đổi kiểu dữ liệu của biến sau khi gán giá trị.

*Ví dụ: Gặp lỗi về sau khi biến bị thay đổi kiểu giá trị.*
```csharp
public static void Main(string[] args)
{
	var name = 5.05f;
	name = "Hello";
	Console.WriteLine(name);
}
```

*Kết quả:*
```powershell
D:\HocTap\Practice\LeetCode\test\TestLeetCode\Program.cs(17,10): error CS0029: Canno
t implicitly convert type 'string' to 'float' [D:\HocTap\Practice\LeetCode\test\Test 
LeetCode\TestLeetCode.csproj]
The build failed. Fix the build errors and run again.
```

- Ngăn ngừa lỗi ép kiểu dữ liệu
- Đặc biệt tiện khi làm việc với các kiểu dữ liệu phức tạp.


---
### **Object**

`object` là một kiểu dữ liệu cơ bản của tất cả kiểu trong C#.

**Tính chất:**
- Được xác định kiểu dữ lệu cơ bản nhất, có thể lưu trữ bất kỳ đối tượng nào.
- Mọi kiểu dữ liệu đều được kế thừa từ `System.Object`
- Thuộc kiểu tham chiếu.
- Kiểu dữ liệu được xác định tại thời điểm biên dịch, nhưng giá trị có thể thay đổi.
- Kiểu `object` có thể gán bất kỳ kiểu dữ liệu nào tại bất kỳ thời điểm nào.

*Ví dụ: về đối tượng object có thể thay đổi kiểu dữ liệu tại bất kỳ thời điểm.*
```csharp
public static void Main(string[] args)
{
	object name = 5.05f;
	name = "Hello";
	Console.WriteLine(name);
}
```

*Kết quả:*
```powershell
Hello
```

---
### **Phân biệt `dynamic`, `object` và `var`**

| Đặc điểm                                              | Object                                            | Var                                                       | Dynamic                            |
| ----------------------------------------------------- | ------------------------------------------------- | --------------------------------------------------------- | ---------------------------------- |
| Là một kiểu dữ liệu                                   | Phải                                              | Không                                                     | Phải                               |
| Phải khởi tạo giá trị khi khai báo                    | Không bắt buộc                                    | Bắt buộc                                                  | Không bắt buộc                     |
| Sử dụng để làm kiểu trả về hoặc tham số cho hàm       | Có                                                | Không                                                     | có                                 |
| Có khả năng ép kiểu qua lại với các kiểu dữ liệu khác | Có                                                | không                                                     | có                                 |
| Thời điểm xác định kiểu dữ liệu thực sự               | Là một kiểu dữ liệu nên không cần xác định gì nữa | Xác định ngay tại khai báo thông qua giá trị được gán vào | Xác định khi chương trình thực thi |


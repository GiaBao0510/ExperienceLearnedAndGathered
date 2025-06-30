
##### **Câu 1: Hãy giải thích sự khác biệt giữa value type và reference type trong C#. Cho ví dụ cụ thể về từng loại.:**
- **Value Type**: Lưu trữ trực tiếp giá trị của biến trong bộ nhớ stack. Khi gán giá trị của một biến value type cho biến khác, nó tạo bản sao độc lập. Ví dụ: int, double, struct, enum.
- **Reference Type**: Lưu trữ tham chiếu (địa chỉ) đến dữ liệu trong bộ nhớ heap. Khi gán, chỉ tham chiếu được sao chép, không phải dữ liệu. Ví dụ: class, interface, string, array, object.

*ví dụ:*
```csharp
int a = 10; // Value type
int b = a;  // b là bản sao của a
b = 20;     // Thay đổi b không ảnh hưởng a
Console.WriteLine(a); // Kết quả: 10

string s1 = "Hello"; // Reference type
string s2 = s1;     // s2 trỏ cùng địa chỉ với s1
s2 = "World";       // s2 trỏ đến một chuỗi mới, không ảnh hưởng s1
Console.WriteLine(s1); // Kết quả: Hello
```

##### **Câu 2: Trong lập trình hướng đối tượng (OOP), 4 nguyên tắc chính là gì? Bạn có thể giải thích ngắn gọn về tính đóng gói (Encapsulation) và cho ví dụ trong C# không?**
- Trong lập trình hướng đối tượng có 4 nguyên tắc chính là tính bao đóng, đa hình, trừu tượng và kế thừa.
- Tính đóng gói (Encapsulation) che giấu chi tiết triển khai bên trong đối tượng và chỉ cung cấp giao diện công khai qua các phương thức getter/setter hoặc thuộc tính (property). Điều này giúp bảo vệ dữ liệu và đảm bảo tính toàn vẹn.
- *Ví dụ:* 
```csharp
public class Safe
{
    private decimal money; // Thuộc tính private để che giấu dữ liệu

    // Sử dụng property thay vì getter/setter truyền thống
    public decimal Money
    {
        get { return money; }
        set
        {
            if (value >= 0) // Kiểm tra tính hợp lệ
                money = value;
            else
                throw new ArgumentException("Số tiền không thể âm!");
        }
    }
}
```


##### **Câu 3: Sự khác biệt giữa .NET Framework và .NET Core là gì? Tại sao .NET Core thường được ưu tiên trong các dự án hiện đại?**
-  **.NET Framework**: Là nền tảng cũ hơn, chỉ chạy trên Windows, phù hợp với các ứng dụng desktop hoặc web truyền thống (như ASP.NET Web Forms). Hỗ trợ nhiều thư viện nhưng ít được cập nhật.
- **.NET Core** (nay là .NET): Đa nền tảng (Windows, macOS, Linux), tối ưu cho hiệu suất cao, ứng dụng cloud, microservices. Hỗ trợ phát triển hiện đại (ASP.NET Core) và có cộng đồng phát triển mạnh mẽ.
- Ví dụ: .NET Core được ưu tiên trong các dự án web hiện đại vì tốc độ nhanh hơn và khả năng triển khai trên container như Docker.

##### **Câu 4: Trong C#, List và Array khác nhau như thế nào? Khi nào bạn sẽ sử dụng List thay vì Array?**
- Array: Kích thước cố định, hiệu suất cao hơn trong trường hợp biết trước số lượng phần tử. Ví dụ: int[] numbers = new int[5];.
- `List<T>`: Linh hoạt, có thể thêm/xóa phần tử dễ dàng. Ví dụ: `List<int> numbers = new List<int>(`); numbers.Add(10);.
- **Khi nào dùng**:
    - Dùng Array khi kích thước cố định và cần tối ưu bộ nhớ (ví dụ: ma trận cố định).
    - Dùng `List<T>` khi số lượng phần tử thay đổi động (ví dụ: danh sách người dùng từ API).

##### **Câu 5: Entity Framework là gì? Hãy giải thích ngắn gọn về Code First và Database First trong Entity Framework.**
- - **Entity Framework**: Là một ORM giúp ánh xạ dữ liệu từ database sang các đối tượng C# (models). Nó giảm thiểu việc viết SQL thủ công và tăng tốc phát triển.
- **Code First**: Lập trình viên định nghĩa các class C# trước, sau đó sử dụng lệnh migration (như Add-Migration và Update-Database) để tạo database.
- **Database First**: Tạo database trước, sau đó dùng công cụ (như EF Designer) để sinh ra các class C# từ database.

- Ví dụ Code First
```csharp
public class User
{
    public int Id { get; set; }
    public string Name { get; set; }
}

public class AppDbContext : DbContext
{
    public DbSet<User> Users { get; set; }
}
```

#### **Câu 6: RESTful API là gì? Một HTTP request với phương thức GET và POST khác nhau như thế nào? Cho ví dụ về một endpoint RESTful mà bạn có thể thiết kế cho một ứng dụng quản lý sách.**
- **RESTful API**: Là kiến trúc API dựa trên giao thức HTTP, sử dụng các phương thức (GET, POST, PUT, DELETE) để thao tác với tài nguyên. Mỗi endpoint đại diện cho một tài nguyên (resource).
- **GET vs POST**:
    - GET: Lấy dữ liệu (ví dụ: GET /api/v1/books để lấy danh sách sách).
    - POST: Tạo mới tài nguyên (ví dụ: POST /api/v1/books để thêm sách mới).

- *Ví dụ thiết kế endpoint cho ứng dụng quản lý sách:*
```text
GET /api/v1/books - Lấy danh sách tất cả sách
GET /api/v1/books/1 - Lấy thông tin sách có ID = 1
POST /api/v1/books - Tạo sách mới
PUT /api/v1/books/1 - Cập nhật sách có ID = 1
DELETE /api/v1/books/1 - Xóa sách có ID = 1
```

#### **Câu 7: Trong C#, try-catch-finally hoạt động như thế nào? Hãy viết một đoạn code mẫu sử dụng try-catch-finally để xử lý một ngoại lệ khi đọc file.**
- Câu lệnh `try-catch-finally` thường áp dụng trong việc bắt và xử lý các exception, cách thức hoạt động:
	- try: tại đây sẽ là nơi chứa các câu lệnh bình thường về một tác vụ nào đó.
	- catch: tại đây sẽ là nơi sẽ bắt các ngoại lệ .Nếu xuất hiện và đồng thời xử lý chúng
	- finally: là nơi dù có ngoại lệ xuất hiện hay không thì nó sẽ vẫn thực thi câu lệnh bên trong
- Ví dụ: xử lý ngoại lệ khi đọc file
```csharp
public async Task ReadFileAsync(string filePath)
{
    try
    {
        using StreamReader reader = new StreamReader(filePath);
        string line;
        while ((line = await reader.ReadLineAsync()) != null)
        {
            Console.WriteLine(line);
        }
    }
    catch (FileNotFoundException ex)
    {
        Console.WriteLine($"Lỗi: Không tìm thấy file - {ex.Message}");
    }
    catch (IOException ex)
    {
        Console.WriteLine($"Lỗi khi đọc file: {ex.Message}");
    }
    finally
    {
        Console.WriteLine("Hoàn thành đọc file.");
    }
}
```

#### **Câu 8: Giả sử bạn có một bảng Users với các cột Id, Name, Email, CreatedDate. Hãy viết một câu lệnh SQL để lấy tất cả người dùng được tạo trong tháng hiện tại.**
 
 ```sql
 select Id, Name, Email
 from User
 Where YEAR(CreatedDate) = YEAR(CURDATE()) AND
 MONTH(CreatedDate) = MONTH(CURDATE());
```

##### **Câu 9: Async và await trong C# được sử dụng để làm gì? Hãy viết một phương thức async đơn giản để gọi một API bất đồng bộ.**
- async/ await là một trong 2 từ khóa để cho biết là đánh đâu phương thức là bất đồng bộ. 
- Từ khóa Async để cho biết là phương thức này đã đánh dấu là phơng bất đồng bộ và có thể sử dụng từ khóa await bên trong phương thức đã đánh dấu là async. Phương thức được đánh dấu aync thường trả về kiểu `Task`  hoặc `Task<T>.`
- Từ khóa await được sử dụng khi phương thức đã đánh dấu là async. Khi gặp từ khóa await thì luồng thực thi sẽ bị tạm dừng và chương trình sẽ chuyển sang thực thi các phần khác của ứng dụng hoặc đợi tác vụ bất đồng bộ hoàn thành.

*Ví dụ:*
```csharp
public async Task<string> GetDataFromApiAsync(string url)
{
    using HttpClient client = new HttpClient();
    try
    {
        HttpResponseMessage response = await client.GetAsync(url);
        response.EnsureSuccessStatusCode();
        return await response.Content.ReadAsStringAsync();
    }
    catch (HttpRequestException ex)
    {
        Console.WriteLine($"Lỗi khi gọi API: {ex.Message}");
        return null;
    }
}
```

##### **Câu 10: Nếu ứng dụng của bạn gặp lỗi NullReferenceException, bạn sẽ làm gì để xác định nguyên nhân và khắc phục?**

- **NullReferenceException**: Xảy ra khi bạn cố truy cập một biến hoặc đối tượng có giá trị null.
- **Cách xử lý**:
    1. **Xác định nguyên nhân**: Dùng công cụ debug (như Visual Studio) để kiểm tra dòng code gây lỗi và giá trị của biến.
    2. **Kiểm tra null**: Sử dụng câu lệnh if (obj != null) hoặc toán tử ?. (null-conditional operator).
    3. **Khắc phục**: Thêm kiểm tra null hoặc khởi tạo đối tượng trước khi sử dụng.

*Ví dụ:*
```csharp
public void ProcessUser(User user)
{
    if (user == null)
    {
        throw new ArgumentNullException(nameof(user));
    }
    Console.WriteLine(user.Name); // An toàn vì đã kiểm tra null
}
```
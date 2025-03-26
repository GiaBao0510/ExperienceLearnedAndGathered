### 1.***Giới thiệu về cú pháp LINQ:***

LINQ (Language Integrated Query) cung cấp hai cách truy vấn:
- **Query Syntax** (Cú pháp giống SQL).
- **Method Syntax** (Cú pháp dùng phương thức mở rộng)

🔹**Query Syntax** thường dễ đọc và dễ hiểu với người quên SQL.
🔹**Method Syntax** linh hoạt hơn và được dùng phổ biến hơn trong thực tế.

***Ví dụ:*** lọc danh sách số chẵn từ một mảng số nguyên
```
int[] nums = {1, 2, 3, 4, 5, 6, 7, 8, 9, 10};

//Query syntax
var evenNumberQuery = from num in nums
						where num % 2 == 0
						select num;

//Method syntax 
var evenNumbersMethod = nums.Where( num => num % 2 == 0);

Console.WriteLine(string.Join(", ", evenNumbersMethod)); // Output: 2, 4, 6, 8, 10
```

---
### 2.***Query Syntax vs Method Syntax:***

##### **2.1 Query Syntax:**

- Bắt đầu với `from`, sau đó là các điều kiện (`where`, `orderby`, `select`,...)
- Giống **SQL** nên dễ đọc đối với người quen với SQL.

***Ví dụ:*** Lọc danh sách sinh viên có điểm > 7:
```
var students = new List<Student>{
	new Student{name = "An", score = 8},
	new Student{name = "Bình", score = 6},
	new Student{name = "Châu", score = 9},
	new Student{name = "Duy", score = 7},
};

//Query syntax
var highScores = from s in students
					where s.score > 7
					select s.Name;

Console.WriteLine(string.Join(", ", highScores)); // Output: An, Châu
```

##### **2.2 Method Syntax:**

- Sử dụng phương thức mở rộng như: Where(), OrderBy(), Select(),...
- Ngắn gọn hơn, dễ kết hợp với các phương thức khác.
- Thường được dùng nhiều trong thực tế do tính linh hoạt

📌 **Cùng bài toán trên nhưng dùng Method Syntax:**
```
var highScorersMethod = students.Where(s => s.score > 7).Select(s => s.Name);

Console.WriteLine(string.Join(", ", highScorersMethod)); // Output: An, Châu
```


>_"Query Syntax phù hợp khi viết các truy vấn đơn giản, dễ đọc, trong khi Method Syntax mạnh hơn khi cần kết hợp nhiều toán tử hoặc xử lý logic phức tạp."_

---
### 3.***Các toán tử cơ bản trong LINQ:***

##### ***3.1 Filtering (Lọc dữ liệu) - Where():***

Dùng để lọc danh sách dựa trên điều kiện.

📃 **Ví dụ: Lọc sản phẩm có giá trên 500K**
```
var products = new List<Product>
{
    new Product { Name = "Laptop", Price = 1500 },
    new Product { Name = "Chuột", Price = 300 },
    new Product { Name = "Bàn phím", Price = 700 }
};

var expensiveProduct = products.Where(p => p.Price >= 500);

foreach (var p in expensiveProducts)
    Console.WriteLine(p.Name); // Output: Laptop, Bàn phím
```

📃 **Ví dụ: Danh sách nhân viên theo lương**
```
var employees = new List<Employee>{
	new Employee { Name = "An", Salary = 5000 }, 
	new Employee { Name = "Bình", Salary = 3000 }, 
	new Employee { Name = "Châu", Salary = 7000 }
};

//Lọc nhân viên có lượng trên 4000
var highEarners = employees.Where(p => p.Salary > 4000);
foreach (var emp in highEarners) 
	Console.WriteLine($"{emp.Name}: {emp.Salary}"); 
	
// Output:
// An: 5000
// Châu: 7000
```

##### ***3.2 Sorting (Sắp xếp) - OrderBy(), OrderByDescending():***

Dùng để sắp xếp dữ liệu tăng/giảm dần.

📃 **Ví dụ: Sắp xếp danh sách sinh viên theo điểm giảm dần
```
var sortedStudents = students.OrderByDescending(s => s.score);
foreach(var i in sortedStudents)
	Console.WriteLine($"name: {i.name} - {i.score}");
```

📃 **Ví dụ: Sắp xếp sản phẩm theo giá
```
var products = new List<Product>
{
	new Product {name = "Laptop", price = 1500},
	new Product {name = "Tai nghe", price = 200},
	new Product {name = "Chuột", price = 400},
	new Product {name = "Bàn phím", price = 450}
};

//Sắp xếp theo giá tăng dần
var sortedProducts = products.OrderBy(p => p.price);
Console.WriteLine("Sắp xếp tăng dần:"); 
foreach (var p in sortedProducts) 
	Console.WriteLine($"{p.Name}: {p.Price}");


//Sắp xếp theo giá giảm dần
var sortedDesc = products.OrderByDescending(p => p.price);
Console.WriteLine("Sắp xếp giảm dần:"); 
foreach (var p in sortedDesc) 
	Console.WriteLine($"{p.Name}: {p.Price}");
```

##### ***3.3 Projection (Chuyển đổi dữ liệu) - Select():***

Dùng để lấy một phần dữ liệu từ danh sách gốc

📃 **Ví dụ: Lấy danh sách tên sinh viên từ danh sách sinh viên
```
var studentNames = students.Select(s => s.name);
Console.WriteLine(string.Join(", ", studentNames)); // Output: An, Bình, Châu
```

📃 **Ví dụ: Chuyển đổi dữ liệu sang định dạng mới
```
var students = new List<Student>
{
    new Student { Name = "An", Score = 8 },
    new Student { Name = "Bình", Score = 6 }
};

// Tạo danh sách chuỗi thông báo kết quả
var results = students.Select(s => $"{s.Name} đạt {s.Score} điểm");
Console.WriteLine(string.Join("\n", results));
// Output:
// An đạt 8 điểm
// Bình đạt 6 điểm
```

##### ***3.4 Set Operators (Toán tử tập hợp):***

📃 **Ví dụ: 

Các toán tử này giúp thao tác trên nhiều tập dữ liệu.

✅ **`Distinct()`** → Lọc bỏ phần tử trùng lặp.

```
int[] nums = { 1, 2, 2, 3, 4, 4, 5 };
var uniqueNums = nums.Distinct();
Console.WriteLine(string.Join(", ",uniqueNums )); // Output: 1, 2, 3, 4, 5
```

📃 **Ví dụ: Loại bỏ trùng lặp trong danh sách đơn hàng:
```
var orders = new List<string> { "Order1", "Order2", "Order1", "Order3", "Order2" };
var uniqueOrders = orders.Distinct();

Console.WriteLine("Danh sách đơn hàng duy nhất:");
Console.WriteLine(string.Join(", ", uniqueOrders));
// Output: Order1, Order2, Order3
```

✅ **`Union()`** → Gộp hai danh sách và loại bỏ phần tử trùng lặp.

```
var list1 = new int[]{1,2,3};
var list2 = new int[]{5,6,7};

var result = list.Union(list2);
Console.WriteLine(string.Join(", ", result)); // Output: 1, 2, 3, 4, 5
```

##### ***3.5 Grouping (Nhóm dữ liệu) - GroupBy():***

Dùng để nhóm dữ liệu theo một thuộc tính nào đó.

📃 **Ví dụ: Nhóm sản phẩm theo giá
```
var groupedProducts = products.GroupBy(p => p.price > 500 ? "Đắt":"Rẻ");

foreach(var group in groupedProducts)
	 Console.WriteLine($"{group.Key}: {string.Join(", ", group.Select(p => p.Name))}");

// Output:
// Đắt: Laptop, Bàn phím
// Rẻ: Chuột
```

📃 **Ví dụ: Nhóm khách hàng theo khu vực:
```
var customers = new List<Customer> { 
	new Customer { Name = "An", Region = "Hà Nội" }, 
	new Customer { Name = "Bình", Region = "TP.HCM" }, 
	new Customer { Name = "Châu", Region = "Hà Nội" } 
};

var groupByRegion = customers.GroupBy(p => p.Region);

foreach (var group in groupedByRegion)
	Console.WriteLine($"{group.Key}: {string.Join(", ", group.Select(c => c.Name))}");
```

---
### 4.***Deferred Execution & Immediate Execution:***

##### **Deferred Execution (thực thi trì hoãn):**

Đây là đặc điểm quan trong trong **LINQ**, đặc biệt với các toán tử như: ==Where()==, ==Select()==, ==OrderBy()==.

_"Hầu hết các toán tử trong **LINQ** sử dụng **Deferred Execution**, nghĩa là truy vấn chỉ được thực hiện khi dữ liệu dược truy cập. (Ví dụ: gọi ==.ToList()==, ==.Count()==, hoặc lặp qua ==foreach==)."_
##### **Immediate Execution:**

Ngược lại, một số phần tử như ==Count()==, ==Sum()==, ==ToList()==, ==Distinct()== thì thực thi ngay lập tức (Immediate Execution). Điều này có thể được đề cập để phân biệt:

_"Các toán tử như ==Count()==,  ==Distinct()== hoặc ==ToList()== thì sẽ thực thi ngay lập tức, không trì hoãn"._

---
### 5.***Tổng kết:***

🔹 **LINQ có 2 cú pháp: Query Syntax và Method Syntax.**  
🔹 **Query Syntax** dễ đọc nhưng ít linh hoạt, còn **Method Syntax** ngắn gọn và phổ biến hơn.  
🔹 **Các toán tử quan trọng:**  
✅ `Where()` → Lọc dữ liệu  
✅ `OrderBy()`, `OrderByDescending()` → Sắp xếp dữ liệu  
✅ `Select()` → Chuyển đổi dữ liệu  
✅ `Distinct()`, `Union()` → Xử lý tập hợp  
✅ `GroupBy()` → Nhóm dữ liệu
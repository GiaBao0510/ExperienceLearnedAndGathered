### 1. **Giới thiệu:**

Bên cạnh các toán tử cơ bản như `Where()`, `OrderBy()`, `Select()`, **LINQ** còn cung cấp nhiều **toán tử nâng cao** để xử lý nhiều loại dữ liệu phức tạp hơn.

Trong bài toán này sẽ tìm hiểu về:
✅ `Join()` → Kết hợp dữ liệu từ nhiều tập hợp
✅ `GroupJoin()` → Nhóm dữ liệu khi kết hợp tập hợp.
✅ `Aggregate()`, `Sum()`, `Average()`, `Min()`, `Max()` → Toán tử tổng hợp. 
✅ `Take()`, `Skip()`, `TakeWhile()`, `SkipWhile` → Phân trang dữ liệu
✅ `Any()`, `All()`, `Contains()` → Kiểm tra điều kiện.
✅ `First()`, `FirstOrDefault()`, `Single()`, `SingleOrDefault()` → Lấy phần tử đầu tiên hoặc duy nhất.

---
### 2. **Toán tử Join (kết hợp dữ liệu):**

Dùng để kết hợp dữ liệu từ nhiều nguồn tập hợp (Tương tự như `INNER JOIN`, `JOIN` trong SQL).

📌 **Ví dụ: Kết hợp danh sách sinh viên và lớp học**
```
using System;
using System.Collections.Generic;
using System.Linq;

class Student{
    public int Id;
    public String Name;
    public int ClassId;
    public Student(int _id, String _name, int _classid ){
        Id = _id;
        Name = _name;
        ClassId = _classid;
    }
}

class Class{
    public int Id;
    public String Name;
    public Class(int _id, String _name){
        Id = _id;
        Name = _name;
    }
}

public class HelloWorld
{
    public static void Main(string[] args)
    {
        
        var students = new List<Student>
        {
            new Student(1, "An", 1),
            new Student(2, "Binh", 2),
            new Student(3, "Chau", 1),
            new Student(4, "Dang", 2),
            new Student(5, "En", 3),
        };
        
        var classes = new List<Class>
        {
            new Class(1, "Class A"),
            new Class(2, "Class B"),
            new Class(3, "Class C")
        };
        
        // Thực hiện JOIN giữa sinh viên và lớp học
        var studentClasses = from s in students
                             join c in classes on s.ClassId equals c.Id
                             select new { s.Name, ClassName = c.Name };
        
        foreach (var sc in studentClasses)
            Console.WriteLine($"{sc.Name} hoc {sc.ClassName}");
    }
}
```

**🎯Output:**
```
An hoc Class A
Binh hoc Class B
Chau hoc Class A
Dang hoc Class B
En hoc Class C
```

👉 **Method Syntax**
```
var studentClassesMethod = students.Join(classes,
	s => s.ClassId,
	c => c.Id
	(s, c) => new {s.Name, ClassName = c.Name}
);

foreach (var sc in studentClassesMethod)
    Console.WriteLine($"{sc.Name} học {sc.ClassName}");
```

📌 **Ví dụ: Kết hợp đơn hàng và khách hàng:

```
var customers = new List<Customer>
{
    new Customer { Id = 1, Name = "An" },
    new Customer { Id = 2, Name = "Bình" }
};

var orders = new List<Order>
{
    new Order { Id = 1, CustomerId = 1, Amount = 500 },
    new Order { Id = 2, CustomerId = 2, Amount = 300 },
    new Order { Id = 3, CustomerId = 1, Amount = 700 }
};

var customerOrders = from c in customers
                     join o in orders on c.Id equals o.CustomerId
                     select new { c.Name, o.Amount };

foreach (var co in customerOrders)
{
    Console.WriteLine($"{co.Name} đã đặt hàng: {co.Amount}");
}
// Output:
// An đã đặt hàng: 500
// Bình đã đặt hàng: 300
// An đã đặt hàng: 700
```


---
### 3. **Toán tử GroupJoin (Nhóm dữ liệu khi JOIN):**

Dùng để nhóm dữ liệu từ nhiều tập hợp.

📌 **Ví dụ: Nhóm danh sách sinh viên theo lớp học**
```
var groupedStudents = classes.GroupJoin(studetns,
	c => c.Id,
	s => s.ClassId,
	(c, sGroup) => new {ClassName = c.Name, Students = sGroup}
);

foreach(var group in groupedStudents){
	Console.WriteLine($"Lớp: {group.ClassName}");
	
	foreach(var student in group.Students)
		Console.WriteLine($"- {student.Name}");
}
```

**🎯Output:**
```
Lớp: Class A
- An
- Chau
Lớp: Class B
- Binh
- Dang
Lớp: Class C
- En
```

##### **Sự khác biệt giữa `Join()` và `GroupJoin()`:**

_"`Join()`tạo ra kết quả phẳng (flat result) với từng cặp tương ứng, trong khi `GroupJoin()` tạo ra kết quả phân cấp (hierarchical result) với các nhóm."_

📌 **Ví dụ: Nhóm đơn hàng theo khách hàng:
```
var groupedOrders = customers.GroupJoin(orders,
    c => c.Id,
    o => o.CustomerId,
    (c, orderGroup) => new { CustomerName = c.Name, Orders = orderGroup });

foreach (var group in groupedOrders)
{
    Console.WriteLine($"Khách hàng: {group.CustomerName}");
    foreach (var order in group.Orders)
    {
        Console.WriteLine($"- Đơn hàng: {order.Amount}");
    }
}
// Output:
// Khách hàng: An
// - Đơn hàng: 500
// - Đơn hàng: 700
// Khách hàng: Bình
// - Đơn hàng: 300
```

---
### 4. **Toán tử Aggregate (Tính toán trên tập dữ liệu):**

Dùng để thực hiện các phép tính như tổng, trung bình, min, max.

📌 **Ví dụ: Tính tổng giá trị đơn hàng**
```
var order = new List<int> {100, 110, 200, 250, 500, 600};

//Tính tổng bằng Aggregate
int totalAmount = order.Aggregate((sum, item) => sum + item);
Console.WriteLine($"Tổng giá trị đơn hàng: {totalAmount}"); 
```

📌 **Ví dụ: Tính tổng hóa đơn với giảm giá:
```
var invoices = new List<int> { 1000, 2000, 1500, 3000 };
var totalWithDiscount = invoices.Aggregate(0, (total, next) => total + (next - next * 10 / 100));

Console.WriteLine($"Tổng sau giảm giá 10%: {totalWithDiscount}");
// Output: Tổng sau giảm giá 10%: 6750
```

✅ Các toán tử hỗ trợ tính toán khác:
```
var numbers  = new List<int>{ 100, 200, 300, 400, 500};

Console.WriteLine($"Sum: {numbers.Sum()}");
Console.WriteLine($"Average: {numbers.Average()}");
Console.WriteLine($"Min: {numbers.Min()}");
Console.WriteLine($"Max: {numbers.Max()}");
```

**🎯Output:**
```
Sum: 1500
Average: 300
Min: 100
Max: 500
```

---
### 5. **Toán tử phân trang (Take, Skip, TakeWhile, SkipWhile):**

Dùng để lấy một phần dữ liệu trong danh sách, thường dùng trong phân trang.

📌 **Ví dụ: Lấy 3 sản phẩm đầu tiên**
```
var products = new List<string>{
	"Usb", "laptop", "mouse", "keyboard",
	"headphones", "wireless headphones"
};

Console.WriteLine(string.Join(", ", products.Take(3) ));

//Output: Usb, laptop, mouse
```

📌 **Ví dụ: Bỏ qua 3 sản phẩm và lấy 2 sản phẩm kế tiếp:**
```
Console.WriteLine(string.Join(", ", products.Skip(3).Take(2) ));
```

📌 `TakeWhile()` lấy các phần tử từ đầu danh sách cho đến khi điều kiện không còn thỏa mãn.

📌 `SkipWhile()`bỏ qua các phần tử từ đầu danh sách cho đến khi điều kiện không còn thỏa mãn

📌 **Ví dụ: TakeWhile() và SkipWhile() - Lọc sản phẩm theo giá:
```
var products = new List<Product>
{
    new Product { Name = "Chuột", Price = 50 },
    new Product { Name = "Bàn phím", Price = 100 },
    new Product { Name = "Laptop", Price = 1500 },
    new Product { Name = "Tai nghe", Price = 200 }
};

// Lấy sản phẩm cho đến khi giá > 1000
var cheapProducts = products.TakeWhile(p => p.Price <= 1000);
Console.WriteLine("Sản phẩm giá rẻ:");
foreach (var p in cheapProducts)
{
    Console.WriteLine($"- {p.Name}: {p.Price}");
}
// Output:
// - Chuột: 50
// - Bàn phím: 100

// Bỏ qua sản phẩm giá < 1000, lấy các sản phẩm còn lại
var expensiveProducts = products.SkipWhile(p => p.Price < 1000);
Console.WriteLine("\nSản phẩm giá cao:");
foreach (var p in expensiveProducts)
{
    Console.WriteLine($"- {p.Name}: {p.Price}");
}
// Output:
// - Laptop: 1500
// - Tai nghe: 200
```

---
### 6. **Toán tử kiểm tra điều kiện (Any, All, Contains):**

Dùng để kiểm tra xem tập dữ liệu có thỏa mãn điều kiện nào đó không.

📌 **`Any()` - Kiểm tra danh sách có phần tử thỏa mãn điều kiện không**

```
bool hasExpensiveProduct = products.Any(p => p.Length > 6);
Console.WriteLine(hasExpensiveProduct);
```

📌 **`All()` - Kiểm tra tất cả phần tử có thỏa mãn điều kiện không**

```
bool allShortNames = products.All(p => p.Length < 10);
Console.WriteLine(allShortNames);
```

📌 **`Contains()` - Kiểm tra xem danh sách có chứa một phần tử không**

```
bool ContainsKeyboard = products.Contains("keyboard");
Console.WriteLine(ContainsKeyboard);
```

📌 **Ví dụ: Any() và All() - Kiểm tra điều kiện kho hàng:
```
var stock = new List<Product>
{
    new Product { Name = "Laptop", Stock = 5 },
    new Product { Name = "Chuột", Stock = 0 },
    new Product { Name = "Bàn phím", Stock = 10 }
};

bool outOfStock = stock.Any(p => p.Stock == 0);
Console.WriteLine("Có sản phẩm hết hàng không? " + (outOfStock ? "Có" : "Không"));
// Output: Có

bool allInStock = stock.All(p => p.Stock > 0);
Console.WriteLine("Tất cả sản phẩm còn hàng không? " + (allInStock ? "Có" : "Không"));
// Output: Không
```

---
### 7. **Toán tử Lấy phần tử đầu tiên, duy nhất (First, FirstOrDefault, Single, SingleOrDefault):**

Dùng để lấy phần tử từ danh sách mà không cần `Where()`.

📌 **`First()` - Lấy phần tử đầu tiên (throw Exception nếu không có phần tử)**

```
var firstProduct = products.First();
Console.WriteLine(firstProduct); // Output: Usb
```

📌 **`FirstOrDefault()` - Lấy phần tử đầu tiên hoặc trả về `null` nếu không có**

```
var emptyList = new List<string>();
var firstOrNull = emptyList.FirstOrDefault();
Console.WriteLine(firstOrNull ?? "Không có sản phẩm"); // Output: Không có sản phẩm
```

📌 **`Single()` - Lấy duy nhất một phần tử (throw Exception nếu có nhiều hơn 1)**

```
var numbers = new List<int>{5};
Console.WriteLine(numbers.Single());
```

📌 **`SingleOrDefault()` - Lấy duy nhất một phần tử hoặc `null` nếu không có**

```
var numbers = new List<int>();
Console.WriteLine(numbers.SingleOrDefault());
```

📌 **`SingleOrDefault()` - Tìm sản phẩm duy nhất:

```
var uniqueProducts = new List<Product>
{
    new Product { Name = "USB", Price = 20 }
};

var singleProduct = uniqueProducts.SingleOrDefault(p => p.Price < 50);

Console.WriteLine(singleProduct != null ? $"Sản phẩm: {singleProduct.Name}" : "Không tìm thấy");
// Output: Sản phẩm: USB
```

📌 `SelectMany()` - Làm phẳng danh sách đơn hàng chi tiết:

```
var customerOrders = new List<CustomerOrder>
{
    new CustomerOrder { Name = "An", Items = new List<string> { "Laptop", "Chuột" } },
    new CustomerOrder { Name = "Bình", Items = new List<string> { "Bàn phím" } }
};

var allItems = customerOrders.SelectMany(co => co.Items);
Console.WriteLine("Tất cả sản phẩm đã đặt:");
Console.WriteLine(string.Join(", ", allItems));
// Output: Laptop, Chuột, Bàn phím
```
##### **Hành vi của `Single()` và `SingleOrDefault()`:**

- "Nếu có nhiều hơn một phần tử thỏa mãn, `Single()` và `SingleOrDefault()` sẽ ném ==InvalidOperationException==."

---
### 8. **Tổng kết:**

🔹 **`Join()` & `GroupJoin()`** → Kết hợp dữ liệu từ nhiều nguồn.  
🔹 **`Aggregate()`, `Sum()`, `Average()`, `Min()`, `Max()`** → Tính toán dữ liệu.  
🔹 **`Take()`, `Skip()`** → Phân trang dữ liệu.  
🔹 **`Any()`, `All()`, `Contains()`** → Kiểm tra điều kiện trong tập dữ liệu.  
🔹 **`First()`, `FirstOrDefault()`, `Single()`, `SingleOrDefault()`** → Lấy phần tử đầu tiên hoặc duy nhất.
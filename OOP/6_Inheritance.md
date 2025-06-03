Kế thừa là một trong 4 tính chất của lập trình hướng đối tượng, nó là khả năng cho phép chúng ta định nghĩa ra một lớp mới dựa trên một lớp khác có sẵn, kế thừa giúp cho việc mở rộng code -bảo trì dễ dàng hơn.
- **Lớp cơ sở** là lớp mà được lớp  khác kế thừa
- **Lớp kế thừa** là lớp kế thừa lại các thuộc tính, phương thức từ lớp cơ sở.

![](http://www.btechsmartclass.com/java/java_images/OOP-Concept-Inheritance.png)

**Tính kế thừa** trong OOP cho phép chúng ta tạo mới các class sử dụng một class có sẳn và có thể mở rộng chúng ra. Các `class con` có hể kế thừa lại các thuộc tính (**properties**) và phương thức (**methods**) của class cha, có thể không cần định nghĩa lại các phương thức hoặc tái định nghĩa (**override**) hoặc thêm các phương thức sử dụng riêng ở lớp con.

**Ưu điểm của kế thừa**
- Cho phép xây dựng 1 lớp mới từ lớp đã có.
	- Lớp mới gọi là **lớp con (subclass)** hay **lớp dẫn xuất (derived class)**.
	- Lớp đã có gọi là **lớp cha (superclass)** hay **lớp cơ sở (base class)**
- Cho phép chia sẽ các thông tin chung nhằm tái sử dụng và đồng thời giúp ra dễ dàng nâng cấp, dễ dang bảo trì.
- Định nghĩa sự tương thích giữa các lớp, nhờ đó ta có thể chuyển kiểu tự động 

---
### **Đơn kế thừa (Single inheritance)**

![](http://www.hiepsiit.com/public/uploads/images/cplusplus/kethua.jpg)

Khi một lớp kế thừa một lớp khác, nó được goi là đơn thừa kế.

***Ví dụ về đơn kế thừa- [ Inherritance ]***

```csharp
public class Base{
	
	public Base()
	{
		Console.WriteLine("Constructor of base class");
	}
	
	public void DisplayMessage(){
		Console.WriteLine("Base class contents");
	}
}

public class Child: Base{
	public Child(){
		Console.WriteLine("Base class contents");
	}
}

public class Program
{
	public static void Main(string[] args){
		Child objchild = new Child();
		objchild.DisplayMessage();
	}
} 
```

Như bạn đã thấy trong ví dụ trước, chúng ta đã tạo một đối tượng `class Child `trong phương thức Main() và sau đó gọi `DisplayMessage()` của class Base. nếu bạn để ý thấy là `class Child` không hề có  phương thức `DisplayMessage()` trong nó. Rõ ràng là nó được kế thừ từ `class Base`. 

---
### **Kế thừa đa cấp (Multilevel Inheritance)**

Là dạng kế thừa theo **tầng/lớp**, lớp con kế thừa từ lớp cha. Sau đó lớp cháu kế thừa từ lớp con đó. Giống như: **Cha $\to$ Con $\to$ Cháu.**

![](http://222.178.203.72:19005/whst/63/=ldchZzfddjrenqfddjrznqf//wp-content/uploads/multi.jpg)

***Ví dụ:***
```csharp
class Person{
	public string Name {get; set;}
}

class Employee: Person{
	public int EmployeeId {get; set;}
}

class Manager: Employee{
	public int TeamSize {get; set;}
}
```

Khi một lớp kế  thừa một lớp khác được kế thừa bởi một lớp khác, nó được gọi là đa kế thừa. Kế thừa là bắc cầu nên lớp dẫn xuất cuối cùng có được tất cả các thành viên của tất cả các lớp cơ sở của nó.

---
### **Kế thừa phân cấp (Hierarchical Inheritance)**

Một lớp cha có thể được kế thừa từ **nhiều lớp con khác nhau.**

![](https://media.geeksforgeeks.org/wp-content/uploads/20210916230620/Hierarchicalgfg.jpg)

***Ví dụ:***
```csharp
class Vehicle {
    public void Start() => Console.WriteLine("Starting...");
}

class Car : Vehicle {
    public void Drive() => Console.WriteLine("Driving Car");
}

class Bike : Vehicle {
    public void Ride() => Console.WriteLine("Riding Bike");
}
```

**💡 Sử dụng:**
```csharp
Car car = new Car();
car.Start();  // từ lớp cha
car.Drive();  // riêng của Car

Bike bike = new Bike();
bike.Start(); // kế thừa từ Vehicle
bike.Ride();  // riêng của Bike
```

---
### **Đa kế thừa (Mutiple Inheritance)**

### ❗**Lưu ý trong C#:**

- **C# không hỗ trợ kế thừa nhiều lớp** trực tiếp để tránh **xung đột đa kế thừa** (vấn đề Diamond Problem).
- Nhưng bạn vẫn có thể dùng **nhiều interface** để đạt được mục đích tương tự.

![](https://www.programtopia.net/wp-content/uploads/2021/01/multiple_0.png)

**🎯 Ví dụ thực tế (qua interface):**

```csharp
interface IWorker {
    void Work();
}

interface IManager {
    void Manage();
}

class TeamLead : IWorker, IManager {
    public void Work() => Console.WriteLine("Working as developer");
    public void Manage() => Console.WriteLine("Managing the team");
}
```

⏩ **Giải pháp C# dùng interface thay cho đa kế thừa**.

---
#### **Lớp niêm phong `sealed`**
Trong lập trình, bạn có thể sẽ đánh dấu một lớp nào đó không bao giờ trở thành lớp cơ cở để phải sinh ra lớp khác - lớp đó gọi là *bị niêm phong*. Muốn niêm phong một lớp thì phải dùng từ khóa `sealed`

![](https://dotnettutorials.net/wp-content/uploads/2018/08/Sealed-Class-and-Sealed-Methods-in-C-with-Examples.png?ezimgfmt=ng:webp/ngcb8)

```csharp
sealed class A{

}

class B : A{     //Xuất hiện lỗi vì kế thừa lớp niêm phong

}
```

---
### **Các vấn đề trong kế thừa**

##### **Vấn đề về phương thức khởi tạo, phương thức hủy bỏ**

**Phương thức khởi tạo** mặc định của lớp cha luôn luôn được gọi mỗi khi có một đối tượng thuộc lớp con khởi tạo. Và được gọi trước phương thức khởi tạo của lớp con.

Nếu như lớp cha có **phương thức khởi tạo có tham số** thì sẽ đòi hỏi lớp con phải có phương thức khởi tạo tương ứng và hiện gọi phương thức khởi tạo của lớp cha thông qua từ khóa `base`.

Khi đối tượng lớp con bị hủy thì phương thức hủy bỏ của nớ sẽ được gọi trước sau đó mới gọi phương thức hủy bỏ của lớp cha để hủy những gì lớp con không hủy được.

##### **Vấn đề hàm trùng tên và cách gọi phương thức của lớp cha**

Ví dụ giải sử ta có `class User` có chứa phương thức `Info()`, nhưng bên trong `class staff` kế thừa từ `class User` nhưng  lại được thiết kế thêm phương thức `Info()` trùng tên với phương thức từ lớp `class User`. Nếu biên dịch thì sẽ nhận một cảnh báo.

Để khắc phục vấn đề thày thì dùng từ khóa `new` vào trước khai báo hàm `Info()` trong lớp `class staff`. Từ khóa `new` nhằm để đánh dấu đây là 1 hàm mới và hàm kế thừa từ lớp cha sẽ bị che khuất đi khiên **bên ngoài** sẽ không gọi được.

```csharp
public new void Info(){

}
```
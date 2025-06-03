### **State Pattern là gì?**

![](https://images.viblo.asia/02b7617f-8ba2-4890-8ae3-c70110fa73ef.png)

- State Pattern là một trong những mẫu kế thuộc nhóm Behavioral Patterns, giúp đối tượng thay đổi hành vi của mình khi trạng thái nội bộ thay đổi mà không cần thay đổi mã nguồn. Mẫu này phân chia logic của các trạng thái vào các lớp riêng biệt, và đối tượng chính sẽ được chuyển đổi giữa các trạng thái này khi cần thiết.

- Tần xuất sử dụng: trung bình

---
### **Tại sao cần dùng State pattern?**

![](https://images.viblo.asia/11c95ce4-736f-4fa6-89ae-8d2843091493.png)

- **Giảm sự phức tạp:** nếu xuất hiện nhiều trạng thái trong một đối tượng và các hành vi khác nhau phụ thuộc vào trạng thái đó, việc viết nhiều câu lệnh **if hoặc switch** để kiểm tra trạng thái sẽ làm cho mã trở nên phức tạp và khó bảo trì.
- **Dễ dàng mở rộng:** Khi cần thêm trạng thái mới thì chỉ cần viết lớp mới mà không cần phải thau đổi logic hiện có.
- **Tăng tính bảo trì:** Các hành vi phụ thuộc trạng thái sẽ được tách riêng, dễ bảo trì và kiểm thử hơn.

---
### **Khi nào nên dùng?**

- Khi một đối tượng có nhiều trạng thái khác nhau và hành vi của nó thay đổi dựa trên trạng thái hiện tại
- Khi muốn tránh việc sử dụng nhiều câu lệnh (if, switch) để kiểm tra trạng thái trong các phương thức.
- Thay đổi hành vi object dựa trên trạng thái object.
- Khi bạn cần thêm mới hoặc thay đổi hành vi của đối tượng mà không muốn làm ảnh hưởng đến logic hiện tại

---
### **Các triển khai State pattern**

State pattern thường bao gồm các thành phần sau:

- Context: Đối tượng chính chứa trạng thái và hoạt động.
- State Interface: Định nghĩa các hành vi mà mỗi trạng thái phải thực hiện
- Concrete State Classes: Các lớp cụ thể triển khai hành vi khác nhau cho từng trạng thái.

Các bước triển khai:
1. Tạo một **interface** hoặc **abstract class** đại diện cho các trạng thái (State).
2. Tạo các lớp con để định nghĩa hành vi cụ thể cho từng trạng thái.
3. Tạo **Context class** chứa tham chiếu đến một đối tượng trạng thái, và có các phương thức để thay đổi hành vi của đối tườn dựa trên trạng thái.

---
### **Demo tính năng quản lý đơn hàng:**

Xây dụng một hệ thống quản lý đơn hàng trực tuyến với các trạng thái:
- Processing (Đang xử lý)
- Shipped (Đang vận chuyển)
- Delivered (Đã giao hàng).

**Bước 1: tạo Interface OrderState**

```csharp
public interface OrderState{
    public void next(OrderContext context);
    public void prev(OrderContext context);
    public void printStatus();
}
```

**Bước 2: Tạo các lớp trạng thái cụ thể**

```csharp
//Trạng thái đang xử lý (Processing)
public class ProcessingState : OrderState{

    public void next(OrderContext context){
        context.setState(new ShippedState());
    }

    public void prev(OrderContext context){
        Console.WriteLine("The order is in its initial state.");
    }

    public void printStatus(){
        Console.WriteLine("Order is begin processed.");
    }
}

  
//Trạng thái đang vận chuyển
public class ShippedState : OrderState{

    public void next(OrderContext context){
        context.setState(new DeliveredState());
    }

    public void prev(OrderContext context){
        context.setState(new ProcessingState());
    }

    public void printStatus(){
        Console.WriteLine("Order has been shipped.");
    }
}

  
//Trạng thái đã vận chuyern xong đơn hàng
public class DeliveredState : OrderState{

    public void next(OrderContext context){
      Console.WriteLine("This order is already delivered.");
    }

    public void prev(OrderContext context){
      context.setState(new ShippedState());
    }

    public void printStatus(){
      Console.WriteLine("Order has been delivered.");
    }
}
```

- Giải thích:
	- **Processing:** khi đơn hàng trong giai đoạn xử lý, bạn không thể quay lại trạng thái trước đó, vì đó là trạng thái đầu tiên
	- **Shipped:** Khi đơn hàng đã được vận chuyển, bạn có thể quay lại trạng thái đang xử lý trước đó, hoặc bạn có thể cập nhật trạng thái đã giao thành công tiếp theo
	- **Delivered:** Khi đơn hàng đã được giao xong, thì không thể chuyển về trạng thái trước đó, mà chỉ có thẻ xem trạng thái hienj tại là "Delivered"

**Bước 3 tạo lớp OrderContext để lưu trữ trạng thái**

```csharp
public class OrderContext{

    private OrderState state;

    //Khởi tạo đơn hàng trong quá trình xử lý
    public OrderContext() =>
        state = new ProcessingState();

    public void setState(OrderState state) =>
        this.state = state;

    public void nextState() =>
        state.next(this);

    public void prevState() =>
        state.prev(this);

    public void printStatus() =>
        state.printStatus();
}
```

**Bước 4: Client Code**
```csharp
public class Solution {

    public static void Main(string[] args) {

		// Khởi tạo một đơn hàng mới
		OrderContext order = new OrderContext();

		// Trong quá trình xử lý
		order.printStatus();
		order.nextState();

		//Trong quá trình vận chuyển
		order.printStatus();
		order.nextState();

		//trong quá trình giao hàng
		order.printStatus();
		order.prevState();

		order.printStatus();
    }
}
```

---
### **Ưu & nhược điểm**

**Ưu điểm:**
- Đảm bảo nguyên tắc Single Resposibility (SRP): tách biệt mỗi state tương ứng với class riêng biệt
- Đảm bảo nguyên tắc Open/Closed Principle (OCP): có thể thêm mootn state mới mà không ảnh hưởng đến State khác hay Context hiện có.
- Giữ hành vi cụ thể tương ứng với mỗi State (trạng thái).
- Giúp chuyển State một cách rõ ràng.
- Loại bỏ các câu lệnh xét trường hợp (If, Switch case) giúp đơn giản code của context.

**Nhược điểm:**
- Việc sử dụng State pattern có thể quá mức cần thiết niếu state machine chỉ có một vài trạng thái hoặc hiếm khi thay đổi có thể dẫn đến việc tăng độ phức tạp của code.
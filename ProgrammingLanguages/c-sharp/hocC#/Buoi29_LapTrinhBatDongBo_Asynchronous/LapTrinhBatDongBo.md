
## **Khóa đối tượng:**
- Để khóa đối tượng hiện tại được tạo ra thì dùng từ khóa **lock** ,Khóa đối tượng giúp cho _các luồng khác_ không thể truy cập cập vào đối tượng đã bị khóa. Khi muốn truy cập vào đối tượng đã bị khóa thì phải đợi cho luồng hiện tại mở khóa đối tượng này mới được truy cập. _Example:_
	lock(X1){
		//Các khối lệnh....
	}; _Phải đợi cho đến khi thực thi các công việc bên trong khối lệnh xong thì các luồng khác mới có thể truy cập vào biến X1 được_

## **Class Task:**
- Một đối tượng lớp **Task** có thể triển khai một tác vụ.
- Khởi tạo:
	- **Task** t2 = **new** **Task**(Action);  _//() => {} (Không tham số không có kiểu trả về)_
	- **Task** t3 = **new** **Task**(Action< object>,object);  //(Object obj) =>{} (Có tham số là một đối tượng bất kỳ).
- Chạy tác vụ, thì sử dụng **phương thức Start()** trên đối tượng **Task**
- Để thực hiện việc xử lý hoàn thành từng tác vụ lần lượt một thì sử dụng phương thức **Wait()**.
- Thay vì phải gọi phương thức **Wait()** cho từng tác vụ để nó thực hiện các thao tác lần lượt. Thì dùng phương thức **Task.WaitAll( _Liệt kê tác vụ có sẳn_ ).**
- Không được khởi chạy 1 **Task** 2 lần.
- Sử dụng **Async** ở đằng trước hàm trả về có kiểu **Task**. Và trong hàm bất đồng bộ thì ít nhất có một **await** Và dùng từ khóa **await** thay thế cho đối tượng dùng phương thức **Wait** sẽ tiện hơn. Vì sau khi dùng từ khóa **await** cho đối tượng kiểu **Task** thì nó sẽ tự động trả về.
- _Example:_
	- x.Wait(); = await x; 
- Để tạo tác vụ có kiểu trả về thì dùng thêm phần như sau: < string>, < int>, < float>,... Trước từ khóa **Task**.
- _Example:_
	- **Task**<**string**> X = **new** **Task**<**String**>(**Func**<**String**>);  //_() => {return string;}_
	- **Task**<**string**> X1 = **new** **Task**<**String**>(**Func**<**object**,**String**>, **object**);  //_(object obj) => {return string;}_
- Khi lấy giá trị của hàm bất ddoonhf bộ có trả về thì tại đối tượng của lớp **Task** dùng phương thức **Result;**
	

## **Asynchronous:**
- Đây là một kỹ thuật lập trình để tạo ra những ứng dụng chạy đa luồng.
- Cho phép chúng ta tạo ra nhiều tác vụ có thể chạy song song và đồng thời với nhau và có thể chạy trên nhiều luồng với nhau.


## **Synchronous:**
- Đây là kỹ thuật lập trình đơn giản ,chạy đơn luồng.
- Khi ứng dụng có nhiều tác vụ thì các tác vụ ấy được viết code sắp xếp theo thứ tự nào đó. Lúc thi hành thì cũng sẽ thi hành theo từng tác vụ đã được sắp xếp. Tác vụ phía trước hoàn thành thì mới đến tác vụ phía sau. 
## **1. Phương thức ảo:**
- Phương thức ảo là phương thức được định nghĩa ở lớp cơ sở(Lớp cha) và phương thức đó bị ghi đè bởi lớp kế thừa(lớp con).
- Để tạo phương thức ảo cần ghi từ khóa **virtual** trước kiểu trả về của phương thức trong lớp cở sở.
- Để ghi đề phương thức cơ sở trong phương thức cha thì tại lớp con cần ghi từ khóa **override** trước kiểu trả về của phương thức của lớp con.
- Cách gọi phương thức được tạo sẵn ở lớp cha thì cần dùng từ khóa **"base".** 

## **2. Lớp trừu trượng:**
- Lớp trừu được được tạo ra dùng để làm cơ sử cho các lớp kế thừa.
- Muốn tạo lớp trừu tượng phải dùng từ khóa **abstract** trước từ khóa **class**;
- Lớp trừu trượng **không được dùng để tạo ra đối tượng**. Nó chỉ được dùng làm lớp cơ sở thôi.
- Trong lớp trừu tượng có thể tạo được lớp trừu tượng.
- Lớp trừu trượng chỉ được tạo thôi không được định nghĩa bên trong nó.
- Dùng từ khóa **abstract** trước kiểu trả về của phương thức để tạo lớp trừu tượng.
- Mục đích của lớp trừu tượng tại lớp cở sở là để cho lớp con ghi đè lên.
- Dùng từ khóa **base** từ lớp con để gọi phương thức trừu trượng từ lớp cha là không được.
- Nếu lớp con kế thừa từ lớp trừu trượng mà bên trong lớp trừu tượng có phương thức trừa tượng thì lớp con phải ghi đè lên lớp đó. Nếu không sẽ báo lỗi. Trước khi ghi đè phải dùng từ khóa **override**.

## **3. Giao diện (Interface):**
- Interface khá giống với lớp trừu tượng cùng không được dùng để tạo ra đối tượng mà chỉ được tạo ra để làm cơ sở cho lớp con kế thừa.
- Khác chỗ là bên trong **Interface** chỉ là phương thức trừu tượng, mà mấy phương thức trừu trượng được khai báo bên trong interface không cần dùng từ khóa **abstract** để khai báo.
- Một lớp kế thừa được nhiều lớp cha.

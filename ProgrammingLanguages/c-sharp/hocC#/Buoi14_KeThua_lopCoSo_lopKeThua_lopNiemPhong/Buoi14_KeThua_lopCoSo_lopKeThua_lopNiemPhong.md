
## **1. Tính kế thừa:**
- Là những lớp khác có thể sử dụng các phương thức , thuộc tính của lớp được kế thừa.
- Nếu kế thừa phương thức và muốn sửa đổi nó từ lớp cha thì thêm từ khóa **new** trước kiểu dữ liệu trả về của phương thức.
- Khi lớp con kế thừa và sửa đổi phương thức kế thừa từ lớp cha, thì mọi hoạt động điều liên quan đến lớp kế thừa đã sửa đổi từ lớp con. Nếu muốn sửa dụng lớp kế thừa chưa bị sửa đổi bới lớp con thì dùng từ khóa **base** sẽ giúp truy cập đến phương thức từ lớp cha.
- 

Ví dụ: Có 2 lớp A và B. Lớp B kế thừa từ lớp A .Vậy lớp B có thể sử dụng thuộc tính và phương thức có **phạm vi khác private từ lớp A**. (_Lớp A gọi là lớp cơ sở, lớp B được goị là lớp con_)

## **2.Lớp niêm phong(sealed):**
Là lớp mà các lớp khác không thể kế thừa được.

## **3. Hoạt động khởi tạo phương thức ở lớp cha và lớp con:**
Khi chạy thì lớp khởi tạo của lớp cha được thi hành trước lớp khởi tạo của lớp con


## **4.Chuyển kiểu ,ép kiểu giữa các đối tượng lớp thuộc cây kế thừa:**

Lớp cơ sở được quyền lấy lớp khởi tạo của lớp con, lớp con không được quyền lấy lớp khởi tạo của lớp cha. 
## **Type:**
- Chứa các thông tin về kiểu dữ liệu nào đó. Chẳng hạn như: _class, struct, int float,..._
- Lớp này là một thành phần cơ bản trong .Net được dùng cho **Reflection** (**Reflection**: lấy thông tin thời điểm thực thi).
- Example:
```
	int X1;

	//Các biến này chứa thông tin về kiểu dữ liệu
	Type type1 = typeof(int),
	type2 = typeof(double);
	var type3 = typeof(Array);

```
- Phương thức **GetType()** để lấy thông tin kiểu dữ liệu của một biến
- Phương thức **FullName** dùng để lấy tên kiểu dữ liệu mà nó tiếp nhận.
- Phương thức **GetProgeties()**: dùng để trả một mảng những thuộc tính của một mảng nào đó.
- Phương thức **GetFields()** dùng để đọc thông tin về các trường dữ liệu.
- Phương thức **GetMethods()** dùng để lấy thông tin về các phương thức có trong dữ liệu kiểu mảng.
- Lấy thông tin về kiểu của một đối tượng nào đó, dựa vào thông tin về kiểu có thể lấy thông tin cụ thể nào đó.
- _Kỹ thuật này được sử dụng trong lập trình nâng cao._

## **Attribute:**
- Thuộc tính bổ sung **attribute** là một phần của siêu dữ liệu.
- Nó cung cấp thông tin cho một lớp hoặc cung cấp thông tin bổ sung cho các thành viên trong lớp.
- Những thông tin được bổ sung bởi **Attribute** sẽ được sử dụng bởi các thư viện, Framwork, các trình biên dịch, Thông tin này có thể được lấy ra và sử dụng ở thời điểm thực thi.
- **_Thuộc tính ObsoleteAttribute:_** này dùng được đánh dấu là sử dụng được cho class, struct, Constructor ,int, Propety, methobs,....Khi thành phần nào đó được bổ sung thêm **ObsoleteAttribute** thì thành phần đó đã lỗi thời => không được sử dụng nữa
- Một số **_Attribute trong .Net_** được sử  dụng nhiều trong ứng dụng web
	- 1. **Required**: Cho biết là phải thiết lập giá trị , không được null.
	- 2. **StringLength**: Cho biết dữ liệu chuỗi nằm trong khoảng nào đó.
	- 3. **DataType**: Cho biết kiểu dữ liệu lưu vào database là kiểu dữ liệu gì.
	- 4. **Range**: Cho biết dữ liệu số nằm trong khoảng nào đó.
	- 5. **Phone**: Cho biết dữ liệu phải là số điện thoại.
	- 6. **EmailAddress**: Cho biết dữ liệu phải là email.
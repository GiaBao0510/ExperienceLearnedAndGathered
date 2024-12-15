## **1.NameSpace:**
- Namespace dùng để gọp các thành phần như **(class, struct, namspace child, enum, interface,...)** để dễ quản lý và sử dụng.
	_Ví dụ cách truy cập đến thành phần nào đó trong lớp:_ 
		<Tên namespace>.<Tên lớp>.<Tên thành phần>;
- Khi muốn sử dụng **_namespace_** khác phải sử dụng từ khóa using.

## **2. Sử dụng partial	chia nhỏ lớp ra thành nhiều file:**

- Từ khóa **partail** dùng để chia một lớp ra làm nhiều phần có thể trên nhiều file.
- Các file khác muốn sử dụng lớp đã phân chia thì phải kèm theo từ khóa **partial** cùng với tên lớp muốn quản lý. Và lớp đó cùng phải chung namespace, chung phạm vi truy cập.
- Sử dụng từ khóa **partail** ngay trước từ khóa **class**.

## **3. Lớp lồng nhau(nested):**
- Trong c# có thể khai báo 1 lớp nằm trong 1 lớp khác.
- Có thể thay đổi phạm vi truy cập của lớp con.
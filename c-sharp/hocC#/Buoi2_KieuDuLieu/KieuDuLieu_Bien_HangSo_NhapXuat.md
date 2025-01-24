## **Buổi 2: Kiểu dữ liệu, biến ,hằng số, nhập xuất dữ liệu**

<h3> > Biến: </h3>là một tên nào đó với kiểu dữ liệu bất kỳ, dùng để trỏ đến ô nhớ trong máy tính, từ tên biến đã được tạo này có thể lưu/đọc bất cứ giá trị nào ,tùy theo kiểu dữ liệu mà biến lưu trữ.

<u>Cú pháp:</u>  Kiểu_dữ_liệu Tên_Biến;

_Có 2 loại khai báo biến:
- _Tường minh:_ Biến này sẽ được quy định sẳn kiểu dữ liệu đàu vào trước khi tạo (ex: int, double, float, string, ...).
- _Ngầm định:_ Biến này sẽ được quy định sẳn kiểu dữ liệu khi nhận giá trị (ex: var). 

<h3> > Kiểu dữ liệu: </h3>Trước khi đặt tên biến thì phải quy định kiểu dữ liệu cho tên biến cần đặt.

<h3> > Tìm hiểu một số phương thức trong Console:</h3>
- **CapsLock:** cho biết là phiếm capslock trên bàn phím đã bật hay chưa (true | false).
- **CursorLeft:** đặt hoặc thiết lập vị trí con trỏ trên màn hình.
- **ForegroundColor:** thiết lập màu chữ.
- **WindowWidth:** thiết lập độ rộng trên màn hình.
- **Title:** thiết lập tiêu đề cho cửa xổ.
- **BackgroundColor**: thiết lập màu nền cho cửa sổ.
- **Beep()**: phát ra tiếng beep khi cửa sổ xuất hiện.
- **Clear()**: xóa sạch màn hình hiển thị.
- **ReadKey()**: dùng để đọc ký tự.
- **ReadLine()**: đọc một dòng ký tự. Tức là để cho người dùng nhập tay, khi người dùng bấm enter thì kết thúc việc đọc.
- **ResetColor:** thiết lập màu nền , màu chữ về trạng thái mặc định.
- **Write():** xuất giá trị dữ liệu ra màn hình. (không xuống dòng)
- **WriteLine()**: xuất giá trị dữ liệu ra màn hình. ( xuống dòng) 

### > Hằng số:
Hằng số là biến khi được tạo giá trị ,thì nó sẽ giữ nguyên giá trị đó và không thay đổi giá trị của nó. Để tạo hằng số thì phải đặt từ khóa **const** trước kiểu dữ liệu của tên biến.


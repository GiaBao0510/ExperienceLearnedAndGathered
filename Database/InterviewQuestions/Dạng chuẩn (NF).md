
### **Các khái niệm cơ bản:**

#### **Phụ thuộc hàm:**
Trong một quan hệ, nếu giá trị của bộ thuộc tính A có thể suy ra được nếu biết giá trị của bộ thuộc tính B ,thì ta nói A phụ thuộc hàm đầy đủ vào B.

***Ví dụ:*** HocSinh(MaHS, TenHS, NgaySinh)
TenHS và NgaySinh điều là thuộc tính phụ thuộc hàm vào MaHS

#### **Phụ thuộc hàm đầy đủ, phụ thuộc bộ phận:**
Trong một quan hệ, bộ thuộc tính A phụ thuộc hàm vào bộ thuộc tính B. Nếu loại bỏ một thuộc tính bất kỳ trong bộ thuộc tính B mà A vẫn phụ thuộc hàm vào B thì A được gọi là phụ thuộc hàm đầy đủ vào B. Ngược lại, A được gọi là phụ thuộc bộ phận vào B.

#### **Phụ thuộc bắc cầu:**
Trong một quan hệ, nếu thuộc tính A mà phụ thuộc vào thuộc tính B mà thuộc tính B cũng phụ thuộc hàm vào thuộc tính C thì thuộc tính A được gọi là phụ thuộc bắc cầu vào thuộc tính C thông qua B

---
### **Chuẩn hóa dữ liệu là gì?**

**Chuẩn hóa dữ liệu** là quá trình biểu diễn cơ sở dữ liệu dưới dạng chuẩn. Đây là một kỹ thuật thiết kế bảng trong cơ sở dữ liệu, chia các bảng lớn thành bảng nhỏ hơn và liên kết chúng bằng các mối quan hệ.
Quá trình này diễn qua với **mục đích** là loại bỏ hoặc làm giảm sự dư thừa và phụ thuộc của dữ liệu.

---
### **Các dạng chuẩn hóa cơ bản**

![](https://images.viblo.asia/f67dc66b-2714-46bc-88f8-ed3ef41a4864.png)

##### **Dạng chuẩn 1 - 1NF (First Normal Form):**

**Định nghĩa:** 
- Một bảng (quan hệ) được gọi là dạng chuẩn 1NF nếu và chỉ nếu toàn bộ giá trị của các cột trong bảng (quan hệ) đề chỉ chứa các giá trị nguyên tử (nguyên tố)
- Giá trị của mỗi thuộc tính trong mỗi bộ phải là giá trị đơn.

**Chuẩn hóa:** Miền giá trị của mỗi thuộc tính chỉ chứa giá trị đơn nguyên tử , không thể tách ra được.

**Xử lý:** Tách các thuộc tính có miền giá trị đa nguyên tử thành các thuộc tính có miền giá trị đơn nguyên tử.

***Ví dụ:***
Một bảng chưa thỏa 1NF:

| MASV | HOTEN | DIACHI         | MAMON     | TENMON         | DIEM |
| ---- | ----- | -------------- | --------- | -------------- | ---- |
| A01  | Jame  | quận 12        | M01M02    | CSDLAnh        | 89   |
| A02  | Join  | quận Gò Vấp    | M01       | CSDL           | 8    |
| A03  | Roman | quận Ninh Kiều | M01M02M03 | CSDLAnh Toán 1 | 689  |
Bảng trên đang vi phạm 1NF tại các cột MAMON, TENMON, DIEM vì chứa nhiều giá trị. Nên các giá trị trong cột không phải là nguyên tố. Sau đây là bảng đã thỏa dạng chuẩn 1NF:

| MASV | HOTEN | DIACHI         | MAMON | TENMON | DIEM |
| ---- | ----- | -------------- | ----- | ------ | ---- |
| A01  | Jame  | quận 12        | M01   | CSDL   | 8    |
| A01  | Jame  | quận 12        | M02   | Anh    | 9    |
| A02  | Join  | quận Gò Vấp    | M01   | CSDL   | 8    |
| A03  | Roman | quận Ninh Kiều | M01   | CSDL   | 6    |
| A03  | Roman | quận Ninh Kiều | M02   | Anh    | 8    |
| A03  | Roman | quận Ninh Kiều | M03   | Toán 1 | 9    |

##### **Dạng chuẩn 2 - 2NF (Second Normal Form):**

Một quan hệ (bảng) đủ tiêu chí thỏa dạng chuẩn 2NF, nếu:
- Là 1NF.
- Các thuộc tính không khóa phải phụ thuộc hàm đầy đủ vào khóa chính.

**Chuẩn hóa:**  Tách các thuộc tính không khóa phụ thuộc bộ phận vào khóa chính thánh quan hệ riêng, khóa của quan hệ mới là khóa bộ phận tương ứng ban đầu.

***Ví dụ 1:*** Loại bỏ thuộc tính không khóa ra khỏi quan hệ gốc **VD**: MuonTra(**SoThe, MaSach**, TenNguoiMuon, TenSach, NgayMuon, NgayTra).
- TenNguoiMuon và TenSach phụ thuộc vào khóa Sach(**MaSach**, TenSach).
- NguoiMuon(**SoThe**, TenNguoiMuon).
- MuonTra(**SoThe, MaSach**, NgayMuon, NgayTra)

***Ví dụ 2:*** Cho một quan hệ R = (A, B, C, D), Khóa chính là AB và tập phụ thuộc hàm F = {AB => C,  AB => D} là dạng quan hệ đạt dạng chuẩn 2.

***Ví dụ 3:*** Cho một quan hệ R = (A, B, C, D), Khóa chính là AB và tập phụ thuộc hàm F = {AB => C,  AB => D, B => DC} là dạng quan hệ không đạt dạng chuẩn 2. Vì B => DC là phụ thuộc hàm không đầy đủ vào khóa chính. Chúng ta đua về dạng chuẩn 2NF như sau:

![](https://funix.edu.vn/wp-content/uploads/2022/04/cd989226-7237-4333-b0e9-a02bfc067bf7.jpg)

Sau đây là bảng từ 1NF được thỏa 2NF, để đạt được dạng chuẩn 2 thì ta sẽ tách ra làm 3 bảng: `SinhVien`, `MonHoc`, `KetQuaMonHoc`

**SinhVien:**

| MASV | HOTEN | DIACHI         |
| ---- | ----- | -------------- |
| A01  | Jame  | quận 12        |
| A02  | Join  | quận Gò Vấp    |
| A03  | Roman | quận Ninh Kiều |

**MonHoc:**

| MAMON | TENMON |
| ----- | ------ |
| M01   | CSDL   |
| M02   | Anh    |
| M03   | Toan 1 |

**KetQuaMonHoc:**

| MASV | MAMON | DIEM |
| ---- | ----- | ---- |
| A01  | M01   | 8    |
| A01  | M02   | 9    |
| A02  | M01   | 8    |
| A03  | M01   | 6    |
| A03  | M02   | 8    |
| A03  | M03   | 9    |

##### **Dạng chuẩn 3 - 3NF (Third Nomal Form):**

Một quan hệ đủ tiêu chí là dạng **chuẩn hóa dữ liệu** 3NF nếu quan hệ đó:
- Là 2NF
- Các thuộc tính không khóa phải phụ thuộc trực tiếp vào khóa chính

**Chuẩn hóa:**
- Tách quan hệ mới gồm các thuộc tính phụ thuộc bắc cầu và thuộc tính không khóa mà nó phụ thuộc vào.
- Loại bỏ các thuộc tính phụ thuộc bắc cầu vào khóa chính trong quan hệ ban đầu.

**Ví dụ 1:** NV(**MaNV**, HoTen, MaDV, TenDV)
- Thấy TenDV phụ thuộc bắc vào khóa chính MaNV thông qua thuộc tính không khóa MaDV
- NV(MaNV, TenNV)
- DV(MaDV, TenDV)

***Ví dụ 2:*** Cho quan hệ R = (ABCDGH), khóa chính là AB và tập phụ thuộc hàm F = {AB => C, AB => D, AB => GH } là quan hệ đạt chuẩn 3NF.

***Ví dụ 3:*** Cho quan hệ R = (ABCDGH), khóa chính là AB và tập phụ thuộc hàm F = {AB => C, AB => D, AB => GH, G => DH } là quan hệ  không đạt chuẩn 3NF. Vì DH phụ thuộc hàm gián tiến vào khóa. Chúng ta đưa nó về dạng chuẩn 3NF như sau:
![](https://funix.edu.vn/wp-content/uploads/2022/04/b62ec803-e3eb-4397-9e44-93cf88d73074.jpg)

##### **Dạng chuẩn BCNF (Boyce Codd Normal Form):**

Một quan hệ đạt chuẩn BCNF nếu quan hệ đó:
- Là 3NF.
- Không có thuộc tính khóa mà phụ thuộc hàm vào thuộc tính không khóa.

**Chuẩn hóa:** 
- Tách các thuộc tính không khoá và thuộc tính khóa phụ thuộc hàm vào nó thành quan hệ mới, thuộc tính không khoá đó trở thành khóa trong quan hệ mới.
- Loại bỏ các thuộc tính khóa ở bước 1 khỏi lượt đồ gốc.
- Bổ sung các thuộc tính không khóa mà thuộc tính (đã loại bỏ ở bước 2) phụ thuộc vào khóa của quan hệ gốc

***Ví dụ 1:*** Cho quan hệ R = (ABCDGH), khóa chính là AB và tập phụ thuộc hàm F = {AB => C, AB => D, AB => GH } là quan hệ đạt chuẩn BCNF.

***Ví dụ 2:*** Cho quan hệ R = (ABCDGH), khóa chính là AB và tập phụ thuộc hàm F = {AB => C, AB => D, AB => GH, H => B } là quan hệ không đạt chuẩn BCNF. Vì có thuộc tính B phụ thuộc hàm vào thuộc tính không khóa H. Chúng ta sẽ đưa về dạng chuẩn BCNF như sau:

![](https://funix.edu.vn/wp-content/uploads/2022/04/e61db351-9cd6-42ed-b2d4-2879534b112d.jpg)

##### Phương pháp chuẩn hóa

---
## **Tài liệu tham khảo:**

1. [Chuẩn hóa dữ liệu là gì? Các dạng chuẩn hóa cơ bản (1NF, 2NF, 3NF, BCNF)](https://funix.edu.vn/chia-se-kien-thuc/chuan-hoa-cac-quan-he-ve-cac-dang-chuan-co-ban/)
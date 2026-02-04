
Kiểm tra phiên bản của git trên máy:
```bash
git --version
```

Cấu hình thông tin cá nhân:
- Vì sao điều này lại quan trọng. Vì khi mỗi lần commt mà mình tạo ra thì sẽ được gắn với tên và email này
```shell
# thiết lập tên (sẽ hiển thị gõ ở phần commit)
git config --global user.name = "Pham Gia B"

#Thiết lập email
git config --global user.email "PhamGiaBao123@gmail.com"

#Kiểm tra cấu hình
git config --list
```

*Ví dụ: kiểm tra cấu hình*
```bash
C:\Users\GIA BAO>git config --list
diff.astextplain.textconv=astextplain
filter.lfs.clean=git-lfs clean -- %f
filter.lfs.smudge=git-lfs smudge -- %f
filter.lfs.process=git-lfs filter-process
filter.lfs.required=true
http.sslbackend=openssl
http.sslcainfo=D:/ChuongTrinh/Git/mingw64/etc/ssl/certs/ca-bundle.crt
core.autocrlf=true
core.fscache=true
core.symlinks=false
pull.rebase=false
credential.helper=manager
credential.https://dev.azure.com.usehttppath=true
init.defaultbranch=main
safe.directory=D:/HocTap/CuuAmChanKinh
user.email=pgiabao2002@gmail.com
user.name=GiaBao0510
```

---
### Các thành phần chính trong Git

#### **Staging Area (Khu vực dàn dựng)**

**Mô tả:** 
Staging Area, còn được gọi là chỉ mục hay bộ nhớ đệm, nó là không gian trung gian nơi mình có thể định dạng và xem lại các sự thay đổi trước khi commit chúng vào kho cục bộ (Local Repository).

Nó hoạt động như một bộ đệm giữa thư mục làm việc (Source code nơi nó đang hiển thị trên máy) và khi lưu trữ cục bộ (Local Repository), cho phép bạn chuẩn bị và tổ chức các commit của mình

**Quy trình làm việc:**

- Thêm thay đổi:
	- `git add <file>` để thêm thay đổi file
	- `git add .` để thêm tất cả file thay đổi vào Staging Area 

- Xem các thay đổi đã staged: sử dụng `git status `hoặc `git diff --staged` để xem những thay đổi nào đã được staged
- Commit các thay đổi đã staged: sử dụng `git commit -m "commit message"` để commit các thay đỏi đã staged vào Local Repository.


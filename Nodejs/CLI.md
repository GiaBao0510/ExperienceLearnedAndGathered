
Kiểm tra phiên bản của node:
```powershell
node -v
npm -v
```


Lệnh khởi tạo dự án:
```powershell
npm init
```

Lệnh tải gói về dự án:
```powershell
npm install [Tên gói cần cài đặt] --save
```

Lệnh kiểm tra các gói toàn cục(global packages): Nghĩa là kiểm tra trên hệ thống đã cài những gói gì rồi. Chứ không phải kiểm tra trên project
```powershell
npm list -g --depth=0
```
- `-g`: liệt kê các gói ở chế độ toàn cục
- `--depth=0`: Chỉ hiển thị cấp 1, không thực hiện các dependency bên trong.
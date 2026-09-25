Bài 01: Sử dụng vòng lặp hiển thị dãy số từ 1 > 100
- Hiển thị dãy số từ 1 đến 100
- Bỏ các con số ngẫu nhiên 6, 48, 75, 89
- Mỗi số cách nhau bằng dấu phẩy
- Bỏ dấu phẩy ở vị trí cuối cùng

```go
func main() {

  

    // Biến khởi tạo
    dichDen := 100
    var xhtml = ""

  

    //Hiển thị dãy số từ 1 đến 100
    for i := 1; i <= dichDen; i++ {

  

        //Bỏ các con số ngẫu nhiên 6, 48, 75, 89
        if i == 6 || i == 48 || i == 75 || i == 89 {
            continue
        }

  

        //Nếu chưa đến đích thì in số hiện tại có kèm dấu phẩy. Ngược lại thì in không kèm dấy phẩy
        if i < dichDen {
            xhtml += fmt.Sprintf("%d ,", i)
        } else {
            xhtml += fmt.Sprintf("%d", i)
        }
    }

    fmt.Println("XHTML:", xhtml)
}
```

---
Bài 02: Sử dụng vòng lặp hiển thị dãy số từ 1 > 100
- Chỉ hiển thị nhưng số lẻ
- Cứ 3 số thì xuống dòng và không có dấu phẩy ở cuối dòng
- Mỗi số cách nhau bằng dấu phẩy
- Bỏ dấu phẩy ở vị trí cuối cùng
```go
func main() {

    // Biến khởi tạo
    dichDen := 100
    demSo := 0
    var xhtml = ""

    //Chỉ hiển thị số lẻ trong dãy số từ 1 đến 100
    for i := 1; i <= dichDen; i += 2 {
        xhtml += fmt.Sprintf("%d", i)
        demSo++ 

        //Cứ 3 số thì xuống dòng và không có dấu phẩy ở cuối dòng
        if demSo%3 == 0 {
            xhtml += "\n"

        } else if i == dichDen-1 {          //Nếu là số cuối cùng thì không thêm dấu phẩy
            continue

        } else {                        //Mỗi số cách nhau bằng dấu phẩy
            xhtml += ", "
        }
    }

    fmt.Println("XHTML:\n", xhtml)
}
```

---
Bài 03: Xây dựng ứng dụng hiển thị bảng cửu chương
- Cho phép người dùng nhập số bắt đầu và kết thúc
- Kiểm tra điều kiện input nếu người dùng nhập sai
- Mỗi bảng cửu chương có tiêu đề hiển thị

```go
func main() {

    // Biến khởi tạ
    var batDau, ketThuc int

    //Nhập thông tin vào biến
    fmt.Print("Vui long nhap so bat dau: ")
    fmt.Scanf("%d\n", &batDau)

    fmt.Print("Vui long nhap so ket thuc: ")
    fmt.Scanf("%d\n", &ketThuc)

  

    //Kiểm tra điều kiện là Kết thúc phải lớn hơn Bắt đầu
    if ketThuc <= batDau {
        fmt.Println("So ket thuc phai lon hon so bat dau, vui long thu lai!")

    } else if batDau == 0 || ketThuc == 0 { //Nếu bắt đầu hoặc kết thúc là 0 thì in thông báo
        fmt.Println("Bat dau va ket thuc phai khac 0, vui long thu lai!")

    } else {
        fmt.Printf("Bat dau la: %d - Ket thuc tai: %d \n", batDau, ketThuc)
  
        //Bảng cửu chương
        for i := batDau; i <= ketThuc; i++ {
            fmt.Printf("Bang cuu chuong: %d \n", i)

            for j := 1; j <= 10; j++ {
                fmt.Printf("%d x %d = %d\n", i, j, i*j)

            }
        }
    }
}
```
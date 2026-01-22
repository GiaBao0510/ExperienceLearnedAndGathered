Bài 01: Xây dựng chức năng tính diện tích và chu vi hình chữ nhật
- Áp dụng các kiến thức đã học về Struct
- Kiểm tra giữ liệu đầu vào

```go
type HinhChuNhat struct{

    chieuDai float32 `desc: "ChieuRong"`
    chieuRong float32 `desc: "Chieu dai"`
}
  
func (hcn *HinhChuNhat) TinhDienTichHinhChuNhat() float32{
    return hcn.chieuDai * hcn.chieuRong
}

func (hcn *HinhChuNhat) TinhChuViHinhChuNhat() float32{
    return (hcn.chieuDai + hcn.chieuRong) * 2
}

func checkInput() HinhChuNhat {

    var Hcn HinhChuNhat

    //Kiểm tra đầu vào chiều rộng
    for {
        fmt.Print("Nhập kích thước chiều rộng: ")
        _, err := fmt.Scanf("%f\n", &Hcn.chieuRong)

        if err == nil && Hcn.chieuRong > 0 {
            break;
        }

        fmt.Println("Chiều rộng không hợp lệ. Vui lòng nhập lại.")
    }

    //Kiểm tra đầu vào chiều dài
    for {

        fmt.Print("Nhập kích thước chiều dài: ")
        _, err := fmt.Scanf("%f\n", &Hcn.chieuDai)
  
        if err == nil && Hcn.chieuDai > 0 {
            break;
        }
        fmt.Println("Chiều dài không hợp lệ. Vui lòng nhập lại.")
    }

    return Hcn
}

func main() {
    hcn := checkInput()
  
    fmt.Printf("Dien tich hinh chu nhat: %.2f\n", hcn.TinhDienTichHinhChuNhat())
    fmt.Printf("Chu vi hinh chu nhat: %.2f\n", hcn.TinhChuViHinhChuNhat())
}
```
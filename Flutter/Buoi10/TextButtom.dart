//Bước 1. Khai báo thư viện
import 'package:flutter/material.dart';
//Bước 1. Khai báo thư viện
import 'package:flutter/material.dart';

//Bước 2: hàm main
void main() {

  //Bước 3: hàm run app giúp khởi dộng app
  runApp(MaterialApp(

    //Bước 4: sử dụng thành phần để xây dựng giao diện

    //Ghi ở phần thân, Thẻ home là thuộc tính trong MaterialApp
    home: SafeArea(     //GIao diện maàn hình

      child: Scaffold(  //Scaffold: Đại dienj cho toàn bộ khung màn hình
        body:  MyWidget(),

      ),
    ),
    //Giúp định dạng font chữ cho toàn bộ project. Làm mặc định có thể bị đè
    theme: ThemeData(
        fontFamily: 'RETROTECH'
    ),

    debugShowCheckedModeBanner: false,  //Xóa nhã bên góc phải màn hình
  ));

}

//Tạo 1 lớp con từ lớp StatelessWidget
class MyWidget extends StatelessWidget{

  //Ghi đè lại hàm build từ StatelessWidget, Đối tượng trả về là Widget
  @override
  Widget build(BuildContext context) {  //Dùng để hiển thị

    //Ném TextButton vào container để có thể dễ dàng đưa thuộc tính margin vào
    return Container(
      margin: EdgeInsets.all(30),

      //Để thêm icon thì phải trỏ đến icon
      child: TextButton.icon(
        //Xử lý click nút bấm
        onPressed: (){          //Nếu trong onPressed đặt giá trị là 'null' thì nút này sẽ bị vô hiệu hóa
          print('Click Me!');
        },

        //Bố trí nút bấm thông qua thuoc tính
        style: TextButton.styleFrom(
          backgroundColor: Colors.lightGreen[900],    //Màu nền của nút bấm
          foregroundColor: Colors.lightGreen[400],    //Màu khi nhận sự kiện click
          minimumSize: Size(200, 80),                 //Chỉnh sử kích cở của buttom với 2 tham số là chiều rộng và chiều cao
          padding: EdgeInsets.all(50),                //Đệm

          //Thay đổi hình dạng của nút bấm
          shape: RoundedRectangleBorder(
            borderRadius: BorderRadius.circular(15)    //Bo góc
          ),

          elevation: 20,                               //Tăng độ rộng của bóng khi đổ
          //Đổ bóng - Với hàm withOpacity() là để làm mờ đổ bóng
          shadowColor: Colors.blue.withOpacity(0.5),

          //Tạo viền - Với hàm BorderSide(). nhận 2 tham số là độ dày và màu
          side: BorderSide(width: 20, color: Colors.amber),

          disabledBackgroundColor: Colors.black,        //Màu nền này sẽ được áp dụng .Nếu nút này bị vô hiệu hóa
          disabledForegroundColor: Colors.white        //Màu chũ này sẽ được áp dụng .Nếu nút này bị vô hiệu hóa
        ),

        //Đặt nội dung của nút trong hàm child()
        label: const Text(
          'Text bottom!',
          style: TextStyle(
            fontSize: 20,
          ),
        ),

        //Thêm ICON - ICON không thể ngang hàng với lớp child .Nên đổi lớp child thành lớp label
        icon: const Icon(Icons.add, size: 30),

        //Vô hiệu
      ),
    );
  }
}
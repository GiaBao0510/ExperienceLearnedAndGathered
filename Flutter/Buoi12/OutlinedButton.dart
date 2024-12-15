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
      margin: const EdgeInsets.all(15),     //Margin

      //Trỏ đển icon để cho biết chuẩn bị tạo
      child: OutlinedButton.icon(
          //Nhận sự kiện khi được Click
          onPressed: () {
            print('Cliked');
          },
          //Nếu bị vô hiệu hóa thì viết như sau
        // onPressed: null,

          //Định dạng nút bấm
          style: OutlinedButton.styleFrom(
            backgroundColor: Colors.black,
            foregroundColor: Colors.amber,
            padding: const EdgeInsets.all(10),            //Padding
            minimumSize: const Size(200, 50),             //Kích thước nút bấm
            shape: RoundedRectangleBorder(                //ĐỊnh dạng
                borderRadius: BorderRadius.circular(15),  //Bo các góc
            ),

            elevation: 10,                                //Độ dày của bóng
            shadowColor: Colors.black,                    //Màu bóng đỏ

            //Viền
            side: const BorderSide(width: 2, color: Colors.amber),

            //Màu này sẽ được chuyển khi b vô hiệu hóa
            disabledBackgroundColor: Colors.grey,
            disabledForegroundColor: Colors.white,
          ),

          //Thường là khu vực chứa nội dung văn bản - Đổi sang label vì có sử dụng đến icon
          label: const Text(
           'Click me . Daddy!',
            style: TextStyle(
              fontSize: 21,
              fontWeight: FontWeight.w500
            ),
          ),

          //Icon
          icon: const Icon(Icons.access_alarm_rounded, size: 30),
      ),
    );
  }
}
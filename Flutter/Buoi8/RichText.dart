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
        body: Center(child: MyWidget()),

      ),
    ),
    //Giúp định dạng font chữ cho toàn bộ project. Làm mặc định có thể bị đè
    theme: ThemeData(
      fontFamily: 'Handmade Notes Regular'
    ),

    debugShowCheckedModeBanner: false,  //Xóa nhã bên góc phải màn hình
  ));

}

//Tạo 1 lớp con từ lớp StatelessWidget
class MyWidget extends StatelessWidget{

  //Ghi đè lại hàm build từ StatelessWidget, Đối tượng trả về là Widget
  @override
  Widget build(BuildContext context) {  //Dùng để hiển thị

    return  RichText(
        text: TextSpan(
          //Phải luôn để trong RichText. Khi có định dạng
          style: DefaultTextStyle.of(context).style,

          //Trong children chứa 1 mangr các TextSpan
          children: const <TextSpan>[
            TextSpan(text: "Hello"),
            TextSpan(text: ' bold', style: TextStyle( fontFamily: 'RETROTECH',fontWeight: FontWeight.bold, fontSize: 30)),
            TextSpan(text: ' world!')
          ]
        )
    );
  }
}
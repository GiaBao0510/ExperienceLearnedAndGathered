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

    //
    return Container(
      //color: Colors.green,    //Đặt màu cho bên trong khung chứa

      //Điều chỉnh chiều rộng và chiều cao
      width: 200,   //Chiều rộng
      height: 200,  //Chiều cao
      
      padding: const EdgeInsets.all(15),    //Đêm
      margin: const EdgeInsets.all(12),

      //Căn chỉnh góc tọa độ của đối tượng bên trong khung chưa. alignment: Alignment, AlignmentDirectional , FractionalOffset
      alignment:  Alignment.centerLeft,

      //decoration: color: , shape: , border: , borderRadius: . Nếu dùng các thuộc tính bên trong mà trùng với bên ngoài thì xóa thuộc tính trùng bên ngoài và giữ lại thuộc tính trùng bên trong
      decoration: BoxDecoration(
        color: Colors.blue,                                     //Màu container
        shape: BoxShape.rectangle,                              //Hình dạng
        borderRadius: BorderRadius.all(Radius.circular(20)),    //Bo góc
        border: Border.all(width: 5, color: Colors.black54 ),   //Viền cho container
      ),

      //Dùng để xoay hay đổi vị trí của container
      transform: Matrix4.rotationZ(0.2),

      child: const Text(
        'Khung chứa',

        //Định dạng nội dung bên trong khung chứa
        style: TextStyle(
          color: Colors.white,
          fontWeight: FontWeight.w400,
          fontSize: 30,
        ),
      ),
    );
  }
}
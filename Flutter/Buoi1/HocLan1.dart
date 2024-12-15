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
          
          appBar:AppBar(  //AppBar: Đại diện cho tiêu đề
            backgroundColor: Colors.red,  //Màu nền
            title: const Text("Tự học flutter") //Dòng tiêu đề
          ),
          
          body: const Center( //Center: Giúp căn chỉnh chữ ở giữa màn hình
              child: Text("Hello world")
          )
      ),
    ),
    
    debugShowCheckedModeBanner: false,  //Xóa nhã bên góc phải màn hình

  ));

}
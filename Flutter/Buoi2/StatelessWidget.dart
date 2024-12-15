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
        body: Center(child: MyWidget(true)),
          // appBar:AppBar(  //AppBar: Đại diện cho tiêu đề
          //   backgroundColor: Colors.red,  //Màu nền
          //   title: const Text("Tự học flutter") //Dòng tiêu đề
          // ),
          // body: const Center( //Center: Giúp căn chỉnh chữ ở giữa màn hình
          //     child: Text("Hello world")
          // )
      ),
    ),

    debugShowCheckedModeBanner: false,  //Xóa nhãn bên góc phải màn hình
  ));

}

//Tạo 1 lớp con từ lớp StatelessWidget
class MyWidget extends StatelessWidget{

  //Tạo các thuộc tính
  final bool loading;

  //Hàm xây dựng
  MyWidget(this.loading);

  //Ghi đè lại hàm build từ StatelessWidget, Đối tượng trả về là Widget. Bắt buộc ghi đề lên hàm build
  @override
  Widget build(BuildContext context) {  //Dùng để hiển thị
    // if(loading == true){
    //   return const CircularProgressIndicator();
    // }else{
    //   return const Text('State404');
    // }
    return loading? const CircularProgressIndicator(): const Text('State404');
  }
}
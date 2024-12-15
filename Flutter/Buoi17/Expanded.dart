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

    //Container , nơi chứa các nội dung bên trong
    return Container(
      //Dịnh dạng nội dung theo chiê ngang
      child: Column(
        //Mảng chứa nội dung theo chiều ngang
        children: [
          //Bố cục 1
          Expanded(
            flex: 1,    //Thuộc tiính này sẽ mở rộng hoặc thu hẹp không gian dựa trên chiều hiển thị trong Expanded.

              //Nội dung
              child: Container(
                color: Colors.green,
                height: 100,
              )
          ),

          //Cũng có thể đưa sizeBox Vào hay thì padding hay margin. Nếu nó nằm ở chính giữa các khối
          const SizedBox(height: 10,),

          //Bố cục 2
          Expanded(
              flex: 1,    //Thuộc tiính này sẽ mở rộng hoặc thu hẹp không gian dựa trên chiều hiển thị trong Expanded
              child: Container(
                color: Colors.yellow,
                height: 100,
              )
          ),

          //Cũng có thể đưa sizeBox Vào hay thì padding hay margin. Nếu nó nằm ở chính giữa các khối
          const SizedBox(height: 10,),

          //Bố cục 3
          Expanded(
              flex: 3,    //Thuộc tiính này sẽ mở rộng hoặc thu hẹp không gian dựa trên chiều hiển thị trong Expanded.
              child: Container(
                color: Colors.pink,
                height: 100,
              )
          ),
        ],
      ),
    );
  }
}
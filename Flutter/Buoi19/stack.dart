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
        fontFamily: 'Time New Roman'
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
      color: Colors.grey,
      width: 500,
      height: 700,

      child: Stack(

        //Thuộc tính textDirection với giá trị TextDirection.ltr hoặc TextDirection.rtl  điều chỉnh hướng hiển thị từ phải sang trái hoặc ngược lại
        //textDirection: TextDirection.ltr,

        //Thuộc tính này trong Stack giúp điều chỉnh hướng hiêển thị ội dung bên trong Stack. Ví dụ: top, topright, center, centerleft,...
        alignment: Alignment.center,

        //Đối với thuộc tính fit khi có giá trị:
             //StackFit.expand và StackFit.passthrough:  Thì đối tượng cuối sẽ ph lên toàn bộ
             //StackFit.loose: Thì trở về mặc địnhS
        fit: StackFit.loose,

        //Với thuộc tính có giá tri như sau: clipBehavior: Clip.none .Bên trong Stack Nó giúp những container có thể hiển thị 1 cách đầy đủ mà không bị cắt khi vượt quá kích thước của Stack
        clipBehavior: Clip.none,

        //Mảng nội dung
        children: [
          //Nội dung 1
          Container(
           color: Colors.blue,
           width: 300,
           height: 300,
          ),

          //Nội dung 2
          //Với Đặt Positioned bên trong Stack thì nó sẽ giúp bạn tự điều chỉnh hướng hiển thị của nội dung tùy thích.Phải có một số thuộc tính sau: top, left , right, bottom bên trong Align
          Positioned(
            left: 5,
            child: Container(
              color: Colors.green,
              width: 200,
              height: 900,
            ),
          ),

          //Nội dung 3
          //Với Đặt Align bên trong Stack thì nó sẽ giúp bạn tự điều chỉnh hướng hiển thị của nội dung tùy thích.Bắt buộc phải có thuộc tính alignment bên trong Align
          Align(
            alignment: Alignment.topRight,
            child: Container(
              color: Colors.yellowAccent,
              width: 100,
              height: 100,
            ),
          ),

          //Nội dung 4
          Container(
            color: Colors.red,
            width: 50,
            height: 50,
          ),
        ],
      ),
    );
  }
}
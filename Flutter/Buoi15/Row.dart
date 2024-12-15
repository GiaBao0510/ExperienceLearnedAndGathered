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
      color: Colors.yellow,
      child: Row(
        //Chỉ bao xung quanh nội dung được tạo (min) - còn (max) thì bao quanh hết chiều ngang
        mainAxisSize: MainAxisSize.max,

        //Dùng để căn chỉnh nội dung bên trong, Có thể: trái , phải hoặc giữa
        mainAxisAlignment: MainAxisAlignment.spaceEvenly,

        //Căn nội theo chiều dọc,Nếu có 1 nội dung nào đó to quá mức, Có the là: trên , giữa hoặc cuối
        crossAxisAlignment: CrossAxisAlignment.start,

        //Đây là 1 mảng, Các thành phần view trong children sẽ phn cách nhau bởi dấu phẩy
        children: [
          ElevatedButton(
            //Xử l sự kiện click
              onPressed: (){},

            //Thiết kế  nội dung
            style: ElevatedButton.styleFrom(
              backgroundColor: Colors.red,
              foregroundColor: Colors.white
            ),

            //1 nội dung
              child: const Text(
                'bottom 1',
                style: TextStyle(
                  fontSize: 15
                ),
              ),
          ),

          Container(
            height: 100,
            child: ElevatedButton(
              //Xử l sự kiện click
              onPressed: (){},

              //Thiết kế  nội dung
              style: ElevatedButton.styleFrom(
                  backgroundColor: Colors.green,
                  foregroundColor: Colors.white
              ),

              //1 nội dung
              child: const Text(
                'bottom 2',
                style: TextStyle(
                    fontSize: 15
                ),
              ),
            ),
          ),
          ElevatedButton(
            //Xử l sự kiện click
            onPressed: (){},

            //Thiết kế  nội dung
            style: ElevatedButton.styleFrom(
                backgroundColor: Colors.blue,
                foregroundColor: Colors.white
            ),

            //1 nội dung
            child: const Text(
              'bottom 3',
              style: TextStyle(
                  fontSize: 15
              ),
            ),
          )
        ],
      ),
    );
  }
}
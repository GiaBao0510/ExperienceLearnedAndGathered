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

    //Sắp xếp các thành phần theo chiều dọc
    return Container(
      color: Colors.orange,

      child: Column(
        //Tạo kích thước bao quanh nội dung, Nếu min thì bao quanh vừa đủ , Ngược lại thì bao quanh hết cột đó
        mainAxisSize: MainAxisSize.max,

        //Căn chỉnh nội dung bên trong
        mainAxisAlignment: MainAxisAlignment.center,

        //Tùy các nội dung còn lại theo chiều ngang .Nếu 1 trong số noi dung nào đó có chiều rộng quá lớn. Thì có thể căn chỉnh theo: phải , trái hoặc giữa
        crossAxisAlignment: CrossAxisAlignment.center,

        //Chứa nội dung theo dạng cột
        children: [
          ElevatedButton(
              onPressed: (){},

              style: ElevatedButton.styleFrom(
                backgroundColor: Colors.red,
                foregroundColor: Colors.white
              ),

              child: const Text(
                'Buttom 1',
                style: TextStyle(
                  fontSize: 20
                ),
              ),
          ),
          Container(
            width: 300,
            child: ElevatedButton(
              onPressed: (){},

              style: ElevatedButton.styleFrom(
                  backgroundColor: Colors.green,
                  foregroundColor: Colors.white
              ),

              child: const Text(
                'Buttom 1',
                style: TextStyle(
                    fontSize: 20
                ),
              ),
            ),
          ),
          ElevatedButton(
            onPressed: (){},

            style: ElevatedButton.styleFrom(
                backgroundColor: Colors.blue,
                foregroundColor: Colors.white
            ),

            child: const Text(
              'Buttom 1',
              style: TextStyle(
                  fontSize: 20
              ),
            ),
          ),
        ],
      ),
    );
  }
}
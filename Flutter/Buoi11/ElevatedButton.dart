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
      margin:const EdgeInsets.all(20),               //Tạo khoảng cách với các phần tử xung quanh

      child: ElevatedButton.icon(               //Nếu cần icon thì phải chon label chứ không phải child
        //Nhận sự kiện khi click
          onPressed: (){
            print('Click OK.');
          },
          //Dùng cho trường hợp vô hiệu hóa
          //onPressed: null,
        
          //Định dang ElevatedButton
          style: ElevatedButton.styleFrom(
            backgroundColor: Colors.pink,
            foregroundColor: Colors.white,        //Màu nội dung
            minimumSize: const  Size(150, 50),    //Lấy định dạng kích cỡ (width, height)
            padding:const EdgeInsets.all(20),     //Khoảng cách giữa khung và nội dung
            shape: RoundedRectangleBorder(
              borderRadius: BorderRadius.circular(40)   //Chỉnh các góc bo
            ),
              elevation: 10,                            //Thuộc tiính này phải luônở trước shadowColor
            shadowColor: Colors.black.withOpacity(0.5), //Đổ bóng
            side: const BorderSide(width: 2, color: Colors.indigo),   //Tạo đường viền với độ rộng và màu sắc

            //Chí cập nhật khi bị vô hiệu hóa
            disabledForegroundColor: Colors.white,    //Màu chữ
            disabledBackgroundColor: Colors.grey      //Màu nền
          ),

          //Dinh dang Van ban ben trong. Vì bên trên child.icon nên bên dưới bắt buộc phải là label
          label: const Text(
            'Click Me!',
              style: TextStyle(
                fontSize: 25,
              ),
          ),
          //ICON
          icon: const Icon(Icons.ice_skating, size: 30),

      ),
    );
  }
}
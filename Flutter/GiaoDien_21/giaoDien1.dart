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

    //
    return Stack(
      children: [

        //Ảnh nền
        Container(

          child: Column(
            mainAxisAlignment: MainAxisAlignment.spaceAround,
            children: [
              Expanded(
                  child: Image.asset('assets/images/NewYorkTimes.jpg',width: 800, height: 1000, fit: BoxFit.cover,),
              ),
            ],
          ),
        ),



        //Nội dung
        Positioned(
          left: 10,
          right: 10,
          bottom: 25,
          child: Card(
            shape: RoundedRectangleBorder(
              borderRadius: BorderRadius.circular(8)
            ),
            child: Padding(
              padding: EdgeInsets.all(15),
              child: Column(
                 children: [
                  Text('New York', style: TextStyle(fontWeight: FontWeight.bold, fontFamily: 'Times New Roman', fontSize: 25),),
                  SizedBox(height: 10,),
                   Text('Công ty được thành lập bởi Henry Jarvis Raymond và George Jones '
                      'tại thành phố New York. Ấn bản đầu tiên của tờ báo The New York Times , '
                      'xuất bản vào ngày 18 tháng 9 năm 1851, nêu rõ: "Hôm nay chúng tôi '
                      'xuất bản số đầu tiên của tờ New-York Daily Times, và chúng tôi '
                      'dự định sẽ phát hành nó vào mỗi buổi sáng (trừ Chủ nhật) '
                      'trong một thời gian không xác định. số năm tới."',
                    textAlign: TextAlign.justify,
                    style: TextStyle(
                      fontSize: 18,
                    ),
                  )
                ],
              ),
            ),
          ),
        ),
      ],
    );
  }
}
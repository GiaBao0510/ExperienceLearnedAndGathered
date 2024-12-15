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

    debugShowCheckedModeBanner: false,  //Xóa nhã bên góc phải màn hình
  ));

}

//Tạo 1 lớp con từ lớp StatelessWidget
class MyWidget extends StatelessWidget{

  //Tạo các thuộc tính
  // final bool loading;
  //
  // //Hàm xây dựng
  // MyWidget(this.loading);

  //Ghi đè lại hàm build từ StatelessWidget, Đối tượng trả về là Widget
  @override
  Widget build(BuildContext context) {  //Dùng để hiển thị
    // if(loading == true){
    //   return const CircularProgressIndicator();
    // }else{
    //   return const Text('State404');
    // }
    return  const Text(''
        'Ung thư sắc tố xảy ra chủ yếu trên da, nhưng cũng có trên niêm mạc của '
        'vùng miệng, bộ phận sinh dục, trực tràng và kết mạc. Các khối u ác tính '
        'cũng có thể phát triển trong lớp màng mạch của mắt, trong màng não '
        '(pia hoặc màng nhện), và trong móng tay. Các khối u ác tính khác nhau '
        'về kích thước, hình dạng và màu sắc (thường là sắc tố) và xu hướng '
        'xâm lấn và di căn của chúng. Di căn xảy ra thông qua mạch bạch huyết và '
        'mạch máu. Sự di căn cục bộ dẫn đến sự hình thành các sẩn hoặc các '
        'nốt vệ tinh gần đó, có hoặc không có sắc tố. Di căn đến da hoặc các '
        'cơ quan nội tạng có thể xảy ra. Đôi khi, các khối u di căn hoặc các '
        'hạch bạch huyết mở rộng được phát hiện trước khi xác định được tổn '
        'thương ban đầu.'
        ,textDirection: TextDirection.ltr
        ,textAlign: TextAlign.justify
        , maxLines: 5
        , overflow: TextOverflow.ellipsis
        , textScaleFactor: 1.5
        ,style: TextStyle(  //TextStype dùng để bo tri văn bản
            color: Colors.white,            //Đổi màu băằng mã màu hoặc tu khóa
            backgroundColor: Colors.indigo, //Màu nền của văn bản
            fontSize: 20,                   //Phông chu đơn vị là pixel
            fontWeight: FontWeight.w400,     //Độ đậm nhạt của chữ
            fontStyle: FontStyle.italic,     //Phông chữ
            fontFamily: 'Times New Roman',   //Áp dụng kiểu phông chữ
            wordSpacing: 1,                  //Khoảng cách giữa các từ đơn vi pixel
            letterSpacing: 1,              //Khoảng cách giữa các ký tu
            decoration: TextDecoration.underline, //Kiểu của dòng văn bản. Ví dụ như: gạch chân, gạch giữa, bình thường,...

      ),
    );
  }
}
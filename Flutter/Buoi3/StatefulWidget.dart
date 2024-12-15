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
        body: Center(child: MyWidget2(false,0)),
      ),
    ),

    debugShowCheckedModeBanner: false,  //Xóa nhã bên góc phải màn hình
  ));

}

/*
  - lớp MyWidget2 là lớp con của lớp StatefulWidget. Nhung khi ghi đè thì lại không có ghi đè lên hàm build
  .Mà phải ghi đề lên hàm createState. Vì vậy để thuc hien ghi đè lên hàm biuld thì phải tạo 1 lớp là lớp con
  cùa hàm State , kiểu trả về là lớp vừa được tạo(VD: MyWidget2). Thì mới xây dựng được hàm build được
*/
class MyWidget2 extends StatefulWidget{
  //Thuộc tính
  final bool loading;
  final int SoLanBam;

  //Hàm xây dựng
  MyWidget2(this.loading, this.SoLanBam);

  //Bước 1(StatefulWidget) :Phải thực hiện ghi đè lên hàm createState(). Đối tượng trả về là State<StatefulWidget>
  @override
  State<StatefulWidget> createState() {
    return MyWidthget2State();
  }
}

//Bước 2 (StatefulWidget) tạo hàm con từ hàm State của thằng lớp thuộc StatefulWidget
class MyWidthget2State extends State<MyWidget2>{

  late bool _localLoading;
  late int _localSoLanBam;
  /*
    - Hàm initState() sẽ được gọi sau hàm MyWidget2. Nhưng s chạy trước hàm build()
    - Hàm này chủ yếu dùng để khởi tạo gía trị ban đầu
  */
  @override
  void initState() {
    //Ví dụ cho biến cục bộ này nhận gia trị của biến loading trong hàm MyWidget2
    _localLoading = widget.loading;
    _localSoLanBam = widget.SoLanBam;
  }

  @override
  Widget build(BuildContext context) {
    //Phải dùng widget để truy những những thuộc tính public. Của lớp MyWidget2

    return _localLoading? const CircularProgressIndicator(): FloatingActionButton(onPressed: NutBam);
  }

  //Tạo hàm nút bấm để thay đổi trang thái của các thuộc tính
  void NutBam(){
    setState(() {
      _localLoading = true;
      _localSoLanBam++;
    });
  }
}
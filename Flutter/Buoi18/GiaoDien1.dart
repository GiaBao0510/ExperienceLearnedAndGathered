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
      padding: const EdgeInsets.all(10),

      child: Column(
        children: [

          //Dòng 1
          Row(
            children:  [
              SizedBox(
                width: 100,
                child: Text('UserName:' ,style: TextStyle(fontSize: 14, color: Colors.grey),)
              ),
              Text('Phạm Gia Bảo', style: TextStyle(fontSize: 15),),
            ],
          ),

          const SizedBox(height: 5,),
          //Dòng 2
          Row(
            children:  [
              SizedBox(
                  width: 100,
                  child: Text('Email:' ,style: TextStyle(fontSize: 14, color: Colors.grey),)
              ),
              Text('pgiabao2002@gmail.com', style: TextStyle(fontSize: 15),),
            ],
          ),

          const SizedBox(height: 5,),
          //Dòng 3
          Row(
            children:  [
              SizedBox(
                  width: 100,
                  child: Text('Address:' ,style: TextStyle(fontSize: 14, color: Colors.grey),)
              ),
              Text('P.Tân An, Q.Ninh Kiều, TP.Cần Thơ', style: TextStyle(fontSize:15),),
            ],
          ),

          const SizedBox(height: 10,),

          //Dòng 4
          Row(
            children: [

              //Nut bấm hủy
              Expanded(
                  flex: 1,
                  child: ElevatedButton(
                    //Xử lý sự kiện
                    onPressed: (){
                      print("Đã hủy");
                    },

                    //Trang trí
                    style: ElevatedButton.styleFrom(
                        backgroundColor: Colors.grey,
                        foregroundColor: Colors.white,
                        shape: RoundedRectangleBorder(
                          borderRadius: BorderRadius.circular(5)
                        ),
                    ),

                    //Nội dung
                    child: Text('Cancel', style: TextStyle(fontSize: 18, fontWeight: FontWeight.w600),),
                  ),
              ),

              const SizedBox(width: 10,),
              //Nut nộp
              Expanded(
                flex: 1,
                child: ElevatedButton(
                  //Xử lý sự kiện
                  onPressed: (){
                    print("Đã nộp");
                  },

                  //Trang trí
                  style: ElevatedButton.styleFrom(
                    backgroundColor: Colors.lightBlue,
                    foregroundColor: Colors.white,
                    shape: RoundedRectangleBorder(
                        borderRadius: BorderRadius.circular(5)
                    ),
                  ),

                  //Nội dung
                  child: Text('Submit', style: TextStyle(fontSize: 18, fontWeight: FontWeight.w600),),
                ),
              ),
            ],
          ),
        ],
      ),
    );
  }
}
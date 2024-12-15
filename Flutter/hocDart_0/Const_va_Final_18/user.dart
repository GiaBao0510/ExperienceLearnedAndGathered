
class User {
  //Khai báo thuộc tính
  //Nếu không khai báo phạm vi thì sẽ mặc định là public
  final int namsinh; //Không xác định giá trị trong khoảng thời gian
  //Nếu khai báo biến mà có dấu gạch dưới trước tên biến thì quy định phạm vị biến đó ở mức private
  final String _mssv , _hoten;
  //Thêm biến với giá trị cố định
  static const int hoangdao = 12;

  //Hàm xây dựng. Nếu có cha thì thêm super
  const User(this.namsinh, this._mssv, this._hoten);

  //Hiển thị
  void InforUser(){
    print('Họ tên: ${_hoten}');
    print('Năm sinh: ${namsinh}');
    print('MSSV: ${_mssv}');
    print('Hoàng đạo: ${hoangdao}');
  }


}

extension ExtraStory on User{
  void additional(){
    print("Đại hoc Cần Thơ");
  }


}
import 'package:flutter/material.dart';
import 'package:tryapp/model/user.dart';

/*
  - Từ khóa "const" được dùng khi giá trị của bến được biêết trước và không đổi
  - final và const trên thực tế thì rất giống nhau. Cả 2 điều không thẻ thay đổi giá trị.
  - final, ít nghiêm ngặt hơn const .Vì trong khoảng thời gian nó có thể không xác định
   giá trị trong khoảng thời gian, ngay kể cả sau khi biên dịch. Khi mà nó mới nhận
   giá trị thì sẽ không thể thay đổi về sau.

   - instance varianbles chỉ có thể là final chứ không thể là const
   - Nếu muốn s dụng cost thì chỉ có thể thêm trước const là từ khóa static
*/
const double pi = 3.14;
final int ThapTu=14;
void main() {
  //runApp(const MyApp());
  User user1 = const User(2002, "B2016947", "Gia Bảo");
  User user2 = const User(2003, "B2016948", "Khá Bá");
  user1.InforUser();

  if( user1 == user2){
    print("Cùng địa chỉ ô nhớ");
  }else{
    print("Khác địa chỉ ô nhớ");
  }

}
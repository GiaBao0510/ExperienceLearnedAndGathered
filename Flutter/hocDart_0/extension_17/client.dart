import 'package:tryapp/model/user.dart';
//Ke thua: chi don ke thua
class Client extends User{
  Client(super.namsinh, super.mssv, super.hoten); //Gọi thàng supper để gọi đến cha

  //Co the lay tu extention cua lop User
  void inforClient(){
    additional();
  }

}
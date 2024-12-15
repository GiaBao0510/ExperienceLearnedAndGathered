import 'package:tryapp/model/basketball.dart';
import 'package:tryapp/model/football.dart';

//Co the mix voi nhieu lop khac
class User with FootBall, Basketball{
  //Khai báo thuộc tính
  //Nếu không khai báo phạm vi thì sẽ mặc định là public
  int namsinh = 0;
  //Nếu khai báo biến mà có dấu gạch dưới trước tên biến thì quy định phạm vị biến đó ở mức private
  String _mssv="Null" , _hoten="Null";

  //Hàm xây dựng. Nếu có cha thì thêm super
  User(this.namsinh, this._mssv, this._hoten);

  //Ham
  void inforFlavourite(){
    inforFootball();
    inforBasketball();
  }

  //Ghi de
  @override
  void play(){
    print('Choi da bong ');
  }
}
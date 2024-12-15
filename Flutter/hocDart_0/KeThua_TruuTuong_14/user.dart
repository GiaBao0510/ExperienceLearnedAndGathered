import 'package:tryapp/model/people.dart';
import 'package:tryapp/model/person.dart';

class User extends People{
  //Khai báo thuộc tính
  //Nếu không khai báo phạm vi thì sẽ mặc định là public
  int namsinh = 0;
  //Nếu khai báo biến mà có dấu gạch dưới trước tên biến thì quy định phạm vị biến đó ở mức private
  String _mssv="Null" , _hoten="Null";

  //Hàm xây dựng. Nếu có cha thì thêm super
  User(this.namsinh, this._mssv, this._hoten);


  //Hàm in
  void infor(){
    print('ID: $_mssv');
    print('Ho ten: $_hoten');
    print('Nam sinh: $namsinh');
  }

  @override
  void thongTin() {
    print('ID: $_mssv');
    print('Ho ten: $_hoten');
    print('Nam sinh: $namsinh');
  }
}
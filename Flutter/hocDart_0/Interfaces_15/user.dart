import 'package:tryapp/model/city.dart';
import 'package:tryapp/model/district.dart';

//Kế thua lớp trừu tượng
class User implements City, District{
  //Khai báo thuộc tính
  //Nếu không khai báo phạm vi thì sẽ mặc định là public
  int namsinh = 0;
  //Nếu khai báo biến mà có dấu gạch dưới trước tên biến thì quy định phạm vị biến đó ở mức private
  String _mssv="Null" , _hoten="Null";

  //Hàm xây dựng. Nếu có cha thì thêm super
  User(this.namsinh, this._mssv, this._hoten);

  //Ghi đè phương thức trong lớp city
  @override
  void showCity(){
    print('override show city.');
  }

  @override
  void showDistrict(){
    print('override show district.');
  }
}

void main(){
    //runApp(const MyApp());
  User ngdung = new User(2002, "b2016947", "Gia Bao");
  ngdung.showCity();
  ngdung.showDistrict();

}
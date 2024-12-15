class User{
  //Khai báo thuộc tính
  //Nếu không khai báo phạm vi thì sẽ mặc định là public
  int namsinh = 0;
  //Nếu khai báo biến mà có dấu gạch dưới trước tên biến thì quy định phạm vị biến đó ở mức private
  String _mssv="Null" , _hoten="Null";   //late nó sẽ mặc định là trước sau cac bien này sẽ có gtri thôi


  //Hàm xây dựng
  User(this._mssv ,this._hoten ,this.namsinh );

  //To String
@override
  String toString() {
    // TODO: implement toString
    return '$_mssv - $_hoten - $namsinh';
  }

  //Getter and Setter
  get hoten => _hoten;

  set hoten(value) {
    _hoten = value;
  }

  String get mssv => _mssv;

  set mssv(String value) {
    _mssv = value;
  }

  //Hàm hiển thị - Đây là hàm public
  void Infor(){
    print('Mssv:' + this._mssv);
    print('Ho ten:'+ this._hoten);
    print('Nam sinh: $namsinh' );
  }

  //Hàm nhập -  Đay là hàm private vì có "_" ở trước tên hàm
  void _EnterInfor(String id, String name, int nam){
    this._hoten = name;
    this._mssv = id;
    this.namsinh = nam;
  }
}
void main() {
  //runApp(const MyApp());
  User ngdung = User("B2016947", "GiaBao",2002);
  print(ngdung.toString());
  ngdung.Infor();

}
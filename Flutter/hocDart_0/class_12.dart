class User{
  //Khai báo thuộc tính
  late String mssv , hoten;   //late nó sẽ mặc định là trước sau cac bien này sẽ có gtri thôi


  //Hàm xây dựng
  User({required this.mssv ,required this.hoten );


  //To String
@override
  String toString() {
    // TODO: implement toString
    return '$mssv - $hoten';
  }
}
void main() {
  //runApp(const MyApp());
  User ngdung = User(mssv:"B2016947", hoten:"GiaBao");
  print(ngdung.toString());
}
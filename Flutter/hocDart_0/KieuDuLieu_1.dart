void main() {
  //Kiểu số nguyên
  int a1 =5, b=10;
  print(a1.runtimeType);
  print(a1+b);
  
  //Kiểu số thực
  double pi = 3.14;
  print(pi.runtimeType);
  print(pi);
  
  //Kiểu boolean
  bool checkin = false;
  print(checkin.runtimeType);
  print(checkin);
  
  //Kiểu chuỗi
  String hoTen = "Gia Bao",
          namSinh = "2002";
  print(hoTen.runtimeType);
  print(hoTen);
  
  //Cách chấp nhận biến có giá trị null: [KiểuDữLiệu]? tên biến;
  int? c;
  print(c);
  print(c.runtimeType);
  
  //Chuyển kiểu chuỗi sang kiểu số
  int namSinhThuc = int.parse(namSinh);
  print(namSinhThuc);
  print(namSinhThuc.runtimeType);
  
  //Chuyển kiểu số sang chuỗi
  String soPi = pi.toString();
  print(soPi);
  print(soPi.runtimeType);
}

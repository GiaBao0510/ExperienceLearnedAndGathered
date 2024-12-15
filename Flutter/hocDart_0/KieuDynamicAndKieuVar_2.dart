void main() {
  /*
   * Kiểu dữ liệu động. nếu ko gán giá trị thì sẽ nhận giá trị null,
   * Tự nhận kiểu dữ liệu dựa vào giá trị.
  */
  dynamic a;
  print(a.runtimeType);
  print(a);
  
  //Chuyển sang kiểu số nguyên
  a = 5;
  print(a.runtimeType);
  print(a);
  
  //Chuyển sang kiểu số thực
  a = 3.14;
  print(a.runtimeType);
  print(a);
  
  //Chuyển sang kiểu chuỗi
  a = "Gia Bảo";
  print(a.runtimeType);
  print(a);
  
  /*Kiểu var, nếu ko gán giá trị thì sẽ nhận giá trị null,
   * Tự nhận kiểu dữ liệu dựa vào giá trị. Gần giống với kiểu dynamic
   * Như nếu biến thuộc kiểu này được gán giá trị ban đầu thì ko như kiểu dynamic nữa
  */
  var b;
  print(b.runtimeType);
  print(b);
  
    //Chuyển sang kiểu số nguyên
  b = 5;
  print(b.runtimeType);
  print(b);
  
  //Chuyển sang kiểu số thực
  b = 3.14;
  print(b.runtimeType);
  print(b);
  
  //Chuyển sang kiểu chuỗi
  b = "Gia Bảo";
  print(b.runtimeType);
  print(b);
  
}
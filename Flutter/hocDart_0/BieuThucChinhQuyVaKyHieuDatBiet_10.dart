void main(){
  //1. [Điều kiện]? true:false
  var check = "hihi",
      check1, 
      record;
  record = (check != null)? 1:0;
  print(record);
  print('---------');
  
  //2. ?? kiểm tra xem biến này có được gán giá trị hay chưa trước khi dùng
  record = check1 ?? "Default";
  print(record);
  print('---------');
  
  //3. (..) Cascades
  List<double> diemso = [];
  diemso.add(5);
  diemso.add(8.2);
  diemso.add(7);
  diemso..add(9.15)..add(4.15)..add(5);
  
  diemso.forEach((i){
    print(i);
  });
  print('---------');
  
  //4. =>   //Đối với câu lệnh 1 dòng
  diemso.forEach((i) => print(i));
}
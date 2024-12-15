//Khai báo list
var ds = [];
List<dynamic> votri = ["BaoK",2002,1.5];
List<int> numbers = []; 

void main(){
  
  //Lấy kích thước danh sách
  print("Lấy số lượng trong danh sách");
  print(ds.length);
  print(numbers.length);
  
  //Thêm phần tử vào danh sách
  ds.add("G.Bao");
  ds.add(2002);
  ds.add("Mạnh giỏi");
  ds.add(369);
  ds.add("Active");
  ds.add(12);
  ds.insert(2, ">>>Insert");    //Thêm phần tử vào vị trí bất kỳ ([vị trí], [giá trị])
  numbers.add(08);
  numbers.add(01);
  numbers.add(2024);
  
  //Thêm danh sách khác vào danh sách
  ds.add(numbers);
  
  //Duyệt danh sách
  print("Duyệt danh sách");
  ds.forEach((i){
    print(i );
  });
  print("---------------");
  
  //Đảo chiều danh sách
  print("Đảo chiều danh sách");
  ds.reversed.forEach((i){
    print(i );
    print(i.runtimeType);
  });
  print("---------------");
  
  //Lấy phần tử đầu
  print("Lấy phần tử đầu");
  print(ds.first);
  
  //Lấy phần tử cuối
  print("Lấy phần tử cuối");
  print(ds.last);
  
  //Lấy phần tử ở vị trí bất kỳ
  print("Lấy phần tử ở vị trí bất kỳ");
  print(ds[1]);
  
  //Kiểm tra danh sách có rỗng không
  print("Kiểm tra danh sách có rỗng không");
  print(ds.isEmpty);
  
  //Kiểm tra danh sách không phải rỗng
  print("Kiểm tra danh sách không phải rỗng");
  print(ds.isNotEmpty);
  
  //Xóa dựa trên giá trị
  print("Xóa dựa trên giá trị");
  ds.remove(">>>Insert");
  ds.forEach((i){
    print(i);
  });
  print("---------------");
  
  //Xóa dựa trên vị trí
  print("Xóa dựa trên vị trí");
  ds.removeAt(3);
  ds.forEach((i){
    print(i);
  });
  print("---------------");
  
  //Xóa phần tử cuối
  print("Xóa phần tử cuối");
  ds.removeLast();
  ds.forEach((i){
    print(i);
  });
  print("---------------");
  
  //Xóa phần tử trong khoảng nào đó
  print("Xóa phần tử trong khoảng nào đó");
  ds.removeRange(0,3);
  ds.forEach((i){
    print(i);
  });
  print("---------------");
  
  //Làm rỗng danh sách
  print("Làm rỗng danh sách");
  ds.clear();
}
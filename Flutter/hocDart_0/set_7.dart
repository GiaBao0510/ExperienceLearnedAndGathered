
//Khai báo
var numbers = <int>{1999,2000,2002,2003,2004,2005};
Set<String> subject = <String>{"maths", "physics",
                   "chemistry", "geography",
                   "history","english"};
Set<String> subject1 = <String>{"informatics","philosophy"};
Set<dynamic> person = <dynamic>{"PhamGiaBao",2002,
                      "B2016947","KhoaHocMayTinh"};

void main(){
  //Lấy kích cỡ
  print("Lấy kích cỡ");
  print(subject.length);
  
  //Duyệt phần tử
  print("Duyệt phần tử");
  subject.forEach((i){
    print(i);
  });
  print("---------");
  
  //Lấy phần tử đầu
  print('Lấy phần tử đầu: '+subject.first);
  
  //Lấy phần tử cuối
  print('Lấy phần tử cuối: '+subject.last);
  
  //Lấy phần tử ở vị trí bất kỳ
  print('Lấy phần tử cuối: '+subject.elementAt(3));
  
  //Kiểm tra xem có rrongx hay không
  print("Kiểm tra xem có rrongx hay không:");
  print(subject.isEmpty);
  
  //Kiểm tra xem có phần tử hay chưa
  print("Kiểm tra xem có phần tử hay chưa:");
  print(subject.isNotEmpty);
  
  print("Kiểm tra bên trong tập set có giá trị hay chưa:");
  print(subject.contains("GiaBao"));
  
  //Thêm phần tử, Nếu giá trị trùng nhau sẽ bị loại bỏ
  subject.add("Arts");
  subject.add("Music");
  subject.add("geography");
  
  //Thêm set vào trong set khác nhưng phải cùng kiểu
  subject.addAll(subject1);
  print("Thêm set vào trong set khác nhưng phải cùng kiểu");
  subject.forEach((i){
    print(i);
  });
  print("---------");
  
  //Lệnh xóa dựa trên giá trị
  subject.remove("Music");
  
  //Lệnh xóa dựa trên nhóm
  subject.removeAll(subject1);
  print("Lệnh xóa dựa trên nhóm");
  subject.forEach((i){
    print(i);
  });
  print("---------");
  
  //Lệnh làm rỗng sets
  subject.clear();
}
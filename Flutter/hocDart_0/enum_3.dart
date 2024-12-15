//Cấu trúc dữ liệu enum
enum season {spring, NULL ,summer, autumn ,winter}

 void main() {
    //  1. Enum
  
  //Lấy tât cả các phần tử trong enum
  print(season.values);     
  
  //Lấy kích thước của đối tượng enum
  print(season.values.length);  
   
  //Lấy phần từ đầu tiên
  print(season.values[0]);    
  print(season.values.last);
   
  //Lấy phần tử cuối
  print(season.values[season.values.length-1]); 
  print(season.values.first);
  print(season.spring);  
   
  //In ra giá trị
  print(season.spring.name);   
  
  //Kiểm tra xem đối tượng enum có rỗng hay không
  print(season.values.isEmpty);
 
  //Kiểm tra xem dối tượn enum không rỗng
  print(season.values.isNotEmpty);
   
  //Duyệt các phần tử trong biến enum
  season.values.forEach((i){
    print(i.name);
  });
   
  //SWITCH - CASE
  season.values.forEach((i){
    switch(i.name){
      case "spring":
        print("Mùa xuân");
      case "summer":
        print("Mùa hè");
      case "autumn":
        print("Mùa thu");
      case "winter":
        print("Mùa đông");
      default:
        print("Các mừa trong năm");
      break;
    }
  });
   
}


















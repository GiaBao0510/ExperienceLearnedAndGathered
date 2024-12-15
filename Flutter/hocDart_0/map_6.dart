//Khai báo 
var person = {};
var mapNull = {};
var map2 = {"Hello":"Xin chào",
           "charge":"Phí",
           "unfortunately": "Không may",
           "accidental": "tình cờ",
           "random":"ngẫu nhiên",
           "default":"Mặc định",
           "object":"Đối tượng",
           "subject": "Đối tượng"};
var trans = {"subject": "Đối tượng",
            "running":"Đang chạy"};

void main(){
  //Kiểm tra số lượng phần tử
  print("Kiểm tra số lượng phần tử");
  print(map2.length);
  
  //Duyệt map
  print("Duyệt map ");
  map2.forEach((key, value){
    print('$key - $value');
  });
  print('------------');
  
  //Thêm phần tử
  person['HoTen'] = "Gia Bảo";
  person['NamSinh'] = 2002;
  person['MSSV'] = 'B2016947';
  print("Thêm phần tử ");
  person.forEach((key, value){
    print('$key - $value');
  });
  print('------------');
  
  //Thêm map này vào map khác
  print("Thêm map này vào map khác");
  mapNull.addAll(map2);
  mapNull.addAll(trans);
  mapNull.forEach((key, value){
    print('$key - $value');
  });
  print('------------');
  
  //Lấy phần tử dựa trên khóa
  print("Lấy phần tử dựa trên khóa");
  print(person['HoTen']);
  
  //Xóa phần tử dựa trên khóa
  print("Xóa phần tử dựa trên khóa");
  mapNull.remove('object');
  mapNull.forEach((key, value){
    print('$key - $value');
  });
  
  //Kiểm tra đối tượng map có tồn tại khóa hay giá trị nào đó không
  print("Kiểm tra đối tượng map có tồn tại khóa hay giá trị nào đó không");
  bool checkKey = mapNull.containsKey("object");
  bool checkValue = mapNull.containsValue("tình cờ");
  print('Check key: $checkKey - Check value: $checkValue');
  
  //Xóa tất cả
  print("Làm rỗng map");
  mapNull.clear();
  print(mapNull.length);
}
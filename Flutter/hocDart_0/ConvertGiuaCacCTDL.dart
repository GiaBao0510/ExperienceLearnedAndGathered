//Thêm thư viện
import 'dart:collection';

//Khai báo list
var ds1 = [1,5,2];
List<String> ds2 = ["tin", "toan", "Vat ly"];

void main(){
  //1.Chuyển kiểu cổ điển
  Set<String> obj1 = {};
  ds2.forEach((i){
    obj1.add(i);
  });
  
    //in
  print(">>>Chuyển kiểu cổ điển");
  obj1.forEach((i){
    print(i);
  });
  print("---------");
  
  //2.All tất cả phần tử
  Set<String> obj2 = {};
  obj2.addAll(ds2);
  print(">>>All tất cả phần tử");
  obj2.forEach((i){
    print(i);
  });
  print("---------");
  
  //3. .from()
  Set<dynamic> obj3 = Set.from(ds2);
  print(">>>.from()");
  obj3.forEach((i){
    print(i);
  });
  print("---------");
  
  //4. .map()
  List<dynamic> obj4 = ds2.map((i){
    return '$i';
  }).toList();
  print(">>>.map()");
  obj4.forEach((i){
    print(i);
  });
  print("---------");
  
} 
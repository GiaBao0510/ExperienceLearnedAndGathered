//Thêm thư viện
import 'dart:collection';

//Khai báo
var kb = Queue();
var subject = Queue<String>();
var subj = ['technology','literature'];
Queue check = Queue();

void main(){
  //Lấy kích cỡ
  print("Lấy kích cỡ");
  print(subject.length);
  
  //Thêm phần tử
  subject.add("philosophy");
  subject.add("physics");
  subject.add("chemistry");
  subject.add("Informatics");
  subject.add("math");
  
  //Thêm danh sách vào hàng đợi
  subject.addAll(subj);
  
  //Có thể thêm vào đầu hoặc cuối
  subject.addFirst("Arts");
  subject.addLast("Music");
  
  //Duyệt phần tử
  print("Duyệt phần tử");
  subject.forEach((i){
    print(i);
  });
  
  //Lấy phần tử dựa trên vị trí
  print('Lấy phần tử dựa trên vị trí: '+subject.elementAt(3));
  
  //Lấy phần tử dựa trên vị trí đầu
  print('Lấy phần tử dựa trên vị trí đầu: '+subject.first);
  
  //Lấy phần tử dựa trên vị trí cuối
  print('Lấy phần tử dựa trên vị trí cuối: '+subject.last);
  
  //Kiểm tra hàng đợi có rỗng không
  print("Kiểm tra hàng đợi có rỗng không");
  print(subject.isEmpty);
  
  //Xóa phần tử dựa trên giá trị
  subject.remove("Music");
  
  //Xóa đầu hoặc cuối
  subject.removeFirst();
  subject.removeLast();
  
  //Làm rỗng hàng đợi
  subject.clear();
} 
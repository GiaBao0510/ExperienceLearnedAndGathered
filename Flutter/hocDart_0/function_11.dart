//Hàm không trả về đối tượng
void syntaxError(){
  print('syntax error.');
}
void logName(String name) => print(name);

//Hàm trả về kết quả
int sum(int a, int b){ 
  return(a+b);
}

//Hàm cho phép chuyền tối thiểu là 0 hoặc 1 tham số
void laySo(int a, [int? b, int? c, int? d]){
  print(a);  
  print(b);
  print(c);
  print(d);
}

//Hàm name parameter: bắt buộc phải điền thên tham số và kèm thepo giá trị. Nhưng không quan trongnj về thứ tự
void layDe({int? a, int? b, int? c}){
  print(a);  
  print(b);
  print(c);
}
void main(){
  print("\n\t>>> Hàm void <<<");
  syntaxError();
  logName("GiaBao");
  
  print("\n\t>>> Hàm trả về kết quả <<<");
  print(sum(5,8));
  
  print("\n\t>>> Hàm nhận ít nhất 0 hoặc 1 tham số <<<");
  laySo(1,5);
  laySo(2);
  
  print("\n\t>>> Hàm name parameter: bắt buộc phải điền thên tham số và kèm thepo giá trị. Nhưng không quan trongnj về thứ tự <<<");
  layDe(b: 5, a:8, c:9);
}  
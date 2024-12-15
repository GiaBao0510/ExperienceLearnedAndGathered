void main(){
  //1. Toán tử số học
  int a=10,
      b=7,
      c=-15;
  print(a+b);
  print(a-b);
  print(a*b);
  print(a/b);
  print(a%b); //Chia lấy phần dư
  print(a ~/ b); //Chia lấy phần nguyên
  print(c.abs());
  print(a++);
  print(++a);
  print('---------');
  
  //2. Toán tử so sánh
  print(a>b);
  print(a<b);
  print(a>=b);
  print(a<=b);
  print(a!=b);
  print(a==b);
  print('---------');
  
  //3. Toán tử kiểm tra kiểu
  print(a is int);
  print(b is! double);
  print('---------');
  
  //4.Toán tử gán
  print(a+=b);
  print(a*=b);
  print(a%=b);
  print(a+=b);
  print('---------');
  
  //5. Toán tử logic
  bool checked1 = true;
  bool checked2 = false;
  print(checked1);
  print(!checked1);
  print(checked1 || checked2);
  print(checked1 && checked2);
  print('---------');
  
  //6.Toán tử Bitwwise và Shift
  print(a & b);
}  
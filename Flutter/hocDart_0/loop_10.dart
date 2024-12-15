List<int> numbers = [4,8,9,3,4,10,9];
void main(){
  print('1.----------');
  //For
  for(int i=0; i<numbers.length; i++){
    print(numbers[i]);
  }
  print('2.----------');
  
  //For in
  for(int i in numbers){
    print(i);
  }
  print('3.----------');
  
  //For each
  numbers.forEach((i)=> print(i)); 
  print('4.----------');
  
  //while
  int e = 0;
  while(true){
    if(e < numbers.length){
      print(numbers[e]);
      e++;
    }else{
      break;
    }
  }
  print('5.----------');
  
  //do while
  int l=0, r = numbers.length;
  do{
    print(numbers[l]);
    l++;
    r--;
  }while(l <= r);
//Kiểu Iterable
var numbers = Iterable.generate(10);

 void main() {

   //For- each
   numbers.forEach((number){
     print(number);
   });
   print('-------------------');
   
   //For
   for(int i=0; i< numbers.length; i++){
     print(i);
   }
   print('-------------------');
   
   //For - in
   for(var i in numbers){
     print(i);
   }
   
   //Độ dài phần tử
   print("Dộ dài");
   print(numbers.length);
   
   //Phần tử đầu
   print("Phần tử đầu");
   print(numbers.first);
   
   //Phần tử cuối
   print("Phần tử cuối");
   print(numbers.last);
}



















using System;
using MyException;

namespace buoi25
{
    class program
    {
        static void Register(string name, int age)
        {
            //Kiem tra nếu ,Tên rỗng thì phát sinh lỗi
            if(string.IsNullOrEmpty(name))
            {
                Exception e = new Exception("Ten khong duoc de rong");
                throw  new _NameEmptyException();
            }
            if (age < 1 || age > 100) {
                _AgeException e = new _AgeException();
                e.Detail();
            }

            Console.WriteLine($"Ten {name} - Tuoi: {age}. Hop le");
        }

        static void Main(string[] args)
        {
            int a = 15, b = 4;
            string path = "../../../bill.txt";
            try
            {

                var c = a / b;
                Console.WriteLine(c);

                int[] arr = { 1, 5, 9 };
                Console.WriteLine(arr[2]);

                //Đọc thông tin file thông qua đường dẫn rồi lưu vào biến
                string content = File.ReadAllText(path);
                Console.WriteLine(content);

                //===2.
                Register("GiaBao", 101);
            }
            catch(IndexOutOfRangeException e)   //Chi ra thong tin loi cu the
            {
                Console.WriteLine("Loi chi so trong mang khong hop le");
            }
            catch(FileNotFoundException fnfe)
            {
                Console.WriteLine(fnfe.Message);
                Console.WriteLine("Khong tim thay file");
            }
            catch(ArgumentNullException ae)
            {
                Console.WriteLine(ae.Message);
                Console.WriteLine("Loi duong dan phai khac null");
            }
            catch (Exception e){
                Console.WriteLine($"\t === Loi thuc thi === ");
                Console.WriteLine($"\n{e.Message}.");       //In nguyên nhân lỗi
                Console.WriteLine($"\n{e.StackTrace}.");    //In dấu vết lỗi
                Console.WriteLine($"\n{e.GetType().Name}.");//lấy kiểu và lấy tên lớp của exception

            }
            Console.WriteLine("ket thuc");
        }
    }
}

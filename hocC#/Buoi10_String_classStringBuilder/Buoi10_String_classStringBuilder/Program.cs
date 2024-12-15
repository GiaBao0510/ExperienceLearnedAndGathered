using System;
using System.Security.Cryptography;
using System.Text;

namespace buoi10
{
    class program
    {
        static void Main(string[] args)
        {
            string lastname = "Pham",
            firstname = "Gia Bao";
            string said = "Please have a seat. Can I offer you a drink?";

            //1.Nối chuỗi
            string fullname = lastname + " "+ firstname;
            Console.WriteLine("==============");
            Console.WriteLine("Full name: \n\t \"{0}\" ",fullname);

            //2.Viết chuỗi trên nhiều dòng
            string LittleMiss = 
@"Little Miss Muffet sat on a tuffet
Eating a Big Mac and fries
Along came a spider and sat down beside her
‘Yuck’, it said, ‘I prefer flies.";
            Console.WriteLine("==============");
            Console.WriteLine(LittleMiss);

            //3.Chèn giá trị vào chuỗi
            Console.WriteLine("==============");
            Console.WriteLine($"First name: {firstname}");

            //4. Chèn giá trị và hiển thị trên nhiều dòng
            string thongBao = $@"
First name: {firstname}
";
            Console.WriteLine("==============");
            Console.WriteLine(thongBao);

            //5. Độ dài của chuỗi
            Console.WriteLine("==============");
            Console.WriteLine($"String length: {said.Length}");

            //6.Chữ hoa
            Console.WriteLine("==============");
            Console.WriteLine($"Upper: {said.ToUpper()}");

            //7. Chữ thường 
            Console.WriteLine("==============");
            Console.WriteLine($"Lower: {said.ToLower()}");

            //8. Tìm chuỗi con và thay thế nó
            Console.WriteLine("==============");
            said = said.Replace("a drink", "some food");
            Console.WriteLine(said);

            //9. Chèn chuỗi vào vị trí cụ thể
            Console.WriteLine("==============");
            said = said.Insert(13, " this");
            Console.WriteLine(said);

            //10. Lấy chuỗi con, dựa trên vị trí bắt đầu và vị trí kết thưc.Nếu vị trí kết thúc không điền thì mặc định lấy đến cuối chuỗi
            Console.WriteLine("==============");
            string substring = said.Substring(24);
            Console.WriteLine(substring);

            //11.Xóa chuỗi con ,dựa trên vị trí bắt đầu và vị trí kết thúc. Nếu vị trí kết thúc không điền thì mặc định lấy đến cuối chuỗi
            Console.WriteLine("==============");
            said = said.Remove(13, 5);
            Console.WriteLine(said);

            //12. Chia chuỗi thành các chuỗi con dựa trên ký tự hoặc chuỗi
            Console.WriteLine("==============");
            string[] StrArr = said.Split(' ');
            foreach(var s in StrArr)
            {
                Console.WriteLine(s);
            }

            //13. Nối các chuỗi con ,dựa trên mảng string
            string said2 = string.Join(" ", StrArr);
            Console.WriteLine("==============");
            Console.WriteLine(said2);


            //14. Lớp StringBuiler (Dùng đối tượng này đỡ tốn bộ nhớ) * Khuyến kích nên dùng *
            StringBuilder announce = new StringBuilder();

                //14.1 thêm chuỗi
            announce.Append("Hello ");
            announce.Append("World!");

            string result = announce.ToString();
            Console.WriteLine("==============");
            Console.WriteLine(announce);

        }
    }
}


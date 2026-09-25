using System;
namespace buoi27
{
    class program
    {
        static void Main(string[] args)
        {
            //1.Hang doi
            Console.WriteLine("\t=======Queue======");
            Queue<string> hoSo = new Queue<string>();

                // Thêm váo cuối hàng đợi
            hoSo.Enqueue("Ho so 1");
            hoSo.Enqueue("Ho so 2");
            hoSo.Enqueue("Ho so 3");
            hoSo.Enqueue("Ho so 4");

                // Xóa phần tử ở đầu hàng đợi
            hoSo.Dequeue();

                //Duyệt phần tử
            Console.WriteLine("[{0}]", string.Join(", ", hoSo));

                //Độ dài hàng đợi
            Console.WriteLine($"Do dai: {hoSo.Count}");

            //2. ngăn xếp
            Console.WriteLine("\n\t=======Stack======");

            Stack<string> sach = new Stack<string>();

                //Thêm phần tử vào danh sách
            sach.Push("Sach giao khoa");
            sach.Push("Truyen");
            sach.Push("Vo ghi bai");
            sach.Push("Tai lieu");
            Console.WriteLine("[{0}]", string.Join(", ", sach));
                
                //Xóa phần tử ở đầu
            sach.Pop();
            Console.WriteLine("[{0}]", string.Join(", ", sach));

                //Độ dài ngăn xếp 
            Console.WriteLine($"Do dai: {sach.Count}");


            //3. Danh sách liên kết
            Console.WriteLine("\n\t=======LinkedList======");
            LinkedList<string> CacBaihoc = new LinkedList<string>();

                //thêm vào đầu
            var bai1 = CacBaihoc.AddFirst("Bai hoc mo dau");
            var bai2 = CacBaihoc.AddAfter(bai1, "Bai hoc 2");

                //Thêm vào cuối
            var bai5 = CacBaihoc.AddLast("Bai hoc cuoi");

                //Thêm sau nút nào đó
            LinkedListNode<string> bai3 = CacBaihoc.AddAfter(bai2, "Bai hoc 3");

                //Them đằng trước nút nào đó
            LinkedListNode<string> bai4 = CacBaihoc.AddBefore(bai5, "Bai hoc 4");

                //Duyệt phần tử
            foreach(var data in CacBaihoc)
            {
                Console.WriteLine(data);
            }

            //4. Từ điển
            Console.WriteLine("\n\t=======Dictionary======");
                //Khởi tạo
            Dictionary<string, string> TuDien = new Dictionary<string, string>()
            {
                ["Hello"]="Xin Chao",
                ["client"] = "khach hang",
                ["summer"] = "mua he",
                ["class"] = "lop",
            };
                //Thêm hoặc cập nhật
            TuDien["object"] = "Doi tuong";            
            
                //Duyệt 
            foreach(KeyValuePair<string, string> K in TuDien)
            {
                Console.WriteLine($"{K.Key} : {K.Value}");
            }

            //5. Hashset
            Console.WriteLine("\n\t=======HashSet======");
            //khởi tạo
            HashSet<int> DanhSo1 = new HashSet<int>() { 1, 2, 3, 4 };
            HashSet<int> DanhSo2 = new HashSet<int>() { 11, 12, 13,14 };
            HashSet<int> DanhSo3 = new HashSet<int>() { 11, 22, 13, 14, 25 ,70, 80 };
            
                //Thêm phần tử
            DanhSo1.Add(5);

                //Xóa phần tử theo giá trị
            DanhSo1.Remove(5);

                //Phép hợp giũa 2 tập hợp
            Console.WriteLine("Phep hop");
            DanhSo1.UnionWith(DanhSo2);
            foreach (var item in DanhSo1)
            {
                Console.WriteLine(item);
            }

            //Phép giao giữa 2 tập hợp
            Console.WriteLine("Phep giao");
            DanhSo2.IntersectWith(DanhSo3);
            foreach(var item in DanhSo2)
            {
                Console.WriteLine(item);
            }
        }
    }
}

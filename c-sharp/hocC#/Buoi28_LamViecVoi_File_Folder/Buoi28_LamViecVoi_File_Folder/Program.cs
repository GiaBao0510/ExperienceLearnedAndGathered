using System;
using System.Text;

namespace buoi28
{
    class Product
    {
        public int Id { get; set; }
        public double price { get; set; }
        public string ProductName { get; set; }

        //Lưu dữ liệu
        public void Save(Stream stream)
        {
            //int -> 4 bytes
            var byte_ID = BitConverter.GetBytes(Id);
            stream.Write(byte_ID, 0, 4);    //Lưu mảng byte_ID, với vị trị bắt đầu lưu trong mảng là 0, và tham số cuỗi là hỏi số byte cần lưu

            //double -> 8 byte
            var byte_Price = BitConverter.GetBytes(price);
            stream.Write(byte_Price, 0, 8);

            //string -> length
            var bytes_name = Encoding.UTF8.GetBytes(ProductName);
            var byte_NameLength = BitConverter.GetBytes(bytes_name.Length);//Lấy chiều dài của thuộc tính SP để biết đường sao lưu lại
            stream.Write(byte_NameLength, 0, 4);
            stream.Write(bytes_name, 0, bytes_name.Length);
        }

        //Khôi phục dữ liệu - Từ file nhị phân
        public void Restore(Stream stream)
        {
            //Đọc vào vùng đệm để phục hồi dữ liệu
            var byte_ID = new byte[4];
            stream.Read(byte_ID, 0, 4);
            Id = BitConverter.ToInt32(byte_ID, 0);

            var byte_Price = new byte[8];
            stream.Read(byte_Price, 0, 8);
            price = BitConverter.ToDouble(byte_Price, 0);

            var byte_NameLength = new byte[4];
            stream.Read(byte_NameLength, 0, 4);
            int LengthName = BitConverter.ToInt32(byte_NameLength);
            var bytes_name = new byte[LengthName];
            stream.Read(bytes_name, 0, LengthName);
            ProductName = Encoding.UTF8.GetString(bytes_name);

        }
    }

    class FileExample
    {
        //phương thức liệt kê tất cả file - thư mục dựa trên thư mục có sẵn
        static void LisFileDirectory(string path)
        {
            String[] directoies = System.IO.Directory.GetDirectories(path);
            String[] files = System.IO.Directory.GetFiles(path);
            foreach (var file in files)
            {
                Console.WriteLine(file);
            }
            foreach(var folder in directoies)
            {
                Console.WriteLine(folder);
                LisFileDirectory(folder);   //Đệ quy
            }
        }

        static void Main(string[] args)
        {
                //1. Lớp DriveInfor
            Console.WriteLine($"\t == DriveInfor ==");
            DriveInfo odia = new DriveInfo("C:");

            //Lấy danh sách tất cả ổ đĩa có sắn
            var drives = DriveInfo.GetDrives();
            foreach (var drive in drives)
            {
                Console.WriteLine("\n\t============");
                Console.WriteLine($"Drive name: {drive.Name}");
                Console.WriteLine($"Drive type: {drive.DriveType}");
                Console.WriteLine($"Drive label: {drive.VolumeLabel}");
                Console.WriteLine($"Drive format: {drive.DriveFormat}");
                Console.WriteLine($"Drive size: {drive.TotalSize}");
                Console.WriteLine($"Drive free: {drive.TotalFreeSpace}");
            }

                //2. Lớp directory
            Console.WriteLine($"\n\t == Directory ==");
            string directoryPath = "Baov2";

            //Tạo thư mục trên đường dẫn hiện tại
            Directory.CreateDirectory( directoryPath );

            //Kiểm tra thư mục có tồn tại hay không
            if (Directory.Exists(directoryPath))
            {
                Console.WriteLine($"Thu muc: {directoryPath} - Ton tai");
            }
            else Console.WriteLine($"Thu muc: {directoryPath} - Khong ton tai");

            //Xóa thư mục dựa trên đường dẫn
            Directory.Delete(directoryPath );

            //Lấy thông tin các file bên trong thư mục
            Console.WriteLine("Danh sach file:");
            var files = Directory.GetFiles("../../../obj");
            foreach (var file in files)
            {
                Console.WriteLine(file);
            }

            //Lấy thông tin danh sách thư mục con
            Console.WriteLine("\nDanh sach thu muc");
            var dsThuMuc = Directory.GetDirectories("../../../../Buoi28_LamViecVoi_File_Folder");
            foreach (var thumuc in dsThuMuc)
            {
                Console.WriteLine(thumuc);
            }

            //Tìm duyệt tất cả file thư mục
            Console.WriteLine("\nDuyet all");
            LisFileDirectory("../../../../Buoi28_LamViecVoi_File_Folder");

                //3. lớp Path
            Console.WriteLine($"\n\t == Path ==");
            Console.WriteLine($"Ky tu phan cach: {Path.DirectorySeparatorChar}");

            //Ket hop cac chuoi thanh duong dan
            var PATH = Path.Combine("../", "../", "../", "/Buoi28_LamViecVoi_File_Folder");
            var path1 = "../../../../Buoi28_LamViecVoi_File_Folder/GhiChu.txt";
            Console.WriteLine(PATH);

            //Thay đổi phần mở rộng của tên file có sản dựa trên đường dẫn
            path1 = Path.ChangeExtension(path1, "md");
            Console.WriteLine(path1);

            //Lấy tên file
            Console.WriteLine($"Ten file: { Path.GetFileName(path1)}");

            //trả về đường dẫn đầy đủ đến file có sẵn
            Console.WriteLine($"Duong daqn tuyet doi: {Path.GetFullPath(path1)}");

            //Tạo tên file ngẫu nhiên
            //-- var RandomFileName = Path.GetRandomFileName();
            //-- Console.WriteLine($"ten file vua tao: {RandomFileName}");

            //Tạo tên file ngẫu nhiên tam thời
            //-- var RandomFileNameTemp = Path.GetTempFileName();
            //-- Console.WriteLine($"ten file temp vua tao: {RandomFileNameTemp}");

                //4. lớp File
            Console.WriteLine($"\n\t == File ==");

            //Đưa từng nội dung vào file
            string HienTai = DateTime.Now.ToString("\ndd/MM/yyyy H:m:s");
            string PathFileGhiChu = "../../../GhiChu1.txt";
            File.AppendAllText(PathFileGhiChu, HienTai);

            //ĐỌc toàn bộ nội dung trong file
            string context = File.ReadAllText(PathFileGhiChu);
            Console.WriteLine($"Context: \n{context}");

            //di chuyển file đế thư mục nào đó(Ví dụ bây h là di chuyển tại thư mục hiện tại luôn)
            // -- File.Move("../../../GhiChu1.txt", "../../../GhiChu1.txt");

            //Copy file
            File.Copy("../../../GhiChu1.txt", "../../../GhiChu1_copy.txt");

            //Xóa file
            File.Delete("../../../GhiChu1_copy.txt");

                //4. lớp FileStream
            Console.WriteLine($"\n\t == FileStream ==");

            //Mở file dựa trên đường dẫn ,Nếu không có thì Tạo file 
            string pathData = "../../../Data.dat";
            using var stream = new FileStream(path: pathData, FileMode.OpenOrCreate);

            //Thực hành doc luu du lieu
            Console.WriteLine("\n\t=== Do luu du lieu");

            //Product sp1 = new Product() {
            //    Id= 12,
            //    price = 15.600,
            //    ProductName = "Ban chai"
            //};

            Product sp1 = new Product() ;

            //sp1.Save(stream);
            sp1.Restore(stream);

            Console.WriteLine($"{sp1.Id} - {sp1.ProductName} - {sp1.price}");



        }
    }
}
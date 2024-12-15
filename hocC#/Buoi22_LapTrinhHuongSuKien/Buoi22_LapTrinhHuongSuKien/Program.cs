using System;

namespace buoi22
{
    //Delegate
    public delegate void SuKienNhapSo(int x);

    //2.Kế thừ từ lớp EventArgs
    class DuLieu : EventArgs
    {
        public int data { get; set; }

        public DuLieu(int x)
        {
            this.data = x;
        } 
    }

    //1.Publisher
    class UserInput
    {
        //Thuộc tính
        public event SuKienNhapSo skns;

        //Eventhandler
        public event EventHandler SuKienNhapSo; //delegate void Kieu(object? sender, EventArgs arg)

        //Phương thức
        public void Input()
        {
            do
            {
                string s;
                Console.Write("Enter interger: ");
                s = Console.ReadLine();

                int K = Int32.Parse(s);     //Chuyển chuỗi nhập vào thành số nguyên
                SuKienNhapSo?.Invoke(this, new DuLieu(K));
            } while (true);
        }
    }

    class TinhToan
    {
        public void MuSoHai(object sender, EventArgs e)
        {
            DuLieu dulieunhap = (DuLieu)e;
            int x = dulieunhap.data;

            Console.WriteLine($"Mu so 2 cua {x}: {Math.Pow(x, 2)}");
        }

        public void CanBacHai(object sender, EventArgs e)
        {
            DuLieu dulieunhap = (DuLieu)e;
            int x = dulieunhap.data;

            Console.WriteLine($"Can bac hai cua {x}: {Math.Sqrt(x)}");
        }

        //Lấy sự kiện nhập số
        public void Sub1(UserInput user)
        {
            user.SuKienNhapSo += MuSoHai;
        }

        public void Sub2(UserInput user)
        {
            user.SuKienNhapSo += CanBacHai;
        }
    }

    class program
    {
        static void Main(string[] args)
        {
            //Bắt sự kiện khi nhấp 1 phím nào đó để hủy
            Console.CancelKeyPress += (sender, e) =>
            {
                Console.WriteLine($"\nApplication Exited.");
            };
            
            //Publisher
            UserInput user1 = new UserInput();

            //Có thể áp dụng biểu thức lambda cho delegate là event
            user1.SuKienNhapSo += (sender, e) =>
            {
                DuLieu dulieunhap = (DuLieu)e;
                int x = dulieunhap.data;

                Console.WriteLine($"Tri tuyet doi cua {x} : {Math.Abs(x)}");
            };

            //Đăng ký nhận sự kiện - subscriber
            TinhToan T1 = new TinhToan(),
                T2 = new TinhToan();
            T1.Sub1(user1);
            T1.Sub2(user1);

            user1.Input();
        }
    }
}


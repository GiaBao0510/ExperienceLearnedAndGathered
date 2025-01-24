using System;

namespace buoi21
{
    class program
    {

        static public void XinChao(string hoten)
        {
            Console.WriteLine($"Hello {hoten}");
        }

        static void Main(string[] args)
        {
            DateTime now = DateTime.Today;

            //1. Lambda & delegate
                //Có tham số đầu vào
            Action<string, string> thongbao;
            thongbao = (firtName, lastName) => Console.WriteLine($"Hello {lastName} {firtName}");
            thongbao?.Invoke("Gia Bao", "Pham");

                //Không có tham số
            Action ngaygio;
            ngaygio = () => Console.WriteLine(now.ToString("dd/MM/yyyy"));
            ngaygio.Invoke();

                //Viết lambda cách khác
            Action<string, string> bangTin;
            bangTin = (TenBangTin, thoidiem) =>
            {
                Console.ForegroundColor = ConsoleColor.Green;
                Console.BackgroundColor = ConsoleColor.White;
                Console.WriteLine($"{TenBangTin} - {thoidiem}");
                Console.ResetColor();
            };

            bangTin?.Invoke("247", now.ToString("dd/MM/yyyy"));

                //Có kiểu trả về
            Func<int, int, int> TinhTong;
            TinhTong = (a, b) =>
            {
                return a + b;
            };

            int A = 15, B = 8;
            Console.WriteLine($"{A} + {B} = {TinhTong?.Invoke(A, B)}");

            //2. Sử dụng lambda trong một số thư viện .NET

                //Sử dụng lambda để căn bậc 2 từng phần tử trong mảng
            int[] arr1 = { 3, 5, 7, 9, 4, 1, 8 , 12 ,14, 19, 22};
            var ketqua = arr1.Select((int x) =>
            {
                return Math.Pow(x, 2);
            });

            Console.WriteLine("[{0}]", string.Join(", ", ketqua));

                //Áp dụng lambda trong Foreach
            arr1.ToList().ForEach(
                (int x) =>
                {
                    if (x % 2 == 0)
                    {
                        Console.WriteLine($"So Chan: {x}");
                    }
                }
            );

            // áp dụng lambda trong where
            var ketqua2 = arr1.Where(
                (x)=>{
                    return (x % 2 != 0);
                }
            );

            Console.WriteLine("So le: [{0}]", string.Join(", ", ketqua2));
        }
    }
}

using System;


namespace buoi11
{
    class VuKhi
    {
            //Thuộc tính
        private string IDvuKhi;
        private string tenVuKhi;
        private int DoSatThuong;

        //>. get & set Thuộc tính. Trong 1 thuộc tính có cả hai set và get hoặc chỉ có 1 set hoặc get
        public string ID_VuKhi
        {
            get{
                return this.IDvuKhi;
            }
        }

        public string TenVuKhi
        {
            set
            {
                this.TenVuKhi = value;
            }
            get
            {
                return this.TenVuKhi;
            }
        }

        public int doSatThuong
        {
            set
            {
                this.DoSatThuong = value;
            }
            get
            {
                return this.DoSatThuong;
            }
        }

        //Ghi tắt thuộc tính
        public string NoiSanXuat { set; get; }

            //Phương thức

        //>1.Phương thức khởi tạo mặc định(Phương thức này thi hành đầu tiên khi được tạo ra)
        public VuKhi() {
            Console.WriteLine("Prepared weapons");
        }

        //>2. Phương thức khởi tạo có giá trị
        public VuKhi(string IDvukhi, string tenVuKhi, int dosatthuong) {
            this.tenVuKhi = tenVuKhi;
            this.IDvuKhi = IDvukhi;
            this.DoSatThuong = dosatthuong;
        }

        //Phương thức hiển thị
        public void show()
        {
            Console.WriteLine("ID: {0}. " +
                "\nWeapon name: {1}. " +
                "\nDamage: {2}." +
                "\nProduction location: {3}"
                ,this.IDvuKhi, this.tenVuKhi, this.DoSatThuong, this.NoiSanXuat
            );
        }

        //Hàm hủy
        ~VuKhi()
        {
            Console.WriteLine("Cancelled!");
        }

    }
}

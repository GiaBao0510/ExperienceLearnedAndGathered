using System;
using System.Runtime.CompilerServices;

namespace buoi14
{
    class program
    {
        //Tạo lớp cơ sở là phương tiện
        class vehicle
        {
            protected string VehicleName;
            protected string topographic;

            //Khoi tao
            public vehicle()
            {
                Console.WriteLine("Initialization Vehicle.");
            }

            public string showVehicle()
            {
                return $"Vehicle name: {this.VehicleName}. \n Topographic: {this.topographic}. \n";
            }

        }

        //Lớp con là lớp xe
        class car : vehicle {
            public int wheelnumber { set; get; }
            public string carname { set; get; }

            public car() {
                this.VehicleName = "Car";
                this.topographic = "ground";
                Console.WriteLine("Initialization Car.");
            }
            public car(int wheelnum, string carname)
            {
                this.VehicleName = "Car";
                this.topographic = "ground";
                this.carname = carname;
                this.wheelnumber = wheelnum;
            }

            //Thay thế phương thức cơ sở từ lớp cha
            public new string showVehicle()
            {
                return $"\t>>> Modified: \nVehicle name: {this.VehicleName}. \nTopographic: {this.topographic}. \n";
            }

            public void showInformation()
            {
                Console.WriteLine($"{this.showVehicle()}"); 

                //Dùng từ khóa base để truy cập đến phương thức chưa sửa đổi bởi lớp con
                Console.WriteLine($"{base.showVehicle()} Wheel number: {this.wheelnumber}. \n Car name: {this.carname}.");
            }
        }

        class bus : car
        {

        }

        class electicbus: bus
        {

        }

        //2.Lớp niêm phong: Lớp khác không tể kế thừa được
        sealed class user
        {
            private string Fname;
            private string Lname;
        }

        static void Main(string[] args)
        {
            car Xe1 = new car(4,"toyota");
            Xe1.showInformation();

            car Xe2;
            bus bus;
            electicbus Ebus;

            Xe2 = new bus();
            bus = new car();
        }
    }
}

### **Ví dụ minh họa với c#:**

**A. Triển khai Singleton với Lazy initialization**

```csharp
public class Singleton1
{
    //Biến static giữ thể hiện duy nhất
    private static Singleton1 _instance;

    //Constructor private để ngăn việc khởi tạo từ bên ngoài
    private Singleton1(){}

    //Phương thức public để truy cập thể hiện duy nhất
    public static Singleton1 GetInstance(){
        if(_instance == null){
            _instance = new Singleton1();
        }
        return _instance;
    }

    //Phương thức ví dụ
    public void DoSomething(){
        Console.WriteLine("Singleton Instance is working!");
    }
}
```

- Cách này chỉ sử dụng tốt trong trường hợp đơn luồng, trường hợp nếu có 2 luồng cùng chạy và cùng gọi đến ==hàm GetInstance()== vào trong cùng 1 thời điểm thì đương nhiên sẽ có ít nhất 2 thể hiện của instance.

**B.Triển khai Singleton với Thread-Safety:**

```csharp
public class Singleton1{

    //Biến static giữ thể hiện duy nhất
    private static Singleton1 _instance;
    private static readonly object _lock = new object();

    //Constructor private để ngăn việc khởi tạo từ bên ngoài
    private Singleton1(){}

    //Phương thức public để truy cập thể hiện duy nhất
    public static Singleton1 GetInstance(){
        if(_instance == null){
            lock(_lock){   //Đảm bảo chỉ có một luồng truy cập
                if(_instance == null){
                    _instance = new Singleton1();
                }
            }
        }
        return _instance;
    }

    //Phương thức ví dụ
    public void DoSomething(){
        Console.WriteLine("Singleton Instance is working!");
    }
}
```

- Cách này được áp dụng trong trường hợp đa luồng.

***Ví dụ quản lý cấu hình:***

```csharp
public class ConfigurationManager{

    //Biến instace là biến tính và cũng là duy nhất
    private static ConfigurationManager _instance;

    //Sử dụng khóa để đảm bảo thread-safe 
    private static readonly object _lock = new Object();

    //Dữ liệu cấu hình (giả lập)
    private Dictionary<string, string> _setting;

    //Để phương thức khởi tạo là private
    private ConfigurationManager(){
        _setting = new Dictionary<string, string>{
            { "DatabaseConnection", "Server=localhost;Database=MyApp;Trusted_Connection=True;" }, 
            { "ApiKey", "12345-abcde" }
        } 
    }

    //Phương thức static để lấy instace duy nhất
    public static ConfigurationManager Instace{
        get{

            //kiểm tra thread-safe
            if(_instance == null){
                lock(_lock){
                    if(_instance == null)
                        _instance = new ConfigurationManager();
                }
            }

            return _instance;
        }
    }

    //phương thức lấy giá trị cấu hình
    public string GetSetting(string key){
        return _setting.ContainsKey(key) ? _setting[key]: null,
    }    
}

//Sử dụng singleton
class Program
{
    static void Main(string[] args){

        //trỏ đến cùng 1 instace, đảm bảo tính nhất quán
        var config1 = ConfigurationManager.Instace;
        var config2 = ConfigurationManager.Instace;

        //Kiểm tra thử cả 2 có trỏ đến cùng 1 instance không
        Console.WriteLine($"Are both instance the same?{ReferenceEquals(config1,config2)}");

        //Lấy cấu hình
        Console.WriteLine($"DatabaseConnection: {config1.GetSetting("DatabaseConnection")}");
        Console.WriteLine($"ApiKey: {config1.GetSetting("ApiKey")}");
    }    
}
```

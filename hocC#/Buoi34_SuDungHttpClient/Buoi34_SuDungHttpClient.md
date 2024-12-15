# **Sử dụng HttpClient thực hiện các truy vấn HTTP (GET, POST)**

### **1. Lớp Uri**: 
- **_System.Uri_** là lớp biểu diễn về địa chỉ URI(URL) dùng để lấy các thông tin thành phần của URL như: host, path, query, post,... .Ngoài ra đối tượng Uri dùng để thực hiện các truy vấn HTTP Request.

### **2. Lớp tĩnh DNS và lớp IPHostEntry:**
- Lớp **DNS(_System.Net.DNS_)** cung cấp các phương thức tĩnh để lấy thông tin vầ host (Website Address, server cung cấp các dịch vụ mạng) hệ thống phân giải tên miền (DNS). Các thông tin truy vấn được trả về đối tượng giao diện **IPHostEntry.**
- ##### Một số phương thức của lớp **DNS:**
| _GetHostName()_                                          | Lấy hostname của máy local                                                                                          |
| -------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------------- |
| **_GetHostEntry(string)_** **_GetHostEntry(IPAddress)_** | Phân giải host hoặc IP thành đối tượng `IPHostEntry`. Đối tượng kiểu IPHostEntry nó chứa thông tin địa chỉ về host. |
- ##### **IPHostEntry** có các thuộc tính để lấy thông tin về host:
| **_HostName_**    | chuỗi chứa hostname của Server                 |
| ----------------- | ---------------------------------------------- |
| **_AddressList_** | Mảng các phần tử IPAddress chứa các địa chỉ IP |

### **3. Lớp Ping (System.Net,NetworkInformation.Ping):**
- Lớp này dùng để kiểm tra một máy chủ từ xa có phản hồi lại không.

### **4. Lớp HttpClient:**
- Lớp _HttpClient_ được sử dụng để gửi truy vấn HTTP (Http Message - request) và nhận phản hồi từ truy vấn đó (Http response Message). Lớp này thuộc **namespace System.Net.Http**(Chứa các lớp giúp tạo ra sự liên lạc giữa clinet - server). Làm việc với HttpClient thì thêm các namespace sau:
```
using System;
using System.Linq;
using System.Net;
using System.Net.Http;
using System.Net.Http.Headers;
using System.Threading;
using System.Threading.Tasks;
using System.IO;
using System.Text;
```

#### Tạo truy vấn GET bất đồng bộ với HttpClient
- Để tạo truy vấn GET đến một địa chỉ URL, thực hiện phương thức **_GetAsync(url)_** ,phương thức này là phương thức bất đồng bộ khi kết thúc nó trả về đối tượng **HttpResponseMessage**. Đối tượng này có thể cho biết kết quả về câu truy vấn.
- Thuộc tính **Headers** của đối tượng kiểu **HttpResponseMessage** dùng để đọc được các Header mà nó trả về
- Đọc HttpResponseMessage bằng ReadAsStreamAsync 
- Dùng để đọc từ byte 1, từng mảng byte một hoặc từng khối byte 1.
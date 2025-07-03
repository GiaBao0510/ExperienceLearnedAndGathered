**MVP (Model-View-Presenter)** là một mô hình kiến trúc lập trình gần giống như MVC. ==MVP có nhiều điểm kế thừa từ MVC nhưng Controller được thay bằng Presenter==. **MVP** được chấp nhận rộng rãi là vì mẫu kiến trúc này cung cấp tính mô-đun, khả năng kiểm thử và cơ sở mã rõ ràng và dễ bảo trì hơn.

**MVP** đưa ra sự phân tách rõ ràng hơn bằng cách đưa logic giao diện vào Presenter thay vì để **View** xử lý.

![](https://miro.medium.com/v2/resize:fit:720/format:webp/1*TuWeZzR14MmB-RBbjtZl-A.png)

---
### **Cấu trúc của MVP:**

##### **Model**
- Cũng giống như bên **MCV**, thì phần **Model** bao gồm các lớp mô tả business logic, định nghĩa business rules cho dữ liệu và giao tiếp với cơ sở dữ liệu và các tầng mạng.
- Nó chịu trách nhiệm xử lý, lưu trữ và quản lý dữ liệu cũng thực hiện mọi quy tắc kinh doanh cần thiết. Mô hình không giao tiếp trực tiếp với **View** hoặc **Presenter**.
- **Vai trò:** Quản lý dữ liệu và logic nghiệp vụ của ứng dụng. Model là nơi chứa các thực thể  (entities) và các dịch vụ (services) để truy xuất và xử lý dữ liệu
- **Chức năng:**
	- Cung cấp dữ liệu cho Present khi có yêu cầu.
	- Không liên quan trực tiếp đến **View** hoặc **Present**.
- **Ví dụ:** Trong ứng dụng quản lý khách hàng, Model sẽ quản lý các thông tin khách hàng và có thể thực hiện các thao tác như thêm, sửa, xóa khách hàng từ cơ sở dữ liệu

##### **View**
- Là thành phần mà người dùng có thể trực tiếp tương tác, gửi yêu cầu. View không bao gồm bất kỳ hoạt động xử lý logic nào.
- Tầng giao diện người dùng (**User interface**), cung cấp trực quan hóa dữ liệu và theo dõi hành động của người dùng để thông báo cho Presenter.
- **Vai trò:** Hiển thị giao diện người dùng và nhận sự kiện từ người dùng.
- **Chức năng:**
	- Hiển thị dữ liệu từ  Model (thông qua **Presenter**)
	- Chuyển các sự kiện người dùng (Như nhấn nút, nhập/ gửi dữ liệu) đến **Presenter**.
- **Tách biệt:** **View** chỉ tập trung vào việc hiển thị dữ liệu và không chứa bất kỳ logic nghiệp vụ nào. Mọi logic điều được xử lý bởi Presenter.
- **Ví dụ:** Trong ứng dụng quản lý khách hàng, View có thể là giao diện cho phép người dùng nhập thông tin khách hàng hoặc hiển thị danh sách khách hàng.

##### **Presenter**
- Đóng vai trò như **middle-man**. Bộ phận này ==nhận input của người dùng thông qua View==, xử lý dữ liệu với sự giúp sức của Model và cuối cùng trả về kết quả cho View. **Presenter** và View giao tiếp với nhau thông qua interface.
- **Vai trò:** Là trung gian giữa **Model** và **View**, chịu trách nhiệm xử lý logic nghiệp vụ và cập nhật giao diện
- **Chức năng:** 
	- Nhận sự kiện từ View và xử lý chúng.
	- Gọi Model để lấy dữ liệu hoặc thực hiện các thao tác cần thiết.
	- Cập nhật lại View với dữ liệu mới sau khi Model thực hiện xong thao tác.
- **Tách biệt:** Presenter không có sự phụ thuộc trực tiếp vào công nghệ giao diện cụ thể mà chỉ giao tiếp với View thông qua các giao diện (interface). Điều này giúp Presenter có thể tái sử dụng với nhiều loại View khác nhau.
- **Ví dụ:** Trong ứng dụng quản lý khách hàng, Presenter sẽ nhận thông tin từ View (ví dụ: tên khách hàng), thực hiện các thao tác với Model (ví dụ: thêm khách hàng vào cơ sở dữ liệu) và sau đó cập nhật View để hiển thị thông tin khách hàng vừa thêm

---
### **Cách thức hoạt động:**

![](https://caodang.fpt.edu.vn/wp-content/uploads/image5-8-768x479.png)

Trông mô hình **MVP**, **View** ==là tầng duy nhất tương tác với người dùng== (khác với **MCV**, cả 2 tầng **Controller** và **View** điều có thể tương tác với người dùng).

Trong khi **View** đảm nhận trình bày thì **Presenter** đảm trách cách **Model** được thao tác và thay đổi như thế nào bởi giao diện người dùng. **Presenter** là nơi chứa các xử lý đặc trưng của ứng dụng (**application logic** so với **business logic** của **Model**). Một điểm đáng chú ý là **Presenter** có khả năng thao tác trực tiếp lên **View** mà nó gắn kết.

Luồng đi của mô hình MVP như sau:
- **View** nhận tương tác từ người dùng  truyền đến $\to$ **Presenter** và xử lý thông qua **Model**. Sau khi **Model** thực hiện thay đổi dữ liệu xong, **Presenter** lấy dữ liệu tư Model để truyền đến cho **View**.

![](https://s3.ap-southeast-1.amazonaws.com/techover.storage/wp-content/uploads/2022/12/07224617/MVP-architecture.png)

---
### **Luồng dữ liệu (Data Flow):**

- Người dùng tương tác với **View**:
	- Người dùng thực hiện hành động trong giao diện (nhấn nút gửi, nhập dữ liệu).
	- View gửi sự kiện đến **presenter**.
- **Presenter** xử lý dữ liệu:
	- **Presenter** nhận sự kiện từ **View** để xử lý logic nghiệp vụ.
	- **Presenter** yêu cầu **Model** thực hiện các tác vụ cần thiết (truy xuất, cập nhật, dữ liệu).
- **Model** trả về kết quả:
	- **Model** thực hiện các tác vụ và trả lại dữ liệu cho **Presenter**.
- **Presenter** cập nhật lại View:
	- **Presenter** quyết định các thực hiển thị dữ liệu và gửi kết quả cho **View**.
	- **View** hiển thị dữ liệu cho người dùng.

---
### **Ưu/nhược điểm:**

##### **Ưu điểm:**
- **MVP** có cấu trúc Code rõ ràng hơn so với MVC
- Dễ dàng viết unit test cho **Presenter** vì nó hoạt động độc lập với **View**.
- Cải thiện sự phân tách mối quan tâm giữa chế độ xem và mô hình.
- **Presenter** tạo điều kiện cho khả năng kiểm tra và mô đun hóa tốt hơn.
- Mỗi thành phấn có thể được sửa đổi hoặc thay thế mà không ảnh hướng đến những thành phần khác.

##### **Nhược điểm:**
- Độ phức tạp tự tăng lên so với MCV truyền thống, do lớp **Presenter** có thêm trách nhiệm (**Presenter** của **MVP** sẽ càng ngày phình to nếu logic được thêm mới). Vì vậy rất khó để chia nhỏ khi **presenter** quá lớn.
- Có thể dẫn đến một cơ sở mã lớn hơn và cần nhiều mã soặn sẵn hơn.

---
#### ***So sánh MCV với MVP:***

| MVC (Model-View-Controller)                                                                                                                           | MVP (Model-View-Presenter)                                                                                                            |
| ----------------------------------------------------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------------------------------- |
| Một trong những kiến trúc phần mềm lâu đời nhất                                                                                                       | Được phát triển như là phiên bản thứ hai của kiến trúc phần mềm, cải tiến từ MVC                                                      |
| Giao diện người dùng (View) và cơ chế truy cập dữ liệu (Model) được kết hợp chặt chẽ với nhau.                                                        | Giải quyết vấn đề phụ thuộc View bằng cách sử dụng **Presenter** (Người thuyết trình) làm kênh liên lạc giữa Model và View.           |
| UI (**View**) và cơ chế truy cập dữ liệu (**Model**) được kết hợp chặt chẽ với nhau                                                                   | Nó giải quyết vấn đề sự phụ thuộc của **View** bằng cách áp dụng **Present** như là một kênh trung gian giữa **Model** và **View**    |
| **Controller** và View tồn tại với mối quan hệ một-nhiều (One-to-many). Một **Controller** có thể chọn một View khác dựa trên hoạt động được yêu cầu. | Mối quan hệ một-một (one-to-one) tồn tại giữa **Presenter** và **View** khi lớp **Presenter** quản lý một **View** tại một thời điểm. |
| **View** không có kiến ​​thức về **Controller**.                                                                                                      | **View** có tham chiếu đến **Presenter**                                                                                              |
| Khó thực hiện thay đổi và sửa đổi các tính năng của ứng dụng vì các tầng mã được liên kết chặt chẽ với nhau.                                          | Các tầng mã được liên kết lỏng lẻo và do đó dễ dàng thực hiện các sửa đổi/thay đổi trong mã ứng dụng.                                 |
| Đầu vào người dùng được xử lý bởi Controller                                                                                                          | View là điểm vào của Ứng dụng                                                                                                         |
| Chỉ lý tưởng cho các dự án quy mô nhỏ.                                                                                                                | Lý tưởng cho các ứng dụng đơn giản và phức tạp.                                                                                       |
| Hỗ trợ hạn chế cho kiểm thử Đơn vị.                                                                                                                   | Dễ dàng thực hiện kiểm thử Đơn vị nhưng sự liên kết chặt chẽ giữa View và Presenter có thể gây khó khăn đôi chút.                     |
| không tuân theo nguyên tắc trách nhiệm đơn lẻ và mô đun hóa.                                                                                          | Tuân theo nguyên tắc trách nhiệm đơn lẻ và mô đun hóa.                                                                                |

---
### **Ví dụ:** về bài toán quản lý khách hàng

#### ***Cây thư mục:***
```text
StudentManagement/
├── Models/
│   ├── Student.cs                    # Entity
│   └── StudentModel.cs               # Business Logic & Data Access
├── Views/
│   ├── IStudentView.cs              # View Interface
│   └── WinForms/
│       └── StudentForm.cs            # Concrete View Implementation
├── Presenters/
│   └── StudentPresenter.cs          # Presenter Logic
├── Tests/
│   └── StudentPresenterTests.cs     # Unit Tests
└── Program.cs                       # Entry Point
```


#### ***Mã triển triển khai:***

 ##### **MODEL:**
```csharp
using System;
using System.Collections.Generic;
using System.Linq;

namespace MVPExample.Models
{
    // Entity - Sinh viên
    public class Student
    {
        public int Id { get; set; }
        public string Name { get; set; }
        public string Email { get; set; }
        public DateTime DateOfBirth { get; set; }
        public string Major { get; set; }
        
        public int Age => DateTime.Now.Year - DateOfBirth.Year;
    }
    
    // Model - Chứa business logic và data access
    public class StudentModel
    {
        private List<Student> _students;
        
        public StudentModel()
        {
            // Khởi tạo dữ liệu mẫu
            _students = new List<Student>
            {
                new Student { Id = 1, Name = "Nguyễn Văn A", Email = "a@gmail.com", DateOfBirth = new DateTime(2000, 5, 15), Major = "Công nghệ thông tin" },
                new Student { Id = 2, Name = "Trần Thị B", Email = "b@gmail.com", DateOfBirth = new DateTime(1999, 8, 20), Major = "Kinh tế" },
                new Student { Id = 3, Name = "Lê Văn C", Email = "c@gmail.com", DateOfBirth = new DateTime(2001, 3, 10), Major = "Kỹ thuật" }
            };
        }
        
        // Lấy tất cả sinh viên
        public List<Student> GetAllStudents()
        {
            return _students.ToList();
        }
        
        // Lấy sinh viên theo ID
        public Student GetStudentById(int id)
        {
            return _students.FirstOrDefault(s => s.Id == id);
        }
        
        // Thêm sinh viên mới
        public bool AddStudent(Student student)
        {
            try
            {
                // Validate business rules
                if (string.IsNullOrEmpty(student.Name))
                    throw new ArgumentException("Tên sinh viên không được để trống");
                
                if (string.IsNullOrEmpty(student.Email))
                    throw new ArgumentException("Email không được để trống");
                
                if (_students.Any(s => s.Email == student.Email))
                    throw new ArgumentException("Email đã tồn tại");
                
                // Tự động tạo ID
                student.Id = _students.Max(s => s.Id) + 1;
                _students.Add(student);
                
                return true;
            }
            catch
            {
                return false;
            }
        }
        
        // Cập nhật sinh viên
        public bool UpdateStudent(Student student)
        {
            try
            {
                var existingStudent = _students.FirstOrDefault(s => s.Id == student.Id);
                if (existingStudent == null)
                    return false;
                
                // Validate business rules
                if (string.IsNullOrEmpty(student.Name))
                    throw new ArgumentException("Tên sinh viên không được để trống");
                
                if (string.IsNullOrEmpty(student.Email))
                    throw new ArgumentException("Email không được để trống");
                
                // Kiểm tra email trùng (trừ chính nó)
                if (_students.Any(s => s.Email == student.Email && s.Id != student.Id))
                    throw new ArgumentException("Email đã tồn tại");
                
                // Cập nhật thông tin
                existingStudent.Name = student.Name;
                existingStudent.Email = student.Email;
                existingStudent.DateOfBirth = student.DateOfBirth;
                existingStudent.Major = student.Major;
                
                return true;
            }
            catch
            {
                return false;
            }
        }
        
        // Xóa sinh viên
        public bool DeleteStudent(int id)
        {
            var student = _students.FirstOrDefault(s => s.Id == id);
            if (student == null)
                return false;
            
            _students.Remove(student);
            return true;
        }
        
        // Tìm kiếm sinh viên theo tên
        public List<Student> SearchStudents(string keyword)
        {
            if (string.IsNullOrEmpty(keyword))
                return GetAllStudents();
            
            return _students.Where(s => 
                s.Name.ToLower().Contains(keyword.ToLower()) ||
                s.Email.ToLower().Contains(keyword.ToLower()) ||
                s.Major.ToLower().Contains(keyword.ToLower())
            ).ToList();
        }
    }
}
```

##### **VIEW INTERFACE:**
```csharp
using System;
using System.Collections.Generic;

namespace MVPExample.Views
{
    // Interface định nghĩa contract giữa View và Presenter
    public interface IStudentView
    {
        // Properties để Presenter có thể đọc/ghi dữ liệu
        string StudentName { get; set; }
        string StudentEmail { get; set; }
        DateTime StudentDateOfBirth { get; set; }
        string StudentMajor { get; set; }
        string SearchKeyword { get; set; }
        
        // Methods để Presenter cập nhật UI
        void DisplayStudents(List<Student> students);
        void DisplayMessage(string message);
        void DisplayError(string error);
        void ClearForm();
        void ShowStudentDetails(Student student);
        
        // Events để thông báo cho Presenter
        event Action LoadStudents;
        event Action<int> SelectStudent;
        event Action AddStudent;
        event Action<int> UpdateStudent;
        event Action<int> DeleteStudent;
        event Action<string> SearchStudents;
    }
}
```

##### **PRESENTER**
```csharp
using System;
using System.Collections.Generic;
using MVPExample.Models;
using MVPExample.Views;

namespace MVPExample.Presenters
{
    public class StudentPresenter
    {
        private readonly IStudentView _view;
        private readonly StudentModel _model;
        private Student _currentStudent;
        
        public StudentPresenter(IStudentView view, StudentModel model)
        {
            _view = view;
            _model = model;
            
            // Đăng ký event handlers
            _view.LoadStudents += OnLoadStudents;
            _view.SelectStudent += OnSelectStudent;
            _view.AddStudent += OnAddStudent;
            _view.UpdateStudent += OnUpdateStudent;
            _view.DeleteStudent += OnDeleteStudent;
            _view.SearchStudents += OnSearchStudents;
        }
        
        // Xử lý sự kiện load danh sách sinh viên
        private void OnLoadStudents()
        {
            try
            {
                var students = _model.GetAllStudents();
                _view.DisplayStudents(students);
                _view.DisplayMessage($"Đã tải {students.Count} sinh viên");
            }
            catch (Exception ex)
            {
                _view.DisplayError($"Lỗi khi tải danh sách: {ex.Message}");
            }
        }
        
        // Xử lý sự kiện chọn sinh viên
        private void OnSelectStudent(int studentId)
        {
            try
            {
                _currentStudent = _model.GetStudentById(studentId);
                if (_currentStudent != null)
                {
                    _view.ShowStudentDetails(_currentStudent);
                    _view.DisplayMessage($"Đã chọn sinh viên: {_currentStudent.Name}");
                }
                else
                {
                    _view.DisplayError("Không tìm thấy sinh viên");
                }
            }
            catch (Exception ex)
            {
                _view.DisplayError($"Lỗi khi chọn sinh viên: {ex.Message}");
            }
        }
        
        // Xử lý sự kiện thêm sinh viên
        private void OnAddStudent()
        {
            try
            {
                var newStudent = new Student
                {
                    Name = _view.StudentName,
                    Email = _view.StudentEmail,
                    DateOfBirth = _view.StudentDateOfBirth,
                    Major = _view.StudentMajor
                };
                
                if (_model.AddStudent(newStudent))
                {
                    _view.DisplayMessage("Thêm sinh viên thành công!");
                    _view.ClearForm();
                    OnLoadStudents(); // Refresh danh sách
                }
                else
                {
                    _view.DisplayError("Không thể thêm sinh viên. Vui lòng kiểm tra thông tin.");
                }
            }
            catch (Exception ex)
            {
                _view.DisplayError($"Lỗi khi thêm sinh viên: {ex.Message}");
            }
        }
        
        // Xử lý sự kiện cập nhật sinh viên
        private void OnUpdateStudent(int studentId)
        {
            try
            {
                var student = _model.GetStudentById(studentId);
                if (student != null)
                {
                    student.Name = _view.StudentName;
                    student.Email = _view.StudentEmail;
                    student.DateOfBirth = _view.StudentDateOfBirth;
                    student.Major = _view.StudentMajor;
                    
                    if (_model.UpdateStudent(student))
                    {
                        _view.DisplayMessage("Cập nhật sinh viên thành công!");
                        _view.ClearForm();
                        OnLoadStudents(); // Refresh danh sách
                    }
                    else
                    {
                        _view.DisplayError("Không thể cập nhật sinh viên. Vui lòng kiểm tra thông tin.");
                    }
                }
                else
                {
                    _view.DisplayError("Không tìm thấy sinh viên cần cập nhật");
                }
            }
            catch (Exception ex)
            {
                _view.DisplayError($"Lỗi khi cập nhật sinh viên: {ex.Message}");
            }
        }
        
        // Xử lý sự kiện xóa sinh viên
        private void OnDeleteStudent(int studentId)
        {
            try
            {
                var student = _model.GetStudentById(studentId);
                if (student != null)
                {
                    if (_model.DeleteStudent(studentId))
                    {
                        _view.DisplayMessage($"Đã xóa sinh viên: {student.Name}");
                        _view.ClearForm();
                        OnLoadStudents(); // Refresh danh sách
                    }
                    else
                    {
                        _view.DisplayError("Không thể xóa sinh viên");
                    }
                }
                else
                {
                    _view.DisplayError("Không tìm thấy sinh viên cần xóa");
                }
            }
            catch (Exception ex)
            {
                _view.DisplayError($"Lỗi khi xóa sinh viên: {ex.Message}");
            }
        }
        
        // Xử lý sự kiện tìm kiếm sinh viên
        private void OnSearchStudents(string keyword)
        {
            try
            {
                var students = _model.SearchStudents(keyword);
                _view.DisplayStudents(students);
                _view.DisplayMessage($"Tìm thấy {students.Count} sinh viên");
            }
            catch (Exception ex)
            {
                _view.DisplayError($"Lỗi khi tìm kiếm: {ex.Message}");
            }
        }
        
        // Method để khởi tạo view
        public void InitializeView()
        {
            OnLoadStudents();
        }
    }
}
```

##### **CONCRETE VIEW (Windows Forms)**
```csharp
using System;
using System.Collections.Generic;
using System.Linq;
using System.Windows.Forms;
using MVPExample.Models;
using MVPExample.Views;

namespace MVPExample.WinForms
{
    public partial class StudentForm : Form, IStudentView
    {
        // Controls (giả sử có các control này trong form)
        private TextBox txtName;
        private TextBox txtEmail;
        private DateTimePicker dtpDateOfBirth;
        private TextBox txtMajor;
        private TextBox txtSearch;
        private ListBox lstStudents;
        private Button btnAdd;
        private Button btnUpdate;
        private Button btnDelete;
        private Button btnSearch;
        private Label lblMessage;
        
        // Implement IStudentView properties
        public string StudentName
        {
            get => txtName.Text;
            set => txtName.Text = value;
        }
        
        public string StudentEmail
        {
            get => txtEmail.Text;
            set => txtEmail.Text = value;
        }
        
        public DateTime StudentDateOfBirth
        {
            get => dtpDateOfBirth.Value;
            set => dtpDateOfBirth.Value = value;
        }
        
        public string StudentMajor
        {
            get => txtMajor.Text;
            set => txtMajor.Text = value;
        }
        
        public string SearchKeyword
        {
            get => txtSearch.Text;
            set => txtSearch.Text = value;
        }
        
        // Implement IStudentView events
        public event Action LoadStudents;
        public event Action<int> SelectStudent;
        public event Action AddStudent;
        public event Action<int> UpdateStudent;
        public event Action<int> DeleteStudent;
        public event Action<string> SearchStudents;
        
        public StudentForm()
        {
            InitializeComponent();
            SetupEventHandlers();
        }
        
        private void SetupEventHandlers()
        {
            // Khi form load, trigger LoadStudents event
            this.Load += (s, e) => LoadStudents?.Invoke();
            
            // Khi chọn sinh viên trong danh sách
            lstStudents.SelectedIndexChanged += (s, e) =>
            {
                if (lstStudents.SelectedItem is Student student)
                {
                    SelectStudent?.Invoke(student.Id);
                }
            };
            
            // Khi click các button
            btnAdd.Click += (s, e) => AddStudent?.Invoke();
            btnUpdate.Click += (s, e) =>
            {
                if (lstStudents.SelectedItem is Student student)
                {
                    UpdateStudent?.Invoke(student.Id);
                }
            };
            btnDelete.Click += (s, e) =>
            {
                if (lstStudents.SelectedItem is Student student)
                {
                    DeleteStudent?.Invoke(student.Id);
                }
            };
            btnSearch.Click += (s, e) => SearchStudents?.Invoke(SearchKeyword);
        }
        
        // Implement IStudentView methods
        public void DisplayStudents(List<Student> students)
        {
            lstStudents.DataSource = null;
            lstStudents.DataSource = students;
            lstStudents.DisplayMember = "Name";
        }
        
        public void DisplayMessage(string message)
        {
            lblMessage.Text = message;
            lblMessage.ForeColor = System.Drawing.Color.Green;
        }
        
        public void DisplayError(string error)
        {
            lblMessage.Text = error;
            lblMessage.ForeColor = System.Drawing.Color.Red;
        }
        
        public void ClearForm()
        {
            txtName.Clear();
            txtEmail.Clear();
            dtpDateOfBirth.Value = DateTime.Now;
            txtMajor.Clear();
        }
        
        public void ShowStudentDetails(Student student)
        {
            StudentName = student.Name;
            StudentEmail = student.Email;
            StudentDateOfBirth = student.DateOfBirth;
            StudentMajor = student.Major;
        }
    }
}
```

##### **PROGRAM - MAIN ENTRY POINT**
```csharp
using System;
using System.Windows.Forms;
using MVPExample.Models;
using MVPExample.Views;
using MVPExample.Presenters;
using MVPExample.WinForms;

namespace MVPExample
{
    class Program
    {
        [STAThread]
        static void Main()
        {
            Application.EnableVisualStyles();
            Application.SetCompatibleTextRenderingDefault(false);
            
            // Khởi tạo các thành phần theo MVP pattern
            var model = new StudentModel();
            var view = new StudentForm();
            var presenter = new StudentPresenter(view, model);
            
            // Khởi tạo view và chạy ứng dụng
            presenter.InitializeView();
            Application.Run(view);
        }
    }
}
```

##### **UNIT TEST EXAMPLE**
```csharp
using Microsoft.VisualStudio.TestTools.UnitTesting;
using MVPExample.Models;
using MVPExample.Views;
using MVPExample.Presenters;
using System;
using System.Collections.Generic;

namespace MVPExample.Tests
{
    // Mock View để test Presenter
    public class MockStudentView : IStudentView
    {
        public string StudentName { get; set; }
        public string StudentEmail { get; set; }
        public DateTime StudentDateOfBirth { get; set; }
        public string StudentMajor { get; set; }
        public string SearchKeyword { get; set; }
        
        public List<Student> DisplayedStudents { get; private set; }
        public string LastMessage { get; private set; }
        public string LastError { get; private set; }
        public bool IsFormCleared { get; private set; }
        public Student DisplayedStudentDetails { get; private set; }
        
        public event Action LoadStudents;
        public event Action<int> SelectStudent;
        public event Action AddStudent;
        public event Action<int> UpdateStudent;
        public event Action<int> DeleteStudent;
        public event Action<string> SearchStudents;
        
        public void DisplayStudents(List<Student> students)
        {
            DisplayedStudents = students;
        }
        
        public void DisplayMessage(string message)
        {
            LastMessage = message;
        }
        
        public void DisplayError(string error)
        {
            LastError = error;
        }
        
        public void ClearForm()
        {
            IsFormCleared = true;
        }
        
        public void ShowStudentDetails(Student student)
        {
            DisplayedStudentDetails = student;
        }
        
        // Helper methods để trigger events
        public void TriggerLoadStudents() => LoadStudents?.Invoke();
        public void TriggerSelectStudent(int id) => SelectStudent?.Invoke(id);
        public void TriggerAddStudent() => AddStudent?.Invoke();
        public void TriggerUpdateStudent(int id) => UpdateStudent?.Invoke(id);
        public void TriggerDeleteStudent(int id) => DeleteStudent?.Invoke(id);
        public void TriggerSearchStudents(string keyword) => SearchStudents?.Invoke(keyword);
    }
    
    [TestClass]
    public class StudentPresenterTests
    {
        private MockStudentView _mockView;
        private StudentModel _model;
        private StudentPresenter _presenter;
        
        [TestInitialize]
        public void Setup()
        {
            _mockView = new MockStudentView();
            _model = new StudentModel();
            _presenter = new StudentPresenter(_mockView, _model);
        }
        
        [TestMethod]
        public void LoadStudents_ShouldDisplayAllStudents()
        {
            // Act
            _mockView.TriggerLoadStudents();
            
            // Assert
            Assert.IsNotNull(_mockView.DisplayedStudents);
            Assert.IsTrue(_mockView.DisplayedStudents.Count > 0);
            Assert.IsTrue(_mockView.LastMessage.Contains("Đã tải"));
        }
        
        [TestMethod]
        public void AddStudent_WithValidData_ShouldSucceed()
        {
            // Arrange
            _mockView.StudentName = "Test Student";
            _mockView.StudentEmail = "test@example.com";
            _mockView.StudentDateOfBirth = new DateTime(2000, 1, 1);
            _mockView.StudentMajor = "Test Major";
            
            // Act
            _mockView.TriggerAddStudent();
            
            // Assert
            Assert.IsTrue(_mockView.LastMessage.Contains("thành công"));
            Assert.IsTrue(_mockView.IsFormCleared);
        }
        
        [TestMethod]
        public void SelectStudent_WithValidId_ShouldShowStudentDetails()
        {
            // Act
            _mockView.TriggerSelectStudent(1);
            
            // Assert
            Assert.IsNotNull(_mockView.DisplayedStudentDetails);
            Assert.AreEqual(1, _mockView.DisplayedStudentDetails.Id);
        }
    }
}
```


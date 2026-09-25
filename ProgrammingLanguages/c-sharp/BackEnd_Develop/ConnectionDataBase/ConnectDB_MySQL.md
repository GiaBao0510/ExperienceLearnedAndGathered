Hướng dẫn cách thiết lập kết nối đến cơ sở dữ liệu MySQL trong ứng dụng .NET, sử dụng `MySql.Data.MySqlClient` và tích hợp với Dapper để truy vấn dữ liệu. Tài liệu phù hợp cho các ứng dụng web (như E-commerce API) hoặc bất kỳ dự án .NET nào cần kết nối MySQL.

---
## Yêu cầu

Trước khi bắt đầu, đảm bảo dự án của bạn đã cài đặt các gói NuGet sau:

- **Microsoft.Extensions.Configuration**: Quản lý cấu hình (như chuỗi kết nối từ appsettings.json).
- **MySql.Data.MySqlClient**: Thư viện kết nối MySQL.

```
dotnet add package Microsoft.Extensions.Configuration
dotnet add package MySql.Data
```

---
## Cấu hình chuỗi kết nối

Chuỗi kết nối đến MySQL được khai báo trong tệp appsettings.json. Ví dụ

```json
"DataBase":{
	"MySQL":"Server=localhost; Port=3306; Database=MyDataBase; Uid = root; Pwd=MyPassword; Connection Timeout=15, DefaultCommandTimeout=30; AllowLoadLocalInfine=false; Pooling=true; Min Pool Size=10; Max Pool Size=20; ConnctionLifeTime=300; ConnectionReset=true;"
} 
```

### Giải thích tham số

| Tham số                   | Ý nghĩa                                                                  | Giá trị đề xuất                         |
| ------------------------- | ------------------------------------------------------------------------ | --------------------------------------- |
| **Server**                | Địa chỉ máy chủ MySQL (host).                                            | localhost hoặc IP (như 127.0.0.1)       |
| **Port**                  | Cổng MySQL lắng nghe kết nối.                                            | 3306 (mặc định)                         |
| **Database**              | Tên cơ sở dữ liệu cần kết nối.                                           | Tên database cụ thể (như MyDatabase)    |
| **Uid**                   | Tên người dùng MySQL.                                                    | root hoặc tài khoản khác                |
| **Pwd**                   | Mật khẩu của người dùng.                                                 | Mật khẩu an toàn                        |
| **ConnectionTimeout**     | Thời gian tối đa (giây) chờ kết nối đến MySQL.                           | 15 (tăng lên 30 nếu mạng không ổn định) |
| **DefaultCommandTimeout** | Thời gian tối đa (giây) cho một lệnh SQL trước khi bị hủy.               | 30                                      |
| **AllowLoadLocalInfile**  | Cho phép dùng lệnh LOAD DATA LOCAL INFILE để nhập dữ liệu từ tệp cục bộ. | false (vì lý do bảo mật)                |
| **Pooling**               | Bật/tắt cơ chế connection pooling (tái sử dụng kết nối).                 | true (để tăng hiệu suất)                |
| **Min Pool Size**         | Số kết nối tối thiểu trong pool khi khởi động.                           | 10 (đủ cho ứng dụng nhỏ)                |
| **Max Pool Size**         | Số kết nối tối đa trong pool.                                            | 20 (tăng nếu tải cao)                   |
| **ConnectionLifeTime**    | Thời gian tối đa (giây) một kết nối được giữ trước khi làm mới.          | 300 (5 phút, tùy chỉnh theo tải)        |
| **ConnectionReset**       | Đặt lại trạng thái kết nối khi tái sử dụng từ pool (đảm bảo nhất quán).  | true (bật để an toàn)                   |

---
## Triển khai lớp DatabaseConnectionFactory

Lớp DatabaseConnectionFactory giúp:

- Tạo kết nối đến MySQL với cơ chế retry khi gặp lỗi tạm thời.
- Khởi động connection pool để cải thiện hiệu suất.
- Kiểm tra tính hợp lệ của kết nối.

Tạo tệp `DatabaseConnectionFactory.cs` trong thư mục `Infrastructure/Data` với nội dung sau:

```csharp
using System;
using System.Collections.Generic;
using System.Data;
using System.Threading.Tasks;
using Microsoft.Extensions.Logging;
using MySql.Data.MySqlClient;

namespace E_commerce.Infrastructure.Data
{
    public class DatabaseConnectionFactory
    {
        private readonly string _connectionString;
        private readonly ILogger<DatabaseConnectionFactory> _logger;
        private readonly int _retryCount = 3; // Số lần thử lại khi kết nối thất bại
        private readonly int _retryDelayMs = 500; // Độ trễ giữa các lần thử (ms)

        // Mã lỗi MySQL có thể thử lại (chủ yếu là lỗi kết nối tạm thời)
        private static readonly HashSet<int> RetryableErrorCodes = new HashSet<int>
        {
            1042, // Không thể kết nối đến máy chủ MySQL
            1043, // Lỗi bắt tay (handshake)
            1045, // Sai thông tin đăng nhập
            1153, // Packet quá lớn
            1158, // Lỗi đọc packet
            1159, // Timeout khi đọc packet
            2002, // Không thể kết nối đến máy chủ
            2003, // Không thể kết nối đến máy chủ
            2004, // Không thể kết nối đến máy chủ
            2005, // Máy chủ không xác định
            2006, // Máy chủ ngắt kết nối
            2013  // Mất kết nối trong khi truy vấn
        };

        /// <summary>
        /// Khởi tạo factory với chuỗi kết nối và logger.
        /// </summary>
        public DatabaseConnectionFactory(IConfiguration configuration, ILogger<DatabaseConnectionFactory> logger)
        {
            _connectionString = configuration["Database:MySQL"]
                ?? throw new InvalidOperationException("Connection string 'Database:MySQL' is not configured.");
            _logger = logger ?? throw new ArgumentNullException(nameof(logger));
        }

        /// <summary>
        /// Tạo kết nối đồng bộ đến MySQL với cơ chế retry.
        /// </summary>
        /// <returns>Kết nối MySQL đã mở.</returns>
        public IDbConnection CreateConnection()
        {
            for (int attempt = 0; attempt < _retryCount; attempt++)
            {
                try
                {
                    var connection = new MySqlConnection(_connectionString);
                    // Mở kết nối nếu chưa mở
                    if (connection.State != ConnectionState.Open)
                    {
                        connection.Open();
                    }
                    _logger.LogInformation("Successfully connected to MySQL database.");
                    return connection;
                }
                catch (MySqlException ex)
                {
                    bool shouldRetry = RetryableErrorCodes.Contains(ex.Number);

                    // Không retry nếu đã hết số lần thử hoặc lỗi không thể retry
                    if (attempt == _retryCount - 1 || !shouldRetry)
                    {
                        _logger.LogError(ex, $"Failed to connect to database after {attempt + 1} attempts: {ex.Message}");
                        throw new InvalidOperationException("Database connection failed.", ex);
                    }

                    _logger.LogWarning($"Connection attempt {attempt + 1} failed: {ex.Message}. Retrying in {_retryDelayMs}ms...");
                    Thread.Sleep(_retryDelayMs);
                }
            }

            throw new TimeoutException($"Failed to connect to database after {_retryCount} attempts.");
        }

        /// <summary>
        /// Tạo kết nối bất đồng bộ đến MySQL với cơ chế retry.
        /// </summary>
        /// <returns>Kết nối MySQL đã mở.</returns>
        public async Task<IDbConnection> CreateConnectionAsync()
        {
            for (int attempt = 0; attempt < _retryCount; attempt++)
            {
                try
                {
                    var connection = new MySqlConnection(_connectionString);
                    // Mở kết nối bất đồng bộ nếu chưa mở
                    if (connection.State != ConnectionState.Open)
                    {
                        await connection.OpenAsync();
                    }
                    _logger.LogInformation("Successfully connected to MySQL database (async).");
                    return connection;
                }
                catch (MySqlException ex)
                {
                    bool shouldRetry = RetryableErrorCodes.Contains(ex.Number);

                    // Không retry nếu đã hết số lần thử hoặc lỗi không thể retry
                    if (attempt == _retryCount - 1 || !shouldRetry)
                    {
                        _logger.LogError(ex, $"Failed to connect to database after {attempt + 1} attempts: {ex.Message}");
                        throw new InvalidOperationException("Database connection failed.", ex);
                    }

                    _logger.LogWarning($"Connection attempt {attempt + 1} failed: {ex.Message}. Retrying in {_retryDelayMs}ms...");
                    await Task.Delay(_retryDelayMs); // Dùng Task.Delay thay vì Thread.Sleep
                }
            }

            throw new TimeoutException($"Failed to connect to database after {_retryCount} attempts.");
        }

        /// <summary>
        /// Kiểm tra xem kết nối có hợp lệ không bằng cách chạy truy vấn thử.
        /// </summary>
        /// <param name="connection">Kết nối cần kiểm tra.</param>
        /// <returns>True nếu kết nối hợp lệ, False nếu không.</returns>
        public bool ValidateConnection(IDbConnection connection)
        {
            if (connection == null || connection.State != ConnectionState.Open)
            {
                _logger.LogWarning("Connection is null or not open.");
                return false;
            }

            try
            {
                using var command = connection.CreateCommand();
                command.CommandText = "SELECT 1";
                command.CommandTimeout = 5; // Timeout sau 5 giây
                var result = command.ExecuteScalar();
                return result != null;
            }
            catch (Exception ex)
            {
                _logger.LogError(ex, $"Connection validation failed: {ex.Message}");
                return false;
            }
        }

        /// <summary>
        /// Khởi động connection pool với số kết nối tối thiểu (10) để cải thiện hiệu suất ban đầu.
        /// </summary>
        public void WarmupConnectionPool()
        {
            try
            {
                const int minPoolSize = 10; // Phải khớp với Min Pool Size trong chuỗi kết nối
                var connections = new List<IDbConnection>();

                _logger.LogInformation($"Warming up connection pool with {minPoolSize} connections...");

                // Tạo các kết nối để làm đầy pool
                for (int i = 0; i < minPoolSize; i++)
                {
                    var connection = CreateConnection();
                    connections.Add(connection);
                }

                _logger.LogInformation($"Successfully established {connections.Count} initial connections.");

                // Giải phóng kết nối để trả về pool
                foreach (var connection in connections)
                {
                    connection.Dispose();
                }
            }
            catch (Exception ex)
            {
                _logger.LogError(ex, $"Failed to warm up connection pool: {ex.Message}");
            }
        }
    }
}
```

---
## Đăng ký dịch vụ

Đăng ký DatabaseConnectionFactory trong Program.cs với lifetime là **Singleton** để đảm bảo chỉ tạo một instance duy nhất:

```csharp
//Đăng ký DatabaseConnectionFactory cho Dapper
builder.Services.AddSingleton<DatabaseConnectionFactory>();
```

---
## Lưu ý

- **Bảo mật**: Không lưu mật khẩu MySQL trực tiếp trong appsettings.json ở môi trường production. Sử dụng biến môi trường hoặc Azure Key Vault.
- **Hiệu suất**: Giữ Pooling=true và điều chỉnh Min Pool Size, Max Pool Size dựa trên tải ứng dụng.
- **Retry**: Điều chỉnh _retryCount và _retryDelayMs nếu ứng dụng chạy trong môi trường mạng không ổn định.
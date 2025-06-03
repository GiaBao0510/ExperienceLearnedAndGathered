Sẽ có thư viện như sau:
```shell

```

Từ khóa **Bulk**

Trong thư viện MySqlConnector sẽ bao gồm lớp `MySqlBulkCopy` cho phép để tải dữ liệu từ một bảng trong MySQL Server một cách hiệu quả. Nó cũng tương tự như lớp SqlBulkCopy trong thư viện SQL Server.

Do một số tính năng bảo mật trong MySQL Server, phần tham số trong chuỗi kết nối phải thiết lập thuộc tính `AllowLoadLocalInfile=true`.

Trong phần này có liên quan đến lớp `MySqlBulkLoader`. cho phép tải dữ liệu từ bảng MySQL Server với định dạng tệp tin là .CSV hoặc TSV hoặc luồng (Stream),

---

### thư viện cần tải

```powershell
 dotnet add package MySqlConnector --version 2.4.0
```

  

# Tạo interface mô tả các chức năng xử lý file

```csharp
public interface IFileDataProccessingWithDatabase{
  
    //Xử lý thêm dữ liệu từ file .csv vào trong bảng cụ thể củ csdl
    public Task<bool> Insert_CSV_fileDataIntoDatabase(IFormFile file, string tableName);

    //Xử lý thêm dữ liệu từ file .tsv vào trong bảng cụ thể củ csdl
    public Task<bool> Insert_TSV_fileDataIntoDatabase(IFormFile file,string tableName);

}

```

  

#  Triển khai giao diện trên lớp cụ thể

```csharp
using Microsoft.AspNetCore.Http;
using MySqlConnector;
using Microsoft.Extensions.Configuration;
using System;
using System.IO;
using System.Threading.Tasks;

public class FileDataProccessingWithDatabase : IFileDataProccessingWithDatabase{

    private readonly IConfiguration _configuration;

    private readonly ILogger _logger;

  

    //Hàm khởi tạo

    public FileDataProccessingWithDatabase(IConfiguration configuration, ILogger logger){

        _configuration = configuration;

        _logger = logger;

    }

  

    public async Task<bool> Insert_CSV_fileDataIntoDatabase(IFormFile file, string tableName){

  

        //Kiểm tra tệp tin không được rỗng

        if(file == null || file.Length == 0)

            return false;

        try{

            //Lưu tệp tin vào vị trí tạm thời
            var tempFilePath = Path.GetTempFileName();
            using(var stream = new FileStream(tempFilePath, FileMode.Create)){
                await file.CopyToAsync(stream);
            }

            //Kết nối cơ sở dữ liệu rồi thêm dữ liệu vào

            await using (MySqlConnection connection = new MySqlConnection(_configuration["Database:MySQL"])){

                await connection.OpenAsync();

  

                MySqlBulkLoader bulkLoader = new MySqlBulkLoader(connection){

                    //Các tùy chọn cho dâu phân cách, ký tự thoát và xuống dòng

                    FiledTerminator = ",",          //Dấy phân cách cho csv
                    FieldQuotationCharacter = '"',  //Ký tự trích dẫn
                    LineTerminator = '\n',          //Ký tự xuống dòng

                    FileName = tempFilePath,
                    TableName = tableName,
                    EscapeCharacter = '\\'          //Ký tự thoát
                };



                //Thực hiện bulk load
                int rowAffected = await bulkLoader.LoadAsync();
                _logger.Info($"Successfully inserted {rowsAffected} rows into {tableName} from CSV.");

  

                // Xóa tệp tạm

                File.Delete(tempFilePath);

                return true;

            }

            //Xóa tệp tin tạm thời

        }catch (MySqlException ex){

            _logger.Error($"Error inserting CSV data into {tableName}: {ex.Message}", ex);

            return false;

        }

        catch (Exception ex)

        {

            _logger.Error($"Unexpected error processing CSV file: {ex.Message}", ex);

            return false;

        }

    }

  

    //Xử lý thêm dữ liệu từ file .tsv vào trong bảng cụ thể củ csdl

    public Task<bool> Insert_TSV_fileDataIntoDatabase(IFormFile file, string tableName){

        //Kiểm tra tệp tin không được rỗng

        if(file == null || file.Length == 0)

            return false;

        try{

            //Lưu tệp tin vào vị trí tạm thời

            var tempFilePath = Path.GetTempFileName();

            using(var stream = new FileStream(tempFilePath, FileMode.Create)){

                await file.CopyToAsync(stream);

            }

            //Kết nối cơ sở dữ liệu rồi thêm dữ liệu vào

            await using (MySqlConnection connection = new MySqlConnection(_configuration["Database:MySQL"])){

                await connection.OpenAsync();

  

                MySqlBulkLoader bulkLoader = new MySqlBulkLoader(connection){

                    //Các tùy chọn cho dâu phân cách, ký tự thoát và xuống dòng

                    FiledTerminator = "\t",          //Dấy phân cách cho csv

                    FieldQuotationCharacter = null,  //Ký tự trích dẫn

                    LineTerminator = '\n',          //Ký tự xuống dòng

                    FileName = tempFilePath,

                    TableName = tableName,

                    EscapeCharacter = '\\'          //Ký tự thoát

                };

  

                //Thực hiện bulk load

                int rowAffected = await bulkLoader.LoadAsync();

                _logger.Info($"Successfully inserted {rowsAffected} rows into {tableName} from CSV.");

  

                // Xóa tệp tạm

                File.Delete(tempFilePath);

                return true;

            }

            //Xóa tệp tin tạm thời

        }catch (MySqlException ex){

            _logger.Error($"Error inserting CSV data into {tableName}: {ex.Message}", ex);

            return false;

        }

        catch (Exception ex)

        {

            _logger.Error($"Unexpected error processing CSV file: {ex.Message}", ex);

            return false;

        }

    }

}

```

  

# Đăng ký dịch vụ

```csharp

services.AddScoped<IFileDataProccessingWithDatabase, FileDataProccessingWithDatabase>();

```

  

# thiết lập endpoit

```csharp

using Microsoft.AspNetCore.Mvc;

using Microsoft.AspNetCore.Http;

using System.Threading.Tasks;

  

[Route("file")]

[ApiController]

public class FileDataProccessingWithDatabaseController: ControllerBase{

  

    private IFileDataProccessingWithDatabase fileDataProccessingWithDatabase;

    public FileDataProccessingWithDatabaseController(IFileDataProccessingWithDatabase _fileDataProccessingWithDatabase){

        fileDataProccessingWithDatabase = _fileDataProccessingWithDatabase;

    }

  

    [HttpPost('process-csv')]

    [Consumes("multipart/form-data")]

    public async Task<IActionResult> Insert_CSV_fileDataIntoDatabase(IFormFile file, [FromQuery] string tablename){

        if (string.IsNullOrEmpty(tableName))

            return BadRequest("Table name is required.");

  

        var result = await _fileDataProccessingWithDatabase.Insert_CSV_fileDataIntoDatabase(file, tableName);

        return result ? Ok("CSV data inserted successfully.") : StatusCode(500, "Failed to insert CSV data.");

    }

  

    [HttpPost('process-tsv')]

    [Consumes("multipart/form-data")]

    public async Task<IActionResult> Insert_TSV_fileDataIntoDatabase(IFormFile file, [FromQuery] string tablename){

        if (string.IsNullOrEmpty(tableName))

            return BadRequest("Table name is required.");

  

        var result = await _fileDataProccessingWithDatabase.Insert_TSV_fileDataIntoDatabase(file, tableName);

        return result ? Ok("TSV data inserted successfully.") : StatusCode(500, "Failed to insert TSV data.");

    }

}

```

---
Hỏi (một cách nhanh chóng)
- Xử lý tác vụ thêm hàng triệu records (có thể từ excel, .csv, .tsv, .sql) vào trong bảng cụ thể với thời gian tối thiểu.
- Xử lý chuyển dữ liệu (với 1 triệu record) từ MySQL sang SQL.
- Xử lý hàng triệu record đã select để insert vào bảng khác. 


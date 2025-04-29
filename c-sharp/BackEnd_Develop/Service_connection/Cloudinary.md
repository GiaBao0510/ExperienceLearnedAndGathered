# Tích hợp Cloudinary trong ứng dụng .NET để quản lý hình ảnh

Cloudinary là một dịch vụ quản lý hình ảnh trên cloud, giúp tải lên, lưu trữ, tối ưu và phân phối hình ảnh hiệu quả. Tài liệu này hướng dẫn cách tích hợp Cloudinary vào ứng dụng .NET (như E-commerce API) để xử lý ảnh sản phẩm (tải lên, xóa, cập nhật).

---

## Yêu cầu

Cài đặt gói NuGet sau để sử dụng Cloudinary:
```bash
dotnet add package CloudinaryDotNet
```

---
## Cấu hình Cloudinary

Cấu hình thông tin tài khoản Cloudinary trong tệp appsettings.json. Bạn có thể lấy thông tin này từ bảng điều khiển Cloudinary.

```json
"CloudinarySetting":{
    "CloudName":"dkafsdfq6",
    "ApiKey":"9821123457895",
    "ApiSecret":"Kyp5451JHJSRSSJDHSY10dzeA0"
 }
```

### Giải thích cấu hình

| Tham số   | Ý nghĩa                                            | Ví dụ                      |
| --------- | -------------------------------------------------- | -------------------------- |
| CloudName | Tên cloud của tài khoản Cloudinary.                | dkafsdfq6                  |
| ApiKey    | Khóa API để xác thực.                              | 9821123457895              |
| ApiSecret | Mật khẩu API để xác thực (bảo mật, không chia sẻ). | Kyp5451JHJSRSSJDHSY10dzeA0 |

**Lưu ý**: Trong môi trường production, lưu ApiSecret trong biến môi trường hoặc Azure Key Vault để tăng bảo mật.

---
## Định nghĩa giao diện ICloudinaryServices

Tạo tệp ICloudinaryServices.cs để định nghĩa các phương thức quản lý hình ảnh:
```csharp
public interface ICloudinaryServices
{

	///<summary>
	/// Thêm ảnh
	/// </summary>
	public Task<ImageUploadResult> AddImageAssync(IFormFile file);

	/// <summary>
	/// Xóa ảnh dựa trên publicId
	/// </summary>
	public Task<DeletionResult> DeleteImageAssync(string publicId);

	/// <summary>
	/// thay đổi ảnh dựa trên public Id
	/// </summary>      
	 public Task<ImageUploadResult> UpdateImageAssync(IFormFile file, string publicId);
}
```

---
## Triển khai dịch vụ CloudinaryServices

Tạo tệp CloudinaryServices.cs để triển khai giao diện:

```csharp
public class CloudinaryServices: ICloudinaryServices
{

	#region  ===[Private Member]===
	private readonly Cloudinary _cloudinary;
	private readonly ILogger _logger;
	#endregion

	///<summary>
	/// Hàm khởi tạo
	/// </summary>
	public CloudinaryServices(
		IConfiguration configuration,
		ILogger logger
	){
		var account = new Account(
		configuration["Authentication:CloudinarySetting:CloudName"],
		configuration["Authentication:CloudinarySetting:ApiKey"],
		configuration["Authentication:CloudinarySetting:ApiSecret"]
		);
		
		_cloudinary = new Cloudinary(account);
		_logger = logger;
	}

	///<summary>
	/// Thêm ảnh
	/// </summary>
	public async Task<ImageUploadResult> AddImageAssync(IFormFile file)
	{

		//Nếu file null thì báo lỗi
		if(file.Length == 0){
			_logger.Error($"File rỗng hoặc không tồn tại. {nameof(file)}");
			throw new 
				ValidationException
				($"File rỗng hoặc không tồn tại. {nameof(file)}");
		}

		try{
			var uploadResult = new ImageUploadResult();
			using (var stream = file.OpenReadStream()){     //Mở luồng đọc file
				var uploadParams = new ImageUploadParams(){   //Tham số upload
					File = new FileDescription(file.FileName, stream),  //Đường dẫn file
					Folder = "NguoiDung",     //Thư mục lưu trữ trên cloudinary
					Transformation = new Transformation()   // Chuyển đổi ảnh
					.Quality("auto")                  // Chất lượng ảnh tự động
					.FetchFormat("auto")              // Định dạng ảnh tự động
					.Flags("preserve_transparency")   // Giữ nguyên độ trong suốt

				};

				return await _cloudinary.UploadAsync(uploadParams);
			}

		}catch(Exception ex){
			_logger.Error($"Lỗi khi upload ảnh lên cloudinary: {ex.Message}", ex);
			throw new DetailsOfTheException(ex);
		}
	}

	/// <summary>
	/// Xóa ảnh dựa trên publicId
	/// </summary>
	public async Task<DeletionResult> DeleteImageAssync(string publicId){
		try
		{
			var deleteParams = new DeletionParams(publicId);
			return await _cloudinary.DestroyAsync(deleteParams);
		}
		catch (Exception ex)
		{

			_logger.Error($"Lỗi khi xóa ảnh trên cloudinary: {ex.Message}", ex);
			throw new DetailsOfTheException(ex);
		}
	}

	/// <summary>
	/// thay đổi ảnh dựa trên public Id
	/// </summary>      
	public async Task<ImageUploadResult> UpdateImageAssync
		(IFormFile file, string publicId){
		try{

			//Xóa ảnh cũ
			await DeleteImageAssync(publicId);

			//Cạp nhật lại ảnh mới
			return await AddImageAssync(file);
		}

		catch (Exception ex)
		{
			_logger
				.Error($"Lỗi khi cập nhật ảnh trên cloudinary: {ex.Message}", ex);
			throw new DetailsOfTheException(ex);
		}
	}
}
```

---
## Đăng ký dịch vụ

Đăng ký CloudinaryServices trong Program.cs với lifetime **Singleton** để tái sử dụng instance duy nhất:

```csharp
builder.Services.AddSingleton<ICloudinaryServices, CloudinaryServices>();
```

---
## Ví dụ sử dụng trong Controller

Dưới đây là cách sử dụng CloudinaryServices trong một API controller để tải ảnh sản phẩm:
```csharp
using Microsoft.AspNetCore.Mvc;
using E_commerce.Infrastructure.Services;

[Route("api/[controller]")]
[ApiController]
public class ImageController : ControllerBase
{
    private readonly ICloudinaryServices _cloudinaryServices;

    public ImageController(ICloudinaryServices cloudinaryServices)
    {
        _cloudinaryServices = cloudinaryServices;
    }

    [HttpPost("upload")]
    public async Task<IActionResult> UploadImage(IFormFile file)
    {
        var result = await _cloudinaryServices.AddImageAsync(file);
        return Ok(new
        {
            PublicId = result.PublicId,
            Url = result.SecureUrl.ToString()
        });
    }

    [HttpDelete("delete/{publicId}")]
    public async Task<IActionResult> DeleteImage(string publicId)
    {
        var result = await _cloudinaryServices.DeleteImageAsync(publicId);
        return Ok(new { Message = $"Image deleted: {result.Result}" });
    }

    [HttpPut("update/{publicId}")]
    public async Task<IActionResult> UpdateImage(IFormFile file, string publicId)
    {
        var result = await _cloudinaryServices.UpdateImageAsync(file, publicId);
        return Ok(new
        {
            PublicId = result.PublicId,
            Url = result.SecureUrl.ToString()
        });
    }
}
```

---
## Kiểm tra tích hợp

1. **Tạo tài khoản Cloudinary**:
    
    - Đăng ký tại Cloudinary.
    - Lấy CloudName, ApiKey, ApiSecret từ bảng điều khiển.
    
2. **Gọi API**:
    
    - Sử dụng Postman hoặc curl để gọi endpoint /api/image/upload với file ảnh.
    - Kiểm tra response để lấy PublicId và Url.        
    
3. **Kiểm tra log**:
    
    - Đảm bảo logger ghi lại thông tin upload/xóa/cập nhật thành công hoặc lỗi.

---
## Lưu ý

- **Bảo mật**: Không lưu ApiSecret trực tiếp trong appsettings.json ở môi trường production. Sử dụng biến môi trường hoặc dịch vụ như Azure Key Vault.
- **Tối ưu**: Điều chỉnh Transformation (chất lượng, định dạng) dựa trên nhu cầu (ví dụ: giảm kích thước ảnh để tiết kiệm băng thông).
- **Kiểm thử**: Viết unit test cho CloudinaryServices bằng cách mock Cloudinary hoặc dùng tài khoản Cloudinary test.
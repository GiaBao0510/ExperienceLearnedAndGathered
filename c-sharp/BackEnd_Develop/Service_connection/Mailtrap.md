# Gửi Email với Mailtrap trong ứng dụng .NET

**Mailtrap** là một dịch vụ giả lập email dùng để kiểm tra và phát triển, cho phép gửi email mà không thực sự gửi đến người nhận (rất hữu ích trong môi trường test). Tài liệu này hướng dẫn cách tích hợp Mailtrap vào ứng dụng .NET (như E-commerce API) để gửi mã OTP qua email, kết hợp Redis để quản lý OTP.

---
## Yêu cầu

Cài đặt gói NuGet sau để gửi email qua Mailtrap:
```bash
dotnet add package RestSharp
```

---
## Cấu hình Mailtrap

Cấu hình thông tin tài khoản Mailtrap trong tệp appsettings.json. Bạn có thể lấy thông tin này từ bảng điều khiển Mailtrap.

```json
{
  "Mailtrap": {
    "UserName": "87a54ba84e6b60",
    "Password": "1xxxxxxxxxxx",
    "EmailTest": "hello@example.com",
    "Port": 2525,
    "Host": "https://sandbox.api.mailtrap.io/api/send/3651870",
    "ApiToken": "2xxxxxxxxxxxxxxxxxxxxxx"
  }
}
```
### Giải thích cấu hình

| Tham số   | Ý nghĩa                      | Ví dụ                                            |
| --------- | ---------------------------- | ------------------------------------------------ |
| UserName  | Tên người dùng Mailtrap.     | 87a54ba84e6b60                                   |
| Password  | Mật khẩu Mailtrap.           | 1xxxxxxxxxxx                                     |
| EmailTest | Email dùng để gửi (giả lập). | hello@example.com                                |
| Port      | Cổng SMTP của Mailtrap.      | 2525                                             |
| Host      | URL API của Mailtrap.        | https://sandbox.api.mailtrap.io/api/send/3651870 |
| ApiToken  | Token API để xác thực.       | 2xxxxxxxxxxxxxxxxxxxxxx                          |

**Lưu ý**: Trong môi trường production, lưu ApiToken và Password trong biến môi trường hoặc Azure Key Vault để tăng bảo mật.

---
## Định nghĩa mô hình dữ liệu

Tạo lớp _EmailModel để biểu diễn email:
```csharp
namespace E_commerce.Core.Entities
{
    public class _EmailModel
    {
        public string EmailFrom { get; set; }
        public string EmailTo { get; set; }
        public string Subject { get; set; }
        public string htmlMessage { get; set; }
    }
}
```

Tạo DTO VerifyOTP_DTO để xác thực OTP:
```csharp
namespace E_commerce.Application.DTOs.Requests
{
    public class VerifyOTP_DTO
    {
        public string Email { get; set; }
        public string OTP { get; set; }
    }
}
```

---
## Định nghĩa giao diện IMailtrapService

Tạo tệp IMailtrapService.cs để định nghĩa phương thức gửi email:
```csharp
namespace E_commerce.Infrastructure.Services
{
    public interface IMailtrapService
    {
        /// <summary>
        /// Gửi email qua Mailtrap.
        /// </summary>
        /// <param name="emailModel">Thông tin email (từ, đến, chủ đề, nội dung).</param>
        /// <returns>True nếu gửi thành công, False nếu thất bại.</returns>
        Task<bool> SendEmailAsync(_EmailModel emailModel);
    }
}
```

---
## Triển khai dịch vụ MailtrapService

Tạo tệp MailtrapService.cs để triển khai gửi email qua Mailtrap:

```csharp
using E_commerce.Core.Entities;
using Microsoft.Extensions.Configuration;
using Microsoft.Extensions.Logging;
using RestSharp;

namespace E_commerce.Infrastructure.Services
{
    public class MailtrapService : IMailtrapService
    {
        private readonly IConfiguration _configuration;
        private readonly ILogger<MailtrapService> _logger;

        public MailtrapService(IConfiguration configuration, ILogger<MailtrapService> logger)
        {
            _configuration = configuration ?? throw new ArgumentNullException(nameof(configuration));
            _logger = logger ?? throw new ArgumentNullException(nameof(logger));
        }

        /// <summary>
        /// Gửi email qua Mailtrap API.
        /// </summary>
        public async Task<bool> SendEmailAsync(_EmailModel emailModel)
        {
            if (emailModel == null || string.IsNullOrEmpty(emailModel.EmailTo))
            {
                _logger.LogError
	                ("Email model or recipient is null or empty.");
                return false;
            }

            var client = new RestClient(_configuration["Mailtrap:Host"]
                ?? throw new InvalidOperationException
	                ("Mailtrap Host is not configured."));
            var request = new RestRequest();

            request.AddHeader("Authorization", 
	            $"Bearer {_configuration["Mailtrap:ApiToken"]}");
            request.AddHeader("Content-Type", "application/json");
            request.AddJsonBody(new
            {
                from = new { email = emailModel.EmailFrom, name = "E-commerce" },
                to = new[] { new { email = emailModel.EmailTo } },
                subject = emailModel.Subject,
                html = emailModel.htmlMessage
            });

            try
            {
                var response = await client.ExecutePostAsync(request);
                if (response.StatusCode == System.Net.HttpStatusCode.OK)
                {
                    _logger.LogInformation
	                    ("Email sent successfully: {Content}", response.Content);
                    return true;
                }

                _logger.LogError
	                ("Failed to send email: {Content}", response.Content);
                return false;
            }
            catch (Exception ex)
            {
                _logger.LogError(ex, "Error sending email: {Message}", ex.Message);
                throw;
            }
        }
    }
}
```

---
## Định nghĩa giao diện IOTPAuthenServices

Tạo tệp IOTPAuthenServices.cs để định nghĩa các phương thức liên quan đến OTP:

```csharp
using E_commerce.Application.DTOs.Requests;

namespace E_commerce.Infrastructure.Services
{
    public interface IOTPAuthenServices
    {
        /// <summary>
        /// Gửi mã OTP đến email và lưu vào Redis.
        /// </summary>
        Task<bool> AddOTP_EmailToRedis_Mailtrap(string email);

        /// <summary>
        /// Xác thực mã OTP từ email.
        /// </summary>
        Task<bool> VerifyOTP_Email(VerifyOTP_DTO info);
    }
}
```

---
## Triển khai dịch vụ OTPAuthenServices
Tạo tệp OTPAuthenServices.cs để triển khai logic OTP:

```csharp
using System;
using System.Collections.Generic;
using System.Linq;
using System.Threading.Tasks;
using E_commerce.Application.Application;
using E_commerce.Application.DTOs.Requests;
using E_commerce.Core.Entities;
using E_commerce.Core.Exceptions;
using E_commerce.Infrastructure.Templates.Email;
using E_commerce.Infrastructure.Utils;
using Microsoft.Extensions.Configuration;

namespace E_commerce.Infrastructure.Services
{
    public class OTPAuthenServices : IOTPAuthenServices
    {
        private readonly IRedisServices _redisServices;
        private readonly IMailtrapService _mailtrapService;
        private readonly IUserRepository _userRepository;
        private readonly IConfiguration _configuration;
        private readonly ILogger<OTPAuthenServices> _logger;

        public OTPAuthenServices(
            IRedisServices redisServices,
            IMailtrapService mailtrapService,
            IUserRepository userRepository,
            IConfiguration configuration,
            ILogger<OTPAuthenServices> logger)
        {
            _redisServices = redisServices ?? 
	            throw new ArgumentNullException(nameof(redisServices));
            _mailtrapService = mailtrapService ?? 
	            throw new ArgumentNullException(nameof(mailtrapService));
            _userRepository = userRepository ?? 
	            throw new ArgumentNullException(nameof(userRepository));
            _configuration = configuration ?? 
	            throw new ArgumentNullException(nameof(configuration));
            _logger = logger ?? throw new ArgumentNullException(nameof(logger));
        }

        private void CheckValid(string content, string fieldName)
        {
            if (string.IsNullOrEmpty(content))
            {
                throw new InvalidOperationException
	                ($"{fieldName} cannot be empty.");
            }
        }

        private async Task<bool> HasExceededRateLimit(string email)
        {
            try
            {
                string rateLimitKey = $"otp:ratelimit:{email}";
                int count = await _redisServices.Get<int>(rateLimitKey);
                int maxCount = 5; // Giới hạn 5 lần gửi OTP trong 1 giờ

                if (count >= maxCount)
                {
                    _logger.LogWarning
	                    ("Rate limit exceeded for email: {Email}", email);
                    return true;
                }

                if (count == 0)
                {
                    await _redisServices.Set(
	                    rateLimitKey, 
	                    1, 
	                    TimeSpan.FromHours(1)
                    );
                }
                else
                {
                    await _redisServices
	                    .Set(rateLimitKey, count + 1, TimeSpan.FromHours(1));
                }

                return false;
            }
            catch (Exception ex)
            {
                _logger.LogError
                (
	                ex, 
	                "Error checking OTP rate limit for email: {Email}", 
	                email
                );
                throw;
            }
        }

        private async Task BeforeSendingOTP(string email)
        {
            CheckValid(email, "Email");

            bool userExists = await _userRepository.IsUserEmailExists(email);
            if (!userExists)
            {
                throw new 
                InvalidOperationException("Email does not exist in the system.");
            }

            bool isKeyExists = await _redisServices.KeyExists(email);
            if (isKeyExists)
            {
                throw new InvalidOperationException
                ("An OTP has already been sent to this email. Please check your inbox.");
            }

            if (await HasExceededRateLimit(email))
            {
                throw new InvalidOperationException
	                ("Too many OTP requests. Please try again after 1 hour.");
            }
        }

        public async Task<bool> AddOTP_EmailToRedis_Mailtrap(string email)
        {
            try
            {
                await BeforeSendingOTP(email);

                string otp = CodeGenerator.GenerateRandomNumber(6);
                await _redisServices.Set(email, otp, TimeSpan.FromMinutes(5));

                string template = VerificationForm.OTPcodeVerificationForm(otp);
                var emailModel = new _EmailModel
                {
                    Subject = "OTP Verification Code",
                    EmailFrom = _configuration["Mailtrap:EmailTest"]
                        ?? 
                        throw new InvalidOperationException("Mailtrap EmailTest is not configured."),
                    EmailTo = email,
                    htmlMessage = template
                };

                bool sent = await _mailtrapService.SendEmailAsync(emailModel);
                if (sent)
                {
                    _logger.LogInformation
	                    ("OTP sent successfully to email: {Email}", email);
                }
                return sent;
            }
            catch (Exception ex)
            {
                _logger.LogError(ex, "Error sending OTP to email: {Email}", email);
                throw;
            }
        }

        public async Task<bool> VerifyOTP_Email(VerifyOTP_DTO info)
        {
            try
            {
                CheckValid(info.Email, "Email");
                CheckValid(info.OTP, "OTP");

                bool userExists = await 
	                _userRepository.IsUserEmailExists(info.Email);
                if (!userExists)
                {
                    throw new InvalidOperationException
	                    ("Email does not exist in the system.");
                }

                bool isKeyExists = await _redisServices.KeyExists(info.Email);
                if (!isKeyExists)
                {
                    throw new InvalidOperationException
	                    ("OTP has expired. Please request a new OTP.");
                }

                string redisOTP = await _redisServices.Get<string>(info.Email);
                if (redisOTP != info.OTP)
                {
                    throw new InvalidOperationException
	                    ("Invalid OTP. Please check your OTP code.");
                }

                bool removed = await _redisServices.Remove(info.Email);
                _logger.LogInformation
	                ("OTP verified and removed for email: {Email}", info.Email);
                return removed;
            }
            catch (Exception ex)
            {
                _logger.LogError
	                (ex, "Error verifying OTP for email: {Email}", info.Email);
                throw;
            }
        }
    }
}
```

---
## Đăng ký dịch vụ

Đăng ký các dịch vụ trong Program.cs với lifetime **Scoped** để đảm bảo mỗi request HTTP tạo một instance mới:
```csharp
builder.Services.AddScoped<IMailtrapService, MailtrapService>();
builder.Services.AddScoped<IOTPAuthenServices, OTPAuthenServices>();
```

---
## Ví dụ sử dụng trong Controller

Dưới đây là cách sử dụng OTPAuthenServices trong một API controller để gửi và xác thực OTP:
```csharp
using E_commerce.Application.DTOs.Requests;
using E_commerce.Infrastructure.Services;
using Microsoft.AspNetCore.Mvc;

[Route("api/[controller]")]
[ApiController]
public class OTPController : ControllerBase
{
    private readonly IOTPAuthenServices _otpService;

    public OTPController(IOTPAuthenServices otpService)
    {
        _otpService = otpService;
    }

    [HttpPost("send")]
    public async Task<IActionResult> SendOTP(string email)
    {
        bool sent = await _otpService.AddOTP_EmailToRedis_Mailtrap(email);
        return sent ? Ok("OTP sent successfully.") : 
	        BadRequest("Failed to send OTP.");
    }

    [HttpPost("verify")]
    public async Task<IActionResult> VerifyOTP([FromBody] VerifyOTP_DTO request)
    {
        bool verified = await _otpService.VerifyOTP_Email(request);
        return verified ? Ok("OTP verified successfully.") : 
	        BadRequest("Invalid or expired OTP.");
    }
}
```

---
## Kiểm tra tích hợp

1. **Tạo tài khoản Mailtrap**:
    - Đăng ký tại Mailtrap.
    - Lấy UserName, Password, ApiToken, và Host từ mục **SMTP/API Settings** trong bảng điều khiển.
    
2. **Gọi API**:
    - Dùng Postman để gọi endpoint /api/otp/send với email (ví dụ: user@example.com).
    - Kiểm tra hộp thư giả lập trên Mailtrap để xem mã OTP.
    - Gọi /api/otp/verify với email và OTP để xác thực.
    
3. **Kiểm tra Redis**:
    - Đảm bảo OTP được lưu trong Redis với TTL 5 phút.
    - Kiểm tra rate limit (5 lần/giờ) bằng cách gửi nhiều yêu cầu.
    
4. **Kiểm tra log**:
    - Xem log để xác nhận email gửi thành công hoặc lỗi (dùng ILogger).

---
## Xử lý lỗi thường gặp

| Lỗi                  | Nguyên nhân                         | Cách khắc phục                            |
| -------------------- | ----------------------------------- | ----------------------------------------- |
| Unauthorized         | Sai ApiToken.                       | Kiểm tra ApiToken trong appsettings.json. |
| Email not found      | Email không tồn tại trong database. | Đảm bảo người dùng đã đăng ký.            |
| Rate limit exceeded  | Gửi quá 5 OTP trong 1 giờ.          | Chờ 1 giờ hoặc tăng maxCount trong code.  |
| Failed to send email | Sai Host hoặc kết nối mạng kém.     | Kiểm tra Host và mạng.                    |

---
## Lưu ý

- **Môi trường test**: Mailtrap chỉ dùng trong môi trường phát triển/test. Trong production, dùng dịch vụ như SendGrid hoặc Amazon SES.
- **Bảo mật**: Lưu ApiToken và Password trong biến môi trường để tránh lộ thông tin.
- **Tối ưu**: Điều chỉnh TTL của OTP (5 phút) và rate limit (5 lần/giờ) dựa trên nhu cầu.
- **Kiểm thử**: Viết unit test cho MailtrapService và OTPAuthenServices bằng cách mock RestClient và IRedisServices.
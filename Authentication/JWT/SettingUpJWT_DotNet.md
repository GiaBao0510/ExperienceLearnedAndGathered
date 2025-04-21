## Thiết lập JSON Web Token (JWT) trong ứng dụng .NET

Tài liệu này hướng dẫn cách thiết lập JSON Web Token (JWT) trong ứng dụng .NET Core để xác thực và phân quyền API. Chúng ta sẽ cấu hình JWT, tạo dịch vụ để sinh token, và hỗ trợ cả **Access Token** và **Refresh Token**.

#### Yêu cầu và cài đặt

Trước khi bắt đầu, cần cài đặt các gói NuGet sau:
```bash
dotnet add package Microsoft.AspNetCore.Authentication.JwtBearer dotnet add package System.IdentityModel.Tokens.Jwt
```

---
#### **Bước 1: Cấu hình trong `appsettings.json`**

Thêm cấu hình JWT vào tệp `appsettings.jso`n. Các thông tin bao gồm `khóa bí mật`, `issuer`, `audience`, và **thời gian hết hạn** cho Access Token/Refresh Token.

```json
{
	"Authentication":{
    "JWT":{
      "secret": "EJ_{&J##l6tEP~ZzGK)f52+wQsdgd2f9RS5ED4AE,6~HxwuLLvue",
      "issuer": "E_commerce_api",
      "audience": "E_commerce_client",
      "AccessTokenExpirationMinutes": 90,
      "RefreshTokenExpirationDays": 7
    }
  }
}
```

*Giải thích các tham số:*
- **secret:** Khóa bí mật dùng để xác minh token
- **issuer:** Tên của server phát hành token
- **audience:** Đối tượng sử dụng token (thường là client)
- **AccessTokenExpirationMinutes:** Thời gian sống accesstoken (phút)
- **RefreshTokenExpirationDays:** Thời gian sống refreshtoken (phút)

> ***⚠️Lưu ý: Lưu khóa bí mật (`secret`) phải dài, ngẫu nhiên và được bảo vệ cẩn thận, không để bị lộ trong mã nguồn.***

---
### **Bước 2: Thiết lập xác thực trong `program.cs`**

Cấu hình xác thực JWT trong `program.cs` để server có thể xác minh token trong các request.
```csharp
using Microsoft.AspNetCore.Authentication.JwtBearer;
using Microsoft.IdentityModel.Tokens;
using System.Text;

var builder = WebApplication.CreateBuilder(args);

// Thêm dịch vụ xác thực JWT
builder.Services.AddAuthentication(options =>
{
    // Thiết lập các cơ chế xác thực mặc định
    options.DefaultAuthenticateScheme = JwtBearerDefaults.AuthenticationScheme; // Xác thực bằng JWT
    options.DefaultChallengeScheme = JwtBearerDefaults.AuthenticationScheme; // Chuyển hướng khi chưa đăng nhập
    options.DefaultForbidScheme = JwtBearerDefaults.AuthenticationScheme; // Từ chối khi không có quyền
    options.DefaultScheme = JwtBearerDefaults.AuthenticationScheme; // Cơ chế mặc định
    options.DefaultSignInScheme = JwtBearerDefaults.AuthenticationScheme; // Đăng nhập bằng JWT
    options.DefaultSignOutScheme = JwtBearerDefaults.AuthenticationScheme; // Đăng xuất
})
.AddJwtBearer(options =>
{
    // Cấu hình các tham số xác minh token
    options.TokenValidationParameters = new TokenValidationParameters
    {
        ValidateIssuer = true, // Kiểm tra issuer của token
        ValidateAudience = true, // Kiểm tra audience của token
        ValidateIssuerSigningKey = true, // Kiểm tra chữ ký của token
        ValidateLifetime = true, // Kiểm tra thời gian hết hạn
        ValidIssuer = builder.Configuration["Authentication:JWT:issuer"], // Issuer hợp lệ
        ValidAudience = builder.Configuration["Authentication:JWT:audience"], // Audience hợp lệ
        IssuerSigningKey = new SymmetricSecurityKey(
            Encoding.UTF8.GetBytes(
                builder.Configuration["Authentication:JWT:secret"] ??
                throw new InvalidOperationException("JWT secret key is missing")
            )
        ),
        ClockSkew = TimeSpan.Zero // Không cho phép chênh lệch thời gian
    };

    // Xử lý các sự kiện liên quan đến token
    options.Events = new JwtBearerEvents
    {
        // Khi xác thực thất bại (ví dụ: token hết hạn)
        OnAuthenticationFailed = context =>
        {
            if (context.Exception is SecurityTokenExpiredException)
            {
                context.Response.Headers.Add("Token-Expired", "true"); // Báo token hết hạn qua header
            }
            return Task.CompletedTask;
        },
        // Hỗ trợ lấy token từ query string (cho WebSocket hoặc Server-Sent Events)
        OnMessageReceived = context =>
        {
            var accessToken = context.Request.Query["access_token"];
            if (!string.IsNullOrEmpty(accessToken))
            {
                context.Token = accessToken; // Gán token để xác thực
            }
            return Task.CompletedTask;
        }
    };
});

// Đăng ký các dịch vụ khác (sẽ giải thích ở bước sau)
builder.Services.AddScoped<ITokenService, TokenService>();

var app = builder.Build();

// Kích hoạt xác thực và phân quyền
app.UseAuthentication();
app.UseAuthorization();

app.Run();
```

>***⚠️ Lưu ý: đừng quên gọi `app.UseAuthentication() `và `app.UseAuthorization()` trong pipeline để kích hoạt xác thực.***

---
### **Bước 3: Tạo giao diện `ITokenService`**

Tạo giao diện `ITokenService.cs` để định nghĩa các phương thức liên quan đến token, bao gồm tạo Access Token, Refresh Token, và kiểm tra tính hợp lệ.

```csharp
public interface ITokenService
{
    // Kiểm tra xem token còn hạn hay không
    Task<bool> CheckIfTokenIsExpired(string token);

    // Kiểm tra xem token có hợp lệ hay không
    Task<bool> CheckIfTokenIsValid(string token);

    // Tạo Access Token và Refresh Token
    Task<TokenDTO> GenerateToken(AccountInforDTO accountInforDTO, int hours = 24);

    // Tạo Refresh Token mới dựa trên Refresh Token cũ
    Task<TokenDTO> RefreshToken(string refreshToken, string userId);
}
```

> ***⚠️ Lưu ý: `TokenDTO` là một DTO để trả về thông tin token (chuỗi token, thời gian hết hạn, User ID)***

---
### **Bước 4: Triển khai `TokenService`**

Tạo lớp `TokenService.cs `để triển khai giao diện `ITokenService`. Lớp này chịu trách nhiệm tạo và xác minh token.

```csharp
using Microsoft.IdentityModel.Tokens;
using System.IdentityModel.Tokens.Jwt;
using System.Security.Claims;
using System.Text;

public class TokenService : ITokenService
{
    #region Private Fields
    private readonly IConfiguration _configuration; // Cấu hình từ appsettings.json
    private readonly ICustomFormat 
	    _customFormat; // Dịch vụ định dạng dữ liệu (giả định)
    private readonly IUserRepository 
	    _userRepository; // Repository để lấy thông tin người dùng
    private readonly ILogger _logger; // Logger để ghi log lỗi
    private readonly SymmetricSecurityKey _secretKey; // Khóa bí mật để ký token
    private readonly string _issuer; // Issuer của token
    private readonly string _audience; // Audience của token
    #endregion

    /// <summary>
    /// Constructor để khởi tạo TokenService
    /// </summary>
    public TokenService(
        IConfiguration configuration,
        ICustomFormat customFormat,
        IUserRepository userRepository,
        ILogger<TokenService> logger)
    {
        _configuration = configuration ?? 
	        throw new ArgumentNullException(nameof(configuration));
        _customFormat = customFormat ?? 
	        throw new ArgumentNullException(nameof(customFormat));
        _userRepository = userRepository ?? 
	        throw new ArgumentNullException(nameof(userRepository));
        _logger = logger ?? 
	        throw new ArgumentNullException(nameof(logger));

        // Kiểm tra và lấy cấu hình JWT
        var secret = configuration["Authentication:JWT:secret"];
        _issuer = configuration["Authentication:JWT:issuer"];
        _audience = configuration["Authentication:JWT:audience"];
        if (
	        string.IsNullOrEmpty(secret) || 
	        string.IsNullOrEmpty(_issuer) || 
	        string.IsNullOrEmpty(_audience)
        )
            throw new InvalidOperationException("JWT configuration is missing");

        // Khởi tạo khóa bí mật
        _secretKey = new SymmetricSecurityKey(Encoding.UTF8.GetBytes(secret));
    }

    /// <summary>
    /// Kiểm tra xem token còn hạn hay không
    /// </summary>
    public async Task<bool> CheckIfTokenIsExpired(string token)
    {
        try
        {
            var tokenHandler = new JwtSecurityTokenHandler();
            var jwtToken = tokenHandler.ReadJwtToken(token);
            var expirationDate = jwtToken.ValidTo;
            var currentDate = DateTime.UtcNow;

            _logger.LogInformation(
	            "Kiểm tra hạn token: {TokenId}, Hết hạn: {Expiration}", 
	            jwtToken.Id, 
	            expirationDate
            );
            return await Task.FromResult(expirationDate > currentDate);
        }
        catch (Exception ex)
        {
            _logger.LogError(
	            ex, 
	            "Lỗi khi kiểm tra hạn token: {Message}", 
	            ex.Message
            );
            return await Task.FromResult(false);
        }
    }

    /// <summary>
    /// Kiểm tra tính hợp lệ của token
    /// </summary>
    public async Task<bool> CheckIfTokenIsValid(string token)
    {
        try
        {
            var tokenHandler = new JwtSecurityTokenHandler();
            var principal = tokenHandler.ValidateToken(
	            token, new TokenValidationParameters
            {
                ValidateIssuer = true,
                ValidateAudience = true,
                ValidateIssuerSigningKey = true,
                ValidateLifetime = true,
                ValidIssuer = _issuer,
                ValidAudience = _audience,
                IssuerSigningKey = _secretKey
            }, out _);

            _logger.LogInformation(
	            "Token hợp lệ: {TokenId}", 
	            principal.Claims
	            .FirstOrDefault(
		            c => c.Type == JwtRegisteredClaimNames.Jti
		        )?.Value
            );
            return await Task.FromResult(true);
        }
        catch (Exception ex)
        {
            _logger.LogError(ex, "Lỗi khi kiểm tra token: {Message}", ex.Message);
            return await Task.FromResult(false);
        }
    }

    /// <summary>
    /// Tạo Access Token và Refresh Token
    /// </summary>
    public async Task<TokenDTO> GenerateToken
	    (AccountInforDTO accountInforDTO, int hours = 24)
    {
        try
        {
            // Tạo chữ ký cho token
            var creds = new SigningCredentials
	            (_secretKey, SecurityAlgorithms.HmacSha256);

            // Tạo danh sách claims (thông tin người dùng)
            var claims = new List<Claim>
            {
                new Claim(
	                JwtRegisteredClaimNames.Jti, 
	                Guid.NewGuid().ToString()
				), // ID duy nhất cho token
                new Claim(
	                ClaimTypes.NameIdentifier, 
	                accountInforDTO.user_id
                ), // ID người dùng
                new Claim(
	                ClaimTypes.Name, 
	                accountInforDTO.user_name
                ), // Tên người dùng
                new Claim(
	                ClaimTypes.Email, 
	                accountInforDTO.email
                ), // Email
                new Claim(
	                ClaimTypes.MobilePhone, 
	                accountInforDTO.phone_num
                ) // Số điện thoại
            };

            // Lấy danh sách vai trò của người dùng từ repository
            var roles = await _userRepository
	            .ListOfRoleNames(accountInforDTO.user_id);
            foreach (var role in roles)
            {
                claims.Add(
	                new Claim(
	                ClaimTypes.Role, role.role_name
                )); // Thêm vai trò vào claims
            }

            // Tạo Access Token
            var accessToken = new JwtSecurityToken(
                issuer: _issuer,
                audience: _audience,
                claims: claims,
                expires: DateTime.UtcNow.AddHours(hours), // Thời gian hết hạn
                signingCredentials: creds
            );

            // Tạo Refresh Token
            var refreshToken = Guid.NewGuid()
	            .ToString(); // Refresh Token là một chuỗi ngẫu nhiên
            
            // Lưu Refresh Token vào database hoặc Redis (giả định)
            await _userRepository
            .SaveRefreshToken(
	            accountInforDTO.user_id, 
	            refreshToken, 
	            DateTime.UtcNow.AddDays(7)
            );

            var handler = new JwtSecurityTokenHandler();
            _logger.LogInformation(
	            "Tạo token thành công cho user: {UserId}", 
	            accountInforDTO.user_id
            );

            return new TokenDTO
            {
                AccessToken = handler
	                .WriteToken(accessToken), // Chuỗi Access Token
                RefreshToken = refreshToken, // Chuỗi Refresh Token
                Expiration = accessToken.ValidTo, // Thời gian hết hạn của Access Token
                UID = accountInforDTO.user_id // ID người dùng
            };
        }
        catch (Exception ex)
        {
            _logger.LogError(
	            ex, 
	            "Lỗi khi tạo token cho user: {UserId}",
	             accountInforDTO.user_id
             );
            throw new InvalidOperationException("Không thể tạo token", ex);
        }
    }

    /// <summary>
    /// Tạo Access Token mới dựa trên Refresh Token
    /// </summary>
    public async Task<TokenDTO> RefreshToken(string refreshToken, string userId)
    {
        try
        {
            // Kiểm tra Refresh Token trong database hoặc Redis
            var storedToken = await _userRepository.GetRefreshToken(
	            userId, refreshToken
            );
            if (storedToken == null || storedToken.Expiration < DateTime.UtcNow)
            {
                _logger.LogWarning(
	                "Refresh Token không hợp lệ hoặc đã hết hạn: {UserId}",
	                 userId);
                throw new SecurityTokenException("Refresh Token không hợp lệ");
            }

            // Tạo Access Token mới
            var accountInforDTO = await _userRepository.GetUserInfo(userId);
            var newTokens = await GenerateToken(accountInforDTO, 24);

            // Xóa Refresh Token cũ và lưu Refresh Token mới
            await _userRepository.DeleteRefreshToken(userId, refreshToken);
            await _userRepository.SaveRefreshToken(
	            userId, newTokens.RefreshToken, DateTime.UtcNow.AddDays(7));

            _logger.LogInformation(
	            "Làm mới token thành công cho user: {UserId}", 
	            userId);
            return newTokens;
        }
        catch (Exception ex)
        {
            _logger.LogError(ex, "Lỗi khi làm mới token cho user: {UserId}", userId);
            throw new InvalidOperationException("Không thể làm mới token", ex);
        }
    }
}
```

> **Lưu ý:**
> - `AccountInforDTO` và `TokenDTO` là các `DTO` cần được định nghĩa (ví dụ: chứa `user_id`, `user_name`, `email`, v.v.).
> - `IUserRepository` cần có các phương thức như `ListOfRoleNames`, `SaveRefreshToken`, `GetRefreshToken`, `DeleteRefreshToken`.

---
## Bước 5: Đăng ký dịch vụ

Đăng ký TokenService và các phụ thuộc khác trong DI container.

```csharp
// Đăng ký TokenService
builder.Services.AddScoped<ITokenService, TokenService>();
```

---
## Bước 6: Sử dụng JWT trong Controller

Sử dụng attribute [Authorize] để bảo vệ các endpoint API và lấy thông tin người dùng từ Claims.
```csharp
[Route("api/[controller]")]
[ApiController]
public class UserController : ControllerBase
{
    private readonly ITokenService _tokenService;

    public UserController(ITokenService tokenService)
    {
        _tokenService = tokenService;
    }

    [HttpPost("login")]
    public async Task<IActionResult> Login([FromBody] LoginDTO loginDTO)
    {
        // Giả định kiểm tra thông tin đăng nhập
        var accountInfor = new AccountInforDTO
        {
            user_id = "123",
            user_name = loginDTO.Username,
            email = "user@example.com",
            phone_num = "1234567890"
        };
        var token = await _tokenService.GenerateToken(accountInfor);
        return Ok(token);
    }

    [HttpGet("profile")]
    [Authorize] // Yêu cầu JWT hợp lệ
    public IActionResult GetProfile()
    {
        // Lấy thông tin từ Claims
        var userId = User.FindFirst(ClaimTypes.NameIdentifier)?.Value;
        var userName = User.FindFirst(ClaimTypes.Name)?.Value;
        return Ok(new { UserId = userId, UserName = userName });
    }
}
```

---
### Bước 7: Gọi API với JWT

Client cần gửi JWT trong header Authorization của request.

**Ví dụ request:**
```http
GET /api/user/profile
Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...
```

Ví dụ response:
```json
{
  "UserId": "123",
  "UserName": "example_user"
}
```

>***Lưu ý: Lưu Access Token và Refresh Token an toàn ở client (ví dụ: HttpOnly cookies cho Refresh Token, localStorage cho Access Token).***

---
## Lưu trữ Refresh Token
Refresh Token nên được lưu trong database hoặc Redis để kiểm tra tính hợp lệ. Ví dụ cấu trúc bảng:

|Column|Type|Description|
|---|---|---|
|UserId|string|ID của người dùng|
|RefreshToken|string|Chuỗi Refresh Token|
|Expiration|datetime|Thời gian hết hạn của token|

Ví dụ lưu trong Redis:
```csharp
await _redisDB.StringSetAsync($"refresh:{userId}:{refreshToken}", refreshToken, TimeSpan.FromDays(7));
```

---

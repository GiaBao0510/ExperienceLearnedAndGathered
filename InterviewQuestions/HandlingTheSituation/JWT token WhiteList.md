- Dùng để lưu các token vào whitelist mỗi khi đăng nhập thành công, đồng thời set TTL(Time To Live) trùng với thời điểm token có hiệu lực và lưu trên whitelisst trong redis. Khi người dùng muốn thao tác tính năng nào có dựa trên API có xác thực. Thì hệ thống sẽ kiểm tra xem token này có trong whitelist không. Nếu không thì cho đăng xuất, ngược lại thì tiếp tục hành động. Khi token bị hủy và hết hạn thì sẽ bị xóa khỏi redis
- Thường thì sẽ lưu token vào **Sorted Set** trên redis, Key thì sẽ lưu thời gian hết hạn, Còn value sẽ lưu thông tin token. Việc lưu thời gian hết hạn trên key vào **Sorted Set** giúp cho quá trình tìm kiếm những token nào đã hết hạn và gần hết hạn 1 cách dễ dàng hơn.

---
### **Vì sao lưu token vào SortedSet?**
- **Sorted Set** trong Redis cho phép lưu trữ theo dạng key-value và sắp xếp theo một giá trị số (**score**), rất thích hợp cho việc lưu token với thời gian hết hạn.
- Quá trình kiểm tra, truy vấn các token gần hết hạn hay đã hết hạn diễn ra nhanh chóng nhờ tính năng sắp xếp.
- Dễ dàng thực hiện thao tác **xóa** các token hết hạn hoặc **làm sạch dữ liệu** định kỳ.

---
#### **Ví dụ:**

1. _Tạo token và lưu vào **Sorted Set** có tên là whileList trong redis_
```
using System; 
using System.Collections.Generic; 
using System.IdentityModel.Tokens.Jwt; 
using System.Security.Claims; 
using Microsoft.IdentityModel.Tokens; 
using StackExchange.Redis;

public class TokenService { 
	private readonly IDatabase _redisDB; 
	private readonly IConfiguration _configuration;

	//Khởi tạo
	public TokenService(IConnectionMultiplexer redisConnection, IConfiguration configuration){
		_redisDB = redisConnection.GetDatabase();
		_configuration = configuration;
	}

	//Tạo token và lưu
	public async Task<string> GenerateTokenAndAddToWhitelist(LoginModel loginModel)
	{
		// 1. Chỉ định các Claims cho người dùng var claims = new List<Claim>() { 
		new Claim("SDT", loginModel.account), 
		new Claim(ClaimTypes.Role, loginModel.Role), 
		new Claim(ClaimTypes.Email, loginModel.Email), 
		new Claim(JwtRegisteredClaimNames.Jti, Guid.NewGuid().ToString()), 
		new Claim("SuDung", loginModel.SuDung.ToString()), new Claim("BiKhoa", loginModel.BiKhoa) 
		};

		//2. Tạo token
		var key = new SymmetricSecurityKey(System.Text.Encoding.UTF8.GetBytes(_configuration["JWT:Key"])); 
		var creds = new SigningCredentials(key, SecurityAlgorithms.HmacSha256); 
		var expires = DateTime.UtcNow.AddHours(3); // Token sống trong 3 giờ 
		var token = new JwtSecurityToken( 
			issuer: _configuration["JWT:Issuer"], 
			audience: _configuration["JWT:Audience"], 
			claims: claims, expires: expires, 
			signingCredentials: creds 
		);

		var tokenString = new JwtSecurityTokenHandler().WriteToken(token);

		// Score là thời gian hết hạn dạng Unix timestamp 
		var expireTimeUnix = ((DateTimeOffset)expires).ToUnixTimeSeconds();

		// Thêm token vào Sorted Set
		await_redisDatabase.SortedSetAddAsync("whitelist", tokenString, expireTimeUnix); 
		return tokenString;
	}


	//Tự động xóa token hết hạn bằng background task
	public async Task RemoveTokenExpiredTokenAsync(){
		var now = DateTimeOffset.UtcNow.ToUnixTimeSeconds();
		
		//Xóa các phần tử có score (Thời gian hết hạn) nhỏ hơn thời gian hiện tại
		await _redisDB.SortedSetRemoveRangeByScoreAsync("whitelist", double.NegativeInfinity, now);
	}
}
```

2. _Kiểm tra Token từ whiteList:_
```
public async Task<bool> IsTokenValidAsync(string token){
	var score = await _redisDB.SortedSetScoreAsync("whitelist", token);
		return score.HasValue; //Token hợp lệ nếu tồn tại trong whitelist
}
```

3. Kiểm tra các token mỗi phút 1 lần và tự động xóa nếu thấy token hết hạn
```
public class TokenCleanupService : BackgroundService { 
	private readonly TokenService _tokenService; 
	public TokenCleanupService(TokenService tokenService) { 
		_tokenService = tokenService; 
	}

	protected override async Task ExecuteAsync(CancellationToken stoppingToken){
		while (!stoppingToken.IsCancellationRequested) { 
		
			// Xóa token hết hạn mỗi phút 
			await _tokenService.RemoveExpiredTokensAsync(); 
			await Task.Delay(TimeSpan.FromMinutes(1), stoppingToken); 
		}
	}
}
```
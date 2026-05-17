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
```csharp
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
	public TokenService
	(IConnectionMultiplexer redisConnection, IConfiguration configuration){
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
		new Claim("SuDung", loginModel.SuDung.ToString()), 
		new Claim("BiKhoa", loginModel.BiKhoa) 
		};

		//2. Tạo token
		var key = new SymmetricSecurityKey
			(System.Text.Encoding.UTF8.GetBytes(_configuration["JWT:Key"])); 
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
		await _redisDB.
			SortedSetAddAsync("whitelist", tokenString, expireTimeUnix); 
		return tokenString;
	}

	//Tự động xóa token hết hạn bằng background task
	public async Task RemoveTokenExpiredTokenAsync(){
		var now = DateTimeOffset.UtcNow.ToUnixTimeSeconds();
		
		//Xóa các phần tử có score (Thời gian hết hạn) nhỏ hơn thời gian hiện tại
		await _redisDB.SortedSetRemoveRangeByScoreAsync
			("whitelist", double.NegativeInfinity, now);
	}
}
```

2. _Kiểm tra Token từ whiteList:_
```csharp
public async Task<bool> IsTokenValidAsync(string token){
	var score = await _redisDB.SortedSetScoreAsync("whitelist", token);
		return score.HasValue; //Token hợp lệ nếu tồn tại trong whitelist
}
```

3. Kiểm tra các token mỗi phút 1 lần và tự động xóa nếu thấy token hết hạn
```csharp
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

---
### **✅Chiến lượt thiết kế an toàn và hiệu quả**

##### 1. **`AccessToken` - Sử dụng WhiteList (Có thể không cần thiết lắm)**

- Không bắt buộc phải lưu vào **whitelist**, vì `accesstoken` có thời gian sống ngắn.
- Tuy nhiên, nếu ứng dụng yêu cầu bảo mật cao, **có thể thu hồi `accesstoken` ngay lập tức (Revoke)**, thì lưu trong **Redis Sorted Set** là hợp lý.
- **Khi nào cần?**  
    Nếu ứng dụng yêu cầu khả năng **thu hồi tức thời (revoke)** AccessToken (ví dụ: người dùng bị khóa tài khoản hoặc đăng xuất từ xa), thì việc lưu vào **Whitelist** (Redis Sorted Set) là hợp lý. Trong tài liệu của bạn, bạn đã triển khai điều này bằng cách lưu token với **score** là thời gian hết hạn (Unix timestamp), giúp dễ dàng kiểm tra và xóa token hết hạn.
- **Khi nào không cần?**  
    Nếu hệ thống ưu tiên hiệu năng cao và chấp nhận AccessToken hết hạn tự nhiên (dựa vào TTL trong JWT), thì không cần lưu vào Whitelist. Điều này giảm tải cho Redis và đơn giản hóa quy trình.

**🔥 Gợi ý:**
- Nếu muốn hiệu năng cao và chấp nhận token hết hạn tự nhiên $\to$ Không cần lưu `AccessToken` vào **WhiteList**
- Nếu muốn **Revoke ngay lập tức** (ví dụ: bị khóa tài khoản, đăng xuất từ xa) $\to$ nê dùng  **WhiteList**
- Thêm thông tin metadata (ví dụ: IP, UserAgent) vào Whitelist để phát hiện hành vi bất thường (như token được dùng từ IP khác).
- Nếu không dùng Whitelist, cần đảm bảo khóa bí mật (secret key) tạo token đủ mạnh và được bảo vệ tốt.

##### 2. **`RefreshToken` - áp dụng WhiteList & BlackList:**

- Vì thời gian sống lâu, nên `RefreshToken` phải được kiểm soát chặt chẽ hơn.
	- Lưu `RefreshToken` vào `WhiteList` (thường là Redis Sorted Set hoặc Hash).
	- Khi bị lạm dụng (reuse, lộ token) $\to$ đưa $\to$ `RefreshToken` vào **BlackList** để chặn lần dùng tiếp theo.
- **Lưu vào Whitelist ⚪:**
	- **Bắt buộc:**  
		RefreshToken phải được lưu vào Whitelist để đảm bảo chỉ những token hợp lệ mới được chấp nhận. Trong tài liệu, bạn đã gợi ý lưu với key dạng refresh:{user_id}:{token_id} và metadata (IP, UserAgent, thời gian hết hạn), điều này rất hợp lý.
	- **Lợi ích:**
		- Kiểm tra nhanh tính hợp lệ của token.
		- Dễ dàng thu hồi khi cần (logout, đổi mật khẩu, khóa tài khoản).
- **Áp dụng Blacklist ⚫:**
	- **Khi nào cần?**  
		Khi RefreshToken bị lạm dụng (reuse, lộ token), hoặc khi người dùng chủ động đăng xuất/đổi mật khẩu, cần đưa token cũ vào **Blacklist** để ngăn sử dụng tiếp. TTL của Blacklist nên bằng thời gian sống còn lại của token.
	- **Cách triển khai:**
		- Lưu trong Redis với key như blacklist:{token_id}.
		- Kiểm tra Blacklist trước khi chấp nhận RefreshToken.

- **Thiết kế tối ưu:**
    - **Whitelist:**
        - Key: refresh:{user_id}:{token_id} (đảm bảo uniqueness cho mỗi token của user).
        - Value: Metadata (IP, UserAgent, expires_at).
        - TTL: Bằng thời gian sống của RefreshToken.
    - **Blacklist:**
        - Key: blacklist:{token_id}.
        - TTL: Thời gian sống còn lại của token cũ.
    - **Quy trình cấp token mới:**
        1. Kiểm tra RefreshToken trong Whitelist.
        2. Kiểm tra không có trong Blacklist.
        3. Nếu hợp lệ: cấp AccessToken mới, tạo RefreshToken mới, xóa token cũ khỏi Whitelist, thêm vào Blacklist.
- **Gợi ý an toàn:**
    - Sử dụng **token rotation**: Mỗi lần refresh, sinh RefreshToken mới và invalidate token cũ. Điều này giảm nguy cơ reuse token bị lộ.
    - Giới hạn số lần refresh cho một chuỗi token để tránh lạm dụng.

**🧠 Gợi ý thiết kế:**
- Khi sinh ra `RefreshToken`, lưu:
```
Key: refresh:{user_id}:{token_id}
Value: metadata (IP, UserAgent, thời gian hết hạn)
TTL: theo thời gian sống của token
```

-  Khi logout, đổi mật khẩu, bị ban,…:
    
    - **Xóa `refresh:{...}` khỏi whitelist**    
    - **Thêm vào `blacklist:{token_id}` với TTL = thời gian còn lại**

- Khi nhận yêu cầu `refresh token` mới:
    
    - Kiểm tra xem `refresh token` có tồn tại trong whitelist.
    - Nếu có, thì cấp `AccessToken` mới và thay `RefreshToken` cũ.
    - **Invalidate** token cũ bằng cách:
        - Xóa khỏi whitelist.
        - Thêm vào blacklist.

|Token|Lưu vào Whitelist|Có Blacklist?|TTL trong Redis|Ghi chú|
|---|---|---|---|---|
|AccessToken|Có thể (nếu cần revoke nhanh)|Không cần thiết|= thời gian sống|Có thể bỏ nếu hệ thống hiệu năng cao|
|RefreshToken|Bắt buộc|Nên có|= thời gian sống|Phải kiểm soát để tránh reuse và lộ|

---
#### **Thông tin bổ sung:** 

###### 1. **Thêm cơ chế Token Rotation:**
Trong `GenerateTokenAndAddToWhitelist`, khi cấp `RefreshToken` mới, `invalidate token cũ` và thêm vào Blacklist:

```csharp
public async Task<(string AccessToken, string RefreshToken)> RefreshTokenAsync(string oldRefreshToken)
{
    if (!await IsTokenValidAsync(oldRefreshToken)) return (null, null);
    
    // Xóa token cũ khỏi Whitelist
    await _redisDB.SortedSetRemoveAsync("whitelist", oldRefreshToken);
    
    // Thêm vào Blacklist
    var expires = GetTokenExpiration(oldRefreshToken); // Hàm tự viết để lấy thời gian hết hạn
    
    await _redisDB.StringSetAsync($"blacklist:{oldRefreshToken}", "revoked", expires - DateTime.UtcNow);
    
    // Tạo token mới
    var newAccessToken = await GenerateTokenAndAddToWhitelist(...);
    var newRefreshToken = GenerateNewRefreshToken(...);
    return (newAccessToken, newRefreshToken);
}
```

###### 2. **Kiểm tra metadata:**
Khi validate RefreshToken, so sánh IP/UserAgent:

```csharp
public async Task<bool> IsTokenValidAsync(string token, string ip, string userAgent)
{
    var metadata = await _redisDB
	    .StringGetAsync($"refresh:metadata:{token}");
    
    if (metadata.IsNull || !metadata.ToString().Contains(ip)) 
	    return false;
    
    return await _redisDB.SortedSetScoreAsync
	    ("whitelist", token).HasValue;
}
```

###### 3. **Tối ưu hiệu năng:**
Dùng Redis EXPIRE cho từng key thay vì background task liên tục quét toàn bộ Sorted Set.

###### 4. **Logging:**
Thêm logging khi token bị xóa hoặc thêm vào Blacklist để dễ dàng kiểm tra sự cố bảo mật.
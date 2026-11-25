# Xác thực Service-to-Service cho Cronjob (Service B gọi API của Service A)

## 1. Bài toán

- **Service B** chạy cronjob theo lịch, trong mỗi lần chạy cần gọi API của **Service A**.
- Các API của Service A yêu cầu token (`Authorization: Bearer <token>`).
- Không có người dùng thật tham gia: đây là **machine-to-machine (M2M)**.

Câu hỏi cần trả lời không chỉ là *"cất token ở đâu"* mà là bốn câu hỏi gắn liền nhau:

| # | Câu hỏi | Vì sao quan trọng |
|---|---------|-------------------|
| 1 | **Danh tính**: B chứng minh "tôi là B" bằng gì? | Quyết định ai chịu trách nhiệm cho request |
| 2 | **Credential**: thứ gì được cấp cho B và sống bao lâu? | Thời gian sống quyết định thiệt hại khi lộ |
| 3 | **Phạm vi quyền**: B được làm gì trên A? | Nguyên tắc least privilege |
| 4 | **Vòng đời**: cấp, làm mới, thu hồi, xoay vòng thế nào? | Quyết định chi phí vận hành lâu dài |

### Giả định (cần xác nhận)

1. A và B do cùng một tổ chức kiểm soát, chạy trong mạng nội bộ hoặc VPC.
2. B chạy trên VM, container hoặc Kubernetes CronJob; tần suất từ vài phút đến vài giờ một lần.
3. Chưa có Identity Provider (IdP) như Keycloak, Auth0, Cognito; chưa chắc đã có service mesh.
4. Stack chính là Go. Ví dụ code dùng Go, ý tưởng áp dụng được cho mọi ngôn ngữ.
5. Hạ tầng chưa chắc ở cloud nào cụ thể.

## 2. Khái niệm nền

**Credential dài hạn và ngắn hạn**

- *Dài hạn* (API key, token không hết hạn): lộ một lần thì kẻ tấn công dùng được đến khi có người phát hiện và thu hồi.
- *Ngắn hạn* (token sống 5–15 phút): lộ cũng chỉ dùng được trong cửa sổ ngắn. Đổi lại cần cơ chế **làm mới** tự động.

**Secret zero**: để lấy một secret từ nơi lưu trữ (Vault, Secrets Manager), B cũng phải xác thực với nơi đó. Nếu "chìa khóa mở két" lại là một secret tĩnh đặt trong cấu hình thì vấn đề chỉ bị dời chỗ. Cách giải quyết bền vững là để **nền tảng chạy workload chứng minh danh tính của B** (Kubernetes ServiceAccount, IAM Role, SPIFFE ID), không dùng secret do con người đặt.

**Định danh B, không định danh người dùng**: token nên đại diện cho *service B* (claim `sub = service-b`), kèm `scope` giới hạn. Token do người dùng đăng nhập sinh ra có hạn ngắn, gắn phiên người dùng, không phù hợp cho cronjob.

## 3. Các phương án

### Phương án A: OAuth2 Client Credentials (có IdP)

```
B ──(client_id + client_secret)──▶ IdP ──▶ access_token (JWT, 5–15 phút)
B ──(Bearer access_token)────────▶ A (xác thực chữ ký JWT bằng JWKS của IdP)
```

- **Cách hoạt động**: B đăng ký như một *client* trên IdP. Mỗi lần chạy, B đổi `client_id/secret` lấy access token ngắn hạn, rồi gọi A. A kiểm tra chữ ký, `exp`, `aud`, `scope`.
- **Ưu**: chuẩn công nghiệp, nhiều thư viện; token ngắn hạn; scope và audience rõ ràng; thu hồi và audit tập trung; A không cần biết secret của B.
- **Nhược**: cần vận hành IdP (hoặc dùng dịch vụ managed); thêm một thành phần phụ thuộc; `client_secret` vẫn là secret tĩnh phải bảo vệ (xem Phương án D để giải bài toán này).
- **Điều kiện áp dụng**: có nhiều service, nhiều client; hoặc đã có IdP; cần audit và quản lý quyền tập trung.

### Phương án B: A tự cấp token cho B (token endpoint tối giản)

```
B ──(client_id + client_secret)──▶ A /oauth/token ──▶ JWT ngắn hạn (ký bằng private key của A)
B ──(Bearer JWT)─────────────────▶ A /api/... (xác thực bằng public key)
```

- **Cách hoạt động**: A tự cài một endpoint theo giao thức Client Credentials, giữ danh sách client (băm secret), ký JWT bằng khóa riêng.
- **Ưu**: không cần IdP; ít thành phần; vẫn có token ngắn hạn và scope; dễ triển khai khi chỉ có một hai client.
- **Nhược**: A gánh thêm trách nhiệm bảo mật (quản lý client, rotate khóa ký, chống brute-force); không tái sử dụng được cho service khác nếu hệ thống lớn lên; dễ tự phát minh sai (ví dụ quên kiểm tra `aud`, `alg`).
- **Điều kiện áp dụng**: hệ thống nhỏ (2–5 service), chưa muốn vận hành IdP, đội đủ kiến thức bảo mật để tự cài.

### Phương án C: mTLS (chứng chỉ hai chiều)

```
B ──(TLS handshake, B trình client cert do CA nội bộ cấp)──▶ A
A kiểm tra cert → biết đó là B (danh tính nằm trong cert, không cần token)
```

- **Cách hoạt động**: danh tính nằm ở chứng chỉ X.509. Thường triển khai qua service mesh (Istio, Linkerd) hoặc SPIFFE/SPIRE; cert có hạn ngắn và được tự động xoay vòng.
- **Ưu**: không có token để lộ trong header; xác thực ở tầng transport; hợp với mô hình zero-trust; xoay vòng tự động.
- **Nhược**: hạ tầng PKI/mesh phức tạp; khó debug; tốn công nếu chưa có nền tảng; mTLS chứng minh *B là ai*, còn *B được làm gì* vẫn cần lớp phân quyền riêng.
- **Điều kiện áp dụng**: đã chạy Kubernetes với mesh, hoặc yêu cầu bảo mật cao (tài chính, y tế).

### Phương án D: Workload Identity (nền tảng chứng minh danh tính cho B)

Dùng danh tính do chính nền tảng chạy cấp, không có secret tĩnh nào do con người đặt.

- **Kubernetes**: *projected ServiceAccount token* (JWT ngắn hạn, kubelet tự làm mới). A (hoặc IdP) xác thực token này qua OIDC/TokenReview.
- **AWS**: IAM Role gắn vào task/pod (IRSA, ECS Task Role) → ký request bằng SigV4, hoặc dùng role đó để đọc secret từ Secrets Manager.
- **GCP/Azure**: Workload Identity / Managed Identity.
- **Vault**: B đăng nhập bằng auth method `kubernetes` hoặc `aws` (thay cho AppRole + secret tĩnh), sau đó đọc secret hoặc xin credential động (dynamic secret) ngắn hạn.

**Ưu**: giải quyết triệt để secret zero; không có secret để rò rỉ trong repo, `.env`, CI; danh tính gắn với workload thật.
**Nhược**: phụ thuộc nền tảng (cloud/K8s); cần cấu hình RBAC/IAM đúng; khó chạy cục bộ nên cần cơ chế fallback cho môi trường dev.
**Điều kiện áp dụng**: chạy trên K8s hoặc cloud lớn. Thường **kết hợp** với A hoặc B: workload identity thay cho `client_secret`.

### Phương án E: API key có quản lý vòng đời

- **Cách hoạt động**: A cấp một API key riêng cho B (có scope, có ngày hết hạn, lưu dạng băm phía A). B giữ key trong secret store và gửi qua header. Có quy trình rotation định kỳ với **hai key cùng hợp lệ trong thời gian chuyển tiếp**.
- **Ưu**: đơn giản nhất; không cần giao thức hay hạ tầng đặc biệt; dễ hiểu với mọi đội.
- **Nhược**: credential vẫn là dài hạn (tính theo tuần/tháng); xoay vòng phụ thuộc kỷ luật vận hành; không tự hết hạn trong vài phút; thiếu chuẩn nên mỗi nơi một kiểu.
- **Điều kiện áp dụng**: giai đoạn đầu, hệ thống nội bộ nhỏ, rủi ro thấp; cần chạy được ngay.

## 4. So sánh

| Tiêu chí | A. Client Credentials + IdP | B. A tự cấp token | C. mTLS | D. Workload Identity | E. API key có vòng đời |
|---|---|---|---|---|---|
| Thời gian sống credential dùng gọi API | Ngắn (phút) | Ngắn (phút) | Ngắn (cert giờ/ngày) | Ngắn (phút/giờ) | Dài (tuần/tháng) |
| Secret tĩnh phải bảo vệ | `client_secret` (trừ khi kết hợp D) | `client_secret` (trừ khi kết hợp D) | Không (khóa riêng trong workload) | Không | API key |
| Phạm vi quyền (scope) | Có, chuẩn | Có, tự cài | Cần lớp riêng | Qua IAM/RBAC | Có, tự cài |
| Thu hồi tức thì | Tốt | Tốt | Trung bình (CRL/OCSP, hoặc chờ cert hết hạn) | Tốt | Tốt (xóa key) |
| Độ phức tạp vận hành | Trung bình–cao | Thấp–trung bình | Cao | Trung bình | Thấp |
| Phụ thuộc hạ tầng | IdP | Không | PKI/mesh | Cloud/K8s | Không |
| Mở rộng khi tăng số service | Rất tốt | Trung bình | Rất tốt | Rất tốt | Kém dần |
| Công sức triển khai ban đầu | Trung bình | Thấp | Cao | Trung bình | Rất thấp |

## 5. Lựa chọn theo bối cảnh

Không có phương án tốt nhất tuyệt đối. Quyết định phụ thuộc vào **hạ tầng sẵn có** và **mức rủi ro chấp nhận**:

| Bối cảnh | Nên cân nhắc | Lý do chính |
|---|---|---|
| Dev/PoC, 1–2 service, rủi ro thấp | **E** | Rẻ nhất, đủ dùng, dễ nâng cấp sau |
| Nhỏ, 2–5 service, chưa có IdP/K8s | **B** + secret store | Có token ngắn hạn mà không cần IdP |
| Đã có IdP, hoặc nhiều service đang tăng | **A** | Chuẩn, tái sử dụng, audit tập trung |
| Chạy trên K8s hoặc cloud | **D** (+ A hoặc B) | Loại bỏ secret tĩnh khỏi cấu hình |
| Yêu cầu zero-trust, đã có mesh | **C** (+ phân quyền) | Danh tính ở tầng transport |

**Lộ trình tiến hóa gợi ý** (mỗi bước độc lập, không phải làm lại từ đầu):

1. **Bước 1**: tách danh tính riêng cho B, đặt scope tối thiểu, đưa secret vào secret store.
2. **Bước 2**: chuyển sang token ngắn hạn (B hoặc A), B cache trong bộ nhớ và làm mới khi sắp hết hạn.
3. **Bước 3**: dùng workload identity để lấy `client_secret` (hoặc bỏ hẳn nó), loại bỏ secret zero.

Với bối cảnh giả định ở mục 1 (nhỏ–vừa, Go, chưa chắc có IdP), tổ hợp **B hoặc A + secret store + workload identity khi có điều kiện** cho cân bằng tốt nhất giữa bảo mật và công sức. Nếu thực tế khác giả định (ví dụ đã có Keycloak, hoặc chạy trên K8s), bảng ở trên đổi kết quả.

## 6. Hướng dẫn triển khai (Phương án A/B: Client Credentials)

### 6.1 Luồng tổng quát một lần chạy cronjob

```
cron tick
  └─ TokenProvider.Get()
        ├─ còn hạn (còn > 30s)  → dùng token trong bộ nhớ
        └─ sắp/đã hết hạn       → POST /oauth/token  (client_id, client_secret, scope)
  └─ gọi A với Authorization: Bearer <token>
        └─ nhận 401 → làm mới token, thử lại đúng 1 lần
```

**Vì sao cache trong bộ nhớ?** Cấp token là thao tác tốn kém hơn dùng token. Cache giúp B gọi endpoint cấp token một lần cho nhiều request thay vì một lần mỗi request. Vì token ngắn hạn, việc cache trong RAM không làm tăng rủi ro đáng kể so với ghi ra đĩa.

**Vì sao làm mới trước khi hết hạn (skew)?** Đồng hồ hai máy lệch nhau và độ trễ mạng khiến token "còn 1 giây" có thể hết hạn khi tới A. Chừa 30 giây là thói quen an toàn.

**Lưu ý với K8s CronJob**: mỗi lần chạy là một process mới nên cache không sống qua các lần chạy. Không sao: mỗi lần chạy lấy token một lần (một request nhỏ), vẫn rẻ hơn nhiều so với nguy cơ lưu token lâu dài.

### 6.2 Phía Service B: TokenProvider (Go)

```go
package auth

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"
	"sync"
	"time"
)

type TokenProvider struct {
	tokenURL     string
	clientID     string
	clientSecret string // nạp từ secret store lúc khởi động, không hard-code
	scope        string
	httpClient   *http.Client

	mu        sync.Mutex
	token     string
	expiresAt time.Time
}

const skew = 30 * time.Second

func NewTokenProvider(tokenURL, clientID, clientSecret, scope string) *TokenProvider {
	return &TokenProvider{
		tokenURL: tokenURL, clientID: clientID, clientSecret: clientSecret, scope: scope,
		httpClient: &http.Client{Timeout: 10 * time.Second},
	}
}

// Get trả token còn hạn; tự làm mới khi sắp hết hạn.
func (p *TokenProvider) Get(ctx context.Context) (string, error) {
	p.mu.Lock()
	defer p.mu.Unlock() // giữ khóa suốt lúc fetch: tránh nhiều goroutine cùng xin token

	if p.token != "" && time.Until(p.expiresAt) > skew {
		return p.token, nil
	}
	return p.refreshLocked(ctx)
}

// Invalidate xóa cache (dùng khi A trả 401).
func (p *TokenProvider) Invalidate() {
	p.mu.Lock()
	p.token, p.expiresAt = "", time.Time{}
	p.mu.Unlock()
}

func (p *TokenProvider) refreshLocked(ctx context.Context) (string, error) {
	form := url.Values{
		"grant_type": {"client_credentials"},
		"scope":      {p.scope},
	}
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, p.tokenURL, strings.NewReader(form.Encode()))
	if err != nil {
		return "", err
	}
	req.SetBasicAuth(p.clientID, p.clientSecret)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	resp, err := p.httpClient.Do(req)
	if err != nil {
		return "", fmt.Errorf("request token: %w", err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("token endpoint returned %d", resp.StatusCode) // không in body/secret
	}

	var out struct {
		AccessToken string `json:"access_token"`
		ExpiresIn   int    `json:"expires_in"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&out); err != nil {
		return "", err
	}
	p.token = out.AccessToken
	p.expiresAt = time.Now().Add(time.Duration(out.ExpiresIn) * time.Second)
	return p.token, nil
}
```

### 6.3 Phía Service B: gọi A kèm thử lại khi 401

```go
func CallA(ctx context.Context, tp *TokenProvider, c *http.Client, req *http.Request) (*http.Response, error) {
	for attempt := 0; attempt < 2; attempt++ {
		tok, err := tp.Get(ctx)
		if err != nil {
			return nil, err
		}
		r := req.Clone(ctx)
		r.Header.Set("Authorization", "Bearer "+tok)

		resp, err := c.Do(r)
		if err != nil {
			return nil, err
		}
		if resp.StatusCode != http.StatusUnauthorized {
			return resp, nil
		}
		resp.Body.Close()
		tp.Invalidate() // token có thể đã bị xoay khóa hoặc thu hồi → xin lại, chỉ 1 lần
	}
	return nil, fmt.Errorf("unauthorized after token refresh")
}
```

> Request có body cần `GetBody` để Clone gửi lại được; với `GET` hoặc body nhỏ có thể tạo lại request.

### 6.4 Phía Service A: cấp và xác thực token (Phương án B)

Dùng thư viện `github.com/golang-jwt/jwt/v5`. Ký bằng **RS256** (khóa bất đối xứng) để các service khác verify chỉ bằng public key.

```go
// Cấp token: POST /oauth/token
func (s *Server) IssueToken(w http.ResponseWriter, r *http.Request) {
	id, secret, ok := r.BasicAuth()
	client, found := s.clients[id]
	// so sánh hằng thời gian trên hash, không so chuỗi thô
	if !ok || !found || !client.VerifySecret(secret) {
		http.Error(w, `{"error":"invalid_client"}`, http.StatusUnauthorized)
		return
	}
	now := time.Now()
	claims := jwt.MapClaims{
		"iss":   "https://service-a.internal",
		"sub":   id,                      // danh tính service B
		"aud":   "service-a",             // token chỉ có giá trị với A
		"scope": client.AllowedScopes,    // ví dụ "reports:read jobs:run"
		"iat":   now.Unix(),
		"exp":   now.Add(10 * time.Minute).Unix(),
		"jti":   uuid.NewString(),        // phục vụ audit/thu hồi
	}
	tok := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)
	tok.Header["kid"] = s.keyID           // hỗ trợ xoay khóa ký
	signed, _ := tok.SignedString(s.privateKey)

	json.NewEncoder(w).Encode(map[string]any{
		"access_token": signed, "token_type": "Bearer", "expires_in": 600,
	})
}
```

```go
// Middleware xác thực + kiểm tra scope
func (s *Server) RequireScope(scope string, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		raw := strings.TrimPrefix(r.Header.Get("Authorization"), "Bearer ")
		claims := jwt.MapClaims{}
		_, err := jwt.ParseWithClaims(raw, claims, func(t *jwt.Token) (any, error) {
			return s.publicKey, nil
		},
			jwt.WithValidMethods([]string{"RS256"}), // chặn tấn công đổi thuật toán (alg=none/HS256)
			jwt.WithAudience("service-a"),
			jwt.WithIssuer("https://service-a.internal"),
			jwt.WithExpirationRequired(),
			jwt.WithLeeway(5*time.Second),
		)
		if err != nil {
			http.Error(w, "unauthorized", http.StatusUnauthorized)
			return
		}
		if !hasScope(claims["scope"], scope) {
			http.Error(w, "forbidden", http.StatusForbidden) // 403: đã xác thực nhưng thiếu quyền
			return
		}
		next.ServeHTTP(w, r)
	})
}
```

**Vì sao ép `WithValidMethods`?** Nếu không, kẻ tấn công có thể gửi token với `alg=none` hoặc đổi sang HS256 dùng chính public key làm khóa HMAC. Đây là lỗi kinh điển khi tự cài JWT.

**Vì sao tách 401 và 403?** 401 = chưa chứng minh được danh tính (B nên thử xin token lại); 403 = biết B là ai nhưng không đủ quyền (thử lại vô ích). Tách rõ giúp B xử lý lỗi đúng.

### 6.5 Nạp `client_secret` an toàn

- **Không** commit vào repo, không in ra log, không đưa vào image Docker.
- **K8s**: dùng `Secret` mount thành file hoặc dùng External Secrets Operator / Secrets Store CSI để đồng bộ từ Vault/Secrets Manager.
- **Vault**: B đăng nhập bằng auth method `kubernetes` (workload identity), đọc `client_secret` lúc khởi động.
- **Khi chạy cục bộ**: dùng biến môi trường chỉ cho dev, với client và scope riêng cho dev, tránh dùng chung với production.

### 6.6 Xoay vòng secret không downtime

1. Tạo **secret mới** cho client B trong khi **secret cũ vẫn còn hợp lệ** (A chấp nhận hai secret).
2. Cập nhật secret store; các lần chạy tiếp theo của B sẽ đọc secret mới.
3. Theo dõi đến khi không còn request dùng secret cũ, rồi thu hồi secret cũ.

Với khóa ký JWT, dùng `kid` + danh sách public key (JWKS) để A xác thực được cả token ký bằng khóa cũ lẫn mới trong thời gian chuyển tiếp.

### 6.7 Quan sát và vận hành

- **Log**: ghi `client_id`, `jti`, `scope`, kết quả cấp/từ chối. **Không bao giờ** ghi token hay secret.
- **Metric**: tỷ lệ 401/403, số lần cấp token, độ trễ endpoint cấp token, số lần làm mới khi 401 (tăng đột biến = có xoay khóa hoặc sự cố).
- **Cảnh báo**: cấp token thất bại liên tục; secret sắp tới hạn xoay; client lạ xin token.
- **Rate limit** endpoint cấp token để chống dò mật khẩu.

## 7. Thông tin cần xác nhận để chốt phương án

1. Token hiện tại đại diện cho *người dùng* hay *service B*? Có thời hạn không?
2. B và A chạy ở đâu (K8s, VM, cloud nào)? Đã có IdP hoặc Vault chưa?
3. Số lượng service sẽ gọi A trong 6–12 tháng tới?
4. Dữ liệu và thao tác A cung cấp nhạy cảm đến mức nào (đọc báo cáo hay thay đổi dữ liệu/tiền)?
5. A có thể sửa code để thêm token endpoint không, hay A là hệ thống đóng?

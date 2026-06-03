package middewarre

import (
	"database/sql"
	"net/http"
	"strings"
	"time"

	"github.com/GiaBao0510/go-apikey-demo/internal/keygen"
	"github.com/gin-gonic/gin"
	"github.com/lib/pq"
)

// APIKeyAuth trả về một Gin Middleware xác thực API key
// - db: kết nối đến cơ sở dữ liệu để kiểm tra API key
// - requiredScope  phạm vi yêu cầu để truy cập tài nguyên
func APIKeyAuth(db *sql.DB, requiredScope string) gin.HandlerFunc {
	return func(c *gin.Context) {
		// ----- Bước 1 lấy Key từ header -----
		rawkey := extractKeyFromRequest(c)
		if rawkey == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "API key is required",
				"hint":  "provide key via X-API-Key header or Authorization: Bearer",
			})
			return
		}

		// ----- Bước 2 kiểm tra key trong DB -----
		keyHash := keygen.Hash(rawkey)

		row := db.QueryRow(`
            SELECT id, user_id, scopes, rate_limit, is_active, expires_at
            FROM api_keys
            WHERE key_hash = $1 LIMIT 1;
        `, keyHash)

		var (
			keyID     string
			userID    int
			scopes    []string
			rateLimit int
			isActive  bool
			expiresAt *time.Time
		)

		// Nếu không tìm thấy key nào khớp với hash, trả về lỗi 401
		err := row.Scan(&keyID, &userID, pq.Array(&scopes), &rateLimit, &isActive, &expiresAt)
		if err == sql.ErrNoRows {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "invalid API key",
			})
			return
		}
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"error": "Internal server error - Details: " + err.Error(),
			})
			return
		}

		// ----- Bước 3 kiểm tra trạng thái key -----
		if !isActive {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "API key is inactive",
			})
			return
		}

		// ----- Bước 4 kiểm tra hạn sử dụng -----
		if expiresAt != nil && time.Now().After(*expiresAt) {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "API key has expired",
			})
			return
		}

		// ----- Bước 5 kiểm tra phạm vi ----
		if requiredScope != "" && !hasScope(scopes, requiredScope) {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{
				"error":    "insufficient scope",
				"required": requiredScope,
			})
			return
		}

		// ----- Bước 6: cập nhật lại thông tin key (last_used) -----
		go func(id string) {
			_, _ = db.Exec(`UPDATE api_keys SET last_used = $1 WHERE id = $2`, time.Now(), id)
		}(keyID)

		// ----- Bước 7: truyền context cho handler phía sau ------
		c.Set("api_key_id", keyID)
		c.Set("user_id", userID)
		c.Set("scopes", scopes)
		c.Next() // tiếp tục xử lý request với handler phía sau
	}
}

// extractKeyFromRequest lấy API key từ header
func extractKeyFromRequest(c *gin.Context) string {
	// Cách 1 lấy Header X-API-KEY
	if key := c.GetHeader("X-API-KEY"); key != "" {
		return key
	}

	// Cách 2 Authorization Bearer <API_KEY>
	auth := c.GetHeader("Authorization")
	if strings.HasPrefix(auth, "Bearer ") {
		return strings.TrimSpace(strings.TrimPrefix(auth, "Bearer "))
	}

	return ""
}

// Kiểm tra xem scopes string có chứa requiredScope  không
func hasScope(scopes []string, required string) bool {
	for _, s := range scopes {
		if strings.TrimSpace(s) == required {
			return true
		}
	}
	return false
}

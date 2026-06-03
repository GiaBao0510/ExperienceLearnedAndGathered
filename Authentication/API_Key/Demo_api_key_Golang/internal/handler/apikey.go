package handler

import (
	"database/sql"
	"net/http"
	"time"

	"github.com/GiaBao0510/go-apikey-demo/internal/keygen"
	"github.com/GiaBao0510/go-apikey-demo/internal/model"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/lib/pq"
)

type APIKeyHandler struct {
	db *sql.DB
}

func NewAPIKeyHandler(db *sql.DB) *APIKeyHandler {
	return &APIKeyHandler{db: db}
}

// CreateAPIKey xử lý ở giao thức HTTP trong phương thức POST để tạo
// POST: /api-keys
func (h *APIKeyHandler) CreateAPIKey(c *gin.Context) {
	var req model.CreateAPIKeyRequest

	// Validate request body
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Đặt giá trị mặt định cho phạm vi
	if len(req.Scopes) == 0 {
		req.Scopes = []string{"read"}
	}

	// --- sinh key ngẫu nhiên và hash key ---
	plainKey, err := keygen.Generate("sk_live")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to generate key"})
		return
	}

	keyHash := keygen.Hash(plainKey) // lưu key mã hóa vào database
	keyHint := plainKey[:10]         // lưu 10 ký tự đầu gơi ý vào database

	// --- lưu vào database ---
	id := uuid.New().String()
	_, err = h.db.Exec(`
	INSERT INTO api_keys (id, user_id, name, key_hint, key_hash, scopes, rate_limit)
	VALUES ($1, $2, $3, $4, $5, $6, $7)`,
		id, req.UserID, req.Name, keyHint, keyHash, pq.Array(req.Scopes), req.RateLimit,
	)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to save key"})
		return
	}

	// trả về response chứa key gốc cho người dùng (Quan trọng là về API key này một lần duy nhất)
	c.JSON(http.StatusCreated, model.CreateAPIKeyResponse{
		APIKey: model.APIKey{
			ID:        id,
			UserID:    req.UserID,
			Name:      req.Name,
			KeyHint:   keyHint,
			Scopes:    req.Scopes,
			RateLimit: req.RateLimit,
			IsActive:  true,
			CreatedAt: time.Now(),
		},
		PlainKey: plainKey,
		Message:  "API Key này chỉ được hiển thị một lần duy nhất, hãy lưu lại cẩn thận!",
	})
}

// ListAPIKeys xử lý GET /api-keys?user_id=xxx để liệt kê tất cả API key của một người dùng
func (h *APIKeyHandler) ListAPIKeys(c *gin.Context) {

	// Lấy user_id từ query parameter
	userID := c.Query("user_id")
	if userID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "user_id is required"})
		return
	}

	// Truy vấn database để lấy danh sách API key của người dùng
	rows, err := h.db.Query(`
	SELECT id, user_id, name, key_hint, scopes, rate_limit, is_active, last_used, created_at
        FROM api_keys
		WHERE user_id = $1
		ORDER BY created_at DESC
	`, userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "query failed - Details: " + err.Error()})
		return
	}
	defer rows.Close() // Đóng kết nối sau khi hoàn thành

	var Keys []model.APIKey
	for rows.Next() {
		var k model.APIKey
		rows.Scan(&k.ID, &k.UserID, &k.Name, &k.KeyHint, &k.Scopes, &k.RateLimit, &k.IsActive, &k.LastUsedAt, &k.CreatedAt)
		Keys = append(Keys, k)
	}
	c.JSON(http.StatusOK, gin.H{
		"api_keys": Keys, "total": len(Keys),
	})
}

// Thu hồi API key bằng cách đặt is_active = false
// DELETE /api-keys/:id
func (h *APIKeyHandler) RevokeAPIKey(c *gin.Context) {
	// Lấy id từ path parameter
	keyID := c.Param("id")
	result, err := h.db.Exec(`UPDATE api_keys SET is_active = false WHERE id = $1`, keyID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "revoke failed - Details: " + err.Error()})
		return
	}
	affected, _ := result.RowsAffected()
	if affected == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "API key not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "API key revoked successfully"})
}

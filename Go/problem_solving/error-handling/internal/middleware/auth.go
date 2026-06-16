package middleware

import (
    "github/GiaBao0510/error-handling/internal/domain"
    "net/http"
    "strings"

    "github.com/gin-gonic/gin"
)

func Auth() gin.HandlerFunc {
    return func(ctx *gin.Context) {
        authHeader := ctx.GetHeader("Authorization")
        
        // Kiểm tra header có tồn tại và bắt đầu với "Bearer "
        if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
            ctx.AbortWithStatusJSON(
                http.StatusUnauthorized,
                domain.NewUnauthorizedError(),
            )
            return
        }

        // Lấy phần token sau "Bearer "
        token := strings.TrimPrefix(authHeader, "Bearer ")

        if !isValidToken(token) {
            ctx.AbortWithStatusJSON(http.StatusUnauthorized, domain.NewTokenInvalidError())
            return
        }

        ctx.Next()
    }
}

func isValidToken(token string) bool {
    return token == "valid-token" 
}
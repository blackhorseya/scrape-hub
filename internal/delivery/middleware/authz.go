package middleware

import (
	"net/http"

	"github.com/auth0/go-jwt-middleware/v2/validator"
	"github.com/gin-gonic/gin"
)

// AuthzMiddleware 提供授權檢查的中介層實作
type AuthzMiddleware struct {
	allowedUserIDs []string
}

// NewAuthzMiddleware 建立新的授權檢查中介層實例
func NewAuthzMiddleware() *AuthzMiddleware {
	return &AuthzMiddleware{
		allowedUserIDs: []string{
			"google-oauth2|117379526721591132148",
			// 在此加入其他允許的使用者 ID
		},
	}
}

// EnsureAuthorized 確保請求的使用者有權限存取資源
func (m *AuthzMiddleware) EnsureAuthorized() gin.HandlerFunc {
	return func(c *gin.Context) {
		claims, exists := c.Get("claims")
		if !exists {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "未經過身分驗證",
			})
			return
		}

		validatedClaims, ok := claims.(*validator.ValidatedClaims)
		if !ok {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "無效的 claims 格式",
			})
			return
		}

		userID := validatedClaims.RegisteredClaims.Subject
		isAllowed := false
		for _, allowedID := range m.allowedUserIDs {
			if userID == allowedID {
				isAllowed = true
				break
			}
		}

		if !isAllowed {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{
				"error": "無權限存取此資源",
			})
			return
		}

		c.Next()
	}
}

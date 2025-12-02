package middleware

import (
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

// AuthRequired: เช็คว่าล็อกอินหรือยัง?
func AuthRequired() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		user := session.Get("user_id")
		if user == nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "กรุณาเข้าสู่ระบบ"})
			return
		}
		c.Next()
	}
}

// AdminOnly: เช็คว่าเป็น Admin เท่านั้น? (ห้าม Staff เข้า)
func AdminOnly() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		role := session.Get("role")

		if role != "admin" {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "คุณไม่มีสิทธิ์เข้าถึงส่วนนี้ (Admin Only)"})
			return
		}
		c.Next()
	}
}

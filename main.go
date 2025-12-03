package main

import (
	"backend/config"
	"backend/controllers"
	"backend/middleware"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

func main() {
	config.ConnectDB()
	r := gin.Default()

	// ✅ 1. แก้ไข CORS: เพิ่ม "PATCH" และ "OPTIONS"
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173", "https://shop9bath.pages.dev", "https://shop9bath.online"},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	store := cookie.NewStore([]byte("secret_key"))
	r.Use(sessions.Sessions("mysession", store))
	r.Static("/uploads", "./uploads")

	api := r.Group("/api")
	{
		api.POST("/register", controllers.Register)
		api.POST("/login", controllers.Login)
		api.POST("/logout", controllers.Logout)

		authorized := api.Group("/")
		authorized.Use(middleware.AuthRequired())
		{
			authorized.GET("/products", controllers.ListProducts)
			authorized.POST("/products", controllers.AddProduct)
			authorized.DELETE("/products/:id", controllers.DeleteProduct)

			// โซน Admin
			admin := authorized.Group("/")
			admin.Use(middleware.AdminOnly())
			{
				admin.GET("/users", controllers.ListUsers)
				admin.DELETE("/users/:id", controllers.DeleteUser)
				// ✅ 2. เพิ่ม Route สำหรับเปลี่ยนสถานะ
				admin.PATCH("/users/:id/status", controllers.UpdateUserStatus)
			}
		}
	}

	r.Run(":8080")
}

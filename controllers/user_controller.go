package controllers

import (
	"backend/config"
	"backend/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// ดึงรายชื่อ
func ListUsers(c *gin.Context) {
	var users []models.User
	config.DB.Select("id, username, fullname, role, status").Find(&users)
	c.JSON(http.StatusOK, users)
}

// ลบ User
func DeleteUser(c *gin.Context) {
	id := c.Param("id")
	if err := config.DB.Delete(&models.User{}, id).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "ลบไม่สำเร็จ"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "ลบสำเร็จ"})
}

// ✅ ฟังก์ชันเปลี่ยนสถานะ (ทำงานจริง)
func UpdateUserStatus(c *gin.Context) {
	id := c.Param("id")

	// รับค่า status จากหน้าบ้าน (active, suspended, fired)
	var input struct {
		Status string `json:"status"`
	}
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ข้อมูลไม่ถูกต้อง"})
		return
	}

	// สั่งอัปเดตลง Database จริงๆ
	if err := config.DB.Model(&models.User{}).Where("id = ?", id).Update("status", input.Status).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "อัปเดตสถานะไม่สำเร็จ"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "อัปเดตสถานะเรียบร้อย", "status": input.Status})
}

package controllers

import (
	"backend/config"
	"backend/models"
	"fmt"
	"net/http"
	"path/filepath"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func ListProducts(c *gin.Context) {
	var products []models.Product
	config.DB.Find(&products)
	c.JSON(http.StatusOK, products) // ส่ง Array JSON กลับไปเลย
}

func AddProduct(c *gin.Context) {
	// ⚠️ ส่วนนี้ยังรับเป็น FormData เหมือนเดิม เพราะต้องอัปโหลดไฟล์
	name := c.PostForm("product_name")
	category := c.PostForm("category")
	price, _ := strconv.ParseFloat(c.PostForm("price"), 64)
	qty, _ := strconv.Atoi(c.PostForm("quantity"))

	file, err := c.FormFile("product_image")
	imageName := "no_image.png"

	if err == nil {
		filename := filepath.Base(file.Filename)
		imageName = fmt.Sprintf("%d_%s", time.Now().Unix(), filename)
		// อย่าลืมสร้างโฟลเดอร์ uploads ไว้ที่ root project ด้วยนะครับ
		c.SaveUploadedFile(file, "uploads/"+imageName)
	}

	product := models.Product{
		ProductName:  name,
		Category:     category,
		Price:        price,
		Quantity:     qty,
		ProductImage: imageName,
		ProductCode:  fmt.Sprintf("P%d", time.Now().Unix()),
	}

	if err := config.DB.Create(&product).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, product)
}

func DeleteProduct(c *gin.Context) {
	id := c.Param("id")
	if err := config.DB.Delete(&models.Product{}, id).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "ลบไม่ได้"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "ลบสำเร็จ"})
}

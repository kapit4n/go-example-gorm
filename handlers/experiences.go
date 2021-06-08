package handlers

import (
	models "example/models"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

// Assign role to person
func AssignTech(c *gin.Context) {
	// db, _ = gorm.Open("mysql", "root:root@tcp(localhost:3306)/peco?charset=utf8&parseTime=True&loc=Local")
	db, err = gorm.Open("sqlite3", "./gorm.db")
	var experience models.Experience
	id := c.Params.ByName("id")
	// get experience by id
	if err := db.Where("id = ?", id).First(&experience).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	}

	db = db.Model(&experience).Preload("Technologies")
	var technology models.Technology

	c.BindJSON(&technology)

	//get technology by id
	techId := technology.ID
	if err := db.Where("id = ?", techId).First(&technology).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	}

	experience.Technologies = append(experience.Technologies, technology)

	db.Save(&experience)
	c.JSON(200, experience)
}

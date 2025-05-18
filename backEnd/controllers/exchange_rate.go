package controllers

import (
	"errors"
	"github.com/gin-gonic/gin"
	"go_code/ginStudy/gindemo/backEnd/global"
	"go_code/ginStudy/gindemo/backEnd/models"
	"gorm.io/gorm"
	"net/http"
	"time"
)

func CreateExchangeRate(c *gin.Context) {
	var exchangeRate models.ExchangeRate
	if err := c.ShouldBindJSON(&exchangeRate); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	exchangeRate.Date = time.Now()

	if err := global.Db.AutoMigrate(exchangeRate); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if err := global.Db.Create(&exchangeRate).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, exchangeRate)
}

func GetExchangeRates(c *gin.Context) {
	var exchangeRate []models.ExchangeRate
	if err := global.Db.Find(&exchangeRate).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
	}
	c.JSON(http.StatusOK, exchangeRate)
}

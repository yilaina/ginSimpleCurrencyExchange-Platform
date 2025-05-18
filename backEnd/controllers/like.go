package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
	"go_code/ginStudy/gindemo/backEnd/global"
	"net/http"
)

func LikeArticle(c *gin.Context) {
	articleID := c.Param("id")

	likeKey := "article:" + articleID + ":likes"

	if err := global.Rdb.Incr(likeKey).Err(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Successful liked the article"})
}

func GetArticleLikes(c *gin.Context) {
	articleID := c.Param("id")

	likeKey := "article:" + articleID + ":likes"

	likes, err := global.Rdb.Get(likeKey).Result()

	if err == redis.Nil {
		likes = "0"
	} else if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"likes": likes})
}

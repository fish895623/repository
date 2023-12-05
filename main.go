package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	. "viewmtechnology.com/dailyreport/reporter"
)

var db = ConnDataBase()

func main() {
	r := gin.Default()
	r.LoadHTMLGlob("templates/*.html")
	r.GET("/api/id", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{})
	})
	r.GET("/api/id/:id", func(c *gin.Context) {
	})
	r.POST("/api/id", func(c *gin.Context) {
		var user User

		if err := c.ShouldBindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		if err := db.Create(&User{Name: user.Name}); err.Error != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"status": "ok", "data": user})
	})
	r.GET("/", func(c *gin.Context) {
		c.HTML(
			http.StatusOK, "index.html",
			gin.H{})
	})
	r.Run(":8080")
}

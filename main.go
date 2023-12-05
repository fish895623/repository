package main

import (
	"net/http"

	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
	. "viewmtechnology.com/dailyreport/reporter"
)

var db = ConnDataBase()

func main() {
	r := gin.Default()
	r.GET("/api/id", func(c *gin.Context) {
		var ids []int64

		if err := db.Model(&User{}).Select("id").Pluck("id", &ids).Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"data": ids})
	}).POST("/api/id", func(c *gin.Context) {
		var user User

		if err := c.ShouldBindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		if err := db.Create(&User{Name: user.Name}); err.Error != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error.Error()})
			return
		}

		db.Last(&user)

		c.JSON(http.StatusOK, gin.H{"data": user})
	}).GET("/api/id/:id", func(c *gin.Context) {
		var user User

		id := c.Params.ByName("id")
		if err := db.Where("id = ?", id).First(&user).Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"data": user})
	})
	r.GET("/day", func(c *gin.Context) {
		c.Redirect(http.StatusTemporaryRedirect, "/day/2021/1/1")
	}).GET("/week", func(c *gin.Context) {
		c.Redirect(http.StatusTemporaryRedirect, "/week/2021/1/1")
	}).GET("/month", func(c *gin.Context) {
		c.Redirect(http.StatusTemporaryRedirect, "/month/2021/1/1")
	}).GET("/quarter", func(c *gin.Context) {
		c.Redirect(http.StatusTemporaryRedirect, "/quarter/2021/1/1")
	}).GET("/year", func(c *gin.Context) {
		c.Redirect(http.StatusTemporaryRedirect, "/year/2021/1/1")
	}).GET("/day/:year/:month/:day", func(c *gin.Context) {
		c.HTML(http.StatusOK, "day.html", gin.H{})
	}).GET("/week/:year/:month/:day", func(c *gin.Context) {
		year := c.Params.ByName("year")
		month := c.Params.ByName("month")
		day := c.Params.ByName("day")
		c.JSON(http.StatusOK, gin.H{"year": year, "month": month, "day": day})
	}).GET("/month/:year/:month/:day", func(c *gin.Context) {
		year := c.Params.ByName("year")
		month := c.Params.ByName("month")
		day := c.Params.ByName("day")
		c.JSON(http.StatusOK, gin.H{"year": year, "month": month, "day": day})
	}).GET("/quarter/:year/:month/:day", func(c *gin.Context) {
		year := c.Params.ByName("year")
		month := c.Params.ByName("month")
		day := c.Params.ByName("day")
		c.JSON(http.StatusOK, gin.H{"year": year, "month": month, "day": day})
	}).GET("/year/:year/:month/:day", func(c *gin.Context) {
		year := c.Params.ByName("year")
		month := c.Params.ByName("month")
		day := c.Params.ByName("day")
		c.JSON(http.StatusOK, gin.H{"year": year, "month": month, "day": day})
	})
	r.Use(static.Serve("/", static.LocalFile("./client/build", true)))

	r.Run(":8080")
}

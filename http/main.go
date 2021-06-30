package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.String(200, "Hello, Geektutu")
	})
	r.GET("/user/:name", func(c *gin.Context) {
		name := c.Param("name")
		role := c.Query("role")
		c.String(http.StatusOK, "Hello, %v=>%v", name, role)
	})
	r.POST("/users", func(c *gin.Context) {
		username := c.PostForm("username")
		age, _ := strconv.Atoi(c.PostForm("age"))
		c.JSON(http.StatusOK, gin.H{"username": username, "age": age})
	})
	r.POST("/users/", func(c *gin.Context) {
		user := &User{}
		if err := c.ShouldBindJSON(user); err != nil {
			c.JSON(http.StatusOK, gin.H{"error": err.Error()})
			return
		}
		username := user.UserName
		age := user.Age
		c.JSON(http.StatusOK, gin.H{"username": username, "age": age})
	})
	r.Run() // listen and serve on 0.0.0.0:8080
}

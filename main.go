package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go_api_server/coffee"
	"net/http"
)

func main() {
	r := gin.Default()
	r.LoadHTMLGlob("templates/*.html")
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "welcome to the coffee shop",
		})
	})
	r.GET("/home", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)

	})
	r.GET("/coffee", getCoffee)
	fmt.Println("Staring a web server")
	r.Run(":8081")
}

func getCoffee(c *gin.Context) {
	coffeeList, _ := coffee.GetCoffees()
	c.String(http.StatusOK, " %s", coffeeList)
}

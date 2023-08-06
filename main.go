package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go_api_server/coffee"
	"net/http"
	"os"
)

func main() {
	portNumber := os.Getenv("APP_PORT")
	r := gin.Default()
	r.LoadHTMLGlob("templates/*.html")
	r.GET("/ping", ping)
	r.GET("/home", coffeeList)
	r.GET("/coffee", getCoffee)
	fmt.Println("Staring a web server")
	fmt.Printf("Starting the app on port %s", portNumber)
	r.Run(fmt.Sprintf(":%s", portNumber))
}

func coffeeList(c *gin.Context) {
	coffees, _ := coffee.GetCoffees()
	c.HTML(
		http.StatusOK,
		"index.html",
		gin.H{
			"list": coffees.List,
		})
}
func getCoffee(c *gin.Context) {
	coffeelist, _ := coffee.GetCoffees()
	c.String(http.StatusOK, " %s", coffeelist)
}

func ping(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "welcome to the coffee-shop",
	})
}

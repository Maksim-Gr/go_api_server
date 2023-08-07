package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"go_api_server/coffee"
	"log"
	"net/http"
	"os"
)

var path = "logs/logs.txt"

func createFile() {
	var _, err = os.Stat(path)
	if os.IsNotExist(err) {
		err = os.MkdirAll("logs", os.ModePerm)
		if err != nil {
			return
		}
		var file, err = os.Create(path)
		if err != nil {
			return
		}
		defer file.Close()
	}
}

func NewLogger() (*zap.Logger, error) {
	cfg := zap.NewProductionConfig()
	cfg.OutputPaths = []string{
		path,
	}
	return cfg.Build()
}

var sugar *zap.SugaredLogger

func init() {
	createFile()
	logger, err := NewLogger()
	if err != nil {
		log.Fatal(err)
	} else {
		sugar = logger.Sugar()
	}
}

func main() {
	portNumber := os.Getenv("APP_PORT")
	r := gin.Default()
	r.LoadHTMLGlob("templates/*.html")
	r.GET("/ping", ping)
	r.GET("/home", coffeeList)
	r.GET("/coffee", getCoffee)
	sugar.Infof("Starting the app on port %s", portNumber)
	r.Run(fmt.Sprintf(":%s", portNumber))
}

func coffeeList(c *gin.Context) {
	coffees, err := coffee.GetCoffees()
	if err != nil {
		sugar.Errorf("Error while getting the coffee list %s", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Error!",
		})
		return
	}
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

func setLogOutput() {
	file, err := os.OpenFile("logs.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}
	log.SetOutput(file)
}

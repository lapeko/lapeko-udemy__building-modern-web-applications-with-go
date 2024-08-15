package main

import (
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	router := gin.Default()
	router.LoadHTMLGlob("templates/*")
	router.GET("/", func(c *gin.Context) {
		log.Println("Visit home page")
		c.HTML(200, "home.page.gohtml", gin.H{
			"Title":   "Home page",
			"Message": "Home page's message",
		})
	})
	_ = router.Run(":3031")
}

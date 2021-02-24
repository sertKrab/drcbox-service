package main

import (
	"fmt"
	"os"

	"gorm.io/gorm"
)

func main() {
	port := os.Getenv("PORT")
	fmt.Print(port)
	// if port == "" {
	// 	port = "8080"
	// }

	// router := gin.New()
	// router.Use(gin.Logger())

	// router.Run(fmt.Sprintf(":%s", port))

	// router.GET("/", func(c *gin.Context) {
	// 	c.JSON(200, gin.H{
	// 		"message": "pong",
	// 	})
	// })

}

// User Model
type User struct {
	gorm.Model
	Name string
}

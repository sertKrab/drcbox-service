package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

func main() {
	viper.SetDefault("port", ":1323")
	viper.AutomaticEnv()
	r := gin.New()
	r.Use(gin.Recovery())
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	port := os.Getenv("PORT")
	srv := &http.Server{
		Addr:    port,
		Handler: r,
	}

	go func() {
		log.Println("Starting server at", port)
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	// graceful shutdown
	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server forced to shutdown:", err)
	}

	log.Println("Server exiting")

	// viper.SetDefault("dsn", "sqlserver://sert11:1234567890@localhost:1433?database=social1")

	// dsn := viper.GetString("dsn")

	// db, err := gorm.Open(mysql.New(mysql.Config{
	// 	DSN:                       dsn,   // data source name
	// 	DefaultStringSize:         256,   // default size for string fields
	// 	DisableDatetimePrecision:  true,  // disable datetime precision, which not supported before MySQL 5.6
	// 	DontSupportRenameIndex:    true,  // drop & create when rename index, rename index not supported before MySQL 5.7, MariaDB
	// 	DontSupportRenameColumn:   true,  // `change` when rename column, rename column not supported before MySQL 8, MariaDB
	// 	SkipInitializeWithVersion: false, // auto configure based on currently MySQL version
	// }), &gorm.Config{})
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// if err := db.AutoMigrate(User{}); err != nil {
	// 	log.Fatal(err)
	// }

}

// User Model
type User struct {
	gorm.Model
	Name string
}

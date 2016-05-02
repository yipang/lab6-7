package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

var (
	db *sql.DB
)

func main() {

	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("$PORT must be set")
	}

	var errd error
	db, errd = sql.Open("postgres", os.Getenv("DATABASE_URL"))
	if errd != nil {
		log.Fatalf("Error opening database: %q", errd)
	}

	router := gin.New()
	router.Use(gin.Logger())
	router.LoadHTMLGlob("html/*.html")
	router.Static("/static", "static")

	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})

	router.GET("/db", func(c *gin.Context) {
		if db != nil {
			c.JSON(http.StatusOK, gin.H{"message": "db created"})
		} else {
			c.JSON(http.StatusOK, gin.H{"message": "db was not created. Check your DATABASE_URL"})
		}
	})

	router.Run(":" + port)
}

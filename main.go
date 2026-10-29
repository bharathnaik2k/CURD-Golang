package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
)

// DB ಅನ್ನೋದು ನಮ್ಮ ಡೇಟಾಬೇಸ್ ಕನೆಕ್ಷನ್ (ಗ್ಲೋಬಲ್ ವೇರಿಯಬಲ್)
var DB *pgxpool.Pool
var version = "v1"

func main() {
	// 1. .env ಫೈಲ್ ಅನ್ನು ಲೋಡ್ ಮಾಡೋಣ
	if err := godotenv.Load(); err != nil {
		log.Println("Warning: No .env file found or error loading it")
	}

	// 2. ಡೇಟಾಬೇಸ್ ಗೆ ಕನೆಕ್ಟ್ ಮಾಡೋಣ
	dbURL := os.Getenv("DB_URL")
	var err error
	DB, err = pgxpool.New(context.Background(), dbURL)
	if err != nil {
		log.Fatalf("❌ Unable to connect to database: %v\n", err)
	}
	defer DB.Close() // ಪ್ರೋಗ್ರಾಮ್ ಸ್ಟಾಪ್ ಆದಾಗ ಕನೆಕ್ಷನ್ ಕ್ಲೋಸ್ ಮಾಡೋಕೆ

	// ಡೇಟಾಬೇಸ್ ಕನೆಕ್ಟ್ ಆಗಿದೆಯಾ ಅಂತ ಚೆಕ್ ಮಾಡೋಕೆ Ping ಮಾಡೋಣ
	if err := DB.Ping(context.Background()); err != nil {
		log.Fatalf("❌ Database ping failed: %v\n", err)
	}
	fmt.Println("✅ Successfully connected to PostgreSQL Database!")

	// 3. Router ಸೆಟಪ್
	r := gin.Default()

	// /v1/health ಅಂತ ಒಂದು ಸಿಂಪಲ್ GET ರಿಕ್ವೆಸ್ಟ್ ಹ್ಯಾಂಡಲ್ ಮಾಡೋಣ
	r.GET("/"+version+"/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status":  true,
			"message": "success",
		})
	})

	r.GET("/"+version+"/check", func(c *gin.Context) {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  true,
			"message": "pass",
		})
	})

	// ಸರ್ವರ್ ರನ್ ಮಾಡೋಣ
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	fmt.Println("Server is running on port: " + port)
	r.Run(":" + port)
}

// package main

// import "fmt"

// func main(){
// 	fmt.Println("hello world")
// }

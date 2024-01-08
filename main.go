package main

import (
	"fmt"
	"log"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/teerapoom/School_MiniApi/Server/controller"
	"github.com/teerapoom/School_MiniApi/Server/database"
	"github.com/teerapoom/School_MiniApi/Server/middleware"
	"github.com/teerapoom/School_MiniApi/Server/model"
)

func main() {
	loadEnv()
	loaddatebase()
	my_server()
}

func loadEnv() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error Load .env file")
	}
	log.Printf("Load .env Successfully")
}

func loaddatebase() {
	database.InitDb()
	database.Db.AutoMigrate(&model.School{})
}

// func CORSMiddleware() gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		c.Header("Content-Type", "application/json")
// 		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
// 		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
// 		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
// 		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

// 		if c.Request.Method == "OPTIONS" {
// 			c.AbortWithStatus(204)
// 			return
// 		}

// 		c.Next()
// 	}
// }

func my_server() {
	gin.SetMode(gin.ReleaseMode)
	router := gin.New()

	// CORS
	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	config.AllowMethods = []string{"POST", "GET", "PUT", "PATCH"}
	config.AllowHeaders = []string{"Origin", "Content-Type", "Authorization", "Accept", "User-Agent", "Cache-Control", "Pragma"}
	config.ExposeHeaders = []string{"Content-Length"}
	config.AllowCredentials = true
	config.MaxAge = 12 * time.Hour
	router.Use(cors.New(config))

	// Middleware
	router.Use(middleware.CheckMethod)

	//Process
	router.POST("/AddStudent", controller.CreateStudent)
	router.GET("/ViewAll", controller.ViewAll)
	router.GET("/View/:id", controller.ViewById)
	router.PATCH("/EditStudent/:id", controller.UpdateStuder)
	router.DELETE("/RemoveStudent/:id", controller.RemoveStuder)

	//Port
	fmt.Println("Server Run....")
	router.Run(":8080")
}

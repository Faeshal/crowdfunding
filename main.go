package main

import (
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"crowdfunding/auth"
	"crowdfunding/campaign"
	"crowdfunding/handler"
	"crowdfunding/helper"
	"crowdfunding/payment"
	"crowdfunding/transaction"
	"crowdfunding/user"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	// * dynamic env path finder
	// const projectDirName = "crowdfunding"
	// projectName := regexp.MustCompile(`^(.*` + projectDirName + `)`)
	// currentWorkDirectory, _ := os.Getwd()
	// rootPath := projectName.Find([]byte(currentWorkDirectory))

	// * load .env file
	// err := godotenv.Load(string(rootPath) + `/.env`)
	err = godotenv.Load(filepath.Join("/var/www/crowdfunding", ".env"))
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	DB_USERNAME := os.Getenv("DB_USERNAME")
	DB_PASSWORD := os.Getenv("DB_PASSWORD")
	DB_NAME := os.Getenv("DB_NAME")

	// dsn := "root:root@tcp(127.0.0.1:3306)/crowdfunding?charset=utf8mb4&parseTime=True&loc=Local"
	dsn := DB_USERNAME + ":" + DB_PASSWORD + "@tcp(127.0.0.1:3306)/" + DB_NAME + "?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err.Error())
	}
	log.Print("DB Connected 🚀...")

	// * Declare Global Dependency
	// repository
	userRepository := user.NewRepository(db)
	campaignRepository := campaign.NewRepository(db)
	transactionRepository := transaction.NewRepository(db)

	// service
	userService := user.NewService(userRepository)
	authService := auth.NewService()
	paymentService := payment.NewService()
	campaignService := campaign.NewService(campaignRepository)
	transactionService := transaction.NewService(transactionRepository, campaignRepository, paymentService)

	// handler
	userHandler := handler.NewUserHandler(userService, authService)
	campaignHandler := handler.NewCampaignHandler(campaignService)
	transactionHandler := handler.NewTransactionHandler(transactionService)

	// * server init
	router := gin.Default()

	// * cors
	router.Use(cors.Default())

	// * static file
	router.Static("/images", "./images")

	// * route
	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status":  200,
			"message": "Welcome to Crowdfunding Service build on top of GIN 🍹",
		})
	})

	api := router.Group("/api/v1")

	api.POST("/users", userHandler.RegisterUser)
	api.POST("/sessions", userHandler.Login)
	api.POST("/email_checkers", userHandler.CheckEmailAvailability)
	api.POST("/avatars", authMiddleware(authService, userService), userHandler.UploadAvatar)
	api.GET("/users/fetch", authMiddleware(authService, userService), userHandler.FetchUser)

	api.GET("/campaigns", campaignHandler.GetCampaigns)
	api.GET("/campaigns/:id", campaignHandler.GetCampaign)
	api.POST("/campaigns", authMiddleware(authService, userService), campaignHandler.CreateCampaign)
	api.PUT("/campaigns/:id", authMiddleware(authService, userService), campaignHandler.UpdateCampaign)
	api.POST("/campaign-images", authMiddleware(authService, userService), campaignHandler.UploadImage)

	api.GET("/campaigns/:id/transactions", authMiddleware(authService, userService), transactionHandler.GetCampaignTransactions)
	api.GET("/transactions", authMiddleware(authService, userService), transactionHandler.GetUserTransactions)
	api.POST("/transactions", authMiddleware(authService, userService), transactionHandler.CreateTransaction)
	api.POST("/transactions/notification", transactionHandler.GetNotification)

	// * Run Gin
	router.Run(":7070")
}

func authMiddleware(authService auth.Service, userService user.Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")

		if !strings.Contains(authHeader, "Bearer") {
			response := helper.APIResponse("Unauthorized", http.StatusUnauthorized, "error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		tokenString := ""
		arrayToken := strings.Split(authHeader, " ")
		if len(arrayToken) == 2 {
			tokenString = arrayToken[1]
		}

		token, err := authService.ValidateToken(tokenString)
		if err != nil {
			response := helper.APIResponse("Unauthorized", http.StatusUnauthorized, "error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		claim, ok := token.Claims.(jwt.MapClaims)

		if !ok || !token.Valid {
			response := helper.APIResponse("Unauthorized", http.StatusUnauthorized, "error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		// karena id kita typenya integer , maka perlu di ubah dari string, float ke int lagi
		userID := int(claim["user_id"].(float64))

		user, err := userService.GetUserByID(userID)
		if err != nil {
			response := helper.APIResponse("Unauthorized", http.StatusUnauthorized, "error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		c.Set("currentUser", user)
	}
}

package main

import (
	"database/sql"
	"final-project-sanbercode/controllers"
	"final-project-sanbercode/database"
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

var (
	DB  *sql.DB
	err error
)

func main() {
	err = godotenv.Load("config/.env")
	if err != nil {
		panic("Error loading .env file")
	}

	psqlInfo := fmt.Sprintf(`host=%s port=%s user=%s password=%s dbname=%s sslmode=disable`,
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
	)

	DB, err = sql.Open("postgres", psqlInfo)
	defer DB.Close()
	err = DB.Ping()
	if err != nil {
		panic(err)
	}

	database.DBMigrate(DB)

	router := gin.Default()
	router.POST("/user/login", controllers.Login)
	router.POST("/user/register", controllers.Register)
	// Group routes that require basic auth
	authorized := router.Group("/api/", gin.BasicAuth(gin.Accounts{
		"admin": "admin",
	}))

	authorized.GET("rooms", controllers.GetAllRoom)
	authorized.POST("rooms", controllers.InsertRoom)
	authorized.GET("rooms/:id", controllers.GetRoomByID)
	authorized.PUT("rooms/:id", controllers.UpdateRoom)
	authorized.DELETE("rooms/:id", controllers.DeleteRoom)

	authorized.GET("bookings", controllers.GetAllBooking)
	authorized.POST("bookings", controllers.InsertBooking)
	authorized.GET("bookings/:id", controllers.GetBookingByID)
	authorized.PUT("bookings/:id", controllers.UpdateBooking)
	authorized.DELETE("bookings/:id", controllers.DeleteBooking)

	authorized.GET("payments", controllers.GetAllPayment)
	authorized.GET("payments/:id", controllers.GetPaymentByID)
	authorized.POST("payments", controllers.InsertPayment)
	authorized.PUT("payments/:id", controllers.UpdatePayment)
	authorized.DELETE("payments/:id", controllers.DeletePayment)

	router.Run(":8080")
}

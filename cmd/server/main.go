package main

import (
	"log"
	"os"

	"database/sql"
	_ "github.com/go-sql-driver/mysql"

	"github.com/gin-gonic/gin"

	"github.com/desafioFinalBack/cmd/server/handler"
	"github.com/desafioFinalBack/internal/dentist"
	"github.com/desafioFinalBack/pkg/middleware"
	_ "github.com/swaggo/swag/cmd/swag"
	store "github.com/desafioFinalBack/pkg/storeDentists"

	"github.com/joho/godotenv"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func main() {
	db, err := sql.Open("mysql", "root:Felisantiago87@tcp(localhost:3306)/desafioFinalBackDB")
	if err != nil {
		log.Fatal(err)
	}

	if err := godotenv.Load("./cmd/server/.env"); err != nil {
		panic("Error loading .env file: " + err.Error())
	}

	storage := store.NewSqlStore(db)
	repo := dentist.NewRepository(storage)
	service := dentist.NewService(repo)
	dentistHandler := handler.NewProductHandler(service)

	r := gin.Default()

	docs.SwaggerInfo.Host = os.Getenv("HOST")
	r.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	dentists := r.Group("/dentists")
	{
		dentists.GET(":id", dentistHandler.GetByID())
		dentists.POST("", middleware.Authentication(), dentistHandler.Post())
		dentists.PUT(":id", middleware.Authentication(), dentistHandler.Put())
		dentists.PATCH(":id", middleware.Authentication(), dentistHandler.Patch())
		dentists.DELETE(":id", middleware.Authentication(), dentistHandler.Delete())
	}
	r.Run(":8080")
}
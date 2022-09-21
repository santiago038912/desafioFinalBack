package main

import (
	"log"

	"database/sql"
	_ "github.com/go-sql-driver/mysql"

	"github.com/gin-gonic/gin"

	"github.com/desafioFinalBack/cmd/server/handler"
	"github.com/desafioFinalBack/internal/dentist"
	"github.com/desafioFinalBack/pkg/middleware"
	store "github.com/desafioFinalBack/pkg/store"

	"github.com/joho/godotenv"
)

func main() {
	db, err := sql.Open("mysql", "root:Felisantiago87@tcp(localhost:3306)/desafioFinalBackDB")
	if err != nil {
		log.Fatal(err)
	}

	if err := godotenv.Load("./cmd/server/.env"); err != nil {
		panic("Error loading .env file: " + err.Error())
	}

	storageDentist := store.NewSqlStoreDentist(db)
	repoDentist := dentist.NewRepository(storageDentist)
	serviceDentist := dentist.NewService(repoDentist)
	dentistHandler := handler.NewProductHandler(serviceDentist)

	r := gin.Default()

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
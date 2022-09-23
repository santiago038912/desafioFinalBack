package main

import (
	"log"

	"database/sql"
	_ "github.com/go-sql-driver/mysql"

	"github.com/gin-gonic/gin"

	"github.com/desafioFinalBack/cmd/server/handler"
	"github.com/desafioFinalBack/internal/dentist"
	"github.com/desafioFinalBack/internal/patient"
	"github.com/desafioFinalBack/internal/turn"
	"github.com/desafioFinalBack/pkg/middleware"
	"github.com/desafioFinalBack/pkg/store"

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

	storagePatient := store.NewSqlStorePatient(db)
	repoPatient := patient.NewRepository(storagePatient)
	servicePatient := patient.NewService(repoPatient)
	patientHandler := handler.NewPatientHandler(servicePatient)

	storageTurn := store.NewSqlStoreTurn(db)
	repoTurn := turn.NewRepository(storageTurn)
	serviceTurn := turn.NewService(repoTurn)
	turnHandler := handler.NewTurnHandler(serviceTurn)

	r := gin.Default()

	dentists := r.Group("/dentists")
	{
		dentists.GET(":id", dentistHandler.GetDentistByID())
		dentists.POST("", middleware.Authentication(), dentistHandler.PostDentist())
		dentists.PUT(":id", middleware.Authentication(), dentistHandler.PutDentist())
		dentists.PATCH(":id", middleware.Authentication(), dentistHandler.PatchDentist())
		dentists.DELETE(":id", middleware.Authentication(), dentistHandler.DeleteDentist())
	}

	patients := r.Group("/patients")
	{
		patients.GET(":id", patientHandler.GetPatientByID())
		patients.POST("", middleware.Authentication(), patientHandler.PostPatient())
		patients.PUT(":id", middleware.Authentication(), patientHandler.PutPatient())
		patients.PATCH(":id", middleware.Authentication(), patientHandler.PatchPatient())
		patients.DELETE(":id", middleware.Authentication(), patientHandler.DeletePatient())
	}

	turns := r.Group("/turns")
	{
		turns.GET(":id", turnHandler.GetTurnByID())
		turns.GET(":dni", turnHandler.GetTurnByDNI())
		turns.POST("", middleware.Authentication(), turnHandler.PostTurn())
		turns.PUT(":id", middleware.Authentication(), turnHandler.PutTurn())
		turns.PATCH(":id", middleware.Authentication(), turnHandler.PatchTurn())
		turns.DELETE(":id", middleware.Authentication(), turnHandler.DeleteTurn())
	}

	r.Run(":8080")
}

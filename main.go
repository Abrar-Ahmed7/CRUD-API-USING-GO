package main

import (
	"errors"
	"log"
	"net/http"

	"github.com/Abrar-Ahmed7/rest-api-go/internal/db"
	"github.com/Abrar-Ahmed7/rest-api-go/internal/handler"
	usergormimpl "github.com/Abrar-Ahmed7/rest-api-go/internal/repo/user/gormimpl"
	"github.com/Abrar-Ahmed7/rest-api-go/internal/service"
	"github.com/gin-gonic/gin"

	"github.com/golang-migrate/migrate/v4"
	m "github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func main() {

	engine := gin.Default()
	router := engine.Group("/api/v1")

	dbConn, err := db.NewDBConn()
	if err != nil {
		panic(err)
	}

	sqlDB, _ := dbConn.DB.DB()
	driver, err := m.WithInstance(sqlDB, &m.Config{})
	if err != nil {
		panic(err)
	}

	mig, err := migrate.NewWithDatabaseInstance(
		"file://db/migration",
		"mysql",
		driver,
	)
	if err != nil {
		panic(err)
	}

	err = mig.Up()
	if err != nil && err != migrate.ErrNoChange {
		panic(err)
	}
	log.Println("Migration Successful or Already Migrated")

	ur := usergormimpl.NewGormUserRepo(dbConn)
	us := service.NewUserService(ur)
	uh := handler.NewUserHandler(us)
	uh.RegisterRoutes(router)

	srv := &http.Server{
		Addr:    ":8080",
		Handler: engine,
	}

	if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
		log.Printf("listen: %s\n", err)
	}
}

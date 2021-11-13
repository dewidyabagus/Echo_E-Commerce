package main

import (
	"fmt"

	echo "github.com/labstack/echo/v4"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	// Module configuration file
	"RESTful/config"

	// Migration
	"RESTful/modules/migration"

	// API
	"RESTful/api"

	// User
	userController "RESTful/api/v1/user"
	userService "RESTful/business/user"
	userRepository "RESTful/modules/user"

	// Auth
	authController "RESTful/api/v1/auth"
	authService "RESTful/business/auth"
)

func newDatabaseConnection(config *config.AppConfig) *gorm.DB {
	strConnection := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%d",
		config.PgHost, config.PgUsername, config.PgPassword, config.PgDbname, config.PgPort,
	)

	db, err := gorm.Open(postgres.Open(strConnection), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	migration.TableMigration(db)

	return db
}

func main() {
	// Get configuration file
	config := config.GetAppConfig()

	// Create new session database
	dbConnection := newDatabaseConnection(config)

	// Initiate user repository
	userRepo := userRepository.NewRepository(dbConnection)

	// Initiate user service
	userSvc := userService.NewService(userRepo)

	// Initiate user controller
	userCtr := userController.NewController(userSvc)

	// Initiate auth user
	authSvc := authService.NewService(userSvc)

	// Initiate auth controller
	authCtr := authController.NewController(authSvc)

	// Initiate echo web framework
	e := echo.New()

	// Initiate routes
	api.RegisterRouters(e, userCtr, authCtr)

	// start echo
	e.Start(":8000")
}

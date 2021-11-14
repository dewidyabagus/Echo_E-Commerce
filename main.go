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

	// Admin
	adminController "RESTful/api/v1/admin"
	adminService "RESTful/business/admin"
	adminRepository "RESTful/modules/admin"

	// Auth
	authController "RESTful/api/v1/auth"
	authService "RESTful/business/auth"

	// Product
	productController "RESTful/api/v1/product"
	productService "RESTful/business/product"
	productRepository "RESTful/modules/product"
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

	// ========= ADMIN ========= //
	// Initiate admin repository
	adminRepo := adminRepository.NewRepository(dbConnection)
	// Initiate admin service
	adminSvc := adminService.NewService(adminRepo)
	// Initiate admin controller
	adminCtr := adminController.NewController(adminSvc)

	// ========= USER ========= //
	// Initiate user repository
	userRepo := userRepository.NewRepository(dbConnection)
	// Initiate user service
	userSvc := userService.NewService(userRepo, adminSvc)
	// Initiate user controller
	userCtr := userController.NewController(userSvc)

	// ========= AIM ========= //
	// Initiate auth user
	authSvc := authService.NewService(userSvc, adminSvc)
	// Initiate auth controller
	authCtr := authController.NewController(authSvc)

	// ========= PRODUCT ========= //
	// Initiate product repository
	productRepo := productRepository.NewRepository(dbConnection)
	// Initiate product service
	productSvc := productService.NewService(productRepo, userSvc)
	// Initiate product controller
	producCtr := productController.NewController(productSvc)

	// Initiate echo web framework
	e := echo.New()

	// Initiate routes && userCtr, authCtr, adminCtr
	api.RegisterRouters(e, &api.Routing{
		User:    userCtr,
		Auth:    authCtr,
		Admin:   adminCtr,
		Product: producCtr,
	})

	// start echo
	e.Start(":8000")
}

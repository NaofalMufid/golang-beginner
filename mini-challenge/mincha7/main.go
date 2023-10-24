package main

import (
	"mincha7/config"
	"mincha7/controller"
	"mincha7/helper"
	"mincha7/model"
	"mincha7/repository"
	"mincha7/router"
	"mincha7/service"
	"net/http"
	"os"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	helper.ErrorPanic(err)

	// Database
	db := config.DBInit()
	validate := validator.New()

	db.Table("orders").AutoMigrate(&model.Orders{})
	db.Table("items").AutoMigrate(&model.Items{})

	// Init Repository
	orderRepository := repository.NewOrdersRepository(db)
	itemRepository := repository.NewItemRepository(db)
	itemService := service.NewItemServiceImpl(itemRepository)

	// Init Service
	orderService := service.NewOrderServiceImpl(orderRepository, itemService, validate)

	// Init Controller
	orderController := controller.NewOrderController(orderService)

	// Router
	routes := router.NewRouter(orderController)

	port := os.Getenv("API_PORT")
	server := &http.Server{
		Addr:           port,
		Handler:        routes,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	err = server.ListenAndServe()
	helper.ErrorPanic(err)
}

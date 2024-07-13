package main

import (
	"fmt"
	"log"
	"onlineShop/apps/auth"
	"onlineShop/apps/products"
	"onlineShop/external/database"
	"onlineShop/internal/config"
	"os"
	"path/filepath"

	"github.com/gofiber/fiber/v2"
)

// Custom Panic function to log and panic
func Panic(v ...any) {
	s := fmt.Sprint(v...)
	log.Output(2, s) // Assuming log is the standard logger or your custom logger
	panic(s)
}

func main() {
	filename := "cmd/api/config.yaml"

	// Get the absolute path to the configuration file
	absPath, err := filepath.Abs(filename)
	if err != nil {
		Panic("Error getting absolute path:", err)
	}
	fmt.Println("Absolute path to config file:", absPath)

	// Ensure the configuration file exists
	if _, err := os.Stat(absPath); os.IsNotExist(err) {
		Panic("Configuration file", absPath, "does not exist")
	}

	// Load the configuration
	err = config.LoadConfig(absPath)
	if err != nil {
		Panic(err)
	}

	// Connect to the database
	db, err := database.ConnectPostgres(config.Cfg.DB)
	if err != nil {
		Panic(err)
	}

	if db != nil {
		log.Println("Connected to database")
	}

	// Set up the Fiber app
	router := fiber.New(fiber.Config{
		Prefork: true,
		AppName: config.Cfg.App.Name,
	})

	// Initialize authentication routes
	auth.Init(router, db)
	products.Init(router, db)

	// Start the server
	if err := router.Listen(config.Cfg.App.Port); err != nil {
		Panic(err)
	}
}

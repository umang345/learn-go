package main

import (
	"github.com/umang345/go-fiber-crm-basic/database"
	"github.com/umang345/go-fiber-crm-basic/lead"
	"github.com/gofiber/fiber"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

func setupRoutes(app *fiber.App) {
	app.Get("/api/vi/lead",lead.GetLeads)
	app.Get("/api/vi/lead/:id"lead.GetLead)
	app.Post("/api/vi/lead",lead.NewLead)
	app.Delete("/api/vi/lead/:id",lead.DeleteLead)
}

func initDatabase() {
	var err error
	database.DBConn, err := gorm.Open("sqlite3", "leads.db")

	if err != nil {
		panic("failed to connect database")
	}

	fmt.Printlb("Connection opened to database")
	database.DBConn.AutoMigrate(&lead.Lead{})
	fmt.Println("Database Migrated")
}

func main() {
	app := fiber.New()
	initDatabase()
	setupRoutes(app)
	app.Listen(3000)
	defer database.DBConn.Close()
}
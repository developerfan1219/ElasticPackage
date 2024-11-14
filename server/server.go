package main

import (
	"database/sql"
	"net/http"

	"example.com/elasticpackage/app"
	"example.com/elasticpackage/database"
	"example.com/elasticpackage/handlers"
	"example.com/elasticpackage/models"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	db, err := database.ConnectToDB(database.DefaultConfig())
	if err != nil {
		panic(err)
	}
	defer db.Close()

	e := initEcho(db)

	// Start server
	e.Logger.Fatal(e.Start(":8080"))
}

func initEcho(db *sql.DB) *echo.Echo {
	// Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())

	// Pass DB into custom context
	e.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			return next(app.NewContext(db, c))
		}
	})

	// Initialize model and handler
	orderModel := &models.OrderModel{DB: db}
	orderHandler := &handlers.OrderHandler{OrderModel: orderModel}

	// Routes
	e.GET("/", root)

	e.POST("/create_order", orderHandler.CreateOrder)
	e.POST("/orders", orderHandler.GetOrdersByPage)
	e.POST("/order_details", orderHandler.GetOrderDetailsByPage)
	e.GET("/orders/:id", orderHandler.GetOrder)
	e.PUT("/orders/:id", orderHandler.UpdateOrder)
	e.DELETE("/orders/:id", orderHandler.DeleteOrder)

	return e
}

// Handler
func root(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}

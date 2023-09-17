package main

import (
	"fmt"
	"os"
	"time"

	"github.com/barealek/minilink/internal/controllers"
	"github.com/barealek/minilink/internal/database"
	"github.com/barealek/minilink/pkg/shutdown"

	gojson "github.com/goccy/go-json"
	"github.com/gofiber/fiber/v2"
)

func main() {
	// Prepare the exit code and defer the exit
	var exitCode int

	defer func() {
		os.Exit(exitCode)
	}()

	// Run the application
	cleanup, err := run()
	defer cleanup()

	// Handle any startup error
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		exitCode = 1
		return
	}

	// Wait for the application to exit and exit gracefully
	shutdown.Shutdown()
}

func run() (func(), error) {
	app, cleanup, err := buildMiniLink()
	if err != nil {
		return nil, err
	}

	// Run the app in a goroutine
	go func() {
		app.Listen(":3000")
	}()

	return func() {
		app.Shutdown()
		cleanup()
	}, nil
}

func buildMiniLink() (*fiber.App, func(), error) {

	// Connect to the MongoDB database
	db, err := database.BootstrapMongo(os.Getenv("MONGO_CONN_URI"), "miniLink", 10*time.Second)
	if err != nil {
		return nil, nil, err
	}

	app := fiber.New(fiber.Config{
		JSONEncoder: gojson.Marshal,
		JSONDecoder: gojson.Unmarshal,
	})

	controllers.RegisterGroups(app)

	return app, func() {
		database.CloseMongo(db)
	}, nil
}

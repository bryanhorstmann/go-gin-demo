// main.go

package main

import (
	"log"
	"net/http"
	"os"

	"github.com/getsentry/sentry-go"
	sentrygin "github.com/getsentry/sentry-go/gin"
	"github.com/gin-gonic/gin"
)

var router *gin.Engine

func main() {
	// Initialise Sentry
	if err := sentry.Init(sentry.ClientOptions{
		Dsn:              os.Getenv("SENTRY_DSN"),
		Debug:            true,
		AttachStacktrace: true,
	}); err != nil {
		log.Printf("Sentry initialization failed: %v\n", err)
	}

	// Set the router as the default one provided by Gin
	router = gin.Default()

	// Load middleware
	router.Use(sentrygin.New(sentrygin.Options{
		Repanic: true,
	}))
	router.Use(sentryHandler())
	router.Use(errorHandler())

	// Process the templates at the start so that they don't have to be loaded
	// from the disk again. This makes serving HTML pages very fast.
	router.LoadHTMLGlob("../templates/*")

	// Initialize the routes
	initializeRoutes()

	// Start serving the application
	router.Run()

}

// Render one of HTML, JSON or CSV based on the 'Accept' header of the request
// If the header doesn't specify this, HTML is rendered, provided that
// the template name is present
func render(c *gin.Context, data gin.H, templateName string) {

	switch c.Request.Header.Get("Accept") {
	case "application/json":
		// Respond with JSON
		c.JSON(http.StatusOK, data["payload"])
	case "application/xml":
		// Respond with XML
		c.XML(http.StatusOK, data["payload"])
	default:
		// Respond with HTML
		c.HTML(http.StatusOK, templateName, data)
	}

}

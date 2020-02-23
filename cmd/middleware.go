package main

import (
	"fmt"
	"log"
	"time"

	"github.com/getsentry/sentry-go"
	sentrygin "github.com/getsentry/sentry-go/gin"
	"github.com/gin-gonic/gin"
)

func errorHandler() gin.HandlerFunc {
	log.Println("this runs only once")

	return func(c *gin.Context) {
		// before request
		t := time.Now()

		// Set example variable
		c.Set("example", "12345")

		c.Next()

		// after request
		latency := time.Since(t)
		log.Print(latency)

		// access the status we are sending
		status := c.Writer.Status()
		log.Println(status)
	}
}

func sentryHandler() gin.HandlerFunc {
	log.Println("Sentry middleware")

	return func(c *gin.Context) {
		// Everything runs after request
		c.Next()

		if c.Writer.Status() >= 400 {
			if hub := sentrygin.GetHubFromContext(c); hub != nil {
				log.Println("there is a hub")
				hub.WithScope(func(scope *sentry.Scope) {
					hub.CaptureException(fmt.Errorf("this is just a random error"))
				})
			}
			log.Println("There was an error")
		}
	}
}

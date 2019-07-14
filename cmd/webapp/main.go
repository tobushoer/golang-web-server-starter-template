// @Version 0.0.1
// @Title Backend API
// @Description Web API server starter
// @ContactName tobushoer
// @ContactEmail tobushoer@gmail.com
// @ContactURL http://example.example
// @TermsOfServiceUrl http://example.example
// @LicenseName MIT
// @LicenseURL https://en.wikipedia.org/wiki/MIT_License
// @Server http://www.fake.com Server-1

package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"template/internal/pkg/config"
	"template/internal/pkg/database/mysql"
	"template/internal/pkg/log"
	"time"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	// Read configs
	cfg := config.New("configs/config-dev.yaml")

	// Create logger
	log.New(cfg.Log.Pkg)
	defer log.Close(cfg.Log.Pkg)

	//  Create db  connection
	db := mysql.New(cfg)
	defer db.Close()

	// Set http router
	router := gin.Default()
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	registryHTTPHandler(router)

	httpServer := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	go func() {
		if err := httpServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Errorf("HTTP server ListenAndServe: %v", err)
		}
		syscall.Kill(syscall.Getpid(), syscall.SIGINT)
	}()

	// Wait for interrupt signal to gracefully shutdown the server with
	// a timeout of 10 seconds.
	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Info("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := httpServer.Shutdown(ctx); err != nil {
		log.Warnf("Server Shutdown: %v", err)
	}

	log.Info("Server exiting")
}

package main

import (
	"context"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/russross/blackfriday/v2"
)

func usage() {
	fmt.Println("Usage: ./preview -file=<markdown-file> [options]")
	fmt.Println("Options: -port=<port-number> (default: 6060)")
	os.Exit(1)
}

func main() {
	// Parse the command line flags
	filePath := flag.String("file", "", "markdown file path")
	port := flag.String("port", "6060", "port to listen on")
	flag.Parse()

	if *filePath == "" {
		usage()
	}

	r := gin.Default()

	r.GET("/preview", func(c *gin.Context) {
		// Read the Markdown file from disk
		mdBytes, err := ioutil.ReadFile(*filePath)
		if err != nil {
			fmt.Println("Error reading file!!")
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		// Convert the Markdown to HTML using Blackfriday
		html := blackfriday.Run(mdBytes)

		// previewHandler(c)
		c.Header("Content-Type", "text/html")
		c.Status(http.StatusOK)
		c.Writer.Write(html)
	})

	server := &http.Server{
		Addr:    ":" + *port,
		Handler: r,
	}

	// Start the server in a separate goroutine
	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	// Wait for interrupt signal to gracefully shut down the server
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	log.Println("Shutting down server...")

	// Shut down the server gracefully
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		log.Fatalf("Server forced to shutdown: %v\n", err)
	}

	log.Println("Server exiting...")
}

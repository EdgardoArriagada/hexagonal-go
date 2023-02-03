package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

const httpAddr = ":8080"

func main() {
	fmt.Println("Starting server on port", httpAddr)

	srv := gin.new()
	srv.GET("/health", healthHandler)

	log.Fatal(srv.Run(httpAddr))
}

func healthHandler(ctx *gin.Context) {
	ctx.String(http.StatusOK, "OK")
}

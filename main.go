package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/pajri/url-shortener/docs" // This is necessary for the generated docs to be included
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"github.com/pajri/url-shortener/handler"
)

func main() {
	// http.HandleFunc("/", handler.HomeHandler)
	// http.HandleFunc("/shorten", handler.ShortenHandler)
	// http.HandleFunc("/r/", handler.RedirectHandler)

	r := gin.Default()

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler)) //access swagger in /swagger/index.html

	r.GET("/", WrapHandler(handler.HomeHandler))
	r.POST("/shorten", WrapHandler(handler.ShortenHandler))
	r.GET("/r/:shortURL", WrapHandler(handler.RedirectHandler))

	fmt.Println("Starting server on :8080")
	log.Fatal(r.Run(":8080"))

	fmt.Println("Starting server on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func WrapHandler(h http.HandlerFunc) gin.HandlerFunc {
	return func(c *gin.Context) {
		h(c.Writer, c.Request)
	}
}

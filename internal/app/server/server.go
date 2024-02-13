package server

import (
	"github.com/gin-gonic/gin"
	"github.com/stepanorama/url-shortener/internal/app/config"
	"github.com/stepanorama/url-shortener/internal/app/handlers"
)

func RunServer() {
	config.ParseFlags()
	r := gin.Default()
	r.POST("/", handlers.CreateShortURL)
	r.GET("/:short_url", handlers.GetFullURL)

	err := r.Run(config.Addresses.Server)
	if err != nil {
		panic(err)
	}
}

package server

import (
	"github.com/gin-gonic/gin"
	"github.com/stepanorama/url-shortener/internal/app/config"
	"github.com/stepanorama/url-shortener/internal/app/handlers"
	"github.com/stepanorama/url-shortener/internal/app/storage"
)

func RunServer() {
	config.ParseFlags()
	r := gin.Default()

	storer := storage.NewMapStorage()         // Create storage
	handler := handlers.NewURLHandler(storer) // Inject storage into handler

	r.POST("/", handler.CreateShortURL)
	r.GET("/:short_url", handler.GetFullURL)

	err := r.Run(config.Addresses.Server)
	if err != nil {
		panic(err)
	}
}

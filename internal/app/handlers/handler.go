package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/stepanorama/url-shortener/internal/app/utils"
	"net/http"
)

// URLStorer defines the interface for URL storage operations. We're going to use it as a dependency for our handler.
// see example https://bryanftan.medium.com/accept-interfaces-return-structs-in-go-d4cab29a301b
type URLStorer interface {
	StoreURL(shortURL, fullURL string) error
	RetrieveURL(shortURL string) (string, bool)
}

type URLHandler struct {
	storer URLStorer
}

func NewURLHandler(storer URLStorer) *URLHandler {
	return &URLHandler{storer: storer}
}

func (h *URLHandler) CreateShortURL(c *gin.Context) {
	// Implementation similar to the original, but use h.storer to store the URL
	c.Header("content-type", "text/plain; charset=utf-8")
	body, err := c.GetRawData()
	if err != nil {
		c.JSON(http.StatusBadRequest, errorResponse(err.Error()))
		return
	}

	if len(body) == 0 {
		c.JSON(http.StatusBadRequest, errorResponse("Body is empty"))
		return
	}

	scheme := "http"
	if c.Request.TLS != nil {
		scheme = "https"
	}
	shortURL := utils.RandString(10)
	// previous solution storage.URLMap[shortURL] = string(body)
	h.storer.StoreURL(shortURL, string(body))

	c.String(http.StatusCreated, "%v://%v%v%v", scheme, c.Request.Host, c.Request.RequestURI, shortURL)
}

func (h *URLHandler) GetFullURL(c *gin.Context) {
	// Implementation similar to the original, but use h.storer to retrieve the URL
	shortURL := c.Params.ByName("short_url")

	if fullURL, ok := h.storer.RetrieveURL(shortURL); ok {
		c.Header("Location", fullURL)
		c.Status(http.StatusTemporaryRedirect)
	} else {
		c.Status(http.StatusBadRequest)
	}
}

func errorResponse(err string) gin.H {
	return gin.H{"error": err}
}

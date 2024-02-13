package utils

import (
	"github.com/stepanorama/url-shortener/internal/app/storage"
	"math/rand"
	"time"
)

var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func init() {
	rand.New(rand.NewSource(time.Now().UnixNano())) // Replace deprecated rand.Seed()
	storage.URLMap = map[string]string{}
}

func RandString(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

package main

import (
	"github.com/damoncheng/gotutorial/pkg/web"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/albums", web.GetAlbums)
	router.GET("/albums/:id", web.GetAlbumByID)
	router.POST("/albums", web.PostAlbums)
	router.Run("localhost:8080")
}

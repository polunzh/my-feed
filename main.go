package main

import (
	"net/http"
	"polunzh/my-feed/dal/feed"
	"polunzh/my-feed/model"

	"github.com/gin-gonic/gin"
)

func handlePostFeeds(ctx *gin.Context) {
	data := &model.Feed{}
	err := ctx.Bind(data)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	newData, err := feed.Add(ctx, *data)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "add feed failed"})
		return
	}

	ctx.JSON(http.StatusOK, newData)
}

func setupRouter() *gin.Engine {
	r := gin.Default()

	api := r.Group("/api")
	api.POST("/feeds", handlePostFeeds)

	return r
}

func main() {
	r := setupRouter()

	r.Run(":8080")
}

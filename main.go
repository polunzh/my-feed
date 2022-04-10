package main

import (
	"net/http"
	"polunzh/my-feed/dal/subscription"
	"polunzh/my-feed/model"

	"github.com/gin-gonic/gin"
)

func handlePostSubscriptions(ctx *gin.Context) {
	data := &model.Subscription{}
	err := ctx.Bind(data)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	newData, err := subscription.Add(ctx, *data)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "add subscription failed"})
		return
	}

	ctx.JSON(http.StatusOK, newData)
}

func setupRouter() *gin.Engine {
	r := gin.Default()

	api := r.Group("/api")
	api.POST("/subscriptions", handlePostSubscriptions)

	return r
}

func main() {
	r := setupRouter()

	r.Run(":8080")
}

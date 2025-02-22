package main

import (
	"fmt"
	"log/slog"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {

	r := gin.Default()

	r.POST("/parcels", func(ctx *gin.Context) {
		var rb struct {
			ID   int    `json:"id"`
			From string `json:"from"`
			To   string `json:"to"`
			Size string `json:"size"`
		}
		if err := ctx.ShouldBindJSON(&rb); err != nil {
			ctx.AbortWithStatus(http.StatusBadRequest)
			return
		}
		slog.Info(fmt.Sprintf("Received parcel %d", rb.ID))
		ctx.JSON(http.StatusOK, rb)
	})

	r.Run(":3000")
}

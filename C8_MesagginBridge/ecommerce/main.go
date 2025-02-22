package main

import (
	"encoding/json"
	"fmt"
	"msoft/g1/hackaton_g1/internal/skafka"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

type product struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Quantity int    `json:"quantity"`
}

type order struct {
	ID       int       `json:"id"`
	Products []product `json:"products"`
}

func main() {
	godotenv.Load()

	p, err := skafka.NewProducer("localhost:9092")
	if err != nil {
		panic(err)
	}

	r := gin.Default()

	r.POST("/", func(ctx *gin.Context) {
		var req order
		if err := ctx.ShouldBindBodyWithJSON(&req); err != nil {
			fmt.Println(err)
			ctx.AbortWithStatus(http.StatusBadRequest)
			return
		}
		data, _ := json.Marshal(&req)
		err = p.Send("orders", string(data))
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
			return
		}
		ctx.Status(http.StatusOK)
	})

	r.Run(":80")
}

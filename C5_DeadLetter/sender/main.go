package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/redis/go-redis/v9"
)

type parcelReq struct {
	ID   int    `json:"id"`
	From string `json:"from"`
	To   string `json:"to"`
	Size string `json:"size"`
}

func main() {
	rdb := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})
	defer rdb.Close()

	client := http.Client{Timeout: 2 * time.Second}

	resendParcels := make(chan struct{})

	sigchan1 := make(chan os.Signal, 2)
	sigchan2 := make(chan os.Signal, 1)
	sigchan3 := make(chan os.Signal, 1)
	signal.Notify(sigchan1, syscall.SIGINT, syscall.SIGTERM)
	signal.Notify(sigchan2, syscall.SIGINT, syscall.SIGTERM)
	signal.Notify(sigchan3, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		s := rdb.Subscribe(context.Background(), "failedParcels")
		defer s.Close()
		ch := s.Channel()
		for {
			select {
			case <-resendParcels:
			case <-sigchan1:
				return
			}
			for {
				breakFor := false
				select {
				case msg := <-ch:
					var p parcelReq
					json.Unmarshal([]byte(msg.Payload), &p)
					client.Post("http://localhost:3000/parcels", "application/json", bytes.NewBuffer([]byte(msg.Payload)))
					slog.Info(fmt.Sprintf("Parcel reenviado %d", p.ID))
				case <-time.After(5 * time.Second):
					resendParcels <- struct{}{}
					breakFor = true
					break
				case <-sigchan1:
					slog.Info("Fin routine secundaria")
					return
				}
				if breakFor {
					break
				}
			}
		}
	}()

	go func() {
		idx := 1
		serviceState := "init" // "restart" | "failed"
		for {
			req := parcelReq{
				ID:   idx,
				From: fmt.Sprintf("City %d", idx),
				To:   fmt.Sprintf("City %d", idx+1),
				Size: fmt.Sprintf("%d kgs", idx*10),
			}
			slog.Info(fmt.Sprintf("Sending parcel %d", idx))
			data, _ := json.Marshal(req)
			_, err := client.Post("http://localhost:3000/parcels", "application/json", bytes.NewBuffer(data))
			if err != nil {
				serviceState = "failed"
				err = rdb.Publish(context.Background(), "failedParcels", string(data)).Err()
				if err != nil {
					slog.Error(err.Error())
				}
				slog.Info(fmt.Sprintf("Parcel failed %d", idx))
			} else if serviceState == "failed" {
				serviceState = "restart"
				resendParcels <- struct{}{}
				<-resendParcels
			}
			time.Sleep(2 * time.Second)
			select {
			case <-sigchan2:
				slog.Info("Fin routine main")
				return
			default:
			}
			idx++
		}
	}()

	<-sigchan3
	slog.Info("End")
}

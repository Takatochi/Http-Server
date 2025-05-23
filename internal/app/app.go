package app

import (
	"context"
	"fmt"
	"htt-server/internal/server"
	"os"
	"os/signal"
	"time"
)

func Run() (err error) {
	//localhost-> 8080
	srv := server.NewHttp(":8080")

	go func() {
		if err := srv.Start(); err != nil {
			panic(err)
		}
	}()

	stop := make(chan os.Signal, 1) // 1
	signal.Notify(stop, os.Interrupt)
	<-stop
	fmt.Print("Shutting down...\n")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Stop(ctx); err != nil {
		return err
	}
	return nil
}

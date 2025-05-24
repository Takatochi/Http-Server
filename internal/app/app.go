package app

import (
	"context"
	"errors"
	"fmt"
	"htt-server/internal"
	"htt-server/internal/server"
	"log"
	"os"
	"os/signal"
	"time"
)

func say() {
	for i := 0; i < 10; i++ {
		fmt.Println("Hello")
		time.Sleep(1 * time.Second)
	}
}

func Run() (err error) {
	//localhost-> 8080
	srv := server.NewHttp(":8080")

	go func() {
		if err := srv.Start(); err != nil {
			var srvE *internal.Error
			if errors.As(err, &srvE) {
				log.Panic(srvE.Error())
			} else {
				log.Panic(srvE.Error())
			}
		}
	}()
	//																		     -> M -> os -> APi-> CPU
	//																		     -> M -> os -> APi-> CPU
	//                                    						    [6,7,5] -> P -> M -> os -> APi-> CPU
	// горутина -> lifo [1,2,3,4,6,7,5,7,10] -> [6,7,5,7,10,3,2,1]  [5,7,10]-> P
	//									  						    [3,2,1]	-> P
	//
	// P:M ->
	// G - горутина (легкий потік) -> обєкт стек
	// P - процес (тяжкий потік) -> абстракція ОС потоком
	// M - потік (machine *thread* ) -> CPU
	//
	go say()

	stop := make(chan os.Signal, 1) // 1
	signal.Notify(stop, os.Interrupt)
	<-stop
	fmt.Print("Shutting down...\n")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	defer close(stop)
	if err := srv.Stop(ctx); err != nil {
		return err
	}

	/// <- cancel()
	/// <- close(stop)
	return nil
}

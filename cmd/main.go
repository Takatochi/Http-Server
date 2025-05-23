package main

import (
	"fmt"
	"htt-server/internal/app"
)

func main() {

	err := app.Run()
	if err != nil {
		fmt.Println(err)
	}
}

package main

import (
	"context"
	"fmt"
	"time"

	govee "github.com/Khabi/go-vee"
)

func main() {
	ctx, cancel := context.WithCancel(context.TODO())
	go func() {
		time.Sleep(30 * time.Second)
		fmt.Println("Canceling!")
		cancel()
	}()
	govee.Run(ctx)
}

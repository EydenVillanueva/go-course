package main

import (
	"context"
	"fmt"
	"time"
)

func ping(ctx context.Context, ch chan<- string) {
	for {
		select {
		case <-ctx.Done():
			return
		case ch <- fmt.Sprintf("ping: %v", time.Now()):
			time.Sleep(1 * time.Second)
		}
	}
}

func pong(ctx context.Context, ch chan<- string) {
	for {
		select {
		case <-ctx.Done():
			return
		case ch <- fmt.Sprintf("pong: %v", time.Now()):
			time.Sleep(1 * time.Second)
		}
	}
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	pingerChan := make(chan string)
	done := make(chan struct{})

	go ping(ctx, pingerChan)
	go pong(ctx, pingerChan)

	go func() {
		timeout := time.After(5 * time.Second)
		for {
			select {
			case <-timeout:
				fmt.Println("Operation completed")
				close(pingerChan)
				done <- struct{}{}
				return

			case msg := <-pingerChan:
				fmt.Println(msg)
			}
		}
	}()

	<-done
	fmt.Println("done")

}

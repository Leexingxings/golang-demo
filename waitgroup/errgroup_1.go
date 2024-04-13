package main

import (
	"context"
	"fmt"
	"golang.org/x/sync/errgroup"
	"log"
	"time"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*2)
	defer cancel()

	g, ctx := errgroup.WithContext(ctx)

	// 协程1
	g.Go(func() error {
		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
			fmt.Println("something1")
			return nil
		}
	})

	// 协程2
	g.Go(func() error {
		select {
		case <-ctx.Done():
			return ctx.Err()
			// 模拟耗时操作3s
		case <-time.After(time.Second * 3):
			fmt.Println("something2")
			return nil
		}
	})

	if err := g.Wait(); err != nil {
		log.Fatal("you are failed: ", err)
	}
}

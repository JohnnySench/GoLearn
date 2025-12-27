package main

import (
	"context"
	"fmt"
	"time"
)

func foo(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("У Foo был отменен контекст")
			return
		default:
			fmt.Println("Foo select default")
		}
		time.Sleep(100 * time.Millisecond)
	}
}

func boo(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("У Boo был отменен контекст")
			return
		default:
			fmt.Println("Boo select default")
		}
		time.Sleep(100 * time.Millisecond)
	}
}

func main() {
	parentCtx, parentCtxCancel := context.WithCancel(context.Background())
	childCtx, childCtxCancel := context.WithCancel(parentCtx)
	go foo(parentCtx)
	go boo(childCtx)

	time.Sleep(1 * time.Second)
	childCtxCancel()

	time.Sleep(1 * time.Second)
	parentCtxCancel()

	time.Sleep(3 * time.Second)
}

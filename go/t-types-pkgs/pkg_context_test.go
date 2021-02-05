package ttp

import (
	"context"
	"fmt"
	"testing"
	"time"
)

func Test_ctxTimeout(t *testing.T) {
	func() {
		ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
		defer cancel()

		go handle(ctx, 500*time.Millisecond)
		// go handle(ctx, 1500*time.Millisecond)
		select {
		case <-ctx.Done():
			fmt.Println("main", ctx.Err())
		}
	}()
}

func Test_ctxWithDeadline(t *testing.T) {
	func() {
		ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(1*time.Second))
		defer cancel()

		go handle(ctx, 500*time.Millisecond)
		select {
		case <-ctx.Done():
			fmt.Println("man", ctx.Err())
		}

	}()
}

func handle(ctx context.Context, duration time.Duration) {
	go handle2(ctx, duration)
	select {
	case <-ctx.Done():
		fmt.Println("handle 1", ctx.Err())
	case <-time.After(duration):
		fmt.Println("process request with", duration)
	}
}

func handle2(ctx context.Context, duration time.Duration) {
	select {
	case <-ctx.Done():
		fmt.Println("handle 2", ctx.Err())
	case <-time.After(duration):
		fmt.Println("process request with", duration)
	}
}

package main

import (
	"context"
	"sync"
	"time"
)

func RunResyncLoop(ctx context.Context, interval time.Duration) {
	t := time.NewTicker(interval)
	defer t.Stop()

	for {
		select {
		case <-ctx.Done():
			println("trigger by context")
			return
		case <-t.C:
		}
		println("trigger by ticker")
	}
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	var wg sync.WaitGroup
	interval := 5000 * time.Millisecond
	//wg.Go(func() {
	//	RunResyncLoop(ctx, interval)
	//})
	wg.Add(1)
	go func() {
		defer wg.Done()
		RunResyncLoop(ctx, interval)
	}()
	time.Sleep(10000 * time.Millisecond)
	cancel()
	wg.Wait()
}

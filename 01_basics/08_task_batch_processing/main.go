package main

import (
	"context"
	"fmt"
	"time"
)

func main() {

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	input := make(chan int)
	go StartBatchProcessor(ctx, input)
	go func() {
		for i := 1; i <= 20; i++ {
			input <- i
		}
	}()
	<-ctx.Done()
	fmt.Println("Main: processing stopped")

}

func StartBatchProcessor(ctx context.Context, input <-chan int) {
	batchSize := 5
	batch := make([]int, 0, batchSize)
	timer := time.NewTimer(2 * time.Second)
	for {
		select {
		case <-ctx.Done():
			fmt.Println("контекст отменен")
			return
		case value := <-input:
			batch = append(batch, value)
			if len(batch) == batchSize {
				fmt.Println("Processed batch:", batch)
				batch = batch[:0]
			}
			if !timer.Stop() {
				<-timer.C
				timer.Reset(2 * time.Second)
			}
		case <-timer.C:
			if len(batch) > 0 {
				fmt.Println("Processed batch (timeout):", batch)
				batch = batch[:0]
				timer.Reset(2 * time.Second)
			}

		}
	}
}

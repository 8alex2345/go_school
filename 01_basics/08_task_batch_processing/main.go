package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	input := make(chan int)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	wg.Add(1)
	go StartBatchProcessor(ctx, input, &wg)
	go func() {
		defer wg.Done()

		for i := 0; i <= 20; i++ {
			select {
			case input <- i:
				fmt.Println("отправялем в канал число - ", i)
			case <-ctx.Done():
				fmt.Println("конеткст отменен. прекращаем отправку")
				return
			}
		}
		close(input)

	}()
	wg.Wait()
	fmt.Println("Main: processing stopped")
}
func StartBatchProcessor(ctx context.Context, input <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	batchSize := 5
	batch := make([]int, 0, batchSize)
	timer := time.NewTimer(2 * time.Second)
	defer timer.Stop()

	for {
		select {
		case value, ok := <-input:
			if !ok {
				if len(batch) > 0 {
					fmt.Println("канал закрыт, обрабатываем последний батч. - ", batch)

				} else {
					fmt.Println("Processed batch:", batch)
				}
				return
			}
			batch = append(batch, value)
			if len(batch) == batchSize {
				fmt.Println("Processed batch:", batch)
				batch = batch[:0]
				if !timer.Stop() {
					<-timer.C
				}
				timer.Reset(2 * time.Second)
			}

		case <-timer.C:
			if len(batch) > 0 {
				fmt.Println("Processed batch:", batch)
				batch = batch[:0]
			}
			timer.Reset(2 * time.Second)

		case <-ctx.Done():
			fmt.Println("контекст отменен")
			if len(batch) > 0 {
				fmt.Println("Processed batch:", batch)
			}
			return

		}

	}
}

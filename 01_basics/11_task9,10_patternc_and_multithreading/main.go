package main

import "sync"

func main() {

}
func Merge(cs ...<-chan int) <-chan int {
	out := make(chan int)
	var wg sync.WaitGroup

	go func() {
		for _, ch := range cs {
			wg.Add(1)
			go func() {
				defer wg.Done()
				for v := range ch {
					out <- v
				}
			}()
		}
		wg.Wait()
		close(out)
	}()

	return out
}

package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

func main() {
	fmt.Println("Задача - 1")
	GoroutineWaitGroup()
	fmt.Println("Задача - 2")
	StartFiveGorotine()
	fmt.Println("Задача - 3")
	var sum int
	for data := range UseChanel() {
		sum = sum + data
	}
	fmt.Println(sum)
	fmt.Println("Задача - 4")
	UseMutex()
	fmt.Println("Задача - 5")
	UseAtomic()
}

// 1. **Запуск горутины и ожидание её завершения**
// Задача: Напишите функцию, которая запускает горутину, выполняющую fmt.Println("Hello from goroutine!"), и использует sync.WaitGroup для ожидания её завершения.
func GoroutineWaitGroup() {
	var wg sync.WaitGroup
	ch := make(chan string)

	wg.Add(2)
	go func() {
		defer wg.Done()
		ch <- "Hello from goroutine!"
	}()
	go func() {
		defer wg.Done()
		msg := <-ch
		fmt.Println(msg)
	}()
	wg.Wait()
}

// Дргуие способы дождаться выполнения горутины
// 1. синхронизация через закрытие канала
//2. sync.Mutex

func StartFiveGorotine() {
	var wg sync.WaitGroup
	ch := make(chan string)
	wg.Add(5)
	for _, name := range []string{
		"Первая горутина",
		"Вторая горутина",
		"Третья горутина",
		"Четвертая горутина",
		"Пятая горутина",
	} {
		go func(n string) {
			defer wg.Done()
			ch <- n
		}(name)
	}
	go func() {
		wg.Wait()
		close(ch)
	}()
	for msg := range ch {
		fmt.Println(msg)
	}
}

//3. **Использование канала для передачи данных**
//Задача: Напишите функцию, которая создает горутину, отправляющую числа от 1 до 5 в канал, а затем в main извлекает их и складывает, результат выводит в консоль.
//Вопрос: когда заканчивается чтение с канала? Что если не закрывать канал?
// Канал лучше с буфером или без? В чём отличие?
// Что будет если использовать select? Как реализовать?

// когда заканчивается чтение с канала? Чтение из канала заканчивается, когдьа канал закрывается.

// Что если не закрывать канал? Цикл будет ждать новые значение. Деадлоок

//Канал лучше с буфером или без? В чём отличие? какой лучше зависит от задач, небуф используется когда важня строгая синхронизаци,
//  а Буф когда нужно временно хранить значения, чтобб отправитель не ждал получателя

// Что будет если использовать select? Как реализовать? селект это как свич касе только для каналов

func UseChanel() chan int {
	ch := make(chan int)
	go func() {
		for i := 1; i <= 5; i++ {
			ch <- i
		}
		close(ch)
	}()
	return ch

}

// 4 Потокобезопасный инкремент - Mutex.**
// Задача: Напишите программу, где 10 горутин инкрементируют один счётчик, защищая его sync.Mutex.
//     1. Что если не обложить мютексом? Воспроизвести race condition.

// несколько горутин будут обращаться к одним и тем же данным. и ркзультат может  быть неправильным

func UseMutex() {
	var wg sync.WaitGroup
	var mu sync.Mutex
	var counter int
	wg.Add(10)
	for i := 0; i < 10; i++ {
		go func() {
			defer wg.Done()
			time.Sleep(time.Microsecond)
			mu.Lock()
			counter++
			mu.Unlock()
		}()
	}
	wg.Wait()
	fmt.Println(counter)
}

//  5. **Потокобезопасный инкремент - Atomic.
// Задача: Напишите программу, где 10 горутин инкрементируют один счётчик без использования мютексов, через атомики.
// Что, если не использовать атомик? Что лучше, атомик или мютекс?

// мьютех работает дольше, но побходит под больше задач.Можно выполнять в лооках несколько действий.
//
//	атомик работает быстро но может быть использован для одной задачи например увеличить счетчкик.
func UseAtomic() {
	var wg sync.WaitGroup
	var count int64
	wg.Add(10)
	for i := 0; i < 10; i++ {
		go func() {
			defer wg.Done()

			atomic.AddInt64(&count, 1)

		}()
	}
	wg.Wait()
	fmt.Println(count)
}

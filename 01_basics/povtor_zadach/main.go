package main

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/dgrijalva/jwt-go"
)

func main() {
	// задача 1
	HelloGoroutine()
	// задача 2
	FiveGoroutine()
	ChannalsFiveGoroutine()
	// задача 3
	for num := range NumberChannals(5) {
		fmt.Println(num)
	}
	// задача 4
	MutexGoroutine()

	// задача 5
	// ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	// defer cancel()
	// input := make(chan int)
	// go StartBatchProcessor(ctx, input)

	// for i := 0; i <= 20; i++ {
	// 	input <- i
	// }
	// close(input)
	// fmt.Println("Main: processing stopped")
	// задача 7
	ctx := context.Background()
	userID := 1234
	ctx, err := AddJWTToCintext(ctx, userID)
	if err != nil {
		fmt.Println("ошибка при добавлении JWT в контекс")
	}
	extractUserID, err := ExtractUserIDFromContext(ctx)
	if err != nil {
		fmt.Println("Ошибка при извлечении из контекса - ", err)
	}
	fmt.Println(extractUserID)

}

// Задача1 : Напишите функцию, которая запускает горутину, выполняющую fmt.Println("Hello from goroutine!"), и использует sync.WaitGroup для ожидания её завершения.
func HelloGoroutine() {
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		fmt.Println("Hello from goroutine!")
	}()
	wg.Wait()
}

// Задача 2: Напишите программу, которая запускает 5 горутин, каждая из которых печатает свой номер (от 1 до 5),
//
//	и использует sync.WaitGroup для их синхронизации(нужно подождать их выполнения).
func FiveGoroutine() {
	var wg sync.WaitGroup
	wg.Add(5)
	for i := 1; i <= 5; i++ {
		go func(num int) {
			defer wg.Done()
			fmt.Println("Горутина номер - ", num)
		}(i)
	}
	wg.Wait()

}
func ChannalsFiveGoroutine() {
	channals := make(chan string)
	var wg sync.WaitGroup
	wg.Add(5)
	for _, name := range []string{
		"первая",
		"вторая",
		"третья",
		"четвертая",
		"пятая",
	} {
		go func(value string) {
			defer wg.Done()
			channals <- name
		}(name)
	}
	go func() {
		wg.Wait()
		close(channals)
	}()
	for msg := range channals {
		fmt.Println(msg)
	}
}

// Задача3 : Напишите функцию, которая создает горутину, отправляющую числа от 1 до 5 в канал, а затем в main извлекает их и складывает, результат выводит в консоль.
func NumberChannals(num int) chan int {
	channals := make(chan int)
	var wg sync.WaitGroup
	wg.Add(num)
	for i := 1; i <= num; i++ {
		go func() {
			defer wg.Done()
			channals <- i
		}()
	}
	go func() {
		wg.Wait()
		close(channals)
	}()

	return channals
}

// Задача 4: Напишите программу, где 10 горутин инкрементируют один счётчик, защищая его sync.Mutex.
func MutexGoroutine() {
	var counter int
	var wg sync.WaitGroup
	var mu sync.Mutex
	wg.Add(10)
	for i := 0; i < 10; i++ {
		go func() {
			defer wg.Done()
			mu.Lock()
			counter++
			mu.Unlock()
		}()
	}
	wg.Wait()
	fmt.Println("счетчик равен = ", counter)
}

// задача 5 :Батчевая обработка
func StartBatchProcessor(ctx context.Context, input <-chan int) {
	batchSize := 5
	batch := make([]int, 0, batchSize)
	timer := time.NewTimer(2 * time.Second)
	defer timer.Stop()
	for {
		select {
		case value, ok := <-input:
			if !ok {
				if len(batch) > 0 {
					fmt.Println("Processed batch:", batch)
				}
				return
			}
			batch = append(batch, value)
			if len(batch) == batchSize {
				fmt.Println("Processed batch:", batch)
				batch = batch[:0]
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
			return
		}
	}

}

// Задача 7 на jwt и контексты
var secretKey = []byte("secret")

type contextKey string

const jwtTokenKey contextKey = "jwtToken"

type Claims struct {
	UserID int `json:"userID"`
	jwt.StandardClaims
}

func CreateToken(userID int) (string, error) {
	claims := Claims{
		UserID: userID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(2 * time.Hour).Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(secretKey)

}
func AddJWTToCintext(ctx context.Context, userID int) (context.Context, error) {
	tokeenInt, err := CreateToken(userID)
	if err != nil {
		fmt.Println("Errors token", err)
		return nil, err
	}
	ctx = context.WithValue(ctx, jwtTokenKey, tokeenInt)

	return ctx, nil

}
func ExtractUserIDFromContext(ctx context.Context) (int, error) {
	if token, ok := ctx.Value(jwtTokenKey).(string); ok {
		fmt.Println("JWT токен из контекста", token)
		claims := &Claims{}
		tokenParsed, err := jwt.ParseWithClaims(token, claims, func(t *jwt.Token) (interface{}, error) {
			return secretKey, nil
		})
		if err != nil {
			return 0, nil
		}
		if tokenParsed.Valid {
			return claims.UserID, nil
		}
	}

	return 0, fmt.Errorf("tокен не найден")
}

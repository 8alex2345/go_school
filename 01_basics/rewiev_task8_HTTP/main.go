package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"sync"
	"time"
)

//## 8. Практическая задача:
// Конкурентный HTTP-запрос к списку URL-адресов с использованием горутин и синхронизации
// > **Задача:**
// **Напишите функцию FetchURLs(urls []string) map[string]string, которая:**
// *Принимает слайс URL-адресов.
// Конкурентно делает HTTP-запросы к каждому URL.
// Собирает результаты (код ответа и часть тела) в map[string]string, где:
// ключ — URL
// значение — содержимое ответа (ограниченное, например, 100 символами)
// Использует sync.WaitGroup и sync.Mutex для защиты записи в map.
// В случае ошибки записывает "error" как значение.*

func main() {
	urls := []string{
		"http://example.com",
		"http://example.org",
	}
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	results := FetchURLs(ctx, urls)
	for url, body := range results {
		fmt.Printf("Ответ от %s:\n%s\n\n", url, body)
	}
}

func FetchURLs(ctx context.Context, urls []string) map[string]string {
	client := &http.Client{
		Timeout: 10 * time.Second,
	}
	var wg sync.WaitGroup
	result := make(map[string]string)
	mu := sync.Mutex{}

	for _, url := range urls {
		wg.Add(1)
		go func(url string) {
			defer wg.Done()
			req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
			if err != nil {
				fmt.Println("Ошибка при создании запроса: ", err)
				return
			}
			resp, err := client.Do(req)
			if err != nil {
				mu.Lock()
				fmt.Println("Ошибка при выполнении запроса: ", err)
				mu.Unlock()
				return
			}
			defer resp.Body.Close()

			body, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				fmt.Println("Ошибка при чтении ответа: ", err)
				return

			}
			content := string(body)
			if len(content) > 100 {
				content = content[:100]
			}
			mu.Lock()
			result[url] = content
			mu.Unlock()
		}(url)
	}
	wg.Wait()
	return result
}

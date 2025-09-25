package main

import (
	"context"
	"errors"
	"fmt"
	"sync"
	"time"

	"github.com/dgrijalva/jwt-go"
)

// > ЗАДАЧА 2. Работа с JWT в контексте
// > **Условие**:
// > Напиши две функции:
// > 1. `AddJWTToContext(ctx context.Context, userID int) (context.Context, error)`
// > 2. `ExtractUserIDFromContext(ctx context.Context) (int, error)`
// > `AddJWTToContext` должен:
// > - Создавать JWT-токен (используй `github.com/golang-jwt/jwt/v5`).
// Зашифровывать в него `userID`.
// Возвращать новый `context.Context`, в который записан JWT-токен.
// `ExtractUserIDFromContext` должен:
// Извлекать JWT-токен из контекста.
// Расшифровывать `userID`.
// Вывести его на экран
// **Дополнительное условие**:
// Создай горутину, в которой будет использоваться `ExtractUserIDFromContext`.
// Покажи, что передача контекста работает и данные можно безопасно извлекать между горутинами.
func main() {
	ctx := context.Background()
	ctx, _ = AddJWTToContext(ctx, 4885)
	var wg sync.WaitGroup
	wg.Add(1)
	go func(ctx context.Context) {
		defer wg.Done()
		userID, err := ExtractUserIDFromContext(ctx)
		if err != nil {
			panic(err)
		}
		fmt.Println("UserID - ", userID)
	}(ctx)
	wg.Wait()
}

type contextKey string

const jwtKey contextKey = "jwt"

func AddJWTToContext(ctx context.Context, userID int) (context.Context, error) {
	var myKey = []byte("secret")
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["userID"] = userID
	claims["exp"] = time.Now().Add(time.Hour * 1).Unix()

	tokenString, err := token.SignedString(myKey)
	if err != nil {
		return nil, err
	}
	return context.WithValue(ctx, jwtKey, tokenString), nil
}
func ExtractUserIDFromContext(ctx context.Context) (int, error) {
	tokenStr, ok := ctx.Value(jwtKey).(string)
	if !ok {
		return 0, errors.New("jwt not found in context")
	}
	var meKey = []byte("secret")
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return meKey, nil
	})
	if err != nil {
		return 0, err
	}
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		userID := int(claims["userID"].(float64))
		return userID, nil
	}
	return 0, errors.New("invalid token")
}

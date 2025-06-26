package main

import (
	"fmt"
	"strings"
)

func main() {
	// ### **Задача 1: Подсчёт количества цифр в числе**
	fmt.Println("Задача 1")
	fmt.Println(QuantityNum(1213124))

	// Задача 2: Подсчёт частоты символов в строке
	fmt.Println("Задача 2")

	str := "Hello"
	res := CounterSymbol(str)
	fmt.Print("map[")
	for key, value := range res {
		fmt.Printf("%c: %d ", key, value)
	}
	fmt.Println("]")

	// Задача 3: Инвертирование мапы
	fmt.Println("Задача 3")
	animals := map[string]string{
		"cat": "кот",
		"dog": "собака",
		"a":   "x",
		"b":   "x",
	}
	fmt.Println(InvertMap(animals))

	// Задача 4: Индексирование слов
	fmt.Println("Задача 4")
	fmt.Println(IndexWords("go is fun and go is fast"))

	// Задача 5: Сравнение двух мап

	fmt.Println("Задача 5")
	map1 := map[string]int{"a": 1, "b": 2}
	map2 := map[string]int{"b": 2, "a": 1}
	fmt.Println(ComparisoMap(map1, map2))
	// Задача 6: Найти ключ по значению
	myMap := map[string]int{
		"x": 10,
		"y": 20,
		"z": 40,
	}
	fmt.Println("Задача 6")
	fmt.Println(FindKey(myMap))
}

// ### **Задача 1: Подсчёт количества цифр в числе**
// Напиши функцию, которая принимает целое число и возвращает
//
//	`map[int]int`, где ключ — цифра, а значение — сколько раз она встречается.
func QuantityNum(numbers int) map[int]int {
	if numbers < 0 {
		numbers = -numbers
	}
	if numbers == 0 {
		return map[int]int{0: 1}
	}
	slicesNum := []int{}
	for numbers > 0 {
		digit := numbers % 10
		slicesNum = append([]int{digit}, slicesNum...)
		numbers = numbers / 10
	}
	myMap := make(map[int]int)
	for _, elem := range slicesNum {
		myMap[elem]++
	}
	return myMap
}

//Задача 2: Подсчёт частоты символов в строке
//Прочитай строку и создай map[rune]int, где хранятся все символы и их количество.

func CounterSymbol(str string) map[rune]int {
	myMap := make(map[rune]int)
	for _, elem := range str {
		myMap[elem]++
	}
	return myMap
}

// Задача 3: Инвертирование мапы
// Есть мапа map[string]string. Поменяй местами ключи и значения.
func InvertMap(myMap map[string]string) map[string]string {
	resMap := make(map[string]string)
	for key, value := range myMap {
		resMap[value] = key
	}
	return resMap
}

// Задача 4: Индексирование слов
// Разбей строку на слова и составь мапу `map[string]int`,
// где каждому слову соответствует индекс его **первого вхождения** в списке.
func IndexWords(word string) map[string]int {
	wordsStr := strings.Fields(word)
	resMap := make(map[string]int)
	for idx, value := range wordsStr {
		if _, exists := resMap[value]; !exists {
			resMap[value] = idx

		}

	}
	return resMap

}

// Задача 5: Сравнение двух мап
// Напиши функцию, которая сравнивает две `map[string]int` и возвращает `true`,
//
//	если они одинаковы (одинаковые ключи и значения), иначе `false`.
func ComparisoMap(map1, map2 map[string]int) bool {
	if len(map1) != len(map2) {

		return false
	}
	for key, value := range map1 {
		if map2Value, exists := map2[key]; !exists || map2Value != value {
			return false
		}

	}

	return true

}

// Задача 6: Найти ключ по значению
//Напиши функцию, которая ищет **ключ по значению** в мапе `map[string]int`.
//  Верни первый найденный ключ (или `"not found"`).

func FindKey(myMap map[string]int) string {
	meaning := 20
	for key, vaule := range myMap {
		if vaule == meaning {
			return key

		}
	}
	return "not found"

}

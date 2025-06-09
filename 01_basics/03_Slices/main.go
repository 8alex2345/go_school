// 1. Удаление элемента по индексу
//Ввод: s := []int{1, 2, 3, 4, 5}, index = 2
// Вывод: [1 2 4 5]

package main

import (
	"fmt"
	"strings"
)

func main() {
	s := []int{1, 2, 3, 4, 5}
	idx := 2
	s = append(s[:idx], s[idx+1:]...)
	fmt.Println(s)

}

// Слайсы с общей памятью
// Создай слайс и подрежь его. Затем измени значение в одном — отразится ли оно на другом?
func SliceMemory() {
	s := []int{1, 2, 3, 4, 5}
	t := s[1:4]
	t[0] = 99
	fmt.Println(s)
	// Полученый слайс [1 99 3 4 5] подслайс t содержит значения с первого по 3 индекс слайса s. для t нклевой индекс это первый для s
}

// Циклический сдвиг влево
// Сдвинь все элементы слайса влево на 1, а первый — в конец.
func CycleShift(s []int) []int {
	s = []int{10, 20, 30, 40}
	first := s[0]
	for i := 0; i < len(s)-1; i++ {
		s[i] = s[i+1]
	}
	s[len(s)-1] = first
	return s
}

// Объединение строк в одну
// Дан слайс строк, например []string{"Go", "is", "awesome"}. Объедини их в одну строку через пробел.
func MergeLines(slice []string) string {
	slice = []string{"Go", "is", "awesome"}
	//result := slice[0] + " " + slice[1] + " " + slice[2]
	result := strings.Join(slice, " ")
	return result

}

// Фильтрация чётных чисел
// Напиши функцию, которая принимает `[]int` и возвращает новый слайс, содержащий только чётные числа.
func FiltrationEven(num []int) []int {
	result := []int{}
	for _, elem := range num {
		if elem%2 == 0 {
			result = append(result, elem)
		}
	}
	return result
}

// Проверка на палиндром слайса
// Проверь, является ли слайс палиндромом (одинаковый слева направо и справа налево).
func Palindrome(slice []int) string {
	left := 0
	right := len(slice) - 1
	for left < right {
		if slice[left] != slice[right] {
			return " Не Палиндром"
		}
		left++
		right--
	}
	return "Палиндром"

}

//

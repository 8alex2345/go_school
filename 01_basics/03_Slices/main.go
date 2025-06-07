// 1. Удаление элемента по индексу
//Ввод: s := []int{1, 2, 3, 4, 5}, index = 2
// Вывод: [1 2 4 5]

package main

import "fmt"

func main() {
	s := []int{1, 2, 3, 4, 5}
	idx := 2
	s = append(s[:idx], s[idx+1:]...)
	fmt.Println(s)
	FiltrationEven([]int{1, 2, 3, 4, 5, 6})
	Palindrome([]int{1, 2, 3, 2, 1})

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
func CycleShift() {
	s := []int{10, 20, 30, 40}
	first := s[0]
	for i := 0; i < len(s)-1; i++ {
		s[i] = s[i+1]
	}
	s[len(s)-1] = first
	fmt.Println(s)
}

// Объединение строк в одну
// Дан слайс строк, например []string{"Go", "is", "awesome"}. Объедини их в одну строку через пробел.
func MergeLines() {
	slice := []string{"Go", "is", "awesome"}
	result := slice[0] + " " + slice[1] + " " + slice[2]
	fmt.Println(result)
}

// Фильтрация чётных чисел
// Напиши функцию, которая принимает `[]int` и возвращает новый слайс, содержащий только чётные числа.
func FiltrationEven(num []int) {
	result := []int{}
	for _, elem := range num {
		if elem%2 == 0 {
			result = append(result, elem)
		}
	}
	fmt.Println(result)
}

// Проверка на палиндром слайса
// Проверь, является ли слайс палиндромом (одинаковый слева направо и справа налево).
func Palindrome(slice []int) {
	for i := 0; i < len(slice)/2; i++ {
		if slice[i] != slice[len(slice)-1-i] {
			fmt.Println(" Не Палиндром")
			return
		}
	}
	fmt.Println("Палиндром")

}

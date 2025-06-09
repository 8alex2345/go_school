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

// НОВЫЕ ЗАДАЧИ НА СЛАЙСЫ !!!!!
// Задача 1: Сумма элементов массива
// Условие:
// Создай массив из 5 целых чисел и вычисли сумму всех его элементов.
func SumElem(arr [5]int) int {
	result := 0
	for _, elem := range arr {
		result = result + elem
	}
	return result
}

// Задача 2: Поиск максимального элемента
// **Условие:**
// Найди наибольшее значение в массиве из 10 чисел.
func MaxNum(arr [10]int) int {
	maxNum := arr[0]
	for _, elem := range arr {
		if maxNum < elem {
			maxNum = elem
		}
	}
	return maxNum
}

//Задача 3: Количество чётных чисел
//Условие:
//Подсчитай, сколько в массиве из 8 элементов чётных чисел.

func EvenNum(arr [8]int) int {
	evenNum := 0
	for _, num := range arr {
		if num%2 == 0 {
			evenNum++
		}
	}
	return evenNum
}

// Задача 4: Реверс массива
//Условие:
// Выведи массив в обратном порядке.

func ReverseArray(arr [5]int) [5]int {
	for i := 0; i < len(arr)/2; i++ {
		arr[i], arr[len(arr)-1-i] = arr[len(arr)-1-i], arr[i]
	}
	return arr

}

// Задача 5: Поиск элемента
// условия:
// ПРоверить есть ли в массиве определенное число, если есть вывестьи индекс

func SearchElement(arr [4]int) int {
	num := 8
	index := 0
	for idx, elem := range arr {
		if elem == num {
			index = idx
		}
	}
	return index
}

// Задача6: Среднее значение
// Условия:
// Вычесли средние арифметические  массива чисел
func AverageValue(arr [5]float64) float64 {
	sum := 0.0
	for _, elem := range arr {
		sum += elem
	}
	result := sum / float64(len(arr))
	return result
}

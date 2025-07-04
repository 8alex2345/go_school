// 1. Вывести все квадраты натуральных чисел, не превосходящие данного числа N.

package main

import (
	"fmt"
	"math"
)

func main() {
	var n int
	fmt.Scan(&n)
	for i := 1; i*i <= n; i++ {
		fmt.Println(i * i)
	}

}

// 2. Вывести на экран кубы чисел от A до B, которые вводит пользователь.

func NumbersCubes(num, num1 int) {
	fmt.Scan(&num, &num1)
	for i := num; i <= num1; i++ {
		fmt.Println(i * i * i)
	}
}

// 3. Возведение числа в степень
// Написать программу, которая возводит число в целочисленную степень.
// Число и степень вводятся с клавиатуры.
func NumbersDegree(numbers, degree int) {
	fmt.Scan(&numbers, &degree)
	if degree == 0 {
		fmt.Println(1)
	} else {
		fmt.Println(math.Pow(float64(numbers), float64(degree)))
	}
}

// 4. Используя цикл написать программу,
// которая выводит на экран таблицу значений функции y = 5 - x2/2 на отрезке [-5; 5] с шагом 0.5.
func TableValue() {
	startX := -5.0
	endX := 5.0
	step := 0.5
	for i := startX; i <= endX; i += step {
		y := 5 - (math.Pow(i, 2) / 2)
		fmt.Printf("%.2f | %.2f\n", i, y)
	}
}

// 5. Вычислить факториал числа, которое ввел пользователь.
func NamderFactorial(f int) {
	res := 1
	for i := 1; i <= f; i++ {
		res = res * i
	}
	fmt.Println(res)
}

// 6. Вывести на экран ряд чисел Фибоначчи, состоящий из N элементов. Значение N вводится с клавиатуры.
func Fibonaccci(n int) {
	fib0 := 0
	fib1 := 1
	for i := 0; i < n; i++ {
		fmt.Println(fib0)
		fib0, fib1 = fib1, fib0+fib1
	}
}

// 7. Проверка гипотезы Сиракуз
// Возьмем любое натуральное число.
// Если оно четное - разделим его пополам, если нечетное - умножим на 3, прибавим 1 и разделим пополам.
// Повторим эти действия с вновь полученным числом. Гипотеза гласит,
// что независимо от выбора первого числа рано или поздно мы получим 1.
func Syracuse() {
	for i := 20; i <= 30; i++ {
		num := i
		for num != 1 {
			if num%2 == 0 {
				num = num / 2
			} else {
				num = num*3 + 1
			}
			fmt.Print("|", num, "|")
		}

	}
}

// 8. Извлечение цифр числа
// С клавиатуры вводится целое число.
// Определить, из каких цифр оно состоит, то есть вывести на экран отдельные цифры числа.
func ExtractDigits(num int) {
	for num > 0 {
		result := num % 10
		fmt.Print(result)
		num = num / 10
	}

}

// 9. Найти сумму четных цифр числа
// Вводится натуральное число. Найти сумму четных цифр, входящих в его состав.
func SumEvenDigits(num int) {
	var sum int
	for num > 0 {
		res := num % 10
		if res%2 == 0 {
			sum = sum + res

		}
		num = num / 10

	}
	fmt.Println(sum)
}

// 10. Количество четных и нечетных цифр числа
// Написать программу, подсчитывающую количество четных и нечетных цифр числа.
func NumEvenOdd(num int) {
	var even int
	var odd int
	for num > 0 {
		if num%2 == 0 {
			even++
		} else {
			odd++
		}
		num = num / 10
	}
	fmt.Println("Четные - ", even, "Нечетные - ", odd)
}

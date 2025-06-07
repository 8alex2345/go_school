//Оценка по баллам
// Пользователь вводит число от 0 до 100. Программа должна вывести оценку по шкале:

package main

import (
	"fmt"
)

func main() {
	var num int
	fmt.Scan(&num)
	if num >= 90 && num <= 100 {
		fmt.Println("Отлично")
	} else if num > 70 && num < 90 {
		fmt.Println("Хорошо")
	} else if num >= 50 && num < 70 {
		fmt.Println("Удовлетворительно")
	} else if num < 50 {
		fmt.Println("Неудовлетворительно")
	}
}

// Четное или нечетное
// Пользователь вводит целое число.
// Программа должна определить, является ли число чётным или нечётным.
func OddEven(num int) {
	fmt.Scan(&num)
	if num%2 == 0 {
		fmt.Println("Четное")
	} else {
		fmt.Println("Нечетное")
	}
}

// Совершеннолетие
// Пользователь вводит свой возраст.
//Если возраст 18 или больше — вывести "Доступ разрешён", иначе — "Доступ запрещён".

func Age(age int) {

	fmt.Scan(&age)
	if age >= 18 {
		fmt.Println("Доступ разрешен")
	} else {
		fmt.Println("Доступ запрещен")
	}

}

//Определение сезона
//Пользователь вводит номер месяца (от 1 до 12).
//Программа должна вывести, к какому времени года он относится:

func Season(month int) {
	fmt.Scan(&month)
	if month == 1 || month == 12 {
		fmt.Println("зима")
	} else if month == 3 || month == 5 {
		fmt.Println("весна")
	} else if month == 6 || month == 8 {
		fmt.Println("лето")
	} else if month == 9 || month == 11 {
		fmt.Println("осень")
	}
}

// Максимум из двух чисел
//Пользователь вводит два числа.
//Программа должна вывести большее из них (или сообщение, что числа равны).

func MaxNumber(num, num1 int) {
	fmt.Scan(&num, &num1)
	if num == num1 {
		fmt.Println("Числа равны")
	} else if num > num1 {
		fmt.Println(num)
	} else {
		fmt.Println(num1)
	}
}

// Число — положительное, отрицательное или ноль
//Пользователь вводит целое число.
//Определите его знак: положительное, отрицательное или ноль.

func Compresion(num int) {
	fmt.Scan(&num)
	if num < 0 {
		fmt.Println("Отрицательное")
	} else if num > 0 {
		fmt.Println("Положительное")
	} else {
		fmt.Println("Ноль")
	}
}

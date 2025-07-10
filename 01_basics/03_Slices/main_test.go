package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// Тест на Задачу 1 Сумма элементов массива

func TestSumEleme(t *testing.T) {
	// Arrange
	// Настройка тестовых данных
	testTable := []struct {
		numbers  [5]int
		expected int
	}{
		{
			numbers:  [5]int{1, 2, 3, 4, 5},
			expected: 15,
		},
		{
			numbers:  [5]int{2, 2, 2, 2, 2},
			expected: 10,
		},
		{
			numbers:  [5]int{},
			expected: 0,
		},
	}
	// Ast
	// вызов тестируемего кода
	for _, testCase := range testTable {
		result := SumElem(testCase.numbers)

		t.Logf("Входящие значения (%v), результат %d\n", testCase.numbers, result)
		assert.Equal(t, testCase.expected, result, "Неккоректный результат- ", testCase.expected)

	}

}

// Тест на Задачу 2 Поиск максимального элемента
func TestMaxNum(t *testing.T) {
	testTable := []struct {
		numbers  [10]int
		expected int
	}{
		{
			numbers:  [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			expected: 10,
		},
		{
			numbers:  [10]int{1, 1, 2, 1, 1, 1, 1, 1, 9, 5},
			expected: 9,
		},
		{
			numbers:  [10]int{1, 1, 2},
			expected: 2,
		},
	}
	for _, testCase := range testTable {
		result := MaxNum(testCase.numbers)
		t.Logf("Входящие значения (%v), Результат %d\n", testCase.numbers, result)
		assert.Equal(t, testCase.expected, result, "Ошибочный результат %d  Ожидалось %d\n", result, testCase.expected)
	}
}

//Тест Задача 3: Количество чётных чисел

func TestEvenNum(t *testing.T) {
	testTable := []struct {
		numbers  [8]int
		expected int
	}{
		{
			numbers:  [8]int{1, 2, 3, 4, 5, 6, 7, 8},
			expected: 4,
		},
		{
			numbers:  [8]int{1, 2},
			expected: 7,
		},
	}
	for _, testCase := range testTable {
		result := EvenNum(testCase.numbers)
		assert.Equal(t, testCase.expected, result)
	}
}

//Тест Задача 4: Реверс массива

func TestReverseArray(t *testing.T) {
	testTable := []struct {
		arr         [5]int
		expectedArr [5]int
	}{
		{
			arr:         [5]int{1, 2, 3, 4, 5},
			expectedArr: [5]int{5, 4, 3, 2, 1},
		},
		{
			arr:         [5]int{6, 2, 3, 22, 5},
			expectedArr: [5]int{5, 22, 3, 2, 6},
		},
	}
	for _, testCase := range testTable {
		res := ReverseArray(testCase.arr)
		assert.Equal(t, testCase.expectedArr, res)
	}
}

//Тест Задача 5: Поиск элемента

func TestSearchElement(t *testing.T) {
	testTable := []struct {
		num      [4]int
		expected int
	}{
		{
			num:      [4]int{1, 2, 3, 8},
			expected: 3,
		},
		{
			num:      [4]int{8},
			expected: 0,
		},
	}
	for _, testCase := range testTable {
		result := SearchElement(testCase.num)
		assert.Equal(t, testCase.expected, result)
	}
}

// Тест Задача6: Среднее значение
func TestAverageValue(t *testing.T) {
	testTable := []struct {
		numbers  [5]float64
		expected float64
	}{
		{
			numbers:  [5]float64{1, 2, 3, 4, 5},
			expected: 3.0,
		},
		{
			numbers:  [5]float64{2, 2, 2, 2, 2},
			expected: 2.0,
		},
		{
			numbers:  [5]float64{},
			expected: 0.0,
		},
	}
	for _, testCase := range testTable {
		res := AverageValue(testCase.numbers)
		assert.Equal(t, testCase.expected, res)
	}
}

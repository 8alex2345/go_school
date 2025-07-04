package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// ### **Задача 1: Подсчёт количества цифр в числе**
func TestQuantityNum(t *testing.T) {
	testTable := []struct {
		numbers  int
		expected map[int]int
	}{
		{
			numbers:  -123,
			expected: map[int]int{1: 1, 2: 1, 3: 1},
		},
		{
			numbers:  0,
			expected: map[int]int{0: 1},
		},
		{
			numbers:  111222,
			expected: map[int]int{1: 3, 2: 3},
		},
		{
			numbers:  1112225678,
			expected: map[int]int{1: 3, 2: 3, 5: 1, 6: 1, 7: 1, 8: 1},
		},
	}
	for _, testCase := range testTable {
		res := QuantityNum(testCase.numbers)
		assert.Equal(t, testCase.expected, res)
	}
}

// Задача 2: Подсчёт частоты символов в строке
func TestCounterSymbol(t *testing.T) {
	testTable := []struct {
		letters  string
		expected map[rune]int
	}{
		{
			letters:  "",
			expected: map[rune]int{},
		},
		{
			letters:  "你好你好",
			expected: map[rune]int{'你': 2, '好': 2},
		},
		{
			letters:  "a!a@",
			expected: map[rune]int{'a': 2, '!': 1, '@': 1},
		},
		{
			letters:  "Aa",
			expected: map[rune]int{'A': 1, 'a': 1},
		},
	}
	for _, testCase := range testTable {
		res := CounterSymbol(testCase.letters)
		assert.Equal(t, testCase.expected, res)
	}
}

// Задача 3: Инвертирование мапы
func TestInvertMap(t *testing.T) {
	testTable := []struct {
		myMap    map[string]string
		expented map[string]string
	}{
		{
			myMap:    map[string]string{},
			expented: map[string]string{},
		},
		{
			myMap:    map[string]string{"a": "x", "b": "x"},
			expented: map[string]string{"x": "b"},
		},
		{
			myMap:    map[string]string{"a": "a"},
			expented: map[string]string{"a": "a"},
		},
	}
	for _, testCase := range testTable {
		res := InvertMap(testCase.myMap)
		assert.Equal(t, testCase.expented, res)
	}

}

// Задача 4: Индексирование слов
func TestIndexWords(t *testing.T) {
	testTable := []struct {
		words    string
		expented map[string]int
	}{
		{
			words:    "",
			expented: map[string]int{},
		},
		{
			words:    "go go go",
			expented: map[string]int{"go": 0},
		},
		{
			words:    "Go go",
			expented: map[string]int{"Go": 0, "go": 1},
		},
		{
			words:    "a b a",
			expented: map[string]int{"a": 0, "b": 1},
		},
	}
	for _, testCase := range testTable {
		res := IndexWords(testCase.words)
		assert.Equal(t, testCase.expented, res)
	}
}

// Задача 5: Сравнение двух мап
func TestComparisoMap(t *testing.T) {
	testTable := []struct {
		map1     map[string]int
		map2     map[string]int
		expected bool
	}{
		{
			map1:     map[string]int{},
			map2:     map[string]int{},
			expected: true,
		},
		{
			map1:     map[string]int{"a": 1},
			map2:     map[string]int{"a": 2},
			expected: false,
		},
		{
			map1:     map[string]int{"a": 1, "b": 2},
			map2:     map[string]int{"b": 2, "a": 1},
			expected: true,
		},
		{
			map1:     map[string]int{"a": 1, "b": 2},
			map2:     map[string]int{"a": 1},
			expected: false,
		},
	}
	for _, testCase := range testTable {
		res := ComparisoMap(testCase.map1, testCase.map2)
		t.Logf("Входящее значение map1(%v), Входящее значение map2 (%v) реезультат %v\n", testCase.map1, testCase.map2, res)
		assert.Equal(t, testCase.expected, res)
	}
}

// Задача 6: Найти ключ по значению
func TestFindKey(t *testing.T) {
	testTable := []struct {
		myMap    map[string]int
		expected string
	}{
		{
			myMap:    map[string]int{},
			expected: "not found",
		},
		{
			myMap:    map[string]int{"x": 10, "y": 20, "z": 20},
			expected: "y",
		},
	}
	for _, testCasse := range testTable {
		res := FindKey(testCasse.myMap)
		assert.Equal(t, testCasse.expected, res)
	}
}

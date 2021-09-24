package main

import (
	"fmt"
	"os"
)

func main() {
	str := input()
	m := findDuplicate(str)
	for key,value := range *m {
		if value > 1 {
			fmt.Printf("%q", key)
		}
	}
}

func input() string {
	fmt.Printf("Введите слово для проверки: ")
	var str string
	_, err := fmt.Scanf("%s", &str)
	if err != nil {
		os.Exit(404)
	}
	return str
}

func findDuplicate(str string) *map[rune]int {
	m := make(map[rune]int)
	for _,char := range str {
		value,ok := m[char]
		if !ok {
			m[char] = 1
		} else {
			m[char] = value + 1
		}
	}
	return &m
}
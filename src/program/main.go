package main

import (
	"fmt"
	"path/filepath"
	"plugin"
)

func main() {
	a := 1
	b := 7

	// Нас интересуют все файлы в каталоге plugins, которые имеют расширение .so
	plugins, err := filepath.Glob("plugins/*.so")
	if err != nil {
		panic(err)
	}

	// Перебираем все найденные файлы, в поисках плагинов, которые нам подойдут
	for _, filename := range plugins {
		p, err := plugin.Open(filename)
		if err != nil {
			fmt.Println(err)
			continue
		}

		// Наш плагин должен содержать функцию или переменную Something
		symbol, err := p.Lookup("Something")
		if err != nil {
			fmt.Println(err)
			continue
		}

		// Функция Something принимает два параметра int и возвращает int
		pFunc := symbol.(func(int, int) int)
		// Общаемся к Something
		result := pFunc(a, b)
		// Выводим путь к файлу и результат
		fmt.Println(filename, result)
	}

	fmt.Println("Success")
}

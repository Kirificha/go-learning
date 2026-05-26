package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup

	//wg.Add() Обязательно до запуска горутин, иначе будет гонка данных
	//Функция может успеть запустить горутину и она создать Add, а может и не успеть
	//Получается неопределённость, не знаем исход кода. Никогда так не делаем!
	wg.Add(3)

	for i := 0; i < 3; i++ {
		go func() {
			fmt.Println("Я горутина номер: ", i)
			wg.Done()
		}()
	}

	wg.Wait()
	fmt.Println("main завершён")
	wrongVersion()
}

func wrongVersion() {
	var wg sync.WaitGroup

	for i := 0; i < 3; i++ {
		go func(n int) {
			wg.Add(1) // ОШИБКА: Add внутри горутины
			fmt.Println("Я горутина номер: ", n)
			wg.Done()
		}(i)
	}

	wg.Wait()
	fmt.Println("main завершён (неправильная версия)")
}

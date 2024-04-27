package main

import "fmt"

type Vertex struct {
	X int
	Y int
}

var (
	m  = Vertex{1, 4}  // Инициализируем m
	m1 = Vertex{X: 5}  // Y:0 является неявным
	m2 = Vertex{}      // X:0 и Y:0
	n  = &Vertex{1, 2} // имеет тип *Vertex
)

func main() {
	a, b := 50, 100

	p := &a         // point to a
	fmt.Println(*p) // read a through the pointer
	*p = 200        // set a through the pointer
	fmt.Println(a)  // see the new value of i

	p = &b                  // point to b
	*p = *p + 1000          // Sum b through the pointer
	fmt.Print(*p, "\n\n\n") // see the new value of b

	v := Vertex{1, 2}
	v.X = 5          // Доступ к структурным полям осуществляется
	fmt.Println(v.X) // при помощи точки

	l := &v   // Создание указателя на структуру v
	l.X = 1e3 // Изменение значения поля X через указатель
	fmt.Print(v, "\n\n\n")

	fmt.Print(m, n, m1, m2, "\n\n\n")

	var Arr [2]string           // Объявление массива элементов String
	Arr[0] = "Privet"           // Инициализация ячейки [0]
	Arr[1] = "Pasha"            // Инициализация ячейки [1]
	fmt.Println(Arr[0], Arr[1]) // Вывод массива Arr

	primes := [5]int{1, 2, 34, 5, 6} // Массив primes c 5 элементами int
	fmt.Print(primes, "\n\n\n")      // Вывод массива primes

}

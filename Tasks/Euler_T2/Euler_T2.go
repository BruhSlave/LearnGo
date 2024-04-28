//Каждый следующий элемент ряда Фибоначчи получается при сложении двух предыдущих.
//Начиная с 1 и 2, первые 10 элементов будут:
//1, 2, 3, 5, 8, 13, 21, 34, 55, 89, ...
//Найдите сумму всех четных элементов ряда Фибоначчи, которые не превышают четыре миллиона.

package main

import "fmt"

func main() {
	var inter, answer int
	fib1 := 1
	fib2 := 2

	for {
		inter = fib1 + fib2
		fib1 = fib2
		fib2 = inter

		if fib2%2 == 0 {
			answer += fib2
			if answer > 4e6 {
				fmt.Println(fib2)
				break
			}
		}
	}
}

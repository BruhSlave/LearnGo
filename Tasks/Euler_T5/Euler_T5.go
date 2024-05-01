/* 2520 - самое маленькое число, которое делится без остатка на все числа от 1 до 10.
Какое самое маленькое число делится нацело на все числа от 1 до 20?
*/

package main

import (
	"fmt"
)

func main() {
	var count int = 0
	for num := 2519; num > 2519; num++ {
		for i := 1; i < 20; i++ {
			if num%i == 0 {
				count++
				fmt.Println(count)
			}
		}
	}
	//fmt.Println(count)
}

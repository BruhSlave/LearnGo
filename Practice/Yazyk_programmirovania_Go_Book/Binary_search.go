package main

import "fmt"

// Бинарный поиск числа
func binary_search(list []int, item int) int {
	low := 0
	high := len(list) - 1

	for low <= high {
		mid := (low + high) / 2
		guess := list[mid]

		if guess == item {
			return mid
		}
		if guess > item {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	return -1 // Возвращаем -1, если элемент не найден
}

func main() {
	my_list := []int{1, 3, 5, 7, 9}

	fmt.Println(binary_search(my_list, 7)) // Должно выводить 4, так как 9 находится на позиции 4
}

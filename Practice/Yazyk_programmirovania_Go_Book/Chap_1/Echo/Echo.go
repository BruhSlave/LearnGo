// Echol выводит аргументы командной строки
package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	var s, sep, sep2 string
	for i, arg := range os.Args[0:] {
		// os.Args является переменной,
		// которая предоставляет доступ к аргументам командной строки
		s += sep + strconv.Itoa(i) + sep2 + arg
		sep = " "
		sep2 = "-"
		fmt.Println(i, "-", os.Args[i])
	}
	fmt.Print("\n", s)
	//fmt.Println(strings.Join(os.Args[0:1], ""))
}

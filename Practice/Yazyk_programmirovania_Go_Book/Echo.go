// Echol выводит аргументы командной строки
package main

import (
	"fmt"
	"os"
)

func main() {
	var s, sep string
	for i := 1; i < len(os.Args); i++ {
		// os.Args является переменной,
		//которая предоставляет доступ к аргументам командной строки
		s += sep + os.Args[i]
		sep = " "
	}
	fmt.Println(s)
}

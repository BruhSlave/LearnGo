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

//Slice length and capacity
func printSlice(ArrLenCap []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(ArrLenCap), cap(ArrLenCap), ArrLenCap)
	//Длина и емкость Slice могут быть получены с помощью выражений len(s) и cap(s).
}

//Creating a slice with make
func printMakeSlice(a string, b []int) {
	fmt.Printf("%s len=%d cap=%d %v\n", a, len(b), cap(b), b)
}

func main() {
	a, b := 50, 100

	p := &a         // point to a
	fmt.Println(*p) // read a through the pointer
	*p = 200        // set a through the pointer
	fmt.Println(a)  // see the new value of a

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

	Mass := [5]int{1, 5, 7, 13, 45} // Масиив Mass
	var g []int = Mass[2:5]         // Slice g равный элеменам массива Mass с 2 по 4
	fmt.Print(g, "\n\n\n")

	Sl1 := Mass[1:3]      // Slice не хранит данных
	Sl2 := Mass[2:4]      // Изменение элементов фрагмента
	fmt.Println(Sl1, Sl2) //изменяет соответствующие элементы базового массива.
	fmt.Println(Sl1, Sl2)
	Sl2[1] = 1000
	fmt.Print(Sl1, Sl2, Mass, "\n\n\n")

	SlLit1 := []int{1, 5, 8, 2} // Slice literal как массив, только без указания длины
	fmt.Println(SlLit1)
	SlLit2 := []string{"Pasha", "Misha", "Sasha"}
	fmt.Println(SlLit2)
	StSliLit := []struct { // Slice literal структуры
		r int
		y string
	}{
		{1, "Pasha"},
		{14, "Misha"},
	}
	fmt.Print(StSliLit, "\n\n\n")

	var ArrDef = []int{1, 2, 4, 6, 47}
	fmt.Println(ArrDef)
	SliceDef := ArrDef[:3] //Значение по умолчанию равно нулю для нижней границы
	SliceDef = ArrDef[2:]  // и длине Slice для верхней границы
	fmt.Print(SliceDef, "\n\n\n")

	ArrLenCap := []int{1, 2, 3, 4, 5, 6} //Slice имеет как длину, так и вместимость.
	printSlice(ArrLenCap)                // Длина Slice - это количество элементов, которые он содержит.
	ArrLenCap = ArrLenCap[:0]            // Емкость Slice - это количество элементов в базовом массиве, отсчитываемое от первого элемента в Slice.
	printSlice(ArrLenCap)                // Вы можете увеличить длину Slice, повторно нарезав его,
	ArrLenCap = ArrLenCap[:5]            //	при условии, что он обладает достаточной емкостью
	printSlice(ArrLenCap)
	ArrLenCap = ArrLenCap[1:4]
	printSlice(ArrLenCap)
	fmt.Print("\n\n\n")

	var ArrNill []int                                                     //Нулевое значение Slice равно нулю.
	fmt.Printf("len=%d cap=%d %v\n", len(ArrNill), cap(ArrNill), ArrNill) //Нулевой Slice имеет длину и емкость, равные 0, и не содержит базового массива.
	fmt.Print("\n\n\n")

	ArrMakeSl1 := make([]int, 5) // Slice можно создавать с помощью встроенной функции make;
	// Чтобы указать емкость, передайте третий аргумент для make
	printMakeSlice("ArrMakeSl1", ArrMakeSl1)
	ArrMakeSl2 := make([]int, 0, 5)
	printMakeSlice("ArrMakeSl2", ArrMakeSl2)
	ArrMakeSl3 := ArrMakeSl1[:2]
	printMakeSlice("ArrMakeSl3", ArrMakeSl3)
	ArrMakeSl4 := ArrMakeSl1[2:5]
	printMakeSlice("ArrMakeSl4", ArrMakeSl4)
	fmt.Print("\n\n\n")

	SliceAppend := append(ArrMakeSl1, 5, 6) // Slice добавляются новые элементы, поэтому в Go предусмотрена встроенная функция append
	fmt.Printf("len=%d cap=%d %v\n", len(SliceAppend), cap(SliceAppend), SliceAppend)
	//Результирующее значение append представляет собой фрагмент,
	// содержащий все элементы исходного фрагмента плюс предоставленные значения.

	/* Если резервный массив  слишком мал, чтобы вместить все заданные значения,
	будет выделен массив большего размера.
	Возвращаемый фрагмент будет указывать на вновь выделенный массив.
	*/
}

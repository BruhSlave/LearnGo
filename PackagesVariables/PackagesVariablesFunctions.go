package main

import (
	"fmt"
	"math"
	"math/cmplx"
	"math/rand"
	"time"
)

func add(x, y int) int {
	return x + y
}

func swap(a, b string) (string, string) {
	return b, a
}

func split(sum int) (x, y int) {
	x = sum*20 + 3
	y = sum*10 + 400
	return
}

var (
	ToBe   bool       = false
	MaxInt uint64     = 1<<64 - 1
	z      complex128 = cmplx.Sqrt(-5 + 12i)
)

const Big = 1451 << 10

func needFloat(x float64) float64 {
	return x
}

func main() {
	fmt.Println("Hello Pasha")
	fmt.Println("The time is  ", time.Now())
	fmt.Println("My favorite number is ", rand.Intn(3))
	fmt.Printf("Now you have %v problems. \n", math.Sqrt(7))

	fmt.Println(add(100, 50))
	fmt.Println(swap("Pasha\n\n", "Privet"))

	fmt.Println(split(5))

	var a, b, c int = 1, 45, 90
	fmt.Println(a, b, c)

	abc, bca := 1000, 2000
	fmt.Println(abc, bca)

	fmt.Printf("Type: %T Value: %v\n", ToBe, ToBe)
	fmt.Printf("Type: %T Value: %v\n", MaxInt, MaxInt)
	fmt.Printf("Type: %T Vlaue: %v\n", z, z)

	var x, y int = 10, 20
	var f float64 = math.Sqrt(float64(x*x + y*y))
	var z uint = uint(f)
	fmt.Println(x, y, z, f)

	fmt.Printf("Value is %v", needFloat(Big))

}

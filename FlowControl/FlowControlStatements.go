package main

//26.04.24
import (
	"fmt"
	"math"
	//"time"
)

func pow(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	} else {
		fmt.Printf("v (%g) >=  lim(%g) \n", v, lim)
	}
	return lim
}

func main() {
	var sum int = 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Printf("Type = %T Sum = %v \n", sum, sum)

	/*InfiniteLoopTire := "-"
	InfiniteLoopReshet := "#"
	var count int = 1
	for {
		fmt.Print(InfiniteLoopTire)
		fmt.Print(InfiniteLoopReshet)
		count++
		time.Sleep(100 * time.Millisecond) // функция задержки; 1/10 Секунды
		if count > 100 {
			break
		}
	}*/

	sum2 := 1
	for sum2 < 1000 {
		sum2 += sum2
	}
	fmt.Printf("Type = %T Sum2 = %v \n", sum2, sum2)

	fmt.Printf("Type = %T FuncPow = %v \n", pow(3, 3, 20), pow(3, 3, 20))
	fmt.Printf("Type = %T FuncPow = %v \n", pow(3, 1, 20), pow(3, 1, 20))
}

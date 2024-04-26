package main

//26.04.24
import (
	"fmt"
	"math"
	"time"
	//"time"
)

func pow(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	} else {
		//fmt.Printf("v (%g) >=  lim(%g) \n", v, lim)
	}
	return lim
}

// Exercise 1
const (
	Delta = 0.000001 // Порог остановки
)

func Sq(x float64) float64 {
	var z float64 = 1
	var zz float64
	for {
		zz = z - ((z*z - x) / (2 * z)) // Вычисление следующего приближения
		if math.Abs(zz-z) < Delta {    // Проверка, изменилось ли значение на достаточно малое
			break
		}
		z = zz
	}
	return z
}

func main() {
	var sum int = 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	//fmt.Printf("Type = %T Sum = %v \n", sum, sum)

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
	/*fmt.Printf("Type = %T Sum2 = %v \n", sum2, sum2)

	fmt.Printf("Type = %T FuncPow = %v \n", pow(3, 3, 20), pow(3, 3, 20))
	fmt.Printf("Type = %T FuncPow = %v \n", pow(3, 1, 20), pow(3, 1, 20))
	*/

	//Exercise 1

	//fmt.Printf(" 	Newton method = %f \n", Sq(2))
	//fmt.Printf(" 	Math.Sqrt = %f", math.Sqrt(2))

	fmt.Println("Когда Пятница?")
	today := time.Now().Weekday()
	switch time.Friday {
	case today + 0:
		fmt.Println("Сегодня")
	case today + 1:
		fmt.Println("Завтра")
	case today + 2:
		fmt.Println("Послезавтра")
	default:
		fmt.Println("Не скоро")
	}

	t := time.Now()
	switch {
	case t.Hour() <= 12:
		fmt.Println("Good Morning")
	case t.Hour() <= 17:
		fmt.Println("Good Afternoon")
	case t.Hour() >= 22:
		fmt.Println("Good Night")
	}

	defer fmt.Println("Privet")
	fmt.Print("Pasha")
	fmt.Println(" and Misha")

	fmt.Println("Counting")
	for i := 1; i < 15; i++ {
		defer fmt.Print(i, " - ")
	}
	fmt.Println("Done")
}

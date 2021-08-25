package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	fmt.Println("Switch in golang")
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("Linux")
	default:
		fmt.Printf("%s \n", os)
	}

	fmt.Println("-----------------")
	fmt.Println("Cuando es Sabado?")
	today := time.Now().Weekday()
	switch time.Saturday {
	case today + 0:
		fmt.Println("Hoy")
	case today + 1:
		fmt.Println("Mañana")
	case today + 2:
		fmt.Println("En dos dias")
	default:
		fmt.Println("Falta un monton")
	}

	fmt.Println("-----------------")
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("Buen dia")
	case t.Hour() < 17:
		fmt.Println("Buenas tardes")
	default:
		fmt.Println("Buenas noches")
	}
}

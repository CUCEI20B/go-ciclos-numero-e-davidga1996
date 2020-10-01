package main

import "fmt"

func main() {
	var e float64
	var exponencial, rev int

	fmt.Print()
	fmt.Scan(&rev)

	e = 1
	for i := 1; i <= rev; i++ {
		exponencial = 1
		for j := 1; j <= i; j++ {
			exponencial = exponencial * j
		}
		e = e + 1/float64(exponencial)
	}

	fmt.Println(e)
}

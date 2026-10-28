package main

import (
	"fmt"
	"math"
)

func main() {
	const TARGET = 10;

	fmt.Print("scrivi 5 valori tra 0 e 20: ");

	var vicino int;

	for i := 0; i < 5; i++ {
		var a int;
		fmt.Scan(&a);

		if(a < 0 || a > 20) {
			i--;
			continue;
		}

		if(math.Abs(float64(TARGET - a)) < math.Abs(float64(TARGET - vicino))) {
			vicino = a;
		}
	}
	fmt.Println(vicino);
}
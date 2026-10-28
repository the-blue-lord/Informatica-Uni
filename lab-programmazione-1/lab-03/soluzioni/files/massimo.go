package main

import (
	"fmt"
	"math/rand"
)

func main() {
	const (
		NUMERO_VALORI = 10;
		MIN = 10;
		MAX = 30
	);

	var max int;

	for i := 0; i < NUMERO_VALORI; i++ {
		r := rand.Intn(MAX-MIN+1)+MIN;
		fmt.Print(r, " ");
		if(r > max) {
			max = r;
		}
	}
	fmt.Println();
	fmt.Println(max);
}
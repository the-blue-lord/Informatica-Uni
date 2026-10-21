package main

import "fmt"

func main() {
	const INTERVALLO_DI_CONFIDENZA = 1e-6;

	var x, y, m, q float64;

	fmt.Print("x e y del punto P: ");
	fmt.Scan(&x, &y);
	fmt.Print("m e q della retta R: ");
	fmt.Scan(&m, &q);

	diff := y - m*x - q;

	fmt.Print("P ");

	if(-INTERVALLO_DI_CONFIDENZA < diff && diff < INTERVALLO_DI_CONFIDENZA) {
		fmt.Print("sulla retta");
	} else if (diff > 0) {
		fmt.Print("sopra");
	} else {
		fmt.Print("sotto");
	}

	fmt.Println(" R");
}
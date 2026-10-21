package main

import "fmt"

func main() {
	var x, y float64;

	fmt.Scan(&x, &y);

	if(x >= 0 && y >= 0) {
		fmt.Print("I");
	} else if(x < 0 && y >= 0) {
		fmt.Print("II");
	} else if(x < 0) {
		fmt.Print("III");
	} else {
		fmt.Print("IV");
	}

	fmt.Println(" quadrante");
}
package main

import "fmt"

func main() {
	var a, b float64;

	fmt.Scan(&a);
	fmt.Scan(&b);

	for(b != 0) {
		fmt.Println(b - a);

		a = b;

		fmt.Scan(&b);
	}
}
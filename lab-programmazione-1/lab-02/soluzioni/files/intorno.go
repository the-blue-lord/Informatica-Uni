package main

import ("fmt"; "math")

func main() {
	const EPSILON = 1e-5;

	var x, y float64;

	fmt.Scan(&x, &y);

	distanza := math.Sqrt(x*x+y*y);

	if(distanza < EPSILON) {
		fmt.Print("molto");
	} else {
		fmt.Print("non");
	}

	fmt.Println(" vicino all'origine");
}
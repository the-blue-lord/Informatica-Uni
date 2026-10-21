package main

import "fmt"

func main() {
	var x_min, x_max, x, y_min, y_max, y int;

	fmt.Print("P1: ");
	fmt.Scan(&x_min, &y_min);

	fmt.Print("P2: ");
	fmt.Scan(&x_max, &y_max);

	fmt.Print("P3: ");
	fmt.Scan(&x, &y);

	if(x_min > x_max) {
		x_min, x_max = x_max, x_min;
	}

	if(y_min > y_max) {
		y_min, y_max = y_max, y_min;
	}

	fmt.Print("P3 ");

	if(x_min < x && x < x_max && y_min < y && y < y_max) {
		fmt.Prinln("interno");
	} else if (x < x_min || x > x_max || y < y_min || y > y_max) {
		fmt.Prinln("esterno");
	} else {
		fmt.Prinln("perimetrale");
	}
}
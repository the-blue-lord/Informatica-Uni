package main

import "fmt"

func main() {
	var min, max int;

	fmt.Print("due int: ");
	fmt.Scan(&min, &max);

	if(min >= max) {
		min, max = max, min;
	}

	fmt.Println(max);
}

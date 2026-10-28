package main

import "fmt"

func main() {
	var a, b int;
	fmt.Scan(&a);
	if(a == -1) {
		fmt.Println();
		fmt.Println("somma: 0");
		return;
	}
	sum := a;
	fmt.Scan(&b);

	for(b != -1) {
		if(b >= a) {
			fmt.Print("+");
		} else {
			fmt.Print("-");
		}
		sum += b;
		a = b;
		fmt.Scan(&b);
	}

	fmt.Println();
	fmt.Println("somma:", sum);
}
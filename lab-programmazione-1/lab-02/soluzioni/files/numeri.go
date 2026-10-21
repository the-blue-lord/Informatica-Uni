package main

import "fmt"

func main() {
	var i int;
	fmt.Print("numero? ");
	fmt.Scan(&i);

	if(0 < i) {
		fmt.Println("positivo");
	}
	if(i % 10 == 0) {
		fmt.Println("divisibile per 10");
	}
}
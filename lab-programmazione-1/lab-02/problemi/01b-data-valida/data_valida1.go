package main

import "fmt"

func main() {
	var g,m,a int;

	fmt.Print("gg mm aaaa? ");
	fmt.Scan(&g, &m, &a);

	fmt.Print("data ");

	if(g <= 0 || m <= 0 || a <= 0 || g>31 || m>12) {
		fmt.Print("non ")
	}

	fmt.Println("valida");
}
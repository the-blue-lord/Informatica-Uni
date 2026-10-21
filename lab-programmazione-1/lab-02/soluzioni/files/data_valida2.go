package main

import "fmt"

func main() {
	var g,m,a int;

	fmt.Print("gg mm aaaa? ");
	fmt.Scan(&g, &m, &a);

	fmt.Print("data ");

	g_max := 31;

	if(m == 4 || m == 6 || m == 9 || m == 11) {
		g_max = 30;
	} else if(m == 2) {
		g_max = 28;
	}

	if(g <= 0 || m <= 0 || g>g_max || m>12) {
		fmt.Print("non ");
	}

	fmt.Println("valida");
}
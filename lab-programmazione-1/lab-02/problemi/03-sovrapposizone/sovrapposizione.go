package main

import "fmt"

func main() {
	var g1, g2, i1, i2, f1, f2 int;

	fmt.Print("appuntamento 1 (gg, start, end): ");
	fmt.Scan(&g1, &i1, &f1);
	fmt.Print("appuntamento 2 (gg, start, end): ");
	fmt.Scan(&g2, &i2, &f2);

	if(g1 != g2 || (i1 > i2 || i2 >= f1) && (i1 >= f2 || f2 > f1)) {
		fmt.Print("non ");
	}

	fmt.Println("si sovrappongono");
}
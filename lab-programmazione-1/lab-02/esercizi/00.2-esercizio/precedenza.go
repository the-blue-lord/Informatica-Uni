package main

import "fmt"

func main() {
	const VALORE_PER_TRUE = "t";

	var as, bs, cs string;
	fmt.Print("tre bool (t/f): ");
	fmt.Scan(&as, &bs, &cs);

	a, b, c := as == VALORE_PER_TRUE, bs == VALORE_PER_TRUE, cs == VALORE_PER_TRUE;

	fmt.Println("a || b && c =", a||b&&c);
	fmt.Println("(a || b) && c =", (a||b)&&c);
}

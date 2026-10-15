package main;

import "fmt";

func main() {
	var a, b int;

	fmt.Print("dividendo: ");
	fmt.Scan(&a);
	fmt.Print("divisore: ");
	fmt.Scan(&b);

	q := a/b;
	r := a%b;

	fmt.Println("quoziente =", q);
	fmt.Println("resto =", r);
}

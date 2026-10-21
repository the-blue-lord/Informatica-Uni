package main

import "fmt"

func main() {
	var n int;

	fmt.Print("un int: ");
	fmt.Scan(&n);

	fmt.Println("n < 10 || n > 20 =", n<10||n>20);
	fmt.Println("!(n >= 10 && n <= 20) =", !(n>=10&&n<=20));
}

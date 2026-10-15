package main;

import "fmt";

func main() {
	// Centigradi per fahrenheit = 5/9 -> fahrenheit per centigrado = 1/(5/9) = 9/5
	const FAHRENHEIT_PER_CENTIGRADO = 9./5.;

	var c, f float64;

	fmt.Print("temperatura in centigradi? ");

	fmt.Scan(&c);

	f = 32. + c * FAHRENHEIT_PER_CENTIGRADO;

	fmt.Println(c, "C =", f, "F");
}

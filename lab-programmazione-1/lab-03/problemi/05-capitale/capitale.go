package main

import "fmt"

func main() {
	var capitale, tasso, obbiettivo float64;

	fmt.Print("capitale iniziale(euro): ");
	fmt.Scan(&capitale);
	fmt.Print("tasso di interesse annuale (%) : ");
	fmt.Scan(&tasso);
	fmt.Print("obbiettivo(euro): ");
	fmt.Scan(&obbiettivo);

	var i int;
	for i = 0; capitale < obbiettivo; i++ {
		capitale *= (100 + tasso)/float64(100);
	}
	fmt.Println("Ci vogliono", i, "anni per arrivare a (o superare) l'obbiettivo");
}
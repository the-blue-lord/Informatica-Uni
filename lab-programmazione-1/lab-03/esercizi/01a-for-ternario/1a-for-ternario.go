package main
import "fmt"
func main() {
    /*
		Chiedi un inpu tintero all'utente e stampa su un'unica riga un numeor di asterischi pari al numero dato.
    */
	var n int;
	fmt.Print("un numero: ");
	fmt.Scan(&n);
	for i := 1; i <= n; i++ {
		fmt.Print("*");
	}
	fmt.Println();
}
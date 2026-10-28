package main
import "fmt"
func main() {
    /*
    	Chiedi all'utente un intero e stampa, separati da spazi, gli interi tra 0 e l'input compresi in ordine decrescente.
    */
	var n int;
	fmt.Print("un numero: ");
	fmt.Scan(&n);
	for i := n; i >= 0; i-- {
		fmt.Print(i, " ");
	}
	fmt.Println();
}
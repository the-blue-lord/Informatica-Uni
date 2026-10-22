package main
import "fmt"
func main() {
    /*
    	Il programma chiede un input intero all'utente e stampa separti da spazio tutti i numeri pari non negativi e non superiori all'intero dato.
    */
	var n int;
	fmt.Print("un numero: ");
	fmt.Scan(&n);
	for i := 0; i <= n; i = i + 2 {
		fmt.Print(i, " ");
	}
	fmt.Println();
}
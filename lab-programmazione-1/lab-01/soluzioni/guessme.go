// L'eseguibile oracolo chiede in input 3 interi, verifica l'input e notifica un eventuale errore di lettura (<nil> in caso di assenza di errore). Informa inoltre l'utente dell'effettivo numero di valori letti e ne stampa la terza parte della loro somma.

package main;

import "fmt";

func main() {
	var a, b, c, n int;

	fmt.Print("fornire tre interi: ");
	n, err := fmt.Scan(&a, &b, &c);

	fmt.Println("errore in lettura:", err);
	fmt.Println("valori letti:", n, "/3");

	m := float64(a + b + c)/3;

	fmt.Println(m);
}

package main

import (
	"fmt";
	"math/rand"
)

/* Scrivi un progoramma che scelga casualemente un intero tra 1 e 10 compresi, poi stampa su una riga l'intero estratto e su quella successiva un valore booleano che indiche se n é 1 oppure 10*/

func main() {
	n := rand.Intn(10) + 1;
	fmt.Println(n);
	fmt.Println(n == 1 || n == 10);
}

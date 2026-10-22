package main

import ("fmt"; "math/rand")

func main() {
    /*
    	Dato un massimo K = 20, continua a estrarre un numero tra 0 e K (inclusi), stampandoli in fila seprati da spazi, fino a quando non estrai 0. Concludi stampando il  numero di numeri estratti (escluso 0).
    */
    const K = 20;
    var n int;

    c := 0;
    for {
        n = rand.Intn(K+1);
        fmt.Print(n, " ");
        if n == 0 {
            break;
        }
        c++;
    }
	fmt.Println();
    fmt.Println(c);
}
package main

import ("fmt"; "math/rand")

func main() {
    /*
    	Siano 20 e 10 i valori chiamati rispettivamente target e massimo, allora continua a sommare numeri casuali compresi tra 0 e massimo (compresi) fino a quando non si raggiunge o si supera la cifra di target. A quel punto stampare il valore così ottenuto dalla serie di somme.
    */
    const TARGET = 20;
    const MAX = 10;
    var n int;

    target_guess := 0;
    for target_guess < TARGET {
        n = rand.Intn(MAX) + 1;
        fmt.Print(n, " ");
        target_guess += n;
    }
    fmt.Print("\n",target_guess,"\n");
}
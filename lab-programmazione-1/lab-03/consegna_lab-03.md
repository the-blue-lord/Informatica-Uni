# Lab03 - for

[statistiche consegne Lab02](Lab02_stats.md)

## Per gli esercizi di scrittura:
1. leggere attentamente le specifiche 
2. individuare le operazioni richieste / progettare la soluzione al problema utilizzando SOLO i costrutti di Go visti finora.
3. implementare la soluzione in Go, eventualmente copiando all'inizio del programma il testo dell'esercizio (tra /* e */).
4. compilare il programma (se non compila, correggere gli errori e ripetere questo passo)
5. (se compila) lanciare il programma su un esempio di input
6. individuare il dominio del problema, e quindi i casi di test su cui testare il programma
7. lanciare il programma sui casi di test, controllando che l'output prodotto sia quello richiesto

Nota: negli esempi di esecuzione, l'input fornito dall'utente è evidenziato in **grassetto**, il resto è output prodotto dal programma.

Raccomandazioni:
- non utilizzate “numeri magici”, ma dichiarate costanti (const)
- scrivete programmi leggibili scegliendo nomi sensati per gli identificatori e indentando correttamente
- rispettare alla lettera le specifiche (questo non vuol dire che non possiate poi scrivere altri programmi di vostra iniziativa)
- leggete anche le note e rispettate gli esempi di esecuzione
- quando viene indicato "il programma legge ... valori" si intende da standard input

## Per gli esercizi di analisi e comprensione del codice:
1. osservate attentamente il codice fornito e usatelo anche come esempio di riferimento per risolvere esercizi di scrittura per problemi simili
2. per comprendere cosa fa, potete sia lanciare l'esecuzione su qualche caso di input, sia tracciarlo passo passo con carta e penna, sempre su qualche input, sia cercare di comprendere cosa fa analizzandolo (e anche tutte e tre le cose).
3. Per avere degli esempi di *specifiche*, basta guardare gli esercizi di scrittura.

## Esercizi con i test (Go testing, non il `blackbox`)
Gli esercizi di cui esistono i test sono contrassegnati da `(T)`. Le istruzioni per lanciare i test, che valgono per ciascuno di quelli etichettati, sono:
- creare una directory chiamata `<programma>`
- inserire il proprio file `<programma>.go` nella directory `<programma>`
- copiare il file `<programma>_test.go` (scaricandolo da questo repo) nella directory `<programma>`
- copiare il file `lib_test.go` (scaricandolo da questo repo) nella directory `<programma>`
- creare il file `go.mod` col comando `go mod init <programma>` lanciato nella directory `<programma>`
- lanciare infine i test con `go build; go test -v` (sempre da dentro la dir `<programma>`)
- si prega di consegnare su upload sia il file `<programma>.go`, sia anche almeno un output dei test catturato tramite redirezione, il nome del file deve essere `<programma>.student-test`


---



## Nota - numeri casuali

Il pacchetto `"math/rand"` mette a disposizione funzioni per generare numeri (pseudo)casuali.
 Ad esempio la funzione 

`rand.Intn(n int)` 

restituisce un numero intero (pseudo)casuale nell'itervallo [0,n).

Ecco un esempio di codice in cui viene generato un numero casuale tra 0 e 10 (compresi) che viene poi utilizzato nel programma per sommarlo ad altri:

```go
var random, sum int
random = rand.Intn(11)
sum += random
```

Se voglio generare dei numeri interi casuali in un intervallo [a,b), con a != 0, ma con la stessa distribuzione uniforme, come posso fare?

Ad esempio, indicando i numeri su una retta, 
se invece che tra 0 e 12

```
|______________________|
0 1 2 3 4 5 6 7 8 9 0 1 2 3 4 5 6 7 8 9
+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-
```

voglio generare numeri casuali tra 3 e 15
```
      |______________________|
0 1 2 3 4 5 6 7 8 9 0 1 2 3 4 5 6 7 8 9
+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-
```
come faccio? Più avanti ci sarà un esercizio su questo. 

---

## Il for ternario

#### Schema di base
```go
for i := first; i <= last; i+= step {
    //do something (repeatedly)
}
```
#### Altri possibili schemi
- con step diverso da incremento di 1 (invece di `i++`: `i--`;  `i+=K`; `i-=K`; `i*=n`; ecc.)
- con elementi da non trattare (`if` nel blocco `for`)
- con un'espressione nella condizione (a sx dell'operatore di confronto, invece di un indice, o a dx)

---

### Esercizio 1a - for ternario - ripetizione semplice

Salvare il seguente programma in un file chiamato `1a-for-ternario.go`, aggiungere come commento le *specifiche* del programma (cioè le indicazioni che dareste al programmatore che dovesse scriverlo) e poi caricarlo su upload.
```go
package main
import "fmt"
func main() {
    /*
	 Scrivere qui le specifiche per questo programma e caricarlo su upload
    */
	var n int
	fmt.Print("un numero: ")
	fmt.Scan(&n)
	for i := 1; i <= n; i++ {
		fmt.Print("*")
	}
	fmt.Println()
}
```
nomefile: 1a-for-ternario.go

---

### Esercizio 1b - for ternario - ripetizione

Salvare il seguente programma in un file chiamato `1b-for-ternario.go`, fare quanto indicato nei commenti e poi caricarlo su upload.

```go
package main

import "fmt"

func main() {

	var first, last, step int = 1, 10, 1

	/* Qui sotto trovate quattro commenti (a., b., c., d.) e poi quattro cicli for.
	Per ciascun ciclo for scegliete il commento corrispondente a quello che fa il ciclo for e spostatelo (taglia e incolla) al posto della riga "//...". */

	//	a. produce 10 stampe del valore di first (1), separate da un tab
	//	b. stampa una riga vericale di asterischi
	//	c. stampa i numeri da 1 a 10, separati da tab
	//	d. stampa una riga orizzontale di asterischi

	// ciclo 1
	//...
	for i := first; i <= last; i += step {
		fmt.Print("*")
	}

	fmt.Println()

	// ciclo 2
	//...
	for i := first; i <= last; i += step {
		fmt.Println("*")
	}

	// ciclo 3
	//...
	for i := first; i <= last; i += step {
		fmt.Print(i, "\t")
	}

	fmt.Println()

	// ciclo 4
	//...
	for i := first; i <= last; i += step {
		fmt.Print(first, "\t")
	}

	fmt.Println()
}
```

nomefile: 1b-for-ternario.go

---

### Esercizio 1c - for ternario - il calcolo di un risultato come obiettivo

Salvare il seguente programma in un file chiamato `1c-for-ternario.go`, aggiungere come commento le *specifiche* del programma, sostituire il nome della variabile s con un nome più significativo e poi caricare il file su upload.
```go
package main

import "fmt"

func main() {
    /*
    Scrivere qui le specifiche per questo programma
    Sostituire il nome della variabile s con un nome più significativo
    Caricare il programma su upload
    */
    const K = 5
	var n int
	s := 0
	for i := 1; i <= K; i++ {
		fmt.Print("un numero: ")
		fmt.Scan(&n)
		s += n  // s = s + n
	}
	fmt.Println(s)
}
```
nomefile: 1c-for-ternario.go

---

### Esercizio 1d - for ternario

Salvare il seguente programma in un file chiamato `1d-for-ternario.go`, fare quanto indicato nei commenti e poi caricarlo su upload.

```go
package main

import "fmt"

func main() {
	/* Qui sotto trovate quattro commenti (a., b., c., d.) e poi quattro cicli for.
	   Per ciascun ciclo for scegliete il commento corrispondente a quello che fa il ciclo for e spostatelo (taglia e incolla) al posto della riga "//...". */

	// a. stampa i numeri tra 0 e n (escluso) espressi in base 3
	// b. stampa le potenze di 2 tra 1 e n (escluso)
	// c. stampa gli interi quadrati perfetti tra 0 e n (escluso)
	// d. calcola la somma dei primi n valori non negativi letti

	// ciclo 1
	//...
	var value, sum float64
	n := 5
	for i := 0; i < n; i++ {
		fmt.Scan(&value)
		if value < 0 {
			i--
		} else {
			sum += value
		}
	}
	fmt.Println(sum)

	// ciclo 2
	//...
	n = 10
	for i := 1; i < n; i *= 2 {
		fmt.Print(i, "\t")
	}
	fmt.Println()

	// ciclo 3
	//...
	const BASE = 3
	n = BASE * BASE
	for i := 0; i < n; i++ {
		fmt.Print(i, "\t", i/BASE, i%BASE, "\n")
	}
	fmt.Println()

	// ciclo 4
	//...
	n = 20
	for i := 0; i*i < n; i++ {
		fmt.Print(i*i, "\t")
	}
	fmt.Println()
}
```

nomefile: 1d-for-ternario.go

---
## Esercizi sul for ternario - ripetizione semplice

### Es. 2
Scrivere un programma `02_doppi.go` che legge K = 5 numeri e di ciascuno stampa il doppio (andando a capo ogni volta).

nomefile: 02_doppi.go

### Es. 3 - uso di `"math/rand"`
Completate il seguente programma (dopo aver copiato e incollato il codice in un file a cui date il nome `03_casuali.go`) e caricatelo su upload.

```go
package main

import (
	"fmt"
	"math/rand"
)

func main() {
	var lower, upper int = 0, 12
	fmt.Println("10 numeri casuali tra", lower, "e", upper)
	for i := 1; i <= 10; i++ {
		fmt.Print(rand.Intn(upper), " ")
	}
	fmt.Println()
	
	/*
	Completare la seguente parte del programma in modo che 
	generi e stampi 10 numeri casuali compresi in un intervallo definito dall'utente.
	La soluzione non deve consistere nel generare anche numeri esterni all'intervallo 
	ed eliminarli (o ignorarli).
	*/

	fmt.Println("Ti fornirò 10 numeri causali in un intervallo a tua scelta")
	fmt.Print("Dammi gli estremi inferiore e superiore dell'intervallo: ")
		//....
}
```
nomefile: 03_casuali.go

---
## Esercizi su for ternario - il calcolo di un risultato 


### Es. 4
Scrivere un programma `04_casuali_pari.go` che genera 10 numeri casuali in [0, 10], li stampa su una riga, ciascuno seguito da uno spazio, e poi a capo stampa quanti sono i valori pari generati, preceduto dall'etichetta "Numeri pari generati: ".

Esempio di esecuzione:

```
5 5 0 2 8 3 8 7 2 9 
Numeri pari generati: 5
```

nomefile: 04_casuali_pari.go

---



### Es. 5
Scrivere un programma `05_somma.go` che genera K = 10 numeri casuali in [1, 10], li stampa ciascuno su una riga, e poi a capo stampa la loro somma, preceduta dall'etichetta "la somma è ".

nomefile: 05_somma.go

---
## Esercizi su for ternario - ripetizione "parametrica"

### Es. 1prm-a-for-ternario.go

Salvare il seguente programma in un file chiamato `1prm-a-for-ternario.go`, aggiungere come commento le *specifiche* del programma, poi caricare il file su upload.

```go
package main
import "fmt"
func main() {
    /*
    Scrivere qui le specifiche per questo programma e caricare il file su upload
    */
	var n int
	fmt.Print("un numero: ")
	fmt.Scan(&n)
	for i := 0; i <= n; i = i + 2 {
		fmt.Print(i, " ")
	}
	fmt.Println()
}
```
nomefile: 1prm-a-for-ternario.go

---
### Es. 1prm-b-for-ternario.go

Salvare il seguente programma in un file chiamato `1prm-b-for-ternario.go`, aggiungere come commento le *specifiche* del programma e caricare il file su upload.

```go
package main
import "fmt"
func main() {
    /*
    Scrivere qui le specifiche per questo programma e caricare il file su upload
    */
	var n int
	fmt.Print("un numero: ")
	fmt.Scan(&n)
	for i := n; i >= 0; i-- {
		fmt.Print(i, " ")
	}
	fmt.Println()
}
```
nomefile: 1prm-b-for-ternario.go

---

### Es. 6
Scrivere un programma `06_stampaMultipli.go` che legge tre numeri *n*, *p*, *q* e stampa su una riga, separati da tab ("\t"), i numeri tra 1 e *n* (compresi) che sono multipli di *p* ma non di *q* (ad esempio multipli di 2 ma non di 5).

nomefile: 06_stampaMultipli.go


### Es. 7
Scrivere un programma `07_tabellina.go` che chiede ("un numero: ") e legge un numero *n*, quindi stampa la tabellina di *n* (cioè `n*1, n*2, ..., n*10`), su una riga, stampando un tab ("\t") dopo ogni valore e un a capo alla fine.

nomefile: 07_tabellina.go

## Esercizi su for ternario - espressione nella condizione

### Es. 1exp-for-ternario.go

Salvare il seguente programma in un file chiamato `1exp-for-ternario.go`, aggiungere come commento le *specifiche* del programma, poi caricare il file su upload.

```go
package main

import "fmt"

func main() {
	/*
	   Scrivere qui le specifiche per questo programma e caricare il file su upload
	*/

	const BASE = 8
	var n int

	fmt.Print("un int tra 1 e 4: ")
	fmt.Scan(&n)

	for i := 0; i < n*BASE; i++ { // i < n*BASE
		fmt.Print("|\t")
	}
	fmt.Println()

	for i := 0; i < n*BASE; i++ {
		fmt.Print(i/BASE, i%BASE, "\t")
	}
	fmt.Println()
}

```
nomefile: 1exp-for-ternario.go

---

### Es. 8
Scrivere un programma `08_quadrati.go` che legge un numero *n* positivo (senza stampare messaggi di richiesta) e stampa, uno per riga, i numeri tra 1 e *n* che sono dei quadrati perfetti, cioè numeri che possono essere espressi come il quadrato di un altro numero intero.\
Trovate una soluzione che non utilizzi la funzione math.Sqrt.\
Fate in modo che il ciclo termini subito dopo aver stampato l'ultimo quadrato, senza continuare a iterare a vuoto: potete usare un break o un'opportuna condizione del for.

nomefile: 08_quadrati.go 

---
## Il for unario
```go
for condition { //while condition is true
    //do something (repeatedly)
}
```
---
### Es. 2a-for-unario
Salvare il seguente programma in un file chiamato `2a-for-unario`, aggiungere come commento le *specifiche* del programma, sostituire il nome della variabile `t` con un nome più significativo e poi caricare il file su upload.

```go
package main
import (
	"fmt"
	"math/rand"
)
func main() {
    /*
    Scrivere qui le specifiche per questo programma, dare un nome più significativo alla variabile t e caricare il file su upload
    */
    const TARGET = 20
    const MAX = 10
    var n int

    t := 0
    for t < TARGET {
        n = rand.Intn(MAX) + 1
        fmt.Print(n, " ")
        t += n
    }
    fmt.Print("\n",t,"\n")
}
```
nomefile: 2a-for-unario.go

### Es. 2b-for-unario
Salvare il seguente programma in un file chiamato `2b-for-unario.go`, aggiungere come commento le *specifiche* del programma, dare nomi più significativi alle variabili c, n, o, e poi caricare il file su upload.

(mancava, aggiunto dopo la lezione)
```go
package main

import "fmt"

func main() {
/*
    Scrivere qui le specifiche per questo programma, dare nomi più significativi alle variabili c, n, o, e poi caricare il file su upload
 */
	var c, n, o int
	fmt.Scan(&n)
	for c != 100 {
		o = n
		fmt.Scan(&n)
		c = (o%100)*100 + n%100
		fmt.Println(c)
	}
}
```
nomefile: 2b-for-unario.go

## Esercizi sul for unario

### Es. 9

Scrivere un programma `09_targetSum.go` che genera numeri casuali tra 1 e 10, stampandoli su una riga ciascuno seguito da uno spazio, fino a quando la loro somma raggiunge un obiettivo fissato nel programma (TARGET), ad esempio 50. Poi stampa il numero di valori generati, la loro somma e la loro media.

Esempio di esecuzione:

```
10 1 10 9 8 7 10 
7 numeri casuali
somma: 55
media: 7.857142857142857
```

nomefile: 09_targetSum.go

---

## Il for zerario
```go
for { // o for ; ; { o for true {
    blocco di istruzioni
    if <condizione>{
        break
    }
    altro blocco di istruzioni
}
```
---
## Esercizi su for zerario

### Es 3-for-zerario
Salvare il seguente programma in un file chiamato `3-for-zerario.go`, aggiungere come commento le *specifiche* del programma, sostituire il nome della variabile `c` con un nome più significativo e poi caricare il file su upload.
```go
package main
import (
    "fmt"
    "math/rand"
)
func main() {
    /*
    Scrivere qui le specifiche per questo programma e caricare il file su upload
    */
    const K = 20
    var n int

    c := 0
    for {
        n = rand.Intn(K+1)
        fmt.Print(n, " ")
        if n == 0 {
            break
        }
        c++
    }
	fmt.Println()
    fmt.Println(c)
}
```

nomefile: 3-for-zerario.go

---
### Es. 10
Scrivere un programma `10_orarioValido.go` che stampa ripetutamente "ore e minuti:" per chiedere due valori h e min fino a ottenere due valori validi, cioè h in [0,23] e min in [0,59], poi stampa "orario valido"

nomefile: 10_orarioValido.go

---
## Frammenti di codice

Associa a ciascuna porzione di codice la descrizione corrispondente: copia le descrizioni riportate sotto in un file frammenti.txt, sostituendo per ciascuna la x con il numero della porzione di codice corrispondente. Carica il file su upload.

Frammenti di codice

![frammenti](Lab03/frammenti-di-codice.png)

x) conteggio di valori con una data proprietà\
x) calcolo di una somma in una serie di numeri\
x) ricerca del min/max in una serie\
x) ripetizione di un'azione\
x) ricerca della posizione di un valore dato in una serie\
x) elaborazione su valori consecutivi di una serie

---

nomefile: frammenti.txt

---
## Esercizi sul for (ternario, unario, zerario)

### Euclide (T)

Scrivere un programma `euclide.go` che legge da standard input due interi a e b, calcola il MCD tra i due numeri usando l'algoritmo di Euclide e lo stampa (seguito da `newline`), senza stampare altro.

Algoritmo di Euclide:

Dati due numeri naturali non negativi a e b:

0. se b > a, scambia a e b
1. controlla se b è zero. 
2. se lo è, a è il MCD. 
3. se non lo è, assegna ad r il resto della divisione a / b, 
   poni a = b e b = r e ripeti da 1.

nomefile: euclide.go


### Massimo

Scrivere un programma `massimo.go` che genera 10 numeri interi casuali tra 10 e 30, li stampa su una riga, separati da uno spazio, e stampa su una nuova riga "il massimo è " seguito dal massimo valore generato. 

nomefile: massimo.go

### Pari

Scrivere un programma `cifra_pari.go` che chiede e legge un intero non negativo da standard input, determina e stampa in che posizione (procedendo da destra a sinistra) si trova la prima cifra pari del numero. Se il numero non contiene cifre pari, il programma stampa che si trova in posizione -1.

Esempio di esecuzione:

un numero intero positivo: **62937777**\
la prima cifra pari (2) è in posizione 7

nomefile: cifra_pari.go


### Differenze (T)

Scrivere un programma `differenze.go` che legge (senza stampare richieste) una serie di valori float64 da tastiera e stampa le differenze, cioè la differenza tra il secondo e il primo, tra il terzo e il secondo, e così via, una per riga (quindi interponendo un `newline`).
Il programma termina quando riceve in input il valore 0 (il programma non deve calcolare la differenza tra 0 e il valore letto precedentemente).

*Domande*: 
- Quanti valori avrà letto il programma quando stampa la prima differenza? E quando stampa la seconda? 
- Quante istruzioni di lettura mi serviranno nel programma?
- Dove la/le devo mettere?

nomefile: differenze.go

### Capitale

Scrivere un programma `capitale.go` che chiede e legge da stardard input tre valori float64 che rappresentano un capitale iniziale (es. 1000.50 euro), un tasso di interesse composto annuale (es. 1.3%) e un obiettivo (es. 2000 euro), e stampa quanti anni occorrono per arrivare a (o superare) l'obiettivo.

Esempio di esecuzione:

capitale iniziale(euro): **456.25**\
tasso di interesse annuale(%): **3.2**\
obiettivo(euro): **700**\
Ci vogliono 14 anni per arrivare a (o superare) l'obiettivo

nomefile: capitale.go

## Esercizi supplementari
Sono per chi è veloce e come materiale in più per prepararsi all'esame.

### Max quadrato

Scrivere un programma `maxQuadrato.go` che, senza stampare richieste, legge da standard input un intero *n* positivo e stampa, utilizzando solo variabili e operazioni sul tipo int, il massimo quadrato (*k*^2) <= *n*.

Per calcolare il quadrato di un numero *n* usate *n*`*`*n*.

nomefile: maxQuadrato.go

### Indovina

Scrivere un programma `indovina_numero.go` che chiede all'utente ("indovina il numero: ") di indovinare un numero intero random *n* tra 1 e MAX, (dove MAX è una costante definita nel programma) e ripete la richiesta fino a che l'utente indovina, oppure fino a un massimo di MAX/2 tentativi.

Se l'utente indovina, il programma stampa il numero di tentativi fatti ("hai indovinato! N. di tentativi: tot", dove tot è il numero di tentativi), altrimenti stampa il messaggio "hai perso, il numero era n", dove n era il numero da indovinare. 

Se il numero digitato dall'utente è fuori dall'intervallo `[1,MAX]`, il tentativo non viene conteggiato e viene visualizzato il messaggio "fuori intervallo!", senza interrompere l'esecuzione. 

Utilizzare la funzione rand.Intn del package "math/rand" per fissare il numero da indovinare.

nomefile: indovina_numero.go


### Vicino

Scrivere un programma `vicino.go` che richiede ("scrivi 5 valori tra 0 e 20: ") e legge da standard input una serie di 5 valori interi tra 0 e 20 (ignorando eventuali valori fuori da questo intervallo), quindi stampa il valore piu' vicino a TARGET, tra quelli letti, dove TARGET è una costante definita nel programma. 
Se i valori  piu' vicini a TARGET sono due (uno minore e uno maggiore), stampa il primo letto tra i due.

Nota. Il package `math` ha una funzione `Abs` che calcola (e restituisce) il valore assoluto.

nomefile: vicino.go

### Andamento (T)

Scrivere un programma `andamento.go` che legge da tastiera, senza stampare richieste, una serie (di almeno un elemento) di numeri interi fino a incontrare il valore -1 e stampa su una riga "+" ogni volta che il nuovo valore (-1 escluso) è maggiore o uguale al precedente e "-" altrimenti; quindi stampa su una nuova riga la somma di tutti i numeri letti (-1 escluso).


#### Esempi di esecuzione
La riga con il prompt `$` fornisce l'input (echo 2 4 7 3 9 1 5 -1 |) all'eseguibile (./andamento), le righe seguenti sono l'output del programma; si presume il programma sia stato compilato.

```
$ echo 2 4 7 3 9 1 5 -1 | ./andamento
++-+-+
somma: 31
```

```
$ echo 20 10 20 -1 | ./andamento
-+
somma: 50
```

nomefile: andamento.go


### Primo

Scrivere un programma `primo.go` che chiede e legge un numero intero da standard input, stampa "è primo" se il numero è primo, "non è primo" altrimenti.

Suggerimento: occorre determinare il primo numero che è un divisore (se c'è).

Domanda: dato in input n, quante iterazioni dovrò fare al massimo?

Esempio di esecuzione:

un numero: **8501**\
8501 è primo

nomefile: primo.go

### Somma cifre

Scrivere un programma `somma_cifre.go` che stampa la somma delle cifre di un numero intero positivo chiesto e fornito da standard input.

Esempio di esecuzione:

un numero intero positivo: **8372395842369**\
la somma delle cifre di 8372395842369 è 69

nomefile: somma_cifre.go


### bin to ten

Scrivere un programma `bin2ten.go` che converte un intero composto di soli 0 e 1 nel valore corrispondente al numero binario rappresentato (es. 101 -> 5) e lo stampi. Nel caso il numero contenga altre cifre, il programma stampa il messaggio di errore "dato non valido".

Esempi di esecuzione:

un numero binario: **11001110110111**\
13239

un numero binario: 121
dato non valido


nomefile: bin2ten.go

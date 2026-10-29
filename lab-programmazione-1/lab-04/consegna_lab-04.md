# Lab 04 - byte, rune, stringhe, for range, cicli annidati

[statistiche consegne Lab03](Lab03_stats.md)

## Problemi di base ricorrenti che richiedono l'iterazione

#### Trovate qui degli [appunti e schemi di soluzioni per i problemi più comuni che richiedono l'uso l'iterazione](Lab04/Problemi-e-schemi-per-l-iterazione.md).


<!--
## bbtest

Ma mano che completate gli esercizi di scrittura, potete controllare che il formato dell'output sia corretto (e che l'output sia corretto sugli input forniti nei bbtest) con lo script  `controlla-consegne.sh` che trovate in  [strumenti](strumenti) (va reso eseguibile con il comando `chmod u+x`).

#### Predisposizione per usare questo script:
- Creare una propria dir di lavoro (la chiamerò prog1, ma potete chiamarla come volete; gli altri nomi di file e dir invece sono fissati) 
- In prog1 salvare e rendere eseguibile (chmod u+x) il file `controlla-consegne.sh` 
- In prog1 salvare il file con le indicazioni sugli esercizi (`Lab<numLab>.md`)
- In prog1 creare una dir `Lab<numLab>` (due cifre)
- In `Lab<numLab>` salvare la dir `bbtest` di quel lab 
- In `Lab<numLab>` creare una sottodir `consegne`
- In `consegne` salvare le proprie soluzioni (i file .go)

A questo punto potete lanciare il comando `controlla-consegne.sh` seguito da  \<numLab\> (due cifre)

#### Struttura di file e dir occorrenti per eseguire lo script (esempio su Lab02)
```
|--controlla-consegne.sh
|--Lab02
|  |--bbtest
|  |  |--myfile.input
|  |  |--myfile.expected
|  |--consegne
|  |  |--myfile.go
|--Lab02.md
```
-->

## Esercizi con i test (Go testing, non il `blackbox`)
Gli esercizi di cui esistono i test sono contrassegnati da `(T)`. Le istruzioni per lanciare i test, che valgono per ciascuno di quelli etichettati, sono:
- creare una directory chiamata `<programma>`
- inserire il proprio file `<programma>.go` nella directory `<programma>`
- copiare il file `<programma>_test.go` (scaricandolo da questo repo) nella directory `<programma>`
- copiare il file `lib_test.go` (scaricandolo da questo repo) nella directory `<programma>`
- creare il file `go.mod` col comando `go mod init <programma>` lanciato nella directory `<programma>`
- lanciare infine i test con `go build; go test -v` (sempre da dentro la dir `<programma>`)
- si prega di consegnare su upload sia il file `<programma>.go`, sia anche almeno un output dei test catturato tramite redirezione, il nome del file deve essere `<programma>.student-test`

### INDICAZIONI per scaricare e utilizzare i test nel file .tar:
- scaricare il file [tar](Lab04/SimilTema.tar) di questo Lab
- scompattarlo seguendo le istruzioni presenti nel [README-studenti](README-studenti.md)
- si otterranno le directory già create con i .go di test
- per ogni esercizio, consegnare su upload sia il file `<programma>.go`, sia anche almeno un output dei test catturato tramite redirezione, il nome del file deve essere `<programma>.student-test`


## Nota - lettura di caratteri
Per leggere un carattere (byte o runa) e salvarlo in una variabile `carattere`, usare la funzione Scanf di `fmt`:
- `fmt.Scanf("%c", &carattere)`

---

# Esercizi Go

Svolgere tutti gli esercizi, sia quelli di analisi e comprensione del codice sia quelli di scrittura, nell'ordine in cui sono proposti.

## Esercizi 0 e 1 di esempio svolti in classe dal docente
(non sono da caricare su upload)

### Es. 0
Scrivere un programma `es0.go` che fa le seguenti cose:
- chiede ("un carattere: ") e legge un byte
- lo stampa
- stampa il byte precedente, il byte stesso, e quello successivo in ordine lessicografico (ASCII), separati da uno spazio. Ad esempio, con 'd' come input, deve stampare: `c d e`
- stabilisce se è una lettera tra A e L, o altro (stampa "A-L" o "altro")
- poi chiede ("una parola: ") e legge una stringa e la stampa in verticale, una runa per riga. 

Ad esempio "città" diventa:
```
c
i
t
t
à
```



### Es 1
Scrivere un programma `pos_alfabeto.go` che
- legge una sequenza di caratteri ASCII (byte) terminati da '.' (es: abc!54LMN.),
- per ciascun carattere letto (tranne '.'), 
    - stabilisce se è una lettera latina **minuscola** (tra 'a' e 'z'), una **cifra** o **altro**
   		- se è una lettera latina minuscola, stabilisce anche la sua posizione nell'alfabeto
    	- se è una cifra, stabilisce il suo valore numerico
-  stampa un messaggio costituito dal carattere letto seguito da:
    - se minuscola, la sua posizione nell'alfabeto (es "f è la 6^a");
    - se cifra, il suo valore
    - altrimenti " - altro" (ad esempio "N - altro")
- quando termina, stampa la somma delle cifre e  "bye"

Esempio di esecuzione

`una sequenza di caratteri terminata da '.': `**abc!54LMN.**
```
a è la 1^a
b è la 2^a
c è la 3^a
! - altro
5 - 5
4 - 4
L - altro
M - altro
N - altro
somma: 9
bye
```

nomefile: pos_alfabeto.go

---


## Esercizi di analisi e comprensione

I testi dei prossimi esercizi si trovano a questo [link](https://forms.gle/aodQYeheiCxGjJLD7).

Analizzate il codice degli esercizi proposti e rispondete alle domande del form.


## Esercizi di scrittura guidata

I testi dei prossimi esercizi si trovano a questo [link](https://forms.gle/MmuNMP4YLZFaGzc39).

Le domande che accompagnano ciascun esercizio servono per progettare la soluzione del problema posto prima di implementarla in Go. 
Prima di scrivere ciascun programma, rispondete alle domande e poi usate le vostre risposte per progettare il programma richiesto. I programmi sono poi da caricare sul sito di upload

#### Lunghezza totale
<!---Scrivere un programma `lunghezza_tot.go` che legge da standard input un int totLen e una sequenza di stringhe (una per riga) e ne somma le lunghezze fino a raggiungere (o superare) totLen.

Quando termina, stampa la somma delle lunghezze e la concatenazione delle stringhe lette.--->

nomefile: lunghezza_tot.go

#### Chars
<!---Scrivere un programma `max_char.go` che legge da standard input una sequenza di 5 caratteri ASCII (byte) e stampa il maggiore in ordine lessicografico (cioè con il codice ASCII più alto).--->

nomefile: max_char.go

#### Trova
<!---Scrivere un programma `trova.go` che legge da standard input un carattere (runa) e una stringa e stampa la posizione del carattere nella stringa, o -1 se il carattere non è presente.--->

nomefile: trova.go

#### Stringa di byte
<!---Scrivere un programma stringaByte.go che legge da standard input una stringa e la analizza byte per byte: stampa "+" quando il carattere considerato è maggiore (ha codice Ascii maggiore) del precedente, stampa "-" quando è minore e stampa "=" quando è uguale.--->

nomefile: stringaByte.go

#### Stringa di rune
<!---Scrivere un programma stringaRune.go che legge da standard input una stringa e la analizza runa per runa: stampa "+" quando il carattere considerato è maggiore (ha codice Unicode maggiore) del precedente, stampa "-" quando è minore e stampa "=" quando è uguale.--->

nomefile: stringaRune.go

---

## Esercizi di scrittura
(da caricare su upload)

**NOTA**: quando di un esercizio non sono mostrati esempi di esecuzione e nelle specifiche non è richiesta esplicitamente la stampa di messaggi per la richiesta dei dati in input, scrivete un programma in cui non sono presenti tali messaggi.

### Crescente
Scrivere un programma `crescente.go` che legge da standard input una stringa di sole lettere (ASCII) minuscole e la stampa inserendo un '-' ogni volta che una lettera viene prima in ordine alfabetico della lettera precedente.

Per esempio, data in input la parola `ambire`, il programma stampa
`am-bir-e`

nomefile: crescente.go

### Slash
Scrivere un programma `disegna_slash.go` che legge un intero positivo *n* e stampa un backslash (\\) di altezza *n* composto da asterischi.

Esempio di esecuzione:

`dimensione \:` **3**
```
*
 *
  *
```

nomefile: disegna_slash.go

### V (T)
Scrivere un programma `disegna_v.go` che legge un intero positivo *n* e stampa una V di altezza (vertice compreso) *n* di asterischi. Il programma **non** deve stampare spazi a destra della V (selezionate con il mouse l'esempio qui sotto per vedere come è composta la figura)

Esempio di esecuzione:
```
*     *
 *   *
  * *
   *
```

nomefile: disegna_v.go

### Fibonacci
Scrivere un programma `fibonacci.go` che legge un intero positivo *n* e stampa i numeri di fibonacci dal primo all'*n*-esimo, rappresentandoli come righe di asterischi, ciascuna lunga quanto il numero da rappresentare.

Esempio di esecuzione:

`un numero:` **6**\
`*`\
`*`\
`**`\
`***`\
`*****`\
`********`


**Nota.** I primi due numeri della serie di Fibonacci sono 1 e 1, dal terzo in poi sono ciascuno la somma dei due numeri che lo precedono nella serie. Quindi i primi numeri della serie sono: 1, 1, 2, 3, 5, 8, 13, 21, ...

nomefile: fibonacci.go

### MinuMaiu
Scrivere un programma `minu_maiu.go` che legge da standard input una stringa (di rune) e stabilisce se la stringa contiene solo minuscole, solo maiuscole o sia minuscole che maiuscole o né maiuscole né minuscole, quindi stampa un messaggio opportuno ("solo minuscole", "solo minuscole", "sia minuscole che maiuscole", "").
Utilizzare le funzioni del pacchetto unicode.

nomefile: minu_maiu.go

---

## Composizione
Una volta identificati i sottoproblemi in un problema, occorre che le soluzioni dei sottoproblemi vengano composte insieme nel programma.

### Esempi di composizione

#### Concatenazione
Quando un sottoproblema non può essere risolto senza i risultati prodotti dalla soluzione di un altro sottoproblema, le loro soluzioni devono essere concatenate una dopo l'altra.

```go
/*
Specifiche: dato un numero intero in forma di stringa, 
stampare le cifre che sono maggiori 
della media delle cifre del numero.
*/

package main

import "fmt"

func main() {
	var numero string
	fmt.Scan(&numero)
	dim := len(numero)

      //soluzione di due sottoproblemi (media e  > media)
      //con due for separati, uno dopo l'altro

      //calcolo della media
	var sum byte
	for i := 0; i < dim; i++ {
		n := numero[i] - '0'
		sum += n
	}
	media := sum/byte(dim) + '0'
	fmt.Println("media", string(media))

      //stampa delle cifre maggiori della media
	for i := 0; i < dim; i++ {
		if numero[i] > media {
			fmt.Println(string(numero[i]))
		}
	}
}
```

#### Fusione

Quando le soluzioni di sottoproblemi presentano alcune parti comuni che possono essere svolte insieme (es. max + min; somma + conteggio, ...), le soluzioni possono essere fuse insieme.
```go
/*
Specifiche: dato un numero intero >=0,
stampare la somma delle cifre e la cifra più grande
*/

package main

import "fmt"

func main() {
	var num, max int
	fmt.Print("un int >=0: ")
	fmt.Scan(&num)

	//soluzione di due sottoproblemi (somma e max)
	//con un unico for
	sum := 0
	for num > 0 {
		cifra := num % 10
		sum += cifra     //calcolo della somma
		if cifra > max { //determinazione del max
			max = cifra
		}
		num /= 10
	}
	fmt.Println("somma:", sum)
	fmt.Println("max:", max)
}

```

#### Annidamento

Quando la soluzione di un problema che richiede un ciclo ha come sottoproblema a suo volta un problema che richiede un ciclo, le soluzioni vanno composte annidando i rispettivi cicli.
```go
/*
Specifiche: data una serie di 10 numeri interi,
stampare il primo numero la cui somma delle cifre è pari
*/

package main

import "fmt"

func main() {
	const DIM = 10
	var num int

	//soluzione di due sottoproblemi
	//(calcolo della somma delle cifre di un numero e ricerca del primo numero con tale somma pari)
	//con due for, uno annidato nell'altro

	for i := 0; i < DIM; i++ { //for esterno per la ricerca sugli input
		fmt.Print("next int: ")
		fmt.Scan(&num)
		sum := 0    //riazzero la somma per ogni nuovo numero
		n := num    //faccio una copia del numero
		for n > 0 { //for interno per la somma
			cifra := n % 10
			sum += cifra
			n /= 10
		}
		//if per interrompere la ricerca (for esterno)
		if sum%2 == 0 { //verifica se la somma è pari
			fmt.Println(num)
			break
		}
	}
}
```
---

### Max somma cifre (T)
Si vuole scrivere un programma `max_somma_cifre.go` che 
- legge da standard input una serie di numeri >= 0, terminata da 999, 
- trova il numero (escludendo 999) la cui somma delle cifre è la maggiore 
- e stampa tale somma.

Copiare il seguente codice in un file  `max_somma_cifre.go` e completarlo in modo che rispetti le specifiche date qui sopra. Ricordarsi di compilarlo e testarlo prima di caricalo su upload.

Si noti che il programma contiene due for annidati, uno servirà per determinare un massimo e uno per calcolare una somma.

```go
func main() {
	var n int

	maxSum := 0
	for {
		fmt.Scan(&n)
		if n == 999 {
			break
		}
		...
		for ... {
			...
		}
		if ... {
			...
		}
	}
	fmt.Println(maxSum)
}
```

nomefile: max_somma_cifre.go

---

### Numero di occorrenze del massimo
Scrivere un programma `num_max.go` che legge una sequenza di 10 interi positivi e stampa, sulla stessa riga, separati da uno spazio, il massimo intero letto e quante volte tale massimo compare nella sequenza.

E` possibile risolvere il problema senza tenere in memoria la sequenza di interi. Che tipo di composizione occorrerà per determinare insieme il calcolo del massimo e il conteggio delle occorrenze di tale massimo?

nomefile: num_max.go

#### Parsons

##### Parsons puzzle per `num_max.go`
In alternativa, o in aggiunta, prova a risolvere il seguente Parsons puzzle e confronta poi la tua soluzione con quella che costruisci qui.

Parsons puzzle: Copiare il codice qui sotto in un file `num_maxP.go`, inserire nel primo blocco e nell'ordine corretto le istruzioni che si trovano nel secondo blocco. Compilare e testare il programma prima di caricalo su upload.


Completare il programma seguente inserendo, nell’ordine corretto, le istruzioni che si trovano più sotto. Chiamare il programma `num_maxP.go`, compilarlo e testarlo, poi caricarlo su upload. 


Si noti che la soluzione che risulterà, invece di trovare prima il massimo e poi contare quante volte compare, utilizza un solo ciclo for, fondendo quindi la soluzione al sotto-problema di trovare il massimo e quella di contare quante volte compare. In questo modo si risolve il problema senza tenere la serie di numeri in memoria.

--- codice da completare con le istruzioni fornite più sotto ----
```go
package main
import "fmt"
func main() {
    var number int
    fmt.Scan(&number)
    max := number
    ....    //aggiungere qui, nell'ordine giusto, le istruzioni fornite sotto
	fmt.Printf("%v %v\n", max, count)
}
```

--- istruzioni da inserire per completare il programma qui sopra ---
```go
}
if number == max {
for i := 1; i<10; i++ {
count ++
max = number
count = 1
count := 1
fmt.Scan(&number)
} else if number > max{
}
```

nomefile: num_maxP.go

---

### Massimo numero di cifre pari
Scrivere un programma `max_num_cifre_pari.go` che, data una sequenza di numeri (da leggere come stringhe), terminata dalla stringa "000", stampa il numero di cifre pari contenute nel numero che ne contiene più di tutti.\
E` possibile risolvere il problema senza tenere in memoria sequenze di numeri. Che tipo di composizione occorrerà per mettere insieme il conteggio delle cifre pari e il calcolo del massimo di tali conteggi?

nomefile: max_num_cifre_pari.go

---

## Problemi

### Cesare (T)
Scrivere un programma `cesare.go` che legge da standard input un numero intero non negativo *k* (la chiave di cifratura), uno spazio e una sequenza di lettere ASCII minuscole consecutive (sulla stessa riga e senza spazi) terminate da invio, cifra le lettere lette con il cifrario di Cesare, usando come chiave *k* (quella fornita dall'utente) e stampa il testo cifrato: ogni lettera del testo in chiaro è sostituita nel testo cifrato dalla lettera che si trova *k* posizioni dopo nell'alfabeto, ritornando alla lettera 'a' dopo la 'z'.

Esempi di esecuzione:

**2** **zaprb**\
`bcrtd`


**100** **abcd**\
`wxyz `

---

nomefile: cesare.go

### L'ultimo giorno di pioggia
Scrivere un programma `ultima_pioggia.go` che legge da standard input una sequenza di numeri interi (terminata da EOF, da tastiera prodotto con ctrl-D) che indicano i mm di pioggia caduti (0 se non ha piovuto) ogni giorno in una sequenza successiva di giorni (numerati a partire da 1) e stampa l'ultimo giorno in cui ha piovuto.

**Nota** su come "monitorare" la lettura con Scan e interrompere un ciclo for quando si intercetta EOF:

```go
_, err := fmt.Scan(&myVar)
		if err == io.EOF {
			break
		}
```

nomefile: ultima_pioggia.go

### Sottosequenze
Scrivere un programma `num_sequenze`.go che legge da standard input una sequenza di uni (1) e zeri (0) (terminata da un 2), che inizia e finisce con 1, e stampa il numero di sottosequenze di zeri.

Ad esempio per input 1 1 0 0 1 0 1 0 0 0 1 1 1 0 1 2, l'output è 4 (vanno considerate anche le sottosequenze di lunghezza 1)

nomefile: num_sequenze.go

### Gradini
Definiamo "gradino" una sequenza di (uno o più) interi non negativi uguali seguita da un'altra sequenza di (uno o più) interi più grandi di 1 (es. 1 1 2 2 2).

Scrivere un programma `maxGradino.go` che, data in input una sequenza di interi tali che ogni intero è >= del precedente, stampa la lunghezza (il numero di interi) del gradino più lungo. (Si noti che i gradini si sovrappongono).

Il programma termina quando legge un numero minore di quello appena letto.

Potete dare per scontato che i numeri in input rispettino la condizione di essere ciascuno uguale o maggiore di una unità rispetto al numero precedente; non occorre quindi fare controlli in questo senso.

Per input 2 2 3 3 4 4 4 5 6 6 6 7 2, l'output è 5 (il gradino 3 3 4 4 4).

nomefile: maxGradino.go

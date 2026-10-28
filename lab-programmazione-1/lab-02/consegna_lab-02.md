# Lab02 - operatori di confronto e logici, selezione (if)

[statistiche consegne Lab01](Lab01_stats.md)

## ESERCIZI DI GO

**Per gli esercizi di scrittura**:
1. leggere attentamente le specifiche 
2. individuare le operazioni richieste / progettare la soluzione al problema utilizzando SOLO i costrutti di Go visti finora.
3. implementare la soluzione in Go, eventualmente copiando all'inizio del programma il testo dell'esercizio (tra /* e */).
4. compilare il programma (se non compila, correggere gli errori)
5. (se compila) lanciare il programma su un esempio di input
6. individuare il dominio del problema, e quindi i casi di test su cui testare il programma
7. lanciare il programma sui casi di test, controllando che l'output prodotto sia quello richiesto

*Nota*: negli esempi di esecuzione, l'input fornito dall'utente è evidenziato in **grassetto**, il resto è output prodotto dal programma.

Raccomandazioni:
- non utilizzate “numeri magici”, ma dichiarate costanti (const)
- scrivete programmi leggibili scegliendo nomi sensati per gli identificatori e indentando correttamente
- rispettare alla lettera le specifiche (questo non vuol dire che non possiate poi scrivere altri programmi di vostra iniziativa)
- leggete anche le note e rispettate gli esempi di esecuzione

**Per gli esercizi di analisi e comprensione del codice**:
1. osservate attentamente il codice fornito e usatelo come esempio di riferimento per risolvere esercizi simili
2. per comprendere cosa fa, potete sia lanciare l'esecuzione su qualche caso di input, sia tracciarlo passo passo con carta e penna, sempre su qualche input, sia cercare di comprendere cosa fa analizzandolo (e anche tutte e tre le cose).
3. Per scrivere le *specifiche*,  ispiratevi a quelle degli esercizi di scrittura.


## "if true else false"

Il seguente codice:
```
if n == 100 {
   fmt.Println(true)
} else {
   fmt.Println(false)
 }  
```
può essere scritto in modo migliore (più compatto e assolutamente equivalente) così:
```
fmt.Println(n == 100)
```
(Provate a tracciare i due frammenti di codice per n = 100 e per n diverso da 100 per convincervi che sono equivalenti)

Usate, quando si può, questa seconda forma più compatta, che rivedremo con le funzioni che restituiscono un valore booleano.

---

## ESERCIZI

### Esercizio 0.0
(Operatori di confronto e logici)

#### a)
Considerate le seguenti righe di codice, cosa fanno?
```
if r == 1 || r == 10 {
	fmt.Println(true)
} else {
	fmt.Println(false)
}
```
Cosa stampano rispettivamente se r=1, se r=2, e se r=10?

---
Considerate la seguente riga di codice, cosa fa?
```
    fmt.Println(r == 1 || r == 10)
```   
Cosa stampa rispettivamente se r=1, se r=2, e se r=10?

---

#### b)
Salvate il programma che segue in un file chiamato **condizioni0.go**, aggiungete come commento le specifiche del programma (cioè le indicazioni che dareste al programmatore che dovesse scriverlo) e poi caricatelo su upload.

```go
package main
import (
	"fmt"
	"math/rand"
)
/* scrivere qui le specifiche per questo programma */
func main() {
	n := rand.Intn(10) + 1
	fmt.Println(n)
	fmt.Println(n == 1 || n == 10)
}
```
nomefile: condizioni0.go

---

### Esercizio 0.1
(Operatori di confronto e logici)

*Specifiche*: Completare il programma Go `condizioni1.go`, che deve verificare, una a una, le condizioni nella tabella che segue. Trovate [qui](Lab02/condizioni1.go) il codice sorgente da copiare e completare. Non dovete usare  nessun `if`.

Il programma, per ogni condizione da verificare, legge un valore da tastiera (del tipo indicato nella colonna “tipo”) e stampa *true* o *false*, a seconda che la condizione sia verificata o no.


Completate il programma una condizione alla volta, testandola su almeno due input (uno che la verifica e uno che la falsifica) prima di procedere alla successiva.

Scopo dell'esercizio è imparare a usare gli operatori di confronto uguale, diverso, maggiore, minore, maggiore o uguale, minore o uguale (==, !=, >, <, >=, <=) e quelli logici NOT, AND, OR (!, &&, ||).

Nota: 
- la notazione [n1, n2] indica l'intervallo di valori da n1 a n2, estremi (cioè i valori n1 e n2) compresi
- la notazione (n1, n2) indica l'intervallo di valori da n1 a n2, estremi esclusi
- deducete voi il significato delle notazioni (n1, n2] e [n1, n2)

#### Tabella

| n. | tipo input | condizione |
|----|------|------------|
| 1| int| uguale a 10 | 
| 2| int| diverso da 10 | 
| 3| int| diverso da 10 e da 20 |
| 4| int| diverso da 10 o da 20 |
| 5| int| maggiore o uguale a 10 |
| 6| int| compreso tra 10 e 20, nell'intervallo [10,20] |
| 7| int| int compreso tra 10 e 20, nell'intervallo (10,20)|
| 8| int| compreso tra 10 e 20, nell'intervallo [10,20) |
| 9| int| esterno all'intervallo [10,20] |
|10| int| tra 10 e 20 (nell'intervallo [10,20]) o tra 30 e 40 ([30,40]) |
|11| int| multiplo di 4 ma non di 100 |
|12| int| dispari e compreso tra 0 e 100 ([0,100]) |
|13| float64| float64 "vicino" a 10.0 (non più lontano di 10^-6) (la notazione Go per 10^-6 è: 1e-6, che sta per 1*10^-6) |
---

Esempi di esecuzione (per la prima condizione):

int uguale a 10: **10**\
true

int uguale a 10: **11**\
false

---

*Domanda*: Avete usato degli if nel programma `condizioni1.go`? Perché sì o perché no o quando?

---

nomefile: condizioni1.go

---

### Esercizio 0.2 - Precedenza degli operatori logici
(Operatori di confronto e logici)

*Specifiche*: Scrivere un programma Go `precedenza.go` che legge tre valori booleani `a`, `b`, `c` (digitare `t` per true e `f` per false) e stampa il valore delle due espressioni che seguono, che differiscono solo per la posizione delle parentesi:

```
a || b && c
(a || b) && c
```

Nota. In Go `&&` ha precedenza su `||`: quindi `a || b && c` equivale a `a || (b && c)` e **non** a `(a || b) && c`. Le due espressioni possono quindi produrre risultati diversi: usate il programma per verificarlo su più input.

Esempio di esecuzione:

`tre bool (t/f): `**t f f**\
`a || b && c = true`\
`(a || b) && c = false`

nomefile: precedenza.go

---

### Esercizio 0.3 - Valutazione a corto circuito
(Operatori di confronto e logici)

*Specifiche*: Scrivere un programma Go `cortocircuito.go` che legge due interi `x` e `y` e stampa `true` se `x` è diverso da zero **e** la divisione intera `y/x` è maggiore di 10, `false` altrimenti. La condizione va scritta come un'unica espressione booleana:

```
x != 0 && y/x > 10
```

Nota. Gli operatori `&&` e `||` sono *valutati a corto circuito*: appena il risultato è determinato, la valutazione si ferma. Quindi, se `x` è 0, la sotto-espressione `y/x` (che causerebbe un errore di esecuzione, essendo una divisione intera per zero) **non viene mai calcolata**. Provate il programma con `x` uguale a 0 per verificare che non si verifichi alcun errore.

Esempio di esecuzione:

`due int (x e y): `**2 50**\
`true`

`due int (x e y): `**0 50**\
`false`

nomefile: cortocircuito.go

---

### Esercizio 0.4 - Leggi di De Morgan
(Operatori di confronto e logici)

*Specifiche*: Scrivere un programma Go `de_morgan.go` che legge un intero `n` e stabilisce se è esterno all'intervallo [10,20] (estremi compresi), calcolando il risultato in due modi diversi che devono dare sempre lo stesso valore:

```
n < 10 || n > 20
!(n >= 10 && n <= 20)
```

Nota. La negazione di una congiunzione equivale alla disgiunzione delle negazioni (e viceversa): `!(a && b)` equivale a `!a || !b`, mentre `!(a || b)` equivale a `!a && !b`. Queste sono le *leggi di De Morgan*.

Esempio di esecuzione:

`un int: `**5**\
`n < 10 || n > 20 = true`\
`!(n >= 10 && n <= 20) = true`

`un int: `**15**\
`n < 10 || n > 20 = false`\
`!(n >= 10 && n <= 20) = false`

nomefile: de_morgan.go

---

### Esercizio 1a - selezione a una via (*guarded action*)

<!---
Specifiche: Scrivere un programma Go voto_valido.go che legge un numero intero. Se il numero non è compreso tra 0 e 30, stampa “voto non valido”, altrimenti non stampa niente. 
Esempi di esecuzione:
```
voto: 42
voto non valido
voto: 28
```
--->

Salvare il seguente programma in un file chiamato `1-if-a-una-via.go`, aggiungere come commento le *specifiche* del programma (cioè le indicazioni che dareste al programmatore che dovesse scriverlo) e poi caricarlo su upload.

```go
package main

import "fmt"

func main() {
	/*
	 Scrivere qui le specifiche per questo programma e caricarlo su upload
	*/

	var voto int

	fmt.Print("voto: ")
	fmt.Scan(&voto)

	if voto < 0 || voto > 30 {
		fmt.Println("voto non valido")
	}
}
```

nomefile: 1-if-a-una-via.go

---

### Esercizio 1b - Maggiore
*Specifiche*: Scrivere un programma Go `maggiore.go` che legga due interi, li salvi in due var `min` e `max` nell'ordine in cui li legge; se non sono in ordine crescente, li sistemi in modo che `min` contenga il minore e `max` il maggiore; infine stampi il contenuto della variabile `max`.

Esempi di esecuzione:

due int: **10 20**\
max: 20

due int: **20 10**\
max: 20

nomefile: maggiore.go

---

### Esercizio 1c - Mattino
*Specifiche*: Scrivere un programma Go `mattino.go` che chieda e legga un orario (intero) e stampi "mattino" se l'orario è compreso tra le 6 (incluso) e le 13 (escluso).

Esempio di esecuzione:

scrivi un orario (int): **10**\
mattino

---

nomefile: mattino.go

---

### Esercizio 1d - Mattino2
*Specifiche*: Scrivere un programma Go `mattino2.go` che legga un orario (due numeri interi, ore e minuti), senza aver stampato una richiesta di input, e stampi "mattino" se l'orario è compreso tra le 5:30 (incluso) e le 12:30 (escluso). Fatelo con un solo if a una via; potranno esservi utili gli operatori logici. 

Esempio di esecuzione:

**5 35**\
mattino

---

Prima di procedere all'implementazione:
- disegnate su un foglio una linea, segnateci sopra i seguenti orari: 5:30, 6:00, 12:00, 12:30, definendo così i tre intervalli corrispondenti. 
- rispondete alle seguenti domande:
	- che caratteristiche hanno gli orari del primo intervallo?
 	- quelli del secondo?
 	- e quelli del terzo?
- traducete ciascuna risposta in una espressione logica che sia vera per gli orari nell'intervallo preso in considerazione e falsa altrimenti
- componete opportunamente le espressioni logiche in una condizione per l'if del vostro programma
- scrivete il resto del programma e testatelo su diversi orari presi opportunamente per essere sicuri che il programma funzioni correttamente su tutti i casi possibili

---

*Domanda*: Quanti `if` avete usato nel vostro programma? Qual è il minimo numero di `if` necessari?

---

nomefile: mattino2.go

---
===>>> TEST DISPONIBILE:
- si può utilizzare il file `mattino2_test.go`
- creare una directory chiamata `mattino2`
- salvare il proprio file `mattino2.go` nella directory
- copiare il file `mattino2_test.go` (scaricandolo su indicazioni del docente) nella directory
- copiare il file `lib_test.go` (scaricandolo su indicazioni del docente) nella directory
- creare il file `go.mod` col comando `go mod init mattino2` lanciato nella directory creata
- lanciare infine i test con `go build; go test -v` (sempre dentro la dir)
- si prega di consegnare su upload anche almeno un output dei test catturato tramite redirezione (`go build; go test -v > mattino2.student-test`), il nome del file deve essere `mattino2.student-test`.

---

### Esercizio 2a - selezione a due vie (due azioni per due situazioni mutuamente esclusive)

<!---Specifiche: Scrivere un programma pari_dispari.go che legge un intero n e, a seconda del valore di n, stampa uno dei messaggi "n è pari" oppure "n è dispari".

Esempi di esecuzione:
```
numero: 4
4 è pari
numero: 15
15 è dispari
```
--->

Salvare il seguente programma in un file chiamato `2-if-a-due-vie.go`, aggiungere come commento le *specifiche* del programma e poi caricarlo su upload.

```go
package main

import "fmt"

func main() {
	/*
	 Scrivere qui le specifiche per questo programma e caricarlo su upload
	*/
	
	var n int
	fmt.Print("numero: ")
	fmt.Scan(&n)

	if n%2 == 0 {
		fmt.Println(n, "è pari")
	} else {
		fmt.Println(n, "è dispari")
	}
}
```

---
nomefile: 2-if-a-due-vie.go

---

### Esercizio 2b - Frazioni
*Specifiche*: Scrivere un programma `frazioni.go` che verifica se due frazioni num1/den1 e num2/den2 sono equivalenti e stampa "equivalenti" o "non equivalenti", a seconda del caso.
Usare il tipo int per i dati in input.
Trovate una soluzione che **non** faccia uso di float64.

Esempi di esecuzione:

num e den fraz 1: **12 32**\
num e den fraz 2: **10 13**\
non equivalenti

num e den fraz 1: **1 2**\
num e den fraz 2: **2 4**\
equivalenti


**Suggerimento**: potete sfruttare la proprietà fondamentale delle proporzioni.

---
nomefile: frazioni.go

---

### Esercizio 2c - Intorno
*Specifiche*: Scrivere un programma `intorno.go` che legge due float64 x e y (senza aver stampato una richiesta di input) e stampa "molto vicino all'origine" se il punto di coordinate (x,y) è distante meno di 10^-5 dall'origine (0,0) del piano cartesiano, cioè effettivamente molto vicino all'origine,  mentre stampa "non vicino all'origine" in caso contrario.\
Impostate una costante usando la notazione scientifica di Go, ad esempio: \
 `const  EPSILON = 1e-5`
 
 Esempio di esecuzione:
 
 **0 0**\
molto vicino all'origine

---

nomefile: intorno.go

---
===>>> TEST DISPONIBILE:
- si può utilizzare il file `intorno_test.go`
- creare una directory chiamata `intorno`
- inserire il proprio file `intorno.go` nella directory
- copiare il file `intorno_test.go` (scaricandolo su indicazioni del docente) nella directory
- copiare il file `lib_test.go` (scaricandolo su indicazioni del docente) nella directory
- creare il file `go.mod` col comando `go mod init intorno` lanciato nella directory creata
- lanciare infine i test con `go build; go test -v` (sempre dentro la dir)
- si prega di consegnare su upload anche almeno un output dei test catturato tramite redirezione, il nome del file deve essere `intorno.student-test`


---

### Esercizio 3a - selezione a più vie (diverse possibilità mutuamente esclusive)

<!---
*Specifiche*: Scrivere un programma `pos_neg_nullo`.go che legge un int da standard input, stabilisce se è positivo, negativo o nullo e stampa il messaggio corrispondente.

Esempi di esecuzione:

un int: **12**\
positivo

un int: **0**\
nullo

un int: **-301**\
negativo
--->

Salvare il seguente programma in un file chiamato `3-if-a-piu-vie.go`, aggiungere come commento le *specifiche* del programma e poi caricarlo su upload.


```go
package main

import "fmt"

func main() {
	/*
	 Scrivere qui le specifiche per questo programma e caricarlo su upload
	*/

	var num int

	fmt.Print("un int: ")
	fmt.Scan(&num)

	if num > 0 {
		fmt.Println("positivo")
	} else if num == 0 {
		fmt.Println("nullo")
	} else { // < 0
		fmt.Println("negativo")
	}
}
```

---
nomefile: 3-if-a-piu-vie.go

---

### Esercizio 3b Indovina 10
*Specifiche*: Scrivere un programma `indovina_10.go` che fissa nel programma un numero intero (a vostra scelta) tra 1 e 10 da indovinare, richiede un numero (stampando "indovina il numero che ho pensato: "), legge un intero da standard input e:
- se il numero in input è fuori dall’intervallo [1-10], stampa “Valore non valido”;
- se il numero è quello fissato, stampa “Hai indovinato!”;
- altrimenti stampa “Non hai indovinato”.

---
nomefile: indovina_10.go

---

### Esercizio 3c Retta e punto
*Specifiche*: Scrivere un programma `punto_retta.go` che chiede (vedi esempio) e legge da standard input le cooordinate (x1,y1) di un punto P e i coefficienti m e q di una retta y = mx + q; controlla se il punto P sta sopra, sotto o appartiene alla retta y data, e stampa rispettivamente "P sopra R", "P sotto R", "P sulla retta R".
Si dichiarino i dati in ingresso come float64.

**Suggerimento**: disegnate su un foglio gli assi cartesiani, tracciate una retta a piacere e fissate un punto a piacere nel piano. Che condizione soddisfa un punto che appartiene a una retta data? Uno che si trova sopra la retta? Uno che sta sotto alla retta? Traducete nelle corrispondenti espressioni booleane e usatele nel vostro programma.

**Nota**: trattando con numeri float, cosideriamo appartenenti alla retta i punti che distano dalla retta meno di 10^-6. In Go 10^-6 può essere scritto come 1e-06 (usare sempre una costante in casi come questo).

Esempi di esecuzione:

x e y del punto P: **9.35 129**\
m e q della retta R: **3.6 12.45**\
P sopra R

x e y del punto P: **24 0**\
m e q della retta R: **5.2 25.25**\
P sotto R

x e y del punto P: **2 7**\
m e q della retta R: **1 5**\
P sulla retta R


---
nomefile: punto_retta.go

---

### Esercizio 3d – Benzarolo
Specifiche: Scrivere un programma `benzarolo.go` che calcola il prezzo di un rifornimento, dati litri (float64) e tipo di carburante (int, 0=benzina, 1=diesel, 2=alcol, 3=cherosene) secondo la seguente tabella dei prezzi:

|carburante|prezzo|
|----------|------|
|0 (benzina)|1.78|
|1 (diesel)|1.98|
|2 (alcol)|1.2|
|3 (cherosene)|1.1|

Esempio di esecuzione:

tipo di carburante? (0 (benzina), 1 (diesel), 2 (alcol), 3 (cherosene)): **1**\
litri? **18.53**\
36.6894

---
nomefile: benzarolo.go
---

### Esercizio 3e - Tariffe scontate
*Specifiche*: Scrivere un programma `tariffe.go` che chiede all’utente l’età (int) e se è studente (bool) e stampa il costo del biglietto di ingresso secondo la tabella che segue.
Provare a risolvere l'esercizio utilizzando solo 4 (e non di più) istruzioni di assegnamento a una variabile `tariffa` e una sola istruzione di stampa finale per la stampa del costo del biglietto.

*Nota*: per fornire in input un valore booleano, digitare `t` o `true` per vero e `f` o `false` per falso.

|età|tariffa|
|---|-------|
|[0 - 9) anni|gratis (0)|
|[9 - 18) anni|5|
|[18 - 26) anni|7.5|
|[26 - 65) anni|10|
|>= 65 anni|7.5|
|studenti >= 18 anni|5|

---
Esempi di esecuzione:

età: **28**\
studente? (t/f): **t**\
ingresso 5 euro

età: **4**\
studente? (t/f): **f**\
ingresso 0 euro

---
nomefile: tariffe.go

---


### Esercizio 4a – scelta sequenziale

<!--Specifiche: Scrivere un programma esiti.go che associa voti in lettere a punteggi secondo la seguente tabella. I voti sono int.

|punteggio|voto|
|---------|----|
|90 - 100| A|
|80 - 89| B|
|70 - 79| C|
|60 - 69| D|
|0 - 59| F|

Più precisamente il programma riceve in ingresso un valore intero tra 0 e 100 e stampa il voto in lettere corrispondente, o "voto non valido" se il voto è fuori dall’intervallo 0-100.

Esempi di esecuzione:
```
Voto: 100
A
Voto: 65
D
Voto: 110
voto non valido
```
-->

Salvare il seguente programma in un file chiamato `4-condizioni-in-sequenza.go`, aggiungere come commento le *specifiche* del programma, rispondere alle domade e poi caricarlo su upload.

```go
package main

import "fmt"

func main() {
	/* 	 
	Scrivere qui le specifiche per questo programma
	*/

	var voto int
	fmt.Print("voto? ")
	fmt.Scan(&voto)

	if voto < 0 || voto > 100 {
		fmt.Println("voto non valido")
		return //interrompe l'esecuzione del programma
	}

	if voto >= 90 {
		fmt.Println("A")
	} else if voto >= 80 {
		fmt.Println("B")
	} else if voto >= 70 {
		fmt.Println("C")
	} else if voto >= 60 {
		fmt.Println("D")
	} else {
		fmt.Println("F")
	}
}
/* DOMANDE */
/*
- D1. Per quali valori è vera la condizione del primo if? (quello subito dopo la Scan)
- R1.

- D2. Per quali valori viene eseguita la seguente istruzione?
    fmt.Println("B")
- R2.

- D3. Giustificate la vostra risposta alla domanda D2
- R3.

- D4. Il programma è impostato correttamente o sarebbe stato (più) corretto scrivere invece così?
 } else if voto >= 80 && voto < 90 {
- R4.

- D5. Perché?
- R5.
*/
```

nomefile: 4-condizioni-in-sequenza.go

---

### Esercizio 4b – Fasi della giornata
*Specifiche*: Scrivere un programma `fasi_giornata.go` che che legge un orario (intero) e stampa la fase della giornata corrispondente, secondo la tabella che segue, o "orario non valido" se l'intero è fuori dall'intervallo 0-23

|ora|fase della giornata|
|---------|----|
|0 - 6| notte|
|7 - 13| mattino|
|14 - 18| pomeriggio
|19 - 23| sera|

---
Esempi di esecuzione:

ora: **8**\
mattino

ora: **24**\
orario non valido

ora: **19**\
sera

---
 *Domanda*: Vi sembra un problema da risolvere con un programma simile a quello dell'esercizio precedente? Perché sì o perché no? Prima di progettare e implementare la soluzione, rispondete alla domanda e salvate la risposta come commento all'inizio del file fasi_giornata.go.

---
nomefile: fasi_giornata.go

---

### Esercizio 5a – azioni indipendenti

<!---
Specifiche: Scrivere un programma fizz_buzz.go che riceve in ingresso un numero intero e stampa "Fizz" se il numero è multiplo di 3, "Buzz" se il numero è multiplo di 5, "Fizz Buzz" se è multiplo sia di 3 sia di 5, niente altrimenti.

Esempi di esecuzione:
```
numero: 5
Buzz
numero: 4
numero: 15
Fizz Buzz
numero: 6
Fizz
```
--->

Salvare il seguente programma in un file chiamato `5-azioni-indipendenti.go`, aggiungere come commento le *specifiche* del programma, rispondere alle domade e poi caricarlo su upload.

```go
package main

import "fmt"

func main() {
	/* 
	Scrivere qui le specifiche per questo programma e caricarlo su upload
	*/

	var num int
	fmt.Print("numero? ")
	fmt.Scan(&num)

	if num%3 == 0 {
		fmt.Println("Fizz")
	}
	if num%5 == 0 {
		fmt.Println("Buzz")
	}
}
/* DOMANDE */
/*
- D1. Se al posto del secondo if ci fosse un else if (legato al primo if), il programma si comporterebbe nello stesso modo? Dareste le stesse specifiche?
- R1.

- D2. Perché? Spiegate la vostra risposta alla domanda D1
- R2.

- D3. Cosa stampa il programma per ciascuno dei seguenti input: 8, 9 10, 15?
- R3
   8: 
   9: 
  10: 
  15:
 

*/
```


nomefile: 5-azioni-indipendenti.go

---

### Esercizio 5b – Numeri
*Specifiche*: Scrivere un programma `numeri.go` che chiede e riceve in ingresso un numero intero, e stampa "divisibile per 10" se il numero è divisibile per 10, "positivo" se il numero è positivo, niente altrimenti.


 *Domanda*: Che tipi di if serviranno per stampare i diversi messaggi coerentemente con l'input ricevuto? If a una via o if a più vie? Quanti del primo tipo e quanti del secondo?
 Prima di progettare e implementare la soluzione, rispondete, argomentando, alla domanda e salvate la risposta come commento all'inizio del file numeri.go.

Esempi di esecuzione:

numero? **100**\
positivo\
divisibile per 10

numero? **123**\
positivo

numero? **-30**\
divisibile per 10

numero? **-486**


---
nomefile: numeri.go

---

### Esercizio 5c – Guasti 
*Specifiche*: Scrivere un programma `guasti.go` che riceve in input tre interi che segnalano eventuali guasti a diverse componenti di un impianto: 0 per nessun guasto, un numero diverso da 0 in caso di guasto. \
Il programma deve stampare "guasti rilevati:" seguito, uno per riga, dai guasto segnalati (vedi esempio di esecuzione).\
La tabella che segue indica, per ciascun input, se è diverso da 0, a che quasto corrisponde.

|input|guasto|
|---------|----|
| 1° input | caricamento acqua |
| 2° input | scarico acqua |
| 3° input | sistema di riscaldamento |


---

Esempio di esecuzione:

componente 1, guasto rilevato? (0 per ok) **3**\
componente 2, guasto rilevato? (0 per ok) **0**\
componente 3, guasto rilevato? (0 per ok) **5**\
guasti rilevati:\
caricamento acqua\
sistema di riscaldamento


---
nomefile: guasti.go

---

<!---
### Esercizio 5d – Appartamento

*Specifiche*: Scrivere un programma `appartamento.go` che riceve in ingresso due numeri (un intero per il piano e un float64 per la superficie) e stampa "piano alto" se il piano è superiore a 5, stampa "ampia metratura" se la superficie è maggiore di 150 metri quadrati, "ampia metratura, piano alto" se valgono entrambe, niente altrimenti.
--->

### Esercizio 6a - if annidati  (azioni condizionate a diversi fattori indipendenti)
Salvare il seguente programma in un file chiamato `6-if-annidati.go`, aggiungere come commento le *specifiche* del programma, rispondere alle domade e poi caricarlo su upload.

```go
package main

import "fmt"

func main() {
	/*
	 Scrivere qui le specifiche per questo programma e caricarlo su upload
	*/
	const (
		RETTANGOLO = 1
		TRIANGOLO  = 2
		PERIMETRO  = 1
		AREA       = 2
	)
	var figura, operazione int

	fmt.Println("di che figura si tratta?")
	fmt.Print("rettangolo (1) o triangolo (2)? ")
	fmt.Scan(&figura)
	fmt.Println("cosa vuoi calcolare?")
	fmt.Print("perimetro (1) o area (2)? ")
	fmt.Scan(&operazione)

	if figura == RETTANGOLO {
		if operazione == PERIMETRO {
			fmt.Println("formula = (base + altezza) * 2")
		} else { //operazione == AREA
			fmt.Println("formula = base * altezza")
		}
	} else { //figura == TRIANGOLO
		if operazione == PERIMETRO {
			fmt.Println("formula = lato1 + lato2 + lato3 ")
		} else { //operazione == AREA
			fmt.Println("formula = (base * altezza)/2")
		}
	}
}
/* DOMANDE */
/*
- D1. Quanti diversi messaggi da stampare ci sono?
- R1.

- D2. Sarebbe possibile ridurre il numero di if/else? Perché?
- R2.
*/
```

nomefile: 6-if-annidati.go

---

### Esercizio 6b - Tasse
*Specifiche*: Scrivere un programma `tasse.go` che chiede reddito (int) e se coniugato (bool) e stampa le tasse da pagare secondo la seguente tabella.
Usate costanti per le aliquote (10% e 24%) e per gli scaglioni (32000 e 64000).

|non coniugato|coniugato|aliquota|
|-------------|---------|--------|
|0 - 32000|0 - 64000|10%|
|> 32000|> 64000|24%|

---
Esempi di esecuzione

reddito? **40000**\
coniugato? (true/false) **t**\
tasse per coniugato con reddito 40000: 4000

reddito? **40000**\
coniugato? (true/false) **f**\
tasse per non coniugato con reddito 40000: 9600

---
*Domanda*. Quante diverse aliquote da applicare ci sono?
Quanti if/else occorrono come minimo? Di che tipo? Perché?


Rispondete, argomentando, alla domanda e salvate la risposta come commento all'inizio del file tasse.go, poi progettare le condizioni per la selezione coerentemente con le vostre risposte, infine implementare la soluzione, testatela e caricatela su upload.

---
nomefile: tasse.go

---

## PROBLEMI

### Problema 1 – Data valida

- P1.a (Anno bisestile) - *Specifiche*: Scrivere un programma `bisestile.go` che legge un intero *n* non negativo corrispondente all’anno di una data, e determina se l’anno *n* è bisestile o no (stampa "bisestile" o "non bisestile"). Vedi sotto la regola per determinare se un anno è bisestile.


- P1.b -*Specifiche*: Scrivere un programma `data_valida1.go` che legge due interi rappresentanti giorno e mese dell’anno, e verifica che sia una data valida. In questa versione si assuma che tutti i mesi abbiano 31 giorni. Il programma stampa "data valida" / "data non valida".

- P1.c - *Specifiche*: Scrivere una seconda versione `data_valida2.go` che legge due interi rappresentanti giorno e mese dell’anno, e verifica che sia una data valida. In questa seconda versione si tenga conto del fatto che solo i mesi 1, 3, 5, 7, 8, 10, 12 hanno 31 giorni, che i mesi 4, 6, 9, 11 ne hanno 30, e si assuma che febbraio ne abbia sempre 28. Il programma stampa "data valida" / "data non valida".

- P1.d - *Specifiche*: Scrivere una terza versione `data_valida3.go` che legge tre interi rappresentanti giorno, mese e anno (anno >=0) e verifica che sia una data valida. Questa versione finale tiene conto anche degli anni bisestili.

---
Esempi di esecuzione di data_valida3.go :

gg mm aaaa? **29 02 2018**\
data non valida

gg mm aaaa? **29 02 2020**\
data valida

---
**Regola per bisestile**

Il calendario gregoriano si applica dal 1582, anno della sua introduzione. Benché, in via teorica, sia possibile estenderlo anche agli anni precedenti, per questi, di norma, non si applicano gli anni bisestili e si usa il calendario giuliano. Sono quindi bisestili tutti gli anni divisibili per 4, compresi quelli secolari (100, 200, ...), dall'anno 0 al 1580 (non consideriamo quelli a.C.). Per gli anni dal 1582 (calendario gregoriano) sono invece bisestili:
- gli anni non secolari il cui numero è divisibile per 4;
- gli anni secolari il cui numero è divisibile per 400.

Per esempio, il 1896 e il 1996 sono stati entrambi bisestili (non secolari divisibili per 4), il 1800 e il 1900 non lo sono stati (secolari non divisibili per 400), mentre il 1600 e il 2000 lo sono stati (secolari divisibili per 400).

---

nomefile: bisestile.go

nomefile: data_valida1.go

nomefile: data_valida2.go

nomefile: data_valida3.go

---


### Problema 2 - Quadrante

Specifiche: Scrivere un programma `quadrante.go` che legge da tastiera le coordinate (*x*, *y*), di tipo float64, di un punto P nel piano cartesiano (senza aver stampato una richiesta di input) e stampa in che quadrante (I, II, III, IV) si trova P. Per i punti che cadono sull’asse delle ordinate e/o delle ascisse, assegnate il quadrante positivo (+).


```
      |
  II  |  I
(-,+) | (+,+)
------+---------
      |
  III | IV
(-,-) | (+,-)

```

Esempi di esecuzione:

**2.2 -5**\
IV quadrante

**-3.7 -25.1**\
III quadrante

---

nomefile: quadrante.go


---
===>>> TEST DISPONIBILE:
- si può utilizzare il file `quadrante_test.go`
- creare una directory chiamata `quadrante`
- inserire il proprio file `quadrante.go` nella directory
- copiare il file `quadrante_test.go` (scaricandolo su indicazioni del docente) nella directory
- copiare il file `lib_test.go` (scaricandolo su indicazioni del docente) nella directory
- creare il file `go.mod` col comando `go mod init quadrante` lanciato nella directory creata
- lanciare infine i test con `go build; go test -v` (sempre dentro la dir)
- si prega di consegnare su upload anche almeno un output dei test catturato tramite redirezione, il nome del file deve essere `quadrante.student-test`



---

### Problema 3 – Sovrapposizione

Specifiche: Scrivere un programma `sovrapposizione.go` che legge da tastiera il giorno [1-31], l'ora di inizio [0-24] e l'ora di fine [0-24] di due appuntamenti e stabilisce se si sovrappongono (anche
parzialmente) o no.

Esempi di esecuzione:

appuntamento 1 (gg, start, end): **28 15 18**\
appuntamento 2 (gg, start, end): **28 16 20**\
si sovrappongono

appuntamento 1 (gg, start, end): **11 15 18**\
appuntamento 2 (gg, start, end): **28 16 20**\
non si sovrappongono

---

*Domande*: disegnate una linea del tempo e due appuntamenti che non si sovrappongono. In quanti diversi modi possono essere disposti?
In quanti diversi modi possono essere disposti invece due appuntamenti che si sovrappongono? Stabilite quale delle due situazioni è più semplice da definire e scrivete la condizione corrispondente, poi usatela nel vostro programma.


---
nomefile: sovrapposizione.go

---

### Problema 4 – Rettangolo
Specifiche: Scrivere un programma `rettangolo.go` che legge da tastiera le coordinate (int) *x* e *y* di tre punti P1, P2, P3 e stabilisce se P3 è interno, perimetrale o esterno al rettangolo che ha P1 e P2 come vertici di una delle sue due diagonali, non importa in che ordine e di quale diagonale (facendo riferimento alla figura,  in input P1 e P2 potrebbero arrivare scambiati tra loro o anche sostituiti da P4 e P5).

In figura un rettangolo e tre possibili punti P3: P3a esterno, P3b interno e P3c perimetrale.
```
      .P3a

P4  ---.P3c------ P2
   |             |
   |      .P3b   |
   |             |
    -------------
P1                P5


```

Esempio di esecuzione:

P1: **0 0**\
P2: **3 5**\
P3: **-1 0**\
P3 esterno

---
Nota: indipendentemente da quali due vertici opposti di un rettangolo riceviamo in input, e in che ordine, il vertice P1 in figura è dato da (min(x1,x2), min(y1,y2)) e il vertice P2 da (max(x1,x2), max(y1,y2)).

*Domande*. (i) Se si volesse definire un'espressione booleana che è vera solo quando P3 è un punto di uno dei due lati verticali, e falsa altrimenti, come andrebbe definita? (ii) E per P3 sui lati orizzontali?

Dichiarate due variabili booleane `suLatoVerticale` e `suLatoOrizzontale` e assegnate alla prima l'espressione (i) e alla seconda l'espressione (ii).

Utilizzate poi queste due variabili nel vostro programma.

In modo analogo, se volete, potete definire una condizione per P3 punto interno o una per P3 punto esterno al rettangolo.

---
nomefile: rettangolo.go

---

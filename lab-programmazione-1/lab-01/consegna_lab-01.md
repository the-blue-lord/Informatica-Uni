# Lab01 - variabili int e float64, costanti, operatori aritmetici, I/O

## COME AFFRONTARE GLI ESERCIZI

Per ogni esercizio:
1. leggere attentamente le specifiche
1. individuare le operazioni richieste / progettare la soluzione al problema utilizzando SOLO i costrutti di Go visti finora.
1. implementare la soluzione in Go, eventualmente copiando all'inizio del programma il testo dell'esercizio (racchiuso tra /* e */). 
1. compilare il programma, e in caso di non compilazione, utilizzare i rimandi del compilatore per correggerlo
1. (se compila) lanciare il programma su un esempio di input
1. verificare che l'output sia quello richiesto sia in termini di valori che in termini di formato
1. individuare il dominio del problema, e quindi i casi di test su cui testare il programma
1. lanciare il programma su ciascuno dei casi di test, controllando ogni volta che l'output prodotto sia quello richiesto

**Nota**: negli *esempi di esecuzione*, l'input fornito dall'utente è evidenziato in **grassetto**, il resto
è l'output prodotto dal programma.

### Raccomandazioni:
- Non utilizzate “numeri magici”, ma dichiarate costanti (const) per fattori di conversione, ecc.
- Scrivete programmi leggibili scegliendo nomi sensati per gli identificatori e indentando correttamente
- Rispettate alla lettera le specifiche degli esercizi proposti (questo non vuol dire che non possiate poi scrivere altri programmi di vostra iniziativa)
- Leggete anche le note e rispettate il formato dell'output, osservando gli esempi di esecuzione

### Notazione sintattica < >
Per specificare la notazione sintattica da usare in un dato contesto, si usa mettere tra < > i segnaposto che andranno sostituiti con le opportune sequenze di caratteri, inseriti nella sequenza di caratteri fissa per quel comando, indirizzo, ecc.

Ecco alcuni esempi per chiarire l'uso di questa notazione.

Per indicare la sintassi di un indirizzo di posta elettronica di uno studente unimi si userà:

`<nome>.<cognome>@studenti.unimi.it`

Per indicare la sintassi di un nome di file con estensione:

`<nomeFile>.<estensione>`

Per indicare un file con estensione txt 

`<nomeFile>.txt` 

Osservate che negli esempi alcune parti (quali?) sono tra < > e altre (quali?) no.

### Notazione sintattica [ ]
Per indicare che un elemento è facoltativo, cioè potrebbe esserci, ma anche non esserci, lo si mette tra parentesi quadre ([]).

Ad esempio

`<nomeFile>[.<estensione>]`

indica la sintassi per un nome di file, con o senza estensione.

### Lettura e stampa 

L'**istruzione per leggere** un valore da standard input e salvarlo in una variabile è:

`fmt.Scan(&<nomeVar>)`

ad esempio `fmt.Scan(&num)`

(notate il simbolo '&' che precede - senza spazi - il nome della variabile)

e l'**istruzione per stampare** è:

`fmt.Println("<testo da stampare>")`

ad esempio `fmt.Println("ciao")`



Nel caso si vogliano stampare valori contenuti in variabili, alternati a testo:

`fmt.Println(["<testo da stampare>",] <nomeVar_1>[, <nomeVar_2>, ..., <nomeVar_n>] [, "<altro testo>"][...])`

ad esempio `fmt.Println("ciao", yourName)`

Quando **non** si vuole avere né un a-capo finale né spazi tra gli elementi separati da virgole, si usa:

`fmt.Print(["<testo da stampare>",] <nomeVar_1>[, <nomeVar_2>, ..., <nomeVar_n>] [, "<altro testo>"][...])`

Provate e osservate la differenza.


### Ridirezione di input e output

E` possibile **ridirigere l'output** di un comando o di un programma da standard output a un file:

- ridirezione con sovrascrittura:
	`<comando> > <nomeOutFile>`

- ridirezione con memorizzazione in coda al file:
		`<comando> >> <nomeOutFile>`

- ad esempio: `echo "ciao" > saluti`

Potete usare poi i comandi `cat/more/less` per ispezionare il contenuto del file, ad esempio:
	`cat outFile`


Analogamente è possibile **ridirigere un file su standard input** per passare dati a un comando o a un programma, al posto di fornirglieli da tastiera: `<comando> < <nomeInFile>`.

Ad esempio: `./myProgram < inputForMyProgram`

## ESERCIZI


## Esercizio 0 - Programma vuoto
Problema: Scrivere un programma Go **main.go** che non abbia nessuna istruzione, ma solo la dichiarazione del package e della funzione main.

nomefile: main.go

## Esercizio 1 - Programma senza output
(Dichiarazioni)

Problema: Scrivere un programma Go **no_out.go** che dichiara una variabile int e una costante = 10, 
assegna alla variabile un valore doppio della costante, poi (in un'istruzione successiva) somma 1 alla variabile, ma non stampa niente.

Nota. Per dichiarare una costante, scegliere un nome tutto maiuscolo e fare un assegnamento preceduto dalla parola "const". Esempio:

`const CAP_MILANO = 20100`

nomefile: no_out.go


## Esercizio 2a - Operazioni aritmetiche
(Operatori aritmetici)

Problema: Completare il programma Go **operazioni.go** che trovate qui sotto in modo che, dati due numeri float64 da standard
input, esegua addizione, sottrazione, moltiplicazione e divisione dei due numeri e stampi i risultati ottenuti.

Testo del programma **operazioni.go** da copiare e incollare in un file e completare:
```go
package main

import "fmt"

func main() {
	//dichiarare qui due variabili x1 e x2 di tipo float64

	fmt.Print("fornire due numeri float: ")
	fmt.Scan(&x1, &x2)
	//inserire qui le istruzioni per svolgere le operazioni
	//salvando i risultati con degli short assignment
	// (i nomi delle variabili da utilizzare per i risultati sono quelli usati nelle istruzioni di stampa qui sotto, cioè somma, differenza, prodotto, quoziente, )

	// osservate qui sotto i diversi formati di argomenti passati alla funzione fmt.Println
	fmt.Println(x1, "+", x2, "=", somma)
	fmt.Println("differenza =", differenza)
	fmt.Println("il prodotto di", x1, "e", x2, "dà", prodotto)
	fmt.Println(x1, "/", x2, "=", quoziente)
}
```

Esempio di esecuzione:

`fornire due numeri float:` **6.6 &nbsp; 3.3**\
`somma = 9.899999999999999`\
`differenza = 3.3`\
`prodotto = 21.779999999999998`\
`quoziente = 2`

Nota. Con la conversione in binario c’è la possibilità, da tenere sempre presente, che si
introducano errori di approssimazione per i float64 e di overflow per gli int (lo vedrete più avanti)


nomefile: operazioni.go

## Esercizio 2b - Operazioni aritmetiche con interi
(Operatori aritmetici)

Problema: Adattare il programma dell'esercizio precedente scrivendo un programma Go **operazioni_int.go** che, dati due numeri interi da standard
input, esegua addizione, sottrazione, moltiplicazione e divisione intera dei due numeri e stampi i risultati ottenuti.

nomefile: operazioni_int.go

## Esercizio 2c - Operazioni aritmetiche con interi
(Operatori aritmetici)

Problema: Adattare il programma dell'esercizio precedente scrivendo un programma Go **operazioni_int2.go** che, dati due numeri interi da standard
input, esegua addizione, sottrazione, moltiplicazione e divisione float dei due numeri e stampi i risultati ottenuti.

nomefile: operazioni_int2.go

## Esercizio 3 - Programma echo
(I/O)

Problema: Scrivere un programma Go **echo.go** che legge un valore di tipo int da tastiera e lo
stampa.

nomefile: echo.go


## Esercizio 4 - Quoziente e resto
(Operatore modulo)

Problema: Scrivere un programma Go **quoziente_resto.go** che legge da standard input un dividendo e un divisore (interi), calcola il quoziente e il resto e li stampa.

Nota. L’operatore per la divisione (/) tra interi calcola la parte intera del risultato; l’operatore per il resto della divisione è %.

Esempio di esecuzione:

`dividendo: `**2500**\
`divisore: `**235**\
`quoziente = 10`\
`resto = 150`

nomefile: quoziente_resto.go


## Esercizio 5 - Angoli di un triangolo
(Costanti)

Problema: Scrivere un programma Go **angolo_triangolo.go** che, date in input le ampiezze in gradi (float64) di due angoli di un triangolo, determini l’ampiezza del terzo angolo.

Nota. La somma degli angoli di un triangolo è sempre 180 gradi. È quindi il caso di utilizzare
una costante per questo valore.

Esempio di esecuzione:

`angolo uno e due:` **25.3 &nbsp; 90.2**\
`ampiezza terzo angolo: 64.49999999999999`

nomefile: angolo_triangolo.go


## Esercizio 6 - Scambio di valori
(Assegnamenti multipli e “blank identifier”)

Problema: Scrivere un programma Go **scambio.go** che, dati in input due valori di tipo int,
controlli che i valori letti siano effettivamente di tipo int (come fare? vedere la documentazione di `fmt.Scan`), li stampi nell’ordine in cui sono stati forniti,
scambi i due valori (senza appoggiarsi a altre variabili), e poi esegua la stessa istruzione di
stampa (per mostrare che i valori sono stati effettivamente scambiati tra le due var).

Provare il programma su input interi, float, testo, e osservare che risultati si ottengono.

Esempio di esecuzione:

`val1 e val2 (int):` **10 &nbsp; &nbsp;20**\
`errore: <nil>`\
`prima dello scambio: 10 20`\
`dopo lo scambio: 20 10`

nomefile: scambio.go

## Esercizio 7 - Senza specifiche

Lanciando il programma eseguibile [**oracolo**](Lab01/oracolo) su diversi input scoprite cosa fa. Quindi scrivetene le specifiche (mettetele come commento all'inizio del file guessme.go) e implementate un programma **guessme.go** che si comporta nello stesso modo.

Per aggiungere i permessi in esecuzione usare il comando `chmod`:

`chmod +x oracolo`

nomefile: guessme.go


## Esercizio 8 - Uso della libreria standard - Import di più pacchetti

Problema: Scrivere un programma Go **random.go** che generi un numero random int tra 0 e 100 e lo stampi.

Nota. Vedere la documentazione della funzione `Intn` del package "math/rand". 

nomefile: random.go


## Esercizio 9 - Conversione da gradi centigradi a Fahrenheit
(Ancora costanti - Conversione tra due grandezze)

Problema: Scrivere un programma Go **centigradi_fahrenheit.go** che converta in gradi Fahrenheit una temperatura espressa in gradi centigradi letta da standard input. 
Si usino variabili di tipo float64.

Nota. Un grado Fahrenheit è 5./9. di un grado centigrado e 0 gradi centigradi corrispondono a 32
gradi Fahrenheit. Usare costanti (const) nella formula di conversione. Non usare direttamente i numeri (*numeri magici*). Attenzione alla differenza in
Go tra 5/9 e 5./9.

Esempio di esecuzione:

`temperatura in centigradi? `**18**\
`18 C = 64.4 F`

nomefile: centigradi_fahrenheit.go


## Esercizio 9 - Conversione da giorni in anni, settimane, giorni
(Ancora costanti - Conversione in serie di grandezze multiple)

Problema: Scrivere un programma Go **conversione_giorni.go** che converta un numero specificato di giorni (letto da tastiera) in anni, settimane, giorni senza mai usare l’operazione di sottrazione. Si considerino tutti gli anni di 365 giorni.

Nota. Dichiarare costanti (const) per il numero di giorni (365) in un anno e il numero di giorni in una
settimana (7). Non usare direttamente i numeri (*numeri magici*).

Esempio di esecuzione:

`numero di giorni da convertire?`\
**1329**\
`1329 giorni = 3 anni + 33 settimane + 3 giorni`

nomefile: conversione_giorni.go


# PROBLEMI

## Problema 1 - Distanza tra due punti
Problema: Scrivere un programma Go **distanza_punti.go** che calcola la distanza tra due punti nel piano cartesiano.

Annotazioni: Siano P1 = (x1, y1) e P2 = (x2, y2). La formula della distanza tra due punti P1 e P2 è:

radiceQuadrata( (x2 − x1)^2 + (y2 − y1)^2 )

dove ^2 indica l'elevamento alla seconda.

Siano x1, y1, x2, y2 tutti valori float64.

Nota. La funzione per calcolare la radice quadrata si chiama `Sqrt` ed è messa a disposizione nel package "math". Per calcolare invece il quadrato z^2 di un numero z si usi z*z.

Esempio di esecuzione:

`x e y del primo punto: `**5 &nbsp; &nbsp;8**\
`x e y del secondo punto: `**7 &nbsp; &nbsp;11**\
`La distanza tra i due punti è: 3.605551275463989`

nomefile: distanza_punti.go


## Problema 2 - Consumo medio e resa di un motore
Problema. Scrivere un programma Go **consumo_resa.go** che calcola il consumo medio e la resa di un motore date la distanza totale percorsa (in km) e la quantità di carburante utilizzata (in litri). I valori in input sono di tipo float64.

Nota. Il consumo medio di carburante si esprime in l/km ed è la quantità di carburante che occorre in media per percorrere un km di strada. La resa di un motore è data dalla distanza percorsa in media con un litro di carburante e si esprime in km/l.

Esempio di esecuzione:

`distanza percorsa (in km): `**50**\
`quantità di carburante utilizzata (in l): `**10**\
`consumo medio: 0.2 l/km`\
`resa media: 5 km/l`

nomefile: consumo_resa.go


## Problema 3 - Conversione da secondi in giorni, ore, minuti, secondi

Problema. Scrivere un programma Go **conversione_secondi.go** che converta un numero dato di secondi (fornito dall’utente) in giorni, ore, minuti, secondi senza mai usare l’operazione di sottrazione.

Esempio di esecuzione:

`numero di secondi: `**123456**\
`g:h:m:s = 1:10:17:36`

Nota: per stampare il risultato usare\
`fmt.Print("g:h:m:s = ", gg, ":", hh, ":", min, ":", sec, "\n")`\
(eventualmente sostituendo i nomei delle variabili con quelli che avete usato voi).

nomefile: conversione_secondi.go


## Problema 4 - Intersezione di due rette
Problema. Scrivere un programma Go **intersezione_rette.go** che, date le equazioni di due rette, stabilisce in che punto si intersecano. I dati in input sono di tipo float64.

Nota 1. L’equazione della retta è: *y = mx + q* \
dove *m* è il coefficiente angolare della retta e *q* l'intercetta dell'asse *y*. 

Nota 2. Dal punto di vista matematico, per determinare il punto di intersezione tra le due rette, occorre risolvere un sistema di primo grado.

Esempio di esecuzione:

`retta 1: m e q? `**1 &nbsp; &nbsp;4**\
`retta 2: m e q? `**2 &nbsp; &nbsp;6**\
`intersezione in (-2,2)`

nomefile: intersezione_rette.go

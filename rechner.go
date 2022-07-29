package main

import "fmt"

func main() {

	var zahl1, zahl2 int
	var operator string
	var erneut string
	var ergebnis int

	fmt.Println("Bitte geben Sie die erste Zahl ein:")
	fmt.Scan(&zahl1)
	fmt.Println("Bitte geben Sie den Operator ein:")
	fmt.Scan(&operator)
	fmt.Println("Bitte geben Sie die zweite Zahl ein:")
	fmt.Scan(&zahl2)

	switch operator {
	case "+":
		ergebnis = zahl1 + zahl2
		fmt.Println("Das Ergebnis ist:", ergebnis)
	case "-":
		ergebnis = zahl1 - zahl2
		fmt.Println("Das Ergebnis ist:", ergebnis)
	case "*":
		ergebnis = zahl1 * zahl2
		fmt.Println("Das Ergebnis ist:", ergebnis)
	case "/":
		ergebnis = zahl1 / zahl2
		fmt.Println("Das Ergebnis ist:", ergebnis)
	default:
		fmt.Println("Ungültiger Operator")

	}

	fmt.Println("Moechten Sie eine weitere Berechnung durchfuehren? (ja/nein)")
	fmt.Scan(&erneut)

	if erneut == "ja" {
		main()
	} else {
		fmt.Println("Bis Bald !!")
	}
}

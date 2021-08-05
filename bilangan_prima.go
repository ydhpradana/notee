package main

import "fmt"

func main() {
	var bilangan, i, faktor int = 0, 0, 1

	fmt.Print("Input Bilangan: ")
	fmt.Scanln(&bilangan)

	for i = 1; i <= bilangan; i++ {
		if bilangan%i == 0 {
			faktor++
		}
	}

	if faktor == 2 {
		fmt.Println("Angka", bilangan ,"adalah Bilangan Prima")
	} else {
		fmt.Println("Angka", bilangan ,"adalah Bukan Bilangan Prima")
	}
}

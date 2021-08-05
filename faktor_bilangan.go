package main

import "fmt"

func main() {
	var bilangan, i int

	fmt.Print("Input Bilangan: ")
	fmt.Scanln(&bilangan)
	fmt.Println("Faktor dari bilangan", bilangan ,"adalah:")
	for i = 1; i <= bilangan; i++ {
		if bilangan%i == 0 {
			fmt.Println(i)
		}
	}
}

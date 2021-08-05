package main

import (
	"fmt"
)

type student struct {
	name       string
	nameEncode string
}

type Chiper interface {
	Encode() string
	Decode() string
}

func (s *student) Encode() string {
	runes := []rune(s.name)
	nameEncode := ""
	for index, char := range runes {
		char = 97 + (26 - (char-96))
		runes[index] = char
	}
	nameEncode = string(runes)
	return nameEncode
}

func (s *student) Decode() string {
	runes := []rune(s.nameEncode)
	nameDecode := ""
	for index, char := range runes {
		char = 96 + (26 + (97-char))
		runes[index] = char
	}
	nameDecode = string(runes)
	return nameDecode
}

func main() {
	var menu int
	var a = student{}
	var c Chiper = &a
	fmt.Print("[1] Encrypt \n[2] Decrypt \nChoose your menu ? ")
	fmt.Scan(&menu)
	if menu == 1 {
		fmt.Print("\nInput Student's Name : ")
		fmt.Scan(&a.name)
		fmt.Print("\nEncode of Student’s Name " + a.name + " is : " + c.Encode())
	} else if menu == 2 {
		fmt.Print("\nInput Student's Encode Name : ")
		fmt.Scan(&a.nameEncode)
		fmt.Print("\nDecode of Student’s Name " + a.nameEncode + " is : " + c.Decode())
	} else {
		fmt.Println("Wrong input name menu abcd")
	}
}

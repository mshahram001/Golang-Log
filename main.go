package main

import (
	"log"
)

func div(firstnum, secondnum int) {
	if secondnum == 0 {
		log.Fatal("Can not be divided by 0!")
	}
	result := firstnum / secondnum
	log.Println("Result is:", result)
}

func main() {
	firstnum := 20
	secondnum := 5
	div(firstnum, secondnum)
}

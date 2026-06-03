package main

import (
	"fmt"
	"kasperaldrin/oeis/pkg/services"
	"log"
	"os"
)

func main() {
	f, err := os.Open("A000001.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	seq, err := services.ParseOEIS(f)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(seq.ID, seq.Name)
	fmt.Println("Terms:", seq.Sequence[:10])
	fmt.Println("Keywords:", seq.Keywords)
}

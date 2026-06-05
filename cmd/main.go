package main

import (
	"fmt"
	"log"

	"github.com/kasperaldrin/go-oeis"
)

func main() {

	client := oeis.NewClient(&oeis.OEISClientConfig{
		Mode:        oeis.ModeOffline,
		OfflinePath: "/Users/kasperaldrin/Documents/Agents/oeis/oeisdata",
	})

	seq, err := client.Get("A000001")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(seq.ID, seq.Name)
	fmt.Println("Terms:", seq.Sequence[:10])
	fmt.Println("Keywords:", seq.Keywords)
	fmt.Println("Programs:", seq.Programs[0].Language, "\n", seq.Programs[0].Code)

	for seq, err := range client.Sequences() {
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(seq.ID, seq.Name)
	}

	/*
		f, err := os.Open("oeisdata/seq/A000/A000001.seq")
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
	*/
}

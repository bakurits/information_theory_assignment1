package main

import (
	"assignment1/filestatistics"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {

	//var symbols = " აბგდევზთიკლმნოპჟრსტუფქღყშჩცძწჭხჯჰ"
	filename := "PublicTests/B/001.dat"
	f, err := os.Open(filename)
	if err != nil {
		log.Fatal("error in opening file")
	}
	defer func() {
		err = f.Close()
		if err != nil {
			log.Fatal(err)
		}
	}()

	singleEntropy := filestatistics.SingleSymbolEntropy(io.Reader(f), 1)
	_, _ = f.Seek(0, 0)
	joinEntropy := filestatistics.SingleSymbolEntropy(io.Reader(f), 2)

	conditionalEntropy := joinEntropy - singleEntropy

	fmt.Println(conditionalEntropy)

}

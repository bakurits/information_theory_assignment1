package main

import (
	"assignment1/encoding"
	"io"
	"log"
	"os"
)

func main() {

	//var symbols = " აბგდევზთიკლმნოპჟრსტუფქღყშჩცძწჭხჯჰ"
	filename := "PublicTests/D/001.dat"
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

	encoding.HuffmanCodes(io.Reader(f), io.Writer(os.Stdout))
}

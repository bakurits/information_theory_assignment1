package main

import (
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

}

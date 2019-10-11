package main

import (
	"assignment1/filestatistics"
	"fmt"
)

func main() {

	var symbols = " აბგდევზთიკლმნოპჟრსტუფქღყშჩცძწჭხჯჰ"

	var mp = filestatistics.GetSubstringDistribution("PublicTests/A/001.dat", 1)

	for _, ch := range symbols {

		fmt.Printf("%c %.7f\n", ch, mp[string(ch)])
	}
}

package filestatistics

import (
	"bufio"
	"io"
	"log"
	"os"
	"strings"
)

func GetSubstringDistribution(filename string, substringLength int) map[string]float64 {
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
	r := bufio.NewReader(f)
	var st []rune
	st = []rune(strings.Repeat(" ", substringLength-1))
	var res = make(map[string]float64)

	var fileLen int64 = 0
	for {
		if ch, _, err := r.ReadRune(); err != nil {

			if err == io.EOF {
				break
			} else {
				log.Fatal(err)
				return nil
			}
		} else {
			fileLen++
			st = append(st, ch)
			if len(st) > substringLength {
				st = st[len(st)-substringLength:]
			}
			str := string(st)
			if _, ok := res[str]; !ok {
				res[str] = 0
			}
			res[str] ++
		}
	}

	for key, element := range res {
		res[key] = element / float64(fileLen)
	}

	return res
}

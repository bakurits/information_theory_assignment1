package filestatistics

import (
	"bufio"
	"io"
	"log"
	"strings"
)

func GetSubstringDistribution(reader io.Reader, substringLength int) map[string]float64 {
	r := bufio.NewReader(reader)
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
			res[str]++
		}
	}

	for key, element := range res {
		res[key] = element / float64(fileLen)
	}

	return res
}

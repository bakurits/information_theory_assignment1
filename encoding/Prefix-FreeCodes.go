package encoding

import (
	"fmt"
	"io"
	"math"
	"sort"
)

func checkCraftInequality(n int, l []int) bool {
	var res float64 = 0

	for i := 0; i < n; i++ {
		res += math.Pow(2, float64(-l[i]))
	}

	return res <= 1
}

func getNewCode(pre string, length int) string {
	var resBytes = []byte(pre)
	var shouldEnlarge = true
	for i := len(pre) - 1; i >= 0; i-- {
		if pre[i] == '0' {
			resBytes[i] = '1'
			shouldEnlarge = false
			break
		}
		resBytes[i] = '0'
	}
	var ans string
	if shouldEnlarge {
		ans = "1" + string(resBytes)
	} else {
		ans = string(resBytes)
	}
	for len(ans) < length {
		ans += "0"
	}
	return ans
}

func constructPrefixFreeCode(n int, l []int, w io.Writer) {
	type codeWithData struct {
		Index  int
		Length int
		Code   string
	}
	var data []codeWithData
	data = make([]codeWithData, n)
	for i := 1; i < n; i++ {
		data[i] = codeWithData{
			Index:  i,
			Length: l[i],
		}
	}

	sort.SliceStable(data, func(i, j int) bool {
		return data[i].Length < data[j].Length
	})

	data[0].Code = "0"

	for i := 1; i < n; i++ {
		data[i].Code = getNewCode(data[i-1].Code, data[i].Length)
	}

	sort.SliceStable(data, func(i, j int) bool {
		return data[i].Index < data[j].Index
	})

	for i := 0; i < n; i++ {
		_, _ = fmt.Fprint(w, data[i].Code)
	}
}

func PrefixFreeCodes(r io.Reader, w io.Writer) {
	var n int
	_, _ = fmt.Fscan(r, &n)

	var l []int
	l = make([]int, n)
	for i := 0; i < n; i++ {
		_, _ = fmt.Fscan(r, &l[i])
	}

	if !checkCraftInequality(n, l) {
		return
	}

	constructPrefixFreeCode(n, l, w)
}

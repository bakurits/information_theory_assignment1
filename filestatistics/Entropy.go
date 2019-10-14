package filestatistics

import (
	"io"
	"math"
)

func SingleSymbolEntropy(reader io.Reader, substringLength int) float64 {
	var dist = GetSubstringDistribution(reader, substringLength)
	var res float64 = 0
	for _, v := range dist {
		res -= math.Log2(v) * v
	}
	return res
}

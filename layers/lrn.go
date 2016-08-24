package layers

import (
	"fmt"
	"math"

	"github.com/gonum/matrix/mat64"
)

// LRN is Local Response Normalization.
type LRN struct {
	*BaseLayer
	N     int
	K     float64
	Alpha float64
	Beta  float64
}

// NewLRNLayer is constructor.
func NewLRNLayer(name, t string, n int, k, alpha, beta float64) *LRN {
	return &LRN{
		BaseLayer: NewBaseLayer(name, t),
		N:         n,
		K:         k,
		Alpha:     alpha,
		Beta:      beta,
	}
}

// Forward forwards one step of the network.
func (lrn *LRN) Forward(input [][][]float32) ([][][]float32, error) {
	output := make([][][]float32, len(input))
	for k := range input {
		s := int(math.Max(0.0, float64(k-lrn.N/2)))
		e := int(math.Min(float64(len(input)-1), float64(k+lrn.N/2)))
		o := ConvertMatrix(input[k])
		var res mat64.Dense
		res.Apply(func(i, j int, v float64) float64 {
			sum := 0.0
			for l := s; l < e; l++ {
				sum += float64(input[l][i][j] * input[l][i][j])
			}
			return float64(lrn.K) + lrn.Alpha*sum
		}, o)
		output[k] = ConvertMat64(&res)
	}
	fmt.Println(len(output), len(output[0]), len(output[0][0]))
	return output, nil
}

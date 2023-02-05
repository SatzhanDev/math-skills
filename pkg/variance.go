package pkg

import "math"

func Variance(arr []float64, average float64) float64 {
	var variance float64
	for i := 0; i < len(arr); i++ {
		variance += float64((arr[i] - average) * (arr[i] - average))
	}
	return math.Round(variance / float64(len(arr)))
}

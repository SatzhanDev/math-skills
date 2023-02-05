package pkg

import "math"

func StandardDeviation(arr []float64, average float64) float64 {
	var result float64
	var super_average float64
	var total float64
	for i := 0; i < len(arr); i++ {
		total += float64(arr[i])
	}
	super_average = math.Round(float64(total) / float64(len(arr)))
	result = float64(Variance(arr, float64(super_average)))
	return math.Sqrt(float64(result))
}

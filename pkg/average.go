package pkg

func Average(arr []float64) float64 {
	var sum float64 = 0
	for i := range arr {
		sum += arr[i]
	}
	return float64(sum) / float64(len(arr))
}

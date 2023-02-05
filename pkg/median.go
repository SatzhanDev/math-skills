package pkg

import "math"

func Median(s []float64) float64 {
	// Buble sort
	for i := 0; i < len(s)-1; i++ {
		for j := 0; j < len(s)-i-1; j++ {
			if s[j] > s[j+1] {
				s[j], s[j+1] = s[j+1], s[j]
			}
		}
	}
	// Median
	if len(s)%2 == 1 {
		index := (len(s) / 2)
		return float64(s[index])
	} else {
		index := len(s) / 2
		sum := math.Round(float64(s[index])+float64(s[index-1])) / 2
		return float64(sum)
	}
}

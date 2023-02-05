package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"

	"mathskill/pkg"
)

func main() {
	text_file := os.Args[1]
	data, _ := os.ReadFile(text_file)
	var number_string string
	number_string = string(data)
	// for i := 0; i < len(data); i++ {
	// 	if data[i] > 47 && data[i] < 58 {
	// 		number_string += string(data[i])
	// 	} else if data[i] == 10 && len(data)-1 != i {
	// 		number_string += " "
	// 	}
	// }
	strs := strings.Split(number_string, "\n")
	ary := make([]float64, len(strs))
	for i := range ary {
		ary[i], _ = strconv.ParseFloat(strs[i], 64)
	}

	fmt.Println("Average:", int(math.Round(pkg.Average(ary))))
	fmt.Println("Median:", int(math.Round(pkg.Median(ary))))
	fmt.Println("Variance:", int(math.Round(pkg.Variance(ary, pkg.Average(ary)))))
	fmt.Println("Standard Deviation:", int(math.Round(pkg.StandardDeviation(ary, pkg.Average(ary)))))
}

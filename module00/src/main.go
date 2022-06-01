package main

import (
	"bufio"
	"flag"
	"fmt"
	scanner2 "go/scanner"
	"math"
	"os"
	"sort"
	"strconv"
)

func isFlagsPassed(metrics map[string]string) bool {
	res := false
	for m := range metrics {
		if res = isFlagPassed(m); res {
			break
		}
	}
	return res
}

func isFlagPassed(name string) bool {
	found := false
	flag.Visit(func(f *flag.Flag) {
		if f.Name == name {
			found = true
		}
	})
	return found
}

func parseNumbers() []float64 {

	var numbers []float64
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		input := scanner.Text()
		num, err := strconv.Atoi(input)
		if err != nil {
			scanner2.PrintError(os.Stderr, err)
			os.Exit(1)
		}
		if num < -100000 || num > 100000 {
			scanner2.PrintError(os.Stderr, fmt.Errorf("error: not a valid number"))
			os.Exit(1)
		}
		numbers = append(numbers, float64(num))
	}
	if numbers == nil {
		scanner2.PrintError(os.Stderr, fmt.Errorf("error: sequence of numbers is not set"))
		os.Exit(1)
	}
	return numbers
}

func mean(numbers []float64) float64 {
	var sumRes float64 = 0
	for _, n := range numbers {
		sumRes += n
	}
	res := sumRes / float64(len(numbers))
	return math.Round(res*100) / 100
}

func median(numbers []float64) float64 {
	sort.Float64s(numbers)
	var res float64
	if len(numbers)%2 == 0 {
		i := len(numbers)/2 - 1
		res = (numbers[i] + numbers[i+1]) / 2
	} else {
		res = numbers[len(numbers)/2]
	}
	return math.Round(res*100) / 100
}

func mode(numbers []float64) float64 {
	m := map[float64]float64{}
	var maxCount float64
	var res float64

	sort.Float64s(numbers)
	for _, a := range numbers {
		m[a]++
		if m[a] > maxCount {
			maxCount = m[a]
			res = a
		}
	}
	return math.Round(res*100) / 100
}

func sd(numbers []float64) float64 {
	total := 0.0
	meanRes := mean(numbers)
	for _, num := range numbers {
		total += math.Pow(num-meanRes, 2)
	}
	res := math.Sqrt(total / float64(len(numbers)-1))
	return math.Round(res*100) / 100
}

func outputAllMetrics(numbers []float64) {
	resMean := mean(numbers)
	fmt.Printf("Mean: %s\n", strconv.FormatFloat(resMean, 'f', -1, 64))

	resMedian := median(numbers)
	fmt.Printf("Median: %s\n", strconv.FormatFloat(resMedian, 'f', -1, 64))

	resMode := mode(numbers)
	fmt.Printf("Mode: %s\n", strconv.FormatFloat(resMode, 'f', -1, 64))

	resSD := sd(numbers)
	fmt.Printf("SD: %s\n", strconv.FormatFloat(resSD, 'f', -1, 64))
}

func main() {

	metrics := map[string]string{
		"mean":   "output the average of a sequence of numbers",
		"median": "output \"middle\" of a sorted sequence of numbers.",
		"mode":   "output the number that occurs most often of a sequence of numbers.\nif numbers are several, the smallest one among those is returned.",
		"sd":     "output the number of variations or  dispersion of a sorted sequence of numbers",
	}

	for m, info := range metrics {
		flag.Bool(m, false, info)
	}
	flag.Parse()
	numbers := parseNumbers()

	if !isFlagsPassed(metrics) {
		outputAllMetrics(numbers)
		return
	}
	if isFlagPassed("mean") {
		res := mean(numbers)
		fmt.Printf("Mean: %s\n", strconv.FormatFloat(res, 'f', -1, 64))
	}
	if isFlagPassed("median") {
		res := median(numbers)
		fmt.Printf("Median: %s\n", strconv.FormatFloat(res, 'f', -1, 64))
	}
	if isFlagPassed("mode") {
		res := mode(numbers)
		fmt.Printf("Mode: %d\n", res)
	}
	if isFlagPassed("sd") {
		res := sd(numbers)
		fmt.Printf("SD: %s\n", strconv.FormatFloat(res, 'f', -1, 64))
	}
}

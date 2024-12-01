package main

import (
	"fmt"
	"math"
	"os"
	"regexp"
	"sort"
	"strconv"
)

func main() {
	inputBytes, err := os.ReadFile("./problem1/input1.txt")
	if err != nil {
		panic(err)
	}
	str := string(inputBytes) // convert content to a 'string'
	var numbersRegex = regexp.MustCompile(`(?m)^(?<first>[0-9]+)\s+(?<second>[0-9]+)$`)
	res := numbersRegex.FindAllStringSubmatch(str, -1)

	firstList := make([]int64, len(res))
	secondList := make([]int64, len(res))
	for i := range res {
		match := res[i]
		if len(match) != 3 {
			continue
		}
		first, _ := strconv.ParseInt(match[1], 10, 64)
		firstList[i] = first
		second, _ := strconv.ParseInt(match[2], 10, 64)
		secondList[i] = second

	}
	sort.Slice(firstList, func(i, j int) bool { return firstList[i] < firstList[j] })
	sort.Slice(secondList, func(i, j int) bool { return secondList[i] < secondList[j] })

	result := 0.0
	for i := range firstList {
		diff := firstList[i] - secondList[i]
		result += math.Abs(float64(diff))
	}

	fmt.Printf("%.0f\n", result)
}

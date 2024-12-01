package main

import (
	"fmt"
	"os"
	"regexp"
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
	secondMap := make(map[int64]int64)
	for i := range res {
		match := res[i]
		if len(match) != 3 {
			continue
		}
		first, _ := strconv.ParseInt(match[1], 10, 64)
		firstList[i] = first
		second, _ := strconv.ParseInt(match[2], 10, 64)

		no, ok := secondMap[second]
		if ok {
			secondMap[second] = no + 1
		} else {
			secondMap[second] = 1
		}

	}

	var result int64 = 0
	for i := range firstList {
		elem := firstList[i]
		no, ok := secondMap[elem]
		if ok {
			result += elem * no
		}
	}

	fmt.Printf("%d\n", result)
}

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	inputNums := []int64{}
	scanner := bufio.NewScanner(os.Stdin)
	for i := 0; i < 10; i++ {
		if scanner.Scan() {
			line := scanner.Text()
			num, err := strconv.ParseInt(line, 10, 64)
			if err != nil {
				panic(err)
			}
			inputNums = append(inputNums, num)
		}
		length := len(inputNums)
		for i := 1; i < length; i++ {
			for j := i; j > 0; j-- {
				if inputNums[j] < inputNums[j-1] {
					inputNums[j], inputNums[j-1] = inputNums[j-1], inputNums[j]
				}
			}
		}
	}
	fmt.Println(inputNums)
}
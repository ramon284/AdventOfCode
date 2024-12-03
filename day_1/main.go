package main

import (
	"fmt"
	"log"
	"os"
	"slices"
	"strconv"
	"strings"
	"time"
)

func SplitList(full_list []int) ([]int, []int) {
	even_list := []int{}
	odd_list := []int{}
	for index, element := range full_list {
		if index%2 == 0 {
			even_list = append(even_list, element)
		} else {
			odd_list = append(odd_list, element)
		}
	}
	return even_list, odd_list
}

func AbsoluteDistance(a, b int) int {
	if a > b {
		return a - b
	}
	return b - a
}

func distanceList(even_list []int, odd_list []int) []int {
	temp_list := []int{}
	for i := range even_list {
		temp_list = append(temp_list, AbsoluteDistance(even_list[i], odd_list[i]))
	}
	return temp_list
}

func similarity_score(even_list []int, odd_list []int) int {
	score := 0
	for _, el := range even_list {
		for _, el2 := range odd_list {
			if el == el2 {
				score += el
			}
		}
	}
	return score
}

func SumList(numbers []int) int {
	sum := 0
	for _, num := range numbers {
		sum += num
	}
	return sum
}

func main() {
	// read the file
	content, err := os.ReadFile("puzzle_input_text.txt")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	} // as strings into a slice
	numberStrings := strings.Fields(string(content))
	numbers := []int{}

	for _, str := range numberStrings { // cast slice of strings to slice of ints
		num, err := strconv.Atoi(str)
		if err != nil {
			fmt.Printf("Error converting %q to integer: %v\n", str, err)
			return
		}
		numbers = append(numbers, num)
	}

	start := time.Now()
	evenList, oddList := SplitList(numbers)
	slices.Sort(evenList)
	slices.Sort(oddList)
	distance_list := distanceList(evenList, oddList)
	method_2 := SumList(distance_list)
	fmt.Println(method_2)

	new_score := similarity_score(evenList, oddList)
	log.Println(new_score)
	elapsed := time.Since(start)
	log.Printf("Binomial took %.4f ms", float64(elapsed.Nanoseconds())/1e6) // Milliseconds

}

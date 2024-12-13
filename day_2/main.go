package main

import (
	"AdventOfCode/benchmark"
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

// Step 1: check if the list == sort  OR equals sort(reverse)

func main() {
	file, err := os.Open("puzzle_input_text.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	var full_list [][]int
	for scanner.Scan() {
		line := scanner.Text()
		split_lines := strings.Fields(line) // turns string into list/slices, split on whitespace
		var row []int                       // empty row list
		for _, str := range split_lines {   // convert str to int, put them in row
			num, err := strconv.Atoi(str)
			if err != nil {
				return
			}
			row = append(row, num)
		}
		full_list = append(full_list, row) // put the row into the final list

	}
	timer := benchmark.Start()
	var count int
	var count_with_dampener int
	for _, row := range full_list {
		count = count + checkList(row, false)
		count_with_dampener += checkList(row, true)
	}
	fmt.Println(count)
	fmt.Println(count_with_dampener)
	timer.PrintElapsed() // 537 microseconds, 0.5ms
}

// This function loops through every iteration of an "incorrect level"
// And then checks it again with the checksList() function
// This function is called inside checksList() when a problem is encountered
func check_alternatives(row []int) bool {
	for i, _ := range row {
		row_copy := deepCopy(row)
		row_copy = append(row_copy[:i], row_copy[i+1:]...)
		succesful := checkList(row_copy, false)
		if succesful == 1 {
			return true
		}
	}
	return false
}

// This function checks for 2 conditions:
//   - Is the list sorted? Either ascending or descending
//   - Is the increase/decrease from number to number larger than 1 and leq than 4?
func checkList(row []int, dampener bool) int {
	ascending := false
	descending := false
	previous_num := 0
	for i, num := range row {
		if i == 0 {
			previous_num = num
		} else if previous_num > num {
			descending = true
			abs_distance := AbsoluteDistance(previous_num, num)
			if abs_distance < 1 || abs_distance > 3 {
				if dampener == true {
					result := check_alternatives(row)
					if result == true {
						return 1
					}
				}
				return 0
			}
			previous_num = num
		} else if num > previous_num {
			ascending = true
			abs_distance := AbsoluteDistance(previous_num, num)
			if abs_distance < 1 || abs_distance > 3 {
				if dampener == true {
					result := check_alternatives(row)
					if result == true {
						return 1
					}
				}
				return 0
			}
			previous_num = num
		} else if num == previous_num {
			if dampener == true {
				result := check_alternatives(row)
				if result == true {
					return 1
				}
			}
			return 0
		}
		if ascending == true && descending == true {
			if dampener == true {
				result := check_alternatives(row)
				if result == true {
					return 1
				}
			}
			return 0
		}
	}
	return 1
}

func deepCopy(slice []int) []int { // makes a new slice of same len() as input, and copies the values
	copySlice := make([]int, len(slice))
	copy(copySlice, slice)
	return copySlice
}

// Checks the absolute distance
func AbsoluteDistance(a, b int) int {
	if a > b {
		return a - b
	}
	return b - a
}

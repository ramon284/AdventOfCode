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

func main() {
	first_list, second_list := retrieve_lists("puzzle_input_text.txt")
	timer := benchmark.Start()
	correct_combinations := check_pages_to_produce(first_list, second_list)
	final_number := count_middle_numbers(correct_combinations)
	fmt.Println(final_number)
	timer.PrintElapsed() // 41 milliseconds
	timer = benchmark.Start()
	correct_combinations_2 := check_and_reorder_pages(first_list, second_list)
	final_number_2 := count_middle_numbers(correct_combinations_2)
	fmt.Println(final_number_2)
	timer.PrintElapsed() // 159 milliseconds
}

func count_middle_numbers(correct_combinations [][]string) int {
	counter := 0
	for _, x := range correct_combinations {
		len_of_x := len(x)
		number, _ := strconv.Atoi(x[len_of_x/2])
		counter += number
	}
	return counter
}

func check_order(first_list []string, first_char string, second_char string) bool {
	for _, combination := range first_list {
		left := combination[:2]
		right := combination[3:5]
		if second_char == left && first_char == right {
			return false
		}
	}
	return true
}

func check_and_reorder_pages(first_list []string, second_list [][]string) [][]string {
	var correct_rows [][]string
	for _, x := range second_list {
		reordered := false // track when a list should be returned or not
		mistake_discovered := true
		for mistake_discovered == true { // keep sorting the list until it's 100% correct
			mistake_discovered = false
			valid := true
			row_len := len(x)
			for i, _ := range x {
				for j := 1; j < row_len-i; j++ {
					valid = check_order(first_list, x[i], x[j+i])
					if valid == false {
						reordered = true          // put this list in the return function
						mistake_discovered = true // re-run the full list again later
						number_to_move := x[j+i]
						x = append(x[:j+i], x[j+i+1:]...)
						x = append(x[:i], append([]string{number_to_move}, x[i:]...)...)
					}
				}
			}
		}
		if reordered == true {
			correct_rows = append(correct_rows, x)
		}
	}
	return correct_rows
}

func check_pages_to_produce(first_list []string, second_list [][]string) [][]string {
	var correct_rows [][]string
	for _, x := range second_list {
		valid := true
		row_len := len(x)
		for i, _ := range x {
			for j := 1; j < row_len-i; j++ {
				valid = check_order(first_list, x[i], x[j+i])
				if valid == false {
					break
				}
			}
			if valid == false {
				break
			}
		}
		if valid == true {
			correct_rows = append(correct_rows, x)
		}
	}
	return correct_rows
}

func retrieve_lists(path string) ([]string, [][]string) {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	var first_list []string
	var second_list [][]string
	first_list_bool := true
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		}
		if len(line) != 5 {
			first_list_bool = false
		}
		if first_list_bool == true {
			first_list = append(first_list, line)
		} else {
			split_lines := strings.Split(line, ",")
			second_list = append(second_list, split_lines)
		}
	}
	return first_list, second_list
}

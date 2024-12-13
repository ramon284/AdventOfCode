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

type numbers_struct struct {
	first_value  int
	second_value []int
}

type numbers_string_struct struct {
	first_value  string
	second_value []string
}

func main() {
	test := load_nrs("puzzle_input_text.txt")
	timer := benchmark.Start()
	correct, incorrect_assigments := assignment_1(test)
	var final_count_1 int
	for _, v := range correct {
		final_count_1 += v.first_value
	}
	fmt.Println(final_count_1)
	timer.PrintElapsed() // 6.3 milliseconds
	timer = benchmark.Start()
	correct_with_pipes := assignment_2(incorrect_assigments)
	var final_count_2 int
	for _, v := range correct_with_pipes {
		final_count_2 += v.first_value
	}
	fmt.Println(final_count_1, "+", final_count_2, "=", final_count_1+final_count_2)
	timer.PrintElapsed() // 217 milliseconds
}

func assignment_2(assignments []numbers_struct) []numbers_struct {
	var correct_equations []numbers_struct
	for _, v := range assignments {
		x := recursive(v.first_value, v.second_value[0], v.second_value[1:], false, true)
		if x {
			correct_equations = append(correct_equations, v)
		}
	}
	return correct_equations
}

func assignment_1(numbers []numbers_struct) ([]numbers_struct, []numbers_struct) {
	var correct_equations []numbers_struct
	var incorrect_equations []numbers_struct
	for _, v := range numbers {
		x := recursive(v.first_value, v.second_value[0], v.second_value[1:], false, false)
		if x {
			correct_equations = append(correct_equations, v)
		} else {
			incorrect_equations = append(incorrect_equations, v)
		}
	}
	return correct_equations, incorrect_equations
}

func recursive(solution_nr int, first_number int, remaining_numbers []int,
	got_it bool, piping bool) bool {
	temp_number := first_number
	if len(remaining_numbers) == 1 {
		if temp_number+remaining_numbers[0] == solution_nr {
			return true // DIT IS TRUE
		} else if temp_number*remaining_numbers[0] == solution_nr {
			return true // DIT IS OOK TRUE
		}
		temp_number, _ = strconv.Atoi(strconv.FormatInt(int64(temp_number), 10) +
			strconv.FormatInt(int64(remaining_numbers[0]), 10))
		if temp_number == solution_nr {
			return true
		} else {
			return false
		}
	}
	next_number := remaining_numbers[0]
	new_number := temp_number + next_number
	correct := recursive(solution_nr, new_number, remaining_numbers[1:], false, piping)
	if correct == false {
		new_number = temp_number * next_number
		correct = recursive(solution_nr, new_number, remaining_numbers[1:], false, piping)
	}
	if correct == false && piping == true {
		new_number, _ = strconv.Atoi(strconv.FormatInt(int64(temp_number), 10) +
			strconv.FormatInt(int64(next_number), 10))
		correct = recursive(solution_nr, new_number, remaining_numbers[1:], false, piping)
	}
	return correct

}

func load_nrs(path string) []numbers_struct {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	nums := []numbers_struct{}
	for scanner.Scan() {
		line := scanner.Text()
		split_lines := strings.Fields(line)
		var temp numbers_struct
		var temp_list []int
		var first_value int
		for _, str := range split_lines {
			length := len(str)
			if str[length-1] == ':' {
				first_value, _ = strconv.Atoi(str[:length-1])
			} else {
				second_value, _ := strconv.Atoi(str)
				temp_list = append(temp_list, second_value)
			}
		}
		temp = numbers_struct{first_value, temp_list}
		nums = append(nums, temp)
	}
	return nums
}

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
	file, err := os.Open("puzzle_input_text.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	var full_text []string
	i := 0
	for scanner.Scan() {
		line := scanner.Text()
		full_text = append(full_text, line)
		i += 1
	}
	full_string := strings.Join(full_text, ",")

	// Challenge #1
	timer := benchmark.Start()
	output_list := findMuls(full_string, false)
	total_number := 0
	for _, row := range output_list {
		multiplied_nr := multiplyNumbers(row)
		total_number += multiplied_nr
	}
	fmt.Println(total_number)
	timer.PrintElapsed() // 30 milliseconds
	timer = benchmark.Start()

	// Challenge #2
	output_list2 := findMuls(full_string, true)
	total_number2 := 0
	for _, row := range output_list2 {
		multiplied_nr2 := multiplyNumbers(row)
		total_number2 += multiplied_nr2
	}
	fmt.Println(total_number2)
	timer.PrintElapsed() // 26 milliseconds
}

func findMuls(string_list string, bonus_assignment bool) [][]string { // returns a list containing parts to be multiplied
	var output_list [][]string
	var left_num string
	var right_num string
	var checking_right_num bool = false
	var inside_mul bool = false
	var looped_over string = "" // make it 7 strings long, to account
	var do bool = true
	for i, str := range string_list {
		if bonus_assignment == true {
			if checkStr(looped_over, "do()", i) {
				do = true
			} else if checkStr(looped_over, "don't()", i) {
				do = false
			}
		}
		if do == false {
			looped_over += string(str)
			continue
		}
		if inside_mul {
			//fmt.Println("yoooo")
			isnr := checkIfNumber(str)
			if isnr {
				if checking_right_num == false {
					left_num += string(str)
				} else {
					right_num += string(str)
				}
			} else if str == ',' {
				checking_right_num = true
			} else if str == ')' { //succesfully captured
				inside_mul = false
				if checking_right_num == true { // succes !
					checking_right_num = false
					output_list = append(output_list, []string{left_num, right_num})
					left_num = ""
					right_num = ""
				} else {
					inside_mul = false
					left_num = ""
					right_num = ""
				}
			} else {
				checking_right_num = false
				inside_mul = false
				left_num = ""
				right_num = ""
			}
		}

		if checkStr(looped_over, "mul", i) && str == '(' {
			//if looped_over[i+4:i+7] == "mul" && str == '(' {
			inside_mul = true
		}
		looped_over += string(str)
	}
	return output_list
}

func checkStr(string_to_check string, substring_to_check string, idx int) bool {
	required_len := len(substring_to_check)
	if required_len > len(string_to_check) {
		return false
	}
	if string_to_check[idx-required_len:idx] == substring_to_check {
		return true
	}
	return false
}

func checkIfNumber(char rune) bool {
	return char >= '0' && char <= '9'
}

func multiplyNumbers(row []string) int {
	nr1, _ := strconv.Atoi(row[0])
	nr2, _ := strconv.Atoi(row[1])
	return nr1 * nr2
}

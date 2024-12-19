package main

import (
	"AdventOfCode/benchmark"
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() { // bro wtf
	file := read_file("puzzle_input_text.txt")
	timer := benchmark.Start()
	some_array := get_files_array(file)
	ordered_array := order_files_array(some_array)
	final_output := get_checksum(ordered_array, 1)
	fmt.Println(final_output)
	timer.PrintElapsed() // 333 milliseconds
	timer = benchmark.Start()
	size_map := create_size_mapping(some_array)
	ordered_array_2 := order_files_array_part_2(some_array, size_map)
	final_output_2 := get_checksum(ordered_array_2, 2)
	fmt.Println(final_output_2)
	timer.PrintElapsed() // 1.17 seconds
}

func order_files_array_part_2(some_array []string, filesize_map map[string]int) []string {
	array_copy := deepCopy(some_array)
	for i := len(array_copy) - 1; i > 0; i-- {
		if array_copy[i] != "." {
			var value = array_copy[i]
			value_size := filesize_map[value]
		forward_loop:
			for j := 0; j < len(array_copy); j++ {
				if array_copy[j] == "." { // we found the dot baby
					if i < j {
						break
					}
					if value_size == 1 {
						array_copy[j] = value
						array_copy[i] = "."
						break
					}
					// indexes are implicit, its j in range value_size and i in range value_size
					for h := 0; h < value_size; h++ {
						if array_copy[j+h] != "." { // not enough space
							j += h
							continue forward_loop
						}
					}
					for h := 0; h < value_size; h++ { // if there IS enough space :
						array_copy[j+h] = value
						array_copy[i-h] = "."
					}
					i -= value_size - 1
					break
				}
			}
		}
	}
	return array_copy
}

func order_files_array(some_array []string) []string {
	array_copy := deepCopy(some_array)
	for i := len(array_copy) - 1; i > 0; i-- { // loop backwards until we find a number
		if array_copy[i] != "." {
			var value string = array_copy[i]
			for j := 0; j < len(array_copy); j++ { // loop forwards until we find a dot
				if array_copy[j] == "." { // now, put "value" in [j] index, and put "." at [i] index
					if i < j {
						return array_copy
					}
					array_copy[j] = value
					array_copy[i] = "."
					break
				}
			}
		}
	}
	return array_copy
}

func create_size_mapping(some_array []string) map[string]int {
	// this function is kind of a mess, but it works
	size := 0
	id := ""
	mymap := make(map[string]int)
	var previous_x = some_array[0]
	for _, x := range some_array {
		if x != "." {
			if x != previous_x && previous_x != "." {
				mymap[id] = size
				size = 0
			}
			id = x
			size += 1
		} else {
			if size != 0 {
				mymap[id] = size
			}
			size = 0
		}
		previous_x = x
	}
	mymap[id] = size
	return mymap
}

func get_checksum(ordered_array []string, assignment int) int {
	var final_output int
	for i, value := range ordered_array {
		if value == "." {
			if assignment == 1 {
				return final_output
			} else if assignment == 2 {
				continue
			}
		}
		int_value, _ := strconv.Atoi(value)
		final_output += i * int_value
	}
	return final_output
}

func get_files_array(file string) []string {
	ID := 0
	var some_array []string
	for i, x := range file {
		if i%2 == 0 { // files
			for j := 0; j < int(x-'0'); j++ {
				some_array = append(some_array, strconv.FormatInt(int64(ID), 10))
			}
			ID += 1
		} else { // Empty blocks
			for j := 0; j < int(x-'0'); j++ {
				some_array = append(some_array, ".")
			}
		}
	}
	return some_array
}

func deepCopy(slice []string) []string { // makes a new slice of same len() as input, and copies the values
	copySlice := make([]string, len(slice))
	copy(copySlice, slice)
	return copySlice
}

func read_file(path string) string {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	var number_string string
	for scanner.Scan() {
		line := scanner.Text()
		number_string += line
	}
	return number_string
}

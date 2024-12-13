package main

import (
	"AdventOfCode/benchmark"
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

type coordinates struct {
	antenna      string
	x_coordinate int
	y_coordinate int
}

func main() {
	input := read_file("puzzle_input_text.txt")
	//input := read_file("test.txt")
	//input := read_file("test2.txt")
	timer := benchmark.Start()
	output_map := assignment(input, false)
	fmt.Println(len(output_map))
	timer.PrintElapsed() // 507 microseconds, 0.5ms
	timer = benchmark.Start()
	output_map_2 := assignment(input, true)
	fmt.Println(len(output_map_2))
	timer.PrintElapsed() // 511 microseconds, 0.5ms
}

func assignment(input []string, assignment_2 bool) map[string]string {
	all_coordinates := []coordinates{}
	for i, row := range input {
		for j, cell := range row {
			if cell != '.' {
				all_coordinates = append(all_coordinates, coordinates{string(cell), i, j})
			}
		}
	}
	width := len(input[0])
	height := len(input)
	mymap := make(map[string]string)
	for _, coordinate := range all_coordinates { // select an antenna tower
		current_antenna := coordinate.antenna
		current_x := coordinate.x_coordinate
		current_y := coordinate.y_coordinate
		for _, other := range all_coordinates { // compare to other antenna towers
			if current_antenna == other.antenna { // same signal?
				if current_x == other.x_coordinate && current_y == other.y_coordinate {
					continue
				}
				if assignment_2 {
					stringified_coordinates := strconv.FormatInt(int64(current_x), 10) + "|" + strconv.FormatInt(int64(current_y), 10)
					mymap[stringified_coordinates] = stringified_coordinates // every antenna should also be a signal spot in assignment_2
				}
				compare_col_loc, compare_row_loc := compare_locations(current_x, current_y, other.x_coordinate, other.y_coordinate)
				new_x, new_y := calc_new_coordinates(current_x, current_y, other.x_coordinate, other.y_coordinate, compare_col_loc, compare_row_loc)
				if new_x < 0 || new_x >= width || new_y < 0 || new_y >= height { // new coord is out of bounds
					continue
				}
				stringified_coordinates := strconv.FormatInt(int64(new_x), 10) + "|" + strconv.FormatInt(int64(new_y), 10)
				mymap[stringified_coordinates] = stringified_coordinates
				if assignment_2 == true {
					distance_x := AbsoluteDistance(current_x, other.x_coordinate)
					distance_y := AbsoluteDistance(current_y, other.y_coordinate)
					for true {
						if compare_col_loc == "left" {
							new_x = new_x + distance_x
						} else if compare_col_loc == "right" {
							new_x = new_x - distance_x
						}
						if compare_row_loc == "below" {
							new_y = new_y - distance_y
						} else if compare_row_loc == "above" {
							new_y = new_y + distance_y
						}
						if new_x < 0 || new_x >= width || new_y < 0 || new_y >= height { // new coord is out of bounds
							break
						}
						stringified_coordinates = strconv.FormatInt(int64(new_x), 10) + "|" + strconv.FormatInt(int64(new_y), 10)
						mymap[stringified_coordinates] = stringified_coordinates
					}
				}
			}
		}
	}
	return mymap
}

func AbsoluteDistance(a, b int) int {
	if a > b {
		return a - b
	}
	return b - a
}

func compare_locations(original_x int, original_y int, other_x int, other_y int) (string, string) {
	compare_col_loc := "same"
	compare_row_loc := "same"
	if original_x < other_x {
		compare_col_loc = "left"
	} else if original_x > other_x {
		compare_col_loc = "right"
	}
	if original_y > other_y {
		compare_row_loc = "below"
	} else if original_y < other_y {
		compare_row_loc = "above"
	}
	return compare_col_loc, compare_row_loc
}

func calc_new_coordinates(current_x int, current_y int, other_x int, other_y int, compare_col_loc string, compare_row_loc string) (int, int) {
	new_x := current_x
	new_y := current_y
	if compare_col_loc == "left" {
		new_x = other_x + AbsoluteDistance(current_x, other_x)
	} else if compare_col_loc == "right" {
		new_x = other_x - AbsoluteDistance(current_x, other_x)
	}
	if compare_row_loc == "below" {
		new_y = other_y - AbsoluteDistance(current_y, other_y)
	} else if compare_row_loc == "above" {
		new_y = other_y + AbsoluteDistance(current_y, other_y)
	}
	return new_x, new_y
}

func read_file(path string) []string {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	var first_list []string
	for scanner.Scan() {
		line := scanner.Text()
		first_list = append(first_list, line)
	}
	return first_list
}

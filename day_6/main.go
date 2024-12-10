package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

// step 1: find the ^ ( or v, <, > )
// step 2: make a pivoted copy, navigate
// step 3: profit

func main() {
	full_list := load_into_slice("puzzle_input_text.txt")
	assignment_1_solution, full_path := assignment_1(full_list)
	fmt.Println(assignment_1_solution)
	fmt.Println(full_path)
	full_list = load_into_slice("puzzle_input_text.txt")
	assignment_2_solution := assignment_2(full_list, full_path)
	fmt.Println(assignment_2_solution)
}

func check_if_meets_path(i int, j int, path []areas) bool {
	for _, coordinate := range path {
		if i == coordinate.i && j == coordinate.j {
			return true
		}
	}
	return false
}

func assignment_2(full_list []string, full_coordinates []areas) int {
	revisit_counts := 0
	for i_obstacle, row := range full_list {
		for j_obstacle, _ := range row {
			if check_if_meets_path(i_obstacle, j_obstacle, full_coordinates) == false {
				continue
			}
			full_list_copy := deepCopy(full_list)
			var next_symbol string
			if string(full_list_copy[i_obstacle][j_obstacle]) == "^" {
				continue
			}
			full_list_copy = add_obstacle(full_list_copy, i_obstacle, j_obstacle)
			testing_obstacle := true
			i, j, dir := find_initial_location(full_list)
			locs := []earlier_locations{}
			for testing_obstacle == true {
				full_list_copy = move_step(full_list_copy, i, j)
				next_symbol = check_next_symbol(full_list_copy, i, j, dir)
				revisited := false
				if next_symbol == "." {
					i, j = next_i_and_j(i, j, dir)
				} else if next_symbol == "#" {
					revisited = check_if_revisit(locs, i, j, dir)
					if revisited == true {
						revisit_counts += 1
						break
					}
					locs = append(locs, earlier_locations{i, j, dir})
					dir = rotate_direction(dir)
					checking := true
					for checking == true {
						if check_double_walls(full_list_copy, i, j, dir) == false {
							break
						}
						dir = rotate_direction(dir)
					}
					i, j = next_i_and_j(i, j, dir)
				} else if next_symbol == "o" {
					revisited = check_if_revisit(locs, i, j, dir)
					if revisited == true {
						revisit_counts += 1
						break
					}
					locs = append(locs, earlier_locations{i, j, dir})
					dir = rotate_direction(dir)
					checking := true
					for checking == true {
						if check_double_walls(full_list_copy, i, j, dir) == false {
							break
						}
						dir = rotate_direction(dir)
					}
					i, j = next_i_and_j(i, j, dir)
				} else if next_symbol == "X" {
					i, j = next_i_and_j(i, j, dir)
				} else if next_symbol == "finished" {
					testing_obstacle = false
				}
			}
		}
	}
	return revisit_counts
}

func check_double_walls(full_list []string, i int, j int, dir string) bool {
	temp_symbol := check_next_symbol(full_list, i, j, dir)
	if temp_symbol == "#" || temp_symbol == "o" || temp_symbol == "finished" {
		return true
	}
	return false
}

type earlier_locations struct { // Which coordinates did we visit, in which direction?
	i   int
	j   int
	dir string
}

func check_if_revisit(locations []earlier_locations, i int, j int, dir string) bool {
	for _, location := range locations {
		if location.i == i && location.j == j && location.dir == dir {
			return true
		}
	}
	return false
}

type areas struct {
	i int
	j int
}

func assignment_1(full_list []string) (int, []areas) {
	still_going := true
	i, j, dir := find_initial_location(full_list)
	areas_slice := []areas{}
	var next_symbol string
	full_list_copy := deepCopy(full_list)
	for still_going == true {
		full_list_copy = move_step(full_list_copy, i, j)
		next_symbol = check_next_symbol(full_list_copy, i, j, dir)
		if next_symbol == "." {
			i, j = next_i_and_j(i, j, dir)
		} else if next_symbol == "#" {
			dir = rotate_direction(dir)
			i, j = next_i_and_j(i, j, dir)
		} else if next_symbol == "X" {
			i, j = next_i_and_j(i, j, dir)
		} else if next_symbol == "finished" {
			still_going = false
		}
	}
	counter := 0
	for i, row := range full_list_copy {
		for j, cell := range row {
			if string(cell) == "X" {
				areas_slice = append(areas_slice, areas{i, j})
				counter += 1
			}
		}
	}
	return counter, areas_slice
}

//func check_if_valid(full_list []string, i int, j int) bool {
//
//}

func rotate_direction(direction string) string {
	if direction == "^" {
		return ">"
	} else if direction == ">" {
		return "v"
	} else if direction == "v" {
		return "<"
	} else {
		return "^"
	}
}

func next_i_and_j(i int, j int, direction string) (int, int) {
	if direction == "^" {
		return i - 1, j
	} else if direction == "v" {
		return i + 1, j
	} else if direction == "<" {
		return i, j - 1
	} else {
		return i, j + 1
	}
}

func add_obstacle(full_list []string, i int, j int) []string {
	row := []rune(full_list[i])
	row[j] = 'o'
	full_list[i] = string(row)
	return full_list
}

func move_step(full_list []string, i int, j int) []string {
	row := []rune(full_list[i])
	row[j] = 'X'
	full_list[i] = string(row)
	return full_list
}

func find_initial_location(full_list []string) (int, int, string) {
	for i, row := range full_list {
		for j, _ := range row {
			if string(row[j]) != "." && string(row[j]) != "#" && string(row[j]) != "X" && string(row[j]) != "o" {
				return i, j, string(row[j])
			}
		}
	}
	return 0, 0, ""
}

func check_next_symbol(full_list []string, i int, j int, direction string) string {
	if direction == "^" {
		if i == 0 {
			return "finished"
		}
		return string(full_list[i-1][j])
	} else if direction == "v" {
		if i+1 == len(full_list) {
			return "finished"
		}
		return string(full_list[i+1][j])
	} else if direction == "<" {
		if j == 0 {
			return "finished"
		}
		return string(full_list[i][j-1])
	} else if direction == ">" {
		if j+1 == len(full_list[0]) {
			return "finished"
		}
		return string(full_list[i][j+1])
	}
	return "If you see this, call 911"
}

func load_into_slice(path string) []string {
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

func deepCopy(slice []string) []string { // makes a new slice of same len() as input, and copies the values
	copySlice := make([]string, len(slice))
	copy(copySlice, slice)
	return copySlice
}

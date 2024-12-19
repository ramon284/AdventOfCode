package main

import (
	"AdventOfCode/benchmark"
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	input := read_file("puzzle_input_text.txt")
	t := benchmark.Start()
	mappie := make(map[string]int)
	mappie2 := make(map[string]int)
	zero_counter := 0
	for i, row := range input {
		for j, v := range row {
			if v == '0' {
				zero_counter += 1
				visited_9 := map[string]int{}
				idx := strconv.FormatInt(int64(i), 10) + "+" + strconv.FormatInt(int64(j), 10)
				paths_found, visited_9 := check_surrounding_nrs(input, i, j, int(v-'0'), idx, visited_9)
				if paths_found > 0 {
					mappie[idx] = paths_found
					mappie2[idx] = 0
					for _, value := range visited_9 {
						mappie2[idx] += value
					}
				}
			}
		}
	}
	number := 0
	for _, value := range mappie { // :(
		number += value
	}
	fmt.Println(number)
	number = 0
	for _, value := range mappie2 {
		number += value
	}
	fmt.Println(number)
	t.PrintElapsed() // 510 microseconds
}

func check_surrounding_nrs(input []string, i int, j int, value int, start_idx string, visited_9 map[string]int) (int, map[string]int) {
	len_row := len(input[0])
	no_of_rows := len(input)
	if j < no_of_rows-1 {
		if int(input[i][j+1]-'0') == value+1 {
			if value+1 == 9 {
				_, exists := visited_9[strconv.FormatInt(int64(i), 10)+strconv.FormatInt(int64(j+1), 10)]
				if exists == false {
					visited_9[strconv.FormatInt(int64(i), 10)+strconv.FormatInt(int64(j+1), 10)] = 1
				} else {
					visited_9[strconv.FormatInt(int64(i), 10)+strconv.FormatInt(int64(j+1), 10)] += 1
				}
			} else {
				_, visited_9 = check_surrounding_nrs(input, i, j+1, value+1, start_idx, visited_9)
			}
		}
	}
	if i < len_row-1 {
		if int(input[i+1][j]-'0') == value+1 {
			if value+1 == 9 {
				_, exists := visited_9[strconv.FormatInt(int64(i+1), 10)+strconv.FormatInt(int64(j), 10)]
				if exists == false {
					visited_9[strconv.FormatInt(int64(i+1), 10)+strconv.FormatInt(int64(j), 10)] = 1
				} else {
					visited_9[strconv.FormatInt(int64(i+1), 10)+strconv.FormatInt(int64(j), 10)] += 1
				}
			} else {
				_, visited_9 = check_surrounding_nrs(input, i+1, j, value+1, start_idx, visited_9)
			}
		}
	}
	if i > 0 {
		if int(input[i-1][j]-'0') == value+1 {
			if value+1 == 9 {
				_, exists := visited_9[strconv.FormatInt(int64(i-1), 10)+strconv.FormatInt(int64(j), 10)]
				if exists == false {
					visited_9[strconv.FormatInt(int64(i-1), 10)+strconv.FormatInt(int64(j), 10)] = 1
				} else {
					visited_9[strconv.FormatInt(int64(i-1), 10)+strconv.FormatInt(int64(j), 10)] += 1
				}
			} else {
				_, visited_9 = check_surrounding_nrs(input, i-1, j, value+1, start_idx, visited_9)
			}
		}
	}
	if j > 0 {
		if int(input[i][j-1]-'0') == value+1 {
			if value+1 == 9 {
				_, exists := visited_9[strconv.FormatInt(int64(i), 10)+strconv.FormatInt(int64(j-1), 10)]
				if exists == false {
					visited_9[strconv.FormatInt(int64(i), 10)+strconv.FormatInt(int64(j-1), 10)] = 1
				} else {
					visited_9[strconv.FormatInt(int64(i), 10)+strconv.FormatInt(int64(j-1), 10)] += 1
				}
			} else {
				_, visited_9 = check_surrounding_nrs(input, i, j-1, value+1, start_idx, visited_9)
			}
		}
	}
	return len(visited_9), visited_9
}

func read_file(path string) []string {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	var string_slice []string
	for scanner.Scan() {
		line := scanner.Text()
		string_slice = append(string_slice, line)
	}
	return string_slice
}

//func check_surrounding_nrs(input []string, i int, j int, value int) bool {
//	// check left, right, above, below, diagonal, 8 checks
//	len_row := len(input[0])
//	no_of_rows := len(input)
//	// example, we have "9" at [5][5]
//	for i_2 := i - 1; i_2 < i+2; i_2++ {
//		if i_2 < 0 || i_2 >= len_row {
//			continue
//		}
//		for j_2 := j - 1; j_2 < j+2; j_2++ {
//			if j_2 < 0 || j_2 >= no_of_rows {
//				continue
//			}
//			if int(input[i_2][j_2]-'0') == value-1 {
//				if value-1 == 0 {
//					return true
//				}
//				//fmt.Println(value, input[i_2][j_2]-'0')
//				found := check_surrounding_nrs(input, i_2, j_2, value-1)
//				if found == false {
//					continue
//				} else {
//					return true
//				}
//			}
//		}
//	}
//	return false
//}

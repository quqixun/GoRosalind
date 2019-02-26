package main

import (
	"fmt"
	"strconv"
	"strings"
	"io/ioutil"
)


func load_numbers(file_path string) (int, int, bool) {

	bs, err := ioutil.ReadFile(file_path)
	if err != nil {
		fmt.Println(err)
		return 0, 0, false
	}

	text := string(bs)
	numbers := strings.Fields(text)

	a, _ := strconv.Atoi(numbers[0])
	b, _ := strconv.Atoi(numbers[1])

	return a, b, true
}


func square_sum(a, b int) (int) {
	return a * a + b * b
}


func main() {

	a, b, is_load_success := load_numbers("rosalind_ini2.txt")

	if is_load_success {
		ss := square_sum(a, b)
		fmt.Println(ss)
	}
}
package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
	"io/ioutil"
)


func load_numbers(file_path string) (int, int, bool) {

	a, b := 0, 0

	bs, err := ioutil.ReadFile(file_path)
	if err != nil {
		fmt.Println(err)
		return a, b, false
	}

	text := string(bs)
	numbers := strings.Fields(text)

	a, _ = strconv.Atoi(numbers[0])
	b, _ = strconv.Atoi(numbers[1])

	return a, b, true
}


func odd_sum(a, b int) (int) {

	sum := 0
	for i := a; i <= b; i++ {
		if math.Mod(float64(i), float64(2)) == 1 {
			sum += i
		}
	}

	return sum
}


func main() {

	a, b, is_load_success := load_numbers("rosalind_ini4.txt")

	if is_load_success {
		os := odd_sum(a, b)
		fmt.Println(os)
	}

}
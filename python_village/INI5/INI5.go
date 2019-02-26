package main

import (
	"fmt"
	"math"
	"bufio"
	"strings"
	"io/ioutil"
)


func load_content(file_path string) ([]string, bool) {

	content := []string{}

	b, err := ioutil.ReadFile(file_path)
	if err != nil {
		fmt.Println(err)
		return content, false
	}

	text := string(b)
	scanner := bufio.NewScanner(strings.NewReader(text))
	for scanner.Scan() {
		content = append(content, scanner.Text())
	}

	return content, true
}


func main() {

	content, is_load_success := load_content("rosalind_ini5.txt")

	if is_load_success {
		for i, s := range content {
			i += 1
			if math.Mod(float64(i), float64(2)) == 0 {
				fmt.Println(s)
			}
		}
	}
}
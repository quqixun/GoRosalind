package main

import (
	"fmt"
	"bufio"
	"strconv"
	"strings"
	"io/ioutil"
)


func LoadData(file_path string) (string, int, int) {

	b, err := ioutil.ReadFile(file_path)
	if err != nil {
		panic(err)
	}

	tb := string(b)
	content := []string{}
	scanner := bufio.NewScanner(strings.NewReader(tb))
	for scanner.Scan() {
		content = append(content, scanner.Text())
	}

	text := content[0]
	num_strs := strings.Fields(content[1])
	k, _ := strconv.Atoi(num_strs[0])
	d, _ := strconv.Atoi(num_strs[1])

	return text, k, d
}


func FindKMers(text string, k int) (map[string]int) {

	kmers := make(map[string]int)
	for i := 0; i <= len(text) - k; i++ {
		kmers[text[i:i + k]] += 1
	}

	return kmers
}


func main() {
	
	text, k, d := LoadData("rosalind_ba1i.txt")
	fmt.Println(d)
	kmers := FindKMers(text, k)
	fmt.Println(kmers)
}
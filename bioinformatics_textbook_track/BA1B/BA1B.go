package main

import (
	"fmt"
	"bufio"
	"strconv"
	"strings"
	"io/ioutil"
)


func LoadData(file_path string) (string, int) {

	b, err := ioutil.ReadFile(file_path)
	if err != nil {
		panic(err)
	}

	text := string(b)
	content := []string{}
	scanner := bufio.NewScanner(strings.NewReader(text))
	for scanner.Scan() {
		content = append(content, scanner.Text())
	}

	DNA := content[0]
	k, _ := strconv.Atoi(content[1])

	return DNA, k
}


func MostFrequentKMers(DNA string, k int) ([]string) {

	all_k_mers := make(map[string]int)
	for i := 0; i <= len(DNA) - k; i++ {
		subDNA := DNA[i:i + k]
		all_k_mers[subDNA] += 1
	}

	max_frequent := 0
	for _, v := range all_k_mers {
		if v > max_frequent {
			max_frequent = v
		}
	}

	mf_k_mers := []string{}
	for k, v := range all_k_mers {
		if v == max_frequent {
			mf_k_mers = append(mf_k_mers, k)
		}
	}

	return mf_k_mers
}


func main() {
	
	DNA, k := LoadData("rosalind_ba1b.txt")
	mf_k_mers := MostFrequentKMers(DNA, k)
	fmt.Println(strings.Join(mf_k_mers, " "))
}
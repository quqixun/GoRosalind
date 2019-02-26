package main

import (
	"fmt"
	"bufio"
	"strconv"
	"strings"
	"io/ioutil"
)


func load_text(file_path string) (string, bool) {

	b, err := ioutil.ReadFile(file_path)
	if err != nil {
		fmt.Println(err)
		return "", false
	}

	text := string(b)
	return text, true
}


func load_table(file_path string) (map[string]float64, bool) {

	b, err := ioutil.ReadFile(file_path)
	if err != nil {
		fmt.Println(err)
		return nil, false
	}

	table := make(map[string]float64)
	scanner := bufio.NewScanner(strings.NewReader(string(b)))

	for scanner.Scan() {
		line_text := scanner.Text()
		words := strings.Fields(line_text)
		table[words[0]], _ = strconv.ParseFloat(words[1], 64)
	}

	return table, true
}


func protein_mass(protein string, table map[string]float64) (float64) {

	mass := 0.0

	for _, c := range(protein) {
		mass += table[string(c)]
	}

	return mass
}


func main() {

	text, is_text_success := load_text("rosalind_prtm.txt")
	table, is_table_success := load_table("monoisotopic_mass_table.txt")

	if is_text_success && is_table_success {
		mass := protein_mass(text, table)
		fmt.Println(mass)
	}
}
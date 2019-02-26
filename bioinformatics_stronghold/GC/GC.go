package main

import (
	"fmt"
	"bufio"
	"strings"
	"io/ioutil"
)


func load_text(file_path string) (string) {

	b, err := ioutil.ReadFile(file_path)
	if err != nil {
		fmt.Println(err)
		return "Error"
	}

	text := string(b)
	return text
}


func load_samples(text string) (map[string]string) {

	samples := make(map[string]string)

	id := ""
	scanner := bufio.NewScanner(strings.NewReader(text))
	for scanner.Scan() {
		// Scan each line in text
		line_text := scanner.Text()
		i := strings.Index(line_text, ">")
		if i > -1 {
			// Line starts with ">" has id
			id = line_text[1:]
			samples[id] = ""
		} else {
			if id != "" {
				// Read sequence of the id
				samples[id] += line_text
			}
		}
	}

	return samples
}


func load_gc_content(samples map[string]string) (map[string]float64) {

	gc_content := make(map[string]float64)

	for k, v := range samples {
		v_length := float64(len(v))
		gc_length := float64(strings.Count(v, "G") + strings.Count(v, "C"))
		gc_content[k] = gc_length / v_length * 100.0
	}

	return gc_content
}


func max_gc_content(gc_content map[string]float64) (string, float64) {

	id := ""
	gc_prop := 0.0

	for k, v := range gc_content {
		if v > gc_prop {
			id = k
			gc_prop = v
		}
	}

	return id, gc_prop
}


func main() {

	text := load_text("rosalind_gc.txt")
	if text != "Error" {
		samples := load_samples(text)
		gc_content := load_gc_content(samples)
		id, gc_prop := max_gc_content(gc_content)

		fmt.Println(id)
		fmt.Printf("%.6f", gc_prop)
	}
}
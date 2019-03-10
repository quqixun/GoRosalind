package main

import (
  "fmt"
  "bufio"
  "strconv"
  "strings"
  "io/ioutil"
)


func LoadData(file_path string) (string, string) {

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

  pattern := content[0]
  genome := content[1]

  return pattern, genome
}


func PatternStartPositions(pattern, genome string) ([]int) {

  len_pattern := len(pattern)
  len_genome := len(genome)

  positions := []int{}
  for i := 0; i <= len_genome - len_pattern; i++ {
    sub_genome := genome[i:i + len_pattern]
    if sub_genome == pattern {
      positions = append(positions, i)
    }
  }

  return positions
}


func IntSliceToString(ints []int) (string) {

  int_strs := []string{}
  for _, p := range ints {
    int_strs = append(int_strs, strconv.Itoa(p))
  }

  str := strings.Join(int_strs, " ")
  return str
}


func main() {
  
  pattern, genome := LoadData("rosalind_ba1d.txt")
  positions := PatternStartPositions(pattern, genome)
  positions_str := IntSliceToString(positions)
  fmt.Println(positions_str)
}
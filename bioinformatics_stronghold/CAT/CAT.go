package main

import (
  "fmt"
  "math"
  "bufio"
  "strings"
  "io/ioutil"
)


func LoadData(file_path string) (string) {

  b, err := ioutil.ReadFile(file_path)
  if err != nil {
    panic(err)
  }
  text := string(b)

  DNA := ""
  scanner := bufio.NewScanner(strings.NewReader(text))
  for scanner.Scan() {
    line_text := scanner.Text()
    if strings.Index(line_text, ">") == -1 {
      DNA += line_text
    }
  }

  return DNA
}


func IsVisited(visited map[string]int, seq string) (bool) {

  for k, _ := range visited {
    if k == seq {
      return true
    }
  }

  return false
}


func CountCatalan(seq string, visited map[string]int) (int) {

  if len(seq) <= 1 {
    return 1
  }

  if IsVisited(visited, seq) {
    return visited[seq]
  }

  sum := 0
  for i := 1; i < len(seq); i += 2 {
    seq0 := string(seq[0])
    seqi := string(seq[i])

    if (seq0 == "A" && seqi == "U") ||
       (seq0 == "U" && seqi == "A") ||
       (seq0 == "C" && seqi == "G") ||
       (seq0 == "G" && seqi == "C") {
      sum += CountCatalan(seq[1:i], visited) *
             CountCatalan(seq[i + 1:len(seq)], visited)
    }
  }

  visited[seq] = int(math.Mod(float64(sum), float64(1000000)))
  return visited[seq]
}


func main() {

  DNA := LoadData("rosalind_cat.txt")
  visited := make(map[string]int)
  cat := CountCatalan(DNA, visited)
  fmt.Println(cat)
}
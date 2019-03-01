package main

import (
  "fmt"
  "math"
  "bufio"
  "strconv"
  "strings"
  "io/ioutil"
)


func load_data(file_path string) (int, float64, string) {

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

  N_x_words := strings.Fields(content[0])
  s := content[1]

  N, _ := strconv.Atoi(N_x_words[0])
  x, _ := strconv.ParseFloat(N_x_words[1], 64)

  return N, x, s
}


func prob(N int, x float64, s string) (float64) {

  gc_p := x / 2.0
  at_p := (1 - x) / 2.0

  p := 1.0
  for _, c := range s {
    sc := string(c)
    if sc == "G" || sc == "C" {
      p *= gc_p
    } else {
      p *= at_p
    }
  }

  fp := 1 - math.Pow((1 - p), float64(N))
  return fp
}


func main() {

  N, x, s := load_data("rosalind_rstr.txt")

  fp := prob(N, x, s)
  fmt.Println(fp)
}
package main

import (
  "fmt"
  "math"
  "bufio"
  "strconv"
  "strings"
  "io/ioutil"
)


func StringToFloats(str string) ([]float64) {

  fs := []float64{}
  f_strs := strings.Split(str, " ")
  for _, f_str := range f_strs {
    f, _ := strconv.ParseFloat(f_str, 64)
    fs = append(fs, f)
  }

  return fs
}


func LoadData(file_path string) ([]float64, []float64) {

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

  s1 := StringToFloats(content[0])
  s2 := StringToFloats(content[1])
  return s1, s2
}


func SpectralConvolution(s1, s2 []float64) ([]float64) {

  sc := []float64{}
  for _, e1 := range s1 {
    for _, e2 := range s2 {
      sc = append(sc, math.Round((e1 - e2) * 100000) / 100000.0)
    }
  }

  return sc
}


func  LargestMultiplicity(conv []float64) (int, float64) {
  
  conv_map := make(map[float64]int)
  lm, x := 0, 0.0

  for _, c := range conv {
    conv_map[c] += 1
    if conv_map[c] > lm {
      lm = conv_map[c]
      x = c
    }
  }

  return lm, x
}


func main() {
  
  s1, s2 := LoadData("rosalind_conv.txt")
  sc := SpectralConvolution(s1, s2)
  lm, x := LargestMultiplicity(sc)

  fmt.Println(lm)
  fmt.Println(x)
}
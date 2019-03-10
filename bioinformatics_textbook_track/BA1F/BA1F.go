package main

import (
  "fmt"
  "strconv"
  "strings"
  "io/ioutil"
)


func LoadData(file_path string) (string) {

  b, err := ioutil.ReadFile(file_path)
  if err != nil {
    panic(err)
  }

  genome := string(b)
  return genome
}


func MinimizingSkew(genome string) ([]int) {

  skew := []int{0}
  for i := 1; i <= len(genome); i++ {
    sub_genome := genome[0:i]
    G_num := strings.Count(sub_genome, "G")
    C_num := strings.Count(sub_genome, "C")
    skew = append(skew, G_num - C_num)
  }

  min_skew := len(genome)
  for _, s := range skew {
    if s < min_skew {
      min_skew = s
    }
  }

  min_skew_i := []int{}
  for i, s := range skew {
    if s == min_skew {
      min_skew_i = append(min_skew_i, i)
    }
  }

  return min_skew_i
}


func IntSliceToString(ints []int) (string) {

  int_strs := []string{}
  for _, i := range ints {
    int_strs = append(int_strs, strconv.Itoa(i))
  }

  str := strings.Join(int_strs, " ")
  return str
}


func main() {
  
  genome := LoadData("rosalind_ba1f.txt")
  min_skew_i := MinimizingSkew(genome)
  fmt.Println(IntSliceToString(min_skew_i))
}
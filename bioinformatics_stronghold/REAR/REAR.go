package main

import (
  "fmt"
  "bufio"
  "strconv"
  "strings"
  "io/ioutil"
)


func String2IntSlice(str string) ([]int) {

  int_strs := strings.Split(str, " ")

  ints := []int{}
  for _, int_str := range int_strs {
    int_n, _ := strconv.Atoi(int_str)
    ints = append(ints, int_n)
  }

  return ints
}


func LoadData(file_path string) ([][][]int) {

  b, err := ioutil.ReadFile(file_path)
  if err != nil {
    panic(err)
  }

  content := [][][]int{}
  text := string(b)
  scanner := bufio.NewScanner(strings.NewReader(text))

  pairs := [][]int{}
  for scanner.Scan() {
    line_text := scanner.Text()
    if line_text == "" {
      content = append(content, pairs)
      pairs = [][]int{}
    } else {
      pairs = append(pairs, String2IntSlice(line_text))
    }
  }

  content = append(content, pairs)
  return content
}


func Reverse(ints []int, start, end int) ([]int) {

  r := make([]int, len(ints))
  copy(r, ints)

  for i, j := start, end; i < j; i, j = i + 1, j - 1 {
        r[i], r[j] = r[j], r[i]
    }

  return r
}


func IsSameIntSlice(ints1, ints2 []int) (bool) {

  for i := 0; i < len(ints1); i++ {
    if ints1[i] != ints2[i] {
      return false
    }
  }

  return true
}


func IsVisited(slice []int, dslice [][]int) (bool) {

  for _, ds := range dslice {
    if IsSameIntSlice(slice, ds) {
      return true
    }
  }

  return false
}


func ReversalDistance(perm1, perm2 []int) (int) {

  len_p := len(perm1)
  p1c, p2c := make([]int, len_p), make([]int, len_p)
  copy(p1c, perm1)
  copy(p2c, perm2)

  if IsSameIntSlice(p1c, p2c) {
    return 0
  }

  visited := [][]int{p2c}
  levels := [][][]int{visited}

  rd := 0
  for {

    this_level := levels[rd]
    next_level := [][]int{}

    for _, pcl := range this_level {
      for i := 0; i < len_p; i++ {
        for j :=  i + 1; j < len_p; j++ {
          tmp := Reverse(pcl, i, j)

          if !IsVisited(tmp, visited) {
            if IsSameIntSlice(p1c, pcl) { return rd + 1 }
            visited = append(visited, tmp)
            next_level = append(next_level, tmp)
          }
        }
      }
    }
    
    levels = append(levels, next_level)
    rd += 1
  }

  return rd
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
  
  content := LoadData("rosalind_rear.txt")
  
  rds := []int{}
  for i, c := range content {
    if i == 1 {
      rd := ReversalDistance(c[0], c[1])
      rds = append(rds, rd)
      break
    }
  }
  // fmt.Println(content)
  fmt.Println(IntSliceToString(rds))
}
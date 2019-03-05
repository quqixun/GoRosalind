package main

import (
  "fmt"
  "bufio"
  "strconv"
  "strings"
  "io/ioutil"
)


func LoadData(file_path string) (int, [][]int) {

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

  n, _ := strconv.Atoi(content[0])
  edges := [][]int{}
  for _, c := range(content[1:len(content)]) {
    c_strs := strings.Fields(c)
    e1, _ := strconv.Atoi(c_strs[0])
    e2, _ := strconv.Atoi(c_strs[1])
    edges = append(edges, []int{e1, e2})
  }

  return n, edges
}


func MinimumEdges(n int, edges [][]int) (int) {

  n_edges := len(edges)
  return n - (n_edges + 1)
}


func main() {
  
  n, edges := LoadData("rosalind_tree.txt")
  me := MinimumEdges(n, edges)
  fmt.Println(me)
}
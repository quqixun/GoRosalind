package main

import (
  "fmt"
  "bufio"
  "strconv"
  "strings"
  "io/ioutil"
)


func LoadData(file_path string) (int, int, [][]int) {

  b, err := ioutil.ReadFile(file_path)
  if err != nil {
    panic(err)
  }

  text := string(b)
  content := [][]int{}
  scanner := bufio.NewScanner(strings.NewReader(text))
  for scanner.Scan() {
    int_strs := strings.Fields(scanner.Text())
    ints := []int{}
    for _, str := range int_strs {
      i, _ := strconv.Atoi(str)
      ints = append(ints, i)
    }
    content = append(content, ints)
  }

  n_vertices := content[0][0]
  n_edges := content[0][1]
  edges := content[1:len(content)]
  return n_vertices, n_edges, edges
}


func Degree(n_vertices, n_edges int, edges [][]int) ([]int) {

  degree := make([]int, n_vertices)

  for i := 0; i < n_edges; i++ {
    edge := edges[i]
    for _, v := range edge {
      degree[v - 1] += 1
    }
  }

  return degree
}


func DoubleDegree(n_vertices, n_edges int, edges [][]int) ([]int) {

  degree := Degree(n_vertices, n_edges, edges)
  ddegree := make([]int, n_vertices)

  for i := 0; i < n_edges; i++ {
    edge := edges[i]

    v1, v2 := edge[0], edge[1]
    ddegree[v1 - 1] += degree[v2 - 1]
    ddegree[v2 - 1] += degree[v1 - 1]
  }

  return ddegree
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
  
  n_vertices, n_edges, edges := LoadData("rosalind_ddeg.txt")
  ddegree := DoubleDegree(n_vertices, n_edges, edges)
  fmt.Println(IntSliceToString(ddegree))
}
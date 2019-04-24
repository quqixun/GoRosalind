package main

import (
  "fmt"
  "bufio"
  "strconv"
  "strings"
  "io/ioutil"
)


func LoadData(file_path string) (int, [][][]int) {

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

  graphs, graph := [][][]int{}, [][]int{}
  for _, c := range content[1:len(content)] {
    if c == "" {
      if len(graph) > 0 {
        graphs = append(graphs, graph)
        graph = [][]int{}
      }
    } else {
      num_strs := strings.Split(c, " ")
      nums := []int{}
      for _, num_str := range num_strs {
        num, _ := strconv.Atoi(num_str)
        nums = append(nums, num)
      }
      graph = append(graph, nums)
    }
  }
  graphs = append(graphs, graph)

  return n, graphs
}


func AdjacentList(n_vertices int, edges [][]int) ([][]int) {

  adjacent := [][]int{}
  for i := 0; i < n_vertices; i++ {
    adjacent = append(adjacent, []int{})
  }

  for _, e := range edges {
    adjacent[e[0] - 1] = append(adjacent[e[0] - 1], e[1] - 1)
    adjacent[e[1] - 1] = append(adjacent[e[1] - 1], e[0] - 1)
  }

  return adjacent
}


func InPairs(pairs [][]int, pair []int) (bool) {

  for _, p := range pairs {
    if (p[0] == pair[0] && p[1] == pair[1]) ||
       (p[0] == pair[1] && p[1] == pair[0]) {
        return true
    }
  }

  return false
}


func SquareCycle(graph [][]int) (int) {

  n_vertices := graph[0][0]
  edges := graph[1:len(graph)]
  adj := AdjacentList(n_vertices, edges)

  pairs := [][]int{}
  for _, nodes := range adj {
    if len(nodes) < 2 {
      continue
    } else {
      for j := 0; j < len(nodes) - 1; j++ {
        for k := j + 1; k < len(nodes); k++ {
          pair := []int{nodes[j], nodes[k]}
          if InPairs(pairs, pair) {
            return 1
          } else {
            pairs = append(pairs, pair)
          }
        }
      }
    }
  }

  return -1
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

  _, graphs := LoadData("rosalind_sq.txt")
  
  results := []int{}
  for _, graph := range graphs {
    res := SquareCycle(graph)
    results = append(results, res)
  }

  fmt.Println(IntSliceToString(results))
}
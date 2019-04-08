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
  }

  return adjacent
}


func DepthFirstSearch(adjacent [][]int, visited, rec_stack []bool, i int) (bool) {

  visited[i] = true
  rec_stack[i] = true

  for _, adj := range adjacent[i] {
    if !visited[adj] {
      if DepthFirstSearch(adjacent, visited, rec_stack, adj) {
        return true
      }
    } else if rec_stack[adj] {
      return true
    }
  }

  rec_stack[i] = false
  return false
}


func  IsCyclic(graph [][]int) (bool) {

  n_vertices := graph[0][0]
  visited := make([]bool, n_vertices)
  rec_stack := make([]bool, n_vertices)

  edges := graph[1:len(graph)]
  adjacent := AdjacentList(n_vertices, edges)

  for node := 0; node < n_vertices; node++ {
    if !visited[node] {
      if DepthFirstSearch(adjacent, visited, rec_stack, node) {
        return true
      }
    }
  }

  return false
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

  _, graphs := LoadData("rosalind_dag.txt")
  
  res := []int{}
  for _, graph := range graphs {

    resb := IsCyclic(graph)

    var resi int
    if resi = -1; !resb {
      resi = 1
    }
    res = append(res, resi)
  }
  fmt.Println(IntSliceToString(res))
}
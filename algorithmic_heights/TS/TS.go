package main

import (
  "os"
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


func TopologicalSortUtil(adj [][]int, v int, visited []bool, stack []int) ([]int) {
  visited[v] = true
  for _, i := range adj[v] {
    if visited[i] == false {
      stack = TopologicalSortUtil(adj, i, visited, stack)
    }
  }
  stack = append([]int{v + 1}, stack...)
  return stack
}


func TopologicalSort(n_vertices int, adjacent [][]int, visited []bool, stack []int) ([]int) {

  for i := 0; i < n_vertices; i++ {
    if visited[i] == false {
      stack = TopologicalSortUtil(adjacent, i, visited, stack)
    }
  }

  return stack
}


func IntSliceToString(ints []int) (string) {

  int_strs := []string{}
  for _, p := range ints {
    int_strs = append(int_strs, strconv.Itoa(p))
  }

  str := strings.Join(int_strs, " ")
  return str
}


func Write2File(file_path string, content []string) {
  f, err := os.Create(file_path)
  if err != nil {
      panic(err)
  }
  defer f.Close()

  for _, c := range content {
     fmt.Fprintln(f, c)
  }
}


func main() {

  n_vertices, _, edges := LoadData("rosalind_ts.txt")

  adjacent := AdjacentList(n_vertices, edges)
  visited := make([]bool, n_vertices)
  stack := make([]int, 0)

  stack = TopologicalSort(n_vertices, adjacent, visited, stack)
  Write2File("ts_output.txt", []string{IntSliceToString(stack)})
}
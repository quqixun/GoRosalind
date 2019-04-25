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

  n_strs := strings.Split(content[0], " ")
  n, _ := strconv.Atoi(n_strs[0])

  graph := [][]int{}
  for _, c := range content[1:len(content)] {
    edge := []int{}
    num_strs := strings.Split(c, " ")
    for _, num_str := range num_strs {
      num, _ := strconv.Atoi(num_str)
      edge = append(edge, num)
    }
    graph = append(graph, edge)
  }

  return n, graph
}


func Dijkstra(n int, graph [][]int) ([]int) {

  distance := []int{}
  for i := 0; i < n + 1; i++ {
    distance = append(distance, -1)
  }
  distance[1] = 0

  for i := 0; i < n; i++ {
    for _, edge := range graph {
      a, b, w := edge[0], edge[1], edge[2]
      if distance[a] > -1 {
        p_distance := distance[a] + w
        if distance[b] == -1 || distance[b] > p_distance {
          distance[b] = p_distance
        }
      }
    }
  }

  return distance[1:len(distance)]
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

  n, graph := LoadData("rosalind_dij.txt")
  
  D := Dijkstra(n, graph)
  fmt.Println(IntSliceToString(D))

}
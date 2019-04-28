package main

import (
  "fmt"
  "math"
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


func BellmanFord(n int, graph [][]int) ([]float64) {

  dist := []float64{}
  for i := 0; i < n; i++ {
    dist = append(dist, math.Inf(1))
  }
  dist[0] = 0

  for i := 0; i < n; i++ {
    for _, edge := range graph {
      u, v, w := edge[0], edge[1], float64(edge[2])
      if dist[u - 1] != math.Inf(1) &&
         dist[u - 1] + w < dist[v - 1]{
        dist[v - 1] = dist[u - 1] + w
      }
    }
  }

  return dist
}


func ResString(dist []float64) (string) {

  res := []string{}
  for _, d := range dist {
    if d != math.Inf(1) {
      res = append(res, strconv.Itoa(int(d)))
    } else {
      res = append(res, "x")
    }
  }

  res_str := strings.Join(res, " ")
  return res_str
}


func main() {

  n, graph := LoadData("rosalind_bf.txt")
  dist := BellmanFord(n, graph)

  dist_str := ResString(dist)
  fmt.Println(dist_str)
}
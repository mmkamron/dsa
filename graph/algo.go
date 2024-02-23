package graph

import (
  "fmt"
)

func BFS(g *Graph, key int) error {
  startVertex, err := g.GetVertex(key)
  if err != nil {
    return fmt.Errorf("vertex with value %d does not exist", key)
  }

  visited := make(map[int]bool)
  queue := []*Vertex{startVertex}

  for len(queue) > 0 {
    current := queue[0]
    queue = queue[1:]

    if !visited[current.key] {
      fmt.Printf("%d->", current.key)
      visited[current.key] = true

      for _, adjacent := range current.adjacent {
        if !visited[adjacent.key] {
          queue = append(queue, adjacent)
        }
      }
    }
  }

  return nil
}

func DFS(g *Graph, key int) error {
  startVertex, err := g.GetVertex(key)
  if err != nil {
    return fmt.Errorf("vertex with value %d does not exist", key)
  }

  var dfsRecursive func(*Vertex, map[int]bool)
  dfsRecursive = func(vertex *Vertex, visited map[int]bool) {
    if !visited[vertex.key] {
      fmt.Printf("%d->", vertex.key)
      visited[vertex.key] = true
      for _, v := range vertex.adjacent {
        if !visited[v.key] {
          dfsRecursive(v, visited)
        }
      }
    }
  }

  visited := make(map[int]bool)
  dfsRecursive(startVertex, visited)

  return nil
}

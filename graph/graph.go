package graph

import "fmt"

type Graph struct {
  vertices []*Vertex
}

type Vertex struct {
  key int
  adjacent []*Vertex
}

func (g *Graph) AddVertex(k int) error {
  if contains(g.vertices, k) {
    return fmt.Errorf("vertex with key %d already exists", k)
  }
  g.vertices = append(g.vertices, &Vertex{key: k})
  return nil
}

func (g *Graph) GetVertex(k int) (*Vertex, error) {
  for i, v := range g.vertices {
    if v.key == k {
      return g.vertices[i], nil
    }
  }
  return nil, fmt.Errorf("vertex with value %d does not exist", k)
}

func (g *Graph) AddEdge(from, to int) error {
  fromVertex, err := g.GetVertex(from);
  if err != nil {
    return fmt.Errorf("from vertex with value %d does not exist", from)
  }
  toVertex, err := g.GetVertex(to);
  if err != nil {
    return fmt.Errorf("to vertex with value %d does not exist", to)
  }
  if contains(fromVertex.adjacent, to) {
    return fmt.Errorf("Edge from %d to %d already exists", from, to)
  }
  fromVertex.adjacent = append(fromVertex.adjacent, toVertex)
  return nil
}

func (g *Graph) Print() {
  for _, v := range g.vertices {
    fmt.Printf("\nVertex %v : ", v.key)
    for _, v := range v.adjacent {
      fmt.Printf("%v ", v.key)
    }
  }
  fmt.Println()
}

func contains(s []*Vertex, k int) bool {
  for _, v := range s {
    if v.key == k {
      return true
    } 
  }
  return false
}

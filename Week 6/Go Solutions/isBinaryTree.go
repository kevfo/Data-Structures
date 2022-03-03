package main

import (
  "fmt"
)

// The Node "class"
type Node struct {
  key     int
  left    int
  right   int
}

func initialize(k, l, r int) *Node {
  return &Node{key: k, left: l, right: r}
}

// ---

func inOrder(tree []*Node) []int {
  var (
    result            []int
    stack             []*Node
    curr              int
    root              *Node
  )
  curr = 0

  for len(stack) != 0 || curr != -1 {
    if curr != -1 {
      root = tree[curr]
      stack = append(stack, root)
      curr = root.left
    } else {
      root = stack[len(stack) - 1] ; stack = stack[0:len(stack) - 1]
      result = append(result, root.key)
      curr = root.right
    }
  }
  return result
}

func isBST(tree []*Node) bool {
  var nodes = inOrder(tree)

  for i := 1; i < len(nodes); i++ {
    if nodes[i] <= nodes[i - 1] {return false}
  }
  return true
}

func main() {
  var (
    t                 []*Node
    k, l, r, numNodes int
  )
  fmt.Scanln(&numNodes)

  for i := 0; i < numNodes; i++ {
    fmt.Scan(&k) ; fmt.Scan(&l) ; fmt.Scan(&r)
    t = append(t, initialize(k, l, r))
  }

  if numNodes == 0 || isBST(t) {
    fmt.Println("CORRECT")
  } else {
    fmt.Println("INCORRECT")
  }
}

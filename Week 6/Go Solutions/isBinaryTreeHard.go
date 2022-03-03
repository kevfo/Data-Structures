package main

import (
  "fmt"
)

type Node struct {
  key     int
  left    int
  right   int
}

func initialize(k, l, r) *Node {
  return &Node{key: k, left: l, right: r}
}

type BSTStack struct {
  min   int
  tree  *Node
  max   int
}

func isBinary(tree []*Node) bool {
  stack := [][]BSTStack{{-999999, tree[0], 999999}}
  for len(stack) != 0 {
    min, root, max := stack[len(stack) - 1][0], stack[len(stack) - 1][1], stack[len(stack) - 1][2]
    stack = stack[0:len(stack) - 1]

    if root.key < min || root.key >= max {
      return false
    }
    if root.left != -1 {stack = append(stack, []BSTStack{min, tree[root.key], root.key})}
    if root.right != -1 {stack = append(stack, []BSTStack{root.key, tree[root.right], max})}
  }
  return true
}

func main() {
  var(
    numNodes, k, l, r    int
    nodes                []*Node
  )
  fmt.Scanln(&numNodes)

  for i := 0; i < numNodes; i++ {
    fmt.Scan(&k) ; fmt.Scan(&l) ; fmt.Scan(&r)
    nodes = append(nodes, initialize(k, l, r))
  }

  if numNodes == 0 || isBinary(nodes) {
    fmt.Println("CORRECT")
  } else {
    fmt.Println("INCORRECT")
  }
}

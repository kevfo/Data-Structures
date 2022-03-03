package main

import (
  "fmt"
)

type Node struct {
  key    int
  left   *Node
  right  *Node
  parent *Node
}

func insert(node *Node, value int) *Node {
  if node == nil {
    node = &Node{key: value, left: nil, right: nil, parent: nil}
  } else if value > node.key {
    node.right = insert(node.right, value)
    node.right.parent = node
  } else {
    node.left = insert(node.left, value)
    node.left.parent = node
  }
  return node
}

func inOrderTraversal(node *Node) {
  if node == nil {return}
  inOrderTraversal(node.left)
  fmt.Print(node.key, " ")
  inOrderTraversal(node.right)
}

func preOrderTraversal(node *Node) {
  if node == nil {return}
  fmt.Print(node.key, " ")
  preOrderTraversal(node.left)
  preOrderTraversal(node.right)
}

func postOrderTraversal(node *Node) {
  if node == nil {return}
  postOrderTraversal(node.left)
  postOrderTraversal(node.right)
  fmt.Print(node.key, " ")
}

func main() {
  var (
    numNodes    int
    v, l, r     int
    values      []int
    node        *Node
    useless     [][]int
  )

  fmt.Scanln(&numNodes)
  for i := 0; i < numNodes; i++ {
    fmt.Scan(&v) ; fmt.Scan(&l) ; fmt.Scan(&r)
    values = append(values, v)
    useless = append(useless, []int{l, r})
  }
  node = nil

  for i := 0; i < len(values); i++ {
    node = insert(node, values[i])
  }

  inOrderTraversal(node)
  fmt.Println("")
  preOrderTraversal(node)
  fmt.Println("")
  postOrderTraversal(node)
  fmt.Println("")
}

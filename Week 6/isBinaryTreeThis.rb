# Solution passes!!!

# Fetches the user's input:
numNodes = gets.chomp.to_i

class Node
  def initialize(a, b, c)
    @key = a
    @left = b
    @right = c
  end

  attr_reader :key, :left, :right
end

def inOrder(tree)
  stack, result, curr = [], [], 0

  while !stack.empty? || curr != -1
    if curr != -1
      root = tree[curr]
      stack << root
      curr = root.left
    else
      root = stack.pop
      result.push(root.key)
      curr = root.right
    end
  end
  result
end

def isBinaryTree?(tree)
  nodes = inOrder(tree)
  (1...nodes.length).each do |i|
    return false if nodes[i] <= nodes[i - 1]
  end
  true
end

# Fetches the user's input, processes it, and displays the output:
t = Array.new()
numNodes.times do |i|
  a, b, c = gets.chomp.split(/ /).map(&:to_i)
  t << Node.new(a, b, c)
end

puts (numNodes == 0 || isBinaryTree?(t)) ? "CORRECT" : "INCORRECT"

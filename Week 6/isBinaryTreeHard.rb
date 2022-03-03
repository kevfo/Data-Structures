# Solution passes!!!

class Node
  def initialize(key, left, right)
    @key = key
    @left = left
    @right = right
  end

  attr_reader :key, :left, :right
end

def isBST?(tree)
  stack = [[-Float::INFINITY, tree[0], Float::INFINITY]]
  while !stack.empty?
    min, root, max = stack.pop
    return false if root.key < min || root.key >= max
    stack << [min, tree[root.left], root.key] if root.left != -1
    stack << [root.key, tree[root.right], max] if root.right != -1
  end
  true
end

# Get the user's input and process it:

numNodes = gets.chomp.to_i

nodes = []
numNodes.times do |i|
  key, left, right = gets.chomp.split(/ /).map(&:to_i)
  nodes << Node.new(key, left, right)
end

puts numNodes == 0 || isBST?(nodes) ? "CORRECT" : "INCORRECT"

# To represent the structure of a vertex in a splay tree...

class Vertex
  def initialize(key, sum, left, right, parent)
    @key, @sum = key, sum
    @left, @right, @parent = left, right, parent
  end

  attr_accessor :key, :left, :right:, :parent, :sum
end

def update(v)
  return if v.nil?

  v.sum = (v.right.nil? ? 0 : v.right.sum) + (v.left.nil? ? 0 : v.left.sum)

  v.left.parent = v if !v.left.nil?
  v.right.parent = v if !v.right.nil?
end

def smallRotation(v)
  parent = v.parent
  return if parent.nil?
  grandparent = v.parent.parent

  if parent.left != v
  else
  end 
end

# Represents the structure of a node in a splay tree...

class Vertex
  def initialize(key, sum, left, right, parent)
    @key, @sum = key, sum
    @left, @right, @parent = left, right, parent
  end

  attr_accessor :key, :sum, :left, :right, :parent
end

def update(v)
  return if v.nil?

  v.sum = (v.left.nil? ? 0 : v.left.sum) + (v.right.nil? ? 0 : v.right.sum) + v.key

  v.right.parent = v if !v.right.nil?
  v.left.parent = v if !v.left.nil?
end

def smallRotation(v)
  parent = v.parent
  return if parent.nil?
  grandparent = v.parent.parent

  if parent.left == v
    m = v.right
    v.right = parent
    parent.left = m
  elsif parent.right == v
    m = v.left
    v.left = parent
    parent.right = m
  end

  update(parent) ; update(v)
  v.parent = grandparent
  if !grandparent.nil?
    if grandparent.left == parent
      grandparent.left = v
    else
      grandparent.right = v
    end
  end
end

def bigRotation(v)
  if v.parent.parent.left == v.left && v.parent.left == v
    smallRotation(v.parent) ; smallRotation(v)
  elsif v.parent.parent.right == v.right && v.parent.right == v
    smallRotation(v.parent) ; smallRotation(v)
  else
    2.times{|i| smallRotation(v)}
  end
end

def splay(v)
  return nil if v.nil?

  while !v.parent.nil?
    if v.parent.parent.nil?
      smallRotation(v)
      break
    end
    bigRotation(v)
  end
  v
end

def find(root, key)
  v, last, n = root, root, nil

  while !v.nil?
    n = v if v.key >= key && (n.nil? || v.key < n.key)
    last = v
    break if v.key == key
    v = v.key < key ? v.right : v.left
  end

  root = splay(last)
  root = splay(n) if !n.nil?
  return n, root
end

def split(root, key)
  result, root = find(root, key)
  return root, nil if result.nil?

  right = splay(result)
  left = right.left
  right.left = nil
  left.parent = nil if !left.nil?
  update(left) ; update(right)
  return left, right
end

def merge(left, right)
  return right if left.nil?
  return left if right.nil?

  right = right.left while !right.left.nil?
  right = splay(right)
  right.left = left
  left.parent = right
  update(right)
  right
end

@root = nil

def insert(x)
  left, right = split(@root, x)
  newVertex = nil
  newVertex = Vertex.new(x, x, nil, nil, nil) if right.nil? || right.key != x
  @root = merge(merge(left, newVertex), right)
end

def erase(x)
  # Implement yourself!
  left, middle = split(@root, x)
  middle, right = split(middle, x + 1)
  @root = merge(left, right)
end

def search(x)
  # Implement yourself!
  result, @root = find(@root, x)
  return result.nil? || result.key != x ? nil : result.key
end

def sum(from, to)
  left, middle = split(@root, from)
  middle, right = split(middle, to + 1)
  ans = 0
  # Implement yourself!
  ans += middle.sum if !middle.nil?
  @root = merge(merge(left, middle), right)
  ans
end

MODULO = 1000000001

numInputs = gets.chomp.to_i

inputs = []
numInputs.times do |i|
  op, n1, n2 = gets.chomp.split(/ /)
  inputs << [op, n1.to_i, n2.to_i]
end

last_sum_result = 0
inputs.each do |i|
  case i[0]
  when '+'
    insert((i[1] + last_sum_result) % MODULO)
  when '-'
    erase((i[1] + last_sum_result) % MODULO)
  when '?'
    puts search((i[1] + last_sum_result) % MODULO) ? "Found" : "Not found"
  when 's'
    res = sum((i[1] + last_sum_result) % MODULO, (i[2] + last_sum_result) % MODULO)
    puts res
    last_sum_result = res % MODULO
  end
end

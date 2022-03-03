class TreeOrder
  def read()
    @numNodes = gets.chomp.to_i
    @keys = Array.new(@numNodes, 0)
    @lefts = Array.new(@numNodes, 0)
    @rights = Array.new(@numNodes, 0)

    (0...@numNodes).each do |i|
      k, l, r = gets.chomp.split(/ /).map(&:to_i)
      @keys[i] = k ; @lefts[i] = l ; @rights[i] = r
    end
  end

  def inOrderTraversal(root = 0)
    return if root == -1
    self.inOrderTraversal(@lefts[root])
    print "#{self.keys[root]} "
    self.inOrderTraversal(@rights[root])
  end

  def preOrderTraversal(root = 0)
    return if root == -1
    print "#{@keys[root]} "
    self.preOrderTraversal(@lefts[root])
    self.preOrderTraversal(@rights[root])
  end

  def postOrderTraversal(root = 0)
    return if root == -1
    self.postOrderTraversal(@lefts[root])
    self.postOrderTraversal(@rights[root])
    print "#{@keys[root]} "
  end
end

tree = TreeOrder.new()
tree.read()
tree.inOrderTraversal()
puts ""
tree.preOrderTraversal()
puts ""
tree.postOrderTraversal()

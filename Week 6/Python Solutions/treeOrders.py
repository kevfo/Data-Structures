class TreeOrder:
    def read(self):
        self.numNodes = int(input())
        self.keys = [0] * self.numNodes
        self.lefts = [0] * self.numNodes
        self.rights = [0] * self.numNodes

        for i in range(self.numNodes):
            k, l, r = [int(i) for i in input().split(" ")]
            self.keys[i] = k ; self.lefts[i] = l ; self.rights[i] = r

    def inOrderTraversal(self, root = 0):
        if root == -1: return
        self.inOrderTraversal(self.lefts[root])
        print(self.keys[root], end = " ")
        self.inOrderTraversal(self.rights[root])

    def preOrderTraversal(self, root = 0):
        if root == -1: return
        print(self.keys[root], end = " ")
        self.preOrderTraversal(self.lefts[root])
        self.preOrderTraversal(self.rights[root])

    def postOrderTraversal(self, root = 0):
        if root == -1: return
        self.postOrderTraversal(self.lefts[root])
        self.postOrderTraversal(self.rights[root])
        print(self.keys[root], end = " ")

tree = TreeOrder()
tree.read()
tree.inOrderTraversal()
print("")
tree.preOrderTraversal()
print("")
tree.postOrderTraversal()

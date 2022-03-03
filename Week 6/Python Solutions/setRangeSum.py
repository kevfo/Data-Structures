# Represents the vertex of a splay tree:

class Vertex:
    def __init__(self, key, sum, left, right, parent):
        self.key, self.sum = key, sum
        self.left, self.right, self.parent = left, right, parent

def update(v):
    if v == None: return

    v.sum = v.key + (v.left.sum if v.left != None else 0) + (v.right.sum if v.right != None else 0)

    if v.left != None: v.left.parent = v
    if v.right != None: v.right.parent = v

def smallRotation(v):
    parent = v.parent
    if parent == None: return
    grandparent = v.parent.parent

    if parent.left == v:
        m = v.right
        v.right = parent
        parent.left = m
    else:
        m = v.left
        v.left = parent
        parent.right = m

    update(parent) ; update(v)
    v.parent = grandparent

    if grandparent != None:
        if grandparent.left == parent:
            grandparent.left = v
        else:
            grandparent.right = v

def bigRotation(v):
    if v.parent.parent.left == v.left and v.parent.left == v:
        smallRotation(v.parent) ; smallRotatation(v)
    elif v.parent.parent.right == v.right and v.parent.right == v:
        smallRotatoin(v.parent) ; smallRotation(v)
    else:
        smallRotation(v) ; smallRotation(v)

def splay(v):
    if v == None: return None

    while v.parent != None:
        if v.parent.parent == None:
            smallRotation(v)
            break
        bigRotation(v)
    return v

def find(root, key):
    v, last, next = root, root, None

    while v != None:
        if v.key >= key and (next == None or v.key < next.key):
            next = v
        last = v
        if v.key == key: break
        if v.key < key:
            v = v.right
        else:
            v = v.left

    root = splay(last)
    if next != None: root = splay(next)
    return(next, root)

def split(root, key):
    result, root = find(root, key)
    if result == None: return(root, None)
    right = splay(result)
    left = right.left
    right.left = None

    if left != None: left.parent = None

    update(left) ; update(right)
    return(left, right)

def merge(left, right):
    if left == None: return right
    if right == None: return left

    while right.left != None: right = right.left

    right = splay(right) ; right.left = left ; left.parent = right
    update(right)
    return right

# Super DUPER important!
root = None

def insert(x):
    global root
    left, right = split(root, x)
    newVertex = None
    if right == None or right.key != x:
        newVertex = Vertex(x, x, None, None, None)
    root = merge(merge(left, newVertex), right)

def erase(x):
    global root
    left, middle = split(root, x)
    middle, right = split(middle, x + 1)
    root = merge(left, right)

def search(x):
    global root
    result, root = find(root, x)
    return(False if result == None or result.key != x else True)

def sum(fr, to):
    global root
    left, middle = split(root, fr)
    middle, right = split(middle, to + 1)
    ans = 0

    if middle != None: ans += middle.sum
    root = merge(merge(left, middle), right)

    return ans

MODULO = 1000000001
last_sum = 0
commands = []

numOperations = int(input())

for i in range(numOperations):
    commands.append(input().split(" "))

for i in commands:
    if i[0] == "+":
        insert((int(i[1]) + last_sum) % MODULO)
    elif i[0] == "-":
        erase((int(i[1]) + last_sum) % MODULO)
    elif i[0] == "?":
        print("Found" if search((int(i[1]) + last_sum) % MODULO) else "Not found")
    elif i[0] == "s":
        l, r = int(i[1]), int(i[2])
        res  = sum((l + last_sum) % MODULO, (r + last_sum) % MODULO)
        print(res)
        last_sum = res % MODULO

#include <iostream>
#include <vector>

using std::cin;
using std::cout;
using std::vector;

class TreeOrder {
public:
  long int numNodes;
  vector<int> keys;
  vector<int> lefts;
  vector<int> rights;

  void read() {
    int k, l, r;

    cin >> numNodes;
    keys.resize(numNodes);
    lefts.resize(numNodes);
    rights.resize(numNodes);

    for (int i = 0; i < numNodes; i++) {
      cin >> k >> l >> r;
      keys[i] = k; lefts[i] = l; rights[i] = r;
    }
  }

  void inOrderTraversal(int root = 0) {
    if (root == -1) {return;}
    inOrderTraversal(this->lefts[root]);
    cout << this->keys[root] << " ";
    inOrderTraversal(this->rights[root]);
  }

  void preOrderTraversal(int root = 0) {
    if (root == -1) {return;}
    cout << this->keys[root] << " ";
    preOrderTraversal(this->lefts[root]);
    preOrderTraversal(this->rights[root]);
  }

  void postOrderTraversal(int root = 0) {
    if (root == -1) {return;}
    postOrderTraversal(this->lefts[root]);
    postOrderTraversal(this->rights[root]);
    cout << this->keys[root] << " ";
  }
};

int main() {
  TreeOrder tree;
  tree.read();
  tree.inOrderTraversal();
  cout << "\n";
  tree.preOrderTraversal();
  cout << "\n";
  tree.postOrderTraversal();
  return 0;
}

#include <iostream>
#include <vector>
#include <uitlity>
#include <string>

using std::cout;
using std::cin;
using std::vector;
using std::pair;
using std::make_pair;
using std::string;

class Vertex {
public:
  long long  key;
  long long  sum;
  Vertex*    left;
  Vertex*    right;
  Vertex*    parent;
};

void update(Vertex* v) {
  if (v == NULL) {return;}

  v->sum = v->key + (v->left == NULL ? 0 : v->left->sum) + (v->right == NULL ? 0 : v->right->sum);

  if (v->left != NULL) {v->left->parent = v;}
  if (v->right != NULL) {v->right->parent = v;}
}

void smallRotation(Vertex* v) {
  Vertex* m = NULL, parent = v->parent, grandparent = v->parent->parent;
  if (parent == NULL) {return;}

  if (parent->left == v) {
    m = v->right;
    v->right = parent;
    parent->left = m;
  } else {
    m = v->left;
    v->left = parent;
    parent->right = m;
  }

  update(parent) ; update(v);
  v->parent = grandparent;

  if (grandparent != NULL) {
    if (grandparent->left == parent) {
      grandparent->left = v;
    } else {
      grandparent->right = v;
    }
  }
}

void bigRotation(Vertex* v) {
  if (v->parent->parent->left == v->left && v->parent->left == v) {
    smallRotation(v->parent) ; smallRotation(v);
  } else if (v->parent->parent->right == v->right && v->parent->right == v) {
    smallRotation(v->parent) ; smallRotation(v);
  } else {
    smallRotation(v) ; smallRotation(v);
  }
}

Vertex* splay(Vertex* v) {
  if (v == NULL) {return NULL;}

  while (v->parent != NULL) {
    if (v->parent->parent == NULL) {
      smallRotation(v);
      break;
    }
    bigRotation(v);
  }
  return v;
}

pair<Vertex*, Vertex*> find(Vertex* root, long long key) {
  Vertex* v = root, last = root, next = NULL;

  while (v != NULL) {
    if (v->key >= key && (next == NULL || v->key < next->key)) {
      next = v;
    }
    last = v;
    if (v->key == key) {break;}
    v = v->key < key ? v->right : v->left;
  }

  root = splay(last);
  if (next != NULL) {root = splay(next);}
  return make_pair(next, root);
}

Vertex* root = NULL;



int main() {
  long long numCommands;
  string command;
  vector<string> commands;

  cin >> numCommands;
  commands.resize(numCommands)

  for (int i = 0; i < numCommands; i++) {
    cin >> command;
    commands[i] = command;
  }

  for (int i = 0; i < command.size(); i++) {

  }
}

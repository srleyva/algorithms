#ifndef HEADER_FILE
#define HEADER_FILE

struct Node *Node_create(int value, struct Node *right, struct Node *left);
struct Tree *Tree_create(void);
void Tree_insert(struct Tree *tree, int value);
struct Node *Tree_find(struct Tree *tree, int value);
void Tree_inorder_traversal(struct Node *root);

#endif

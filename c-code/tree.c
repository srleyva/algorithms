#include <stdio.h>
#include <stdlib.h>
#include "algos.h"

struct Node {
    int value;
    struct Node *right;
    struct Node *left;
};

struct Node *Node_create(int value, struct Node *right, struct Node *left)
{
    struct Node *rval = malloc(sizeof(struct Node));
    rval->value = value;
    rval->right = right;
    rval->left = left;
    return rval;
}

struct Tree {
    struct Node *root;
};

struct Tree *Tree_create(void)
{
    struct Tree *rval = malloc(sizeof(struct Tree));
    rval->root = NULL;
    return rval;

}

void Tree_insert(struct Tree *tree, int value)
{
    struct Node *newNode = Node_create(value, NULL, NULL);

    if (!tree->root)
    {
        tree->root = newNode;
        return;
    }

    struct Node *current = tree->root;

    while(current)
    {
        if (current->value == value)
        {
            printf("WARNING: %d already present in tree. Doing nothing...\n", value);
        } 

        else if (current->value > value)
        {
            if (!current->left) {
                current->left = newNode;
                break;
            } else {
                current = current->left;
            }
        } 

        else if (current->value < value)
        {
            if (!current->right) {
                current->right = newNode;
                break;
            } else {
                current = current->right;
            }
        }
    }
}

struct Node *Tree_find(struct Tree *tree, int value)
{
    struct Node *current = tree->root;

    while(current)
    {
        if (current->value == value)
        {
            return current;
        } 

        else if (current->value > value)
        {
            current = current->left;
        } 

        else if (current->value < value)
        {
            current = current->right;
        }
    }

    printf("%d not found in tree\n", value);
    return NULL;
}

void Tree_inorder_traversal(struct Node *root)
{
    if (!root)
    {
        return;
    }

    Tree_inorder_traversal(root->left);
    printf("Value at node %p: %d\n", root, root->value);
    Tree_inorder_traversal(root->right);
}

int main(void)
{
    struct Tree *myTree = Tree_create();
    Tree_insert(myTree, 36);
    Tree_insert(myTree, 50);
    Tree_insert(myTree, 40);
    Tree_insert(myTree, 12);
    Tree_insert(myTree, 24);

    Tree_inorder_traversal(myTree->root);

    printf("Root Node at %p with value of %d\n", myTree->root, myTree->root->value); 

    struct Node *node = Tree_find(myTree, 40);
    printf("Node with value of %d found at %p\n", node->value, node);

    return 0;
}

# AVL tree for key-value store

This package implements an AVL tree for key-value store.

## What is an AVL tree?

An AVL tree is a balancing search tree.

## Features

| Feature                 | Implemented |
|-------------------------|-------------|
| Insertion               | yes         |
| Deletion                | yes         |
| Searching               | yes         |
| Get inorder successor   | yes         |
| Get inorder predecessor | yes         |
| Inorder traversal       | yes         |
| Preorder traversal      | yes         |
| Postorder traversal     | yes         |
| Balancing               | yes (AVL)   |

Serialization of the tree will not be implemented since the key and value can be an arbitrary type. It would be better
for the caller to implement the serialization that fits the use case the best.

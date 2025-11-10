package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"cmp"
)

type Seq2[K, V any] func(yield func(K, V) bool)

type Node[K cmp.Ordered, V any] struct {
	Key   K
	Value V
	Left  *Node[K, V]
	Right *Node[K, V]
}

type BST[K cmp.Ordered, V any] struct {
	Root *Node[K, V]
}

func (bst *BST[K, V]) Insert(key K, value V) {
	bst.Root = bst.insert(bst.Root, key, value)
}

func (bst *BST[K, V]) insert(node *Node[K, V], key K, value V) *Node[K, V] {
	if node == nil {
		return &Node[K, V]{Key: key, Value: value}
	}
	if key < node.Key {
		node.Left = bst.insert(node.Left, key, value)
	} else if key > node.Key {
		node.Right = bst.insert(node.Right, key, value)
	} else {
		node.Value = value
	}
	return node
}

func (bst *BST[K, V]) Find(key K) (V, bool) {
	return bst.find(bst.Root, key)
}

func (bst *BST[K, V]) find(node *Node[K, V], key K) (V, bool) {
	if node == nil {
		var zero V
		return zero, false
	}
	if key < node.Key {
		return bst.find(node.Left, key)
	} else if key > node.Key {
		return bst.find(node.Right, key)
	} else {
		return node.Value, true
	}
}

func (bst *BST[K, V]) Delete(key K) {
	bst.Root = bst.delete(bst.Root, key)
}

func (bst *BST[K, V]) delete(node *Node[K, V], key K) *Node[K, V] {
	if node == nil {
		return nil
	}
	if key < node.Key {
		node.Left = bst.delete(node.Left, key)
	} else if key > node.Key {
		node.Right = bst.delete(node.Right, key)
	} else {
		if node.Left == nil {
			return node.Right
		} else if node.Right == nil {
			return node.Left
		}

		minNode := bst.findMin(node.Right)
		node.Key = minNode.Key
		node.Value = minNode.Value
		node.Right = bst.delete(node.Right, minNode.Key)
	}
	return node
}

func (bst *BST[K, V]) findMin(node *Node[K, V]) *Node[K, V] {
	current := node
	for current.Left != nil {
		current = current.Left
	}
	return current
}

func (bst *BST[K, V]) InOrder() Seq2[K, V] {
	return func(yield func(K, V) bool) {
		bst.inOrder(bst.Root, yield)
	}
}
func (bst *BST[K, V]) inOrder(node *Node[K, V], yield func(K, V) bool) bool {
	if node == nil {
		return true
	}

	if !bst.inOrder(node.Left, yield) {
		return false
	}

	if !yield(node.Key, node.Value) {
		return false
	}

	return bst.inOrder(node.Right, yield)
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	tree := &BST[int, string]{}

	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Fields(line)
		if len(parts) == 0 {
			continue
		}

		switch parts[0] {
		case "Insert":
			if len(parts) != 3 {
				continue
			}
			key, _ := strconv.Atoi(parts[1])
			value := parts[2]
			tree.Insert(key, value)

		case "Find":
			if len(parts) != 2 {
				continue
			}
			key, _ := strconv.Atoi(parts[1])
			if value, found := tree.Find(key); found {
				fmt.Println(value)
			} else {
				fmt.Println("not found")
			}

		case "Delete":
			if len(parts) != 2 {
				continue
			}
			key, _ := strconv.Atoi(parts[1])
			tree.Delete(key)

		case "InOrder":
			var result []string
			iter := tree.InOrder()
			iter(func(key int, value string) bool {
				result = append(result, fmt.Sprintf("%d: %s\n", key, value))
				return true
			})
			fmt.Println(strings.Join(result, " "))
		}
	}
}

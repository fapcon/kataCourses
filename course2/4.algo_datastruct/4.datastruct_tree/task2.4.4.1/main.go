package main

import (
	"fmt"
	"math/rand"
	"time"
)

type User struct {
	ID   int
	Name string
	Age  int
}

type Node struct {
	index int
	data  *User
	left  *Node
	right *Node
}

type BinaryTree struct {
	root *Node
}

func (t *BinaryTree) insert(user *User) *BinaryTree {
	if t.root == nil {
		t.root = &Node{data: user}
	} else {
		t.root.insert(user)
	}
	return t
}

func (n *Node) insert(user *User) {
	if user.ID < n.data.ID {
		if n.left == nil {
			n.left = &Node{data: user}
		} else {
			n.left.insert(user)
		}
	} else if user.ID > n.data.ID {
		if n.right == nil {
			n.right = &Node{data: user}
		} else {
			n.right.insert(user)
		}
	}
}

func (t *BinaryTree) search(key int) (*User, error) {
	if t.root == nil {
		return nil, fmt.Errorf("1")
	}
	user, err := t.root.search(key)
	if err != nil {
		return nil, err
	}
	return user, err
}

func (n *Node) search(key int) (*User, error) {
	if n == nil {
		return nil, fmt.Errorf("2")
	}
	if key == n.data.ID {
		return n.data, nil
	}
	if key < n.data.ID {
		return n.left.search(key)
	}
	return n.right.search(key)
}

func generateData(n int) *BinaryTree {
	rand.Seed(time.Now().UnixNano())
	bt := &BinaryTree{}
	for i := 0; i < n; i++ {
		val := rand.Intn(100)
		bt.insert(&User{
			ID:   val,
			Name: fmt.Sprintf("User%d", val),
			Age:  rand.Intn(50) + 20,
		})
	}
	return bt
}

//func (root *Node) GetTreeNodeNum() int {
//	if root == nil {
//		return 0
//	} else {
//		return root.left.GetTreeNodeNum() + root.right.GetTreeNodeNum() + 1
//	}
//}
//
//func (root *Node) GetTreeDegree() int {
//	maxDegree := 0
//
//	if root == nil {
//		return maxDegree
//	}
//
//	if root.left.GetTreeDegree() > root.right.GetTreeDegree() {
//		maxDegree = root.left.GetTreeDegree()
//	} else {
//		maxDegree = root.right.GetTreeDegree()
//	}
//
//	return maxDegree + 1
//}

func main() {
	bt := generateData(50)
	user, err := bt.search(30)
	if user != nil {
		fmt.Printf("Найден пользователь: %+v\n", user)
	} else if err != nil {
		fmt.Println("Пользователь не найден")
	}
	//fmt.Println(bt.root.GetTreeNodeNum())
	//
	//fmt.Println(bt.root.GetTreeDegree())
}

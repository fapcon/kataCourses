---
menu:
    after:
        name: binary_tree
        weight: 2
title: Построение сбалансированного бинарного дерева
---

# Задача построить сбалансированное бинарное дерево
Используя AVL дерево, постройте сбалансированное бинарное дерево, на текущей странице.

Нужно написать воркер, который стартует дерево с 5 элементов, и каждые 5 секунд добавляет новый элемент в дерево.

Каждые 5 секунд на странице появляется актуальная версия, сбалансированного дерева.

При вставке нового элемента, в дерево, нужно перестраивать дерево, чтобы оно оставалось сбалансированным.

Как только дерево достигнет 100 элементов, генерируется новое дерево с 5 элементами.

```go
package binary

import (
"fmt"
"math/rand"
"time"
)

type Node struct {
	Key    int
	Height int
	Left   *Node
	Right  *Node
}

type AVLTree struct {
	Root *Node
}

func NewNode(key int) *Node {
	return &Node{Key: key, Height: 1}
}

func (t *AVLTree) Insert(key int) {
	t.Root = insert(t.Root, key)
}

func (t *AVLTree) ToMermaid() string {

}

func height(node *Node) int {

}

func max(a, b int) int {

}

func updateHeight(node *Node) {

}

func getBalance(node *Node) int {

}

func leftRotate(x *Node) *Node {

}

func rightRotate(y *Node) *Node {

}

func insert(node *Node, key int) *Node {

}

func GenerateTree(count int) *AVLTree {

}
```

Не обязательно использовать выше описанный код, можно использовать любую реализацию, выдающую сбалансированное бинарное дерево.

## Mermaid Chart

[MermaidJS](https://mermaid-js.github.io/) is library for generating svg charts and diagrams from text.

## Пример вывода

{{< columns >}}
```tpl
{{</*/* mermaid [class="text-center"]*/*/>}}
graph TD
55 --> 34
34 --> 22
22 --> 16
16 --> 9
9 --> 0
0 --> 4
9 --> 15
16 --> 18
18 --> 17
22 --> 27
27 --> 26
27 --> 30
30 --> 28
30 --> 31
34 --> 45
45 --> 40
40 --> 36
40 --> 42
42 --> 44
45 --> 52
52 --> 49
49 --> 46
52 --> 54
55 --> 73
73 --> 68
68 --> 56
56 --> 62
68 --> 71
71 --> 69
71 --> 72
73 --> 88
88 --> 81
81 --> 74
74 --> 76
81 --> 83
83 --> 82
83 --> 85
88 --> 93
93 --> 92
93 --> 96
96 --> 94
96 --> 97

{{</*/* /mermaid */*/>}}
```

{{< /columns >}}

{{< mermaid >}}
graph TD
55 --> 34
34 --> 22
22 --> 16
16 --> 9
9 --> 0
0 --> 4
9 --> 15
16 --> 18
18 --> 17
22 --> 27
27 --> 26
27 --> 30
30 --> 28
30 --> 31
34 --> 45
45 --> 40
40 --> 36
40 --> 42
42 --> 44
45 --> 52
52 --> 49
49 --> 46
52 --> 54
55 --> 73
73 --> 68
68 --> 56
56 --> 62
68 --> 71
71 --> 69
71 --> 72
73 --> 88
88 --> 81
81 --> 74
74 --> 76
81 --> 83
83 --> 82
83 --> 85
88 --> 93
93 --> 92
93 --> 96
96 --> 94
96 --> 97

{{< /mermaid >}}

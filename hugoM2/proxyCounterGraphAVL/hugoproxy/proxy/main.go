package main

import (
	"fmt"
	"github.com/go-chi/chi"
	_ "github.com/swaggo/http-swagger/example/go-chi/docs"
	"log"
	"math/rand"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
	"strings"
	"time"
)

// @title swagger API
// @version 1.0
// @description API geogrpc

// @host localhost:1313
// @BasePath /

func main() {
	//creds := client.Credentials{
	//	ApiKeyValue:    "c4774295a9b0dc55c9b8e73cb299d5c5faeb2da8",
	//	SecretKeyValue: "efefac572c92aaddf6fe54d04a31f00d72702ed1",
	//}
	//
	//api := dadata.NewCleanApi(client.WithCredentialProvider(&creds))
	//query := "королев пионерская 4"
	//result, err := api.Address(context.Background(), query)
	//if err != nil {
	//	fmt.Errorf("zapros minus")
	//}
	//
	//for _, addr := range result {
	//	fullAddr := &Address{Value: addr.Street} //GeoLon: addr.GeoLon, GeoLat: addr.GeoLat, Source: addr.Source, Result: addr.Result, PostalCode: addr.PostalCode, Country: addr.Country, Region: addr.Region, CityArea: addr.CityArea, CityDistrict: addr.CityDistrict, Street: addr.Street, House: addr.House}
	//	fmt.Println(fullAddr)
	//}

	r := chi.NewRouter()

	// ...
	proxy := NewReverseProxy("hugo", "1313")
	r.Use(proxy.ReverseProxy)
	r.Get("/*", ApiHandler)

	go WorkerTest()

	go graphWorker()

	go treeWorker()

	http.ListenAndServe(":8080", r)
}

type ReverseProxy struct {
	host string
	port string
}

func ApiHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("qwerty"))
}

func NewReverseProxy(host, port string) *ReverseProxy {
	return &ReverseProxy{
		host: host,
		port: port,
	}
}

func (rp *ReverseProxy) ReverseProxy(next http.Handler) http.Handler {
	//targetURL := &url.URL{Scheme: "http", Host: rp.host + ":" + rp.port}
	// Create a reverse proxy with the target URL
	targetURL, _ := url.Parse(fmt.Sprintf("http://" + rp.host + ":" + rp.port))
	proxy := httputil.NewSingleHostReverseProxy(targetURL)
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/api" {
			next.ServeHTTP(w, r)
		} else {
			proxy.ServeHTTP(w, r)
		}
	})
}

const content = `---
menu:
    before:
        name: tasks
        weight: 5
title: Обновление данных в реальном времени
---

# Задача: Обновление данных в реальном времени

Напишите воркер, который будет обновлять данные в реальном времени, на текущей странице.
Текст данной задачи менять нельзя, только время и счетчик.

Файл данной страницы: /app/static/tasks/_index.md

Должен меняться счетчик и время:

Текущее время: %s

Счетчик: %d
`

const graphContent = `---
menu:
    after:
        name: graph
        weight: 1
title: Построение графа
---

# Построение графа

Нужно написать воркер, который будет строить граф на текущей странице, каждые 5 секунд
От 5 до 30 элементов, случайным образом. Все ноды графа должны быть связаны.

` + "```go" + `
type Node struct {
	ID int
	Name string
	Form string // "circle", "rect", "square", "ellipse", "round-rect", "rhombus"
	Links []*Node
}
` + "```" + `

## Mermaid Chart

[MermaidJS](https://mermaid-js.github.io/) is library for generating svg charts and diagrams from text.

## Пример

{{< columns >}}
` + "```tpl" + `
{{</*/* mermaid [class="text-center"]*/*/>}}
%s
{{</*/* /mermaid */*/>}}
` + "```" + `

<--->

{{< mermaid >}}
%s
{{< /mermaid >}}

{{< /columns >}}
`

func WorkerTest() {
	t := time.NewTicker(5 * time.Second)
	var b byte = 0
	for {
		select {
		case <-t.C:
			err := os.WriteFile("/app/static/_index.md", []byte(fmt.Sprintf(content, time.Now().Format("02 Jan 06 15:04 MST"), b)), 0644)
			if err != nil {
				log.Println(err)
			}
			b++
		}
	}
}

type Nodee struct {
	ID    int
	Name  string
	Form  string // "circle", "rect", "square", "ellipse", "round-rect", "rhombus"
	Links []*Nodee
}

func NewGraph() *Nodee {
	n := rand.Intn(6) + 5
	firstNode := &Nodee{
		ID:    0,
		Name:  "Node0",
		Form:  "circle",
		Links: make([]*Nodee, n),
	}

	for i := 0; i < n; i++ {
		firstNode.Links[i] = &Nodee{
			ID:    i,
			Name:  fmt.Sprintf("Node%d", i),
			Form:  "circle",
			Links: make([]*Nodee, 0),
		}
	}
	//
	//	firstNode.Links = append(firstNode.Links, node)
	//
	//	firstNode = node
	//}
	//
	//firstNode.Links = append(firstNode.Links, firstNode)
	//
	//return firstNode
	return firstNode
}

func MermaidGraph(nodes *Nodee) string {
	var builder strings.Builder
	builder.WriteString("graph LR\n")
	builder.WriteString(fmt.Sprintf("%s --> %s\n", nodes.Name, nodes.Links[0].Name))
	for i := 1; i < len(nodes.Links)-1; i++ {
		for j := i + 1; j < len(nodes.Links)-2; j++ {
			builder.WriteString(fmt.Sprintf("%s --> %s\n", nodes.Links[i].Name, nodes.Links[j].Name))
		}
	}
	return builder.String()
}

func graphWorker() {
	t := time.NewTicker(5 * time.Second)
	for {
		select {
		case <-t.C:
			node := NewGraph()
			merm := MermaidGraph(node)
			err := os.WriteFile("/app/static/_graph.md", []byte(fmt.Sprintf(graphContent, merm, merm)), 0644)
			if err != nil {
				log.Println(err)
			}
		}
	}
}

const treeContent = `---
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

` + "```go" + `
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
` + "```" + `

Не обязательно использовать выше описанный код, можно использовать любую реализацию, выдающую сбалансированное бинарное дерево.

## Mermaid Chart

[MermaidJS](https://mermaid-js.github.io/) is library for generating svg charts and diagrams from text.

## Пример вывода

{{< columns >}}
` + "```tpl" + `
{{</*/* mermaid [class="text-center"]*/*/>}}
%s
{{</*/* /mermaid */*/>}}
` + "```" + `

{{< /columns >}}

{{< mermaid >}}
%s
{{< /mermaid >}}
`

//func NewTree() *avl.Tree {
//	n := 5
//	tree := avl.NewWithIntComparator()
//	for i := 1; i < n+1; i++ {
//		tree.Put(i, rand.Intn(100))
//	}
//	return tree
//}
//
//func MermaidTree(n *avl.Tree) string {
//	if n.Root == nil {
//		return ""
//	}
//	if n.Left() != nil {
//		res += fmt.Sprintf("%d --> %d\n", n.Root.Value, n.Left().Value)
//		res += n.Left()
//	}
//	if n.Right != nil {
//		res += fmt.Sprintf("%d --> %d\n", n.Root.Value, n.Right().Value)
//		res += n.Right.MermaidTree()
//	}
//	return res
//
//	//var builder strings.Builder
//	//builder.WriteString("graph LR\n")
//	//builder.WriteString(fmt.Sprintf("%s --> %s\n", tree., nodes.Links[0].Name))
//	//for i := 0; i < len(nodes.Links)-1; i++ {
//	//	for j := i + 1; j < len(nodes.Links)-2; j++ {
//	//		builder.WriteString(fmt.Sprintf("%s --> %s\n", nodes.Links[i].Name, nodes.Links[j].Name))
//	//	}
//	//}
//	//return builder.String()
//}

func treeWorker() {
	t := time.NewTicker(5 * time.Second)
	tree := GenerateTree(5)
	for {
		select {
		case <-t.C:
			tree.Insert(rand.Intn(100))
			merm := tree.ToMermaid()
			err := os.WriteFile("/app/static/_binary.md", []byte(fmt.Sprintf(treeContent, "graph TD\n"+merm, "graph TD\n"+merm)), 0644)
			if err != nil {
				log.Println(err)
			}
		}
	}
}

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

func height(node *Node) int {
	if node == nil {
		return 0
	}
	return node.Height
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func updateHeight(node *Node) {
	node.Height = 1 + max(height(node.Left), height(node.Right))
}

func getBalance(node *Node) int {
	if node == nil {
		return 0
	}
	return height(node.Left) - height(node.Right)
}

func leftRotate(x *Node) *Node {
	y := x.Right
	T2 := y.Left
	y.Left = x
	x.Right = T2
	updateHeight(x)
	updateHeight(y)
	return y
}

func rightRotate(y *Node) *Node {
	x := y.Left
	T2 := x.Right
	x.Right = y
	y.Left = T2
	updateHeight(y)
	updateHeight(x)
	return x
}

func insert(node *Node, key int) *Node {
	if node == nil {
		return NewNode(key)
	}
	if key < node.Key {
		node.Left = insert(node.Left, key)
	} else if key > node.Key {
		node.Right = insert(node.Right, key)
	} else {
		return node
	}
	updateHeight(node)
	balance := getBalance(node)
	if balance > 1 && key < node.Left.Key {
		return rightRotate(node)
	}
	if balance < -1 && key > node.Right.Key {
		return leftRotate(node)
	}
	if balance > 1 && key > node.Left.Key {
		node.Left = leftRotate(node.Left)
		return rightRotate(node)
	}
	if balance < -1 && key < node.Right.Key {
		node.Right = rightRotate(node.Right)
		return leftRotate(node)
	}
	return node
}

func GenerateTree(count int) *AVLTree {
	tree := &AVLTree{}
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < count; i++ {
		tree.Insert(rand.Intn(100))
	}
	return tree
}

func (t *AVLTree) ToMermaid() string {
	return t.Root.MermaidTree()
}

func (n *Node) MermaidTree() string {
	if n == nil {
		return ""
	}
	res := ""
	if n.Left != nil {
		res += fmt.Sprintf("%d --> %d\n", n.Key, n.Left.Key)
		res += n.Left.MermaidTree()
	}
	if n.Right != nil {
		res += fmt.Sprintf("%d --> %d\n", n.Key, n.Right.Key)
		res += n.Right.MermaidTree()
	}
	return res
}

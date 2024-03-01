package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/brianvoe/gofakeit/v6"
	"io/ioutil"
	"time"
)

type DoubleLinkedList struct {
	head *Node // начальный элемент в списке
	tail *Node // последний элемент в списке
	curr *Node // текущий элемент меняется при использовании методов next, prev
	len  int   // количество элементов в списке
}

type LinkedLister interface {
	LoadData(path string) error
	Init(c []Commit)
	Len() int
	SetCurrent(n int) error
	Current() *Node
	Next() *Node
	Prev() *Node
	Insert(n int, c Commit) error
	Push(c Commit) error
	Delete(n int) error
	DeleteCurrent() error
	Index() (int, error)
	GetByIndex(n int) (*Node, error)
	Pop() *Node
	Shift() *DoubleLinkedList
	SearchUUID(uuID string) *Node
	Search(message string) *Node
	Reverse() *DoubleLinkedList
}

func quickSort(arr []Commit, low, high int) []Commit {
	if low < high {
		pi := partition(arr, low, high)
		quickSort(arr, low, pi-1)
		quickSort(arr, pi+1, high)
	}
	return arr
}

func partition(arr []Commit, low, high int) int {
	pivot := arr[high]
	i := low - 1
	for j := low; j < high; j++ {
		if int(arr[j].Date.Unix()) < int(pivot.Date.Unix()) {
			i++
			arr[i], arr[j] = arr[j], arr[i]
		}
	}
	arr[i+1], arr[high] = arr[high], arr[i+1]
	return i + 1
}

func (d *DoubleLinkedList) Init(c []Commit) {
	for _, commit := range c {
		d.Push(commit)
	}
}

func (d *DoubleLinkedList) Push(c Commit) error {
	newNode := &Node{
		data: &c,
	}
	if d.head == nil {
		d.head = newNode
		d.tail = newNode
	} else {
		currentNode := d.head
		//i := 0
		for currentNode.next != nil {
			//currentNode.data = &c[i]
			currentNode = currentNode.next
			//i++
		}
		newNode.prev = currentNode
		currentNode.next = newNode
		d.tail = newNode
	}
	d.len++
	return nil
}

// LoadData загрузка данных из подготовленного json файла
func (d *DoubleLinkedList) LoadData(path string) error {
	// отсортировать список используя самописный QuickSort
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}
	var c []Commit
	err = json.Unmarshal(data, &c)
	if err != nil {
		return err
	}
	quickSort(c, 0, len(c)-1)
	d.Init(c)
	d.curr = d.head
	//d.curr.data = c
	return err
}

func (d *DoubleLinkedList) SetCurrent(n int) error {
	if n < 0 || n > d.len {
		return errors.New("invalid index")
	}
	d.curr = d.head
	for i := 0; i < d.Len(); i++ {
		if i == n {
			return nil
		}
		d.curr = d.curr.next
	}
	return nil
}

func (d *DoubleLinkedList) GetByIndex(n int) (*Node, error) {
	if n < 0 || n > d.len {
		return nil, errors.New("invalid index")
	}
	d.curr = d.head
	for i := 0; i < d.Len(); i++ {
		if i == n {
			return d.curr, nil
		}
		d.curr = d.curr.next
	}
	return nil, nil
}

// Len получение длины списка
func (d *DoubleLinkedList) Len() int {
	return d.len
}

// Current получение текущего элемента
func (d *DoubleLinkedList) Current() *Node {
	return d.curr
}

// Next получение следующего элемента
func (d *DoubleLinkedList) Next() *Node {
	return d.curr.next
}

// Prev получение предыдущего элемента
func (d *DoubleLinkedList) Prev() *Node {
	return d.curr.prev
}

// Insert вставка элемента после n элемента
func (d *DoubleLinkedList) Insert(n int, c Commit) error {
	if n < 0 || n > d.len {
		return errors.New("invalid index")
	}
	newNode := &Node{
		data: &c,
	}
	if d.len == 0 {
		d.head = newNode
		d.tail = newNode
		d.curr = newNode
	} else if n == 0 {
		newNode.next = d.head
		d.head.prev = newNode
		d.head = newNode
	} else if n == d.len {
		newNode.prev = d.tail
		d.tail.next = newNode
		d.tail = newNode
	} else {
		curr := d.head
		for i := 0; i < n-1; i++ {
			curr = curr.next
		}
		newNode.next = curr.next
		newNode.prev = curr
		curr.next.prev = newNode
		curr.next = newNode
	}
	d.len++
	return nil
}

// Delete удаление n элемента
func (d *DoubleLinkedList) Delete(n int) error {
	if n < 0 || n > d.len {
		return errors.New("invalid index")
	}
	if n == 0 {
		d.head.data = nil
		d.head = d.head.next
	}
	d.curr = d.head
	for i := 0; i < d.Len(); i++ {
		if i == n {
			d.curr.data = nil
			d.curr.prev.next = d.curr.prev.next.next
			d.len--
			return nil
		}
	}
	d.curr = d.curr.next
	//d.curr.prev = d.curr.next
	//d.curr.next = d.curr.prev
	return nil
}

// DeleteCurrent удаление текущего элемента
func (d *DoubleLinkedList) DeleteCurrent() error {
	d.curr.data = nil
	d.curr.prev.next = d.curr.prev.next.next
	d.len--
	return nil
}

// Index получение индекса текущего элемента
func (d *DoubleLinkedList) Index() (int, error) {
	index := 0
	curr := d.head
	for curr != nil {
		if curr == d.curr {
			return index, nil
		}
		curr = curr.next
		index++
	}
	return -1, nil
}

// Pop Операция Pop
func (d *DoubleLinkedList) Pop() *Node {
	var res *Node
	res = d.tail
	d.tail = d.tail.prev
	d.tail.next = nil
	d.curr = d.tail
	d.len--
	return res
}

// Shift операция shift
func (d *DoubleLinkedList) Shift() *Node {
	var res *Node
	res = d.head
	d.head = d.head.next
	d.head.prev = nil
	d.curr = d.head
	d.len--
	return res
}

// SearchUUID поиск коммита по uuid
func (d *DoubleLinkedList) SearchUUID(uuID string) *Node {
	curr := d.head
	for {
		if uuID == curr.data.UUID {
			return curr
		} else {
			curr = curr.next
		}
	}
}

// Search поиск коммита по message
func (d *DoubleLinkedList) Search(message string) *Node {
	curr := d.head
	for {
		if message == curr.data.Message {
			return curr
		} else {
			curr = curr.next
		}
	}
}

// Reverse возвращает перевернутый список
func (d *DoubleLinkedList) Reverse() *DoubleLinkedList {
	//var reversed *DoubleLinkedList
	reversed := &DoubleLinkedList{len: d.Len(), head: d.tail, tail: d.head, curr: d.head}
	//var revCommit []Commit
	//reversed.Init(revCommit)
	d.curr = d.tail
	reversed.curr = reversed.head
	for d.curr != nil {
		reversed.curr.data = d.curr.data
		reversed.curr.next = d.curr.prev
		reversed.curr.prev = d.curr.next
		//revCommit = append(revCommit, *d.curr.data)
		reversed.curr = reversed.curr.next
		d.curr = d.curr.prev
	}

	//reversed.Init(revCommit)
	//reversed.curr = reversed.head
	return reversed
}

type Node struct {
	data *Commit
	prev *Node
	next *Node
}

type Commit struct {
	Message string    `json:"message"`
	UUID    string    `json:"uuid"`
	Date    time.Time `json:"date"`
}

func GenerateData(n int) []Commit {
	gofakeit.Seed(time.Now().UnixNano())
	comms := make([]Commit, n)
	for i := range comms {
		comms[i] = Commit{
			UUID:    gofakeit.UUID(),
			Message: gofakeit.Sentence(5),
			Date:    gofakeit.Date(),
		}
	}
	return comms
}

func main() {
	path := "/Users/fc/go/src/gitlab.com/fcons/go-kata/course2/4.algo_datastruct/3.datastruct_list/task2.4.3.1/test.json"
	var dll DoubleLinkedList
	comm := Commit{
		Message: "AAAAAAAAAAAAA",
		UUID:    "BBBBBBBBBBBBB",
		Date:    time.Now(),
	}
	dll.LoadData(path)
	fmt.Println(dll.curr.data)
	fmt.Println(dll.curr.next.data)
	fmt.Println(dll.curr.next.next.data)
	fmt.Println(dll.curr.next.next.next.data)
	fmt.Println(dll.curr.next.next.next.next.data)
	fmt.Println(dll.curr.next.next.next.next.next.data)
	//fmt.Println(dll.curr.next.next.next.next)
	//fmt.Println()
	dll.Insert(1000, comm)
	dll.SetCurrent(3)
	fmt.Println(dll.curr.data)
	fmt.Println(dll.Index())
	//fmt.Println(dll.GetByIndex(4))
	fmt.Println()
	//dll.Delete(3)
	//d, _ := dll.GetByIndex(3)
	//fmt.Println(d.data)
	dll.DeleteCurrent()
	fmt.Println(dll.curr.data)
	d, _ := dll.GetByIndex(3)
	fmt.Println(d.data)

	//fmt.Println()
	//fmt.Println(dll.curr.data.Message)
	//fmt.Println(dll.Search("We need to copy the mobile SMTP pixel!"))
	//fmt.Println(dll.SearchUUID("6956dcd2-875b-11ed-8150-acde48001122"))
	////dll.Search("We need to copy the mobile SMTP pixel!")
	//dll.Insert(1133, comm)
	//fmt.Println(dll.Search("AAAAAAAAAAAAA"))
	//fmt.Println(dll.SearchUUID("BBBBBBBBBBBBB"))
	//fmt.Println(dll.Index())
	//fmt.Println(dll.curr.next.next.next.next.next.data)
	////dll.curr = dll.curr.next.next.next.next.next
	//dll.DeleteCurrent()
	//fmt.Println()
	//fmt.Println(dll.Search("AAAAAAAAAAAAA"))
	//fmt.Println(dll.curr.next.next.next.next.next.data)
	//fmt.Println(dll.Index())
	fmt.Println()
	fmt.Println()
	//dll.SetCurrent(0)
	//fmt.Println(dll.curr.data)
	//reverse := dll.Reverse()
	//reverse.SetCurrent(2000) // ?????????????????????????????????? 2000 elem dont have prev & next pointers
	//fmt.Println(reverse.curr)
	////shift := dll.Shift()
	////fmt.Println(shift)
	//fmt.Println(reverse.Pop().data)
	//fmt.Println(reverse.curr.data)

	dll.SetCurrent(1999)
	fmt.Println(dll.Index())
	//fmt.Println(dll.curr.data)
	//fmt.Println(dll.Pop().data)
	//fmt.Println(dll.curr.data)
	//fmt.Println()
	//dll.SetCurrent(0)
	//fmt.Println(dll.curr.data)
	//fmt.Println(dll.Shift().data)
	//fmt.Println(dll.curr.data)

}

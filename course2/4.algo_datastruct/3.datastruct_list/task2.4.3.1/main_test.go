package main

import (
	"reflect"
	"testing"
	"time"
)

func TestDoubleLinkedList_Current(t *testing.T) {
	type fields struct {
		head *Node
		tail *Node
		curr *Node
		len  int
	}
	tests := []struct {
		name   string
		fields fields
		want   *Node
	}{
		{
			name:   "c1",
			fields: fields{curr: &Node{data: &Commit{UUID: "123", Message: "123", Date: time.Unix(123, 456)}}},
			want:   &Node{data: &Commit{UUID: "123", Message: "123", Date: time.Unix(123, 456)}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &DoubleLinkedList{
				head: tt.fields.head,
				tail: tt.fields.tail,
				curr: tt.fields.curr,
				len:  tt.fields.len,
			}
			if got := d.Current(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Current() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDoubleLinkedList_Delete(t *testing.T) {
	type fields struct {
		head *Node
		tail *Node
		curr *Node
		len  int
	}
	type args struct {
		n int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name:    "c1",
			fields:  fields{head: &Node{next: &Node{}, data: &Commit{UUID: "123", Message: "123", Date: time.Unix(123, 456)}}},
			args:    args{n: 1},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &DoubleLinkedList{
				head: tt.fields.head,
				tail: tt.fields.tail,
				curr: tt.fields.curr,
				len:  tt.fields.len,
			}
			if err := d.Delete(tt.args.n); (err != nil) != tt.wantErr {
				t.Errorf("Delete() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestDoubleLinkedList_DeleteCurrent(t *testing.T) {
	type fields struct {
		head *Node
		tail *Node
		curr *Node
		len  int
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		//{
		//	name:    "c1",
		//	fields:  fields{len: 1, head: &Node{data: &Commit{UUID: "123", Message: "123", Date: time.Unix(123, 456)}}, tail: &Node{data: &Commit{UUID: "123", Message: "123", Date: time.Unix(123, 456)}}, curr: &Node{data: &Commit{UUID: "123", Message: "123", Date: time.Unix(123, 456)}}},
		//	wantErr: false,
		//},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &DoubleLinkedList{
				head: tt.fields.head,
				tail: tt.fields.tail,
				curr: tt.fields.curr,
				len:  tt.fields.len,
			}
			if err := d.DeleteCurrent(); (err != nil) != tt.wantErr {
				t.Errorf("DeleteCurrent() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestDoubleLinkedList_GetByIndex(t *testing.T) {
	type fields struct {
		head *Node
		tail *Node
		curr *Node
		len  int
	}
	type args struct {
		n int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *Node
		wantErr bool
	}{
		{
			name:    "c1",
			fields:  fields{head: &Node{next: &Node{}, data: &Commit{UUID: "123", Message: "123", Date: time.Unix(123, 456)}}},
			args:    args{n: 0},
			want:    nil,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &DoubleLinkedList{
				head: tt.fields.head,
				tail: tt.fields.tail,
				curr: tt.fields.curr,
				len:  tt.fields.len,
			}
			got, err := d.GetByIndex(tt.args.n)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetByIndex() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetByIndex() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDoubleLinkedList_Index(t *testing.T) {
	type fields struct {
		head *Node
		tail *Node
		curr *Node
		len  int
	}
	tests := []struct {
		name    string
		fields  fields
		want    int
		wantErr bool
	}{
		{
			name:    "c1",
			fields:  fields{len: 2, head: &Node{next: &Node{}, data: &Commit{UUID: "123", Message: "123", Date: time.Unix(123, 456)}}, tail: &Node{data: &Commit{UUID: "123", Message: "123", Date: time.Unix(123, 456)}}, curr: &Node{data: &Commit{UUID: "123", Message: "123", Date: time.Unix(123, 456)}}},
			want:    -1,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &DoubleLinkedList{
				head: tt.fields.head,
				tail: tt.fields.tail,
				curr: tt.fields.curr,
				len:  tt.fields.len,
			}
			got, err := d.Index()
			if (err != nil) != tt.wantErr {
				t.Errorf("Index() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Index() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDoubleLinkedList_Init(t *testing.T) {
	type fields struct {
		head *Node
		tail *Node
		curr *Node
		len  int
	}
	type args struct {
		c []Commit
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name:   "c1",
			fields: fields{head: &Node{data: &Commit{UUID: "123", Message: "123", Date: time.Unix(123, 456)}}},
			args:   args{[]Commit{{UUID: "123", Message: "123", Date: time.Unix(123, 456)}}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &DoubleLinkedList{
				head: tt.fields.head,
				tail: tt.fields.tail,
				curr: tt.fields.curr,
				len:  tt.fields.len,
			}
			d.Init(tt.args.c)
		})
	}
}

func TestDoubleLinkedList_Insert(t *testing.T) {
	type fields struct {
		head *Node
		tail *Node
		curr *Node
		len  int
	}
	type args struct {
		n int
		c Commit
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name:    "c1",
			fields:  fields{head: &Node{data: &Commit{UUID: "123", Message: "123", Date: time.Unix(123, 456)}}},
			args:    args{0, Commit{UUID: "123", Message: "123", Date: time.Unix(123, 456)}},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &DoubleLinkedList{
				head: tt.fields.head,
				tail: tt.fields.tail,
				curr: tt.fields.curr,
				len:  tt.fields.len,
			}
			if err := d.Insert(tt.args.n, tt.args.c); (err != nil) != tt.wantErr {
				t.Errorf("Insert() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestDoubleLinkedList_Len(t *testing.T) {
	type fields struct {
		head *Node
		tail *Node
		curr *Node
		len  int
	}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		{
			name:   "c1",
			fields: fields{len: 1, head: &Node{data: &Commit{UUID: "123", Message: "123", Date: time.Unix(123, 456)}}},
			want:   1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &DoubleLinkedList{
				head: tt.fields.head,
				tail: tt.fields.tail,
				curr: tt.fields.curr,
				len:  tt.fields.len,
			}
			if got := d.Len(); got != tt.want {
				t.Errorf("Len() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDoubleLinkedList_LoadData(t *testing.T) {
	type fields struct {
		head *Node
		tail *Node
		curr *Node
		len  int
	}
	type args struct {
		path string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			//[{"message":"We need to program the bluetooth ADP protocol!","uuid":"6956d3f4-875b-11ed-8150-acde48001122","date":"2002-01-19T16:00:14.747278511Z"}
			name:    "c1",
			fields:  fields{head: &Node{data: &Commit{UUID: "6956d3f4-875b-11ed-8150-acde48001122", Message: "We need to program the bluetooth ADP protocol!", Date: time.Unix(123, 456)}}},
			args:    args{"/Users/fc/go/src/gitlab.com/fcons/go-kata/course2/4.algo_datastruct/3.datastruct_list/task2.4.3.1/testing.json"},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &DoubleLinkedList{
				head: tt.fields.head,
				tail: tt.fields.tail,
				curr: tt.fields.curr,
				len:  tt.fields.len,
			}
			if err := d.LoadData(tt.args.path); (err != nil) != tt.wantErr {
				t.Errorf("LoadData() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestDoubleLinkedList_Next(t *testing.T) {
	type fields struct {
		head *Node
		tail *Node
		curr *Node
		len  int
	}
	tests := []struct {
		name   string
		fields fields
		want   *Node
	}{
		//{
		//	name:   "c1",
		//	fields: fields{len: 2, head: &Node{next: &Node{data: &Commit{UUID: "1111", Message: "2222", Date: time.Unix(123, 456)}}, data: &Commit{UUID: "123", Message: "123", Date: time.Unix(123, 456)}}},
		//	want:   &Node{data: &Commit{UUID: "1111", Message: "2222", Date: time.Unix(123, 456)}},
		//},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &DoubleLinkedList{
				head: tt.fields.head,
				tail: tt.fields.tail,
				curr: tt.fields.curr,
				len:  tt.fields.len,
			}
			if got := d.Next(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Next() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDoubleLinkedList_Pop(t *testing.T) {
	type fields struct {
		head *Node
		tail *Node
		curr *Node
		len  int
	}
	tests := []struct {
		name   string
		fields fields
		want   *Node
	}{
		//{
		//	name:   "c1",
		//	fields: fields{tail: &Node{data: &Commit{UUID: "123", Message: "123", Date: time.Unix(123, 456)}}},
		//	want:   &Node{data: &Commit{UUID: "123", Message: "123", Date: time.Unix(123, 456)}},
		//},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &DoubleLinkedList{
				head: tt.fields.head,
				tail: tt.fields.tail,
				curr: tt.fields.curr,
				len:  tt.fields.len,
			}
			if got := d.Pop(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Pop() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDoubleLinkedList_Prev(t *testing.T) {
	type fields struct {
		head *Node
		tail *Node
		curr *Node
		len  int
	}
	tests := []struct {
		name   string
		fields fields
		want   *Node
	}{}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &DoubleLinkedList{
				head: tt.fields.head,
				tail: tt.fields.tail,
				curr: tt.fields.curr,
				len:  tt.fields.len,
			}
			if got := d.Prev(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Prev() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDoubleLinkedList_Push(t *testing.T) {
	type fields struct {
		head *Node
		tail *Node
		curr *Node
		len  int
	}
	type args struct {
		c Commit
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name:    "c1",
			fields:  fields{head: &Node{data: &Commit{UUID: "123", Message: "123", Date: time.Unix(123, 456)}}},
			args:    args{Commit{UUID: "123", Message: "123", Date: time.Unix(123, 456)}},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &DoubleLinkedList{
				head: tt.fields.head,
				tail: tt.fields.tail,
				curr: tt.fields.curr,
				len:  tt.fields.len,
			}
			if err := d.Push(tt.args.c); (err != nil) != tt.wantErr {
				t.Errorf("Push() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestDoubleLinkedList_Reverse(t *testing.T) {
	type fields struct {
		head *Node
		tail *Node
		curr *Node
		len  int
	}
	tests := []struct {
		name   string
		fields fields
		want   *DoubleLinkedList
	}{}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &DoubleLinkedList{
				head: tt.fields.head,
				tail: tt.fields.tail,
				curr: tt.fields.curr,
				len:  tt.fields.len,
			}
			if got := d.Reverse(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Reverse() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDoubleLinkedList_Search(t *testing.T) {
	type fields struct {
		head *Node
		tail *Node
		curr *Node
		len  int
	}
	type args struct {
		message string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Node
	}{}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &DoubleLinkedList{
				head: tt.fields.head,
				tail: tt.fields.tail,
				curr: tt.fields.curr,
				len:  tt.fields.len,
			}
			if got := d.Search(tt.args.message); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Search() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDoubleLinkedList_SearchUUID(t *testing.T) {
	type fields struct {
		head *Node
		tail *Node
		curr *Node
		len  int
	}
	type args struct {
		uuID string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Node
	}{}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &DoubleLinkedList{
				head: tt.fields.head,
				tail: tt.fields.tail,
				curr: tt.fields.curr,
				len:  tt.fields.len,
			}
			if got := d.SearchUUID(tt.args.uuID); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SearchUUID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDoubleLinkedList_SetCurrent(t *testing.T) {
	type fields struct {
		head *Node
		tail *Node
		curr *Node
		len  int
	}
	type args struct {
		n int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name:    "c1",
			fields:  fields{head: &Node{data: &Commit{UUID: "123", Message: "123", Date: time.Unix(123, 456)}}},
			args:    args{0},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &DoubleLinkedList{
				head: tt.fields.head,
				tail: tt.fields.tail,
				curr: tt.fields.curr,
				len:  tt.fields.len,
			}
			if err := d.SetCurrent(tt.args.n); (err != nil) != tt.wantErr {
				t.Errorf("SetCurrent() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestDoubleLinkedList_Shift(t *testing.T) {
	type fields struct {
		head *Node
		tail *Node
		curr *Node
		len  int
	}
	tests := []struct {
		name   string
		fields fields
		want   *Node
	}{}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &DoubleLinkedList{
				head: tt.fields.head,
				tail: tt.fields.tail,
				curr: tt.fields.curr,
				len:  tt.fields.len,
			}
			if got := d.Shift(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Shift() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGenerateData(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want []Commit
	}{}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GenerateData(tt.args.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GenerateData() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_partition(t *testing.T) {
	type args struct {
		arr  []Commit
		low  int
		high int
	}
	tests := []struct {
		name string
		args args
		want int
	}{}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := partition(tt.args.arr, tt.args.low, tt.args.high); got != tt.want {
				t.Errorf("partition() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_quickSort(t *testing.T) {
	type args struct {
		arr  []Commit
		low  int
		high int
	}
	tests := []struct {
		name string
		args args
		want []Commit
	}{}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := quickSort(tt.args.arr, tt.args.low, tt.args.high); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("quickSort() = %v, want %v", got, tt.want)
			}
		})
	}
}

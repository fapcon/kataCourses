package main

import (
	"reflect"
	"testing"
)

func TestMerge(t *testing.T) {
	type args struct {
		arr1 []User
		arr2 []User
	}
	tests := []struct {
		name string
		args args
		want []User
	}{
		{
			name: "c1",
			args: args{arr1: []User{{Name: "aaa", Age: 111, ID: 12}}, arr2: []User{{Name: "bbb", Age: 222, ID: 34}}},
			want: []User{{Name: "aaa", Age: 111, ID: 12}, {Name: "bbb", Age: 222, ID: 34}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Merge(tt.args.arr1, tt.args.arr2); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Merge() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_generateUsers(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want []User
	}{}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := generateUsers(tt.args.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("generateUsers() = %v, want %v", got, tt.want)
			}
		})
	}
}

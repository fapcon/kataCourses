package main

import (
	"hash"
	"hash/crc32"
	"reflect"
	"testing"
	"time"
)

func TestHashMap_Get(t *testing.T) {
	type fields struct {
		hasher hash.Hash
		arr    []*bucket
	}
	type args struct {
		key string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   interface{}
		want1  bool
	}{
		{
			name: "c1",
			fields: fields{
				hasher: crc32.New(crc32.MakeTable(crc32.IEEE)),
				arr:    []*bucket{}},
			args:  args{key: "key"},
			want:  "value",
			want1: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			hm := &HashMap{
				hasher: tt.fields.hasher,
				arr:    tt.fields.arr,
			}
			got, got1 := hm.Get(tt.args.key)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Get() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("Get() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestHashMap_Set(t *testing.T) {
	type fields struct {
		hasher hash.Hash
		arr    []*bucket
	}
	type args struct {
		key   string
		value interface{}
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name: "c1",
			fields: fields{
				hasher: crc32.New(crc32.MakeTable(crc32.IEEE)),
				arr:    []*bucket{}},
			args: args{key: "key", value: "value"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			hm := &HashMap{
				hasher: tt.fields.hasher,
				arr:    tt.fields.arr,
			}
			hm.Set(tt.args.key, tt.args.value)
		})
	}
}

func TestMeassureTime(t *testing.T) {
	type args struct {
		a func()
	}
	tests := []struct {
		name string
		args args
		want time.Duration
	}{
		{
			name: "c1",
			args: args{a: func() {
			}},
			want: 1 * time.Microsecond,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MeassureTime(tt.args.a); got != tt.want {
				t.Errorf("MeassureTime() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewHashMap(t *testing.T) {
	type args struct {
		size int
		opts []func(hm *HashMap)
	}
	tests := []struct {
		name string
		args args
		want *HashMap
	}{
		{
			name: "c1",
			args: args{size: 16, opts: []func(hm *HashMap){WithHashCRC32()}},
			want: &HashMap{hasher: crc32.New(crc32.MakeTable(crc32.IEEE)), arr: make([]*bucket, 16)},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewHashMap(tt.args.size, tt.args.opts...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewHashMap() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestWithHashCRC16(t *testing.T) {
	tests := []struct {
		name string
		want func(hm *HashMap)
	}{
		{
			name: "c1",
			want: WithHashCRC16(),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := WithHashCRC16(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("WithHashCRC16() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestWithHashCRC32(t *testing.T) {
	tests := []struct {
		name string
		want func(hm *HashMap)
	}{
		{
			name: "c1",
			want: WithHashCRC32(),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := WithHashCRC32(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("WithHashCRC32() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestWithHashCRC64(t *testing.T) {
	tests := []struct {
		name string
		want func(hm *HashMap)
	}{
		{
			name: "c1",
			want: WithHashCRC64(),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := WithHashCRC64(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("WithHashCRC64() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestWithHashCRC8(t *testing.T) {
	tests := []struct {
		name string
		want func(hm *HashMap)
	}{
		{
			name: "c1",
			want: WithHashCRC8(),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := WithHashCRC8(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("WithHashCRC8() = %v, want %v", got, tt.want)
			}
		})
	}
}

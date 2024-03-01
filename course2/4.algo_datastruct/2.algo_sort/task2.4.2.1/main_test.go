package main

import (
	"reflect"
	"testing"
	"time"
)

func TestByCount_Len(t *testing.T) {
	tests := []struct {
		name string
		bc   ByCount
		want int
	}{
		{
			name: "c1",
			bc:   []Product{{Count: 1}, {Count: 3}},
			want: 2,
		},
		{
			name: "c2",
			bc:   []Product{{Count: 1}, {Count: 3}, {Count: 6}},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.bc.Len(); got != tt.want {
				t.Errorf("Len() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestByCount_Less(t *testing.T) {
	type args struct {
		i int
		j int
	}
	tests := []struct {
		name string
		bc   ByCount
		args args
		want bool
	}{
		{
			name: "c1",
			bc:   []Product{{Count: 1}, {Count: 3}},
			args: args{i: 0, j: 1},
			want: true,
		},
		{
			name: "c2",
			bc:   []Product{{Count: 1}, {Count: 3}},
			args: args{i: 1, j: 0},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.bc.Less(tt.args.i, tt.args.j); got != tt.want {
				t.Errorf("Less() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestByCount_Swap(t *testing.T) {
	type args struct {
		i int
		j int
	}
	tests := []struct {
		name string
		bc   ByCount
		args args
	}{
		{
			name: "c1",
			bc:   []Product{{Count: 1}, {Count: 3}},
			args: args{i: 0, j: 1},
		},
		{
			name: "c1",
			bc:   []Product{{Count: 1}, {Count: 3}},
			args: args{i: 1, j: 0},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.bc.Swap(tt.args.i, tt.args.j)
		})
	}
}

func TestByCreatedAt_Len(t *testing.T) {
	tests := []struct {
		name string
		bc   ByCreatedAt
		want int
	}{
		{
			name: "c1",
			bc:   []Product{{CreatedAt: time.Now()}, {CreatedAt: time.Now().Add(-24 * time.Hour)}},
			want: 2,
		},
		{
			name: "c2",
			bc:   []Product{{CreatedAt: time.Now()}, {CreatedAt: time.Now().Add(24 * time.Hour)}, {}},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.bc.Len(); got != tt.want {
				t.Errorf("Len() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestByCreatedAt_Less(t *testing.T) {
	type args struct {
		i int
		j int
	}
	tests := []struct {
		name string
		bc   ByCreatedAt
		args args
		want bool
	}{
		{
			name: "c1",
			bc:   []Product{{CreatedAt: time.Now()}, {CreatedAt: time.Now().Add(24 * time.Hour)}},
			args: args{i: 0, j: 1},
			want: true,
		},
		{
			name: "c2",
			bc:   []Product{{CreatedAt: time.Now()}, {CreatedAt: time.Now().Add(24 * time.Hour)}},
			args: args{i: 1, j: 0},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.bc.Less(tt.args.i, tt.args.j); got != tt.want {
				t.Errorf("Less() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestByCreatedAt_Swap(t *testing.T) {
	type args struct {
		i int
		j int
	}
	tests := []struct {
		name string
		bc   ByCreatedAt
		args args
	}{
		{
			name: "c1",
			bc:   []Product{{CreatedAt: time.Now()}, {CreatedAt: time.Now().Add(24 * time.Hour)}},
			args: args{i: 0, j: 1},
		},
		{
			name: "c2",
			bc:   []Product{{CreatedAt: time.Now()}, {CreatedAt: time.Now().Add(24 * time.Hour)}},
			args: args{i: 1, j: 0},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.bc.Swap(tt.args.i, tt.args.j)
		})
	}
}

func TestByPrice_Len(t *testing.T) {
	tests := []struct {
		name string
		bp   ByPrice
		want int
	}{
		{
			name: "c1",
			bp:   []Product{{Count: 1}, {Count: 3}},
			want: 2,
		},
		{
			name: "c2",
			bp:   []Product{{Count: 1}, {Count: 3}, {Count: 6}},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.bp.Len(); got != tt.want {
				t.Errorf("Len() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestByPrice_Less(t *testing.T) {
	type args struct {
		i int
		j int
	}
	tests := []struct {
		name string
		bp   ByPrice
		args args
		want bool
	}{
		{
			name: "c1",
			bp:   []Product{{Price: 1.11}, {Price: 3.33}},
			args: args{i: 0, j: 1},
			want: true,
		},
		{
			name: "c2",
			bp:   []Product{{Price: 1.11}, {Price: 3.33}},
			args: args{i: 1, j: 0},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.bp.Less(tt.args.i, tt.args.j); got != tt.want {
				t.Errorf("Less() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestByPrice_Swap(t *testing.T) {
	type args struct {
		i int
		j int
	}
	tests := []struct {
		name string
		bp   ByPrice
		args args
	}{
		{
			name: "c1",
			bp:   []Product{{Price: 1.11}, {Price: 3.33}},
			args: args{i: 0, j: 1},
		},
		{
			name: "c2",
			bp:   []Product{{Price: 1.11}, {Price: 3.33}},
			args: args{i: 1, j: 0},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.bp.Swap(tt.args.i, tt.args.j)
		})
	}
}

func TestProduct_String(t *testing.T) {
	type fields struct {
		Name      string
		Price     float64
		CreatedAt time.Time
		Count     int
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name:   "c1",
			fields: fields{Name: "aaa", Price: 1.11, CreatedAt: time.Unix(92834, 1234), Count: 10},
			want:   "Name: aaa, Price: 1.110000, Count: 10",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := Product{
				Name:      tt.fields.Name,
				Price:     tt.fields.Price,
				CreatedAt: tt.fields.CreatedAt,
				Count:     tt.fields.Count,
			}
			if got := p.String(); got != tt.want {
				t.Errorf("String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_generateProducts(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want []Product
	}{}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := generateProducts(tt.args.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("generateProducts() = %v, want %v", got, tt.want)
			}
		})
	}
}

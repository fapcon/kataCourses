package main

import (
	"reflect"
	"testing"
)

func TestNewUser(t *testing.T) {
	type args struct {
		id      int
		options []UserOption
	}
	tests := []struct {
		name string
		args args
		want *User
	}{
		{
			name: "c1",
			args: args{id: 12, options: []UserOption{WithUsername("abcd"), WithRole("qwer"), WithEmail("g@vn")}},
			want: &User{ID: 12, Username: "abcd", Role: "qwer", Email: "g@vn"},
		},
		{
			name: "c2",
			args: args{id: 1, options: []UserOption{WithUsername("abcd"), WithRole("qwer"), WithEmail("g@vn")}},
			want: &User{ID: 12, Username: "abcd", Role: "qwer", Email: "g"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewUser(tt.args.id, tt.args.options...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewUser() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestWithEmail(t *testing.T) {
	type args struct {
		value string
	}
	tests := []struct {
		name string
		args args
		want UserOption
	}{
		{
			name: "c1",
			args: args{value: "g@vn"},
			want: func(user *User) {
				user.Email = "g@vn"
			},
		},
		{
			name: "c2",
			args: args{value: "g@vn"},
			want: func(user *User) {
				user.Email = "g"
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := WithEmail(tt.args.value); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("WithEmail() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestWithRole(t *testing.T) {
	type args struct {
		value string
	}
	tests := []struct {
		name string
		args args
		want UserOption
	}{
		{
			name: "c1",
			args: args{value: "qwer"},
			want: func(user *User) {
				user.Role = "qwer"
			},
		},
		{
			name: "c2",
			args: args{value: "qwer"},
			want: func(user *User) {
				user.Role = "q"
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := WithRole(tt.args.value); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("WithRole() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestWithUsername(t *testing.T) {
	type args struct {
		value string
	}
	tests := []struct {
		name string
		args args
		want UserOption
	}{
		{
			name: "c1",
			args: args{value: "qwer"},
			want: func(user *User) {
				user.Username = "qwer"
			},
		},
		{
			name: "c2",
			args: args{value: "qwer"},
			want: func(user *User) {
				user.Role = "q"
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := WithUsername(tt.args.value); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("WithUsername() = %v, want %v", got, tt.want)
			}
		})
	}
}

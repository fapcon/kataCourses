package main

import (
	"reflect"
	"testing"
)

func TestGenerateUserInserts(t *testing.T) {
	type args struct {
		a int
	}
	tests := []struct {
		name string
		args args
		want []string
	}{}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GenerateUserInserts(tt.args.a); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GenerateUserInserts() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGoFakeitGenerator_GenerateFakeUser(t *testing.T) {
	tests := []struct {
		name string
		want User
	}{}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := &GoFakeitGenerator{}
			if got := g.GenerateFakeUser(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GenerateFakeUser() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSQLiteGenerator_CreateInsertSQL(t *testing.T) {
	type args struct {
		model Tabler
	}
	tests := []struct {
		name string
		args args
		want string
	}{}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &SQLiteGenerator{}
			if got := s.CreateInsertSQL(tt.args.model); got != tt.want {
				t.Errorf("CreateInsertSQL() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSQLiteGenerator_CreateTableSQL(t *testing.T) {
	type args struct {
		table Tabler
	}
	tests := []struct {
		name string
		args args
		want string
	}{}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &SQLiteGenerator{}
			if got := s.CreateTableSQL(tt.args.table); got != tt.want {
				t.Errorf("CreateTableSQL() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUser_TableName(t *testing.T) {
	type fields struct {
		ID        int
		FirstName string
		LastName  string
		Email     string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := User{
				ID:        tt.fields.ID,
				FirstName: tt.fields.FirstName,
				LastName:  tt.fields.LastName,
				Email:     tt.fields.Email,
			}
			if got := u.TableName(); got != tt.want {
				t.Errorf("TableName() = %v, want %v", got, tt.want)
			}
		})
	}
}

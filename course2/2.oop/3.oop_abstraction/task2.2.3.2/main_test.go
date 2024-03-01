package main

import (
	"reflect"
	"testing"
)

func TestUser_CreateInsertSQL(t *testing.T) {
	type fields struct {
		ID        int
		FirstName string
		LastName  string
		Email     string
	}
	type args struct {
		model Tabler
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := User{
				ID:        tt.fields.ID,
				FirstName: tt.fields.FirstName,
				LastName:  tt.fields.LastName,
				Email:     tt.fields.Email,
			}
			if got := u.CreateInsertSQL(tt.args.model); got != tt.want {
				t.Errorf("CreateInsertSQL() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUser_CreateTableSQL(t *testing.T) {
	type fields struct {
		ID        int
		FirstName string
		LastName  string
		Email     string
	}
	type args struct {
		table Tabler
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := User{
				ID:        tt.fields.ID,
				FirstName: tt.fields.FirstName,
				LastName:  tt.fields.LastName,
				Email:     tt.fields.Email,
			}
			if got := u.CreateTableSQL(tt.args.table); got != tt.want {
				t.Errorf("CreateTableSQL() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUser_GenerateFakeUser(t *testing.T) {
	type fields struct {
		ID        int
		FirstName string
		LastName  string
		Email     string
	}
	tests := []struct {
		name   string
		fields fields
		want   User
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := User{
				ID:        tt.fields.ID,
				FirstName: tt.fields.FirstName,
				LastName:  tt.fields.LastName,
				Email:     tt.fields.Email,
			}
			if got := u.GenerateFakeUser(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GenerateFakeUser() = %v, want %v", got, tt.want)
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
	}{
		// TODO: Add test cases.
	}
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

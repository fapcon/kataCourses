package main

import (
	"database/sql"
	"reflect"
	"testing"
)

func TestMigrator_Migrate(t *testing.T) {
	type fields struct {
		db           *sql.DB
		sqlGenerator SQLiteGenerator
	}
	type args struct {
		model []Tabler
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &Migrator{
				db:           tt.fields.db,
				sqlGenerator: tt.fields.sqlGenerator,
			}
			if err := m.Migrate(tt.args.model...); (err != nil) != tt.wantErr {
				t.Errorf("Migrate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestNewMigrator(t *testing.T) {
	type args struct {
		db        *sql.DB
		generator SQLiteGenerator
	}
	tests := []struct {
		name string
		args args
		want *Migrator
	}{}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewMigrator(tt.args.db, tt.args.generator); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewMigrator() = %v, want %v", got, tt.want)
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

package main

import (
	"reflect"
	"testing"
)

func TestNewCustomer(t *testing.T) {
	type args struct {
		id      int
		options []Option
	}
	tests := []struct {
		name string
		args args
		want *Customer
	}{
		{
			name: "c1",
			args: args{id: 1, options: []Option{WithAccount(&SavingsAccount{balance: 123.0}), WithName("qwer")}},
			want: &Customer{ID: 1, Account: &SavingsAccount{balance: 123.0}, Name: "qwer"},
		},
		{
			name: "c2",
			args: args{id: 1, options: []Option{WithAccount(&SavingsAccount{balance: 123.0}), WithName("qwer")}},
			want: &Customer{ID: 1, Account: &SavingsAccount{balance: 123}, Name: "qw"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewCustomer(tt.args.id, tt.args.options...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewCustomer() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSavingsAccount_Balance(t *testing.T) {
	type fields struct {
		balance float64
	}
	tests := []struct {
		name   string
		fields fields
		want   float64
	}{
		{
			name:   "c1",
			fields: fields{balance: 123.0},
			want:   123.0,
		},
		{
			name:   "c2",
			fields: fields{balance: 123.0},
			want:   123,
		},
		{
			name:   "c3",
			fields: fields{balance: 123.0},
			want:   124,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &SavingsAccount{
				balance: tt.fields.balance,
			}
			if got := s.Balance(); got != tt.want {
				t.Errorf("Balance() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSavingsAccount_Deposit(t *testing.T) {
	type fields struct {
		balance float64
	}
	type args struct {
		amount float64
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name:   "c1",
			fields: fields{balance: 400.0},
			args:   args{amount: 500.0},
		},
		{
			name:   "c2",
			fields: fields{balance: 400.0},
			args:   args{amount: 200.0},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &SavingsAccount{
				balance: tt.fields.balance,
			}
			s.Deposit(tt.args.amount)
		})
	}
}

func TestSavingsAccount_Withdraw(t *testing.T) {
	type fields struct {
		balance float64
	}
	type args struct {
		amount float64
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name:    "c1",
			fields:  fields{balance: 400.0},
			args:    args{amount: 500.0},
			wantErr: true,
		},
		{
			name:    "c2",
			fields:  fields{balance: 1000.0},
			args:    args{amount: 500.0},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &SavingsAccount{
				balance: tt.fields.balance,
			}
			if err := s.Withdraw(tt.args.amount); (err != nil) != tt.wantErr {
				t.Errorf("Withdraw() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestWithAccount(t *testing.T) {
	type args struct {
		value *SavingsAccount
	}
	tests := []struct {
		name string
		args args
		want Option
	}{
		{
			name: "c1",
			args: args{value: &SavingsAccount{balance: 400.0}},
			want: func(customer *Customer) {
				customer.Account = &SavingsAccount{balance: 400.0}
			},
		},
		{
			name: "c2",
			args: args{value: &SavingsAccount{balance: 400.0}},
			want: func(customer *Customer) {
				customer.Account = &SavingsAccount{balance: 300.0}
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := WithAccount(tt.args.value); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("WithAccount() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestWithName(t *testing.T) {
	type args struct {
		value string
	}
	tests := []struct {
		name string
		args args
		want Option
	}{
		{
			name: "c1",
			args: args{value: "qwer"},
			want: func(customer *Customer) {
				customer.Name = "qwer"
			},
		},
		{
			name: "c2",
			args: args{value: "qwer"},
			want: func(customer *Customer) {
				customer.Name = "q"
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := WithName(tt.args.value); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("WithName() = %v, want %v", got, tt.want)
			}
		})
	}
}

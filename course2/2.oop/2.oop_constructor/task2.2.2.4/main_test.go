package main

import (
	"reflect"
	"testing"
)

func TestLogSystem_Log(t *testing.T) {
	type fields struct {
		logger FileLogger
	}
	type args struct {
		m string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &LogSystem{
				logger: tt.fields.logger,
			}
			l.Log(tt.args.m)
		})
	}
}

func TestNewLogSystem(t *testing.T) {
	type args struct {
		option LogOption
	}
	tests := []struct {
		name string
		args args
		want *LogSystem
	}{}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewLogSystem(tt.args.option); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewLogSystem() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestWithLogger(t *testing.T) {
	type args struct {
		value FileLogger
	}
	tests := []struct {
		name string
		args args
		want LogOption
	}{}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := WithLogger(tt.args.value); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("WithLogger() = %v, want %v", got, tt.want)
			}
		})
	}
}

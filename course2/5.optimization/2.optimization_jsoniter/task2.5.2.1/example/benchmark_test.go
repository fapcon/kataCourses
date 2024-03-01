package example

import (
	"encoding/json"
	jsoniter "github.com/json-iterator/go"
	"github.com/mailru/easyjson"
	"testing"
)

func BenchmarkJSONEasyMarshal(b *testing.B) {
	var (
		err    error
		target []byte
	)
	px := &Record{
		ID:    0,
		Email: "2",
		//Amount:    0,
		Profile:   Profile{},
		Password:  "23",
		Username:  "23",
		CreatedAt: "123",
		CreatedBy: "123",
	}
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		target, err = easyjson.Marshal(px)
		if err != nil {
			panic(err)
		}
		_ = target
	}
}
func BenchmarkJSONIterMarshal(b *testing.B) {
	var (
		err    error
		target []byte
	)
	px := &Record{
		ID:    0,
		Email: "2",
		//Amount:    0,
		Profile:   Profile{},
		Password:  "23",
		Username:  "23",
		CreatedAt: "123",
		CreatedBy: "123",
	}
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		target, err = jsoniter.Marshal(px)
		if err != nil {
			panic(err)
		}
		_ = target
	}
}

func BenchmarkJSONMarshal(b *testing.B) {
	var (
		err    error
		target []byte
	)
	px := &Record{
		ID:    0,
		Email: "2",
		//Amount:    0,
		Profile:   Profile{},
		Password:  "23",
		Username:  "23",
		CreatedAt: "123",
		CreatedBy: "123",
	}
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		target, err = json.Marshal(px)
		if err != nil {
			panic(err)
		}
		_ = target
	}
}

//
//func BenchmarkJSONEasyUnmarshal(b *testing.B) {
//	// JSON data as a byte slice
//	jsonData := []byte(`{"name": "John", "age": 30}`)
//
//	for i := 0; i < b.N; i++ {
//		// Unmarshal the JSON data
//		var data easyjson.Unmarshaler
//		err := easyjson.Unmarshal(jsonData, data)
//		if err != nil {
//			b.Errorf("Error: %s", err)
//		}
//	}
//}

func BenchmarkJSONIterUnmarshal(b *testing.B) {
	// JSON data as a byte slice
	jsonData := []byte(`{"name": "John", "age": 30}`)

	for i := 0; i < b.N; i++ {
		// Unmarshal the JSON data
		var data map[string]interface{}
		err := jsoniter.Unmarshal(jsonData, &data)
		if err != nil {
			b.Errorf("Error: %s", err)
		}
	}
}

func BenchmarkJSONUnmarshal(b *testing.B) {
	// JSON data as a byte slice
	jsonData := []byte(`{"name": "John", "age": 30}`)

	for i := 0; i < b.N; i++ {
		// Unmarshal the JSON data
		var data map[string]interface{}
		err := json.Unmarshal(jsonData, &data)
		if err != nil {
			b.Errorf("Error: %s", err)
		}
	}
}

package example

import (
	"encoding/json"
	jsoniter "github.com/json-iterator/go"
	"github.com/mailru/easyjson"
)

//go:generate easytags $GOFILE json,db,db_ops,db_type,db_default,db_index
type Welcome struct {
	Records     []Record `json:"records"`
	Skip        int64    `json:"skip"`
	Limit       int64    `json:"limit"`
	TotalAmount int64    `json:"totalAmount"`
}

type Record struct {
	ID        int64     `json:"id"`
	Email     string    `json:"email"`
	Profile   Profile   `json:"profile"`
	Password  string    `json:"password"`
	Username  string    `json:"username"`
	CreatedAt string    `json:"createdAt"`
	CreatedBy CreatedBy `json:"createdBy"`
}

type Profile struct {
	Dob        string `json:"dob"`
	Avatar     string `json:"avatar"`
	LastName   string `json:"lastName"`
	FirstName  string `json:"firstName"`
	StaticData string `json:"staticData"`
}

type CreatedBy string

////go:generate easyjson -no_std_marshalers client.go

type easyJ struct {
}

////easyjson:json
//type JSONData struct {
//	Data []string
//}

type iterJ struct {
}

type stand struct {
}

const (
	System CreatedBy = "system"
)

func main() {

}

type MarshalUnmarshaler interface {
	Marshal(v any) ([]byte, error)
	Unmarshal(data []byte, v any) error
}

func (iJ *iterJ) Marshal(v any) ([]byte, error) {
	return jsoniter.Marshal(v)
}

func (iJ *iterJ) Unmarshal(data []byte, v any) error {
	return jsoniter.Unmarshal(data, &v)
}

func (eJ *easyJ) Marshal(v easyjson.Marshaler) ([]byte, error) {
	rawBytes, err := easyjson.Marshal(v) //easyjson.Marshal(v)
	return rawBytes, err

}

func (eJ *easyJ) Unmarshal(data []byte, v easyjson.Unmarshaler) error {
	err := easyjson.Unmarshal(data, v)
	return err
}

func (s *stand) Marshal(v any) ([]byte, error) {
	return json.Marshal(v)
}

func (s *stand) Unmarshal(data []byte, v any) error {
	return json.Unmarshal(data, &v)
}

package models

type Pet struct {
	ID     int64  `json:"id" db:"id" db_type:"SERIAL PRIMARY KEY"`
	Name   string `json:"name" db:"name" db_type:"VARCHAR(100)"`
	Status string `json:"status" db:"status" db_type:"VARCHAR(100)"`
}

func (p Pet) TableName() string {
	return "pets"
}

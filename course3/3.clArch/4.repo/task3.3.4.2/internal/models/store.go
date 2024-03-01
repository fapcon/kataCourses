package models

type Order struct {
	ID       int64  `json:"id" db:"id" db_type:"SERIAL PRIMARY KEY"`
	PetID    int64  `json:"petId" db:"pet_id" db_type:"INTEGER"`
	Quantity int64  `json:"quantity" db:"quantity" db_type:"INTEGER"`
	ShipDate string `json:"shipDate" db:"ship_date" db_type:"VARCHAR(100)"`
	Status   string `json:"status" db:"status" db_type:"VARCHAR(100)"`
	Complete bool   `json:"complete" db:"complete" db_type:"BOOLEAN"`
}

func (o Order) TableName() string {
	return "orders"
}

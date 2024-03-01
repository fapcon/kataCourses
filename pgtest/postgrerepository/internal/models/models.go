package models

type User struct {
	ID        int    `db:"id" db_type:"SERIAL PRIMARY KEY" json:"id"`
	FirstName string `db:"first_name" db_type:"VARCHAR(100)" json:"first_name"`
	LastName  string `db:"last_name" db_type:"VARCHAR(100)" json:"last_name"`
	Username  string `db:"username" db_type:"VARCHAR(100)" json:"username"`
	Email     string `db:"email" db_type:"VARCHAR(100)" json:"email"`
	Address   string `db:"address" db_type:"VARCHAR(100)" json:"address"`
	Deleted   bool   `db:"deleted" db_type:"VARCHAR(100)" json:"deleted"`
}

func (m User) TableName() string {
	return "users"
}

type Conditions struct {
	Equal       map[string]interface{} `json:"equal"`
	NotEqual    map[string]interface{} `json:"not_equal"`
	Order       []*Order               `json:"order"`
	LimitOffset *LimitOffset           `json:"limit_offset"`
	ForUpdate   bool                   `json:"for_update"`
	Upsert      bool                   `json:"upsert"`
}

// Order is ... JSON
// swagger:model
type Order struct {
	Field string `json:"field"`
	Asc   bool   `json:"asc"`
}

// LimitOffset is JSON ...
// swagger:model
type LimitOffset struct {
	Offset int64 `json:"offset"`
	Limit  int64 `json:"limit"`
}

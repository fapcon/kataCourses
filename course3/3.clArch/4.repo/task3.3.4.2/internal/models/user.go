package models

type User struct {
	ID         int64  `json:"id" db:"id" db_type:"SERIAL PRIMARY KEY"`
	Username   string `json:"username" db:"username" db_type:"VARCHAR(100)"`
	FirstName  string `json:"firstName" db:"first_name" db_type:"VARCHAR(100)"`
	LastName   string `json:"lastName" db:"last_name" db_type:"VARCHAR(100)"`
	Email      string `json:"email" db:"email" db_type:"VARCHAR(100)"`
	Password   string `json:"password" db:"password" db_type:"VARCHAR(100)"`
	Phone      string `json:"phone" db:"phone" db_type:"VARCHAR(100)"`
	UserStatus int64  `json:"userStatus" db:"user_status" db_type:"INTEGER"`
	Token      string `json:"token" db:"token" db_type:"VARCHAR(100)"`
}

func (u User) TableName() string {
	return "users"
}

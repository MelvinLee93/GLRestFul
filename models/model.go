package models
import "time"

type Users struct {
    Id          int    `json:"id" db:"id"`
    Username    string `json:"username" db:"username"`
    Password    string `json:"password" db:"password"`
    Email       string `json:"email" db:"email"`
    Created_At  time.Time `json:"created_at" db:"created_at"`
	Updated_At  time.Time `json:"updated_at" db:"updated_at"`
}
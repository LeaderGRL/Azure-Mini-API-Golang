package API

import "time"

type Users struct {
	Id         int       `json:"id"`
	Username   string    `json:"username"`
	Password   string    `json:"password"`
	Email      string    `json:"email"`
	Created_At time.Time `json:"created_at"`
	Updated_At time.Time `json:"updated_at"`
}

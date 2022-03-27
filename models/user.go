package models

type User struct {
	Id       uint   `json:"id"`
	UserName string `json:"username"`
	Email    string `gorm:"unique" json:"email"`
	Password []byte `json:"password"`
}
type StatusResponse struct {
	Message string                 `json:"message"`
	Status  int16                  `json:"status"`
	Data    map[string]interface{} `json:"data"`
}

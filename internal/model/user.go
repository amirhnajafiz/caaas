package model

type User struct {
	ID          uint   `json:"id"`
	AppID       uint   `json:"app_id"`
	Credentials string `json:"credentials"`
}

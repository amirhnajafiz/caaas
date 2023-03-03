package response

import "time"

type App struct {
	Name      string    `json:"name"`
	AppKey    string    `json:"app_key"`
	URL       string    `json:"url"`
	CreatedAt time.Time `json:"created_at"`
}

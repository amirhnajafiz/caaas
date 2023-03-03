package response

type App struct {
	Name      string `json:"name"`
	AppKey    string `json:"app_key"`
	URL       string `json:"url"`
	CreatedAt string `json:"created_at"`
}

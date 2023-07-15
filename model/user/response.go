package user

type Response struct {
	StatusCode int    `json:"statusCode"`
	Error      string `json:"error"`
	Message    string `json:"message"`
	Data       *User  `json:"data"`
}

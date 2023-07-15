package user

type User struct {
	ID           int    `json:"id"`
	UserName     string `json:"userName"`
	UserAge      int    `json:"userAge"`
	UserLocation string `json:"userLocation"`
}

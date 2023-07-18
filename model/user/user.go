package user

type User struct {
	UserId       int    `json:"userId"`
	UserName     string `json:"userName"`
	UserAge      int    `json:"userAge"`
	UserLocation string `json:"userLocation"`
}

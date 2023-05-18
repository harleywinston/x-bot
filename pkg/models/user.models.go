package models

type UserModel struct {
	ChatID     int64  `json:"chatID"`
	Username   string `json:"username"`
	Email      string `json:"email"`
	FuckedUser bool   `json:"fucked_user"`
}

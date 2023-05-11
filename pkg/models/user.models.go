package models

type UserModel struct {
	Username   string `json:"username"`
	Email      string `json:"email"`
	FuckedUser bool   `json:"fucked_user"`
}

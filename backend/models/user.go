package models

//User ...
type User struct {
	ID           int32
	Username     string
	Email        string
	PasswordHash string
}

type Login struct {
	Login        string
	PasswordHash string
}

type LoginGoodResponse struct {
	Status int32
	ID     int32
}

type LoginBadResponse struct {
	Error string
	ID    int32
}

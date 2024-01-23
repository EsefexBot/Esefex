package userdb

type User struct {
	ID     string  `json:"id"`
	Tokens []Token `json:"tokens"`
}

type Token string

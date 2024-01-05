package userdb

import "esefexapi/types"

type User struct {
	ID     types.UserID `json:"id"`
	Tokens []Token      `json:"tokens"`
}

type Token string

package d

import "street/ent"

type Token struct {
	*ent.Token
}

func TokenFromEnt(token *ent.Token) *Token {
	return &Token{
		Token: token,
	}
}

type TokenForm struct {
	Token string `header:"Authorization" binding:"required"`
}

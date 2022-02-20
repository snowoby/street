package d

import "street/ent"

type Token struct {
	*ent.Token
	NoEdges
	ValueType
}

func TokenFromEnt(token *ent.Token) *Token {
	return &Token{
		Token:     token,
		ValueType: ValueType{"token"},
	}
}

type TokenForm struct {
	Token string `header:"Authorization" binding:"required"`
}

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

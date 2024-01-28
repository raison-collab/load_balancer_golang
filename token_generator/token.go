package token_generator

import (
	"crypto/md5"
	"encoding/hex"
)

type Token struct {
	Token string `json:"token"`
}

func (t *Token) GenerateToken(hashedData string) {
	token := md5.New()
	token.Write([]byte(hashedData))

	t.Token = hex.EncodeToString(token.Sum(nil))
}

package token_generator

import (
	"golang.org/x/crypto/bcrypt"
	"log"
)

type Hash struct {
	Hash string `json:"hash"`
}

func (h *Hash) GenerateHash(data string) {
	hash, err := bcrypt.GenerateFromPassword([]byte(data), bcrypt.DefaultCost)
	if err != nil {
		log.Fatal(err)
	}

	h.Hash = string(hash)
}

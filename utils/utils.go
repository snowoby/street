package utils

import (
	"crypto/rand"
	"github.com/alexedwards/argon2id"
	"math/big"
)

func HashPassword(password string) (string, error) {
	hash, err := argon2id.CreateHash(password, argon2id.DefaultParams)
	return hash, err
}

func CheckPassword(password, hash string) bool {
	match, err := argon2id.ComparePasswordAndHash(password, hash)
	return match && err == nil
}

//func Sha256Hash(plain string) string {
//	hash := sha256.Sum256([]byte(plain))
//	hashString := hex.EncodeToString(hash[:])
//	return hashString
//}

var letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

func RandomString(n int) string {
	ret := make([]byte, n)
	for i := 0; i < n; i++ {
		num, err := rand.Int(rand.Reader, big.NewInt(int64(len(letters))))
		if err != nil {
			panic(err)
		}
		ret[i] = letters[num.Int64()]
	}
	return string(ret)
}

func StrongPassword(rawPassword string) error {
	//TODO
	return nil
}

func Encrypt(rawPassword string) (string, error) {
	return HashPassword(rawPassword)
}

func Validate(rawPassword string, recordPassword string) bool {
	return CheckPassword(rawPassword, recordPassword)
}

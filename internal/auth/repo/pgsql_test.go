package repo

import (
	"BookStore/pkg/utils"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"testing"
)

const PwSaltBytes = 32

func Test_genPass(t *testing.T) {
	login := "admin"
	pass := "admin"
	salt := utils.RandString(64)

	h := sha256.New()
	h.Write([]byte(login + salt + pass))
	hashedPassword := hex.EncodeToString(h.Sum(nil))

	fmt.Println(login, pass, salt, hashedPassword)

	login = "user"
	pass = "user"
	salt = utils.RandString(64)

	h = sha256.New()
	h.Write([]byte(login + salt + pass))
	hashedPassword = hex.EncodeToString(h.Sum(nil))

	fmt.Println(login, pass, salt, hashedPassword)
}

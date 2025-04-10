package models

import (
	"bytes"
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"errors"
	"fmt"
	uuid2 "github.com/gofrs/uuid/v5"
	"os"
	"strings"

	"golang.org/x/crypto/pbkdf2"
	"log/slog"
)

type AppUser struct {
	Uuid     string
	Username string
	Password string
}

func NewAppUser(username string, password string) *AppUser {
	uuid, _ := uuid2.NewV4()
	user := &AppUser{Uuid: uuid.String(), Username: username}
	user.SetPassword(password)
	return user
}

func Base64Dec(encoded string) ([]byte, error) {
	decoded, err := base64.StdEncoding.DecodeString(encoded)
	if err != nil {
		fmt.Println("Error decoding:", err)
		return nil, errors.New("Error decoding string " + encoded)
	}
	return decoded, nil
}

func MakeSalt() ([]byte, error) {
	randomBytes := make([]byte, 16)
	_, err := rand.Read(randomBytes)
	if err != nil {
		return nil, err
	}
	return randomBytes, nil
}

func (u *AppUser) SetPassword(password string) error {
	salt, err := MakeSalt()
	if err != nil {
		return err
	}
	passwordBytes := []byte(password)
	passwordEncrypted := pbkdf2.Key(passwordBytes, salt, 4096, 32, sha256.New)
	u.Password = base64.StdEncoding.EncodeToString(salt) + ":" + base64.StdEncoding.EncodeToString(passwordEncrypted)
	return nil
}

func (u *AppUser) GetPassword() string {
	return u.Password
}

func (u *AppUser) CheckPassword(givenPassword string) (bool, error) {
	log := slog.New(slog.NewJSONHandler(os.Stdout, nil))

	passString := strings.Split(u.Password, ":")
	// len
	if len(passString) != 2 {
		log.Info("Password is malformed for ", "username", u.Username)
		return false, errors.New("givenPassword is malformed")
	}

	salt, origPassHash := passString[0], passString[1]
	passwordBytes := []byte(givenPassword)
	saltBytes, err := Base64Dec(salt)
	if err != nil {
		return false, err
	}
	origPassBytes, err := Base64Dec(origPassHash)
	if err != nil {
		return false, err
	}
	salted := pbkdf2.Key(passwordBytes, saltBytes, 4096, 32, sha256.New)
	if bytes.Equal(salted, origPassBytes) {
		return true, nil
	}
	return false, errors.New("givenPassword is incorrect")
}

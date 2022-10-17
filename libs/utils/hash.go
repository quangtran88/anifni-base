package baseUtils

import "golang.org/x/crypto/bcrypt"

type HashGenerator struct {
}

func (h HashGenerator) HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func (h HashGenerator) CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

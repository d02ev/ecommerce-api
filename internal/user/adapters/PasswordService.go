package adapters

import "golang.org/x/crypto/bcrypt"

type PasswordService struct {}

func NewPasswordService() *PasswordService {
	return &PasswordService{}
}

func (ps *PasswordService) Hash(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost);
	return string(hash), err
}

func (ps *PasswordService) Compare(hash, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
package password

import "golang.org/x/crypto/bcrypt"

func HashPassword (rawPassword string) (string, error){
	hash, err := bcrypt.GenerateFromPassword([]byte(rawPassword), bcrypt.DefaultCost)
	return string(hash), err
}

func CheckPassword (rawPassword, hashedPassword string) bool{
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(rawPassword))
	return err == nil
}
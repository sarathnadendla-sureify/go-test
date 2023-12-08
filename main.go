package main

import (
	"fmt"
	"log"

	"golang.org/x/crypto/bcrypt"
)

func main() {

	password := "123456789"

	hashPass, err := hashPassword(password)
	if err != nil {
		panic(err)
	}
	fmt.Println("Hash Password", hashPass)
	err = comparePassword(password, hashPass)
	if err != nil {
		log.Fatalln("Error", err)

	}
	log.Println("Logged In")

}

func hashPassword(password string) ([]byte, error) {
	bs, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, fmt.Errorf("Error while generatig bcrypt from hashpassword %w", err)
	}
	return bs, nil

}

func comparePassword(password string, hashPassword []byte) error {
	err := bcrypt.CompareHashAndPassword(hashPassword, []byte(password))
	if err != nil {
		return fmt.Errorf("Invalid Password: %w", err)
	}
	return nil
}

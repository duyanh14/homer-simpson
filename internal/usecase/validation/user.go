package validation

import (
	"errors"
	"net/mail"
	"regexp"
)

func ValidationEmail(email string) error {
	if email == "" {
		return errors.New("email empty")
	}
	_, err := mail.ParseAddress(email)
	if err != nil {
		return errors.New("email invalid")
	}
	return nil
}

func ValidationPhone(phone string) error {
	if phone == "" {
		return errors.New("phone empty")
	}
	reg := regexp.MustCompile(`(0|\+84)[0-9]{9}$`)
	if !reg.MatchString(phone) {
		return errors.New("phone invalid")
	}
	return nil
}

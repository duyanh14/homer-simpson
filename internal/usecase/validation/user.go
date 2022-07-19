package validation

import (
	"errors"
	"net/mail"
	"regexp"
	"unicode"

	"golang.org/x/crypto/bcrypt"
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

/*
 * Password rules:
 * at least 7 letters
 * at max 20 letters
 * at least 1 number
 * at least 1 upper case
 * at least 1 special character
 */
func ValidatorPw(pw string) error {
	var (
		hasUpper   = false
		hasLower   = false
		hasNumber  = false
		hasSpecial = false
	)
	if len(pw) < 7 {
		return errors.New("password least 7 letters")
	}
	if len(pw) > 20 {
		return errors.New("password max 20 letters")
	}
	for _, charStr := range pw {
		switch {
		case unicode.IsUpper(charStr):
			hasUpper = true
		case unicode.IsLower(charStr):
			hasLower = true
		case unicode.IsNumber(charStr):
			hasNumber = true
		case unicode.IsPunct(charStr) || unicode.IsSymbol(charStr):
			hasSpecial = true
		}
	}
	if !hasUpper {
		return errors.New("password least 1 upper case")
	}
	if !hasNumber {
		return errors.New("password least 1 number")
	}
	if !hasLower {
		return errors.New("password least 1 lower case")
	}
	if !hasSpecial {
		return errors.New("password least 1 special character")
	}

	return nil
}

func HashPw(pw string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(pw), 14)
	return string(bytes), err
}

func CheckPasswordHash(hashedPassword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

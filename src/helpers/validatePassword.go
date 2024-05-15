package helpers

import (
	"fmt"
	"regexp"
)

func ValidatePassword(password string) (interface{}, error) {

	if len(password) < 8 {
		return nil, fmt.Errorf("password must be at least 8 characters")
	}
	// isValidate := regexp.MustCompile(`^(?=.*[A-Z])(?=.*[a-z])(?=.*[0-9])(?=.*[!@#\$%\^&\*]).+$`)

	uppercasePassword := regexp.MustCompile(`[A-Z]`)
	lowercasePassword := regexp.MustCompile(`[a-z]`)
	numberPassword := regexp.MustCompile(`[0-9]`)
	charPassword := regexp.MustCompile(`[_!@#\$%\^&\*]`)
	if !uppercasePassword.MatchString(password){
		return nil, fmt.Errorf("password must contain at least one uppercase letter")
	}
	if !lowercasePassword.MatchString(password){
		return nil, fmt.Errorf("password must contain at least one lowercase letter")
	}
	if !numberPassword.MatchString(password){
		return nil, fmt.Errorf("password must contain at least one digit number")
	}
	if !charPassword.MatchString(password){
		return nil, fmt.Errorf("password must contain at least one special character")
	}
	return true, nil

}
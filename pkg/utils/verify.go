package utils

import "regexp"

func VerifyPhone(phone string) bool {
	regex := regexp.MustCompile(`^$`)
	return regex.MatchString(phone)

}

func VerifyEmail(email string) bool {
	regex := regexp.MustCompile(`^$`)
	return regex.MatchString(email)
}

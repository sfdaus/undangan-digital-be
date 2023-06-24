package utils

import (
	"strings"
)

func CensorName(name string) string {
	parts := strings.Split(name, " ")
	for i, part := range parts {
		firstLetter := string(part[0])
		censoredPart := firstLetter + strings.Repeat("*", len(part)-1)
		parts[i] = censoredPart
	}
	return strings.Join(parts, " ")
}

func InitialName(name string) string {
	var initials string
	nameSplit := strings.Split(name, " ")
	for i := 0; i < len(nameSplit); i++ {
		if i >= 2 {
			break
		}
		initials += strings.ToUpper(string(nameSplit[i][0]))
	}
	return initials
}

func CensorEmail(email string) string {
	atIndex := strings.Index(email, "@")
	if atIndex <= 1 {
		// Invalid email address, return the original value
		return email
	}
	username := email[:atIndex]
	if len(username) <= 2 {
		// Username too short, return the original value
		return email
	}
	first := string(username[0])
	last := string(username[len(username)-1])
	censored := first + strings.Repeat("*", 4) + last
	return censored + email[atIndex:]
}

func CensorPhoneNumber(phoneNumber string) string {
	length := len(phoneNumber)
	if length <= 4 {
		return phoneNumber
	}
	prefix := phoneNumber[0:2]
	suffix := phoneNumber[length-2:]
	asterisks := strings.Repeat("*", length-4)
	return prefix + asterisks + suffix
}


func CensorKTPNumber(ktpNumber string) string {
	length := len(ktpNumber)
	if length != 16 {
		return ktpNumber
	}
	prefix := ktpNumber[0:2]
	suffix := ktpNumber[14:]
	asterisks := strings.Repeat("*", 16-4)
	return prefix + asterisks + suffix
}
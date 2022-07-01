package misc

import (
	"math/rand"
	"regexp"
	"strings"
)

// RandomString returns randomized string
func RandomString(letters []byte, n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

// IsMailDomainOnly メールドメイン ( @ 含む ) かどうかを検証します
func IsMailDomainOnly(mail string) bool {
	reg := regexp.MustCompile(`^@[a-zA-Z0-9-]+(?:\.[a-zA-Z0-9-]+)*$`)
	return reg.MatchString(mail)
}

// IsSameMailDomain メールアドレスのドメイン部分が同一かどうかを判定します
func IsSameMailDomain(mail1, mail2 string) bool {
	arr1 := strings.Split(mail1, "@")
	arr2 := strings.Split(mail2, "@")
	if len(arr1) != 2 || len(arr2) != 2 {
		return false
	}
	return arr1[1] == arr2[1]
}

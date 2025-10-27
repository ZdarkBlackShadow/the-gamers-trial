package util

import (
	"crypto/rand"
	"crypto/subtle"
	"encoding/base64"
	"errors"
	"fmt"
	"net"
	"os"
	"regexp"
	"strings"

	"golang.org/x/crypto/argon2"
)

const saltLength = 16


func generateSalt() ([]byte, error) {
	salt := make([]byte, saltLength)
	if _, err := rand.Read(salt); err != nil {
		return nil, err
	}
	return salt, nil
}

func HashPassword(password string) (string, error) {
	pepper := os.Getenv("PEPPER")
	if pepper == "" {
		return "", errors.New("PEPPER undefined in .env")
	}

	salt, err := generateSalt()
	if err != nil {
		return "", err
	}

	hash := argon2.IDKey([]byte(password+pepper), salt, 1, 64*1024, 4, 32)

	b64Salt := base64.RawStdEncoding.EncodeToString(salt)
	b64Hash := base64.RawStdEncoding.EncodeToString(hash)

	return fmt.Sprintf("%s$%s", b64Salt, b64Hash), nil
}
func VerifyPassword(password, hashed string) (bool, error) {
	pepper := os.Getenv("PEPPER")
	if pepper == "" {
		return false, errors.New("PEPPER undefined in .env")
	}

	parts := strings.SplitN(hashed, "$", 2)
	if len(parts) != 2 {
		return false, errors.New("malformed hash")
	}

	salt, err := base64.RawStdEncoding.DecodeString(parts[0])
	if err != nil {
		return false, errors.New("decoding error salt:" + err.Error())
	}

	expectedHash, err := base64.RawStdEncoding.DecodeString(parts[1])
	if err != nil {
		return false, errors.New("decoding error hash:" + err.Error())
	}

	newHash := argon2.IDKey([]byte(password+pepper), salt, 1, 64*1024, 4, 32)

	if len(newHash) != len(expectedHash) {
		return false, nil
	}

	match := subtle.ConstantTimeCompare(newHash, expectedHash) == 1
	return match, nil
}

func IsPasswordStrong(password string) error {
	if len(password) < 8 {
		return errors.New("the password must have minimum 8 character")
	}

	if !regexp.MustCompile(`[A-Z]`).MatchString(password) {
		return errors.New("the password must have at least one uppercase letter")
	}
	if !regexp.MustCompile(`[0-9]`).MatchString(password) {
		return errors.New("the password must have at least one digit")
	}

	if !regexp.MustCompile(`[!@#\$%\^&\*\(\)_\+\-=\[\]\{\}\\|;:'", <>\./\?]`).MatchString(password) {
		return errors.New("the password must have at least one special character")
	}
	return nil
}

func IsValidEmail(email string) bool {
	var re = regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	if !re.MatchString(email) {
		return false
	}
	domain := strings.Split(email, "@")[1]
	if !strings.Contains(domain, ".") {
		return false
	}
	_, err := net.LookupMX(domain)
	return err == nil
}

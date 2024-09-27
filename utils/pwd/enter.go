package pwd

import (
	"golang.org/x/crypto/bcrypt"
	"todoBackend/utils/loggers"
)

// Hash 密码
func HashPassword(pwd string) string {
	hash, err := bcrypt.GenerateFromPassword([]byte(pwd), bcrypt.MinCost)
	if err != nil {
		loggers.TodoLogger.Error(err)
	}
	return string(hash)
}
func CheckPasswordHash(pwd string, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(pwd))
	if err != nil {
		loggers.TodoLogger.Error(err)
		return false
	}
	return true
}

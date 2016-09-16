package managers

import (
	"os"

	"github.com/Gwennin/IntelligentNetwork_Go/src/errors"
	"github.com/Gwennin/IntelligentNetwork_Go/src/models"
)

type Authenticator interface {
	Authenticate(username string, password string) (bool, *errors.INError)
}

func GetAuthenticator() Authenticator {
	authType := os.Getenv("AUTH_TYPE")

	switch authType {
	case "DB":
		return newDbAuthenticator()
	default:
		return nil
	}
}

type dbAuthenticator struct {
}

func newDbAuthenticator() Authenticator {
	return dbAuthenticator{}
}

func (dbAuthenticator) Authenticate(username string, password string) (bool, *errors.INError) {
	return models.PasswordMatch(username, password)
}

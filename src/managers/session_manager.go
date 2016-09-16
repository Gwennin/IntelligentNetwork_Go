package managers

import (
	"github.com/Gwennin/IntelligentNetwork_Go/src/models"
	"github.com/nu7hatch/gouuid"
	"log"
	"os"
	"strconv"
	"sync"
	"time"
)

var _sessions map[string]*models.Session = make(map[string]*models.Session)
var mutex = &sync.Mutex{}

func OpenSession(username string) string {
	token, _ := uuid.NewV4()

	session := new(models.Session)
	session.Username = username
	session.OpenedOn = time.Now()
	session.Token = token.String()

	mutex.Lock()
	_sessions[session.Token] = session
	mutex.Unlock()

	return session.Token
}

func GetSessionUser(token string) *string {
	mutex.Lock()
	session, ok := _sessions[token]
	mutex.Unlock()

	if ok {
		return &session.Username
	}
	return nil
}

func GetSessionOpenedDate(token string) *time.Time {
	mutex.Lock()
	session, ok := _sessions[token]
	mutex.Unlock()

	if ok {
		return &session.OpenedOn
	}
	return nil
}

func CloseSession(token string) {
	mutex.Lock()
	delete(_sessions, token)
	mutex.Unlock()
}

func CleanSessions() {
	sessionCleanUp := os.Getenv("SESSION_CLEANUP")
	cleanUpSeconds, _ := strconv.Atoi(sessionCleanUp)
	ticker := time.NewTicker(time.Duration(cleanUpSeconds) * time.Second)
	go func() {
		for t := range ticker.C {
			log.Println("Session cleaning at", t)
			mutex.Lock()

			var toRemove []string

			for token, session := range _sessions {
				sessionTimeout := os.Getenv("SESSION_TIMEOUT")
				seconds, _ := strconv.Atoi(sessionTimeout)
				expireDate := session.OpenedOn.Add(time.Duration(seconds) * time.Second)

				if !time.Now().Before(expireDate) {
					toRemove = append(toRemove, token)
				}
			}

			for _, token := range toRemove {
				log.Println("Session", token, "cleaned")
				delete(_sessions, token)
			}

			mutex.Unlock()
		}
	}()
}

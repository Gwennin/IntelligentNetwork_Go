package models

import "time"

type Session struct {
	Username string
	OpenedOn time.Time
	Token    string
}

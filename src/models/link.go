package models

import (
	"github.com/Gwennin/IntelligentNetwork_Go/src/errors"
	"github.com/Gwennin/IntelligentNetwork_Go/src/managers/db"
	"time"
)

type Link struct {
	Id       int       `json:"link_id"`
	Link     string    `json:"link"`
	PostedBy string    `json:"posted_by"`
	PostedIn string    `json:"posted_in"`
	PostedOn time.Time `json:"posted_on"`
	Readed   bool      `json:"readed" gorm:"-"`
}

type readedLink struct {
	ReadLink int
	Reader   string
	ReadOn   time.Time
}

func ListLinks(space string, username string) ([]Link, *errors.INError) {
	database, mutex := db.GetDB()

	if database != nil {
		var links []Link
		database.Table("links").Select("links.*, (r.read_id IS NOT NULL) AS readed").
			Joins("LEFT JOIN readed_links r ON links.id = r.read_link AND r.reader = ?", username).
			Where("posted_in = ?", space).Find(&links)

		mutex.Unlock()

		return links, nil
	}

	mutex.Unlock()

	err := errors.FatalError(1, "Unable to access to the database. May be the connection is closed.")
	return []Link{}, err
}

func AddLink(link *Link) (*Link, *errors.INError) {
	database, mutex := db.GetDB()

	if database != nil {
		database.Create(link)

		mutex.Unlock()
		return link, nil
	}

	mutex.Unlock()

	err := errors.FatalError(1, "Unable to access to the database. May be the connection is closed.")
	return nil, err
}

func DeleteLink(id int, name string) *errors.INError {
	database, mutex := db.GetDB()
	var err *errors.INError = nil

	if database != nil {
		database.Where("id = ? AND posted_in = ?", id, name).Delete(&Link{})
	} else {
		err = errors.FatalError(1, "Unable to access to the database. May be the connection is closed.")
	}

	mutex.Unlock()

	return err
}

func SetLinkRead(id int, by string) *errors.INError {
	database, mutex := db.GetDB()
	var err *errors.INError = nil

	if database != nil {
		readed := new(readedLink)
		readed.ReadLink = id
		readed.ReadOn = time.Now()
		readed.Reader = by

		database.Create(readed)
	} else {
		err = errors.FatalError(1, "Unable to access to the database. May be the connection is closed.")
	}

	mutex.Unlock()

	return err
}

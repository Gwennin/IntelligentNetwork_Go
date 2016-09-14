package models

import (
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

func ListLinks(space string, username string) []Link {
	database, mutex := db.GetDB()

	if database != nil {
		var links []Link
		database.Table("links").Select("links.*, (r.read_id IS NOT NULL) AS readed").
			Joins("LEFT JOIN readed_links r ON links.id = r.read_link AND r.reader = ?", username).
			Where("posted_in = ?", space).Find(&links)

		mutex.Unlock()

		return links
	}

	mutex.Unlock()
	return []Link{}
}

func AddLink(link *Link) *Link {
	database, mutex := db.GetDB()

	if database != nil {
		database.Create(link)

		mutex.Unlock()
		return link
	}

	mutex.Unlock()
	return nil
}

func DeleteLink(id int, name string) {
	database, mutex := db.GetDB()

	if database != nil {
		database.Where("id = ? AND posted_in = ?", id, name).Delete(&Link{})
	}

	mutex.Unlock()
}

func SetLinkRead(id int, by string) {
	database, mutex := db.GetDB()

	if database != nil {
		readed := new(readedLink)
		readed.ReadLink = id
		readed.ReadOn = time.Now()
		readed.Reader = by

		database.Create(readed)
	}

	mutex.Unlock()
}

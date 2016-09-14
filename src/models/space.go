package models

import (
	"github.com/Gwennin/IntelligentNetwork_Go/src/managers/db"
)

type Space struct {
	Name  string `json:"name"`
	Owner string `json:"owner"`
}

type UserSpace struct {
	UserId  string `json:"user_id"`
	SpaceId string `json:"space_id"`
}

type newSpace struct {
	Name   string
	Owner  string
	Public bool
}

func (newSpace) TableName() string {
	return "spaces"
}

func ListPublicSpaces() []Space {
	database, mutex := db.GetDB()

	if database != nil {
		var spaces []Space
		database.Model(&Space{}).Where("public = TRUE").Find(&spaces)

		mutex.Unlock()

		return spaces
	}

	mutex.Unlock()
	return []Space{}
}

func AddSpace(space *Space) {
	addSpace(space, true)
}

func AddPrivateSpace(space *Space) {
	addSpace(space, false)
}

func addSpace(space *Space, public bool) {
	database, mutex := db.GetDB()

	if database != nil {
		creatingSpace := new(newSpace)
		creatingSpace.Name = space.Name
		creatingSpace.Owner = space.Owner
		creatingSpace.Public = public

		database.Create(creatingSpace)
	}

	mutex.Unlock()
}

func DeleteSpace(name string) {
	database, mutex := db.GetDB()

	if database != nil {
		database.Where("name = ?", name).Delete(&Space{})
	}

	mutex.Unlock()
}

package models

import (
	"github.com/Gwennin/IntelligentNetwork_Go/src/errors"
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

func ListPublicSpaces() ([]Space, *errors.INError) {
	database, mutex := db.GetDB()

	if database != nil {
		var spaces []Space
		err := database.Model(&Space{}).Where("public = TRUE").Find(&spaces).Error

		mutex.Unlock()

		if err != nil {
			return []Space{}, errors.NewError(5, "An error occured while accessing to the database.")
		}

		return spaces, nil
	}

	mutex.Unlock()

	err := errors.FatalError(1, "Unable to access to the database. May be the connection is closed.")
	return []Space{}, err
}

func AddSpace(space *Space) *errors.INError {
	return addSpace(space, true)
}

func AddPrivateSpace(space *Space) *errors.INError {
	return addSpace(space, false)
}

func addSpace(space *Space, public bool) *errors.INError {
	database, mutex := db.GetDB()
	var err *errors.INError = nil

	if database != nil {
		creatingSpace := new(newSpace)
		creatingSpace.Name = space.Name
		creatingSpace.Owner = space.Owner
		creatingSpace.Public = public

		err = database.Create(creatingSpace).Error
	} else {
		err = errors.FatalError(1, "Unable to access to the database. May be the connection is closed.")
	}

	mutex.Unlock()

	return err
}

func DeleteSpace(name string) *errors.INError {
	database, mutex := db.GetDB()
	var err *errors.INError = nil

	if database != nil {
		err = database.Where("name = ?", name).Delete(&Space{}).Error
	} else {
		err = errors.FatalError(1, "Unable to access to the database. May be the connection is closed.")
	}

	mutex.Unlock()

	return err
}

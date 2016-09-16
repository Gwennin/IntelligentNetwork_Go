package models

import (
	"github.com/Gwennin/IntelligentNetwork_Go/src/errors"
	"github.com/Gwennin/IntelligentNetwork_Go/src/managers/db"
)

type User struct {
	Username string `json:"username"`
	Password string `json:"-"`
}

type NewUser struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (NewUser) TableName() string {
	return "users"
}

func IsUserExist(alias string) (bool, *errors.INError) {
	database, mutex := db.GetDB()

	if database != nil {
		var count int
		database.Model(&User{}).Where("username = ?", alias).Count(&count)
		mutex.Unlock()

		return count == 1, nil
	}

	mutex.Unlock()

	err := errors.FatalError(1, "Unable to access to the database. May be the connection is closed.")
	return false, err
}

func PasswordMatch(alias string, password string) (bool, *errors.INError) {
	database, mutex := db.GetDB()

	if database != nil {
		var count int
		database.Model(&User{}).Where("username = ?", alias).Where("password = ?", password).Count(&count)
		mutex.Unlock()

		return count == 1, nil
	}

	mutex.Unlock()

	err := errors.FatalError(1, "Unable to access to the database. May be the connection is closed.")
	return false, err
}

func ListUsers() ([]User, *errors.INError) {
	database, mutex := db.GetDB()

	if database != nil {
		var users []User
		database.Model(&User{}).Find(&users)

		mutex.Unlock()

		return users, nil
	}

	mutex.Unlock()

	err := errors.FatalError(1, "Unable to access to the database. May be the connection is closed.")
	return []User{}, err
}

func AddUser(user *NewUser) (string, *errors.INError) {
	database, mutex := db.GetDB()

	if database != nil {
		database.Create(user)

		mutex.Unlock()
		return user.Username, nil
	}

	mutex.Unlock()

	err := errors.FatalError(1, "Unable to access to the database. May be the connection is closed.")
	return "", err
}

func DeleteUser(alias string) *errors.INError {
	database, mutex := db.GetDB()
	var err *errors.INError = nil

	if database != nil {
		database.Where("username = ?", alias).Delete(&User{})
	} else {
		err = errors.FatalError(1, "Unable to access to the database. May be the connection is closed.")
	}

	mutex.Unlock()

	return err
}

func ChangePassword(alias string, password string) *errors.INError {
	database, mutex := db.GetDB()
	var err *errors.INError = nil

	if database != nil {
		database.Model(&User{}).Where("username = ?", alias).Update("password", password)
	} else {
		err = errors.FatalError(1, "Unable to access to the database. May be the connection is closed.")
	}

	mutex.Unlock()

	return err
}

func AddUserSpace(userSpace *UserSpace) *errors.INError {
	database, mutex := db.GetDB()
	var err *errors.INError = nil

	if database != nil {
		database.Create(userSpace)
	} else {
		err = errors.FatalError(1, "Unable to access to the database. May be the connection is closed.")
	}

	mutex.Unlock()

	return err
}

func DeleteUserSpace(userSpace *UserSpace) *errors.INError {
	database, mutex := db.GetDB()
	var err *errors.INError = nil

	if database != nil {
		database.Where("user_id = ? AND space_id = ?", userSpace.UserId, userSpace.SpaceId).Delete(&UserSpace{})
	} else {
		err = errors.FatalError(1, "Unable to access to the database. May be the connection is closed.")
	}

	mutex.Unlock()

	return err
}

func ListUserSpaces(alias string) ([]string, *errors.INError) {
	database, mutex := db.GetDB()

	if database != nil {
		var spaces []string
		database.Table("user_spaces").Where("user_id = ?", alias).Pluck("space_id", &spaces)

		mutex.Unlock()

		return spaces, nil
	}

	mutex.Unlock()

	err := errors.FatalError(1, "Unable to access to the database. May be the connection is closed.")
	return []string{}, err
}

func ListOwnedSpaces(alias string) ([]string, *errors.INError) {
	database, mutex := db.GetDB()

	if database != nil {
		var spaces []string
		database.Table("spaces").Where("owner = ?", alias).Pluck("name", &spaces)

		mutex.Unlock()

		return spaces, nil
	}

	mutex.Unlock()

	err := errors.FatalError(1, "Unable to access to the database. May be the connection is closed.")
	return []string{}, err
}

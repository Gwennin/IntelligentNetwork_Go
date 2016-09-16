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
		err := database.Model(&User{}).Where("username = ?", alias).Count(&count).Error
		mutex.Unlock()

		if err != nil {
			return false, errors.NewError(5, "An error occured while accessing to the database.")
		}

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
		err := database.Model(&User{}).Where("username = ?", alias).Where("password = ?", password).Count(&count).Error
		mutex.Unlock()

		if err != nil {
			return false, errors.NewError(5, "An error occured while accessing to the database.")
		}

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
		err := database.Model(&User{}).Find(&users).Error

		mutex.Unlock()

		if err != nil {
			return []User{}, errors.NewError(5, "An error occured while accessing to the database.")
		}

		return users, nil
	}

	mutex.Unlock()

	err := errors.FatalError(1, "Unable to access to the database. May be the connection is closed.")
	return []User{}, err
}

func AddUser(user *NewUser) (string, *errors.INError) {
	database, mutex := db.GetDB()

	if database != nil {
		err := database.Create(user).Error

		mutex.Unlock()

		if err != nil {
			return "", errors.NewError(5, "An error occured while accessing to the database.")
		}

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
		err = database.Where("username = ?", alias).Delete(&User{}).Error
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
		err = database.Model(&User{}).Where("username = ?", alias).Update("password", password)
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
		err = database.Create(userSpace).Error
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
		err = database.Where("user_id = ? AND space_id = ?", userSpace.UserId, userSpace.SpaceId).Delete(&UserSpace{}).Error
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
		err := database.Table("user_spaces").Where("user_id = ?", alias).Pluck("space_id", &spaces).Error

		mutex.Unlock()

		if err != nil {
			return []string{}, errors.NewError(5, "An error occured while accessing to the database.")
		}

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
		err := database.Table("spaces").Where("owner = ?", alias).Pluck("name", &spaces).Error

		mutex.Unlock()

		if err != nil {
			return []string{}, errors.NewError(5, "An error occured while accessing to the database.")
		}

		return spaces, nil
	}

	mutex.Unlock()

	err := errors.FatalError(1, "Unable to access to the database. May be the connection is closed.")
	return []string{}, err
}

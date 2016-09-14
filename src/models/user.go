package models

import (
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

func IsUserExist(alias string) bool {
	database, mutex := db.GetDB()

	if database != nil {
		var count int
		database.Model(&User{}).Where("username = ?", alias).Count(&count)
		mutex.Unlock()

		return count == 1
	}

	mutex.Unlock()
	return false
}

func PasswordMatch(alias string, password string) bool {
	database, mutex := db.GetDB()

	if database != nil {
		var count int
		database.Model(&User{}).Where("username = ?", alias).Where("password = ?", password).Count(&count)
		mutex.Unlock()

		return count == 1
	}

	mutex.Unlock()
	return false
}

func ListUsers() []User {
	database, mutex := db.GetDB()

	if database != nil {
		var users []User
		database.Model(&User{}).Find(&users)

		mutex.Unlock()

		return users
	}

	mutex.Unlock()
	return []User{}
}

func AddUser(user *NewUser) string {
	database, mutex := db.GetDB()

	if database != nil {
		database.Create(user)

		mutex.Unlock()
		return user.Username
	}

	mutex.Unlock()
	return ""
}

func DeleteUser(alias string) {
	database, mutex := db.GetDB()

	if database != nil {
		database.Where("username = ?", alias).Delete(&User{})
	}

	mutex.Unlock()
}

func ChangePassword(alias string, password string) {
	database, mutex := db.GetDB()

	if database != nil {
		database.Model(&User{}).Where("username = ?", alias).Update("password", password)
	}

	mutex.Unlock()
}

func AddUserSpace(userSpace *UserSpace) {
	database, mutex := db.GetDB()

	if database != nil {
		database.Create(userSpace)
	}

	mutex.Unlock()
}

func DeleteUserSpace(userSpace *UserSpace) {
	database, mutex := db.GetDB()

	if database != nil {
		database.Where("user_id = ? AND space_id = ?", userSpace.UserId, userSpace.SpaceId).Delete(&UserSpace{})
	}

	mutex.Unlock()
}

func ListUserSpaces(alias string) []string {
	database, mutex := db.GetDB()

	if database != nil {
		var spaces []string
		database.Table("user_spaces").Where("user_id = ?", alias).Pluck("space_id", &spaces)

		mutex.Unlock()

		return spaces
	}

	mutex.Unlock()
	return []string{}
}

func ListOwnedSpaces(alias string) []string {
	database, mutex := db.GetDB()

	if database != nil {
		var spaces []string
		database.Table("spaces").Where("owner = ?", alias).Pluck("name", &spaces)

		mutex.Unlock()

		return spaces
	}

	mutex.Unlock()
	return []string{}
}

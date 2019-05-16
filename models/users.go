package models

import (
	u "../utils"
	"fmt"
	"github.com/jinzhu/gorm"
	"strings"
)

type User struct {
	gorm.Model
	FirstName string `json:"first_name"`
	LastName string `json:"last_name"`
	DOB string `json:"dob"`
	Email string `json:"email"`
}

func (user *User) Validate() (map[string]interface{}, bool)  {
	if !strings.Contains(user.Email,"@"){
		return u.Message(false, "Email address is required"), false
	}
	if user.Email == ""{
		return u.Message(false, "Email address should not be empty"), false
	}
	if user.FirstName == "" {
		return u.Message(false, "User first name should be on the payload"), false
	}
	if user.DOB == "" {
		return u.Message(false, "User date of birth should be on the payload"), false
	}

	return u.Message(false, "Requirement passed"), true
}

func (user *User)Create()(map[string]interface{})  {

	if resp, ok := user.Validate(); !ok {
		return resp
	}

	GetDB().Create(user)

	resp := u.Message(true, "success")
	resp["user"] = user
	return resp
}

func GetRegisteredUser (id uint)(*User){

	user := &User{}
	err := GetDB().Table("users").Where("id = ?", id).First(user).Error
	if err != nil {
		return nil
	}
	return user
}
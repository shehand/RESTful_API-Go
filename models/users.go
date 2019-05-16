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
		return u.Message(false, "Contact first name should be on the payload"), false
	}
	
	return u.Message(false, "Requirement passed"), true
}
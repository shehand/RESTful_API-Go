package models

import (
	u "../utils"
	"github.com/jinzhu/gorm"
	"fmt"
)

type Contact struct {
	gorm.Model
	Name string `json:"name"`
	Phone string `json:"phone"`
	UserId uint `json:"user_id"` //The user that this contact belongs to
}

func (contact *Contact) Validate() (map[string] interface{}, bool) {

	if contact.Name == "" {
		return u.Message(false, "Contact name should be on the payload"), false
	}

	if contact.Phone == "" {
		return u.Message(false, "Phone number should be on the payload"), false
	}

	if contact.UserId <= 0 {
		return u.Message(false, "User is not recognized"), false
	}

	//All the required parameters are present
	return u.Message(true, "success"), true
}

func (contact *Contact) Create() (map[string] interface{}) {

	if resp, ok := contact.Validate(); !ok {
		return resp
	}

	GetDB().Create(contact)

	resp := u.Message(true, "success")
	resp["contact"] = contact
	return resp
}

func GetContact(id uint) (*Contact) {

	contact := &Contact{}
	err := GetDB().Table("contacts").Where("id = ?", id).First(contact).Error
	if err != nil {
		return nil
	}
	return contact
}

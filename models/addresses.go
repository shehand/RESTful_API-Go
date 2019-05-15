package models

import (
	u "../utils"
	"fmt"
	"github.com/jinzhu/gorm"
)

type Address struct {
	gorm.Model
	Number string `json:"number"`
	FirstLine string `json:"first_line"`
	Town string `json:"town"`
	PostalCode string `json:"postal_code"`
	Province string `json:"province"`
	Country string `json:"country"`
	UserId uint `json:"user_id"` //The user that this contact belongs to
}

func (address *Address) Validate() (map[string] interface{}, bool) {

	if address.PostalCode == "" {
		return u.Message(false, "Postal code should be on the payload"), false
	}

	if address.Country == "" {
		return u.Message(false, "Country should be on the payload"), false
	}

	if address.UserId <= 0 {
		return u.Message(false, "User is not recognized"), false
	}

	//All the required parameters are present
	return u.Message(true, "success"), true
}
func (address *Address) Create()(map[string]interface{}){
	if resp, ok := address.Validate(); !ok {
		return resp
	}

	GetDB().Create(address)

	resp := u.Message(true, "success")
	resp["address"] = address
	return resp
}

func GetAddress(id uint)(*Address){
	address := &Address{}
	err := GetDB().Table("addresses").Where("id = ?", id).First(address).Error
	if err != nil {
		return nil
	}
	return address
}

func GetAddress(id uint)(*Address){
	address := &Address{}
	err := GetDB().Table("addresses").Where("id = ?", id).First(address).Error
	if err != nil {
		return nil
	}
	return address
}
package models

import (
	"fmt"
	"github.com/jinzhu/gorm"
	u "../utils"

)

type User struct {
	FirstName string `json:"first_name"`
	LastName string `json:"last_name"`
	DOB string `json:"dob"`
	Email string `json:"email"`
} 
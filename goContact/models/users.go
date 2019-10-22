package models

import (
	"fmt"
	u "../utils"
)

type User struct {
	Id      uint  `json:"id"`
	Name    string `json:"name,omitempty"`
	SurName string `json:"surname,omitempty"`
	Age     int   `json:"age,omitempty"`
}

func (user *User) Validate() (map[string]interface{}, bool) {

	if user.Name == "" {
		return u.Message(false, "Contact name should be on the payload"), false
	}

	if user.SurName == "" {
		return u.Message(false, "Phone number should be on the payload"), false
	}

	if user.Id <= 0 {
		return u.Message(false, "User is not recognized"), false
	}

	if user.Age <= 0 {
		return u.Message(false, "User is not recognized"), false
	}

	//All the required parameters are present
	return u.Message(true, "success"), true
}

func (user *User) Create() (map[string]interface{}) {
	if resp, ok := user.Validate(); !ok {
		return resp
	}
	//GetDB().Create(user)
	resp := u.Message(true, "success" )
	resp["user"] = user
	return resp
}

func GetUser(id uint) (*User) {
	s := User{Name: "Sean", SurName: "Smith", Id: id, Age: 50}
	user := &s
	/*err := GetDB().Table("user").Where("id = ?", id).First(user).Error
	if err != nil {
		return nil
	}*/
	fmt.Print("Obteniendo el user")
	return user
}

func GetUsers(user uint) ([]*User) {

	users := make([]*User, 0)
	/*err := GetDB().Table("user").Where("user_id = ?", user).Find(&user).Error
	if err != nil {
		fmt.Println(err)
		return nil
	}*/
	fmt.Print("Obteniendo lista de users")
	return users
}
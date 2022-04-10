package db

import (
	"github.com/jinzhu/gorm"
	"simpleApi/seriallzer"
	"strings"
)

var NotFound = gorm.ErrRecordNotFound

func InsertUser(user *User) (bool, error) {
	res := engine.Create(user)
	if res.Error != nil {
		return false, res.Error
	}

	return true, nil
}

func GetUsers(query map[string]interface{}) (error, []*User) {
	var users []*User
	res := engine.Where(query).Find(&users)

	return res.Error, users
}

func GetUserByID(id int) (*User, error) {
	user := &User{}
	res := engine.Where(map[string]interface{}{"id": id}).First(user)

	return user, res.Error
}


func InsertRelation(relation *Relation) (bool, error) {
	//gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
	//	return defaultTableName + "_01"
	//}
	//gorm.DefaultTableNameHandler(engine, "relation")

	res := engine.Create(relation)
	if res.Error != nil {
		return false, res.Error
	}

	return true, nil
}

func UpdateUser(id int, user *seriallzer.UserUpdateReq) error {
	m := make(map[string]interface{})

	if user.Weight > 0 {
		m["weight"] = user.Weight
	}
	if user.Stature > 0 {
		m["stature"] = user.Stature
	}
	address := strings.Trim(user.Address, " ")
	if address != "" {
		m["address"] = address
	}
	occupation := strings.Trim(user.Occupation, " ")
	if occupation != "" {
		m["occupation"] = occupation
	}

	res := engine.Model(&User{}).Where(map[string]interface{}{"id": id}).Update(m)
	return res.Error
}

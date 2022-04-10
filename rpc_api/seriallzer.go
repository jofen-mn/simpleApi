package rpc_api

import (
	"simpleApi/db"
	simple "simpleApi/protos"
)

func resetUserResponse(user *db.User, res *simple.UserResponse) {
	res.Name = user.Name
	res.Number = user.Number
	res.Phone = user.Phone
	res.Age = int32(user.Age)
	res.Gender = int32(user.Gender)
	res.Weight = int32(user.Weight)
	res.Stature = int32(user.Stature)
	res.Address = user.Address
	res.Occupation = user.Occupation
}

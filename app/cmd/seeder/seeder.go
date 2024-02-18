package main

import (
	"github.com/PunkPlusPlus/cources_service/app/internal/storage"
	"github.com/PunkPlusPlus/cources_service/app/internal/users"
)

func main() {
	var s = storage.GetStorage()
	err := s.DB.AutoMigrate(&users.DbUser{})
	if err != nil {
		panic(err)
	}
	var email = "lilimomlnt@gmail.com"
	var user = users.DbUser{
		Username: "IdDev",
		Email:    &email,
		Password: "root",
	}
	var result = s.DB.Create(&user)
	if result.Error != nil {
		panic(result.Error)
	}
}

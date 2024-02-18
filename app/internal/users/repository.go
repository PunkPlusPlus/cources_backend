package users

import (
	"github.com/PunkPlusPlus/cources_service/app/internal/storage"
)

func DoLogin(dto *Login) *DbUser {
	var s = storage.GetStorage()
	var user = DbUser{}
	var result = s.DB.Where(&dto).First(&user)
	//var result = s.DB.First(&user)
	if result.Error != nil {
		return nil
	}
	if result.RowsAffected > 0 {
		return &user
	}
	return nil
}

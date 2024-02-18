package auth

import (
	"github.com/PunkPlusPlus/cources_service/app/internal/storage"
	"github.com/PunkPlusPlus/cources_service/app/internal/users"
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
)

func payload(data any) jwt.MapClaims {
	if v, ok := data.(*users.DbUser); ok {
		return jwt.MapClaims{
			"username": v.Username,
		}
	}
	return jwt.MapClaims{}
}

func identity(c *gin.Context) interface{} {
	claims := jwt.ExtractClaims(c)
	return &users.DbUser{
		Username: claims[identityKey].(string),
	}
}

func authenticator(c *gin.Context) (any, error) {
	var loginVals users.Login
	if err := c.ShouldBind(&loginVals); err != nil {
		return "", jwt.ErrMissingLoginValues
	}
	var user = users.DoLogin(&loginVals)
	if user == nil {
		return nil, jwt.ErrFailedAuthentication
	}
	return user, nil

	//return nil, jwt.ErrFailedAuthentication
}

func authorizator(data any, c *gin.Context) bool {
	user, ok := data.(*users.DbUser)
	if !ok {
		return false
	}
	var s = storage.GetStorage()
	var result = s.DB.Where("username=?", user.Username).First(&user)
	if result.Error != nil || result.RowsAffected == 0 {
		return false
	}
	return true

}

func unauthorized(c *gin.Context, code int, message string) {
	c.JSON(code, gin.H{
		"code":    code,
		"message": message,
	})
}

package controllers

import (
	"net/http"

	"huckleberry.app/server/auth"
	"huckleberry.app/server/database"
	"huckleberry.app/server/models"
	"huckleberry.app/server/utils/formaterror"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func Login(c *gin.Context) {
	user := models.User{}
	if err := c.ShouldBind(&user); err != nil {
		formattedError := formaterror.FormatError(http.StatusBadRequest)
		c.JSON(http.StatusBadRequest, formattedError)
		return
	}

	user, token, err := signIn(user.Username, user.Password)
	if err != nil {
		formattedError := formaterror.FormatError(http.StatusUnprocessableEntity)
		c.JSON(http.StatusUnprocessableEntity, formattedError)
		return
	}

	userDTO := user.ToDTO()
	userDTO.Token = token

	c.JSON(http.StatusOK, userDTO)
}

func signIn(username, password string) (models.User, string, error) {
	var err error
	user := models.User{}
	err = database.DB.Debug().Where("username = ?", username).First(&user).Error
	if err != nil {
		return user, "", err
	}
	err = models.VerifyPassword(user.Password, password)
	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
		return user, "", err
	}
	token, err := auth.CreateToken(user.ID)

	return user, token, err
}

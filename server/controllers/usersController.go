package controllers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"huckleberry.app/server/dtos"
	"huckleberry.app/server/models"
	"huckleberry.app/server/utils/formaterror"
)

func UsernameExists(c *gin.Context) {
	username := c.Param("username")

	usernameExists, err := models.UsernameExists(username)
	if gorm.IsRecordNotFoundError(err) {
		formattedError := formaterror.FormatError(http.StatusNotFound)
		c.JSON(http.StatusNotFound, formattedError)
		return
	}

	if err != nil {
		formattedError := formaterror.FormatError(http.StatusBadRequest)
		c.JSON(http.StatusBadRequest, formattedError)
		return
	}

	resourceExistsDTO := dtos.ResourceExistsDTO{
		usernameExists,
		"That username has been taken. Please choose another.",
	}

	c.JSON(http.StatusOK, resourceExistsDTO)
}

func EmailExists(c *gin.Context) {
	email := c.Param("email")

	emailExists, err := models.EmailExists(email)
	if gorm.IsRecordNotFoundError(err) {
		formattedError := formaterror.FormatError(http.StatusNotFound)
		c.JSON(http.StatusNotFound, formattedError)
		return
	}

	if err != nil {
		formattedError := formaterror.FormatError(http.StatusBadRequest)
		c.JSON(http.StatusBadRequest, formattedError)
		return
	}

	resourceExistsDTO := dtos.ResourceExistsDTO{
		emailExists,
		"That email has been associated with another account.",
	}

	c.JSON(http.StatusOK, resourceExistsDTO)
}

// CreateUser registers a new user
func CreateUser(c *gin.Context) {
	user := models.User{}
	if err := c.ShouldBind(&user); err != nil {
		formattedError := formaterror.FormatError(http.StatusBadRequest)
		c.JSON(http.StatusBadRequest, formattedError)
		return
	}

	userCreated, err := user.Create()
	if err != nil {
		formattedError := formaterror.FormatError(http.StatusBadRequest)
		c.JSON(http.StatusBadRequest, formattedError)
		return
	}

	c.JSON(http.StatusOK, userCreated)
}

func FindNotificationsByUsername(c *gin.Context) {
	username := c.Param("username")
	user := models.User{Username: username}
	user.FindByUsername()

	if user.ID == 0 {
		formattedError := formaterror.FormatError(http.StatusNotFound)
		c.JSON(http.StatusNotFound, formattedError)
		return
	}

	notificationsDTO, err := models.FindNotificationsByUser(&user)
	if err != nil && !gorm.IsRecordNotFoundError(err) {
		formattedError := formaterror.FormatError(http.StatusBadRequest)
		c.JSON(http.StatusBadRequest, formattedError)
		return
	}
	c.JSON(http.StatusOK, notificationsDTO)
}

func NotificationResponse(c *gin.Context) {
	// TODO: Know if user is owner of such notification
	username := c.Param("username")
	user := models.User{Username: username}
	user.FindByUsername()

	if user.ID == 0 {
		formattedError := formaterror.FormatError(http.StatusNotFound)
		c.JSON(http.StatusNotFound, formattedError)
		return
	}

	ID := c.Param("id")
	IDuint64, err := strconv.ParseUint(ID, 10, 64)
	if err != nil {
		formattedError := formaterror.FormatError(http.StatusBadRequest)
		c.JSON(http.StatusBadRequest, formattedError)
		return
	}

	notificationResponse := dtos.NotificationResponseDTO{}
	if err := c.ShouldBind(&notificationResponse); err != nil {
		formattedError := formaterror.FormatError(http.StatusBadRequest)
		c.JSON(http.StatusBadRequest, formattedError)
		return
	}

	notification := models.Notification{}
	notification.ID = IDuint64

	fmt.Println(notificationResponse.BookmarksShare.IsAccepted)
	if notificationResponse.BookmarksShare.IsAccepted {
		// if is accepted: patch share & delete notification return notification deleted
		err = notification.AcceptBookmarksShare()
	} else {
		// if not accepted: delete notification & delete share return notification deleted
		err = notification.RejectBookmarksShare()
	}

	if err != nil {
		formattedError := formaterror.FormatError(http.StatusBadRequest)
		c.JSON(http.StatusBadRequest, formattedError)
		return
	}

	c.JSON(http.StatusOK, notification.ToDTO())

}

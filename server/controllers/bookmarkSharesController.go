package controllers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"huckleberry.app/server/dtos"
	"huckleberry.app/server/models"
	"huckleberry.app/server/utils/formaterror"
)

func CreateShare(c *gin.Context) {
	username := c.Param("username")
	user := models.User{Username: username}

	user.FindByUsername()
	if user.ID == 0 {
		formattedError := formaterror.FormatError(http.StatusBadRequest)
		c.JSON(http.StatusBadRequest, formattedError)
		return
	}

	bookmarksShare := models.BookmarksShare{}
	if err := c.ShouldBind(&bookmarksShare); err != nil {
		formattedError := formaterror.FormatError(http.StatusBadRequest)
		c.JSON(http.StatusBadRequest, formattedError)
		return
	}

	if username == bookmarksShare.Receiver.Username {
		c.JSON(http.StatusForbidden, gin.H{"error": "Can't share your bookmarks with yourself"})
		return
	}

	receiver := models.User{Username: bookmarksShare.Receiver.Username}
	receiver.FindByUsername()
	if receiver.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Username not found"})
		return
	}

	bookmarksShare.User = user
	bookmarksShare.Receiver = receiver
	bookmarksShare.UserID = user.ID
	bookmarksShare.ReceiverID = receiver.ID

	isAlreadyInDB := bookmarksShare.FindByUserAndReceiver()
	if isAlreadyInDB {
		c.JSON(http.StatusConflict, gin.H{"error": "You have already shared your bookmarks with this user."})
		return
	}

	bookmarksShareCreated, err := bookmarksShare.Create()

	if err != nil {
		fmt.Println(err)
		formattedError := formaterror.FormatError(http.StatusInternalServerError)
		c.JSON(http.StatusInternalServerError, formattedError)
		return
	}
	c.JSON(http.StatusOK, bookmarksShareCreated.ToDTO())
}

func FindSharesByUsername(c *gin.Context) {
	username := c.Param("username")
	user := models.User{Username: username}
	user.FindByUsername()
	if user.ID == 0 {
		formattedError := formaterror.FormatError(http.StatusNotFound)
		c.JSON(http.StatusNotFound, formattedError)
		return
	}
	var bookmarksShareDTO []dtos.BookmarksShareDTO
	bookmarksShareDTO = models.FindSharesFromUser(user)

	c.JSON(http.StatusOK, bookmarksShareDTO)
}

func DeleteShare(c *gin.Context) {
	// TODO: Know if user is owner of such share
	// username := c.Param("username")
	ID := c.Param("id")
	bookmarksShare := models.BookmarksShare{}
	IDuint64, err := strconv.ParseUint(ID, 10, 64)
	if err != nil {
		formattedError := formaterror.FormatError(http.StatusBadRequest)
		c.JSON(http.StatusBadRequest, formattedError)
		return
	}
	bookmarksShare.ID = IDuint64

	deletedBookmarksShare, err := bookmarksShare.Delete()
	if err != nil {
		formattedError := formaterror.FormatError(http.StatusInternalServerError)
		c.JSON(http.StatusInternalServerError, formattedError)
		return
	}
	// Delete notification generated by share, if any
	notification := models.Notification{BookmarksShareID: bookmarksShare.ID}
	notification.DeleteByBookmarksShareID()

	c.JSON(http.StatusOK, deletedBookmarksShare.ToDTO())
}

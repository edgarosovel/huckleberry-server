package controllers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"huckleberry.app/server/dtos"
	"huckleberry.app/server/models"
	"huckleberry.app/server/scraper"
	"huckleberry.app/server/utils/formaterror"
)

func CreateBookmark(c *gin.Context) {
	username := c.Param("username")
	user := models.User{Username: username}
	user.FindByUsername()
	if user.ID == 0 {
		formattedError := formaterror.FormatError(http.StatusBadRequest)
		c.JSON(http.StatusBadRequest, formattedError)
		return
	}

	bookmark := models.Bookmark{}
	if err := c.ShouldBind(&bookmark); err != nil {
		fmt.Println(err)
		formattedError := formaterror.FormatError(http.StatusBadRequest)
		c.JSON(http.StatusBadRequest, formattedError)
		return
	}

	bookmark.UserID = user.ID

	if err := scraper.GetBookmarkInfo(&bookmark); err != nil {
		formattedError := formaterror.FormatError(http.StatusUnprocessableEntity)
		c.JSON(http.StatusUnprocessableEntity, formattedError)
		return
	}

	bookmarkCreated, err := bookmark.Create()
	if err != nil {
		formattedError := formaterror.FormatError(http.StatusBadRequest)
		c.JSON(http.StatusBadRequest, formattedError)
		return
	}

	c.JSON(http.StatusOK, bookmarkCreated)

}

func FindBookmarksByUsername(c *gin.Context) {
	username := c.Param("username")
	user := models.User{Username: username}
	bookmarksResponse := dtos.BookmarksResponseDTO{}

	// Find user's own bookmarks
	err := user.FindOwnBookmarks()
	if err != nil { // username was not found
		formattedError := formaterror.FormatError(http.StatusBadRequest)
		c.JSON(http.StatusBadRequest, formattedError)
		return
	}
	bookmarksResponse.OwnBookmarks = user.ToDTO().Bookmarks

	// Find shared bookmarks from other users
	bookmarksResponse.Shared = models.FindSharesToUser(user)

	c.JSON(http.StatusOK, bookmarksResponse)
}

func DeleteBookmark(c *gin.Context) {
	// TODO: Know if user is owner of such bookmark
	// username := c.Param("username")
	ID := c.Param("id")
	bookmark := models.Bookmark{}
	IDuint64, err := strconv.ParseUint(ID, 10, 64)
	if err != nil {
		formattedError := formaterror.FormatError(http.StatusBadRequest)
		c.JSON(http.StatusBadRequest, formattedError)
		return
	}
	bookmark.ID = IDuint64
	deletedBookmark, err := bookmark.Delete()
	if err != nil {
		formattedError := formaterror.FormatError(http.StatusNotFound)
		c.JSON(http.StatusNotFound, formattedError)
		return
	}
	c.JSON(http.StatusOK, deletedBookmark)
}

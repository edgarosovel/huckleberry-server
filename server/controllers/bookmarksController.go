package controllers

import (
	"fmt"
	"net/http"

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

// func (server *Server) DeletePost(w http.ResponseWriter, r *http.Request) {

// 	vars := mux.Vars(r)

// 	// Is a valid post id given to us?
// 	pid, err := strconv.ParseUint(vars["id"], 10, 64)
// 	if err != nil {
// 		responses.ERROR(w, http.StatusBadRequest, err)
// 		return
// 	}

// 	// Is this user authenticated?
// 	uid, err := auth.ExtractTokenID(r)
// 	if err != nil {
// 		responses.ERROR(w, http.StatusUnauthorized, errors.New("Unauthorized"))
// 		return
// 	}

// 	// Check if the post exist
// 	post := models.Post{}
// 	err = server.DB.Debug().Model(models.Post{}).Where("id = ?", pid).Take(&post).Error
// 	if err != nil {
// 		responses.ERROR(w, http.StatusNotFound, errors.New("Unauthorized"))
// 		return
// 	}

// 	// Is the authenticated user, the owner of this post?
// 	if uid != post.AuthorID {
// 		responses.ERROR(w, http.StatusUnauthorized, errors.New("Unauthorized"))
// 		return
// 	}
// 	_, err = post.DeleteAPost(server.DB, pid, uid)
// 	if err != nil {
// 		responses.ERROR(w, http.StatusBadRequest, err)
// 		return
// 	}
// 	w.Header().Set("Entity", fmt.Sprintf("%d", pid))
// 	responses.JSON(w, http.StatusNoContent, "")
// }

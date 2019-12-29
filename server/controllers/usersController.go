package controllers

import (
	"net/http"

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

// func (server *Server) CreateUser(w http.ResponseWriter, r *http.Request) {

// 	body, err := ioutil.ReadAll(r.Body)
// 	if err != nil {
// 		responses.ERROR(w, http.StatusUnprocessableEntity, err)
// 	}
// 	user := models.User{}
// 	err = json.Unmarshal(body, &user)
// 	if err != nil {
// 		responses.ERROR(w, http.StatusUnprocessableEntity, err)
// 		return
// 	}
// 	user.Prepare()
// 	err = user.Validate("")
// 	if err != nil {
// 		responses.ERROR(w, http.StatusUnprocessableEntity, err)
// 		return
// 	}
// 	userCreated, err := user.SaveUser(server.DB)

// 	if err != nil {

// 		formattedError := formaterror.FormatError(err.Error())

// 		responses.ERROR(w, http.StatusInternalServerError, formattedError)
// 		return
// 	}
// 	w.Header().Set("Location", fmt.Sprintf("%s%s/%d", r.Host, r.RequestURI, userCreated.ID))
// 	responses.JSON(w, http.StatusCreated, userCreated)
// }

// func (server *Server) GetUser(w http.ResponseWriter, r *http.Request) {

// 	vars := mux.Vars(r)
// 	uid, err := strconv.ParseUint(vars["id"], 10, 32)
// 	if err != nil {
// 		responses.ERROR(w, http.StatusBadRequest, err)
// 		return
// 	}
// 	user := models.User{}
// 	userGotten, err := user.FindUserByID(server.DB, uint32(uid))
// 	if err != nil {
// 		responses.ERROR(w, http.StatusBadRequest, err)
// 		return
// 	}
// 	responses.JSON(w, http.StatusOK, userGotten)
// }

// func (server *Server) UpdateUser(w http.ResponseWriter, r *http.Request) {

// 	vars := mux.Vars(r)
// 	uid, err := strconv.ParseUint(vars["id"], 10, 32)
// 	if err != nil {
// 		responses.ERROR(w, http.StatusBadRequest, err)
// 		return
// 	}
// 	body, err := ioutil.ReadAll(r.Body)
// 	if err != nil {
// 		responses.ERROR(w, http.StatusUnprocessableEntity, err)
// 		return
// 	}
// 	user := models.User{}
// 	err = json.Unmarshal(body, &user)
// 	if err != nil {
// 		responses.ERROR(w, http.StatusUnprocessableEntity, err)
// 		return
// 	}
// 	tokenID, err := auth.ExtractTokenID(r)
// 	if err != nil {
// 		responses.ERROR(w, http.StatusUnauthorized, errors.New("Unauthorized"))
// 		return
// 	}
// 	if tokenID != uint32(uid) {
// 		responses.ERROR(w, http.StatusUnauthorized, errors.New(http.StatusText(http.StatusUnauthorized)))
// 		return
// 	}
// 	user.Prepare()
// 	err = user.Validate("update")
// 	if err != nil {
// 		responses.ERROR(w, http.StatusUnprocessableEntity, err)
// 		return
// 	}
// 	updatedUser, err := user.UpdateAUser(server.DB, uint32(uid))
// 	if err != nil {
// 		formattedError := formaterror.FormatError(err.Error())
// 		responses.ERROR(w, http.StatusInternalServerError, formattedError)
// 		return
// 	}
// 	responses.JSON(w, http.StatusOK, updatedUser)
// }

// func (server *Server) DeleteUser(w http.ResponseWriter, r *http.Request) {

// 	vars := mux.Vars(r)

// 	user := models.User{}

// 	uid, err := strconv.ParseUint(vars["id"], 10, 32)
// 	if err != nil {
// 		responses.ERROR(w, http.StatusBadRequest, err)
// 		return
// 	}
// 	tokenID, err := auth.ExtractTokenID(r)
// 	if err != nil {
// 		responses.ERROR(w, http.StatusUnauthorized, errors.New("Unauthorized"))
// 		return
// 	}
// 	if tokenID != 0 && tokenID != uint32(uid) {
// 		responses.ERROR(w, http.StatusUnauthorized, errors.New(http.StatusText(http.StatusUnauthorized)))
// 		return
// 	}
// 	_, err = user.DeleteAUser(server.DB, uint32(uid))
// 	if err != nil {
// 		responses.ERROR(w, http.StatusInternalServerError, err)
// 		return
// 	}
// 	w.Header().Set("Entity", fmt.Sprintf("%d", uid))
// 	responses.JSON(w, http.StatusNoContent, "")
// }

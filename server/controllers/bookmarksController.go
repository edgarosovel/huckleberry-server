package controllers

// func (server *Server) CreateBookmark(w http.ResponseWriter, r *http.Request) {

// 	body, err := ioutil.ReadAll(r.Body)
// 	if err != nil {
// 		responses.ERROR(w, http.StatusUnprocessableEntity, err)
// 		return
// 	}
// 	post := models.Post{}
// 	err = json.Unmarshal(body, &post)
// 	if err != nil {
// 		responses.ERROR(w, http.StatusUnprocessableEntity, err)
// 		return
// 	}
// 	post.Prepare()
// 	err = post.Validate()
// 	if err != nil {
// 		responses.ERROR(w, http.StatusUnprocessableEntity, err)
// 		return
// 	}
// 	uid, err := auth.ExtractTokenID(r)
// 	if err != nil {
// 		responses.ERROR(w, http.StatusUnauthorized, errors.New("Unauthorized"))
// 		return
// 	}
// 	if uid != post.AuthorID {
// 		responses.ERROR(w, http.StatusUnauthorized, errors.New(http.StatusText(http.StatusUnauthorized)))
// 		return
// 	}
// 	postCreated, err := post.SavePost(server.DB)
// 	if err != nil {
// 		formattedError := formaterror.FormatError(err.Error())
// 		responses.ERROR(w, http.StatusInternalServerError, formattedError)
// 		return
// 	}
// 	w.Header().Set("Location", fmt.Sprintf("%s%s/%d", r.Host, r.URL.Path, postCreated.ID))
// 	responses.JSON(w, http.StatusCreated, postCreated)
// }

// func (server *Server) GetBookmarks(w http.ResponseWriter, r *http.Request) {

// 	bookmark := models.Bookmark{}

// 	bookmarks, err := bookmark.FindAllPosts(server.DB)
// 	if err != nil {
// 		responses.ERROR(w, http.StatusInternalServerError, err)
// 		return
// 	}
// 	responses.JSON(w, http.StatusOK, bookmarks)
// }

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

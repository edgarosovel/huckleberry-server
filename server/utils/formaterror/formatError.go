package formaterror

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func FormatError(errCode int) gin.H {

	switch errCode {

	case http.StatusBadRequest:
		return gin.H{"error": "Bad request"}

	case http.StatusUnprocessableEntity:
		return gin.H{"error": "Wrong username or password"}

	case http.StatusNotFound:
		return gin.H{"error": "Resource not found"}

	case http.StatusForbidden:
		return gin.H{"error": "Access to resource forbidden"}

	default:
		return gin.H{"error": "Unknown error. Try again later"}
	}

	return gin.H{"error": "Unknown error"}
}

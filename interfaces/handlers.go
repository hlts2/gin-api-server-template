package interfaces

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hlts2/gin-server-template/application"
	"github.com/hlts2/gin-server-template/domain"
	"github.com/jinzhu/gorm"
)

// ErrIDNotFull is error that id is not full
var ErrIDNotFull = errors.New("id is not full")

// GetUser is handler to get user
func GetUser(ctxt *gin.Context) {
	id := ctxt.GetString("id")
	if len(id) == 0 {
		ctxt.JSON(http.StatusBadRequest, ErrIDNotFull)
		return
	}

	var user domain.User

	var err error
	if user, err = application.GetUser(id); err != nil {
		errorResponse(ctxt, err)
		return
	}

	ctxt.JSON(200, user)
}

// GetUsers is handler to get users
func GetUsers(ctxt *gin.Context) {
	var users domain.Users

	var err error
	if users, err = application.GetUsers(); err != nil {
		errorResponse(ctxt, err)
		return
	}

	ctxt.JSON(200, users)
}

func errorResponse(ctxt *gin.Context, err error) {
	switch err {
	case gorm.ErrRecordNotFound:
		ctxt.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
	default:
		ctxt.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
}

package interfaces

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hlts2/gin-server-template/application"
	"github.com/hlts2/gin-server-template/domain"
)

// GetTweet is handler to get tweet
func GetTweet(ctxt *gin.Context) {
	var tweet domain.Tweet

	id := ctxt.GetString("id")
	if len(id) == 0 {
		ctxt.JSON(http.StatusBadRequest, errors.New("id is not full"))
		return
	}

	var err error
	if tweet, err = application.GetTweet(id); err != nil {
		ctxt.JSON(http.StatusBadRequest, err)
		return
	}

	ctxt.JSON(200, tweet)
}

// GetTweets is handler to get tweets
func GetTweets(ctxt *gin.Context) {
}

package interfaces

import (
	"github.com/gin-gonic/gin"
	"github.com/hlts2/gin-server-template/application"
)

// GetTweet is handler to get tweet
func GetTweet(ctxt *gin.Context) {
	id := ctxt.GetString("id")
	if len(id) == 0 {
		return
	}

	application.GetTweet(id)
}

// GetTweets is handler to get tweets
func GetTweets(ctxt *gin.Context) {
	application.GetTweets()
}

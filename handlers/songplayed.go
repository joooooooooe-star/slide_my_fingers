package handlers

import (
	"errors"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
)

func PostSongPlayed(c *gin.Context) {

	jsonData, err := io.ReadAll(c.Request.Body)
	if err != nil {
		c.Error(errors.New("error reading JSON data"))
		return
	}

	c.String(http.StatusAccepted, string(jsonData))
}

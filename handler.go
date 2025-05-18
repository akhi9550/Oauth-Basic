package main

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/oauth2"
)

func (a *App) loginHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", nil)
}

func (a *App) oAuthHandler(c *gin.Context) {
	url := a.config.AuthCodeURL("hello world", oauth2.AccessTypeOffline)
	c.Redirect(http.StatusTemporaryRedirect, url)
}

func (a *App) oAuthCallbackHandler(c *gin.Context) {
	code := c.Query("code")
	token, err := a.config.Exchange(context.Background(), code)
	if err != nil {
		c.String(http.StatusBadRequest, "Token exchange error: %v", err)
		return
	}

	client := a.config.Client(context.Background(), token)

	resp, err := client.Get("https://www.googleapis.com/oauth2/v2/userinfo")
	if err != nil {
		c.String(http.StatusBadRequest, "Error fetching user info: %v", err)
		return
	}
	defer resp.Body.Close()

	var userInfo map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&userInfo); err != nil {
		c.String(http.StatusInternalServerError, "Error decoding user info: %v", err)
		return
	}

	c.JSON(http.StatusOK, userInfo)
}

package utilities

import (
	"encoding/json"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/mrcunninghamz/tprkpr/platform/authenticator"
)

func GetUser(ctx *gin.Context) (authenticator.UserProfile, error) {
	userProfile := authenticator.UserProfile{}
	session := sessions.Default(ctx)
	profile, error := json.Marshal(session.Get("profile").(map[string]interface{}))

	if error != nil {
		return authenticator.UserProfile{}, error
	}
	error = json.Unmarshal(profile, &userProfile)

	if error != nil {
		return authenticator.UserProfile{}, error
	}

	return userProfile, nil
}

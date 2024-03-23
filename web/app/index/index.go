package index

import (
	"encoding/json"
	"github.com/a-h/templ"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/mrcunninghamz/tprkpr/platform/authenticator"
	"github.com/mrcunninghamz/tprkpr/web/views"
)

func Handler(ctx *gin.Context) {
	session := sessions.Default(ctx)
	profile := session.Get("profile")
	userProfile := authenticator.UserProfile{}

	if profile != nil {
		profileJson, _ := json.Marshal(session.Get("profile").(map[string]interface{}))

		json.Unmarshal(profileJson, &userProfile)
	}

	templ.Handler(views.Index(userProfile)).ServeHTTP(ctx.Writer, ctx.Request)

}

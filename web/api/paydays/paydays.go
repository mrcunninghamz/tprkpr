package paydays

import (
	"encoding/json"
	"github.com/a-h/templ"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/mrcunninghamz/tprkpr/platform/authenticator"
	"github.com/mrcunninghamz/tprkpr/platform/services"
	"github.com/mrcunninghamz/tprkpr/web/views"
)

func Get(paydayService services.Paydays) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		session := sessions.Default(ctx)
		profile, _ := json.Marshal(session.Get("profile").(map[string]interface{}))

		userProfile := authenticator.UserProfile{}
		json.Unmarshal(profile, &userProfile)

		paydays := paydayService.Get(userProfile.Id)

		templ.Handler(views.PayDays(paydays)).ServeHTTP(ctx.Writer, ctx.Request)
	}
}

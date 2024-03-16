package index

import (
	"github.com/a-h/templ"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/mrcunninghamz/tprkpr/web/views"
)

func Handler(ctx *gin.Context) {
	session := sessions.Default(ctx)
	profile := session.Get("profile")

	templ.Handler(views.Index(profile)).ServeHTTP(ctx.Writer, ctx.Request)

}

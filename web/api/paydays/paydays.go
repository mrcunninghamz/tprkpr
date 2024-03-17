package paydays

import (
	"github.com/a-h/templ"
	"github.com/gin-gonic/gin"
	"github.com/mrcunninghamz/tprkpr/platform/data"
	"github.com/mrcunninghamz/tprkpr/web/views"
)

func Get(ctx *gin.Context) {
	templ.Handler(views.PayDays(data.TprKprContext)).ServeHTTP(ctx.Writer, ctx.Request)
}

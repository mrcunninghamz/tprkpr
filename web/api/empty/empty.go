package empty

import (
	"github.com/a-h/templ"
	"github.com/gin-gonic/gin"
	"github.com/mrcunninghamz/tprkpr/web/views"
)

func New(ctx *gin.Context) {
	templ.Handler(views.Empty()).ServeHTTP(ctx.Writer, ctx.Request)
}

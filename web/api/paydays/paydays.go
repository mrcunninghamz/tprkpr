package paydays

import (
	"github.com/a-h/templ"
	"github.com/gin-gonic/gin"
	"github.com/mrcunninghamz/tprkpr/platform/data"
	"github.com/mrcunninghamz/tprkpr/platform/data/models"
	"github.com/mrcunninghamz/tprkpr/web/views"
)

func Get(dbcontext *data.DataContext) gin.HandlerFunc {
	var paydays []models.Payday
	dbcontext.DB.Preload("Bills").Find(&paydays)
	return func(ctx *gin.Context) {
		templ.Handler(views.PayDays(paydays)).ServeHTTP(ctx.Writer, ctx.Request)
	}
}

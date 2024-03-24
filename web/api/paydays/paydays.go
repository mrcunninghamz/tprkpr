package paydays

import (
	"encoding/json"
	"github.com/a-h/templ"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/gofrs/uuid"
	"github.com/mrcunninghamz/tprkpr/platform/authenticator"
	"github.com/mrcunninghamz/tprkpr/platform/data/models"
	"github.com/mrcunninghamz/tprkpr/platform/services"
	"github.com/mrcunninghamz/tprkpr/web/views"
	"net/http"
)

func Get(paydayService services.Paydays) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		session := sessions.Default(ctx)
		profile, _ := json.Marshal(session.Get("profile").(map[string]interface{}))

		userProfile := authenticator.UserProfile{}
		json.Unmarshal(profile, &userProfile)

		paydays := paydayService.GetPaydays(userProfile.Id)

		templ.Handler(views.PayDays(paydays)).ServeHTTP(ctx.Writer, ctx.Request)
	}
}

func View(paydayService services.Paydays) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id := ctx.Param("id")
		paydayID, err := uuid.FromString(id)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid payday ID"})
			return
		}
		payday, err := paydayService.GetPayday(paydayID)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get payday"})
			return
		}

		templ.Handler(views.PaydayCard(payday)).ServeHTTP(ctx.Writer, ctx.Request)
	}
}

func NewForm(ctx *gin.Context) {
	emptyBill := &models.Payday{}
	templ.Handler(views.PaydayForm(emptyBill)).ServeHTTP(ctx.Writer, ctx.Request)
}

func Create(paydayService services.Paydays) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		session := sessions.Default(ctx)
		profile, _ := json.Marshal(session.Get("profile").(map[string]interface{}))

		userProfile := authenticator.UserProfile{}
		json.Unmarshal(profile, &userProfile)

		newPayday := &models.Payday{}

		err := ctx.ShouldBindJSON(newPayday)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		newPayday.UserID = userProfile.Id

		createdPayday, err := paydayService.CreatePayday(newPayday)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create payday"})
			return
		}

		templ.Handler(views.PaydayCard(createdPayday)).ServeHTTP(ctx.Writer, ctx.Request)
	}
}

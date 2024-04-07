package paydays

import (
	"github.com/a-h/templ"
	"github.com/gin-gonic/gin"
	"github.com/gofrs/uuid"
	"github.com/mrcunninghamz/tprkpr/platform/data/models"
	"github.com/mrcunninghamz/tprkpr/platform/services"
	"github.com/mrcunninghamz/tprkpr/web/api/utilities"
	"github.com/mrcunninghamz/tprkpr/web/views"
	"net/http"
)

func EditForm(paydayService services.Paydays) gin.HandlerFunc {
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

		templ.Handler(views.PaydayForm(&payday)).ServeHTTP(ctx.Writer, ctx.Request)
	}
}

func Get(paydayService services.Paydays) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		userProfile, err := utilities.GetUser(ctx)

		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to user profile"})
			return
		}

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
		userProfile, err := utilities.GetUser(ctx)

		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to user profile"})
			return
		}

		newPayday := &models.Payday{}

		err = ctx.ShouldBindJSON(newPayday)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		newPayday.UserID = userProfile.Id

		_, err = paydayService.CreatePayday(newPayday)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create payday"})
			return
		}

		paydays := paydayService.GetPaydays(userProfile.Id)

		templ.Handler(views.PayDays(paydays)).ServeHTTP(ctx.Writer, ctx.Request)
	}
}

func Delete(paydayService services.Paydays) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id := ctx.Param("id")
		paydayID, err := uuid.FromString(id)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid payday ID"})
			return
		}

		err = paydayService.DeletePayday(paydayID)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete payday"})
			return
		}

		templ.Handler(views.Empty()).ServeHTTP(ctx.Writer, ctx.Request)
	}
}

func Update(paydayService services.Paydays) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id := ctx.Param("id")
		paydayID, err := uuid.FromString(id)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid payday ID"})
			return
		}

		updatedPayday := &models.Payday{
			ID: paydayID,
		}

		err = ctx.ShouldBindJSON(updatedPayday)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		err = paydayService.UpdatePayday(updatedPayday)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update payday"})
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

package worksheets

import (
	"fmt"
	"github.com/a-h/templ"
	"github.com/gin-gonic/gin"
	"github.com/gofrs/uuid"
	"github.com/mrcunninghamz/tprkpr/platform/data/models"
	"github.com/mrcunninghamz/tprkpr/platform/services"
	"github.com/mrcunninghamz/tprkpr/web/api/utilities"
	"github.com/mrcunninghamz/tprkpr/web/views"
	"net/http"
)

func Get(worksheetService services.Worksheets) func(c *gin.Context) {
	return func(ctx *gin.Context) {
		userProfile, err := utilities.GetUser(ctx)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		worksheets, err := worksheetService.GetWorksheets(userProfile.Id)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		groupedWorksheets := GroupWorksheetsByYearAndMonth(worksheets)
		templ.Handler(views.Worksheets(groupedWorksheets)).ServeHTTP(ctx.Writer, ctx.Request)
	}
}

func Delete(worksheetService services.Worksheets) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id := ctx.Param("id")
		worksheetId, err := uuid.FromString(id)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid worksheet ID"})
			return
		}

		err = worksheetService.DeleteWorksheet(worksheetId)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete worksheet"})
			return
		}

		templ.Handler(views.Empty()).ServeHTTP(ctx.Writer, ctx.Request)
	}
}

func GroupWorksheetsByYearAndMonth(worksheets []models.Worksheet) map[string][]models.Worksheet {
	result := make(map[string][]models.Worksheet)

	for _, worksheet := range worksheets {
		key := fmt.Sprintf("%d-%d", worksheet.Year, worksheet.Month)

		result[key] = append(result[key], worksheet)
	}

	return result
}

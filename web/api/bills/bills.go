package bills

import (
	"github.com/a-h/templ"
	"github.com/gin-gonic/gin"
	"github.com/gofrs/uuid"
	"github.com/mrcunninghamz/tprkpr/platform/data/models"
	"github.com/mrcunninghamz/tprkpr/platform/services"
	"github.com/mrcunninghamz/tprkpr/web/views"
	"net/http"
)

func EditForm(billService services.Bills) func(c *gin.Context) {
	return func(ctx *gin.Context) {
		id := ctx.Param("id")

		billID, err := uuid.FromString(id)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid bill ID"})
			return
		}

		bill, err := billService.GetBill(billID)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get bill"})
			return
		}
		templ.Handler(views.BillForm(&bill)).ServeHTTP(ctx.Writer, ctx.Request)
	}
}

func Update(billService services.Bills) func(c *gin.Context) {
	return func(ctx *gin.Context) {
		id := ctx.Param("id")
		bill := &models.Bill{}

		err := ctx.ShouldBindJSON(bill)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
			return
		}

		billID, err := uuid.FromString(id)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid bill ID"})
			return
		}

		bill.ID = billID
		updatedBill, err := billService.UpdateBill(bill)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update bill"})
			return
		}
		templ.Handler(views.BillDetail(&updatedBill, true)).ServeHTTP(ctx.Writer, ctx.Request)
	}
}

func View(billService services.Bills) func(c *gin.Context) {
	return func(ctx *gin.Context) {
		id := ctx.Param("id")

		billID, err := uuid.FromString(id)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid bill ID"})
			return
		}

		bill, err := billService.GetBill(billID)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get bill"})
			return
		}
		templ.Handler(views.BillDetail(&bill, true)).ServeHTTP(ctx.Writer, ctx.Request)
	}
}

func NewForm(ctx *gin.Context) {
	paydayId := ctx.Param("paydayId")

	paydayID, err := uuid.FromString(paydayId)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid payday ID"})
		return
	}

	emptyBill := &models.Bill{
		Name:     "New Bill",
		PaydayId: paydayID,
	}
	templ.Handler(views.BillForm(emptyBill)).ServeHTTP(ctx.Writer, ctx.Request)
}

func Create(billService services.Bills) func(c *gin.Context) {
	return func(ctx *gin.Context) {
		bill := &models.Bill{}

		err := ctx.ShouldBindJSON(bill)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
			return
		}

		createdBill, err := billService.CreateBill(bill)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create bill"})
			return
		}
		templ.Handler(views.BillDetail(&createdBill, true)).ServeHTTP(ctx.Writer, ctx.Request)
	}
}

func Delete(billService services.Bills) func(c *gin.Context) {
	return func(ctx *gin.Context) {
		id := ctx.Param("id")

		billID, err := uuid.FromString(id)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid bill ID"})
			return
		}

		err = billService.DeleteBill(billID)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete bill"})
			return
		}
		templ.Handler(views.Empty()).ServeHTTP(ctx.Writer, ctx.Request)
	}
}

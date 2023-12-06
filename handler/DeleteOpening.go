package handler

import (
	"fmt"
	"net/http"

	"github.com/C4ndid0/gojobsearh/schemas"
	"github.com/gin-gonic/gin"
)

func DeleteOpeningHandler(ctx *gin.Context) {
	id := ctx.Query("id")
	if id == "" {
		sendError(ctx, http.StatusBadRequest, errParamsIsRequired("id", "queryParameter").Error())
		return
	}

	opening := schemas.Opening{}
	// Find Opening
	if err := db.First(&opening, id).Error; err != nil {
		sendError(ctx, http.StatusFound, fmt.Sprintf("opening whit id: %s not found", id))
		return
	}

	// Delete Opening

	if err := db.Delete(&opening).Error; err != nil {
		sendError(ctx, http.StatusInternalServerError, fmt.Sprintf("error deleting opening %s", id))
		return
	}

	sendSuccess(ctx, "delete-opening", opening)
}

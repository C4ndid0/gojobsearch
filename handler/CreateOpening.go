package handler

import (
	"net/http"

	"github.com/C4ndid0/gojobsearh/schemas"
	"github.com/gin-gonic/gin"
)
    
 
        
func CreateOpeningHandler(ctx *gin.Context) {
  request := CreateOpeningRequest{}

  ctx.BindJSON(&request)

  if err := request.validate(); err != nil{
    logger.Errorf("validation error: %v", err.Error())
    sendError(ctx, http.StatusBadRequest, err.Error())
    return 
  }

  opening := schemas.Opening{

    Role: request.Role,
    Company: request.Company,
    Location: request.Location,
    Remote: *request.Remote,
    Link: request.Link,
    Salary: int(request.Salary),
  }

  if err := db.Create(&request).Error; err != nil {
    logger.Errorf("error creating opening: %v", err.Error())
    sendError(ctx, http.StatusInternalServerError, "error creting opening on database")
    return 
  }

  sendSuccess(ctx,"create-opening", opening)

}




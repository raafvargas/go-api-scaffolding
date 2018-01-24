package controllers

import (
	"strconv"

	"go-api-scaffolding/db"

	"github.com/gin-gonic/gin"
)

//ModelController ...
type ModelController struct{}

// @Summary Search something.
// @Description Get something by filter.
// @Produce  json
// @Param   size     query    int     true        "Size"
// @Param   from     query    int     true        "From"
// @Param   filter     query    string     true        "Filter"
// @Success 200 {object} models.Model
// @Failure 401 {object} models.Error
// @Failure 400 {object} models.Error
// @Router /api/models [get]
func (controller ModelController) Find(context *gin.Context) {
	from := context.Query("from")
	size := context.Query("size")
	filter := context.Query("filter")

	if filter == "" {
		context.JSON(400, Error("Filter is missing."))
		return
	}

	sizeInt, err := strconv.ParseInt(size, 10, 32)
	if err != nil {
		context.JSON(400, Error("Size is missing."))
		return
	}

	fromInt, err := strconv.ParseInt(from, 10, 32)
	if err != nil {
		context.JSON(400, Error("From is missing."))
		return
	}

	modelRep := new(db.ModelRepository)
	models := modelRep.Find(filter, int(fromInt), int(sizeInt))
	context.JSON(200, models)
}

// @Summary Search something by id.
// @Description Gets something by id.
// @Produce  json
// @Param   modelID     path    string     true        "Id"
// @Success 200 {object} models.Model
// @Failure 401 {object} models.Error
// @Failure 400 {object} models.Error
// @Router /api/models/:modelID [get]
func (controller ModelController) Get(context *gin.Context) {
	modelID := context.Param("modelID")

	if modelID == "" {
		context.JSON(400, Error("modelID is missing."))
		return
	}

	modelRep := new(db.ModelRepository)
	model := modelRep.Get(modelID)
	context.JSON(200, model)
}

//Error ...
func Error(message string) gin.H {
	return gin.H{"message": message}
}

package routergin

import (
	"antia/internal/infrastructure/api/handlers"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (router *RouterGin) AddRuneForTeam(ctx *gin.Context) {
	var req handlers.RelationsRequest
	if err := ctx.ShouldBind(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	result, err := router.hs.AddRuneForTeam(ctx, &req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, result)

}

func (router *RouterGin) DeleteRelationByID(ctx *gin.Context) {
	var req handlers.DeleteRelationsRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	err := router.hs.DeleteRelationByID(ctx, &req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, "ok")

}

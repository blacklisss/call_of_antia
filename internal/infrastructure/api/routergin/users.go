package routergin

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type getUserRequest struct {
	ID uint64 `uri:"id" binding:"required"`
}

func (router *RouterGin) GetByUserID(ctx *gin.Context) {
	var req getUserRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	response, err := router.hs.GetUserByID(ctx, req.ID)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.HTML(http.StatusOK, "index.html", gin.H{
		"title":    response.User.Name,
		"response": response,
	})

}

package presentationbreakdown

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hernangonzalez1987/scriptBreakdown/internal/domain/_interfaces"
	"github.com/hernangonzalez1987/scriptBreakdown/internal/domain/entity"
)

type PresentationBreakdown struct {
	service _interfaces.ScriptBreakdownRequestUseCase
}

func New(service _interfaces.ScriptBreakdownRequestUseCase) *PresentationBreakdown {
	return &PresentationBreakdown{
		service: service,
	}
}

func (ref *PresentationBreakdown) ProcessFile(ctx *gin.Context) {
	request := BreakdownRequest{}

	err := ctx.ShouldBind(&request)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, NewErrorResponse(err))

		return
	}

	fileHeader, err := ctx.FormFile("file")
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, NewErrorResponse(err))

		return
	}

	file, err := fileHeader.Open()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, NewErrorResponse(err))

		return
	}

	result, err := ref.service.RequestScriptBreakdown(ctx, entity.ScriptBreakdownRequest{TempScriptFile: file})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, NewErrorResponse(err))

		return
	}

	ctx.JSON(http.StatusOK, NewBreakdownRequestResponse(*result))
}

package presentationbreakdown

import (
	"errors"
	"io"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/hernangonzalez1987/scriptBreakdown/internal/_interfaces"
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

func (ref *PresentationBreakdown) GetResult(ctx *gin.Context) {
	breakdownID := ctx.Param("breakdownID")
	if breakdownID == "" {
		ctx.JSON(http.StatusBadGateway, NewErrorResponse(
			errors.New("breakdownID is mandatory"),
		))

		return
	}

	result, err := ref.service.GetResult(ctx, breakdownID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, NewErrorResponse(err))

		return
	}

	if result.Content != nil {
		tempFile, err := os.CreateTemp("", "")
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, NewErrorResponse(err))

			return
		}

		_, err = io.Copy(tempFile, result.Content)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, NewErrorResponse(err))

			return
		}

		tempFile.Close()
		result.Content.Close()

		ctx.File(tempFile.Name())

		return
	}

	ctx.JSON(http.StatusOK, NewBreakdownRequestResponse(*result))
}

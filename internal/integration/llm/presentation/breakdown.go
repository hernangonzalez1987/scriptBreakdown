package presentationbreakdown

import (
	"net/http"
	"path/filepath"

	"github.com/gin-gonic/gin"
	"github.com/hernangonzalez1987/scriptBreakdown/internal/domain/_interfaces"
	"github.com/hernangonzalez1987/scriptBreakdown/internal/domain/entity"
	"github.com/rs/zerolog"
)

type PresentationBreakdown struct {
	service _interfaces.ScriptBreakdownUseCase
}

func New(service _interfaces.ScriptBreakdownUseCase) *PresentationBreakdown {
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

	dst := filepath.Base(request.File.Filename)

	err = ctx.SaveUploadedFile(request.File, dst)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, NewErrorResponse(err))

		return
	}

	result, err := ref.service.ScriptBreakdown(ctx, entity.ScriptBreakdownRequest{ScriptFileName: dst})
	if err != nil {
		zerolog.Ctx(ctx).Err(err).Msg("error on script breakdown")
		ctx.JSON(http.StatusInternalServerError, NewErrorResponse(err))

		return
	}

	ctx.File(result.BreakdownTempFileName)
}

package presentationbreakdown

import (
	"net/http"
	"path/filepath"

	"github.com/gin-gonic/gin"
	"github.com/hernangonzalez1987/scriptBreakdown/internal/domain/_interfaces"
	"github.com/hernangonzalez1987/scriptBreakdown/internal/domain/entity"
)

type presentationBreakdown struct {
	_interfaces.BreakdownPresentation
	service _interfaces.ScriptBreakdownUseCase
}

func New(service _interfaces.ScriptBreakdownUseCase) _interfaces.BreakdownPresentation {
	return &presentationBreakdown{
		service: service,
	}
}

func (ref *presentationBreakdown) ProcessFile(ctx *gin.Context) {

	request := BreakdownRequest{}

	err := ctx.ShouldBind(&request)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"error": err.Error()})
		return
	}

	dst := filepath.Base(request.File.Filename)

	err = ctx.SaveUploadedFile(request.File, dst)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	_, err = ref.service.ProcessFile(ctx, entity.ScriptBreakdownRequest{FilePath: dst})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"status": "successful"})

}

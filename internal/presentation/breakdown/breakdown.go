package presentationbreakdown

import (
	"io"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/hernangonzalez1987/scriptBreakdown/api"
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

// BrakdownScript godoc
// @Summary      Creates a breakdown script
// @Description  Creates a breakdown script requests, the result should be async obtain later from GET method.
// @Tags         breakdwn
// @Accept       multipart/form-data
// @Produce      json
// @Success      201  {object}  api.BreakdownRequestResponse
// @Failure      400  {object}  api.ErrorResponse
// @Failure      500  {object}  api.ErrorResponse
// @Router       /script/breakdown [post]
func (ref *PresentationBreakdown) BreakdownScript(ctx *gin.Context) {
	request := api.BreakdownRequest{}

	err := ctx.ShouldBind(&request)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, api.NewErrorResponse(err))

		return
	}

	fileHeader, err := ctx.FormFile("file")
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, api.NewErrorResponse(err))

		return
	}

	file, err := fileHeader.Open()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, api.NewErrorResponse(err))

		return
	}

	result, err := ref.service.RequestScriptBreakdown(ctx, entity.ScriptBreakdownRequest{TempScriptFile: file})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, api.NewErrorResponse(err))

		return
	}

	ctx.JSON(http.StatusCreated, api.NewBreakdownRequestResponse(*result))
}

// BrakdownScript godoc
// @Summary      Gets a breakdown script result
// @Description  Gets the result of a breakdown script
// @Tags         breakdwn
// @Produce      json
// @Param 		 breakdown_id path string true "BreakdownID"
// @Success      201  {object}  api.BreakdownRequestResponse
// @Failure      400  {object}  api.ErrorResponse
// @Failure      404  {object}  api.ErrorResponse
// @Failure      500  {object}  api.ErrorResponse
// @Router       /script/breakdown/{breakdown_id} [get]
func (ref *PresentationBreakdown) GetResult(ctx *gin.Context) {
	breakdownID := ctx.Param("breakdownID")

	result, err := ref.service.GetResult(ctx, breakdownID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, api.NewErrorResponse(err))

		return
	}

	if result.Content != nil {
		tempFile, err := os.CreateTemp("", "")
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, api.NewErrorResponse(err))

			return
		}

		_, err = io.Copy(tempFile, result.Content)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, api.NewErrorResponse(err))

			return
		}

		tempFile.Close()
		result.Content.Close()

		ctx.File(tempFile.Name())

		return
	}

	ctx.JSON(http.StatusOK, api.NewBreakdownRequestResponse(*result))
}

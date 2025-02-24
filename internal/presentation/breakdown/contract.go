package presentationbreakdown

import (
	"errors"
	"mime/multipart"

	"github.com/hernangonzalez1987/scriptBreakdown/internal/domain/entity"
	valueobjects "github.com/hernangonzalez1987/scriptBreakdown/internal/domain/valueObjects"
)

type BreakdownRequest struct {
	File *multipart.FileHeader `binding:"required" form:"file"`
}

type ErrorResponse struct {
	Code        string `json:"code"`
	Description string `json:"description"`
}

func NewErrorResponse(err error) ErrorResponse {
	var customErr valueobjects.CustomError

	if errors.As(err, &customErr) {
		return ErrorResponse{
			Code:        customErr.Code,
			Description: customErr.Desc,
		}
	}

	return ErrorResponse{
		Code:        "UNKNOWN",
		Description: err.Error(),
	}
}

type BreakdownRequestResponse struct {
	BreakdownID string `json:"breakdownId"`
}

func NewBreakdownRequestResponse(result entity.ScriptBreakdownResult) BreakdownRequestResponse {
	return BreakdownRequestResponse{
		BreakdownID: result.BreakdownID,
	}
}

package presentationbreakdown

import "mime/multipart"

type BreakdownRequest struct {
	File *multipart.FileHeader `form:"file" binding:"required"`
}

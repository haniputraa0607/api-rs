package utility

import (
	"api-rs/schemas"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Code    int           `json:"code"`
	Message string        `json:"message"`
	Data    any           `json:"data"`
	Meta    *schemas.Meta `json:"meta,omitempty"`
}

func ApiResponse(ctx *gin.Context, code int, message string, data any, meta *schemas.Meta) {
	jsonResponse := Response{
		Code:    code,
		Message: message,
		Data:    data,
	}
	if meta != nil {
		jsonResponse.Meta = meta
	}

	ctx.JSON(code, jsonResponse)
}

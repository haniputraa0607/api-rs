package examplehandler

import (
	"api-rs/utility"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ExampleHandler interface {
	Example(ctx *gin.Context)
}

type exampleHandler struct {
}

func NewExampleHandler() *exampleHandler {
	return &exampleHandler{}
}

func (h *exampleHandler) Example(ctx *gin.Context) {
	utility.ApiResponse(ctx, http.StatusOK, "success", "TEST EXAMPLE", nil)
}

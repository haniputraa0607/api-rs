package utility

import (
	"bytes"
	"errors"
	"fmt"
	"io"

	"github.com/gin-gonic/gin"
)

func ShouldBindJSON(ctx *gin.Context, request any) error {
	bodyBytes, err := io.ReadAll(ctx.Request.Body)
	if err != nil {
		message := fmt.Sprintf("error reading the request body: %v", err)
		return errors.New(message)
	}

	ctx.Request.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))
	err = ctx.ShouldBindJSON(&request)
	ctx.Request.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))
	if err != nil {
		return err
	}

	return nil
}

func ShouldBind(ctx *gin.Context, request any) error {
	var buf bytes.Buffer
	tee := io.TeeReader(ctx.Request.Body, &buf)

	bodyBytes, err := io.ReadAll(tee)
	if err != nil {
		message := fmt.Sprintf("error reading the request body: %v", err)
		return errors.New(message)
	}

	ctx.Request.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))
	err = ctx.ShouldBind(request)
	if err != nil {
		error := FormatValidationError(err)
		return errors.New(error)
	}
	ctx.Request.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))

	return nil
}

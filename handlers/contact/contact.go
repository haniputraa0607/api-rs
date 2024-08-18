package contacthandler

import (
	"api-rs/schemas"
	contactservice "api-rs/services/contact"
	"api-rs/utility"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ContactHandler interface {
	GetContacts(ctx *gin.Context)
	CreateContact(ctx *gin.Context)
	DeleteContact(ctx *gin.Context)
	GetContact(ctx *gin.Context)
	UpdateContact(ctx *gin.Context)
}

type contactHandler struct {
	contactService contactservice.ContactService
}

func NewContactHandler(contactService contactservice.ContactService) *contactHandler {
	return &contactHandler{contactService}
}

func (h *contactHandler) GetContacts(ctx *gin.Context) {
	var request schemas.Common
	err := utility.ShouldBind(ctx, &request)
	if err != nil {
		error := utility.FormatValidationError(err)
		utility.ApiResponse(ctx, http.StatusBadRequest, error, nil, nil)
		return
	}

	contacts, meta, err := h.contactService.ListContactsPagination(request)
	if err != nil {
		utility.ApiResponse(ctx, http.StatusBadRequest, err.Error(), nil, nil)
		return
	}

	utility.ApiResponse(ctx, http.StatusOK, "success", contacts, meta)
}

func (h *contactHandler) CreateContact(ctx *gin.Context) {
	var request schemas.CreateUpdateContactRequest
	err := utility.ShouldBind(ctx, &request)
	if err != nil {
		error := utility.FormatValidationError(err)
		utility.ApiResponse(ctx, http.StatusBadRequest, error, nil, nil)
		return
	}

	err = h.contactService.CreateContact(request)
	if err != nil {
		utility.ApiResponse(ctx, http.StatusBadRequest, err.Error(), nil, nil)
		return
	}

	utility.ApiResponse(ctx, http.StatusOK, "success", nil, nil)
}

func (h *contactHandler) DeleteContact(ctx *gin.Context) {
	contactID := ctx.Param("id")

	err := h.contactService.DeleteContact(contactID)
	if err != nil {
		utility.ApiResponse(ctx, http.StatusBadRequest, err.Error(), nil, nil)
		return
	}

	utility.ApiResponse(ctx, http.StatusOK, "success", nil, nil)
}

func (h *contactHandler) GetContact(ctx *gin.Context) {
	contactID := ctx.Param("id")

	contact, err := h.contactService.GetContact(contactID)
	if err != nil {
		utility.ApiResponse(ctx, http.StatusBadRequest, err.Error(), nil, nil)
		return
	}

	utility.ApiResponse(ctx, http.StatusOK, "success", contact, nil)
}

func (h *contactHandler) UpdateContact(ctx *gin.Context) {
	var request schemas.CreateUpdateContactRequest
	err := utility.ShouldBind(ctx, &request)
	if err != nil {
		error := utility.FormatValidationError(err)
		utility.ApiResponse(ctx, http.StatusBadRequest, error, nil, nil)
		return
	}

	contactID := ctx.Param("id")
	err = h.contactService.UpdateContact(contactID, request)
	if err != nil {
		utility.ApiResponse(ctx, http.StatusBadRequest, err.Error(), nil, nil)
		return
	}

	utility.ApiResponse(ctx, http.StatusOK, "success", nil, nil)
}

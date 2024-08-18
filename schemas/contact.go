package schemas

type ListContactResponse struct {
	ID    uint64 `json:"id"`
	Name  string `json:"name"`
	Value string `json:"value"`
}

type CreateUpdateContactRequest struct {
	Name  string  `form:"name" binding:"required"`
	Value string  `form:"value" binding:"required"`
	Icon  *string `form:"icon"`
}

type ContactResponse struct {
	ID    uint64  `json:"id"`
	Name  string  `json:"name"`
	Value string  `json:"value"`
	Icon  *string `json:"icon"`
}

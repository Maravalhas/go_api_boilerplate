package structs

type Pagination struct {
	Offset   *int    `json:"offset" form:"offset" validate:"omitempty,min=0"`
	Limit    *int    `json:"limit" form:"limit" validate:"omitempty,min=1,max=100"`
	OrderBy  *string `json:"order_by" form:"order_by" validate:"omitempty"`
	OrderDir *string `json:"order_dir" form:"order_dir" validate:"omitempty"`
	Search   *string `json:"search" form:"search" validate:"omitempty"`
}

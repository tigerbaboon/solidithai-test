package base

type Pagination struct {
	Page    int    `form:"page"`
	Size    int    `form:"size"`
	OrderBy string `form:"order_by"`
	SortBy  string `form:"sort_by"`
}

var (
	DEFAULT_PAGINATION = Pagination{
		Page:    1,
		Size:    10,
		OrderBy: `asc`,
		SortBy:  `created_at`,
	}
)

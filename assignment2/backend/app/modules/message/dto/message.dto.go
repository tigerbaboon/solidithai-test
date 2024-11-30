package messagedto

type CreateMessageRequest struct {
	UserID string `json:"userId"`
	Text   string `json:"text"`
}

type GetMessageListRequest struct {
	Page    int    `form:"page"`
	Size    int    `form:"size"`
	SortBy  string `form:"sort_by"`
	OrderBy string `form:"order_by"`
	Search  string `form:"search"`
}

type MessageResponse struct {
	ID        string `json:"id"`
	UserID    string `json:"userId"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Avatar    string `json:"avatar"`
	Text      string `json:"text"`
	CreatedAt string `json:"createdAt"`
}

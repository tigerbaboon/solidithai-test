package userdto

type CreateUserRequest struct {
	FirstName string `json:"firstName" binding:"required"`
	LastName  string `json:"lastName" binding:"required"`
	Email     string `json:"email" binding:"required,email"`
	Password  string `json:"password" binding:"required,min=8"`
	Avatar    string `json:"avatar" binding:"required"`
}

type UpdateUserRequest struct {
	FirstName string `json:"firstName" binding:"required"`
	LastName  string `json:"lastName" binding:"required"`
	Email     string `json:"email" binding:"required,email"`
	Avatar    string `json:"avatar" binding:"required"`
}

type GetUserByIDRequest struct {
	ID string `uri:"id" binding:"required"`
}

type GetUserListRequest struct {
	Page    int    `form:"page"`
	Size    int    `form:"size"`
	SortBy  string `form:"sort_by"`
	OrderBy string `form:"order_by"`
	Search  string `form:"search"`
}

type UserResponse struct {
	ID        string `json:"id"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email     string `json:"email"`
	Avatar    string `json:"avatar"`
}

type UpdatePasswordRequest struct {
	Password string `json:"password" binding:"required,min=8"`
}

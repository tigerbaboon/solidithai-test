package base

type ResponseStatus struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type Response[T any] struct {
	*ResponseStatus
	Data     T                 `json:"data"`
	Paginate *ResponsePaginate `json:",omitempty"`
}

type ResponsePaginate struct {
	Page  int64 `json:"page"`
	Size  int64 `json:"size"`
	Total int64 `json:"total"`
}

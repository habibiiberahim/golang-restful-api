package model

type WebResponse[T any] struct {
	Data   T             `json:"data"`
	Paging *PageMetaData `json:"paging,omitempty"`
	Errors string        `json:"errors,omitempty"`
}

type PageResponse[T any] struct {
	Data         []T          `json:"data,omitempty"`
	PageMetaData PageMetaData `json:"paging,omitempty"`
}

type PageMetaData struct {
	Page      int `json:"page"`
	Size      int `json:"size"`
	TotalItem int `json:"total_item"`
	TotalPage int `json:"total_page"`
}

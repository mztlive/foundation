package response

// Collection 集合
type Collection[T any] struct {
	Total int64 `json:"total"`
	Items []T   `json:"items"`
}

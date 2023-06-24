package binding

type Delete struct {
	ID        string `json:"id"`
	DeletedAt int64  `json:"deleted_at"`
	DeletedBy string `json:"deleted_by"`
}

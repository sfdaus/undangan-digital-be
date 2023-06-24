package binding

type Update struct {
	ID        string `json:"id"`
	UpdatedAt int64  `json:"updated_at"`
	UpdatedBy string `json:"updated_by"`
	IsActive  bool   `json:"is_active"`
	Scope     string `json:"scope"`
	Name      string `json:"name"`
}

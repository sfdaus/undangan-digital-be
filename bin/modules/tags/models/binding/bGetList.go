package binding

type GetList struct {
	Page    int    `query:"page"`
	PerPage int    `query:"per_page"`
	Search  string `query:"search"`

	ID        string `query:"id"`
	DeletedAt int64  `query:"deleted_at"`
	DeletedBy string `query:"deleted_by"`
	IsActive  bool   `query:"is_active"`
	Scope     string `query:"scope"`
	Name      string `query:"name"`
}

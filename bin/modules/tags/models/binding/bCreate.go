package binding

type Create struct {
	Scope string `json:"scope" validate:"required,lte=36"`
	Name  string `json:"name" validate:"required"`
}

package models

var Table = "tags"

type Tags struct {
	ID        string `json:"id" gorm:"primaryKey;type:varchar(36)"`
	CreatedAt int64  `json:"created_at"`
	CreatedBy string `json:"created_by" gorm:"type:varchar(36)"`
	UpdatedAt int64  `json:"updated_at"`
	UpdatedBy string `json:"updated_by" gorm:"type:varchar(36)"`
	DeletedAt int64  `json:"deleted_at"`
	DeletedBy string `json:"deleted_by" gorm:"type:varchar(36)"`
	IsActive  bool   `json:"is_active"`
	Scope     string `json:"scope" gorm:"not null;type:varchar(36)"`
	Name      string `json:"name" gorm:"not null"`
}

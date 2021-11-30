package models

type BaseModel struct {
	ID uint64 `gorm:"column:id;primaryKey;autoIncrement;not null"`
	CreatedAt int64 `gorm:"column:created_at;index;autoCreateTime"`
	UpdatedAt int64 `gorm:"column:updated_at;index;autoCreateTime"`
}

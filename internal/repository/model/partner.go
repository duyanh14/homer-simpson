package model

type (
	Partner struct {
		ID         uint   `gorm:"primaryKey"`
		Name       string `gorm:"column:name"`
		Code       string `gorm:"column:code"`
		Descrition string `gorm:"column:description"`
		Alias      string `gorm:"column:alias"`
		Address    string `gorm:"column:address"`
		IsDisable  string `gorm:"column:is_disable"` // default active(false), unactive(true)
		CreatedBy  string `gorm:"column:created_by"`
		UpdatedBy  string `gorm:"column:updated_by"`
		CreatedAt  uint64 `gorm:"column:created_at"`
		UpdatedAt  uint64 `gorm:"column:updated_at"`
		DeletedAt  uint64 `gorm:"column:deleted_at"`
		ExpiredAt  uint64 `gorm:"column:expired_at"`
	}
)

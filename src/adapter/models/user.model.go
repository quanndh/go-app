package models

type User struct {
	ID        uint   `json:"id" gorm:"primaryKey"`
	Username  string `json:"username"`
	Password  string `json:"password"`
	CreatedAt int64  `json:"createdAt" gorm:"autoCreateTime:milli"`
	UpdatedAt int64  `json:"updatedAt" gorm:"autoUpdateTime:milli"`
}

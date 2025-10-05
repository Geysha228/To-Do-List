package models

import "time"

type User struct {
	ID         uint      `gorm:"primaryKey;autoIncrement"`
	TelegramID int64     `gorm:"unique;not null"`
	Username   string    `gorm:"type:varchar(50)"`
	FirstName  string    `gorm:"type:varchar(100)"`
	CreatedAt  time.Time `gorm:"autoCreateTime"`
	UpdatedAt  time.Time `gorm:"autoUpdateTime"`

	Tasks []Task `gorm:"foreignKey:CreatorID"`
}

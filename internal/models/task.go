package models

import "time"

type Task struct {
	ID             uint      `gorm:"primaryKey;autoIncrement"`
	Name           string    `gorm:"type:varchar(100);not null"`
	CreatorID      uint      `gorm:"not null"`
	Creator        User      `gorm:"foreignKey:CreatorID;constraint:OnDelete:CASCADE"`
	CreatedAt      time.Time `gorm:"autoCreateTime"`
	UpdatedAt      time.Time `gorm:"autoUpdateTime"`
	DateOfRemind   *time.Time
	DateOfAppoint  *time.Time
	DateOfComplete *time.Time
	Description    string `gorm:"type:text"`
}

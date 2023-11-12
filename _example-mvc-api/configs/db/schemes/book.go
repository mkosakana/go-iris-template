package schemes

import "time"

type Book struct {
	ID        uint      `gorm:"primaryKey;autoIncrement;auto"`
	CreatedAt time.Time `gorm:"autoCreateTime:nano"`
	UpdatedAt time.Time `gorm:"autoUpdateTime:nano"`
	Title     string    `gorm:"not null"`
	Author    string    `gorm:"not null"`
}

package entity

import "time"

type User struct {
	ID         int64      `gorm:"column:user_id;primary_key"`
	KeycloakID string     `gorm:"column:keycloak_id"`
	Username   string     `gorm:"column:username"`
	Password   string     `gorm:"column:password"`
	Email      string     `gorm:"column:email"`
	Desc       string     `gorm:"column:user_desc"`
	FullName   string     `gorm:"column:full_name"`
	Phone      string     `gorm:"column:phone"`
	Salt       string     `gorm:"column:salt"`
	Role       string     `gorm:"column:role"`
	Disabled   bool       `gorm:"column:disabled"`
	CreatedAt  *time.Time `gorm:"column:created_at"`
	UpdatedAt  *time.Time `gorm:"column:updated_at"`
	LoginAt    *time.Time `gorm:"column:login_at"`
}

func (u *User) TableName() string {
	return "tusers"
}

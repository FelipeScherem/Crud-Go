package modelUsuario

import "time"

type Usuario struct {
	ID       uint   `gorm:"primaryKey" json:"id"`
	Nome     string `json:"nome"`
	Email    string `json:"email"`
	Situacao bool   `json:"situacao"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

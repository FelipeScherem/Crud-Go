package modelLivros

import "time"

type Livros struct {
	ID         uint      `gorm:"primaryKey" json:"id"`
	Nome       string    `json:"username"`
	Autor      string    `json:"email"`
	Lancamento string    `json:"lancamento"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

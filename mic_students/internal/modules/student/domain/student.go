package domain

type Student struct {
	ID    uint   `gorm:"primary_key" json:"id"`
	Name  string `json:"name"`
	Grade int    `json:"grade"`
}

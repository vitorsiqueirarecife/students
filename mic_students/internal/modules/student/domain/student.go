package domain

type Student struct {
	ID    string `gorm:"primary_key" json:"id"`
	Name  string `json:"name"`
	Grade int    `json:"grade"`
}

type GetAllStudentsResponse struct {
	Students   []Student `json:"students"`
	Total      int       `json:"total"`
	TotalPages int       `json:"totalPages"`
}

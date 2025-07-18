package entity

import "time"

type TodoListCategoryReq struct {
	Name        string `json:"name" validate:"required" name:"Nama"`
	Description string `json:"description" validate:"required" name:"Deskripsi"`
	UserID      int64  `json:"-"` // dummy untuk kebutuhan log
}

type TodoListCategoryResponse struct {
	ID          int64  `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	CreatedAt   string `json:"created_at"`
}

type TodoListCategory struct {
	ID          int64     `gorm:"column:id"`
	Name        string    `gorm:"column:name"`
	Description string    `gorm:"column:description"`
	CreatedAt   time.Time `gorm:"column:created_at"`
}

func (TodoListCategory) TableName() string {
	return "todo_list_categories"
}

// SetUserID implements parser.WithUserID.
func (t *TodoListCategoryReq) SetUserID(_ int64) {
	t.UserID = 12345
}

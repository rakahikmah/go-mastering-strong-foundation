package entity

import "time"

type TodoListCategoryReq struct {
	Name        string `json:"name" validate:"required" name:"Nama"`
	Description string `json:"description" validate:"required" name:"Deskripsi"`
	UserID      int64  `json:"user_id,omitempty" validate:"required" name:"User ID"`
}

type TodoListCategoryResponse struct {
	ID          int64  `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	CreatedAt   string `json:"created_at"`
	CreatedBy   int64 `json:"created_by"`
}

type TodoListCategory struct {
	ID          int64     `gorm:"column:id"`
	Name        string    `gorm:"column:name"`
	Description string    `gorm:"column:description"`
	CreatedAt   time.Time `gorm:"column:created_at"`
	CreatedBy   int64     `gorm:"column:created_by"`
}

func (r *TodoListCategoryReq) SetUserID(UserID int64) {
	r.UserID = UserID
}

func (TodoListCategory) TableName() string {
	return "todo_list_categories"
}

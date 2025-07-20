package todo_list_category_usecase

import (
	"context"
	"fmt"

	// "time"

	// errwrap "github.com/pkg/errors"
	generalEntity "github.com/rahmatrdn/go-skeleton/entity"
	// "github.com/rahmatrdn/go-skeleton/internal/helper"
	"github.com/rahmatrdn/go-skeleton/internal/helper"
	"github.com/rahmatrdn/go-skeleton/internal/repository/mysql"
	myentity "github.com/rahmatrdn/go-skeleton/internal/repository/mysql/entity"
	"github.com/rahmatrdn/go-skeleton/internal/usecase/todo_list_category/entity"
	// "github.com/rahmatrdn/go-skeleton/internal/usecase"
	// "github.com/rahmatrdn/go-skeleton/internal/usecase/todo_list/entity"
)

// TODO
// - Create
// - Read
// - Update
// - Delete

type CrudTodoListCategory struct {
	TodoListCategoryRepo mysql.ITodoListCategoryRepository
}

func NewCrudTodoListCategory(
	TodoListCategoryRepo mysql.ITodoListCategoryRepository,
) *CrudTodoListCategory {
	return &CrudTodoListCategory{TodoListCategoryRepo}
}

type ICrudTodoListCategory interface {
	Create(ctx context.Context, req entity.TodoListCategoryReq) error
	GetAll(ctx context.Context) ([]entity.TodoListCategoryResponse, error)
	Update(ctx context.Context, id int64, req entity.TodoListCategoryReq) error // <-- Tambahan
	Delete(ctx context.Context, id int64) error
}

func (u *CrudTodoListCategory) Create(ctx context.Context, req entity.TodoListCategoryReq) error {
	funcName := "CrudTodoListCategory.Create"
	logFields := generalEntity.CaptureFields{
		"user_id": "12345",
	}

	// TODO
	// ambil request data
	// insert ke database
	data := &myentity.TodoListCategory{
		Name:        req.Name,
		Description: req.Description,
		CreatedAt:   helper.DatetimeNowJakarta(),
	}
	err := u.TodoListCategoryRepo.Create(ctx, nil, data, false)
	if err != nil {
		helper.LogError(funcName, "TodoListCategoryRepo.Create", err, logFields, "")
		return err
	}

	return nil
}

func (u *CrudTodoListCategory) GetAll(ctx context.Context) ([]entity.TodoListCategoryResponse, error) {
	funcName := "CrudTodoListCategory.GetAll"
	logFields := generalEntity.CaptureFields{
		"layer": "usecase",
	}

	// Ambil data dari repository
	data, err := u.TodoListCategoryRepo.GetAll(ctx)
	if err != nil {
		helper.LogError(funcName, "TodoListCategoryRepo.GetAll", err, logFields, "")
		return nil, err
	}

	// Mapping ke response DTO
	var result []entity.TodoListCategoryResponse
	for _, row := range data {
		result = append(result, entity.TodoListCategoryResponse{
			ID:          row.ID,
			Name:        row.Name,
			Description: row.Description,
			CreatedAt:   helper.ConvertToJakartaTime(row.CreatedAt),
		})
	}

	return result, nil
}

func (u *CrudTodoListCategory) Update(ctx context.Context, id int64, req entity.TodoListCategoryReq) error {
	funcName := "CrudTodoListCategory.Update"
	logFields := generalEntity.CaptureFields{
		"user_id": "12345",
		"id":      fmt.Sprintf("%d", id),
	}

	oldData, err := u.TodoListCategoryRepo.GetByID(ctx, id)
	if err != nil {
		helper.LogError(funcName, "GetByID", err, logFields, "")
		return err
	}

	changes := &myentity.TodoListCategory{
		Name:        req.Name,
		Description: req.Description,
	}

	err = u.TodoListCategoryRepo.Update(ctx, nil, oldData, changes)
	if err != nil {
		helper.LogError(funcName, "Update", err, logFields, "")
		return err
	}

	return nil
}

func (u *CrudTodoListCategory) Delete(ctx context.Context, id int64) error {
	funcName := "CrudTodoListCategory.Delete"
	logFields := generalEntity.CaptureFields{
		"user_id": "12345",
		"id":      fmt.Sprintf("%d", id),
	}

	// Validasi dulu apakah data dengan ID tersebut ada
	_, err := u.TodoListCategoryRepo.GetByID(ctx, id)
	if err != nil {
		helper.LogError(funcName, "GetByID", err, logFields, "")
		return err
	}

	// Lakukan delete
	err = u.TodoListCategoryRepo.DeleteByID(ctx, nil, id)
	if err != nil {
		helper.LogError(funcName, "DeleteByID", err, logFields, "")
		return err
	}

	return nil
}


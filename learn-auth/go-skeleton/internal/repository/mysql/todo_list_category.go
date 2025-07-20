package mysql

import (
	"context"

	"github.com/rahmatrdn/go-skeleton/config"
	"github.com/rahmatrdn/go-skeleton/internal/helper"
	"github.com/rahmatrdn/go-skeleton/internal/repository/mysql/entity"

	apperr "github.com/rahmatrdn/go-skeleton/error"

	errwrap "github.com/pkg/errors"
	"gorm.io/gorm"
)

type ITodoListCategoryRepository interface {
	TrxSupportRepo
	GetByID(ctx context.Context, ID int64) (e *entity.TodoListCategory, err error)
	Create(ctx context.Context, dbTrx TrxObj, params *entity.TodoListCategory, nonZeroVal bool) error
	Update(ctx context.Context, dbTrx TrxObj, params *entity.TodoListCategory, changes *entity.TodoListCategory) (err error)
	DeleteByID(ctx context.Context, dbTrx TrxObj, id int64) error
	GetAll(ctx context.Context) (result []*entity.TodoListCategory, err error)
}

type TodoListCategoryRepository struct {
	GormTrxSupport
}

func NewTodoListCategoryRepository(mysql *config.Mysql) *TodoListCategoryRepository {
	return &TodoListCategoryRepository{GormTrxSupport{db: mysql.DB}}
}

func (r *TodoListCategoryRepository) GetAll(ctx context.Context) (result []*entity.TodoListCategory, err error) {
	funcName := "TodoListCategoryRepository.GetAll"

	if err := helper.CheckDeadline(ctx); err != nil {
		return nil, errwrap.Wrap(err, funcName)
	}

	err = r.db.Raw("SELECT * FROM todo_list_categories").Scan(&result).Error
	if errwrap.Is(err, gorm.ErrRecordNotFound) {
		return nil, apperr.ErrRecordNotFound()
	}

	return result, err
}

func (r *TodoListCategoryRepository) GetByID(ctx context.Context, ID int64) (result *entity.TodoListCategory, err error) {
	funcName := "TodoListCategoryRepository.GetByID"

	if err := helper.CheckDeadline(ctx); err != nil {
		return nil, errwrap.Wrap(err, funcName)
	}

	err = r.db.Raw("SELECT * FROM todo_list_categories WHERE id = ? LIMIT 1", ID).Scan(&result).Error
	if errwrap.Is(err, gorm.ErrRecordNotFound) {
		return nil, apperr.ErrRecordNotFound()
	}

	return result, err
}

func (r *TodoListCategoryRepository) Create(ctx context.Context, dbTrx TrxObj, params *entity.TodoListCategory, nonZeroVal bool) error {
	funcName := "TodoListCategoryRepository.Create"

	if err := helper.CheckDeadline(ctx); err != nil {
		return errwrap.Wrap(err, funcName)
	}

	cols := helper.NonZeroCols(params, nonZeroVal)
	return r.Trx(dbTrx).Select(cols).Create(&params).Error
}

func (r *TodoListCategoryRepository) Update(ctx context.Context, dbTrx TrxObj, params *entity.TodoListCategory, changes *entity.TodoListCategory) error {
	funcName := "TodoListCategoryRepository.Update"

	if err := helper.CheckDeadline(ctx); err != nil {
		return errwrap.Wrap(err, funcName)
	}

	if params.ID == 0 {
		return errwrap.Wrap(apperr.ErrInvalidRequest(), funcName)
	}

	db := r.Trx(dbTrx).Model(params)

	var err error
	if changes != nil {
		err = db.Updates(*changes).Error
	} else {
		err = db.Updates(helper.StructToMap(params, false)).Error
	}

	if err != nil {
		return errwrap.Wrap(err, funcName)
	}

	return nil
}

func (r *TodoListCategoryRepository) DeleteByID(ctx context.Context, dbTrx TrxObj, id int64) error {
	funcName := "TodoListCategoryRepository.DeleteByID"

	if err := helper.CheckDeadline(ctx); err != nil {
		return errwrap.Wrap(err, funcName)
	}

	err := r.Trx(dbTrx).Where("id = ?", id).Delete(&entity.TodoListCategory{}).Error
	if err != nil {
		return err
	}

	return nil
}

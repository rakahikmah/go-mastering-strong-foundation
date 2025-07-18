package handler

import (
	"net/http"

	"github.com/rahmatrdn/go-skeleton/internal/parser"
	"github.com/rahmatrdn/go-skeleton/internal/presenter/json"
	todo_list_category_usecase "github.com/rahmatrdn/go-skeleton/internal/usecase/todo_list_category"
	"github.com/rahmatrdn/go-skeleton/internal/usecase/todo_list_category/entity"

	fiber "github.com/gofiber/fiber/v2"
)

type TodoListCategoryHandler struct {
	parser                      parser.Parser
	presenter                   json.JsonPresenter
	CrudTodoListCategoryUsecase todo_list_category_usecase.ICrudTodoListCategory
}

func NewTodoListCategoryHandler(
	parser parser.Parser,
	presenter json.JsonPresenter,
	CrudTodoListCategoryUsecase todo_list_category_usecase.ICrudTodoListCategory,
) *TodoListCategoryHandler {
	return &TodoListCategoryHandler{parser, presenter, CrudTodoListCategoryUsecase}
}

func (w *TodoListCategoryHandler) Register(app fiber.Router) {
	app.Post("/todo-list-category", w.Create)
	app.Get("/todo-list-category", w.GetAll)     // <-- Tambahan ini
	app.Put("/todo-list-category/:id", w.Update) // <-- tambahkan ini
	app.Delete("/todo-list-category/:id", w.Delete)

}

func (w *TodoListCategoryHandler) Create(c *fiber.Ctx) error {
	var req entity.TodoListCategoryReq

	err := w.parser.ParserBodyRequest(c, &req)
	if err != nil {
		return w.presenter.BuildError(c, err)
	}

	err = w.CrudTodoListCategoryUsecase.Create(c.Context(), req)
	if err != nil {
		return w.presenter.BuildError(c, err)
	}

	return w.presenter.BuildSuccess(c, nil, "Success", http.StatusOK)
}

func (w *TodoListCategoryHandler) GetAll(c *fiber.Ctx) error {
	result, err := w.CrudTodoListCategoryUsecase.GetAll(c.Context())
	if err != nil {
		return w.presenter.BuildError(c, err)
	}

	return w.presenter.BuildSuccess(c, result, "Success", http.StatusOK)
}

func (w *TodoListCategoryHandler) Update(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil || id <= 0 {
		return w.presenter.BuildError(c, err)
	}

	var req entity.TodoListCategoryReq
	err = w.parser.ParserBodyRequest(c, &req)
	if err != nil {
		return w.presenter.BuildError(c, err)
	}

	err = w.CrudTodoListCategoryUsecase.Update(c.Context(), int64(id), req)
	if err != nil {
		return w.presenter.BuildError(c, err)
	}

	return w.presenter.BuildSuccess(c, nil, "Updated", http.StatusOK)
}

func (w *TodoListCategoryHandler) Delete(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil || id <= 0 {
		return w.presenter.BuildError(c, err)
	}

	err = w.CrudTodoListCategoryUsecase.Delete(c.Context(), int64(id))
	if err != nil {
		return w.presenter.BuildError(c, err)
	}

	return w.presenter.BuildSuccess(c, nil, "Deleted", http.StatusOK)
}

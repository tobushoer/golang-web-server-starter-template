package http

import (
	"net/http"
	"template/internal/pkg/errordef"
	"template/internal/pkg/log"
	"template/pkg/todo"

	"github.com/gin-gonic/gin"
)

// ToDoHandler to do
type ToDoHandler struct {
	toDoUseCase todo.UseCases
}

// NewHandler create to do http handler
func NewHandler(toDoUseCase todo.UseCases) *ToDoHandler {
	return &ToDoHandler{toDoUseCase}
}

type postToDoRequest struct {
	Title   string `json:"title,required" example:"To Do Title" description:"to do title"`
	Content string `json:"content" example:"this is to do content" description:"to do content"`
}

func (p *postToDoRequest) Bind(c *gin.Context) error {
	if err := c.ShouldBindJSON(p); err != nil {
		log.Info(err.Error())
		return err
	}
	return nil
}

func (p *postToDoRequest) Check() error {
	if p.Title == "" {
		return errordef.ErrMsgInvalidRequest
	}
	return nil
}

type postToDoResponse struct {
	ID uint32 `json:"id" example:"1"`
}

// @Title Create a todo.
// @Description Create a todo.
// @Param  todo  body  postToDoRequest  true  "request body"
// @Success  200  object  postToDoResponse  "success"
// @Failure  400  object  errordef.ErrorInfo  "invalid request"
// @Failure  500  object  errordef.ErrorInfo  "internal error"
// @Resource todo
// @Route /api/v1/todo [post]
func (h *ToDoHandler) createToDoEndpoint(c *gin.Context) {
	request := &postToDoRequest{}
	if err := request.Bind(c); err != nil {
		errordef.GinHTTPResponse(c, errordef.ErrorInfo{
			Code:    errordef.ErrCodeInvalidRequest,
			Message: err.Error(),
		})
		return
	}
	if err := request.Check(); err != nil {
		errordef.GinHTTPResponse(c, errordef.ErrorInfo{
			Code:    errordef.ErrCodeInvalidRequest,
			Message: err.Error(),
		})
		return
	}

	id, err := h.toDoUseCase.CreateToDo(todo.ToDo{
		Title:   request.Title,
		Content: request.Content,
	})
	if err != nil {
		errCode := errordef.ErrCodeInternalError
		if err == errordef.ErrMsgInvalidRequest {
			errCode = errordef.ErrCodeInvalidRequest
		}
		errordef.GinHTTPResponse(c, errordef.ErrorInfo{
			Code:    errCode,
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, postToDoResponse{id})
}

type listToDoResp struct {
	Data []todo.ToDo `json:"data"`
}

// @Title List todos.
// @Description List todos.
// @Success  200  object  listToDoResp  "success"
// @Failure  500  object  errordef.ErrorInfo  "internal error"
// @Resource todo
// @Route /api/v1/todos [get]
func (h *ToDoHandler) listToDoEndpoint(c *gin.Context) {
	todos, err := h.toDoUseCase.ListToDos()
	if err != nil {
		errordef.GinHTTPResponse(c, errordef.ErrorInfo{
			Code:    errordef.ErrCodeInternalError,
			Message: err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, listToDoResp{
		Data: todos,
	})
}

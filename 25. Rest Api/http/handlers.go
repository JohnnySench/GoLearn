package http

import (
	"encoding/json"
	"errors"
	"net/http"
	"restapi/todo"
	"time"

	"github.com/gorilla/mux"
)

type HTTPHandlers struct {
	todoList *todo.List
}

func NewHTTPHandlers(todoList *todo.List) *HTTPHandlers {
	return &HTTPHandlers{
		todoList: todoList,
	}
}

/*
pattern: /tasks
method: POST
info: JSON ih HTTP request body
succeed:
  - status code: 201 Created
  - response body: return JSON created task

failed:
  - status code: 400, 409, 500...
  - response body: JSON with error + time
*/
func (h *HTTPHandlers) HandleCreateTask(w http.ResponseWriter, r *http.Request) {
	var taskDTO TaskDTO

	if err := json.NewDecoder(r.Body).Decode(&taskDTO); err != nil {
		errorDTO := ErrorDTO{
			Message: err.Error(),
			Time:    time.Now(),
		}
		http.Error(w, errorDTO.ToString(), http.StatusBadRequest)
		return
	}

	if err := taskDTO.ValidateForCreate(); err != nil {
		errorDTO := ErrorDTO{
			Message: err.Error(),
			Time:    time.Now(),
		}
		http.Error(w, errorDTO.ToString(), http.StatusBadRequest)
		return
	}

	task := todo.NewTask(taskDTO.Title, taskDTO.Description)
	if err := h.todoList.AddTask(task); err != nil {
		errorDTO := ErrorDTO{
			Message: err.Error(),
			Time:    time.Now(),
		}
		if errors.Is(err, todo.ErrTaskAlreadyExists) {
			http.Error(w, errorDTO.ToString(), http.StatusConflict)
		} else {
			http.Error(w, errorDTO.ToString(), http.StatusInternalServerError)
		}
		return
	}

	b, err := json.MarshalIndent(task, "", "    ")
	if err != nil {
		panic(err)
	}
	w.WriteHeader(http.StatusCreated)
	if _, err := w.Write(b); err != nil {
		println("Error:", err.Error())
		return
	}
}

/*
pattern: /tasks/{title}
method: DELETE
info: in pattern

succeed:
  - status code: 204 No content
  - response body: EMPTY

failed:
  - status code: 400, 404, 500...
  - response body: JSON with errr + time
*/
func (h *HTTPHandlers) HandleDeleteTask(w http.ResponseWriter, r *http.Request) {
	title := mux.Vars(r)["title"]
	if err := h.todoList.DeleteTask(title); err != nil {
		errDTO := ErrorDTO{
			Message: err.Error(),
			Time:    time.Now(),
		}
		if errors.Is(err, todo.ErrTaskNotFound) {
			http.Error(w, errDTO.ToString(), http.StatusNotFound)
		} else {
			http.Error(w, errDTO.ToString(), http.StatusInternalServerError)
		}
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

/*
pattern: /tasks/{title}
method: PATCH
info: in pattern + JSON body request

succeed:
  - status code: 200 Ok
  - response body: JSON completed task

failed:
  - status code: 400, 409, 500...
  - response body: JSON with errr + time
*/
func (h *HTTPHandlers) HandleCompleteTask(w http.ResponseWriter, r *http.Request) {
	var completeTask CompleteTaskDTO

	if err := json.NewDecoder(r.Body).Decode(&completeTask); err != nil {
		errDTO := ErrorDTO{
			Message: err.Error(),
			Time:    time.Now(),
		}
		http.Error(w, errDTO.ToString(), http.StatusBadRequest)
	}

	title := mux.Vars(r)["title"]

	if completeTask.Complete {
		task, err := h.todoList.CompleteTask(title)
		if err != nil {
			errDTO := ErrorDTO{
				Message: err.Error(),
				Time:    time.Now(),
			}
			if errors.Is(err, todo.ErrTaskNotFound) {
				http.Error(w, errDTO.ToString(), http.StatusNotFound)
			} else {
				http.Error(w, errDTO.ToString(), http.StatusInternalServerError)
			}
			return
		}
		b, err := json.MarshalIndent(task, "", "    ")
		if err != nil {
			panic(err)
		}
		w.WriteHeader(http.StatusOK)
		w.Write(b)
	} else {
		task, err := h.todoList.UncompleteTask(title)
		if err != nil {
			errDTO := ErrorDTO{
				Message: err.Error(),
				Time:    time.Now(),
			}
			if errors.Is(err, todo.ErrTaskNotFound) {
				http.Error(w, errDTO.ToString(), http.StatusNotFound)
			} else {
				http.Error(w, errDTO.ToString(), http.StatusInternalServerError)
			}
			return
		}
		b, err := json.MarshalIndent(task, "", "    ")
		if err != nil {
			panic(err)
		}
		w.WriteHeader(http.StatusOK)
		w.Write(b)
	}

}

/*
pattern: /tasks/{title}
method: GET
info: in pattern

succeed:
  - status code: 200 Ok
  - response body: JSON found task

failed:
  - status code: 400, 404, 500...
  - response body: JSON with error + time
*/
func (h *HTTPHandlers) HandleGetTask(w http.ResponseWriter, r *http.Request) {
	title := mux.Vars(r)["title"]
	task, err := h.todoList.GetTask(title)
	if err != nil {
		errDTO := ErrorDTO{
			Message: err.Error(),
			Time:    time.Now(),
		}
		if errors.Is(err, todo.ErrTaskNotFound) {
			http.Error(w, errDTO.ToString(), http.StatusNotFound)
		} else {
			http.Error(w, errDTO.ToString(), http.StatusInternalServerError)
		}
		return
	}

	b, err := json.MarshalIndent(task, "", "    ")
	if err != nil {
		panic(err)
	}
	w.WriteHeader(http.StatusOK)
	w.Write(b)
}

/*
pattern: /tasks
method: GET
info: -

succeed:
  - status code: 200 Ok
  - response body: JSON found tasks

failed:
  - status code: 400, 404, 500...
  - response body: JSON with error + time
*/
func (h *HTTPHandlers) HandleGetAllTasks(w http.ResponseWriter, r *http.Request) {
	tasks := h.todoList.ListTasks()
	b, err := json.MarshalIndent(tasks, "", "    ")
	if err != nil {
		panic(err)
	}
	w.WriteHeader(http.StatusOK)
	w.Write(b)
}

/*
pattern: /tasks?completed=false
method: GET
info: query params

succeed:
  - status code: 200 Ok
  - response body: JSON found uncompleted tasks

failed:
  - status code: 400, 404, 500...
  - response body: JSON with error + time
*/
func (h *HTTPHandlers) HandleGetAllUncompletedTasks(w http.ResponseWriter, r *http.Request) {
	tasks := h.todoList.ListUncompletedTasks()
	b, err := json.MarshalIndent(tasks, "", "    ")
	if err != nil {
		panic(err)
	}
	w.WriteHeader(http.StatusOK)
	w.Write(b)
}

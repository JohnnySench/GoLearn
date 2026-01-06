package http

import (
	"net/http"

	"github.com/gorilla/mux"
)

type HTTPServer struct {
	httpHandlers *HTTPHandlers
}

func NewHTTPServer(h *HTTPHandlers) *HTTPServer {
	return &HTTPServer{
		httpHandlers: h,
	}
}

func (h *HTTPServer) StartServer() error {
	router := mux.NewRouter()

	// HandleCreateTask
	router.Path("/tasks").Methods(http.MethodPost).HandlerFunc(h.httpHandlers.HandleCreateTask)

	// HandleDeleteTask
	router.Path("/tasks/{title}").Methods(http.MethodDelete).HandlerFunc(h.httpHandlers.HandleDeleteTask)

	// HandleCompleteTask
	router.Path("/tasks/{title}").Methods(http.MethodPatch).HandlerFunc(h.httpHandlers.HandleCompleteTask)

	// HandleGetTask
	router.Path("/tasks/{title}").Methods(http.MethodGet).HandlerFunc(h.httpHandlers.HandleGetTask)

	// HandleGetAllTasks
	router.Path("/tasks").Methods(http.MethodGet).HandlerFunc(h.httpHandlers.HandleGetAllTasks)

	// HandleGetAllUncompletedTasks
	router.Path("/tasks").Methods(http.MethodGet).Queries("completed", "false").HandlerFunc(h.httpHandlers.HandleGetAllUncompletedTasks)

	return http.ListenAndServe(":9091", router)
}

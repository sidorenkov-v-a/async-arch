package api

import (
	"encoding/json"
	"net/http"

	"github.com/google/uuid"

	api_client "async-arch/billing/api/generated"
	"async-arch/billing/internal/pkg/domain"
	"async-arch/billing/internal/pkg/usecase/create_task"
	"async-arch/billing/internal/pkg/usecase/reassign_tasks"
)

type server struct {
	createTaskUsecase    create_task.Usecase
	reassignTasksUsecase reassign_tasks.Usecase
}

func NewServer(createTaskUsecase create_task.Usecase, reassignTasksUsecase reassign_tasks.Usecase) *server {
	return &server{createTaskUsecase: createTaskUsecase, reassignTasksUsecase: reassignTasksUsecase}
}

func badRequest(w http.ResponseWriter, err error) {
	w.WriteHeader(http.StatusBadRequest)

	out := api_client.Error{
		Message: err.Error(),
	}

	_ = json.NewEncoder(w).Encode(out)
}

func internalError(w http.ResponseWriter, err error) {
	w.WriteHeader(http.StatusInternalServerError)

	out := api_client.Error{
		Message: err.Error(),
	}

	_ = json.NewEncoder(w).Encode(out)
}

func (s *server) CreateTask(w http.ResponseWriter, r *http.Request) {
	in := api_client.TaskCreate{}

	err := json.NewDecoder(r.Body).Decode(&in)
	if err != nil {
		internalError(w, err)
		return
	}

	reporterID, err := uuid.Parse(r.Header.Get("userID"))
	if err != nil {
		badRequest(w, err)
		return
	}

	task, err := s.createTaskUsecase.Run(r.Context(), create_task.In{
		ReporterID:  reporterID,
		AssigneeID:  in.AssigneeID,
		Title:       in.Title,
		Description: in.Description,
		Status:      domain.TaskStatusNew,
	})
	if err != nil {
		badRequest(w, err)
		return
	}

	out := api_client.Task{
		AssigneeID: task.AssigneeID,
		CreatedAt:  task.CreatedAt,
		JiraID:     int(task.JiraID),
		ReporterID: task.ReporterID,
		Status:     string(task.Status),
		Title:      task.Title,
		UpdatedAt:  task.UpdatedAt,
	}

	err = json.NewEncoder(w).Encode(out)
	if err != nil {
		internalError(w, err)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (s *server) ReassignTasks(w http.ResponseWriter, r *http.Request) {
	err := s.reassignTasksUsecase.Run(r.Context())
	if err != nil {
		badRequest(w, err)
		return
	}

	out := api_client.Ok{
		Status: "Ok",
	}

	err = json.NewEncoder(w).Encode(out)
	if err != nil {
		internalError(w, err)
		return
	}

	w.WriteHeader(http.StatusOK)
}

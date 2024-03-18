// Package api_client provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen/v2 version v2.1.0 DO NOT EDIT.
package api_client

import (
	"time"

	openapi_types "github.com/oapi-codegen/runtime/types"
)

// Error defines model for Error.
type Error struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

// Ok defines model for Ok.
type Ok struct {
	Status string `json:"status"`
}

// Task defines model for Task.
type Task struct {
	AssigneeID openapi_types.UUID `json:"assigneeID"`
	CreatedAt  time.Time          `json:"createdAt"`
	JiraID     int                `json:"jiraID"`
	ReporterID openapi_types.UUID `json:"reporterID"`
	Status     string             `json:"status"`
	Title      string             `json:"title"`
	UpdatedAt  time.Time          `json:"updatedAt"`
}

// TaskCreate defines model for TaskCreate.
type TaskCreate struct {
	AssigneeID  openapi_types.UUID `json:"assigneeID"`
	Description string             `json:"description"`
	Title       string             `json:"title"`
}

// ReassignTasksJSONBody defines parameters for ReassignTasks.
type ReassignTasksJSONBody interface{}

// CreateTaskJSONRequestBody defines body for CreateTask for application/json ContentType.
type CreateTaskJSONRequestBody = TaskCreate

// ReassignTasksJSONRequestBody defines body for ReassignTasks for application/json ContentType.
type ReassignTasksJSONRequestBody ReassignTasksJSONBody

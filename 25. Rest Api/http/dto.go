package http

import (
	"encoding/json"
	"errors"
	"time"
)

type TaskDTO struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}

type TaskGetDTO struct {
	Title string `json:"title"`
}

type CompleteTaskDTO struct {
	Complete bool `json:"complete"`
}

type ErrorDTO struct {
	Message string
	Time    time.Time
}

func (e ErrorDTO) ToString() string {
	b, err := json.MarshalIndent(e, "", "    ")
	if err != nil {
		panic(err)
	}
	return string(b)
}

func (t TaskDTO) ValidateForCreate() error {
	if t.Description == "" {
		return errors.New("field description is empty")
	}
	if t.Title == "" {
		return errors.New("field title is empty")
	}
	return nil
}

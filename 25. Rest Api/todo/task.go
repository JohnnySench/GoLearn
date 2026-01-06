package todo

import "time"

type Task struct {
	Title       string
	Description string
	Completed   bool

	CreatedAt   time.Time
	CompteledAt *time.Time
}

func NewTask(title string, description string) Task {
	return Task{
		Title:       title,
		Description: description,
		Completed:   false,
		CreatedAt:   time.Now(),
		CompteledAt: nil,
	}
}

func (t *Task) Complete() {
	doneTime := time.Now()
	t.Completed = true
	t.CompteledAt = &doneTime
}

func (t *Task) Uncomplete() {
	t.Completed = false
	t.CompteledAt = nil
}

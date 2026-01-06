package todo

/*
	can call: repository
*/

type List struct {
	tasks map[string]Task
}

func NewList() *List {
	return &List{
		tasks: make(map[string]Task),
	}
}

func (l *List) AddTask(task Task) error {
	if _, ok := l.tasks[task.Title]; ok {
		return ErrTaskAlreadyExists
	}
	l.tasks[task.Title] = task
	return nil
}

func (l *List) ListTasks() map[string]Task {
	tmp := make(map[string]Task, len(l.tasks))

	for k, v := range l.tasks {
		tmp[k] = v
	}
	return tmp
}

func (l *List) ListUncompletedTasks() map[string]Task {
	tmp := make(map[string]Task)

	for k, v := range l.tasks {
		if !v.Completed {
			tmp[k] = v
		}
	}

	return tmp
}

func (l *List) CompleteTask(title string) (Task, error) {
	task, ok := l.tasks[title]
	if !ok {
		return Task{}, ErrTaskNotFound
	}
	task.Complete()

	l.tasks[title] = task

	return task, nil
}

func (l *List) UncompleteTask(title string) (Task, error) {
	task, ok := l.tasks[title]
	if !ok {
		return Task{}, ErrTaskNotFound
	}
	task.Uncomplete()

	l.tasks[title] = task

	return task, nil
}

func (l *List) DeleteTask(title string) error {
	_, ok := l.tasks[title]
	if !ok {
		return ErrTaskNotFound
	}
	delete(l.tasks, title)

	return nil
}

func (l *List) GetTask(title string) (Task, error) {
	task, ok := l.tasks[title]
	if !ok {
		return Task{}, ErrTaskNotFound
	}
	return task, nil
}

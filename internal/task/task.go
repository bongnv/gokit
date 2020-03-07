package task

// Task presents a task. It has Do function to execute the task.
type Task interface {
	Do() error
}

// Group is a group of tasks. Sub-tasks will be executed sequentially.
type Group []Task

// Do executes sub-tasks sequentially. It returns error if one sub-task returns error.
func (tasks Group) Do() error {
	for _, t := range tasks {
		if err := t.Do(); err != nil {
			return err
		}
	}

	return nil
}

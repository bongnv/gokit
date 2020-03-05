package command

type task interface {
	do() error
}

type taskGroup []task

func (tasks taskGroup) do() error {
	for _, t := range tasks {
		if err := t.do(); err != nil {
			return err
		}
	}

	return nil
}

package taskhandler

/*
TODO: Task might take a while.
It should be possible for client to send task and get it status until it is finished
*/

type task interface {
	Work(int) int
}

type taskHolder struct {
	work   task
	status int
}

type Taskhandler struct {
	tasks []taskHolder
}

func (h *Taskhandler) AddWork(n task) (int, error) {
	temp := new(taskHolder)
	temp.work = n

	h.tasks = append(h.tasks, *temp)

	return len(h.tasks), nil
}

func (h *Taskhandler) GetWorks() (int, error) {
	return 1, nil
}

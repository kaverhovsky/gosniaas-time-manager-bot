package scheduler

type scheduler struct {
}

var sch *scheduler

func Get() Scheduler {
	if sch == nil {
		sch = create()
	}
	return sch
}

func create() *scheduler {
	return &scheduler{}
}

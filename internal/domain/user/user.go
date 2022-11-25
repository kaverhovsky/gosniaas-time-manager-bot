package user

type WorkSchedule string

func (ws WorkSchedule) String() string {
	return string(ws)
}

const (
	FullSchedule WorkSchedule = "full"
	HalfSchedule              = "half"
)

type User struct {
	ID        int64
	Username  string
	Firstname string
	Lastname  string
	Schedule  WorkSchedule
}

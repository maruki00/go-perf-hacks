package boxinginterface


type Worker interface {
    Work()
}

type LargeJob struct {
    payload [4096]byte
}

func NewLargeJob() *LargeJob{
	return &LargeJob{
		payload: [4096]byte{},
	}
}

func (LargeJob) Work() {}

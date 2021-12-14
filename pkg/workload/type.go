package workload

type (
	JobType uint8
)

const (
	JobLoad JobType = 1 + iota
	JobRun
)

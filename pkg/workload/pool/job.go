package pool

// Job defined an interface to allow worker pools to perform different kinds of
// work so long as the passed struct conforms to the Job interface
type Job interface {
	Execute()
}

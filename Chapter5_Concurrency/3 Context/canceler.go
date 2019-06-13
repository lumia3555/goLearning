type canceler interface {
	cancel(removeFromParent bool, err error)
	Done()
}
package syncs

// Limiter 等待器.
type Limiter interface {
	Wait()
	Try() error
}

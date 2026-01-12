package listing

import (
	"errors"
	"fmt"
)

var ErrNotRunning = errors.New("not running")

var ErrFailedToStop = errors.New("failed to stop")

func (s *Service) StopFeed() error {
	s.mu.Lock()
	defer s.mu.Unlock()

	if !s.started {
		return ErrNotRunning
	}

	if s.cancelFunc == nil {
		return fmt.Errorf("%w cancel func is undefined", ErrFailedToStop)
	}
	s.cancelFunc()
	s.started = false
	return nil
}

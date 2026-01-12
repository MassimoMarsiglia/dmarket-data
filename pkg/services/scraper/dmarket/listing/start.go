package listing

import (
	"context"
	"errors"
	"fmt"
)

var ErrRunning = errors.New("already running")

var ErrFailedToStart = errors.New("failed to start")

func (s *Service) StartFeed() error {
	s.mu.Lock()
	defer s.mu.Unlock()

	if s.started {
		return ErrRunning
	}
	ctx, cancel := context.WithCancel(s.ctx)
	s.cancelFunc = cancel
	_, err := s.Pool.Pool.Spawn(ctx, s.numWorkers)
	if err != nil {
		cancel()
		return err
	}

	err = s.Pool.Pool.Start(ctx, s.startDelay)
	if err != nil {
		cancel()
		return fmt.Errorf("%w %v", ErrFailedToStart, err)
	}
	sub := s.Pool.Pool.Publisher.Subscribe()
	go func() {
		defer sub.Close()

		for {
			select {
			case resp, ok := <-sub.C:
				if !ok {
					return
				}
				s.publisher.Notify(resp)

			case err, ok := <-sub.E:
				if !ok {
					return
				}
				s.publisher.NotifyErr(err)

			case <-ctx.Done():
				return
			}
		}
	}()
	s.started = true
	return nil
}

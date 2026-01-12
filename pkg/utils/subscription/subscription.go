package subscription

import "sync"

type (
	Subscribable[T any] interface {
		Subscribe() Subscription[T]
		Notify(T)
		NotifyErr(error)
	}

	Subscription[T any] struct {
		C     <-chan T
		E     <-chan error
		close func()
	}

	Publisher[T any] struct {
		mu   sync.Mutex
		subs map[uint64]subscriber[T]
		next uint64
	}

	subscriber[T any] struct {
		c chan T
		e chan error
	}
)

func NewPublisher[T any]() *Publisher[T] {
	return &Publisher[T]{
		mu:   sync.Mutex{},
		next: 0,
		subs: make(map[uint64]subscriber[T]),
	}
}

func (s Subscription[T]) Close() {
	s.close()
}

var _ Subscribable[any] = (*Publisher[any])(nil)

func (p *Publisher[T]) Notify(v T) {
	p.mu.Lock()
	defer p.mu.Unlock()

	for _, sub := range p.subs {
		sub.c <- v
	}
}

func (p *Publisher[T]) NotifyErr(err error) {
	p.mu.Lock()
	defer p.mu.Unlock()

	for _, sub := range p.subs {
		sub.e <- err
	}
}

func (p *Publisher[T]) Subscribe() Subscription[T] {
	p.mu.Lock()
	defer p.mu.Unlock()

	id := p.next
	p.next++

	c := make(chan T, 100)
	e := make(chan error, 100)

	p.subs[id] = subscriber[T]{c: c, e: e}

	return Subscription[T]{
		C: c,
		E: e,
		close: func() {
			p.mu.Lock()
			defer p.mu.Unlock()

			if sub, ok := p.subs[id]; ok {
				close(sub.c)
				close(sub.e)
				delete(p.subs, id)
			}
		},
	}
}
